package logger

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LogOutput string `yaml:"log_output"`
	Level     string `yaml:"level"`
}

func InitializeLogger() error {
	data, err := os.ReadFile("log_config.yaml")
	if err != nil {
		slog.Error("Failed to read config file", "error", err)
		return err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		slog.Error("Failed to unmarshal config", "error", err)
		return err
	}
	logConfig := &slog.HandlerOptions{}
	switch config.Level {
	case "DEBUG":
		logConfig.Level = slog.LevelDebug
	case "INFO":
		logConfig.Level = slog.LevelInfo
	case "WARN":
		logConfig.Level = slog.LevelWarn
	case "ERROR":
		logConfig.Level = slog.LevelError
	default:
		logConfig.Level = slog.LevelDebug
	}
	var logger *slog.Logger
	if config.LogOutput != "" {
		logFile, err := os.OpenFile(config.LogOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			slog.Error("Failed to open log file", "error", err)
			return err
		}
		writer := io.MultiWriter(logFile, os.Stdout)
		logger = slog.New(slog.NewJSONHandler(writer, logConfig))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, logConfig))
	}
	slog.SetDefault(logger)
	slog.Info("Logger initialized")
	return nil
}
