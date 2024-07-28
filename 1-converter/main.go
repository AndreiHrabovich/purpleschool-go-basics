package main

import "fmt"

const usdEur = 0.92
const usdRub float64 = 86
const eurRub = (1 / usdEur) / (1 / usdRub)

func main() {
	fmt.Print(eurRub)
}

/*
	1 usd = 0.92 eur
	1 usd = 86 rub
	0.92 = 86
	1 / 0.92 = 1.08696
	1 / 86 = 0.01163
	1.08696 / 0.01163 = 93.46
*/
