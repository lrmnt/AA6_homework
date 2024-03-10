package main

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"github.com/lrmnt/AA6_homework/auth/ent"
	"github.com/lrmnt/AA6_homework/auth/internal/server"
	service2 "github.com/lrmnt/AA6_homework/auth/internal/service"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"go.uber.org/zap"
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

	client, err := ent.Open("postgres", "sslmode=disable host=localhost port=5432 user=postgres dbname=auth password=password")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		l.Fatal("failed creating schema resources: %v", zap.Error(err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userProducer, err := kafka.NewProducer(ctx, "localhost:9092", "users_stream_v1")
	if err != nil {
		l.Fatal("can not init kafka producer", zap.Error(err))
	}

	service := service2.New(l, client, userProducer)

	srv, err := server.New("1234", ":8091", l, service)
	if err != nil {
		l.Fatal("can not init server", zap.Error(err))
	}

	// rus server
	go func() {
		err = srv.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			l.Error("error on running server", zap.Error(err))
		}

		cancel()
	}()

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

}
