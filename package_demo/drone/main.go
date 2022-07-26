package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	body := strings.NewReader(
		os.Getenv("PLUGIN_BODY"),
	)

	_, err := http.NewRequest(
		os.Getenv("PLUGIN_METHOD"),
		os.Getenv("PLUGIN_URL"),
		body,
	)
	if err != nil {
		os.Exit(1)
	}
}
