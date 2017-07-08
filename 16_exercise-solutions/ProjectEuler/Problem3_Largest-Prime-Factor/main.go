package main

import (
	"fmt"
	"math"
)

/*******************************************************************/
/*        Project Euler - Challenge 3: Largest Prime factor        */
/* The prime factors of 13195 are 5, 7, 13 and 29.                 */
/* What is the largest prime factor of the number 600851475143 ?   */
/*******************************************************************/
func getPrimeNumbers(n int) []int {
	var numList, primeFactorList []int

	//Get square root of argument number passed in. Prime factors will be lower than the square root.
	var square float64 = math.Sqrt(float64(n))
	x := int(square)
	fmt.Printf("The Square root of the number %d is: %d\n", n, x)

	//A prime number will always satisfy the formula 6n-1 or 6n+1 (result being a natural number).
	//Loop over numbers from 2 to the square root of the input argument and append to a slise - numList.
	//NOTE: This will include numbers that aren't prime.
	for i := 2; i < x; i++ {
		form1 := (i - 1) % 6
		form2 := (i + 1) % 6
		if form1 == 0 || form2 == 0 && !(form1 == 0 && form2 == 0) {
			numList = append(numList, i)
		}
	}

	//Loop over numList to filter out any numbers that aren't prime
	for _, v := range numList {
		var count int
		var appendNum bool

		//NOTE: A prime number is divisible by itself and 1 (two factors).
		//For each value in the slise, check wether it is a factor of the input arg
		//and iterate down to 1 keeping a count of the factors.
		for i := v; i >= 1; i-- {
			if count > 2 || (n%v != 0) {
				appendNum = false
				break
			} else if v%i == 0 {
				count++
				appendNum = true
			}
		}
		//IF the factor count for the value in the slise is 2 then it is a primeFactor
		if count <= 2 && appendNum {
			primeFactorList = append(primeFactorList, v)
		}
	}
	return primeFactorList
}

//Func to return the max value for an input slise of ints.
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
	primeFactorList := getPrimeNumbers(13195)
	fmt.Println("The list of prime factors are: ", primeFactorList)
	fmt.Println(getMaxPrimeFactor(primeFactorList...))
	primeFactorList = getPrimeNumbers(600851475143)
	fmt.Println("The list of prime factors are: ", primeFactorList)
	fmt.Println(getMaxPrimeFactor(primeFactorList...))
}
