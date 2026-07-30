package main

import (
	"fmt"
	"sync"
)

// HealthState represents the persistence layer for node health
var (
	healthRegistry = make(map[string]bool)
	mu             sync.RWMutex
)

// SetNodeHealth updates the health status in the persistent registry
func SetNodeHealth(nodeID string, healthy bool) {
	mu.Lock()
	defer mu.Unlock()
	healthRegistry[nodeID] = healthy
}

// GetNodeHealth retrieves the health status from the persistent registry
func GetNodeHealth(nodeID string) bool {
	mu.RLock()
	defer mu.RUnlock()
	val, exists := healthRegistry[nodeID]
	if !exists {
		return true // Default to healthy if no state exists
	}
	return val
}

func main() {
	fmt.Println("APISIX Health Checker State Manager Initialized")
}