package service

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "github.com/bdandy/go-socks4"
	"github.com/imzoloft/soundcheck/config"
	"github.com/imzoloft/soundcheck/model"
	"github.com/imzoloft/soundcheck/util"
	"github.com/imzoloft/soundcheck/util/inout"
	"golang.org/x/net/proxy"
)

type Soundcloud struct {
	BaseURL string
}

func (s *Soundcloud) Checker(email string, proxyString string, goRoutineID int) {
	transport := s.getDefaultTransport()

	if config.ProxyType != "none" {
		proxyInfo := util.ParseProxy(proxyString)
		transport = s.getProxyTransport(proxyInfo)
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	parsedURL, err := url.Parse(s.BaseURL)
	if err != nil {
		inout.FatalError("INVALID URL", err.Error())
	}

	parsedURL.RawQuery = s.addQueryParameter(parsedURL, email)

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		inout.FatalError("INVALID REQUEST", err.Error())
	}

	s.addHeader(req, goRoutineID)

	resp, err := client.Do(req)
	if err != nil {
		if config.Debug {
			inout.PrintError("ERROR DOING REQUEST", err.Error())
		}

		s.Checker(email, util.GetRandomProxy(goRoutineID), goRoutineID)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if config.Debug {
			inout.PrintError("ERROR READING BODY", err.Error())
		}

		s.Checker(email, util.GetRandomProxy(goRoutineID), goRoutineID)
		return
	}

	err = s.verifyIfEmailInUse(body)
	if err != nil {
		if config.Debug {
			inout.PrintError("ERROR", err.Error())
		}

		if strings.Contains(err.Error(), "invalid_identifier") {
			return
		}

		if strings.Contains(err.Error(), "Limited") {
			time.Sleep(30 * time.Second)
		}

		s.Checker(email, util.GetRandomProxy(goRoutineID), goRoutineID)
		return
	}
}

func (s *Soundcloud) verifyIfEmailInUse(body []byte) error {
	var response model.Response

	if strings.Contains(string(body), "Please slow down.") {
		return fmt.Errorf("[%s!%s] Limited | Waiting 30 seconds",
			config.TextRed, config.TextReset)
	}

	err := json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Error != "" {
		return errors.New(response.Error)
	}

	fileName := inout.DisplayResult(response)
	util.WriteToFile(fileName, response)

	return nil
}

func (s *Soundcloud) getProxyTransport(parsedProxy *util.ProxyInfo) *http.Transport {
	transport := s.getDefaultTransport()

	switch config.ProxyType {
	case "https":
		proxyURL := fmt.Sprintf("http://%s:%s", parsedProxy.IP, parsedProxy.Port)

		if parsedProxy.Username != "" && parsedProxy.Password != "" {
			proxyURL = fmt.Sprintf("http://%s:%s@%s:%s", parsedProxy.Username, parsedProxy.Password, parsedProxy.IP, parsedProxy.Port)
		}
		parsedURL, _ := url.Parse(proxyURL)
		transport.Proxy = http.ProxyURL(parsedURL)

	case "socks4", "socks5":
		var dialer proxy.Dialer
		timeout := time.Duration(5 * time.Second)

		proxyURL := fmt.Sprintf("%s://%s:%s", config.ProxyType, parsedProxy.IP, parsedProxy.Port)

		if parsedProxy.Username != "" && parsedProxy.Password != "" {
			proxyURL = fmt.Sprintf("%s://%s:%s@%s:%s", config.ProxyType, parsedProxy.Username, parsedProxy.Password, parsedProxy.IP, parsedProxy.Port)
		}
		parsedURL, _ := url.Parse(proxyURL)

		switch config.ProxyType {
		case "socks4":
			dialer, _ = proxy.FromURL(parsedURL, &net.Dialer{Timeout: timeout})
		case "socks5":
			var auth *proxy.Auth

			if parsedProxy.Username != "" && parsedProxy.Password != "" {
				auth = &proxy.Auth{User: parsedProxy.Username, Password: parsedProxy.Password}
			}
			dialer, _ = proxy.SOCKS5("tcp", parsedProxy.IP+":"+parsedProxy.Port, auth, &net.Dialer{Timeout: timeout})
		}
		transport.Dial = dialer.Dial
	}
	return transport
}

func (s *Soundcloud) getDefaultTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}
