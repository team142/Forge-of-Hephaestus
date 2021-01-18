package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"sync"
)

type GenericNamedLocationRepository struct {
	mutex            sync.Mutex
	locations        []domain.NamedLocation
	persist          domain.FuncPersistNamedLocations
	connectionString string
}

func NewNamedLocationRepository(connectionString string, init domain.FuncInitNamedLocations, persist domain.FuncPersistNamedLocations) *GenericNamedLocationRepository {
	result := &GenericNamedLocationRepository{
		persist:          persist,
		connectionString: connectionString,
	}
	init(result.connectionString, result.initLocations)
	return result
}

func (h *GenericNamedLocationRepository) GetConnectionString() string {
	return h.connectionString
}

func (h *GenericNamedLocationRepository) GetAll() []domain.NamedLocation {
	h.mutex.Lock()
	result := h.locations
	h.mutex.Unlock()
	return result
}

func (h *GenericNamedLocationRepository) AddLocation(location domain.NamedLocation) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.locations = append(h.locations, location)
	l := h.locations
	h.persist(h.connectionString, l)
}

func (h *GenericNamedLocationRepository) initLocations(locations []domain.NamedLocation) {
	h.locations = locations
}
