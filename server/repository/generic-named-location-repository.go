package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/model"
	"sync"
)

type GenericNamedLocationRepository struct {
	mutex            sync.Mutex
	locations        []model.NamedLocation
	persist          model.FuncPersistNamedLocations
	connectionString string
}

func NewNamedLocationRepository(connectionString string, init model.FuncInitNamedLocations, persist model.FuncPersistNamedLocations) *GenericNamedLocationRepository {
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

func (h *GenericNamedLocationRepository) GetAll() []model.NamedLocation {
	h.mutex.Lock()
	result := h.locations
	h.mutex.Unlock()
	return result
}

func (h *GenericNamedLocationRepository) AddLocation(location model.NamedLocation) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.locations = append(h.locations, location)
	l := h.locations
	h.persist(h.connectionString, l)
}

func (h *GenericNamedLocationRepository) initLocations(locations []model.NamedLocation) {
	h.locations = locations
}
