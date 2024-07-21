package application

import (
	"fmt"
)

type Recipe struct {
	Dependencies []string
	Factory      func(dependencies map[string]interface{}) (interface{}, error)
}

type Factory interface {
	RegisterRecipe(name string, recipe Recipe)
	Create(name string) (interface{}, error)
}

type factory struct {
	serviceLocator ServiceLocator
	recipes        map[string]Recipe
}

func NewFactory(locator ServiceLocator) Factory {
	return &factory{
		serviceLocator: locator,
		recipes:        make(map[string]Recipe),
	}
}

func (f *factory) RegisterRecipe(name string, recipe Recipe) {
	fmt.Printf("Registering recipe: %s\n", name)
	f.recipes[name] = recipe
}

func (f *factory) Create(name string) (interface{}, error) {
	recipe, exists := f.recipes[name]
	if !exists {
		return nil, fmt.Errorf("recipe %s not found", name)
	}

	dependencies := make(map[string]interface{})
	for _, depName := range recipe.Dependencies {
		dep, err := f.serviceLocator.Resolve(depName)
		if err != nil {
			return nil, err
		}
		dependencies[depName] = dep
	}

	fmt.Printf("Creating use case: %s\n", name)
	return recipe.Factory(dependencies)
}
