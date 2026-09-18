package config

import (
	"github.com/MakMoinee/go-boiler-plate/internal/common"
	"github.com/MakMoinee/go-mith/pkg/mithconfigs"
)

func Set() {
	mithconfigs.Set()
	common.SERVER_PORT = mithconfigs.Registry.GetString("SERVER_PORT")
	common.SLOG_LEVEL = mithconfigs.Registry.GetInt("SLOG_LEVEL")
	common.SERVICE_VERSION = mithconfigs.Registry.GetString("SERVICE_VERSION")
	common.USER_API = mithconfigs.Registry.GetString("USER_API")
}
