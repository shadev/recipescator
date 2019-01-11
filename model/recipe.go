package model

import "time"

const (
	// INTERNET is the SourceType for all recipes found on the internet
	INTERNET SourceType = "INTERNET"

	// COOKBOOK is the SourceType for all recipes found in a cook-book
	COOKBOOK SourceType = "COOKBOOK"
)

type (
	// Recipe represents a recipe
	Recipe struct {
		Rid         string
		Title       string
		Rating      uint
		Servings    uint
		Tags        []string
		PreparedOn  []string
		Comments    []string
		Time        Time
		Source      Source
		Ingredients []Ingredient
	}

	// SourceType is one of the constants available
	SourceType string

	// Time defines the time it takes to complete this recipe
	Time struct {
		Active   time.Duration
		Inactive time.Duration
		Prep     time.Duration
	}

	// Source defines where to find this recipe
	Source struct {
		SourceType SourceType
		Title      string
		Ref        string
	}

	// Ingredient represents a single ingredient
	Ingredient struct {
		Name        string
		Amount      string
		Preparation string
	}
)
