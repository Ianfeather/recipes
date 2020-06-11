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
					Unit:     "gram",
				},
			},
			expect: Ingredient{
				Unit:     "gram",
				Quantity: 1,
			},
		},
		{
			name: "simple addition",
			a: []common.Ingredient{
				common.Ingredient{
					Name:     "mince",
					Quantity: "1",
					Unit:     "gram",
				},
				common.Ingredient{
					Name:     "mince",
					Quantity: "2",
					Unit:     "gram",
				},
			},
			expect: Ingredient{
				Unit:     "gram",
				Quantity: 3,
			},
		},
		{
			name: "addition over threshold",
			a: []common.Ingredient{
				common.Ingredient{
					Name:     "mince",
					Quantity: "500",
					Unit:     "gram",
				},
				common.Ingredient{
					Name:     "mince",
					Quantity: "600",
					Unit:     "gram",
				},
			},
			expect: Ingredient{
				Unit:     "kilogram",
				Quantity: 1.1,
			},
		},
		{
			name: "addition over threshold",
			a: []common.Ingredient{
				common.Ingredient{
					Name:     "milk",
					Quantity: "500",
					Unit:     "millilitre",
				},
				common.Ingredient{
					Name:     "milk",
					Quantity: "600",
					Unit:     "millilitre",
				},
			},
			expect: Ingredient{
				Unit:     "litre",
				Quantity: 1.1,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			recipes := []common.Recipe{common.Recipe{
				Ingredients: tc.a,
			}}
			result := CombineIngredients(recipes)
			if *result[tc.a[0].Name] != tc.expect {
				t.Errorf("expected %v but got %v", tc.expect, *result[tc.a[0].Name])
			}
		})
	}
}
