package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/urhStop/vremenske-postaje/vreme"
)

func main() {
	cmd := &cli.Command{
		Name:  "weatherstation",
		Usage: "Program za delo z vremenskimi postajami",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			if minTemp > maxTemp {
				return fmt.Errorf("Najnižja temperatura ne sme biti višja od najvišje.")
			}

			if stMeritev < 1 {
				return fmt.Errorf("Število meritev mora biti vsaj 1.")
			}

			postaje := make(map[string]vreme.Postaja)

			// Dodajanje vremenskih postaj
			vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
			vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")

			// Meritve za postajo KP01
			vreme.DodajMeritev(postaje, "KP01", 24.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KP01", 27, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KP01", 26.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KP01", 28.4, minTemp, maxTemp)

			// Meritve za postajo KR01
			vreme.DodajMeritev(postaje, "KR01", 2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KR01", -1.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KR01", 3.5, minTemp, maxTemp)

			// Meritve za postajo LJ01
			vreme.DodajMeritev(postaje, "LJ01", 18.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 21.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 19.8, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 16.5, minTemp, maxTemp)

			// Meritve za postajo MB01
			vreme.DodajMeritev(postaje, "MB01", 12.4, minTemp, maxTemp)

			fmt.Println("\nVse meritve:")
			vreme.IzpisMeritev(postaje)

			fmt.Println("\nTemperaturno poročilo:")
			vreme.IzpisPorocila(postaje, stMeritev)
			fmt.Println()

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
