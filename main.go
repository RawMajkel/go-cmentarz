package main

import (
	"fmt"
)

type Znicz string
type Wiazanka string

type Przedmiot interface {
	Znicz | Wiazanka
}

var magazynZnicze chan Znicz
var magazynWiazanki chan Wiazanka
var koszZnicze chan Znicz
var koszWiazanki chan Wiazanka

func StworzMagazyn[T Przedmiot](ilosc int, nazwa string) chan T {

	magazyn := make(chan T, ilosc)

	for i := 1; i <= ilosc; i++ {
		magazyn <- T(fmt.Sprintf("%s%d", nazwa, i))
	}

	return magazyn
}

func PracaPoslancaZniczy(numer int) {

	for len(magazynZnicze) > 0 {
		przedmiot := <-magazynZnicze
		fmt.Printf("poslaniecZniczy_%d pobranie: %s\n", numer, przedmiot)

		for {
			if len(koszZnicze) <= 10 {
				koszZnicze <- przedmiot
				fmt.Printf("poslaniecZniczy_%d skladowanie: %s\n", numer, przedmiot)

				break
			}
		}
	}
}

func PracaPoslancaWiazanek(numer int) {

	for len(magazynWiazanki) > 0 {
		przedmiot := <-magazynWiazanki
		fmt.Printf("poslaniecWiazanek_%d pobranie: %s\n", numer, przedmiot)

		for {
			if len(koszZnicze) <= 10 {
				koszWiazanki <- przedmiot
				fmt.Printf("poslaniecWiazanek_%d skladowanie: %s\n", numer, przedmiot)

				break
			}
		}
	}
}

func PracaBabki(numer int) {

	for {
		if (len(koszZnicze) >= 2) && (len(koszWiazanki) >= 1) {

			wiazanka := <-koszWiazanki
			znicz1 := <-koszZnicze
			znicz2 := <-koszZnicze

			fmt.Printf("babka_%d: %s, %s, %s\n", numer, wiazanka, znicz1, znicz2)
		}
	}
}

func main() {

	magazynZnicze = StworzMagazyn[Znicz](100, "znicz")
	magazynWiazanki = StworzMagazyn[Wiazanka](50, "wiazanka")

	koszZnicze = make(chan Znicz, 10)
	koszWiazanki = make(chan Wiazanka, 10)

	for i := 1; i <= 2; i++ {
		go PracaPoslancaZniczy(i)
	}

	for i := 1; i <= 2; i++ {
		go PracaPoslancaWiazanek(i)
	}

	for i := 1; i <= 5; i++ {
		go PracaBabki(i)
	}

	for {
		if len(magazynZnicze)+len(magazynWiazanki)+len(koszZnicze)+len(koszWiazanki) == 0 {
			fmt.Println("KONIEC")
			return
		}
	}
}
