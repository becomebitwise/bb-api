package main

import (
	"log"
	"os"
	"strconv"

	"github.com/becomebitwise/bb-api/http"
)

func main() {
	host := getEnvString("BB_HOST", "localhost")
	port := getEnvUint("BB_PORT", 1337)
	if err := http.Serve(host, port); err != nil {
		log.Fatalf("HTTP Server failed: %s", err)
	}
}

func getEnvString(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}

func getEnvUint(key string, def uint) uint {
	val, err := strconv.ParseUint(os.Getenv(key), 10, 32)
	if err != nil {
		return def
	}

	return uint(val)
}
