package commands

import (
	"fmt"
	"math"
	"time"
)

// Costant for different percent
// For sum more than 210 000 it's 8% per month and 0.26% per day
// For sum less than 210 000 it's 7.16%% per month and 0.2327% per day
const (
	minTax    float64 = 0.0716
	maxTax    float64 = 0.08
	dayMinTax float64 = 0.002327
	dayMaxTax float64 = 0.0026
	threshold int     = 210000
)

var dep, dayly, money, relax, first float64
var day int

func Calculate(dep float64, dur int) (float64, int) {

	days := daysInMonth(dur)
	fmt.Println("all days is: ", days)

	//months := math.Round(float64(days) / 30)
	//fmt.Println("all days is: ", days)

	// Calculating refound for the deposit period without reinvesting
	relax = dep + dep*minTax*float64(dur)
	if dep > float64(threshold) {
		relax = dep + dep*maxTax*float64(dur)
	}
	fmt.Println("\nХорошо, без реинвестирования ваш депозит через", dur, "месяцаев, составит:")
	fmt.Printf("%.2f\n", relax)

	// Calculating the first interest payment (which comes the next day, and increases every day)
	first = dep * dayMinTax
	if dep > float64(threshold) {
		first = dep * dayMaxTax
	}
	fmt.Printf("Ежедневно вам будет начисляться процент, начиная с\n%.2f\n", first)

	return dep, dur
}

func daysInMonth(month int) int {
	var now = time.Now()
	//fmt.Scanln(&month)
	for month < 1 {
		fmt.Println("Введите пожалуйста значение больше нуля")
		fmt.Scanln(&month)
	}
	var se = now.AddDate(0, month, 0)
	dur := se.Sub(now)
	allday := dur.Hours() / 24
	return int(allday)
}

// Function for calculate hard percent (menu #1)
func calculateold() {
	// Say "Hello" and ask a sum of deposit
	fmt.Println("\nПриветствую! Это калькулятор сложного процента для проекта SpaceBot!\n\nВведите сумму депозита в рублях:")
	fmt.Scanln(&dep)

	// Check positive sum
	// If error - show error message and ask again
	for dep < 1 {
		fmt.Println("Введите пожалуйста значение больше нуля")
		fmt.Scanln(&dep)
	}
	money = dep

	// Ask a deposit period (month)
	fmt.Println("\nНа сколько месяцев вносим депозит?")
	days := calct(day)
	months := math.Round(float64(days) / 30)

	// Calculating refound for the deposit period without reinvesting
	relax = dep + dep*minTax*float64(months)
	if dep > float64(threshold) {
		relax = dep + dep*maxTax*float64(months)
	}
	fmt.Println("\nХорошо, без реинвестирования ваш депозит через", months, "месяцаев, составит:")
	fmt.Printf("%.2f\n", relax)

	// Calculating the first interest payment (which comes the next day, and increases every day)
	first = money * dayMinTax
	if money > float64(threshold) {
		first = money * dayMaxTax
	}
	fmt.Printf("Ежедневно вам будет начисляться процент, начиная с\n%.2f\n", first)

	// Reinvestment calculation for deposit period with dayly reinvesting
	for i := 0; i < days; i++ {
		if money > float64(threshold) {
			dayly = money * dayMinTax
			money = dayly + money
		} else {
			dayly = money * dayMaxTax
			money = dayly + money
		}
	}
	fmt.Println("\nВаш депозит через", months, "месяцяев при ежедневном реинвестировании:")
	fmt.Printf("%.2f\n", money)
}

// Function for calculation of calendar days depending on the current date
// Ask a number of months, then check positive date
// If error - show error message and ask again
func calct(month int) int {
	var now = time.Now()
	fmt.Scanln(&month)
	for month < 1 {
		fmt.Println("Введите пожалуйста значение больше нуля")
		fmt.Scanln(&month)
	}
	var se = now.AddDate(0, month, 0)
	dur := se.Sub(now)
	allday := dur.Hours() / 24
	return int(allday)
}
