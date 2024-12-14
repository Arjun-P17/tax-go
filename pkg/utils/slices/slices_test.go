package slices

import "testing"

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		value  int
		expect bool
	}{
		{"Value exists", []int{1, 2, 3, 4, 5}, 3, true},
		{"Value does not exist", []int{1, 2, 3, 4, 5}, 6, false},
		{"Empty slice", []int{}, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.value)
			if got != tt.expect {
				t.Errorf("Contains() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		value  string
		expect bool
	}{
		{"Value exists", []string{"apple", "banana", "cherry"}, "banana", true},
		{"Value does not exist", []string{"apple", "banana", "cherry"}, "grape", false},
		{"Empty slice", []string{}, "banana", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.value)
			if got != tt.expect {
				t.Errorf("Contains() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestContainsEmptyValue(t *testing.T) {
	tests := []struct {
		name   string
		slice  []string
		value  string
		expect bool
	}{
		{"Empty value", []string{"apple", "banana", "cherry"}, "", false},
		{"Empty value in empty slice", []string{}, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.slice, tt.value)
			if got != tt.expect {
				t.Errorf("Contains() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		filter func(int) bool
		expect []int
	}{
		{"Even numbers", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 }, []int{2, 4}},
		{"Odd numbers", []int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 != 0 }, []int{1, 3, 5}},
		{"Empty slice", []int{}, func(v int) bool { return v%2 == 0 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.slice, tt.filter)
			if len(got) != len(tt.expect) {
				t.Errorf("Filter() = %v, want %v", got, tt.expect)
			}
			for i := range got {
				if got[i] != tt.expect[i] {
					t.Errorf("Filter() = %v, want %v", got, tt.expect)
				}
			}
		})
	}
}
