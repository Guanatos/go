package main

import "fmt"
import "os"

// Main Function
func main(){
	file := createFile("dealers.txt")
	defer closeFile(file)
	writeToFile(file,"Esta es una prueba")
}

// Create a File
func createFile(path string) *os.File {
	fmt.Println("Creating file...")
	file,err := os.Create(path)
	if err != nil {
	   panic(err)
	}
	return file
}

// Write to a File
func writeToFile(file *os.File, dealerName string){
	fmt.Println("Writing to file...")
	fmt.Fprintln(file,dealerName)
}

// Close file
func closeFile(file *os.File){
	fmt.Println("Closing file...")
	file.Close()
}
