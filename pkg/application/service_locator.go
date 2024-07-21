package application

import (
	"fmt"
	"sync"
)

type ServiceLocator interface {
	Register(name string, dependency interface{})
	Resolve(name string) (interface{}, error)
}

type serviceLocator struct {
	dependencies map[string]interface{}
	mu           sync.RWMutex
}

func NewSimpleServiceLocator() *serviceLocator {
	return &serviceLocator{
		dependencies: make(map[string]interface{}),
	}
}

func (sl *serviceLocator) Register(name string, dependency interface{}) {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	fmt.Printf("Registering dependency: %s\n", name)
	sl.dependencies[name] = dependency
}

func (sl *serviceLocator) Resolve(name string) (interface{}, error) {
	sl.mu.RLock()
	defer sl.mu.RUnlock()
	dependency, exists := sl.dependencies[name]
	if !exists {
		return nil, fmt.Errorf("dependency %s not found", name)
	}
	fmt.Printf("Resolving dependency: %s\n", name)
	return dependency, nil
}
