package service

import (
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	serviceop "github.com/ez-deploy/ezdeploy/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/wuhuizuo/sqlm"
)

type ServiceVersionManager struct {
	Tables *db.Tables
}

/* CreateServiceVersion Create Service Version */
func (s *ServiceVersionManager) CreateServiceVersion(params serviceop.CreateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	allowed, err := s.checkPermissionOnService(principal.UserInfo.ID, params.Body.ServiceID, models.RolePermissionPermissionWrite)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewCreateServiceVersionInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewCreateServiceVersionForbidden().WithPayload(errBody)
	}

	createSVC := params.Body
	createSVC.CreateAt = time.Now().Unix()

	id, err := s.Tables.ServiceVersion.Insert(createSVC)
	if err != nil {
		errBody := newError("create service version failed", err.Error())
		return serviceop.NewCreateServiceVersionInternalServerError().WithPayload(errBody)
	}
	createSVC.ID = id

	return serviceop.NewCreateServiceVersionCreated().WithPayload(createSVC)
}

/* ListServiceVersion List Service Version by service ID. */
func (s *ServiceVersionManager) ListServiceVersion(params serviceop.ListServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	allowrd, err := s.checkPermissionOnService(principal.UserInfo.ID, params.ServiceID, models.RolePermissionPermissionRead)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewListServiceVersionInternalServerError().WithPayload(errBody)
	}

	if !allowrd {
		errBody := newError("permission denied")
		return serviceop.NewListServiceVersionForbidden().WithPayload(errBody)
	}

	filter := sqlm.SelectorFilter{"service_id": params.ServiceID}
	getItems, err := s.Tables.ServiceVersion.List(filter, sqlm.ListOptions{AllColumns: true})
	if err != nil {
		errBody := newError("list service version failed", err.Error())
		return serviceop.NewListServiceVersionInternalServerError().WithPayload(errBody)
	}

	var serviceVersions []*models.ServiceVersion
	for _, item := range getItems {
		serviceVersions = append(serviceVersions, item.(*models.ServiceVersion))
	}

	respBody := &serviceop.ListServiceVersionOKBody{Versions: serviceVersions}
	return serviceop.NewListServiceVersionOK().WithPayload(respBody)
}

/* ListServiceVersion List Service Version by service ID. */
func (s *ServiceVersionManager) GetServiceVersion(params serviceop.GetServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	allowed, err := s.checkPermissionOnService(principal.UserInfo.ID, params.ServiceID, models.RolePermissionPermissionRead)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewGetServiceVersionInternalServerError().WithPayload(errBody)
	}

	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewGetServiceVersionForbidden().WithPayload(errBody)
	}

	filter := sqlm.SelectorFilter{"id": params.VersionID}
	getVersion := &models.ServiceVersion{}
	if err := s.Tables.ServiceVersion.Get(filter, getVersion); err != nil {
		errBody := newError("get service version failed", err.Error())
		return serviceop.NewGetServiceVersionInternalServerError().WithPayload(errBody)
	}

	return serviceop.NewGetServiceVersionOK().WithPayload(getVersion)
}

func (s *ServiceVersionManager) checkPermissionOnService(userID, svcID int64, permission string) (bool, error) {
	svc, err := s.getServiceByID(svcID)
	if err != nil {
		return false, err
	}

	// Stop And Delete Service First.
	allowed, err := newRBACManager(s.Tables).Check(userID, svc.ProjectID, permission)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

func (s *ServiceVersionManager) getServiceByID(serviceID int64) (*models.ServiceInfo, error) {
	filter := sqlm.SelectorFilter{"id": serviceID}

	getService := &models.ServiceInfo{}
	if err := s.Tables.ServiceInfo.Get(filter, getService); err != nil {
		return nil, err
	}

	return getService, nil
}
