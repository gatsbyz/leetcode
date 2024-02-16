func maxProfit(prices []int) int {
    profit := 0
    buyPrice := prices[0]
    for _, price := range prices[1:] {
        if price < buyPrice {
            buyPrice = price
        } else {
            if price - buyPrice > profit {
                profit = price - buyPrice
            }
        }
    }
    return profit
}