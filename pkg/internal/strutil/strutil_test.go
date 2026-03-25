package strutil

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single lowercase char", "a", "A"},
		{"single uppercase char", "A", "A"},
		{"already capitalized", "Hello", "Hello"},
		{"all lowercase", "hello", "Hello"},
		{"multi-word", "hello world", "Hello world"},
		{"all uppercase", "HELLO", "HELLO"},
		{"numeric first char", "123abc", "123abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Capitalize(tt.input)
			if got != tt.expected {
				t.Errorf("Capitalize(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestPluralLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"standard type Movie", "Movie", "movies"},
		{"standard type Person", "Person", "people"},
		{"already lowercase", "movie", "movies"},
		{"ends in s", "Status", "statuses"},
		{"ends in x", "Index", "indices"},
		{"ends in ch", "Church", "churches"},
		{"ends in sh", "Wish", "wishes"},
		{"consonant+y", "Category", "categories"},
		{"multi-word", "MovieGenre", "movieGenres"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PluralLower(tt.input)
			if got != tt.expected {
				t.Errorf("PluralLower(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestPluralCapitalized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"standard type Movie", "Movie", "Movies"},
		{"ends in s", "Status", "Statuses"},
		{"ends in x", "Index", "Indices"},
		{"consonant+y", "Category", "Categories"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PluralCapitalized(tt.input)
			if got != tt.expected {
				t.Errorf("PluralCapitalized(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
