package handle

import (
	"log"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/handle/project"
	"github.com/ez-deploy/ezdeploy/handle/user"

	_ "github.com/go-sql-driver/mysql"
)

// handerImpl impl restapi.Handler interface.
type handlerImpl struct {
	*ConfigurableImpl
	*user.UserOperationImpl
	*project.ProjectOperationImpl
}

const dsn = "kratos:123456@tcp(localhost:3306)/ezdeploy?charset=utf8mb4&parseTime=True"

func New() *handlerImpl {
	tables, err := db.NewTables(dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &handlerImpl{
		ConfigurableImpl:     &ConfigurableImpl{},
		UserOperationImpl:    &user.UserOperationImpl{Tables: tables},
		ProjectOperationImpl: &project.ProjectOperationImpl{Tables: tables},
	}
}
