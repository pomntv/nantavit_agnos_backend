package handlers

import "testing"

func TestValidatePassword(t *testing.T) {
	// Test cases
	cases := []struct {
		password string
		expected int
	}{
		{"Abcdefg123", 0},            // Case 0: Valid password
		{"Abcdefg", 1},               // Case 1: Missing number
		{"abcdefg123", 1},            // Case 2: Missing uppercase letter
		{"ABCDEFG123", 1},            // Case 3: Missing lowercase letter
		{"Abcde555fga", 1},           // Case 4: Repeating character
		{"aA1", 1},                   // Case 5: Short password
		{"Abcd!fg12345678901234", 1}, // Case 6: Long password
		{"A11", 2},                   // Case 7: Short and Missing lowercase password
		{"a23", 2},                   // Case 8: Short and Missing uppercase password
		{"AAA", 4},                   // Case 7: Short and Missing number password and Repeating and lowercase
		{"111", 4},                   // Case 7: Short and Missing Repeating password and lowercase and uppercase

		{"aaaaaaaaaaaaaaaaaaaaaaaaa", 4}, // Case 9: Long and Missing uppercase and repeating password and Missing number
		{"AAAAAAAAAAAAAAAAAAAAAAAAA", 4}, // Case 10: Long and Missing uppercase and repeating
		{"a.1", 2},                       //Short and
	}

	// Iterate over test cases
	for _, c := range cases {
		// Call the function being tested
		result := validatePassword(c.password)
		// Check if the result matches the expected value
		if result != c.expected {
			t.Errorf("validatePassword(%q) == %d, expected %d", c.password, result, c.expected)
		}
	}
}
