package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/adityasanka/ipaddr"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	ipInfo, err := ipaddr.GetIPInfo()

	if err != nil {
		printError(err)
	} else {
		printIPInfo(ipInfo)
	}
}

func printError(err error) {
	redStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E88388"))

	fmt.Println(redStyle.Render(err.Error()))
}

func printIPInfo(ipInfo ipaddr.IPInfo) {
	var rows []string

	rValue := reflect.ValueOf(ipInfo)
	for i := 0; i < rValue.NumField(); i++ {
		key := rValue.Type().Field(i).Name
		val := rValue.Field(i).String()

		// Highlight IP address in green color
		if key == "IP" {
			greenStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A8CC8C"))

			val = greenStyle.Render(val)
		}

		rows = append(rows, fmt.Sprintf("%-10s : %s", key, val))
	}

	dialogStyle := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		Align(lipgloss.Left).
		Padding(1)

	fmt.Println(dialogStyle.Render(strings.Join(rows, "\n")))
}
