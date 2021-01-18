package model

type FuncInit func(connectionString string, initializeLocations func(locations []NamedLocation))
type FuncPersist func(connectionString string, locations []NamedLocation)
