package dao

import (
	"fmt"

	"github.com/juankis/registry/api/src/models"
)

//InsertRegistry insert
func InsertRegistry(registry *models.Registry) error {
	res, err := Db.NamedExec(`INSERT INTO registry (name, second_name, email, cel, type_customer) VALUES (:name, :second_name, :email, :cel, :type_customer)`, &registry)
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

//GetRegistry Returns error and registry if exist
func GetRegistry(registry *models.Registry) error {
	err := Db.Get(registry, "SELECT * FROM `registry` where id = ? limit 1", registry.ID)
	if err != nil {
		fmt.Errorf("Error getting registry from database %v \n", err)
		return err
	}
	return nil
}

/*

 */

//UpdateRegistry update registry
func UpdateRegistry(registry *models.Registry) error {
	fmt.Printf("update registry : %+v", registry)
	_, err := Db.NamedExec(`UPDATE registry set name = :name,
											second_name = :second_name,
											cel= :cel,
											type_customer = :type_customer where id = :id`, &registry)
	if err != nil {
		fmt.Errorf("Error updating registry in database", err)
		return err
	}
	return nil
}

//DeleteRegistry delete registry
func DeleteRegistry(registry *models.Registry) error {
	_, err := Db.Exec("DELETE FROM registry WHERE id = ?", registry.ID)
	if err != nil {
		fmt.Errorf("Error deleting registry from database", err, nil)
		return err
	}
	return nil
}
