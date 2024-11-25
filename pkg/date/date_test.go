package date

import (
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	// Test Case 1: Regular date with non-zero time
	inputDate := time.Date(2024, time.November, 25, 10, 30, 45, 0, time.UTC)
	newDate := NewDate(inputDate)
	expectedDate := time.Date(2024, time.November, 25, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}

	// Test Case 2: Midnight input date (time already at 00:00:00)
	inputDate = time.Date(2024, time.November, 25, 0, 0, 0, 0, time.UTC)
	newDate = NewDate(inputDate)
	expectedDate = time.Date(2024, time.November, 25, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}

	// Test Case 3: Leap year date (February 29th)
	inputDate = time.Date(2024, time.February, 29, 15, 45, 0, 0, time.UTC) // Leap year date
	newDate = NewDate(inputDate)
	expectedDate = time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}

	// Test Case 4: January 1st with non-zero time
	inputDate = time.Date(2024, time.January, 1, 23, 59, 59, 999999999, time.UTC)
	newDate = NewDate(inputDate)
	expectedDate = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}

	// Test Case 5: December 31st with non-zero time
	inputDate = time.Date(2024, time.December, 31, 22, 15, 30, 0, time.UTC)
	newDate = NewDate(inputDate)
	expectedDate = time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}

	// Test Case 7: A random date far in the future
	inputDate = time.Date(3000, time.March, 12, 5, 15, 0, 0, time.UTC)
	newDate = NewDate(inputDate)
	expectedDate = time.Date(3000, time.March, 12, 0, 0, 0, 0, time.UTC)

	if !newDate.date.Equal(expectedDate) {
		t.Errorf("Expected date %v, but got %v", expectedDate, newDate.date)
	}
}

func TestDateToString(t *testing.T) {
	// Test the DateToString method with a valid date
	inputDate := time.Date(2024, time.November, 25, 10, 30, 45, 0, time.UTC)
	newDate := NewDate(inputDate)
	result := newDate.DateToString(newDate.date)
	expectedResult := "2024-11-25"

	if result != expectedResult {
		t.Errorf("Expected string %s, but got %s", expectedResult, result)
	}
}
