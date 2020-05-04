package infrastructure

import (
	"fmt"
	"os"
	"strings"
)

// AppConfigLoader can load the configuration for this application
type AppConfigLoader struct {
	Logger *Logger
}

// LoadFromEnv loads the configuration from environment variables
func (cl *AppConfigLoader) LoadFromEnv() map[string]string {
	// Setup default config
	config := map[string]string{
		"PORT":           "3000",
		"REDIS_ADDR":     "localhost:6379",
		"REDIS_PASSWORD": "",
		"REDIS_TLS":      "disabled",
		"LOG_LEVEL":      "3",
	}

	for key, value := range config {
		if parsed := os.Getenv(key); parsed != "" {
			config[key] = parsed
		} else {
			config[key] = value
		}
	}

	cl.Logger.Info(formatConfig(&config))

	return config
}

func formatConfig(config *map[string]string) string {
	output := "app config: {"

	for key, value := range *config {
		if !strings.Contains(strings.ToUpper(key), "PASSWORD") {
			output += fmt.Sprintf(` "%s": "%s",`, key, value)
		} else {
			output += fmt.Sprintf(` "%s": "<HIDDEN>",`, key)
		}
	}

	output = output[0:len(output)-1] + "}"

	return output
}
