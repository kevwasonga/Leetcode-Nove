// 121. Best Time to Buy and Sell Stock     https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
package main

import "fmt"

func maxProfit(prices []int) int {
	left, right := 0, 1 // left is the buy day, right is the sell day
	maxProfit := 0      // Initialize the maximum profit as 0

	for right < len(prices) {
		// If the selling price is greater than the buying price, we calculate profit
		if prices[right] > prices[left] {
			profit := prices[right] - prices[left]
			// Update maxProfit if we find a better profit
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			// If the selling price is less, move the left pointer to right
			// because a lower price means a better buying opportunity
			left = right
		}
		// Move the right pointer to the next day
		right++
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max Profit:", maxProfit(prices)) // Should print 5
}
