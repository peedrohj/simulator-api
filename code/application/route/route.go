package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string      `json:"id"`
	ClientID  string      `json:"clientId"`
	Positions []Positions `json:"positions"`
}

type Positions struct {
	Lat  float64 `json:"latittude"`
	Long float64 `json:"longitude"`
}

type ParcialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (route *Route) LoadPositions() error {
	if route.ID == "" {
		return errors.New("Route id cannot be null")
	}

	file, err := os.Open("destinations/" + route.ID + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		route.Positions = append(route.Positions, Positions{
			Lat:  lat,
			Long: long,
		})
	}

	return nil
}

func (route *Route) ExportJsonPositions() ([]string, error) {
	var routeData ParcialRoutePosition
	var result []string
	total := len(route.Positions)

	for key, value := range route.Positions {
		routeData.ID = route.ID
		routeData.ClientID = route.ClientID
		routeData.Position = []float64{value.Lat, value.Long}
		routeData.Finished = false

		if (total - 1) == key {
			routeData.Finished = true
		}

		jsonRoute, err := json.Marshal(routeData)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}
