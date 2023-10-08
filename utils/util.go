package utils

import (
	"log/slog"
	"math"
	"os"
)

func Round2Dec(val float64) float64 {
	return math.Round(val*100.0) / 100.0
}

func LogError(message string, err error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Error(message, "errMsg", err)
	os.Exit(1)
}
