package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &PizzaDecorator{}

	pizzaResult, _ := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"

	// strings.Contains will compare two parameters whether they are equal or not.
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator it must return the text %s the expected text, not '%s'", pizzaResult, expectedText)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &Onion{}
	onionResult, err := onion.AddIngredient()
	// because the onion must based on the core decorator : pizza.
	// must implement '	onion = &Onion{&PizzaDecorator{}}' first.
	// or it would return error.
	// if no error, that means your code is wrong!
	if err == nil {
		t.Errorf("When calling AddIngredient on the onion decorator withoout an IngredientAdd on its Ingredient field must return an error, not a string with '%s'", onionResult)
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it must return a text with the word 'onion, not '%s'", onionResult)
	}
}

func TestMead_AddIngredient(t *testing.T) {
	meat := &Meat{}
	meatResult, err := meat.AddIngredient()

	if err == nil {
		t.Errorf("When calling AddIngredient on the meat decorator withoout an IngredientAdd on its Ingredient field must return an error, not a string with '%s'", meatResult)
	}
	// if you did not use &PizzaDecorator{}
	// it will occur compile error
	// because you have to implement the IngredientAdd first.
	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()

	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the add ingredient of the onion decorator it must return a text with the word 'meat', not '%s'", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}

	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned string must contain the text '%s' nut '%s' didn't have it", expectedText, pizzaResult)
	}

	t.Log(pizzaResult)
}
