package main

import (
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(event events.SQSEvent) error {
	slog.Info("Received event", "event", event)

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
