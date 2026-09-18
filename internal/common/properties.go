package common

import "log/slog"

var (
	SERVER_PORT     string
	SLOG_LEVEL      int
	LOGGER          *slog.Logger
	SERVICE_VERSION string
	USER_API        string
)

const (
	GetUsersPath = "/users"
)
