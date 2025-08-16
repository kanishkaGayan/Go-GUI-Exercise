package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	calculatorApp := app.New()
	calculatorWindow := calculatorApp.NewWindow("Calculator")
	calculatorWindow.Resize(fyne.NewSize(400, 300))

	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Enter first number")

	operator := widget.NewSelect([]string{"+", "-", "*", "/"}, func(selected string) {
		// Handle operator selection
		
	})
	operator.PlaceHolder = "Select operator"

	input2 := widget.NewEntry()
	input2.SetPlaceHolder("Enter second number")

	result := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {
		calculatedResult := calculateResult(operator.Selected, input1.Text, input2.Text)
		result.SetText("Result: " + calculatedResult)
	})

	calculatorWindow.SetContent(container.NewVBox(input1, operator, input2, calculateButton, result))
	calculatorWindow.ShowAndRun()
}


func calculateResult(operator string, num1, num2 string) string {
	n1, err1 := strconv.ParseFloat(num1, 64)
	n2, err2 := strconv.ParseFloat(num2, 64)
	if err1 != nil || err2 != nil {
		return "Invalid input"
	}

	switch operator {
	case "+":
		return strconv.FormatFloat(n1+n2, 'f', -1, 64)
	case "-":
		return strconv.FormatFloat(n1-n2, 'f', -1, 64)
	case "*":
		return strconv.FormatFloat(n1*n2, 'f', -1, 64)
	case "/":
		if n2 == 0 {
			return "Cannot divide by zero"
		}
		return strconv.FormatFloat(n1/n2, 'f', -1, 64)
	default:
		return "Invalid operator"
	}
}

// The code implements a basic calculator GUI application using the Fyne framework.
// It allows users to input two numbers and select an operator to perform calculations.
// The result is displayed when the "Calculate" button is clicked.
// If the input is invalid, it shows an error message in the result label.
// The application is structured with a main function that initializes the GUI components and sets up the event

