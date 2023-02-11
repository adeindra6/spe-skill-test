package main

import (
	"fmt"
	"math"
)

func main() {
	// Narcissistic Number
	fmt.Println("=== Narcissistic Number ===")
	fmt.Println(narcissistic(153))
	fmt.Println(narcissistic(111))

	// Parity Outliers
	fmt.Println("=== Parity Outliers ===")
	x := []int64{2, 4, 0, 100, 4, 11, 2602, 36}
	num, avail := outliers(x)
	if avail {
		fmt.Println(num)
	} else {
		fmt.Println(avail)
	}

	x = []int64{160, 3, 1719, 19, 11, 13, -21}
	num, avail = outliers(x)
	if avail {
		fmt.Println(num)
	} else {
		fmt.Println(avail)
	}

	x = []int64{11, 13, 15, 19, 9, 13, -21}
	num, avail = outliers(x)
	if avail {
		fmt.Println(num)
	} else {
		fmt.Println(avail)
	}
}

func narcissistic(number int64) bool {
	var narcissistic bool

	num := number
	slc := []int64{}
	for number > 0 {
		slc = append(slc, number%10)
		number /= 10
	}

	slc_int := []int64{}
	for i := range slc {
		slc_int = append(slc_int, slc[len(slc)-1-i])
	}

	lastIndex := len(slc_int) - 1
	result := 0

	for i := range slc_int {
		result = result + int(math.Pow(float64(slc_int[i]), float64(slc_int[lastIndex])))
	}

	if num == int64(result) {
		narcissistic = true
	} else {
		narcissistic = false
	}

	return narcissistic
}

func outliers(arr_num []int64) (int64, bool) {
	if len(arr_num) < 3 {
		fmt.Println("Minimum array length is 3")
		return 0, false
	}

	var even_num []int64
	var odd_num []int64

	for i := range arr_num {
		if arr_num[i]%2 == 0 {
			even_num = append(even_num, int64(arr_num[i]))
		} else {
			odd_num = append(odd_num, int64(arr_num[i]))
		}
	}

	var outliers_num int64
	var outliers_avail bool
	if len(even_num) == 1 {
		outliers_num = even_num[0]
		outliers_avail = true
	} else if len(odd_num) == 1 {
		outliers_num = odd_num[0]
		outliers_avail = true
	} else {
		outliers_num = 0
		outliers_avail = false
	}

	return outliers_num, outliers_avail
}
