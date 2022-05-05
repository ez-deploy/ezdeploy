package service

import (
	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	serviceop "github.com/ez-deploy/ezdeploy/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
)

type ServiceOperationImpl struct {
	Tables *db.Tables
	*ServiceVersionManager
}

/* CreateService Create Service */
func (s *ServiceOperationImpl) CreateService(params serviceop.CreateServiceParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* DeleteService Delete Service */
func (s *ServiceOperationImpl) DeleteService(params serviceop.DeleteServiceParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* ListService List Service by project ID, service ID, service name. */
func (s *ServiceOperationImpl) ListService(params serviceop.ListServiceParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* UpdateServiceDescription Update Service Description */
func (s *ServiceOperationImpl) UpdateServiceDescription(params serviceop.UpdateServiceDescriptionParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

/* UpdateServiceVersion Update Service Version */
func (s *ServiceOperationImpl) UpdateServiceVersion(params serviceop.UpdateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
