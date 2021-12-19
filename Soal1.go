package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{4, 11, 2, 20, 59, 80}
	i := 2
	profit := MaxProfit(prices, i)
	fmt.Println("Hasil : ",profit)
}

func MaxProfit(prices []int, n int) int {
	result := 0
	profit := make([]int, n)
	for i:=0; i<n; i++ {
		profit[i] = 0
	}
	max_price := profit[n - 1];
	for i:=n-2; i >= 0; i-- {
		if prices[i] > max_price {
			max_price = prices[i];
		}
		maxData := math.Max(float64(profit[i + 1]), float64(max_price - prices[i]))
		profit[i] = int(maxData)
	}
	min_price := prices[0];
	for i:=1; i<n; i++ {
		if prices[i] < min_price {
			min_price = prices[i];
		}
		minData := math.Max(float64(profit[i - 1]), float64(profit[i] + (prices[i] - min_price)))
		profit[i] = int(minData)
	}
	result = profit[n - 1];
	return result
}