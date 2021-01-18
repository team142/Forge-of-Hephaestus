package persistence

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/team142/Forge-of-Hephaestus/server/model"
	"io/ioutil"
)

var RepositoryPersistToDisk model.FuncPersistNamedLocations = func(connectionString string, locations []model.NamedLocation) {
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

var RepositoryInitFromDisk model.FuncInitNamedLocations = func(connectionString string, initializeLocations func(locations []model.NamedLocation)) {
	b, err := ioutil.ReadFile(connectionString)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not read file", err))
		//TODO: what should know about it?
		return
	}
	result := make([]model.NamedLocation, 0)
	err = json.Unmarshal(b, &result)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("%s with error: %s", "could not write file", err))
		//TODO: what should know about it?
		return
	}
	initializeLocations(result)

}
