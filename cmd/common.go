package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func handleErrors(hasError bool) {
	warn := viper.GetString("warn")
	if hasError && warn != "true" {
		os.Exit(1)
	}
	if hasError && warn == "true" {
		log.Println("Warning: 1 or more errors occurred but --warn-only is set.")
	}

}
