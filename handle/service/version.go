package service

import (
	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	serviceop "github.com/ez-deploy/ezdeploy/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

type ServiceVersionManager struct {
	Tables *db.Tables
}

/* CreateServiceVersion Create Service Version */
func (s *ServiceVersionManager) CreateServiceVersion(params serviceop.CreateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* ListServiceVersion List Service Version by service ID. */
func (s *ServiceVersionManager) ListServiceVersion(params serviceop.ListServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* ListServiceVersion List Service Version by service ID. */
func (s *ServiceVersionManager) GetServiceVersion(params serviceop.GetServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
