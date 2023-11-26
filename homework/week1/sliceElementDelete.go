package main

import "reflect"

func DeleteElementByIndex[T any](input []T, index int) []T {
	var updatedResult []T
	updatedResult = append(input[:index], input[index+1:]...)
	return updatedResult
}

func DeleteElementByValue[T any](input []T, value T) []T {
	var updatedResult []T
	for _, item := range input {
		if !reflect.DeepEqual(item, value) {
			updatedResult = append(updatedResult, item)
		}
	}
	return updatedResult
}

func DeleteElementByValueForComparableElements[T comparable](input []T, value T) []T {
	var updatedResult []T
	for _, item := range input {
		if item != value {
			updatedResult = append(updatedResult, item)
		}
	}
	return updatedResult
}
