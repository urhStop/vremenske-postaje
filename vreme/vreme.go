// Package vreme omogoča delo z vremenskimi postajami in njihovimi meritvami.
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z njenimi podatki in meritvami temperature.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajPostajo doda novo vremensko postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda temperaturno meritev izbrani postaji, če je temperatura znotraj dovoljenih mej.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	el, ok := postaje[oznaka]
	if ok {
		if temperatura >= minTemp && temperatura <= maxTemp {
			el.meritve = append(el.meritve, temperatura)
			postaje[oznaka] = el
		} else {
			fmt.Printf("Neveljavna meritev za postajo %s: %.1f °C (dovoljen razpon: %.1f do %.1f °C).\n", oznaka, temperatura, minTemp, maxTemp)
		}
	} else {
		fmt.Println("Postaja ne obstaja.")
	}
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, ok := postaje[oznaka]

	if !ok || len(p.meritve) == 0 || len(p.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = p.meritve[0]
	najvisja = p.meritve[0]
	vsota := 0.0

	for _, val := range p.meritve {
		if val < najnizja {
			najnizja = val
		}
		if val > najvisja {
			najvisja = val
		}
		vsota += val
	}

	povprecna = vsota / float64(len(p.meritve))

	return najnizja, najvisja, povprecna, true
}

// IzpisMeritev izpiše vse vremenske postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	for k, v := range postaje {
		fmt.Printf("%s - %s (%s): %v \n", k, v.ime, v.kraj, v.meritve)
	}
}

// IzpisPorocila izpiše temperaturno poročilo za vse postaje, ki imajo dovolj meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for k, v := range postaje {
		min, max, povp, ok := statistika(postaje, k, stMeritev)

		if ok {
			opis := ""
			switch {
			case povp >= 25:
				opis = "Vroče"
			case povp >= 15 && povp < 25:
				opis = "Toplo"
			case povp >= 5 && povp < 15:
				opis = "Hladno"
			case povp < 5:
				opis = "Mrzlo"
			}
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", v.ime, min, max, povp, opis)
		} else {
			fmt.Printf("%s: premalo meritev\n", v.ime)
		}
	}
}
