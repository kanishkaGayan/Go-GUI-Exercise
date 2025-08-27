
# 🧮 Counter App

A simple, elegant counter application built with Go and the Fyne GUI toolkit. This project demonstrates fundamental GUI programming concepts in Go, including event handling, state management, and responsive UI design.

## 🎯 Project Overview

This counter application provides a clean, intuitive interface for incrementing a numerical counter. Built as a learning exercise to explore Go's GUI capabilities through the Fyne framework, it showcases modern cross-platform desktop application development.

## ✨ Features

- **Simple Counter Interface**: Clean, minimalist design with a counter display and increment button
- **Real-time Updates**: Instant visual feedback when the counter value changes
- **Cross-platform Compatibility**: Runs on Windows, macOS, and Linux
- **Responsive Design**: Fixed window size (300x200) for consistent user experience
- **Modern GUI**: Built with Fyne v2 for native look and feel

## 🛠️ Technologies Used

- **Go (Golang)**: Core programming language
- **Fyne v2**: Cross-platform GUI toolkit for Go
- **Go Modules**: Dependency management

## 📦 Installation & Setup

### Prerequisites
- Go 1.16 or higher
- Git

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/kanishkagayan/go-counter-app.git
   cd go-counter-app
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run counter.go
   ```

## 🏗️ Project Structure

```
Exercise-01-Counter-App/
├── counter.go      # Main application file
├── go.mod          # Go module definition
├── go.sum          # Dependency checksums
└── README.md       # Project documentation
```

## 🚀 Usage

1. Launch the application by running `go run counter.go`
2. A window titled "Counter App" will appear
3. Click the "Increment" button to increase the counter value
4. The counter display updates in real-time

## 🧠 What I Learned Through This Project

### 1. **Fyne GUI Framework Fundamentals**
- Understanding Fyne's architecture and component system
- Working with widgets (Labels, Buttons) and containers
- Managing window properties and sizing

### 2. **Go GUI Programming Concepts**
- Event-driven programming in Go
- Callback functions and closure usage
- State management in GUI applications

### 3. **Application Structure & Design**
- Organizing code for GUI applications
- Separation of UI logic and application logic
- Creating responsive and user-friendly interfaces

### 4. **Cross-Platform Development**
- Building applications that work across different operating systems
- Understanding platform-agnostic GUI development

### 5. **Go Module Management**
- Working with external dependencies
- Understanding `go.mod` and `go.sum` files
- Package importing and management

### 6. **Event Handling & State Management**
- Implementing button click handlers
- Managing application state (counter variable)
- Real-time UI updates based on state changes

### 7. **Container Layouts**
- Using VBox containers for vertical arrangement
- Understanding Fyne's layout system
- Creating structured UI components

## 🔧 Code Highlights

- **Clean Architecture**: Simple, readable code structure
- **Event Handling**: Proper implementation of button click events
- **State Management**: Efficient counter state updates
- **UI Composition**: Effective use of containers and widgets

## 🚀 Potential Enhancements

- Add decrement functionality
- Implement reset button
- Add keyboard shortcuts
- Include counter history
- Add custom themes
- Implement counter limits

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](../../issues).

## 👨‍💻 Author

**Your Name**
- GitHub: [@kanishkagayan](https://github.com/kanishkagayan)

---

*This project was created as part of my Go learning journey, focusing on GUI development and cross-platform application creation.*
