package main

import (
	"errors"
	"reflect"
)

/**
 * GetIndex function
 *
 * This function takes in an array and an index and returns the value at that index.
 * If the array is not a slice or an array, it returns an error.
 * If the index is out of range, it returns an error.
 *
 */
func GetIndex(X interface{}, i int) (interface{}, error) {
	v := reflect.ValueOf(X)

	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, errors.New("variable is not a slice or array")
	}

	if i < 0 || i >= v.Len() {
		return nil, errors.New("index out of range")
	}

	return v.Index(i).Interface(), nil
}