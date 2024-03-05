package main

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"github.com/lrmnt/AA6_homework/tasks/ent"
	"github.com/lrmnt/AA6_homework/tasks/internal/kafka/consumer"
	"github.com/lrmnt/AA6_homework/tasks/internal/kafka/producer"
	"github.com/lrmnt/AA6_homework/tasks/internal/server"
	"github.com/lrmnt/AA6_homework/tasks/internal/service"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can not init logger: %v", err)
	}

	client, err := ent.Open("postgres", "sslmode=disable host=localhost port=5432 user=postgres dbname=tasks password=password")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasksProducer, err := producer.New(ctx, "localhost:9092", "tasks")
	if err != nil {
		l.Fatal("can not init kafka producer", zap.Error(err))
	}

	srv := server.New(client, "http://localhost:8091", ":8092", l, tasksProducer)

	userConsumer := consumer.NewReader([]string{"localhost:9092"}, "users")
	loader := service.New(l, client, userConsumer)

	eg, ctx := errgroup.WithContext(ctx)

	// run server
	eg.Go(func() error {
		err = srv.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			l.Error("error on running server", zap.Error(err))
			return err
		}

		l.Debug("http server stopped")

		return nil
	})

	// run load users
	eg.Go(func() error {
		loader.Run(ctx)

		l.Debug("kafka loader stopped")

		return nil
	})

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)

	signal.Notify(interrupt, os.Interrupt)
	defer signal.Stop(interrupt)

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	termCtx, termCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer termCancel()

	err = srv.Stop(termCtx)
	if err != nil {
		l.Error("error on stopping server", zap.Error(err))
	}

	cancel()
	_ = eg.Wait()
}
