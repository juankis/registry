package dao

import (
	"fmt"

	"github.com/juankis/registry/api/src/models"
	"github.com/mercadolibre/goutils/logger"
)

//InsertRegistry insert
func InsertRegistry(registry *models.Registry) error {
	res, err := Db.NamedExec(`INSERT INTO registry (name, email, cel, type_customer) VALUES (:name, :email, :cel, :type_customer)`, &registry)
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Error inserting registry to the database, registry: %v\n", registry), err, nil)
		return err
	}
	id, _ := res.LastInsertId()
	registry.ID = int(id)
	return nil
}

//GetRegistryAll Returns error and checks if exist
func GetRegistryAll() ([]models.Check, error) {
	var checks []models.Check
	err := Db.Select(&checks, "SELECT * FROM `registry` ORDER BY id ASC")
	if err != nil {
		logger.Errorf("Error getting registry from database", err, nil)
		return checks, err
	}
	return checks, nil
}
