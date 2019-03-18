package dao

import (
	"fmt"

	"github.com/juankis/registry/api/src/models"
)

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
