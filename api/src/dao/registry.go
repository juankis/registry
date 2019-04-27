package dao

import (
	"fmt"

	"github.com/juankis/registry/api/src/models"
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

//GetRegistryAll Returns error and registry if exist
func GetRegistryAll() ([]models.Registry, error) {
	var registry []models.Registry
	err := Db.Select(&registry, "SELECT * FROM `registry` ORDER BY id ASC")
	if err != nil {
		fmt.Errorf("Error getting registry from database %v \n", err)
		return registry, err
	}
	return registry, nil
}
