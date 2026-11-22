# Vremenske postaje (weatherstation)

Preprost aplikacija, napisana v programskem jeziku Go, za delo z vremenskimi postajami in njihovimi temperaturnimi meritvami.

Omogoča dodajanje vremenskih postaj, shranjevanje temperaturnih meritev, preverjanje veljavnosti meritev ter izpis osnovnega temperaturnega poročila.

## Struktura projekta

Projekt je razdeljen na dva glavna dela:

- **`vreme`** – paket, ki vsebuje podatkovno strukturo vremenske postaje ter funkcije za dodajanje postaj, shranjevanje meritev in izpis podatkov ter statistike.
- **`cmd/weatherstation`** – glavni program (main.go), ki omogoča zagon aplikacije in nastavljanje parametrov prek ukazne vrstice.

## Zagon

V korenski mapi projekta izvedite ukaz:

```bash
go run ./cmd/weatherstation
```

Aplikacija se zažene s privzetimi vrednostmi stikal in izpiše vse shranjene meritve ter temperaturno poročilo.

## Stikala ukazne vrstice

| Stikalo | Opis | Privzeta vrednost |
|---|---|---|
| `--minTemp` | Najnižja dovoljena temperatura meritve (°C). | `-50` |
| `--maxTemp` | Najvišja dovoljena temperatura meritve (°C). | `50` |
| `--stMeritev` | Najmanjše število meritev, potrebno za izračun statistike. | `3` |

### Primer uporabe

```bash
go run ./cmd/weatherstation --minTemp=-15 --maxTemp=25 --stMeritev=2
```

V tem primeru:
- se upoštevajo samo meritve med -15 °C in 25 °C;
- za izračun statistike morata biti na postaji vsaj dve veljavni meritvi.

Program preveri tudi pravilnost nastavitev. Najnižja dovoljena temperatura ne sme biti višja od najvišje, število zahtevanih meritev pa mora biti vsaj 1.

## Funkcionalnosti paketa `vreme`

Paket `vreme` ponuja naslednje javne funkcije:

- `DodajPostajo` – doda novo vremensko postajo.
- `DodajMeritev` – doda temperaturno meritev izbrani postaji, če je znotraj dovoljenega temperaturnega območja.
- `IzpisMeritev` – izpiše vse vremenske postaje in njihove meritve.
- `IzpisPorocila` – izpiše temperaturno poročilo z najnižjo, najvišjo in povprečno temperaturo ter besedilnim opisom vremenskih razmer.

Izračun statistike se izvaja z interno funkcijo `statistika`.
