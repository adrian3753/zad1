package main

import (
	"fmt"
	"strconv"
)

func wartoscCiaguFibonacci(n int64) int64 {

	if n < 2 {
		return n
	}

	var q11, q12, q21, q22 int64
	var w11, w12, w21, w22 int64
	var p11, p12, p21, p22 int64

	q11, q12, q21 = 1, 1, 1
	q22 = 0

	w11, w22 = 1, 1
	w12, w21 = 0, 0

	n = n - 1

	for n != 0 {
		if (n & 1) == 1 {
			p11 = w11*q11 + w12*q21
			p12 = w11*q12 + w12*q22
			p21 = w21*q11 + w22*q21
			p22 = w21*q12 + w22*q22

			w11 = p11
			w12 = p12
			w21 = p21
			w22 = p22
		}

		n = n >> 1

		if n == 0 {
			break
		}

		p11 = q11*q11 + q12*q21
		p12 = q11*q12 + q12*q22
		p21 = q21*q11 + q22*q21
		p22 = q21*q12 + q22*q22

		q11 = p11
		q12 = p12
		q21 = p21
		q22 = p22
	}
	return w11
}

func wyswietlDane() {
	fmt.Println("-------------------------")
	fmt.Println("Nazwa programu: FibCalc")
	fmt.Println("Autor: Adrian Szafrański")
	fmt.Println("Grupa dziekańska: I2S 1.5")
	fmt.Println("-------------------------")
	fmt.Println()
}

func main() {

	var numElemFibString string
	bledyWalidacji := true
	wyswietlDane()

	for bledyWalidacji {
		fmt.Print("Podaj numer elementu ciągu Fibonacciego (zakres 0 - 92): ")
		fmt.Scan(&numElemFibString)

		numElemFib, err := strconv.ParseInt(numElemFibString, 10, 64)

		if err == nil && (numElemFib >= 0 && numElemFib <= 92) {
			fmt.Printf("Wprowadzony numer elementu ciągu: %d. Jego wartość to: %d", numElemFib, wartoscCiaguFibonacci(numElemFib))
			fmt.Println()
			bledyWalidacji = false
		} else {
			fmt.Println("Nie podałeś liczby całkowitej z zakresu 0 - 92")
			fmt.Println()
		}
	}
}
