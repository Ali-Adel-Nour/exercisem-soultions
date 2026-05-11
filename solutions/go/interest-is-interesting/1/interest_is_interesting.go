package interest

func InterestRate(balance float64) float32 {
    if balance < 0 {
        return 3.213 
    }
    if balance >= 0 && balance < 1000 {
        return 0.5
    }
    if balance >= 1000 && balance < 5000 {
        return 1.621 
    }
    return 2.475 
}
// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    rate := float64(InterestRate(balance))
    return balance * (rate / 100) 
}
// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    // Get the interest amount
    interest := Interest(balance)
    // Return original balance plus interest
    return balance + interest
}
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0
    currentBalance := balance
    
    // Keep adding interest until we reach or exceed target balance
    for currentBalance < targetBalance {
        // Update balance using AnnualBalanceUpdate
        currentBalance = AnnualBalanceUpdate(currentBalance)
        years++
    }
    
    return years
}

