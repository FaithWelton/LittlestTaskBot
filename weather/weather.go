package weather

import (
	"errors"
	"fmt"
	"os"

	owm "github.com/briandowns/openweathermap"
)

type Weather struct {
	client *owm.CurrentWeatherData
	unit   string
}

// https://core.telegram.org/bots/features#web-apps Maybe upgrade to a weather webapp in the future for this?

func New(language string) (*Weather, error) {
	var unit = "C"
	var lang = language

	w, err := owm.NewCurrent(unit, lang, os.Getenv("OPENWEATHER_APITOKEN"))
	if err != nil {
		return nil, errors.New("[WEATHER]: Location Error")
	}

	return &Weather{
		client: w,
		unit:   unit,
	}, nil
}

func (w *Weather) ChangeUnit() {
	if w.unit == "C" {
		w.unit = "F"
	} else {
		w.unit = "C"
	}
}

func (w *Weather) Get(location map[string]float64, language string) (string, error) {
	fmt.Println("\n[WEATHER]: New Location Received: ")
	fmt.Println(location)

	if len(location) == 0 {
		return fmt.Sprintln(location), errors.New("\n[WEATHER]: No Location Provided")
	}

	coords := &owm.Coordinates{
		Longitude: location["lon"],
		Latitude:  location["lat"],
	}

	w.client.CurrentByCoordinates(coords)

	temperature := fmt.Sprintf(" %.4g %s", w.client.Main.Temp, w.unit)
	feelsLike := fmt.Sprintf("%.4g %s", w.client.Main.FeelsLike, w.unit)
	humidity := fmt.Sprintf("%d%%", w.client.Main.Humidity)
	highLow := fmt.Sprintf("High of %.4g %s and Low of %.4g %s", w.client.Main.TempMax, w.unit, w.client.Main.TempMin, w.unit)
	precipitation := fmt.Sprintf("%.4g%% Chance of Precipitation in the next hour", w.client.Rain.OneH)

	weather := fmt.Sprintf("Weather Data for your location: \nCurrent Temperature: %s\nFeels Like: %s\nHumidity: %s\n%s\nand a %s\n", temperature, feelsLike, humidity, highLow, precipitation)
	return weather, nil
}
