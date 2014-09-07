package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(".env file not found.")
	}

	syslogEnabled := os.Getenv("SYSLOG_ENABLED")
	fmt.Println("Syslog enabled:", syslogEnabled)
}
