package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	config := loadConfig("config/config.json")
	for site, siteConfig := range config {
		fmt.Printf("Starting to scrape %s\n\n", site)
		fmt.Println(siteConfig)
	}
}

func loadConfig(filePath string) map[string]interface{} {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	var config map[string]interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding config JSON: %v", err)
	}
	return config
}
