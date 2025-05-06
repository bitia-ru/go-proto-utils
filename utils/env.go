package utils

import (
	"fmt"
	"os"
	"strings"
)

func FetchEnvVariable(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("%s environment variable is required", name))
	}
	return value
}

func FetchEnvList(name string) []string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("%s environment variable is required", name))
	}
	return strings.Split(value, ",")
}
