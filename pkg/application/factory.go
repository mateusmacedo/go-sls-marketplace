package application

import (
	"errors"
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
	f.recipes[name] = recipe
}

func (f *factory) Create(name string) (interface{}, error) {
	recipe, exists := f.recipes[name]
	if !exists {
		return nil, errors.New("recipe not found")
	}

	dependencies := make(map[string]interface{})
	for _, depName := range recipe.Dependencies {
		dep, err := f.serviceLocator.Resolve(depName)
		if err != nil {
			return nil, err
		}
		dependencies[depName] = dep
	}

	return recipe.Factory(dependencies)
}
