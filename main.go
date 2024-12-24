package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ubuntupunk/obsidian-cli/cmd"
	"github.com/ubuntupunk/obsidian-cli/config"
)

func main() {
	// Initialize configuration
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	// Execute root command
	if err := cmd.Execute(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
