# Unit Converter (Go + Fyne)


A comprehensive unit converter application built with Go and the Fyne GUI framework. This application allows users to convert between various units of measurement across multiple categories.

[![Fyne.io Showcase](https://img.shields.io/badge/Fyne.io-Showcase-blue?logo=go&logoColor=white)](https://apps.fyne.io/apps/com.sprintora.Unit_Converter.html)
[![Go Version](https://img.shields.io/badge/Go-1.24.4-blue?logo=go&logoColor=white)](https://golang.org/)
[![Fyne](https://img.shields.io/badge/Fyne-v2.6.2-orange)](https://fyne.io/)

## 📱 Featured on Fyne.io

This application is officially listed on the [Fyne.io Apps Showcase](https://apps.fyne.io/apps/com.sprintora.Unit_Converter.html), demonstrating quality and design standards of the Fyne community.

## ✨ Features

### Supported Categories
- **Weight**: Kilogram, Gram, Milligram
- **Length**: Meter, Centimeter, Millimeter
- **Liquid**: Liter, Milliliter
- **Time**: Hour, Minute, Second
- **Temperature**: Celsius, Fahrenheit, Kelvin
- **Area**: Square Meter, Square Kilometer, Square Centimeter, Square Millimeter, Hectare, Acre
- **Energy**: Joule, Kilojoule, Calorie, Kilocalorie, Watt-hour, Kilowatt-hour
- **Power**: Watt, Kilowatt, Horsepower
- **Electric Current**: Ampere, Milliampere, Kiloampere
- **Data Storage**: Byte, Kilobyte, Megabyte, Gigabyte, Terabyte

### Key Features
- ✅ **Bidirectional Conversion**: Enter values in either field for instant conversion
- ✅ **Real-time Updates**: Conversions happen as you type
- ✅ **Smart Defaults**: Automatically selects sensible default units
- ✅ **High Precision**: Uses floating-point arithmetic for accurate conversions
- ✅ **Modern GUI**: Clean and intuitive interface built with Fyne
- ✅ **Cross-platform**: Works on Windows, macOS, and Linux
- ✅ **Lightweight**: Minimal resource usage

## 🚀 Quick Start

### Prerequisites
- Go 1.24.4 or higher
- C compiler (for Fyne dependencies)

### Installation

#### Option 1: Build from Source
```bash
# Clone the repository
git clone https://github.com/kanishkaGayan/Go-GUI-Exercise.git
cd Exercise-03-Unit-Converter-App

# Build the application
go build -o unitConverter unitConverter.go

# Run the application
./unitConverter
```

#### Option 2: Using Go Install
```bash
go install unitConverter
```

#### Option 3: System Installation (Linux)
```bash
# Install to system
sudo make install

# Or install to user directory
make user-install
```

### Running the Application
```bash
./unitConverter
```

## 🖥️ Screenshots

The application features a clean, modern interface with:
- Category selection dropdown
- From/To unit selection
- Bidirectional input fields
- Real-time conversion results

![Screenshot 1]<img width="548" height="685" alt="Screenshot from 2025-08-16 14-14-46" src="https://github.com/user-attachments/assets/43ae3be5-3d69-4ff0-99d3-dd46c60f4662" />

![Screenshot 2]<img width="550" height="687" alt="Screenshot from 2025-08-16 14-15-23" src="https://github.com/user-attachments/assets/c2e29255-2724-4aae-9ad2-4e7b19752cb1" />


## 🏗️ Architecture

### Core Components

1. **Conversion Engine** (`convert` function)
   - Handles all unit conversions using a standardized base unit approach
   - Supports 10 different measurement categories
   - Implements bidirectional conversion logic

2. **GUI Layer** (Fyne Framework)
   - Category selection with dynamic unit loading
   - Dual input fields for bidirectional conversion
   - Real-time update mechanism with feedback loop prevention

3. **Data Model**
   - Structured unit definitions in Go maps
   - Extensible design for adding new categories

## 🔧 Dependencies

- **Fyne v2**: Modern Go GUI framework
- **Standard Library**: strconv, fmt, net/url

```go
require fyne.io/fyne/v2 v2.6.2
```

## 💻 Development

### Building for Different Platforms

#### Windows
```bash
GOOS=windows GOARCH=amd64 go build -o unitConverter.exe unitConverter.go
```

#### macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o unitConverter unitConverter.go
```

#### Linux
```bash
GOOS=linux GOARCH=amd64 go build -o unitConverter unitConverter.go
```

### Package for Distribution
```bash
# Create application bundle
fyne package -os windows
fyne package -os darwin
fyne package -os linux
```

## 🤝 Contributing

Contributions are welcome! Here are some ways you can help:

1. **Add New Categories**: Extend the `units` map and `convert` function
2. **Improve UI/UX**: Enhance the interface design
3. **Add Features**: Implement history, favorites, or custom units
4. **Bug Fixes**: Report and fix any conversion errors
5. **Documentation**: Improve code comments and documentation

### Adding New Unit Categories

1. Add the category and units to the `units` map
2. Implement conversion logic in the `convert` function
3. Follow the base unit conversion pattern
4. Test thoroughly with known conversion values

## 📝 License

This project is open source. Please check the license file for details.

## 👨‍💻 Author

**Kanishka Meddegoda**  
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?logo=linkedin)](https://www.linkedin.com/in/kanishka-me/)

## 🙏 Acknowledgments

- Built with [Fyne](https://fyne.io/) - Cross-platform GUI framework for Go
- Featured on [Fyne.io Apps Showcase](https://apps.fyne.io/)
- Inspired by the need for a simple, fast unit converter

## 📊 Technical Details

- **Language**: Go 1.24.4
- **GUI Framework**: Fyne v2.6.2
- **Architecture**: Single binary, no external dependencies
- **Performance**: Near-instantaneous conversions
- **Memory Usage**: Minimal footprint
- **Platform Support**: Windows, macOS, Linux

---

*This application demonstrates the power and simplicity of Go combined with the Fyne framework for creating modern, cross-platform desktop applications.*
