package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// for clear screen function
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

//here above function close

type Student struct {
	firstName  string
	secondName string
	mobile     string
	email      string
	password   string
}
type Admin struct {
	id         string
	firstName  string
	secondName string
	mobile     string
	email      string
	password   string
}
type Book struct {
	bookName   string
	autherName string
	bookPages  uint
	bookPrice  float64
}

var OpenDatabase string
var SystemPassword = "#465Ui0407"
var s Student
var a Admin
var b Book
var userChoice string
var flag bool = true

func main() {

	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	if flag == true {
		Developer()
	}

	fmt.Print("\nPress 1 for login")
	fmt.Print("\nPress 2 for signup")
	fmt.Print("\nPress 3 for close\n")
	fmt.Print("\nPress your choice:")
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1":
		userLoginFinction()
	case "2":
		userSignupFunction()
	case "3":
		closeFunction()
	default:
		fmt.Print("\nInvalid Input!")
		time.Sleep(2 * time.Second)
		clearScreen()
		main()
	}

}

func Developer() {
	fmt.Println("\nDeveloper: Mr. Umar\t\tLanguage: Golang/Go")
	fmt.Print("\n \t\tProcessing...")
	time.Sleep(10 * time.Second)
	clearScreen()
	flag = false
	main()
}

func closeFunction() {
	fmt.Print("\nThanks for using")
	time.Sleep(5 * time.Second)
	os.Exit(1)

}
func userLoginFinction() {
	clearScreen()
	var userChoice string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nPress 1 for Admin login")
	fmt.Print("\nPress 2 for Student login")
	fmt.Print("\nPress 3 for back\n")
	fmt.Print("\nPress your choice: ")
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1":
		adminLogin()
	case "2":
		studentLogin()
	case "3":
		clearScreen()
		time.Sleep(2 * time.Second)
		main()
	default:
		userLoginFinction()

	}
}

func adminLogin() {
	clearScreen()
	var userEmail string
	var userPassword string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nPlease enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Print("\nPlease enter your password: ")
	fmt.Scan(&userPassword)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT `adminEmail`, `adminPassword` FROM `adminlogin`")
	if err != nil {
		fmt.Println("Error in Databese fetching!", err)
	}
	defer row.Close()

	var dbEmail string
	var dbPassword string

	for row.Next() {
		row.Scan(&dbEmail, &dbPassword)
		if userEmail == dbEmail {
			if userPassword == dbPassword {
				fmt.Print("\n\nLogin Successfully")
				time.Sleep(2 * time.Second)
				adminLoginFunctionality()
			} else {
				fmt.Print("\n\nInvalid Password")
				time.Sleep(2 * time.Second)
				adminLogin()
			}

		}
	}
	fmt.Println("\nInvalid Email Address")
	time.Sleep(3 * time.Second)
	adminLogin()
}

func studentLogin() {
	clearScreen()
	var userEmail string
	var userPassword string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nPlease enter your email: ")
	fmt.Scan(&userEmail)
	fmt.Print("\nPlease enter your password: ")
	fmt.Scan(&userPassword)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT `stdEmail`, `stdPassword` FROM `studentlogin`")
	if err != nil {
		fmt.Println("Error in Databese fetching!", err)
	}
	defer row.Close()

	var dbEmail string
	var dbPassword string

	for row.Next() {
		row.Scan(&dbEmail, &dbPassword)
		if userEmail == dbEmail {
			if userPassword == dbPassword {
				fmt.Print("\n\nLogin Successfully")
				time.Sleep(2 * time.Second)
				studentLoginFunctionality()
			} else {
				fmt.Print("\n\nInvalid Password")
				time.Sleep(2 * time.Second)
				studentLogin()
			}

		}
	}
	fmt.Println("\nInvalid Email Address")
	time.Sleep(3 * time.Second)
	studentLogin()
}

func adminLoginFunctionality() {

	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Println("1. Add book information")
	fmt.Println("2. Display all books information")
	fmt.Println("3. List all book of given auther")
	fmt.Println("4. Delete a book record")
	fmt.Println("5. Exit")
	fmt.Print("\nEnter one of above: ")
	fmt.Scan(&userChoice)
	if userChoice == "1" {
		addBook()
	} else if userChoice == "2" {
		displayBook()
	} else if userChoice == "3" {
		listOfAutherBooks()
	} else if userChoice == "4" {
		deleteBook()
	} else if userChoice == "5" {
		userLoginFinction()
	} else {
		fmt.Println("Invalid Input!")
		time.Sleep(2 * time.Second)
		adminLoginFunctionality()
	}
}

func addBook() {
	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nEnter book name: ")
	fmt.Scan(&b.bookName)
	fmt.Print("\nEnter author name: ")
	fmt.Scan(&b.autherName)
	fmt.Print("\nEnter book pages: ")
	fmt.Scan(&b.bookPages)
	fmt.Print("\nEnter book price: ")
	fmt.Scan(&b.bookPrice)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	} else {
		fmt.Print("Connection establish successfullly\n")
	}
	quary := "INSERT INTO `books` (`bookName`, `bookAuther`, `bookPage`, `bookPrice`, `addDate`) VALUES (?, ?, ?, ?, ?);"
	insertedData, err := db.ExecContext(context.Background(), quary, b.bookName, b.autherName, b.bookPages, b.bookPrice, time.Now())
	if err != nil {
		fmt.Println("Error in Data entry", err)
	}
	systemID, err := insertedData.LastInsertId()
	if err != nil {
		fmt.Println("Error in ID returining", err)
	}
	fmt.Print("\n\nSystem generated id for this book is: ", systemID)

	fmt.Print("\n\nA new book added Successfully ")
	defer db.Close()
	time.Sleep(5 * time.Second)
	adminLoginFunctionality()
}
func displayBook() {

	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Println("\nOur Database have following books")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT `bookName`, `bookAuther`, `bookPage`, `bookPrice` FROM `books`")
	if err != nil {
		fmt.Println("Error in Databese fetching!", err)
	}
	defer row.Close()

	var bookName string
	var bookAuther string
	var bookPage string
	var bookPrice string

	for row.Next() {
		row.Scan(&bookName, &bookAuther, &bookPage, &bookPrice)
		fmt.Println("Name: ", bookName, "Auther: ", bookAuther, "Page: ", bookPage, "Price: ", bookPrice)
	}
	time.Sleep(15 * time.Second)
	adminLoginFunctionality()
}
func listOfAutherBooks() {

	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT `bookName`, `bookAuther`, `bookPage`, `bookPrice` FROM `books`")
	if err != nil {
		fmt.Println("Error in Databese fetching!", err)
	}
	defer row.Close()

	var autherName string
	var bookAuther string
	var bookName string
	var bookPage string
	var bookPrice string
	fmt.Printf("\nEnter auther name: ")
	fmt.Scan(&autherName)

	for row.Next() {
		row.Scan(&bookName, &bookAuther, &bookPage, &bookPrice)
		if bookAuther == autherName {
			fmt.Println("Name: ", bookName, "Page: ", bookPage, "Price: ", bookPrice)
		}
	}
	fmt.Print("\n\nNo book found with auther name: ", autherName)
	time.Sleep(15 * time.Second)
	adminLoginFunctionality()
}
func deleteBook() {
	clearScreen()
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
	if err != nil {
		fmt.Println("Error in Databese Connection!", err)
	}
	defer db.Close()
	row, err := db.Query("SELECT `bookID`, `bookName`, `bookAuther`, `bookPage`, `bookPrice` FROM `books`")
	if err != nil {
		fmt.Println("Error in Databese fetching!", err)
	}
	defer row.Close()

	var autherName string
	var bookAuther string
	var bookName string
	var bookPage string
	var bookPrice string
	var bookID int
	var allId []int
	fmt.Printf("\nEnter auther name: ")
	fmt.Scan(&autherName)

	for row.Next() {
		row.Scan(&bookID, &bookName, &bookAuther, &bookPage, &bookPrice)
		if bookAuther == autherName {
			fmt.Println("Id: ", bookID, "Name: ", bookName, "Page: ", bookPage, "Price: ", bookPrice)
			allId = append(allId, bookID)
		}
	}
	fmt.Println("\nTotal id of book for this autre are : ", allId)
	fmt.Println("Please select one of them: ")
	var id int
	fmt.Scan(&id)
	quary := "DELETE FROM `books` WHERE bookID = ?;"
	db.ExecContext(context.Background(), quary, id)
	fmt.Print("\n\nNo book found with auther name: ", autherName)
	time.Sleep(15 * time.Second)
	adminLoginFunctionality()
}

func studentLoginFunctionality() {
	clearScreen()
	fmt.Print("\nI am student login")
	time.Sleep(2 * time.Second)
	main()
}

func userSignupFunction() {
	clearScreen()
	var userChoice string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nPress 1 for Admin Signup")
	fmt.Print("\nPress 2 for Student Signup")
	fmt.Print("\nPress 3 for back\n")
	fmt.Print("\nPress your choice: ")
	fmt.Scan(&userChoice)
	switch userChoice {
	case "1":
		adminSignupFunction()
	case "2":
		studentSignupFinction()
	case "3":
		main()
	default:
		userSignupFunction()

	}
}

func studentSignupFinction() {

	clearScreen()
	var rePassword string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nEnter Your 1st name: ")
	fmt.Scan(&s.firstName)
	fmt.Print("\nEnter Your 2nd name: ")
	fmt.Scan(&s.secondName)
	fmt.Print("\nEnter Your mobile number: ")
	fmt.Scan(&s.mobile)
	fmt.Print("\nEnter Your email address: ")
	fmt.Scan(&s.email)
	fmt.Print("\nEnter Your pasword: ")
	fmt.Scan(&s.password)
	fmt.Print("\nReenter your password: ")
	fmt.Scan(&rePassword)

	if s.password == rePassword {
		fullName := s.firstName + " " + s.secondName
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
		if err != nil {
			fmt.Println("Error in Databese Connection!", err)
		} else {
			fmt.Print("Connection establish successfullly\n")
		}
		quary := "INSERT INTO `studentlogin` (`stdName`, `stdNumber`, `stdEmail`, `stdPassword`, `loginDate`) VALUES (?, ?, ?, ?, ?);"
		insertedData, err := db.ExecContext(context.Background(), quary, fullName, s.mobile, s.email, s.password, time.Now())
		if err != nil {
			fmt.Println("Error in Data entry", err)
		}
		systemID, err := insertedData.LastInsertId()
		if err != nil {
			fmt.Println("Error in ID returining", err)
		}
		fmt.Print("\n\nYour system generated id is: ", systemID)

		fmt.Print("\n\nYou are register Successfully ")
		defer db.Close()
		time.Sleep(5 * time.Second)
		main()
	} else {
		fmt.Println("Password not match ")
		time.Sleep(2 * time.Second)
		studentSignupFinction()
	}

}
func adminSignupFunction() {

	clearScreen()
	var rePassword, password string
	fmt.Println("\n*********###### WELLCOME TO E-LIBRARY ######*********")
	fmt.Print("\nEnter the system password:")
	fmt.Scan(&password)
	if password == SystemPassword {
		fmt.Print("\nEnter your id: ")
		fmt.Scan(&a.id)
		fmt.Print("\nEnter Your 1st name: ")
		fmt.Scan(&a.firstName)
		fmt.Print("\nEnter Your 2nd name: ")
		fmt.Scan(&a.secondName)
		fmt.Print("\nEnter Your mobile number: ")
		fmt.Scan(&a.mobile)
		fmt.Print("\nEnter Your email address: ")
		fmt.Scan(&a.email)
		fmt.Print("\nEnter Your pasword: ")
		fmt.Scan(&a.password)
		fmt.Print("\nReenter your password: ")
		fmt.Scan(&rePassword)

		if a.password == rePassword {
			fullName := a.firstName + " " + a.secondName
			db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/elibrarymanagement")
			if err != nil {
				fmt.Println("Error in Databese Connection!", err)
			} else {
				fmt.Print("Connection establish successfullly\n")
			}
			quary := "INSERT INTO `adminlogin` (`adminID`, `adminName`, `adminNumber`, `adminEmail`, `adminPassword`, `loginDate`) VALUES (?, ?, ?, ?, ?, ?)"
			insertedData, err := db.ExecContext(context.Background(), quary, a.id, fullName, a.mobile, a.email, a.password, time.Now())
			if err != nil {
				fmt.Println("Error in Data entry", err)
			}
			systemID, err := insertedData.LastInsertId()
			if err != nil {
				fmt.Println("Error in ID returining", err)
			}
			fmt.Print("\n\nYour system generated id is: ", systemID)

			fmt.Print("\n\nYou are register Successfully ")
			defer db.Close()
			time.Sleep(5 * time.Second)
			main()
		}
		fmt.Print("\nPassword not match")
		time.Sleep(2 * time.Second)
		adminSignupFunction()

	}

	fmt.Println("System Password not match ")
	time.Sleep(2 * time.Second)
	main()

}
