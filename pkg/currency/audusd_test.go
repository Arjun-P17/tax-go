package currency

import (
	"testing"
	"time"

	"github.com/Arjun-P17/tax-go/pkg/date"
	"github.com/stretchr/testify/assert"
)

func TestCreateCurrencyMap(t *testing.T) {
	tests := []struct {
		name    string
		rows    [][]string
		want    map[date.Date]float64
		wantErr bool // Whether we expect an error to be returned
	}{
		{
			name: "Valid input with sequential dates",
			rows: [][]string{
				{"Date", "Price"},      // Header
				{"01/01/2023", "0.75"}, // Day 1
				{"02/01/2023", "0.76"}, // Day 2
			},
			want: map[date.Date]float64{
				parseDate("01/01/2023", t): 0.75,
				parseDate("02/01/2023", t): 0.76,
			},
			wantErr: false,
		},
		{
			name: "Valid input with missing dates",
			rows: [][]string{
				{"Date", "Price"},      // Header
				{"01/01/2023", "0.75"}, // Day 1
				{"05/01/2023", "0.78"}, // Day 5
			},
			want: map[date.Date]float64{
				parseDate("01/01/2023", t): 0.75,
				parseDate("02/01/2023", t): 0.75,
				parseDate("03/01/2023", t): 0.75,
				parseDate("04/01/2023", t): 0.75,
				parseDate("05/01/2023", t): 0.78,
			},
			wantErr: false,
		},
		{
			name: "Malformed date in rows",
			rows: [][]string{
				{"Date", "Price"},       // Header
				{"01/01/2023", "0.75"},  // Valid row
				{"InvalidDate", "0.76"}, // Malformed date
				{"03/01/2023", "0.77"},  // Valid row
			},
			wantErr: true,
		},
		{
			name: "Malformed price in rows",
			rows: [][]string{
				{"Date", "Price"},              // Header
				{"01/01/2023", "0.75"},         // Valid row
				{"02/01/2023", "InvalidPrice"}, // Malformed price
			},
			wantErr: true,
		},
		{
			name:    "Empty input rows",
			rows:    [][]string{},
			want:    map[date.Date]float64{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := createCurrencyMap(tt.rows)

			if tt.wantErr {
				assert.Error(t, err, "Expected error, but got nil")
			} else {
				assert.NoError(t, err, "Did not expect error, but got")

				// Validate the resulting map
				if !equalCurrencyMaps(got, tt.want) {
					t.Errorf("createCurrencyMap() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// Helper function to parse dates and fail tests on errors
func parseDate(dateStr string, t *testing.T) date.Date {
	parsedDate, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		t.Fatalf("Failed to parse date %q: %v", dateStr, err)
	}
	return date.NewDate(parsedDate)
}

// Helper function to compare two maps of time.Time to float64
func equalCurrencyMaps(a, b map[date.Date]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
