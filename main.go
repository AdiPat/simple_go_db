package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runCommand(command string) int {
	fmt.Println("You entered: ", command)
	return 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Simple Go DB. ")
	fmt.Println("1. SELECT DATABASE {database_name}")
	fmt.Println("2. CREATE DATABASE {database_name}")
	fmt.Println("3. CREATE TABLE {table_name}")
	fmt.Println("4. INSERT INTO {table_name} VALUES {value1, value2, ...}")
	fmt.Println("5. SELECT * FROM {table_name}")
	fmt.Println("6. DELETE FROM {table_name} WHERE {condition}")
	fmt.Println("7. UPDATE {table_name} SET {column_name} = {value} WHERE {condition}")
	fmt.Println("8. DROP DATABASE {database_name}")
	fmt.Println("9. DROP TABLE {table_name}")
	fmt.Println("10. EXIT")

	for {
		fmt.Println("Enter your command: ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		runCommand(command)
	}
}
