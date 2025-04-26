package system

import (
	"fmt"

	"github.com/vantu2801se/product-manager-system/client/rds"
	"github.com/vantu2801se/product-manager-system/config"
	"go.uber.org/zap"
)

type SystemContext struct {
	Config *config.Config
	Logger *zap.SugaredLogger
	RDSCli rds.Client
}

func NewSystemContext(config *config.Config) (*SystemContext, error) {
	logger, err := newLogger(fmt.Sprintf(config.LogFolder, config.AppName))
	if err != nil {
		return nil, fmt.Errorf("failed to init logger. %s", err.Error())
	}

	rdsCli, err := rds.NewRDSClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to init rds client. %s", err.Error())
	}

	return &SystemContext{
		Config: config,
		Logger: logger,
		RDSCli: rdsCli,
	}, nil
}
