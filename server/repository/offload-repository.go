package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewOffloadRepository() *GenericNamedLocationRepository {
	return NewNamedLocationRepository("state/offload-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
