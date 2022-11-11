package main

import (
	"container/list"
	"fmt"
)

type Magazyn struct {
	znicze *list.List
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

type Poslaniec interface {
	PobierzPrzedmiot(m *Magazyn)
}

func (p PoslaniecWiazanek) PobierzPrzedmiot(m *Magazyn) {
	wiazanka := m.wiazanki.Back()

	m.wiazanki.Remove(wiazanka)

	fmt.Println(fmt.Sprintf("%v pobiera %v z magazynu (pozostaje %d wiazanek)", p, wiazanka.Value, m.wiazanki.Len()))
}

func (p PoslaniecZniczy) PobierzPrzedmiot(m *Magazyn) {
	znicz := m.znicze.Back()

	m.znicze.Remove(znicz)

	fmt.Println(fmt.Sprintf("%v pobiera %v z magazynu (pozostaje %d zniczy)", p, znicz.Value, m.znicze.Len()))
}

// 100 zniczy, 50 wiazanek, 4 poslancow (2x znicze, 2x wiazanki), kosz_na_znicze (poj 10), kosz_na_wiazanki (poj 10), 5 babek
// var znicze [100]Znicz
func main() {

	// stwórz magazyn
	magazyn := Magazyn{list.New(), list.New()}
	magazyn.Init()

	// fmt.Println(magazyn.znicze.Len())
	// fmt.Println(magazyn.wiazanki.Len())
	// fmt.Println(magazyn.wiazanki.Back().Value)
	// for e := magazyn.wiazanki.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	// stwórz posłańców
	pz_1 := PoslaniecZniczy{"poslaniecZniczy_1"}
	pz_2 := PoslaniecZniczy{"poslaniecZniczy_2"}
	pw_1 := PoslaniecWiazanek{"poslaniecWiazanek_1"}
	pw_2 := PoslaniecWiazanek{"poslaniecWiazanek_1"}
	// p1.PobierzPrzedmiot(&magazyn)

	// for temp := magazyn.wiazanki.Front(); temp != nil; temp = temp.Next() {
	// 	fmt.Println(temp.Value)
	// }

	// znicze := list.New()
	// for i := 0; i < 100; i++ {
	// 	znicze.PushBack(fmt.Sprintf("znicz%d", i+1))
	// }

	// wiazanki := list.New()
	// for i := 0; i < 50; i++ {
	// 	wiazanki.PushBack(fmt.Sprintf("wiązanka%d", i+1))
	// }

	// poslaniecZniczy1
	// poslaniecZniczy2
	// poslaniecWiazanek1
	// poslaniecWiazanek2


	// for temp := wiazanki.Front(); temp != nil; temp = temp.Next() {
	// 	fmt.Println(temp.Value)
	// }

	// messages := make(chan string)
	
	// go func() {

	// }();
}