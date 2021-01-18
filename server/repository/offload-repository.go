package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewOffloadRepository() *GenericRepository {
	return NewRepository("state/offload-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
