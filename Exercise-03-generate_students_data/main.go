package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
    // Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

    // Define data pools
	firstNamesMale := []string{
		"Liam", "Noah", "William", "James", "Oliver", "Benjamin", "Elijah", 
		"Lucas", "Mason", "Ethan", "Alexander", "Henry", "Jacob", "Michael", 
		"Daniel", "Logan", "Jackson", "Levi", "Sebastian", "Mateo",
	}

	// firstNamesFemale := []string{
	// 	"Emma", "Olivia", "Ava", "Isabella", "Sophia", "Charlotte", "Mia", 
	// 	"Amelia", "Harper", "Evelyn", "Abigail", "Emily", "Ella", "Elizabeth", 
	// 	"Camila", "Luna", "Sofia", "Avery", "Mila", "Scarlett",
	// }

	lastNames := []string{
		"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", 
		"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", 
		"Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin", 
		"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", 
		"Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King", 
		"Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green", 
		"Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", 
		"Carter", "Roberts",
	}

	// countries := []string{
	// 	"United States", "Canada", "United Kingdom", "Australia", "Germany", 
	// 	"France", "Japan", "Brazil", "India", "Italy", "Spain", "Mexico", 
	// 	"South Africa", "Netherlands", "South Korea",
	// }

	// Create a map to track used emails and ensure uniqueness
	usedEmails := make(map[string]bool)

    // Create and open the SQL file
	file, err := os.Create("students_data_3_col.sql")
	if err != nil {
		panic(err)
	}
	defer file.Close()

    // Write the table creation SQL
	file.WriteString("CREATE TABLE Students (\n")
	file.WriteString("    StudentID INT AUTO_INCREMENT PRIMARY KEY,\n")
	// file.WriteString("    FirstName VARCHAR(50) NOT NULL,\n")
	// file.WriteString("    LastName VARCHAR(50) NOT NULL,\n")
	// file.WriteString("    Gender ENUM('Male', 'Female', 'Other'),\n")
	file.WriteString("    Email VARCHAR(100) NOT NULL UNIQUE,\n")
	// file.WriteString("    Country VARCHAR(50),\n")
	// file.WriteString("    PostalCode VARCHAR(20)\n")
	file.WriteString(");\n\n")
	// file.WriteString("INSERT INTO Students (FirstName, LastName, Gender, Email, Country, PostalCode) VALUES\n")
	file.WriteString("INSERT INTO Students (FirstName,Email) VALUES\n")

    // Generate 1000 students
	for i := 0; i < 1000; i++ {
		//gender := "Male"
		firstNamePool := firstNamesMale
		
		// if rand.Intn(2) == 1 {
		// 	gender = "Female"
		// 	firstNamePool = firstNamesFemale
		// }

		firstName := firstNamePool[rand.Intn(len(firstNamePool))]
		lastName := lastNames[rand.Intn(len(lastNames))]
		//country := countries[rand.Intn(len(countries))]

		// Generate unique email
		var email string
		for {
			emailNum := rand.Intn(99999) + 1
			email = fmt.Sprintf("%s.%s.%d@email.com", firstName, lastName, emailNum)
			
			if !usedEmails[email] {
				usedEmails[email] = true
				break
			}
		}

		// Generate a simple postal code based on country
		//postalCode := fmt.Sprintf("%05d", rand.Intn(99999))
		// if country == "Canada" {
		// 	postalCode = fmt.Sprintf("%c%c%c %c%c%c", 
		// 		'A'+rune(rand.Intn(26)), 
		// 		'0'+rune(rand.Intn(10)), 
		// 		'A'+rune(rand.Intn(26)),
		// 		'0'+rune(rand.Intn(10)),
		// 		'A'+rune(rand.Intn(26)),
		// 		'0'+rune(rand.Intn(10)),
		// 	)
		// } else if country == "United Kingdom" {
		// 	postalCode = fmt.Sprintf("%c%c%c %c%c%c", 
		// 		'A'+rune(rand.Intn(26)), 
		// 		'A'+rune(rand.Intn(26)), 
		// 		'0'+rune(rand.Intn(10)),
		// 		'0'+rune(rand.Intn(10)),
		// 		'A'+rune(rand.Intn(26)),
		// 		'A'+rune(rand.Intn(26)),
		// 	)
		// }

		// Write the SQL value tuple
		value := fmt.Sprintf("('%s', '%s')", 
			firstName,email)
		
		if i < 999 {
			value += ","
		} else {
			value += ";"
		}
		
		file.WriteString(value + "\n")
	}

	fmt.Println("SQL file generated successfully: students_data.sql")
	fmt.Println("Total records: 1000")
	fmt.Println("Unique emails guaranteed:", len(usedEmails))
}