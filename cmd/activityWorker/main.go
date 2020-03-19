package main

import (
	"os"
	"os/signal"

	"github.com/uber-go/tally"
	"go.uber.org/zap"

	"go.temporal.io/temporal/activity"
	"go.temporal.io/temporal/client"
	"go.temporal.io/temporal/worker"

	"github.com/temporalio/temporal-go-demo/activities"
	"github.com/temporalio/temporal-go-demo/common"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap logger created")
	scope := tally.NoopScope

	// The client is a heavyweight object that should be created once per process.
	serviceClient, err := client.NewClient(client.Options{
		HostPort:     common.Host,
		DomainName:   common.Domain,
		MetricsScope: scope,
	})
	if err != nil {
		logger.Fatal("Unable to create client", zap.Error(err))
	}

	worker := worker.New(serviceClient, common.ActivityTaskList, worker.Options{
		Logger:                logger,
		DisableWorkflowWorker: true,
	})

	worker.RegisterActivityWithOptions(activities.Deposit, activity.RegisterOptions{Name: "deposit"})
	worker.RegisterActivityWithOptions(activities.Withdraw, activity.RegisterOptions{Name: "withdraw"})

	err = worker.Start()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}
	// The workers are supposed to be long running process that should not exit.
	waitCtrlC()

	// Stop worker, close connection, clean up resources.
	worker.Stop()
	_ = serviceClient.CloseConnection()
}

func waitCtrlC() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
