package main

import (
	"container/list"
	"fmt"
)

// zmienne globalne
var m Magazyn
var kz KoszNaZnicze
var kw KoszNaWiazanki

func MagazynyIKoszePuste() bool {
	return m.wiazanki.Len() == 0 && m.znicze.Len() == 0 && kz.znicze.Len() == 0 && kw.wiazanki.Len() == 0
}

func MoznaPobracWiazanke() bool {
	return m.wiazanki.Len() > 0 && kw.wiazanki.Len() < 10
}

func MoznaPobracZnicz() bool {
	return m.znicze.Len() > 0 && kz.znicze.Len() < 10
}

func MoznaPobracZKosza() bool {
	// fmt.Println("kosz na wiazanki:", kw.wiazanki.Len(), "kosz na znicze:", kz.znicze.Len())
	return kw.wiazanki.Len() >= 1 && kz.znicze.Len() >= 2
}

type Magazyn struct {
	znicze   *list.List
	wiazanki *list.List
}

func (m *Magazyn) Init() {
	for i := 1; i <= 100; i++ {
		m.znicze.PushFront(Znicz{fmt.Sprintf("znicz%d", i)})
	}

	for i := 1; i <= 50; i++ {
		m.wiazanki.PushFront(Wiazanka{fmt.Sprintf("wiazanka%d", i)})
	}
}

type Znicz struct {
	nazwa string
}
type Wiazanka struct {
	nazwa string
}

type PoslaniecZniczy struct {
	nazwa string
}

type PoslaniecWiazanek struct {
	nazwa string
}

func (p PoslaniecWiazanek) PobierzPrzedmiot() {
	wiazanka := m.wiazanki.Back()

	m.wiazanki.Remove(wiazanka)

	fmt.Println(fmt.Sprintf("%v pobiera %v z magazynu (pozostaje %d wiazanek)", p, wiazanka.Value, m.wiazanki.Len()))
}

func (p PoslaniecZniczy) PobierzPrzedmiot() {
	znicz := m.znicze.Back()

	m.znicze.Remove(znicz)
	kz.znicze.PushFront(znicz)

	fmt.Println(fmt.Sprintf("%v pobiera %v z magazynu i odkłada do kosza (pozostaje %d zniczy, kosz zawiera %d przedmiotów)", p, znicz.Value, m.znicze.Len(), kz.znicze.Len()))
}

func PobierzZnicz() *list.Element {
	znicz := m.znicze.Back()

	m.znicze.Remove(znicz)
	kz.znicze.PushFront(znicz)

	return znicz
}

func PobierzWiazanke() *list.Element {
	wiazanka := m.wiazanki.Back()

	m.wiazanki.Remove(wiazanka)
	kw.wiazanki.PushFront(wiazanka)

	return wiazanka
}

type Pobranie struct {
	wiazanka *list.Element
	znicz1   *list.Element
	znicz2   *list.Element
}

func PobierzZKosza() Pobranie {
	wiazanka := kw.wiazanki.Back()
	znicz1 := kz.znicze.Back()
	znicz2 := kz.znicze.Back()

	kw.wiazanki.Remove(wiazanka)
	kz.znicze.Remove(znicz1)
	kz.znicze.Remove(znicz2)

	return Pobranie{wiazanka, znicz1, znicz2}
}

type KoszNaZnicze struct {
	znicze *list.List
}

type KoszNaWiazanki struct {
	wiazanki *list.List
}

type Babka struct {
	nazwa string
}

func main() {
	// stwórz magazyn
	m = Magazyn{list.New(), list.New()}
	m.Init()

	// stwórz posłańców
	// pz_1 := PoslaniecZniczy{"poslaniecZniczy_1"}
	// pz_2 := PoslaniecZniczy{"poslaniecZniczy_2"}
	// pw_1 := PoslaniecWiazanek{"poslaniecWiazanek_1"}
	// pw_2 := PoslaniecWiazanek{"poslaniecWiazanek_1"}
	// // pz_1.PobierzPrzedmiot(&m)
	// // pw_1.PobierzPrzedmiot(&m)

	// // stwórz babki
	// b1 := Babka{"babka_1"}
	// b2 := Babka{"babka_2"}
	// b3 := Babka{"babka_3"}
	// b4 := Babka{"babka_4"}
	// b5 := Babka{"babka_5"}

	// stwórz kosze (zmienne globalne)
	kz = KoszNaZnicze{list.New()}
	kw = KoszNaWiazanki{list.New()}

	// for kz.znicze.Len() < 10 {
	// 	pz_1.PobierzPrzedmiot()
	// }

	ch_znicze_1 := make(chan *list.Element)
	ch_znicze_2 := make(chan *list.Element)
	ch_wiazanki_1 := make(chan *list.Element)
	ch_wiazanki_2 := make(chan *list.Element)

	ch_babka_1 := make(chan Pobranie)
	ch_babka_2 := make(chan Pobranie)
	ch_babka_3 := make(chan Pobranie)
	ch_babka_4 := make(chan Pobranie)
	ch_babka_5 := make(chan Pobranie)

	fmt.Printf("Stan magazynu: %d zniczy, %d wiazanek\n", m.znicze.Len(), m.wiazanki.Len())
	fmt.Printf("Kosze: %d zniczy, %d wiazanek\n", kz.znicze.Len(), kw.wiazanki.Len())

	go func() {
		for {
			if MoznaPobracZnicz() {
				ch_znicze_1 <- PobierzZnicz()
			}
		}
	}()

	go func() {
		for {
			if MoznaPobracZnicz() {
				ch_znicze_2 <- PobierzZnicz()
			}
		}
	}()

	go func() {
		for {
			if MoznaPobracWiazanke() {
				ch_wiazanki_1 <- PobierzWiazanke()
			}
		}
	}()

	go func() {
		for {
			if MoznaPobracWiazanke() {
				ch_wiazanki_2 <- PobierzWiazanke()
			}
		}
	}()

	go func() {
		for {
			if MoznaPobracZKosza() {
				ch_babka_1 <- PobierzZKosza()
			}
		}
	}()
	go func() {
		for {
			if MoznaPobracZKosza() {
				ch_babka_2 <- PobierzZKosza()
			}
		}
	}()
	go func() {
		for {
			if MoznaPobracZKosza() {
				ch_babka_3 <- PobierzZKosza()
			}
		}
	}()
	go func() {
		for {
			if MoznaPobracZKosza() {
				ch_babka_4 <- PobierzZKosza()
			}
		}
	}()
	go func() {
		for {
			if MoznaPobracZKosza() {
				ch_babka_5 <- PobierzZKosza()
			}
		}
	}()

	// główna pętla
	for {
		select {
		//babka 1
		case rec5 := <-ch_babka_1:
			fmt.Println(fmt.Sprintf("babka_1 pobiera i sprzedaje %v, %v i %v", rec5.wiazanka, rec5.znicz1, rec5.znicz2))
		//babka 2
		case rec6 := <-ch_babka_2:
			fmt.Println(fmt.Sprintf("babka_2 pobiera i sprzedaje %v, %v i %v", rec6.wiazanka, rec6.znicz1, rec6.znicz2))
		//babka 3
		case rec7 := <-ch_babka_3:
			fmt.Println(fmt.Sprintf("babka_3 pobiera i sprzedaje %v, %v i %v", rec7.wiazanka, rec7.znicz1, rec7.znicz2))
		//babka 4
		case rec8 := <-ch_babka_4:
			fmt.Println(fmt.Sprintf("babka_4 pobiera i sprzedaje %v, %v i %v", rec8.wiazanka, rec8.znicz1, rec8.znicz2))
		//babka 5
		case rec9 := <-ch_babka_5:
			fmt.Println(fmt.Sprintf("babka_5 pobiera i sprzedaje %v, %v i %v", rec9.wiazanka, rec9.znicz1, rec9.znicz2))
		//poslaniec zniczy 1
		case rec1 := <-ch_znicze_1:
			fmt.Println(fmt.Sprintf("poslaniec_zniczy_1 pobiera %v i odklada do kosza", rec1.Value))
			// fmt.Println(fmt.Sprintf("Magazyn: %d, Kosz %d", m.znicze.Len(), kz.znicze.Len()))
		//poslaniec zniczy 2
		case rec2 := <-ch_znicze_2:
			fmt.Println(fmt.Sprintf("poslaniec_zniczy_2 pobiera %v i odklada do kosza", rec2.Value))
			// fmt.Println(fmt.Sprintf("Magazyn: %d, Kosz %d", m.znicze.Len(), kz.znicze.Len()))
		//poslaniec wiazanek 1
		case rec3 := <-ch_wiazanki_1:
			fmt.Println(fmt.Sprintf("poslaniec_wiazanek_1 pobiera %v i odklada do kosza", rec3.Value))
			// fmt.Println(fmt.Sprintf("Magazyn: %d, Kosz %d", m.wiazanki.Len(), kw.wiazanki.Len()))
		//poslaniec wiazanek 2
		case rec4 := <-ch_wiazanki_2:
			fmt.Println(fmt.Sprintf("poslaniec_wiazanek_2 pobiera %v i odklada do kosza", rec4.Value))
			// fmt.Println(fmt.Sprintf("Magazyn: %d, Kosz %d", m.wiazanki.Len(), kw.wiazanki.Len()))
		}
	}

	fmt.Println("Koniec programu - magazyny i kosze puste")

	// for m.znicze.Len() > 0 {
	// 	fmt.Println(fmt.Sprintf("Magazyn zniczy nie jest pusty (%d)", m.znicze.Len()))

	// 	pz_1.PobierzPrzedmiot(&m)
	// 	time.Sleep(10 * time.Millisecond)
	// }

	// for {
	// 	select {
	// 	case c <- x:
	// 		x, y = y, x+y
	// 	case <-quit:
	// 		fmt.Println("quit")
	// 		return
	// 	}
	// }
}
