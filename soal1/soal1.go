package main

import (
	"fmt"
	"sort"
)

func MoneyChange(money int) []int {
	listMoney := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	var res []int

	if money < 100 {
		res = append(res, money)
	}
	i := 0
	for money >= 0 && i < len(listMoney) {
		if money-listMoney[i] >= 0 {
			res = append(res, listMoney[i])
			money -= listMoney[i]
			//fmt.Println(money)
			if money <= 99 && money >= 1 {
				res = append(res, listMoney[9])
			}
		} else {
			i++
		}
	}
	//fmt.Println(res)
	return res

}

func Sort(items int) map[int]int {

	result := MoneyChange(items)

	//var res []int
	buat := map[int]int{}
	for i := 0; i < len(result); i++ {
		if _, found := buat[result[i]]; found {
			buat[result[i]]++
		} else {
			buat[result[i]] = 1
		}
	}
	sortVal := []int{}
	for _, val := range buat {
		sortVal = append(sortVal, val)
	}
	sort.Ints(sortVal)

	// for index, value := range buat {
	// 	if value != 1 {
	// 		buat[index] = -1
	// 	}
	// }
	//fmt.Println(buat)
	return buat
}

func main() {
	fmt.Println(Sort(145000))
	fmt.Println(Sort(2050))

}
