package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewRefuelRepository() *GenericRepository {
	return NewRepository("state/refuel-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
