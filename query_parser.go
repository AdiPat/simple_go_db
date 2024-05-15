package main

import "strings"

type QueryOperation string 

const (
	INSERT QueryOperation = "INSERT"
	SELECT QueryOperation = "SELECT"
	UPDATE QueryOperation = "UPDATE"
	DELETE QueryOperation = "DELETE"
	CREATE QueryOperation = "CREATE"
	DROP QueryOperation = "DROP"
)

type Query struct {
	raw string
	operation QueryOperation
	table string
	columns []string
	values []string
}

func ParseQuery(query string) Query {
	// Split the query by space
	tokens := strings.Split(query, " ")
	// Get the operation
	operation := QueryOperation(tokens[0])
	// Get the table
	table := tokens[2]
	// Get the columns
	columns := strings.Split(tokens[3], ",")
	// Get the values
	values := strings.Split(tokens[5], ",")
	// Return the Query struct
	return Query{
		raw: query,
		operation: operation,
		table: table,
		columns: columns,
		values: values,
	}
}

