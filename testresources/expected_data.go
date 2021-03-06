package testresources

import (
	"github.com/shadev/recipescator/model"
	"time"
)

// SampleRecipes returns two sample recipes
func SampleRecipes() []*model.Recipe {
	recipes := []*model.Recipe{
		{
			Rid:      "123456789",
			Title:    "Tiramisu Bars",
			Rating:   5,
			Servings: 4,
			Tags:     []string{"Dessert", "Quick", "Vegan"},
			Ingredients: []model.Ingredient{
				{Name: "Cashew", Amount: "1 Cup"},
				{Name: "Cocoa", Amount: "0.5 Cup"},
			},
			Source:     model.Source{SourceType: model.COOKBOOK, Title: "Richa's Everyday Kitchen", Ref: "Page 214"},
			Time:       model.Time{Active: 10 * time.Minute, Inactive: 1 * time.Hour, Prep: 10 * time.Minute},
			PreparedOn: []string{"2018-12-25"},
			Comments:   []string{"Tastes great", "One of my favourite desserts"},
		},
		{
			Rid:      "abcdefgh",
			Title:    "Cholent",
			Rating:   5,
			Servings: 4,
			Tags:     []string{"Stew", "Quick", "Vegan"},
			Ingredients: []model.Ingredient{
				{Name: "TVP", Amount: "1 Cup"},
				{Name: "Onion", Amount: "1 large", Preparation: "Diced"},
			},
			Source:     model.Source{SourceType: model.INTERNET, Title: "Veganomicon", Ref: "https://postpunkkitchen.de/cholent"},
			Time:       model.Time{Active: 50 * time.Minute, Inactive: 40 * time.Minute, Prep: 20 * time.Minute},
			PreparedOn: []string{"2019-01-06"},
			Comments:   []string{"Great winter dish", "So yummy"},
		},
	}
	return recipes
}

// SampleRecipesAsInterface returns the two sample recipes as interface{}
func SampleRecipesAsInterface() []interface{} {
	recipes := SampleRecipes()
	return []interface{}{recipes[0], recipes[1]}
}

// NewRecipeToBeInserted returns a third recipe used for insertion tests
func NewRecipeToBeInserted() model.Recipe {
	return model.Recipe{
		Title:    "Fig and Shallot Tarte",
		Rating:   5,
		Servings: 4,
		Tags:     []string{"Snack", "Fancy", "Vegan"},
		Ingredients: []model.Ingredient{
			{Name: "Cashew", Amount: "1 Cup"},
			{Name: "Shallots", Amount: "4 large"},
		},
		Source:     model.Source{SourceType: model.COOKBOOK, Title: "Mississippi Vegan", Ref: "Page 82"},
		Time:       model.Time{Active: 30 * time.Minute, Inactive: 30 * time.Minute, Prep: 30 * time.Minute},
		PreparedOn: []string{"2019-01-22"},
		Comments:   []string{"Our favourite snack", "So fancy"},
	}
}
