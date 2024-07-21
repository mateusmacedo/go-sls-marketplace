package application

import (
	"fmt"
	"testing"
)

func TestServiceLocator(t *testing.T) {
	tests := []struct {
		name          string
		registerCalls []struct {
			name       string
			dependency interface{}
		}
		resolveCalls []struct {
			name        string
			expected    interface{}
			expectedErr error
		}
	}{
		{
			name: "Successful registration and resolution",
			registerCalls: []struct {
				name       string
				dependency interface{}
			}{
				{"ServiceA", "A"},
				{"ServiceB", "B"},
			},
			resolveCalls: []struct {
				name        string
				expected    interface{}
				expectedErr error
			}{
				{"ServiceA", "A", nil},
				{"ServiceB", "B", nil},
			},
		},
		{
			name: "Dependency not found",
			registerCalls: []struct {
				name       string
				dependency interface{}
			}{},
			resolveCalls: []struct {
				name        string
				expected    interface{}
				expectedErr error
			}{
				{"UnknownService", nil, fmt.Errorf("dependency UnknownService not found")},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locator := NewSimpleServiceLocator()

			// Register dependencies
			for _, call := range tt.registerCalls {
				locator.Register(call.name, call.dependency)
			}

			// Resolve dependencies
			for _, call := range tt.resolveCalls {
				result, err := locator.Resolve(call.name)

				if call.expectedErr != nil {
					if err == nil || err.Error() != call.expectedErr.Error() {
						t.Errorf("expected error %v, got %v", call.expectedErr, err)
					}
				} else {
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if result != call.expected {
						t.Errorf("expected %v, got %v", call.expected, result)
					}
				}
			}
		})
	}
}
