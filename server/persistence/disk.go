package persistence

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/team142/Forge-of-Hephaestus/server/domain"
	"io/ioutil"
	"strings"
)

var RepositoryPersistToDisk domain.FuncPersistNamedLocations = func(connectionString string, locations []domain.NamedLocation) {
	b, err := json.Marshal(locations)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not marshal to json", err))
		//TODO: what should know about it?
		return
	}
	err = ioutil.WriteFile(connectionString, b, 644)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not marshal to json", err))
		//TODO: what should know about it?
		return
	}

}

var RepositoryInitFromDisk domain.FuncInitNamedLocations = func(connectionString string, initializeLocations func(locations []domain.NamedLocation)) {
	b, err := ioutil.ReadFile(connectionString)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not read file", err))
		//TODO: what should know about it?
		return
	}
	result := make([]domain.NamedLocation, 0)
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not write file", err))
		//TODO: what should know about it?
		return
	}
	initializeLocations(result)

}

var RepositoryOrchestrationsPersistToDisk domain.FuncPersistOrchestration = func(connectionString string, list []domain.Orchestration) {
	panic("RepositoryOrchestrationsPersistToDisk not implemented")
}

var RepositoryOrchestrationsInitFromDisk domain.FuncInitOrchestration = func(connectionString string, initializeItems func([]domain.Orchestration)) {
	files, err := ioutil.ReadDir(connectionString)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.Println(">  files found:", len(files))
	items := make([]domain.Orchestration, 0)
	for _, f := range files {
		if !strings.Contains(f.Name(), ".json") {
			continue
		}
		path := fmt.Sprint(connectionString, "/", f.Name())
		logrus.Println("... loading", path)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not read file", err))
			//TODO: what should know about it?
			continue
		}
		result := domain.Orchestration{}
		err = json.Unmarshal(b, &result)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not read file", err))
			//TODO: what should know about it?
			continue
		}
		items = append(items, result)
	}
	initializeItems(items)
}
