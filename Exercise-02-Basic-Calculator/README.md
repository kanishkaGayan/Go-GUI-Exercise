# 🧮 Calculator App

A simple GUI calculator application built with Go and the Fyne framework. This calculator provides basic arithmetic operations with a clean, user-friendly interface.

## Features

- **Basic Arithmetic Operations**: Addition (+), Subtraction (-), Multiplication (*), Division (/)
- **GUI Interface**: Clean and intuitive graphical user interface using Fyne
- **Error Handling**: Handles invalid inputs and division by zero
- **Real-time Calculation**: Click the "Calculate" button to get instant results

## Screenshots
<img width="553" height="480" alt="image" src="https://github.com/user-attachments/assets/a207bb5a-aa5a-46ed-9bd4-87e673845fc0" />

The calculator features a simple vertical layout with:
- Two input fields for numbers
- A dropdown selector for operations
- A calculate button
- A result display area

## Prerequisites

- Go 1.19 or later
- Fyne framework dependencies (automatically handled by Go modules)

## Installation

1. **Clone or download the project**:
   ```bash
   git clone <repository-url>
   cd Exercise-02-Basic-Calculator
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run calculator.go
   ```

## Building the Application

To build an executable:

```bash
go build -o calculator calculator.go
```

To build for different platforms:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o calculator.exe calculator.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o calculator-mac calculator.go

# Linux
GOOS=linux GOARCH=amd64 go build -o calculator-linux calculator.go
```

## Usage

1. **Launch the application**
2. **Enter the first number** in the top input field
3. **Select an operation** from the dropdown menu (+, -, *, /)
4. **Enter the second number** in the second input field
5. **Click "Calculate"** to see the result

### Example Operations

- **Addition**: 5 + 3 = 8
- **Subtraction**: 10 - 4 = 6
- **Multiplication**: 7 * 6 = 42
- **Division**: 15 / 3 = 5

## Error Handling

The calculator handles several error conditions:

- **Invalid Input**: Non-numeric values will display "Invalid input"
- **Division by Zero**: Attempting to divide by zero shows "Cannot divide by zero"
- **Invalid Operator**: Unrecognized operators display "Invalid operator"

## Code Structure

### Main Components

- **GUI Setup**: Creates the main window and UI components
- **Input Fields**: Two entry widgets for number input
- **Operator Selection**: Dropdown widget for operation selection
- **Calculate Function**: Processes the calculation and displays results

### Key Functions

- `main()`: Initializes the GUI and sets up event handlers
- `calculateResult(operator, num1, num2)`: Performs the actual calculations with error handling

## Dependencies

The project uses the following Go modules:

- **fyne.io/fyne/v2**: GUI framework for creating the calculator interface
- **strconv**: For string to number conversions

## Project Structure

```
Exercise-02-Basic-Calculator/
├── calculator.go    # Main application code
├── go.mod          # Go module definition
├── go.sum          # Dependency checksums
└── README.md       # This file
```

## Technical Details

- **Framework**: Fyne v2.6.2
- **Language**: Go 1.24.4
- **Architecture**: Cross-platform GUI application
- **Window Size**: 400x300 pixels

## Learning Objectives

This exercise demonstrates:

- Building GUI applications in Go
- Using the Fyne framework for cross-platform development
- Handling user input and events
- String to number conversion and error handling
- Basic arithmetic operations implementation
- Application structure and organization

## Troubleshooting

### Common Issues

1. **Application won't start**:
   - Ensure Go is properly installed
   - Run `go mod tidy` to install dependencies

2. **Build errors**:
   - Check that you're using Go 1.19 or later
   - Verify all dependencies are installed

3. **GUI doesn't appear**:
   - Ensure your system supports GUI applications
   - On Linux, you may need X11 or Wayland support

## Future Enhancements

Potential improvements for this calculator:

- [ ] Memory functions (M+, M-, MR, MC)
- [ ] Scientific operations (sin, cos, tan, log, etc.)
- [ ] History of calculations
- [ ] Keyboard input support
- [ ] Copy/paste functionality
- [ ] Decimal precision settings
- [ ] Themes and customization options

## Contributing

Feel free to fork this project and submit pull requests for any improvements or bug fixes.

## License

This project is created for educational purposes as part of Go programming exercises.

---

**Note**: This is Exercise 02 in a series of Go programming exercises focused on building practical applications and learning GUI development with the Fyne framework.

