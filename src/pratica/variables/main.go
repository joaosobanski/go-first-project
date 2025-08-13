package main

import (
	"fmt"
	"time"
)

func revertString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		opp := n - 1 - i
		runes[i], runes[opp] = runes[opp], runes[i]
		fmt.Printf("Troca %d: %c com %c\n", i, runes[i], runes[opp])
	}

	return string(runes)
}

func revertAllSlice(v []int) []int {
	length := len(v)
	for i := 0; i < length/2; i++ {
		opp := length - 1 - i
		v[i], v[opp] = v[opp], v[i]
	}
	return v
}

func revertSlice(v []int, p1 int, p2 int) []int {
	if p1 < 0 || p2 >= len(v) || p1 >= p2 {
		return v
	}

	v[p1], v[p2] = v[p2], v[p1]

	return v
}

func main() {
	x := "string1"
	fmt.Println(x)
	y := "string2"
	fmt.Println(y)
	z := x + y
	fmt.Println(z)

	i1 := 10
	f1 := 20.1
	f2 := float64(i1)
	fmt.Println(f2 + f1)

	var gnirts = revertString("string1")
	fmt.Println(gnirts)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	revertedSlice := revertAllSlice(slice)
	fmt.Println(revertedSlice)

	otherSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(otherSlice)

	revertedSlice2 := revertSlice(otherSlice, 0, 1)
	fmt.Println(revertedSlice2)

	var sumSlice = []int{}

	for i := 0; i < len(slice); i++ {
		sumSlice = append(sumSlice, slice[i]+otherSlice[i])
	}
	fmt.Println("Slice:", slice)
	fmt.Println("Other Slice:", otherSlice)
	fmt.Println("Sum Slice:", sumSlice)

	// 	Memorização fácil: A data de referência é 1 2 3 4 5 6 7 em ordem:
	// 01 = Mês (January)
	// 02 = Dia
	// 03 = Hora (15h = 3pm)
	// 04 = Minuto
	// 05 = Segundo
	// 06 = Ano (2006)
	// 07 = Dia da semana
	// esta data pega a hora da maquina e não a UTC
	// Para utilizar UTC: now := time.Now().UTC()

	now := time.Now().UTC()
	fiveDaysAgo := now.AddDate(0, 0, -5)
	fmt.Println("Agora:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("5 dias atrás:", fiveDaysAgo.Format("2006-01-02 15:04:05"))

	tenMinutesAgo := now.Add(-10 * time.Minute)
	fmt.Println("10 minutos atrás:", tenMinutesAgo.Format("2006-01-02 15:04:05"))

	tenSecundsAgo := now.Add(-10 * time.Second)
	fmt.Println("10 segundos atrás:", tenSecundsAgo.Format("2006-01-02 15:04:05"))

}
