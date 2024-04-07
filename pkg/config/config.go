package config

import (
	"bufio"
	"os"
	"strings"
)

// LoadEnv reads a .env file and sets the environment variables
func LoadEnv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") { // skip empty strings and comments
			continue
		}

		// Split strings into 2 parts, skip those that throw errors
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value) // set the environment variables
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
