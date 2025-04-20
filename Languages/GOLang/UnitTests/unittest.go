package main

import "fmt"

// Function to return the monthly price in pennies
func getMonthlyPrice(tier string) int {
    switch tier {
    case "basic":
        return 100 * 100 // 100 dollars in pennies
    case "premium":
        return 150 * 100 // 150 dollars in pennies
    case "enterprise":
        return 500 * 100 // 500 dollars in pennies
    default:
        return 0 // Return 0 for invalid tiers
    }
}

func main() {
    // Test cases
    fmt.Println("Basic Plan Price:", getMonthlyPrice("basic"), "pennies")
    fmt.Println("Premium Plan Price:", getMonthlyPrice("premium"), "pennies")
    fmt.Println("Enterprise Plan Price:", getMonthlyPrice("enterprise"), "pennies")
    fmt.Println("Invalid Plan Price:", getMonthlyPrice("unknown"), "pennies")
}
