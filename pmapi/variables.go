package pmapi

import (
	"log"
	"net/http"
	"proxysql-master/admin/variables"

	"github.com/gin-gonic/gin"
)

//查询出所有变量的内容
func (pmapi *PMApi) ListPsVariables(c *gin.Context) {

	c.JSON(http.StatusOK, variables.GetProxySqlVariables(pmapi.Apidb))
}

func (pmapi *PMApi) UpdateOneVariables(c *gin.Context) {
	args := struct {
		VariableName  string `json:"variable_name"`
		VariableValue string `json:"variable_value"`
	}{}

	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": err})
	}

	psv := new(variables.Variables)
	log.Print("UpdateOneVariables", args)

	psv.VariablesName = args.VariableName
	psv.Value = args.VariableValue

	pret, _ := psv.UpdateOneVariable(pmapi.Apidb)
	if pret == 1 {
		c.JSON(http.StatusExpectationFailed, gin.H{"result": "UpdateOneVariable Failed"})
	}
	c.JSON(http.StatusOK, gin.H{"result": "OK"})
}
