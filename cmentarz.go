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

//[poslaniecZniczy_1,poslaniecZniczy_2,poslaniecWiazanek_1,poslaniecWiazanek_2];
// type Poslaniec interface {
// 	PobierzPrzedmiot() string
// }

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