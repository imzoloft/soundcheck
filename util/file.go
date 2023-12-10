package util

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/imzoloft/soundcheck/config"
	"github.com/imzoloft/soundcheck/model"
	"github.com/imzoloft/soundcheck/util/inout"
)

var fileMutex sync.Mutex

func WriteToFile(fileName string, response model.Response) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString(response.Identifier + "\n"); err != nil {
		inout.FatalError("Error writing to file", err.Error())
	}
}

func ReadFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		inout.FatalError("ERROR READING FILE", err.Error())
	}

	return file
}

func LoadProxies() {
	proxiesFile, err := os.ReadFile("list/proxies.txt")

	if err != nil {
		inout.FatalError("ERROR READING PROXIES FILE", err.Error())
	}

	config.Proxies = strings.Split(string(proxiesFile), "\n")

	if len(config.Proxies) == 0 {
		inout.FatalError("NO PROXIES FOUND", "No proxies were loaded from the file")
	}
}
