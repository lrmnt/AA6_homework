package main

import (
	"context"
	"errors"
	_ "github.com/lib/pq"
	"github.com/lrmnt/AA6_homework/analytics/ent"
	"github.com/lrmnt/AA6_homework/analytics/internal/server"
	"github.com/lrmnt/AA6_homework/analytics/internal/service/consumer"
	"github.com/lrmnt/AA6_homework/analytics/internal/service/reports"
	"github.com/lrmnt/AA6_homework/lib/kafka"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"log"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can not init logger: %v", err)
	}

	client, err := ent.Open("postgres", "sslmode=disable host=localhost port=5432 user=postgres dbname=analytics password=password")
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

	reportService := reports.New(l, client)

	srv := server.New("http://localhost:8091", ":8094", l, reportService)

	userConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "users_stream_v1", "analytics")
	taskConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "tasks_stream_v1", "analytics")
	billingEventConsumerV1 := kafka.NewReader([]string{"localhost:9092"}, "billing_events_v1", "analytics")
	loader := consumer.New(l, client, userConsumerV1, taskConsumerV1, billingEventConsumerV1)

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
		err := loader.RunConsumeBillingEventV1(ctx)
		l.Debug("kafka loader stopped")
		return err
	})
	eg.Go(func() error {
		err := loader.RunConsumeUserMessageV1(ctx)
		l.Debug("kafka loader stopped")
		return err
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
