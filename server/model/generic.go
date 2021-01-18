package model

type FuncInitNamedLocations func(connectionString string, initializeLocations func(locations []NamedLocation))
type FuncPersistNamedLocations func(connectionString string, locations []NamedLocation)
