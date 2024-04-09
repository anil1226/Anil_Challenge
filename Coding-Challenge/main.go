package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidFormat(cardNumber string) bool {
	pattern := `^[456]\d{3}(-?\d{4}){3}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(cardNumber)
}

func hasRepeatedDigits(cardNumber string) bool {
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return true
		}
	}
	return false
}

func validateCreditCard(cardNumber string) string {
	if isValidFormat(cardNumber) && !hasRepeatedDigits(strings.ReplaceAll(cardNumber, "-", "")) {
		return "Valid"
	}
	return "Invalid"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the number of Creditcards to be validated")
	scanner.Scan()
	n := scanner.Text()
	count := 0

	fmt.Sscanf(n, "%d", &count)

	fmt.Println("Enter Card details one after one")
	for i := 0; i < count && scanner.Scan(); i++ {

		cardNumber := strings.TrimSpace(scanner.Text())
		fmt.Println(validateCreditCard(cardNumber))
	}
}
