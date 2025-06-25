package main

import (
	"log/slog"
	"match-sim/mock"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(event events.SQSEvent) error {
	slog.Info("Received event", "event", event)

	// TODO: Get match from lambda sqs event
	match := mock.Match1

	match.SimulateMatch(12345, false)

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
