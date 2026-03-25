package strutil

import (
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

var pluralClient = pluralize.NewClient()

// Capitalize returns the string with its first character uppercased.
// GraphQL type names are ASCII, so byte-level uppercasing is safe.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// PluralLower returns the lowercase-first plural form.
// e.g. "Movie" → "movies", "Status" → "statuses", "Category" → "categories".
func PluralLower(name string) string {
	if name == "" {
		return name
	}
	p := pluralClient.Plural(name)
	return strings.ToLower(p[:1]) + p[1:]
}

// PluralCapitalized returns the plural form with the first character capitalized.
// e.g. "Movie" → "Movies", "Status" → "Statuses".
func PluralCapitalized(name string) string {
	if name == "" {
		return name
	}
	return pluralClient.Plural(name)
}
