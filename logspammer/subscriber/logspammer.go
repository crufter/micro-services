package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	logspammer "logspammer/proto/logspammer"
)

type Logspammer struct{}

func (e *Logspammer) Handle(ctx context.Context, msg *logspammer.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *logspammer.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
