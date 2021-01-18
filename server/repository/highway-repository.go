package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewHighwayRepository() *GenericRepository {
	return NewRepository("state/highway-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
