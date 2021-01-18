package repository

import (
	"github.com/team142/Forge-of-Hephaestus/server/model"
	"sync"
)

type GenericRepository struct {
	mutex            sync.Mutex
	locations        []model.NamedLocation
	persist          model.FuncPersist
	connectionString string
}

func NewRepository(connectionString string, init model.FuncInit, persist model.FuncPersist) *GenericRepository {
	result := &GenericRepository{
		persist:          persist,
		connectionString: connectionString,
	}
	init(result.connectionString, result.initLocations)
	return result
}

func (h *GenericRepository) GetConnectionString() string {
	return h.connectionString
}

func (h *GenericRepository) GetAll() []model.NamedLocation {
	h.mutex.Lock()
	result := h.locations
	h.mutex.Unlock()
	return result
}

func (h *GenericRepository) AddLocation(location model.NamedLocation) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.locations = append(h.locations, location)
	l := h.locations
	h.persist(h.connectionString, l)
}

func (h *GenericRepository) initLocations(locations []model.NamedLocation) {
	h.locations = locations
}
