package application

import (
	"fmt"
	"testing"
)

type MockServiceLocator struct {
	services map[string]interface{}
	err      error
}

func NewMockServiceLocator(services map[string]interface{}, err error) *MockServiceLocator {
	return &MockServiceLocator{
		services: services,
		err:      err,
	}
}

func (m *MockServiceLocator) Resolve(name string) (interface{}, error) {
	if m.err != nil {
		return nil, m.err
	}
	service, exists := m.services[name]
	if !exists {
		return nil, fmt.Errorf("service %s not found", name)
	}
	return service, nil
}

func (m *MockServiceLocator) Register(name string, dependency interface{}) {
	m.services[name] = dependency
}

func TestFactory_Create(t *testing.T) {
	tests := []struct {
		name        string
		recipes     map[string]Recipe
		serviceMap  map[string]interface{}
		serviceErr  error
		recipeName  string
		expectedErr error
	}{
		{
			name: "Successful creation",
			recipes: map[string]Recipe{
				"MyService": {
					Dependencies: []string{"ServiceA", "ServiceB"},
					Factory: func(dependencies map[string]interface{}) (interface{}, error) {
						return fmt.Sprintf("Created with %v and %v", dependencies["ServiceA"], dependencies["ServiceB"]), nil
					},
				},
			},
			serviceMap: map[string]interface{}{
				"ServiceA": "A",
				"ServiceB": "B",
			},
			serviceErr:  nil,
			recipeName:  "MyService",
			expectedErr: nil,
		},
		{
			name:    "Recipe not found",
			recipes: map[string]Recipe{},
			serviceMap: map[string]interface{}{
				"ServiceA": "A",
				"ServiceB": "B",
			},
			serviceErr:  nil,
			recipeName:  "UnknownService",
			expectedErr: fmt.Errorf("recipe UnknownService not found"),
		},
		{
			name: "Service resolution error",
			recipes: map[string]Recipe{
				"MyService": {
					Dependencies: []string{"ServiceA", "ServiceB"},
					Factory: func(dependencies map[string]interface{}) (interface{}, error) {
						return fmt.Sprintf("Created with %v and %v", dependencies["ServiceA"], dependencies["ServiceB"]), nil
					},
				},
			},
			serviceMap: map[string]interface{}{
				"ServiceA": "A",
			},
			serviceErr:  fmt.Errorf("service not available"),
			recipeName:  "MyService",
			expectedErr: fmt.Errorf("service not available"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockLocator := NewMockServiceLocator(tt.serviceMap, tt.serviceErr)
			factory := NewFactory(mockLocator)

			for name, recipe := range tt.recipes {
				factory.RegisterRecipe(name, recipe)
			}

			result, err := factory.Create(tt.recipeName)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else {
					t.Logf("result: %v", result)
				}
			}
		})
	}
}
