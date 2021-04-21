package pkg

func CalculateInterest(savingsBalance float32) float32 {
	var balanceToAdd float32
	var interest float32

	if savingsBalance > 0 && savingsBalance <= 500 {
		interest = 1
		balanceToAdd = interestCalculation(interest, savingsBalance)
	}

	if savingsBalance > 500 && savingsBalance <= 1000 {
		interest = 2
		balanceToAdd = interestCalculation(interest, savingsBalance)
	}

	if savingsBalance > 1000 {
		interest = 3
		balanceToAdd = interestCalculation(interest, savingsBalance)
	}

	return balanceToAdd
}

func interestCalculation(interest float32, savingsBalance float32) float32 {
	return (interest / 100) * savingsBalance
}
