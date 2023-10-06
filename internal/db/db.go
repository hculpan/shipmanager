package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hculpan/shipmanager/internal/data"
	"github.com/hculpan/shipmanager/internal/util"
)

func createDataDirectory() {
	path := "./data"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func LoadShip(id string) (*data.Ship, error) {
	jdata, err := os.ReadFile("./data/" + id + ".json")
	if err != nil {
		err := errors.New(fmt.Sprintf("File for ship id '%s' not found", id))
		util.LogError(err.Error())
		return nil, err
	}

	ship := &data.Ship{}
	err = json.Unmarshal(jdata, ship)
	return ship, err
}

func SaveShip(ship *data.Ship) error {
	createDataDirectory()

	json, err := json.Marshal(ship)
	if err != nil {
		return err
	}
	if err := os.WriteFile("./data/"+ship.Id+".json", json, 0644); err != nil {
		return err
	}

	return nil
}
