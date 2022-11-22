package main

import (
	"context"
	"app"
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logger.Sync() }()
	a := app.App{}
	if err := a.Init(context.Background(), logger); err != nil {
		log.Fatal(err)
	}
	if err := a.Serve(); err != nil {
		log.Fatal(err)
	}
}
