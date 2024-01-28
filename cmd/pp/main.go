package main

import (
	"fmt"
)

const (
	ProductsPerVydacha int = 48
)

func readIntFromConsole(textAskingInt string) (int, error) {
	fmt.Print(textAskingInt)
	var resultInt int
	_, err := fmt.Scanf("%d", &resultInt)
	if err != nil {
		return -1, err
	}

	return resultInt, nil
}

func generateVydacha(vydacha [ProductsPerVydacha]int64, PPvydacha []int64, PP_place int) [ProductsPerVydacha]int64 {
	var vydachaWithPP [ProductsPerVydacha]int64
	var PPvydachaCounter uint = 0
	var vydachaCounter uint = 0
	for vydachaWithPPcounter := 0; vydachaWithPPcounter < ProductsPerVydacha; vydachaWithPPcounter++ {
		fmt.Println(vydachaWithPPcounter, PPvydachaCounter, vydachaCounter)
		if vydachaWithPPcounter%PP_place == 0 && PPvydachaCounter < uint(len(PPvydacha)) {
			vydachaWithPP[vydachaWithPPcounter] = PPvydacha[PPvydachaCounter]
			PPvydachaCounter++
		} else {
			vydachaWithPP[vydachaWithPPcounter] = vydacha[vydachaCounter]
			vydachaCounter++
		}
	}
	return vydachaWithPP
}

func main() {
	products_ids := [48]int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}
	PPproducts_len, err := readIntFromConsole("Enter a number of PP products: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	PP_place, err := readIntFromConsole("Enter a number of PP place: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Enter %d PP products ids: ", PPproducts_len)
	PPproducts_ids := make([]int64, PPproducts_len)
	for i := 0; i < PPproducts_len; i++ {
		fmt.Scan(&PPproducts_ids[i])
	}

	vydacha := generateVydacha(products_ids, PPproducts_ids, PP_place)
	fmt.Println(vydacha)
}
