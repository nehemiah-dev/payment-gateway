package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/nehemiah-dev/payment-gateway/internal/logger"
	"github.com/nehemiah-dev/payment-gateway/internal/router"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "application error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger.Init()
	defer logger.Sync()

	r := router.New()
	if err := r.Run(":8443"); err != nil {
		zap.L().Error("failed to run server", zap.Error(err))
		return err
	}

	return nil
}
