package domain

//NamedLocations
type FuncInitNamedLocations func(connectionString string, initializeLocations func(locations []NamedLocation))
type FuncPersistNamedLocations func(connectionString string, locations []NamedLocation)

//Orchestrations
type FuncInitOrchestration func(connectionString string, initialize func(list []Orchestration))
type FuncPersistOrchestration func(connectionString string, list []Orchestration)
