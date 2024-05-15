package main

import (
	"reflect"
	"testing"
)

func TestParseQuery(t *testing.T) {
    query := "INSERT INTO users name,email VALUES john, john@example.com"
    expected := Query{
        raw:       query,
        operation: "INSERT",
        table:     "users",
        columns:   []string{"name", "email"},
        values:    []string{"john", "john@example.com"},
    }

    result := ParseQuery(query)

    if result.raw != expected.raw {
        t.Errorf("Expected raw %v, but got %v", expected.raw, result.raw)
    }

    if result.operation != expected.operation {
        t.Errorf("Expected operation %v, but got %v", expected.operation, result.operation)
    }

    if result.table != expected.table {
        t.Errorf("Expected table %v, but got %v", expected.table, result.table)
    }

    if !reflect.DeepEqual(result.columns, expected.columns) {
        t.Errorf("Expected columns %v, but got %v", expected.columns, result.columns)
    }

    if !reflect.DeepEqual(result.values, expected.values) {
        t.Errorf("Expected values %v, but got %v", expected.values, result.values)
    }
}