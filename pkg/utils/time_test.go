package utils

import (
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	tests := []struct {
		name        string
		dateString  string
		expected    time.Time
		expectError bool
	}{
		{
			name:        "valid date string",
			dateString:  "2023-11-25, 14:30:00",
			expected:    time.Date(2023, 11, 25, 14, 30, 0, 0, time.UTC),
			expectError: false,
		},
		{
			name:        "invalid date string",
			dateString:  "invalid-date",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToTime(tt.dateString)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !result.Equal(tt.expected) {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestTimeToString(t *testing.T) {
	input := time.Date(2023, 11, 25, 14, 30, 0, 0, time.UTC)
	expected := "2023-11-25:14:30:00"

	result := TimeToString(input)
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestIsOneYearGreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{
			name:     "exactly one year apart - false",
			date1:    time.Date(2022, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "more than one year apart - true",
			date1:    time.Date(2022, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 26, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "less than one year apart - false",
			date1:    time.Date(2022, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 24, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "same dates - false",
			date1:    time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "date1 after date2 - false",
			date1:    time.Date(2024, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "exactly one year minus a day apart - false",
			date1:    time.Date(2022, 11, 25, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 11, 24, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "date2 much later - true",
			date1:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOneYearGreaterThan(tt.date1, tt.date2)
			if result != tt.expected {
				t.Errorf("expected %v, got %v (date1: %v, date2: %v)", tt.expected, result, tt.date1, tt.date2)
			}
		})
	}
}

func TestGetFYYearString(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "date before July",
			input:    time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			expected: "2022-2023",
		},
		{
			name:     "date after July",
			input:    time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC),
			expected: "2023-2024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFYYearString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
