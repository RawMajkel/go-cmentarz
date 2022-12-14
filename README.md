# Obsługa Cmentarza - Interface / Channels

```
docker build -t go-cmentarz .
docker run go-cmentarz
```

W magazynie znajduje się 100 zniczy oraz 50 wiązanek.

Rzeczy pobierają 4 posłańcy: 2 posłańców pobiera tylko znicze i 2 posłańców pobiera tylko wiązanki;
[poslaniecZniczy_1,poslaniecZniczy_2,poslaniecWiazanek_1,poslaniecWiazanek_2];

Każdy posłaniec może jednorazowo pobrać tylko jedną sztukę rzeczy.

Pobrane znicze i wiązanki są chwilowo składowane w koszach, przed pobraniem ich przez Babki - sprzedawczynie.
Istnieją więc: kosz_na_znicze (pojemność 10) oraz kosz_na_wiązanki (pojemność 10)
Żaden posłaniec nie może włożyć rzeczy do kosza, jeśli jest on pełny, musi czekać aż Babka pobierze rzecz i zwolni miejsce.
Każdy posłaniec raportuje wykonanie zadań, przykładowo:
poslaniec1: pobranie znicz4;
poslaniec2: składowanie w kosz_na_znicze znicz6;

Babek jest 5. [babka1,babka2,...,babka5].
Każda pobiera (jeśli jest możliwość) 1 wiązankę i 2 znicze, przy okazji wypisując na ekran swój numer oraz numery pobranych rzeczy, przykładowo:
babka2: znicz4, znicz6, wiązanka9;

Możliwość pobrania jest wtedy, gdy kosz zawiera element. Jeśli kosz jest pusty babka czeka na doniesienie towaru przez posłańca.

```
go run .\cmentarz.go
```
