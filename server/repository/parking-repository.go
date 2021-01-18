package repository

import "github.com/team142/Forge-of-Hephaestus/server/persistence"

func NewParkingRepository() *GenericRepository {
	return NewRepository("state/parking-repository.json", persistence.RepositoryInitFromDisk, persistence.RepositoryPersistToDisk)
}
