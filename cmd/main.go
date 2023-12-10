package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/imzoloft/soundcheck/config"
	"github.com/imzoloft/soundcheck/service"
	"github.com/imzoloft/soundcheck/util"
	"github.com/imzoloft/soundcheck/util/inout"
)

func main() {
	inout.ClearScreen()
	inout.SelectProxyType()
	routinesNumber := inout.SelectRoutineNumber()
	time.Sleep(2 * time.Second)

	inout.ClearScreen()
	inout.ShowBanner()

	sc := &service.Soundcloud{
		BaseURL: "https://api-auth.soundcloud.com/web-auth/identifier",
	}

	emailsFile := util.ReadFile("list/emails.txt")
	emails := strings.Split(string(emailsFile), "\n")

	if len(emails) == 1 {
		inout.FatalError("NO EMAILS FOUND", "No emails were loaded from the file")
	}

	if config.ProxyType != "none" {
		util.LoadProxies()
	}

	var wg sync.WaitGroup

	emailChan := make(chan string)

	for i := 0; i < routinesNumber; i++ {
		wg.Add(1)

		go func(goRoutineID int) {
			defer wg.Done()
			for email := range emailChan {
				sc.Checker(email, util.GetRandomProxy(goRoutineID), goRoutineID)
			}
		}(i)
	}

	for _, email := range emails {
		emailChan <- email
	}
	close(emailChan)
	wg.Wait()
	fmt.Printf("[%s!%s] %sFinished checking %d emails%s\n", config.TextGreen, config.TextReset, config.TextGreen, len(emails), config.TextReset)
}
