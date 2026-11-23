package vreme

import (
	"errors"
	"fmt"
	"slices"
)

var (
	postaje map[string]Postaja

	minTemp float64 = -50
	maxTemp float64 = 50
	nSamples int = 3
)

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(ime string, oznaka string, kraj string) bool {
	if len(ime) < 1 || len(oznaka) < 1 || len(kraj) < 1 {
		return false
	}
	_, ok := postaje[oznaka]
	if ok {
		return false
	}

	postaja := Postaja{
		ime,
		kraj,
		[]float64{},
	}

	postaje[oznaka] = postaja
	return true
}

func Init(minTemperature float64, maxTemperature float64, numberOfSamples int) {
	minTemp = minTemperature
	maxTemp = maxTemperature
	nSamples = numberOfSamples
}

func dodajMeritev(oznaka string, temperatura float64) error {
	postaja, ok := postaje[oznaka]
	if !ok {
		return errors.New("Postaja ne obstaja")
	}

	if temperatura < minTemp || temperatura > maxTemp {
		return errors.New("Nepravilna temperatura")
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
	return nil
}

func statistika(oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]
	if !ok || len(postaja.meritve) < nSamples {
		return 0, 0, 0, false
	}

	najnizja = slices.Min(postaja.meritve)
	najvisja = slices.Max(postaja.meritve)
	sum := 0.
	for _, meritev := range postaja.meritve {
		sum += meritev
	}
	povprecna = sum / float64(len(postaja.meritve))

	return najnizja, najvisja, povprecna, true
}

func IzpisMeritev() {
	for oznaka, postaja := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, postaja.ime, postaja.kraj, postaja.meritve)
	}
}

func IzpisPorocila() {
	for oznaka, postaja := range postaje {
		min, max, avg, ok := statistika(oznaka)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", postaja.ime)
			continue
		}
		fmt.Printf("%s: min %.2f °C, max %.2f °C, povprečje %.2f °C -> ", postaja.ime, min, max, avg)
		switch {
		case avg >= 25:
			fmt.Println("Vroče")
		case avg >= 15:
			fmt.Println("Toplo")
		case avg >= 5:
			fmt.Println("Hladno")
		case avg < 5:
			fmt.Println("Mrzlo")
		}
	}
}
