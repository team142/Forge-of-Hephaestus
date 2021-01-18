package repository

import (
	"github.com/sirupsen/logrus"
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"sync"
)

type OrchestrationRepository struct {
	mutex            sync.Mutex
	list             []domain.Orchestration
	persist          domain.FuncPersistOrchestration
	connectionString string
}

func NewOrchestrationRepository(connectionString string, init domain.FuncInitOrchestration, persist domain.FuncPersistOrchestration) *OrchestrationRepository {
	logrus.Infoln("> Loading orchestrations")
	result := &OrchestrationRepository{
		persist:          persist,
		connectionString: connectionString,
	}
	init(result.connectionString, result.initLocations)
	return result
}

func (h *OrchestrationRepository) GetConnectionString() string {
	return h.connectionString
}

func (h *OrchestrationRepository) GetAll() []domain.Orchestration {
	h.mutex.Lock()
	result := h.list
	h.mutex.Unlock()
	return result
}

func (h *OrchestrationRepository) initLocations(list []domain.Orchestration) {
	h.list = list
}
