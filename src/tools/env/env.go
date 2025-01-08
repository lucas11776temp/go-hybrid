package env

import (
	"os"
	"strings"
)

func Load(path string) error {
	bytes, err := os.ReadFile(path)

	if err != err {
		return err
	}

	env := strings.Split(string(bytes), "\r\n")

	for k := range env {
		item := strings.Split(env[k], "=")

		if len(item) < 2 {
			os.Setenv(strings.Trim(item[0], " "), "")
			continue
		}

		os.Setenv(strings.Trim(item[0], " "), strings.Trim(strings.Trim(item[1], " "), "\""))
	}

	return nil
}

func LoadDefault() error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	Load(cwd + "\\.env")

	return nil
}

func Get(name string) string {
	return os.Getenv(name)
}
