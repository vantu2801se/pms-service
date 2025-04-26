package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vantu2801se/product-manager-system/config"
	"github.com/vantu2801se/product-manager-system/handler"
	"github.com/vantu2801se/product-manager-system/static"
	"github.com/vantu2801se/product-manager-system/system"
)

func main() {
	var exitCode = static.ExitOK
	defer exit(&exitCode)

	cfg, err := config.NewConfig("./config/config.toml")
	if err != nil {
		log.Printf("failed to read config. err: %s", err.Error())
		exitCode = static.ExitStartFailed
		return
	}

	sysCtx, err := system.NewSystemContext(cfg)
	if err != nil {
		log.Printf("failed to create system context. err: %s", err.Error())
		exitCode = static.ExitStartFailed
		return
	}

	h := handler.NewHttpHandler(sysCtx)
	h.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := h.Shutdown(ctx); err != nil {
		sysCtx.Logger.Errorf("failed to graceful shutdown: %v", err)
		exitCode = static.ExitError
		return
	}

	os.Exit(static.ExitOK)
}

func exit(exitCode *int) {
	if err := recover(); err != nil {
		panicCode := static.ExitPanic
		log.Printf("panic. err: %v", err)
		exitCode = &panicCode
	}
	os.Exit(*exitCode)
}
