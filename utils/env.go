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

func FetchEnvMap(name string) map[string]string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("%s environment variable is required", name))
	}
	pairs := strings.Split(value, ",")
	result := make(map[string]string)
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			panic(fmt.Sprintf("Invalid format for %s environment variable", name))
		}
		result[kv[0]] = kv[1]
	}
	return result
}
