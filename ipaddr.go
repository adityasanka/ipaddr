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
	IP       string
	City     string
	Region   string
	Country  string
	Location string
	Org      string
	Timezone string
}

func GetIPInfo() (IPInfo, error) {
	var resp getIPInfoResponse

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

	return resp.toIPInfo(), nil
}

type getIPInfoResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

func (r getIPInfoResponse) toIPInfo() IPInfo {
	myCountry := countries.ByName(r.Country)

	return IPInfo{
		IP:       r.IP,
		City:     r.City,
		Region:   r.Region,
		Country:  fmt.Sprintf("%s (%s)", myCountry, myCountry.Alpha2()),
		Location: r.Location,
		Org:      r.Org,
		Timezone: r.Timezone,
	}
}
