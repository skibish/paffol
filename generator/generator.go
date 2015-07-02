package generator

import (
	"html"
	"unicode"

	"github.com/skibish/paffol/helpers"
	"github.com/skibish/paffol/librarian"
)

// Generator is structure to store all variables
type Generator struct {
	Input      string
	lowerCased string
	converted  string
}

// escape is a method to escape string
func (g *Generator) escape() string {
	return html.EscapeString(g.Input)
}

// toLowerCase is a method to convert input string to lower case
// and remove unnecessary characters (kind of string normalization)
func (g *Generator) toLowerCase() {
	for _, runeValue := range g.Input {
		g.lowerCased = g.lowerCased + string(unicode.ToLower(runeValue))
	}

	g.lowerCased = helpers.CleanString(g.lowerCased)
}

// convertCyrillicToLatin converts cyrillic symbols to latin
func (g *Generator) convertCyrillicToLatin() {
	for _, runeValue := range g.lowerCased {
		g.converted = g.converted + librarian.Dictionary[string(runeValue)]
	}
}

// mapToAssociations randomly create string from passed string
// using provided associations
func (g *Generator) mapToAssociations() string {
	var output string
	var blockSize int
	for _, runeValue := range g.converted {
		blockSize = len(librarian.Associations[string(runeValue)])
		if blockSize > 0 {
			output = output + librarian.Associations[string(runeValue)][helpers.Random(blockSize)]
		} else {
			output = output + string(runeValue)
		}
	}

	return output
}

// Generate returns string with generated password
func (g *Generator) Generate() string {
	g.escape()                   // Step 1 is to escape input
	g.toLowerCase()              // Step 2 is to convert string to lower case
	g.convertCyrillicToLatin()   // Step 3 is to convert cyrillic phrase to latin
	return g.mapToAssociations() // Step 4 is to generate random string from associations
}
