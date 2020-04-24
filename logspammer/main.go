package main

import (
	"logspammer/handler"
	"logspammer/subscriber"
	"time"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	logspammer "logspammer/proto/logspammer"
)

func init() {
	go func() {
		for {
			log.Info("These logs will happen until you stop me! Never stop never stopping!")
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.logspammer"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	logspammer.RegisterLogspammerHandler(service.Server(), new(handler.Logspammer))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.logspammer", service.Server(), new(subscriber.Logspammer))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
