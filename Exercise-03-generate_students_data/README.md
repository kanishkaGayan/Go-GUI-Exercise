# Go Student Data Generator

A bulk test data generator built with Go for creating realistic student records in SQL format. When I was learning SQL, I needed a large dataset for testing queries, so I created this codebase to generate test user data efficiently.

## 🎯 Project Overview

This utility generates 1000 realistic student records with unique emails, names, and personal information. The output is a ready-to-use SQL file that can be directly imported into any MySQL/MariaDB database for testing SQL queries, practicing database operations, or populating development environments.

## ✨ Features

- **Bulk Data Generation**: Creates 1000 student records in seconds
- **Realistic Data**: Uses common first names, last names, and email patterns
- **Unique Email Generation**: Ensures no duplicate emails across all records
- **SQL Ready Output**: Generates complete SQL file with table creation and insert statements
- **Customizable Schema**: Easy to modify fields and data pools
- **Random Data Distribution**: Uses proper randomization for realistic data spread
- **Multiple Data Fields**: Supports FirstName, LastName, Gender, Email, Country, PostalCode
- **Commented Code**: Well-documented for easy customization

## 🛠️ Technologies Used

- **Go (Golang)**: Core programming language
- **Standard Library**: `fmt`, `math/rand`, `os`, `time`
- **SQL**: MySQL/MariaDB compatible syntax
- **File I/O**: Direct SQL file generation

## 📦 Installation & Setup

### Prerequisites
- Go 1.16 or higher

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/kanishkaGayan/Go-GUI-Exercise.git
   cd Go-GUI-Exercise/Exercise-03-generate_students_data
   ```

2. Run the generator:
   ```bash
   go run main.go
   ```

3. Import the generated SQL file:
   ```bash
   mysql -u username -p database_name < students_data.sql
   ```

## 🏗️ Project Structure

```
Exercise-03-generate_students_data/
├── main.go           # Main generator application
├── go.mod            # Go module definition
└── students_data.sql # Generated SQL output file
```

## 🚀 Usage

1. **Generate Data**: Run `go run main.go`
2. **Check Output**: The file `students_data.sql` will be created
3. **Import to Database**: Use the SQL file in your database
4. **Test Queries**: Start practicing SQL with 1000 realistic records

### Sample Generated Data:
```sql
CREATE TABLE Students (
    StudentID INT AUTO_INCREMENT PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Gender ENUM('Male', 'Female', 'Other'),
    Email VARCHAR(100) NOT NULL UNIQUE,
    Country VARCHAR(50),
    PostalCode VARCHAR(20)
);

INSERT INTO Students (FirstName, LastName, Gender, Email, Country, PostalCode) VALUES
('Mason', 'White', 'Male', 'Mason.White.48236@email.com', 'Australia', '70072'),
('Benjamin', 'Brown', 'Male', 'Benjamin.Brown.92497@email.com', 'Australia', '81596'),
...
```

## 🧠 What I Learned Through This Project

### 1. **Random Data Generation**
- Using `math/rand` for realistic randomization
- Seeding random generators with timestamps
- Creating diverse data distributions

### 2. **File I/O Operations**
- Creating and writing to files in Go
- Handling file operations safely with defer
- Formatting large text outputs efficiently

### 3. **Data Uniqueness Management**
- Implementing unique constraint handling
- Using maps for fast duplicate checking
- Ensuring data integrity across large datasets

### 4. **SQL Schema Design**
- Creating proper table structures
- Understanding data types and constraints
- Writing efficient INSERT statements

### 5. **String Manipulation & Formatting**
- Using `fmt.Sprintf` for dynamic content
- Building complex formatted strings
- Managing large string concatenations

### 6. **Data Structure Management**
- Working with slices and arrays
- Managing multiple data pools
- Implementing efficient data selection

### 7. **Performance Optimization**
- Generating large datasets efficiently
- Memory management for bulk operations
- Optimizing random selection algorithms

## 🔧 Code Highlights

- **Unique Email Generation**: Guarantees no duplicate emails using map tracking
- **Realistic Data Pools**: Curated lists of common names and countries
- **Flexible Schema**: Easy to modify fields and add new data types
- **Error Handling**: Proper file operation error management
- **Clean Output**: Well-formatted SQL ready for production use

## 🚀 Potential Enhancements

- Add command-line arguments for record count
- Support multiple output formats (JSON, CSV)
- Add more data fields (phone numbers, addresses)
- Implement different naming conventions
- Add database connection for direct insertion
- Create configuration file for data pools
- Add data validation and constraints
- Support for different SQL dialects

## ⚡ Use Cases

- **SQL Learning**: Practice SELECT, JOIN, UPDATE, DELETE queries
- **Database Testing**: Populate test environments quickly
- **Performance Testing**: Load testing with realistic data
- **Development**: Mock data for application development
- **Training**: Teaching database concepts with real-like data

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](../../issues).

## 👨‍💻 Author

**Kanishka Gayan**
- GitHub: [@kanishkaGayan](https://github.com/kanishkaGayan)

---

*This project was created as part of my SQL learning journey, providing realistic test data for database query practice and development.*
