package main

import (
	"fmt"
)

var magazynZnicze chan Znicz
var magazynWiazanki chan Wiazanka
var koszZnicze chan Znicz
var koszWiazanki chan Wiazanka

type Znicz struct {
	nazwa string
}

type Wiazanka struct {
	nazwa string
}

func StworzMagazynZniczy(quantity int) chan Znicz {
	magazyn := make(chan Znicz, quantity)

	for i := 1; i <= quantity; i++ {
		magazyn <- Znicz{fmt.Sprintf("znicz%d", i)}
	}

	return magazyn
}

func StworzMagazynWiazanek(quantity int) chan Wiazanka {
	magazyn := make(chan Wiazanka, quantity)

	for i := 1; i <= quantity; i++ {
		magazyn <- Wiazanka{fmt.Sprintf("wiazanka%d", i)}
	}

	return magazyn
}

func PracaPoslancaZniczy(numer int) {

	for len(magazynZnicze) > 0 {
		produkt := <-magazynZnicze
		fmt.Printf("poslaniec_zniczy_%d pobral z magazynu: %v \n", numer, produkt)

		for {
			if len(koszZnicze) <= 10 {
				koszZnicze <- produkt
				fmt.Printf("poslaniec_zniczy_%d umiescil %v w koszu \n", numer, produkt)

				break
			} else {
				fmt.Printf("poslaniec_zniczy_%d czeka na wolny kosz \n", numer)
			}
		}
	}
	fmt.Printf("poslaniec_zniczy_%d zakonczyl prace", numer)
}

func PracaPoslancaWiazanek(numer int) {

	for len(magazynWiazanki) > 0 {
		produkt := <-magazynWiazanki
		fmt.Printf("poslaniec_wiazanek_%d pobral z magazynu: %v \n", numer, produkt)

		for {
			if len(koszZnicze) <= 10 {
				koszWiazanki <- produkt
				fmt.Printf("poslaniec_wiazanek_%d umiescil %v w koszu \n", numer, produkt)

				break
			} else {
				fmt.Printf("poslaniec_wiazanek_%d czeka na wolny kosz \n", numer)
			}
		}
	}
	fmt.Printf("poslaniec_wiazanek_%d zakonczyl prace", numer)
}

func PracaBabek() {

	for {
		for i := 1; i <= 5; i++ {

			if (len(koszZnicze) >= 2) && (len(koszWiazanki) >= 1) {

				znicz1 := <-koszZnicze
				znicz2 := <-koszZnicze
				wiazanka1 := <-koszWiazanki

				response := fmt.Sprintf("babka_%d pobrala: %v | %v | %v", i, znicz1, znicz2, wiazanka1)
				fmt.Println(response)
			}
		}
	}
}

func main() {

	// magazyn
	magazynZnicze = StworzMagazynZniczy(100)
	magazynWiazanki = StworzMagazynWiazanek(50)

	// stwórz kosze
	koszZnicze = make(chan Znicz, 10)
	koszWiazanki = make(chan Wiazanka, 10)

	for {
		select {
		case znicz := <-koszZnicze:
			{
				if (len(koszZnicze) >= 2) && (len(koszWiazanki) >= 1) {
					
				}
			}
		case znicz := <-magazynZnicze:
			{
				fmt.Printf("poslaniec_zniczy_1 pobral i zlozyl do kosza %v\n", znicz)
			}
		case znicz := <-magazynZnicze:
			{
				fmt.Printf("poslaniec_zniczy_2 pobral i zlozyl do kosza %v\n", znicz)
			}
		case wiazanka := <-magazynWiazanki:
			{
				fmt.Printf("poslaniec_wiazanek_1 pobral i zlozyl do kosza %v\n", wiazanka)
			}
		case wiazanka := <-magazynWiazanki:
			{
				fmt.Printf("poslaniec_wiazanek_2 pobral i zlozyl do kosza %v\n", wiazanka)
			}
		}
	}

	// for i := 1; i <= 2; i++ {
	// 	go PracaPoslancaZniczy(i)
	// }

	// for i := 1; i <= 2; i++ {
	// 	go PracaPoslancaWiazanek(i)
	// }

	// go PracaBabek()

	// for {
	// 	if len(magazynZnicze) == 0 && len(magazynWiazanki) == 0 && len(koszZnicze) == 0 && len(koszWiazanki) == 0 {
	// 		fmt.Println("KONIEC")
	// 		break
	// 	}
	// }
}
