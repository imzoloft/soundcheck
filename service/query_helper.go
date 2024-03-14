package service

import (
	"net/http"
	"net/url"

	"github.com/imzoloft/soundcheck/config"
	"github.com/imzoloft/soundcheck/util"
)

func (s *Soundcloud) addHeader(req *http.Request, goRoutineID int) {
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
		"Origin":       {"https://secure.soundcloud.com"},
		"Referer":      {"https://secure.soundcloud.com/"},
		"User-Agent":   {config.UserAgents[util.GetRandomNumber(goRoutineID, len(config.UserAgents)-1)]},
	}
}

func (s *Soundcloud) addQueryParameter(url *url.URL, email string) string {
	query := url.Query()
	query.Set("q", email)
	query.Set("client_id", "rVtnkH7kI646kRDwGSONEc7euMBMyJwv")

	return query.Encode()
}
