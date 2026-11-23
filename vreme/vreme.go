package vreme

import (
	"errors"
	"fmt"
	"slices"
)

var postaje map[string]Postaja

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(ime string, oznaka string, kraj string) bool {
	if len(ime) < 1 || len(oznaka) < 1 || len(kraj) < 1 {
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

func dodajMeritev(oznaka string, temperatura float64) error {
	postaja, ok := postaje[oznaka]
	if !ok {
		return errors.New("Postaja ne obstaja")
	}

	if temperatura < -50 || temperatura > 50 {
		return errors.New("Nepravilna temperatura")
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
	return nil
}

func statistika(oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	postaja, ok := postaje[oznaka]
	if !ok || len(postaja.meritve) < 3 {
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
