package main

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"github.com/lrmnt/AA6_homework/billing/ent"
	"github.com/lrmnt/AA6_homework/billing/internal/server"
	"github.com/lrmnt/AA6_homework/billing/internal/service/consumer"
	"github.com/lrmnt/AA6_homework/billing/internal/service/cron"
	"github.com/lrmnt/AA6_homework/billing/internal/service/reports"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"

	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can not init logger: %v", err)
	}

	client, err := ent.Open("postgres", "sslmode=disable host=localhost port=5432 user=postgres dbname=billing password=password")
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

	reportService := reports.New(l, client)

	srv := server.New("http://localhost:8091", ":8093", l, reportService)

	userConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "users_stream_v1", "billing")
	taskConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "tasks_stream_v1", "billing")
	taskEventConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "tasks_event_v1", "billing")
	loader := consumer.New(l, client, userConsumerV1, taskConsumerV1, taskEventConsumerV1)

	billingProducer, err := kafka.NewProducer(ctx, "localhost:9092", "billing_events_v1")
	if err != nil {
		l.Fatal("can not init kafka producer", zap.Error(err))
	}

	cronService := cron.New(l, client, billingProducer)

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
		err := loader.RunConsumeTaskMessageV1(ctx)
		l.Debug("kafka loader stopped")
		return err
	})
	eg.Go(func() error {
		err := loader.RunConsumeTaskEventV1(ctx)
		l.Debug("kafka loader stopped")
		return err
	})
	eg.Go(func() error {
		err := loader.RunConsumeUserMessageV1(ctx)
		l.Debug("kafka loader stopped")
		return err
	})

	eg.Go(func() error {
		return cronService.Run(ctx)
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
