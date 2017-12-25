package proxysql

import (
	"database/sql"
	"fmt"

	"github.com/juju/errors"
	//"fmt"
)

type (
	Variables struct {
		VariablesName string `db:"Variable_name" json:"variable_name"`
		Value         string `db:"Value" json:"variable_value"`
	}
)

const (
	StmtGlobalVariables   = `SHOW GLOBAL VARIABLES`
	StmtUpdateOneVariable = `
	UPDATE 
		global_variables 
	SET 
		variable_value=%q 
	WHERE variable_name = %q`
)

func UpdateOneConfig(db *sql.DB, var_name string, var_value string) error {
	st := fmt.Sprintf(StmtUpdateOneVariable, var_value, var_name)

	_, err := db.Exec(st)
	if err != nil {
		return errors.Trace(err)
	}

	LoadMySQLVariablesToRuntime(db)
	LoadAdminVariablesToRuntime(db)
	SaveMySQLVariablesToDisk(db)
	SaveAdminVariablesToDisk(db)
	return nil
}

func GetConfig(db *sql.DB) ([]Variables, error) {
	var tmparray []Variables
	var tmp Variables

	rows, err := db.Query(StmtGlobalVariables)
	if err != nil {
		return []Variables{}, errors.Trace(err)
	}

	for rows.Next() {
		tmp = Variables{}
		err = rows.Scan(&tmp.VariablesName, &tmp.Value)
		tmparray = append(tmparray, tmp)
	}

	return tmparray, nil
}
