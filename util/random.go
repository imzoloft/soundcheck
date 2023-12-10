package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/imzoloft/soundcheck/config"
)

func GetRandomNumber(goRoutineID int, max int) int {
	s := rand.NewSource(time.Now().Unix() + int64(goRoutineID))
	r := rand.New(s)
	return r.Intn(max)
}

func GetRandomProxy(goRoutineID int) string {
	if config.ProxyType == "none" {
		return ""
	}

	time.Sleep(500 * time.Millisecond)
	proxyCount := len(config.Proxies)

	if proxyCount == 0 {
		return ""
	} else if proxyCount == 1 {
		return strings.TrimSpace(
			strings.TrimRight(config.Proxies[0], "\r"),
		)
	}
	return strings.TrimSpace(
		strings.TrimRight(config.Proxies[GetRandomNumber(goRoutineID, proxyCount-1)], "\r"),
	)
}
