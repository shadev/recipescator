package model

import "time"

const (
	INTERNET SourceType = "INTERNET"
	COOKBOOK SourceType = "COOKBOOK"
)

type (
	Recipe struct {
		Rid         string
		Title       string
		Rating      uint
		Servings    uint
		Tags        []string
		PreparedOn  []time.Time
		Comments    []string
		Time        Time
		Source      Source
		Ingredients []Ingredient
	}

	SourceType string

	Time struct {
		Active   time.Duration
		Inactive time.Duration
		Prep     time.Duration
	}

	Source struct {
		SourceType SourceType
		Title      string
		Ref        string
	}

	Ingredient struct {
		Name        string
		Amount      string
		Preparation string
	}
)
