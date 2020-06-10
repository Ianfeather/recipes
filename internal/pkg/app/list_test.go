package app

import (
	"recipes/internal/pkg/common"
	"testing"
)

func TestCombineIngredients(t *testing.T) {
	tests := []struct {
		name   string
		a      []common.Ingredient
		expect Ingredient
	}{
		{
			name: "no addition",
			a: []common.Ingredient{
				common.Ingredient{
					Name:     "mince",
					Quantity: "1",
					Unit:     "g",
				},
			},
			expect: Ingredient{
				Unit:     "g",
				Quantity: 1,
			},
		},
		{
			name: "simple addition",
			a: []common.Ingredient{
				common.Ingredient{
					Name:     "mince",
					Quantity: "1",
					Unit:     "g",
				},
				common.Ingredient{
					Name:     "mince",
					Quantity: "2",
					Unit:     "g",
				},
			},
			expect: Ingredient{
				Unit:     "g",
				Quantity: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recipes := []common.Recipe{common.Recipe{
				Ingredients: tc.a,
			}}
			result := CombineIngredients(recipes)
			if *result["mince"] != tc.expect {
				t.Errorf("expected %v but got %v", tc.expect, *result["mince"])
			}
		})
	}
}
