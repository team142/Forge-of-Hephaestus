package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewRefuelRepository() *GenericNamedLocationRepository {
	return NewNamedLocationRepository("state/refuel-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
