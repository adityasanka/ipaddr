package ipaddr

import (
	"context"
	"errors"
	"fmt"

	"github.com/biter777/countries"
	"github.com/carlmjohnson/requests"
)

var ErrServerConnection = errors.New("Unable to connect to server. Please check your internet connection.")

var ipinfoUrl = "https://ipinfo.io"
var Token string

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func GetIPInfo() (IPInfo, error) {
	var resp IPInfo

	err := requests.
		URL(ipinfoUrl).
		Param("token", Token).
		ToJSON(&resp).
		Fetch(context.Background())

	if err != nil {
		if errors.Is(err, requests.ErrTransport) {
			return IPInfo{}, ErrServerConnection
		}

		return IPInfo{}, err
	}

	resp.Country = fmtCountry(resp.Country)
	return resp, nil
}

// Format country display text
func fmtCountry(isoCode string) string {
	return fmt.Sprintf("%s (%s)", countries.ByName(isoCode), isoCode)
}
