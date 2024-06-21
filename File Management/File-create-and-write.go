package main

import (
	"fmt"
	"os"
)

func main() {
	createFile()
	readDirectory()
	deleteFile()
	readFile()
	fmt.Println("File does not exist:", fileExists("sample.txt"))
	appendContent()
}

// Create and Write content into File

func createFile() {
	file, err := os.Create("example.go")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	} else {
		fmt.Println("File created successfully!", file)
	}

	defer file.Close()

	content := "This is my first ever script!"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return
	} else {
		fmt.Println("Content written successfully!")
	}
}

// Reading Directory Contents

func readDirectory() {
	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory: ", err)
	} else {
		fmt.Println("Directory read successfully!\n", dir)
	}
}

// Delete File Content

func deleteFile() {
	err := os.Remove("sample.txt")
	if err != nil {
		fmt.Println("Error deleting file: ", err)
		return
	} else {
		fmt.Println("File deleted successfulyy!")
	}
}

// Reading from a file

func readFile() {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	} else {
		fmt.Println("File read sucessfully!\n", string(file))
	}
}

// Checking if a file exists

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)

}

// Appending content into a file

func appendContent() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return
	}
	defer file.Close()

	additionalContent := "\nIt makes me so happy."
	_, err = file.WriteString(additionalContent)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	} else {
		fmt.Println("Content appended successfully.")
	}
}

/*
Flags (os.O_APPEND|os.O_WRONLY):

The second argument specifies the flags that control how the file is opened. Multiple flags can be combined using the bitwise OR operator (|).
os.O_APPEND: This flag tells the system to append data to the end of the file when writing.
If the file exists, the file offset is positioned at the end of the file before each write.
os.O_WRONLY: This flag opens the file for write-only access.

*/
