package main

import (
	"fmt"
	"math"
)

func getPrimeNumbers(n int) []int {
	var primeNumList, primeFactorList []int
	var square float64 = math.Sqrt(float64(n))
	x := int(square)
	fmt.Println(x)

	for i := 2; i < x; i++ {
		form1 := (i - 1) % 6
		form2 := (i + 1) % 6
		if form1 == 0 || form2 == 0 && !(form1 == 0 && form2 == 0) {
			primeNumList = append(primeNumList, i)
		}
	}

	for _, v := range primeNumList {
		var count int
		var appendNum bool

		for i := v; i >= 1; i-- {
			if count > 2 || (n%v != 0) {
				appendNum = false
				break
			} else if v%i == 0 {
				count++
				appendNum = true
			}
		}
		if count <= 2 && appendNum {
			primeFactorList = append(primeFactorList, v)
		}
	}
	return primeFactorList
}

func getMaxPrimeFactor(n ...int) int {
	var maxPrimeFactor int
	for _, v := range n {
		if v > maxPrimeFactor {
			maxPrimeFactor = v
		}
	}
	return maxPrimeFactor
}

func main() {
	primeFactorList := getPrimeNumbers(600851475143)
	fmt.Println(primeFactorList)
	fmt.Println(getMaxPrimeFactor(primeFactorList...))
	//fmt.Println(primeNumber(factorList...))
}

//600,851,475,143
