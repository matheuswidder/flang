package runtime

import (
	"os"
	"strings"
)

// LoadEnv reads a .env file and sets environment variables.
func LoadEnv(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil // .env is optional
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		// Remove surrounding quotes
		if len(val) >= 2 && (val[0] == '"' || val[0] == '\'') {
			val = val[1 : len(val)-1]
		}

		os.Setenv(key, val)
	}

	return nil
}

// GetEnv gets env variable with fallback.
func GetEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
