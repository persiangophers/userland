package zap_logger

import "go.uber.org/zap"

func InitLogger(path string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.OutputPaths = []string{
		path,
	}

	return cfg.Build()
}

func SyncLogger(logger *zap.Logger) {
	err := logger.Sync()
	if err != nil {
		logger.Error("Error syncing logger", zap.String("error", err.Error()))
	}
}
