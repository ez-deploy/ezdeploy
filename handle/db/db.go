package db

import (
	"log"

	"github.com/ez-deploy/ezdeploy/models"
	"github.com/wuhuizuo/sqlm"
)

type Tables struct {
	User                *sqlm.Table
	Token               *sqlm.Table
	ProjectInfo         *sqlm.Table
	ServiceInfo         *sqlm.Table
	ServiceVersion      *sqlm.Table
	EnvironmentVariable *sqlm.Table
}

var (
	// Tables row models.
	userRowModel               = func() interface{} { return &models.UserInfo{} }
	tokenRowModel              = func() interface{} { return &models.Token{} }
	projectRowModel            = func() interface{} { return &models.ProjectInfo{} }
	serviceInfoRowModel        = func() interface{} { return &models.ServiceInfo{} }
	serviceVersionRowModel     = func() interface{} { return &models.ServiceVersion{} }
	enviromentVariableRowModel = func() interface{} { return &models.EnvironmentVariable{} }
)

// NewTables connect db and create Tables by dsn.
func NewTables(dsn string) (*Tables, error) {
	database := &sqlm.Database{
		Driver: sqlm.DriverMysql,
		DSN:    dsn,
	}
	if err := database.Create(); err != nil {
		return nil, err
	}

	retTables := &Tables{
		User:                newSqlmTableWithModeler(database, "user", userRowModel),
		Token:               newSqlmTableWithModeler(database, "token", tokenRowModel),
		ProjectInfo:         newSqlmTableWithModeler(database, "project_info", projectRowModel),
		ServiceInfo:         newSqlmTableWithModeler(database, "service_info", serviceInfoRowModel),
		ServiceVersion:      newSqlmTableWithModeler(database, "service_version", serviceVersionRowModel),
		EnvironmentVariable: newSqlmTableWithModeler(database, "enviroment_variable", enviromentVariableRowModel),
	}

	return retTables, nil
}

func newSqlmTableWithModeler(database *sqlm.Database, tableName string, modeler func() interface{}) *sqlm.Table {
	retTable := &sqlm.Table{
		Database:  database,
		TableName: tableName,
	}

	retTable.SetRowModel(modeler)

	if err := retTable.Create(); err != nil {
		log.Fatalln("create table", tableName, "error, get err = ", err)
	}

	return retTable
}
