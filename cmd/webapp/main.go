package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/MakMoinee/go-boiler-plate/cmd/webapp/config"
	"github.com/MakMoinee/go-boiler-plate/internal/common"
	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {
	config.Set()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(common.SLOG_LEVEL),
	}))
	httpService := goserve.NewService(common.SERVER_PORT)
	logger.Info(fmt.Sprintf("Server Starting in Port %s", common.SERVER_PORT))
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
