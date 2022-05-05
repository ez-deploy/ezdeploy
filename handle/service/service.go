package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/handle/k8s"
	"github.com/ez-deploy/ezdeploy/handle/rbac"
	"github.com/ez-deploy/ezdeploy/models"
	serviceop "github.com/ez-deploy/ezdeploy/restapi/operations/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/wuhuizuo/sqlm"
)

type ServiceOperationImpl struct {
	Tables *db.Tables
	*ServiceVersionManager
	K8SManager *k8s.K8SManager
}

// UpdateServiceVersion Update Service Version
// You Can Call this api to re deploy/ expose this service.
func (s *ServiceOperationImpl) UpdateServiceVersion(params serviceop.UpdateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
	// check permission.
	svc, err := s.getServiceByID(params.Body.ID)
	if err != nil {
		errBody := newError("get service failed", err.Error())
		return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
	}

	// Stop And Delete Service First.
	allowed, err := newRBACManager(s.Tables).Check(principal.UserInfo.ID, svc.ProjectID, models.RolePermissionPermissionDelete)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewUpdateServiceVersionForbidden().WithPayload(errBody)
	}

	// get project info.
	project, err := s.getProjectByServiceInfo(svc)
	if err != nil {
		errBody := newError("get project info by projectID failed", err.Error())
		return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
	}
	// get service version info.
	svcVersion, err := s.getServiceVersionByID(params.Body.ID, params.Body.VersionID)
	if err != nil {
		errBody := newError("get service version info failed", err.Error())
		return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
	}

	// Set Service Info.
	if params.Body.ExposeType == models.ServiceInfoExposeTypeNone {
		// Do Not Expose Service In K8S.
		err := s.K8SManager.DeleteService(params.HTTPRequest.Context(), project.Name, params.Body.Name)
		if err != nil && !errors.Is(err, k8s.ERRServiceNotFound) {
			errBody := newError("delete service failed", err.Error())
			return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
		}
	} else {
		// Expose Service In K8S.
		fmt.Println(project.Name, params.Body, svcVersion)
		err := s.K8SManager.SetService(context.Background(), project.Name, params.Body, svcVersion)
		if err != nil {
			errBody := newError("set service failed", err.Error())
			return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
		}
	}

	// Set Deployment Info.
	if !params.Body.Running {
		// Delete Deployment.
		err := s.K8SManager.DeleteDeployment(params.HTTPRequest.Context(), project.Name, params.Body.Name)
		if err != nil && !errors.Is(err, k8s.ERRDeploymentNotFound) {
			errBody := newError("delete deployment failed", err.Error())
			return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
		}
	} else {
		// Create Deployment.
		err := s.K8SManager.SetDeployment(params.HTTPRequest.Context(), project.Name, params.Body, svcVersion)
		if err != nil {
			errBody := newError("set deployment failed", err.Error())
			return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
		}
	}

	// Write Result Into DB.
	selector := sqlm.SelectorFilter{"id": params.Body.ID}
	updateFields := map[string]interface{}{
		"running":         params.Body.Running,
		"version_id":      params.Body.VersionID,
		"replica":         params.Body.Replica,
		"expose_type":     params.Body.ExposeType,
		"in_cluster_port": params.Body.InClusterPort,
		"node_port":       params.Body.NodePort,
		"update_at":       time.Now().Unix(),
	}

	if err := s.Tables.ServiceInfo.Update(selector, updateFields); err != nil {
		errBody := newError("update service failed", err.Error())
		return serviceop.NewUpdateServiceVersionInternalServerError().WithPayload(errBody)
	}

	return serviceop.NewUpdateServiceVersionOK().WithPayload(params.Body)
}

// CreateService Create Service.
// CreateService will only create service, not create service version.
// To Deploy and expose version, user should create service version first, and then call UpdateServiceVersion function.
func (s *ServiceOperationImpl) CreateService(params serviceop.CreateServiceParams, principal *models.AuthInfo) middleware.Responder {
	allowed, err := newRBACManager(s.Tables).Check(principal.UserInfo.ID, params.Body.ProjectID, models.RolePermissionPermissionWrite)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewCreateServiceInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewCreateServiceForbidden().WithPayload(errBody)
	}

	createSVC := params.Body
	createSVC.CreateAt = time.Now().Unix()
	createSVC.UpdateAt = createSVC.CreateAt
	createSVC.Running = false
	createSVC.Replica = 0
	createSVC.ExposeType = models.ServiceInfoExposeTypeNone
	createSVC.InClusterPort = 0
	createSVC.NodePort = 0
	createSVC.VersionID = 0

	id, err := s.Tables.ServiceInfo.Insert(createSVC)
	if err != nil {
		errBody := newError("create service failed", err.Error())
		return serviceop.NewCreateServiceInternalServerError().WithPayload(errBody)
	}
	createSVC.ID = id

	return serviceop.NewCreateServiceCreated().WithPayload(createSVC)
}

// DeleteService Delete Service.
func (s *ServiceOperationImpl) DeleteService(params serviceop.DeleteServiceParams, principal *models.AuthInfo) middleware.Responder {
	svc, err := s.getServiceByID(params.ID)
	if err != nil {
		errBody := newError("get service failed", err.Error())
		return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
	}

	// Stop And Delete Service First.
	allowed, err := newRBACManager(s.Tables).Check(principal.UserInfo.ID, svc.ProjectID, models.RolePermissionPermissionDelete)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewDeleteServiceForbidden().WithPayload(errBody)
	}

	// Stop Service in k8s.
	if svc.Running {
		if err := s.stopAndDeleteServiceInCluster(params.HTTPRequest.Context(), svc); err != nil {
			errBody := newError("stop service failed", err.Error())
			return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
		}
	}

	// Delete Service in database.
	if err := s.deleteServiceByID(params.ID); err != nil {
		errBody := newError("delete service failed", err.Error())
		return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
	}

	return serviceop.NewDeleteServiceOK()
}

/* ListService List Service by project ID, service ID, service name. */
func (s *ServiceOperationImpl) ListService(params serviceop.ListServiceParams, principal *models.AuthInfo) middleware.Responder {
	matchedServices, err := s.listService(params.ProjectID, params.ServiceID, params.ServiceName)
	if err != nil {
		errBody := newError("list service failed", err.Error())
		return serviceop.NewListServiceInternalServerError().WithPayload(errBody)
	}

	var allowToReadServices []*models.ServiceInfo
	for _, svc := range matchedServices {
		allowed, err := newRBACManager(s.Tables).Check(principal.UserInfo.ID, svc.ProjectID, models.RolePermissionPermissionRead)
		if err != nil {
			errBody := newError("get permission info failed", err.Error())
			return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
		}

		if allowed {
			allowToReadServices = append(allowToReadServices, svc)
		}
	}

	retBody := &serviceop.ListServiceOKBody{Services: allowToReadServices}
	return serviceop.NewListServiceOK().WithPayload(retBody)
}

/* UpdateServiceDescription Update Service Description */
func (s *ServiceOperationImpl) UpdateServiceDescription(params serviceop.UpdateServiceDescriptionParams, principal *models.AuthInfo) middleware.Responder {
	svc, err := s.getServiceByID(params.Body.ID)
	if err != nil {
		errBody := newError("get service failed", err.Error())
		return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
	}

	// Stop And Delete Service First.
	allowed, err := newRBACManager(s.Tables).Check(principal.UserInfo.ID, svc.ProjectID, models.RolePermissionPermissionDelete)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return serviceop.NewDeleteServiceInternalServerError().WithPayload(errBody)
	}
	if !allowed {
		errBody := newError("permission denied")
		return serviceop.NewDeleteServiceForbidden().WithPayload(errBody)
	}

	filter := sqlm.SelectorFilter{
		"id": params.Body.ID,
	}

	err = s.Tables.ServiceInfo.Update(filter, map[string]interface{}{"description": params.Body.Description})
	if err != nil {
		errBody := newError("update service description failed", err.Error())
		return serviceop.NewUpdateServiceDescriptionInternalServerError().WithPayload(errBody)
	}

	return serviceop.NewUpdateServiceDescriptionOK().WithPayload(params.Body)
}

func (s *ServiceOperationImpl) listService(projID, svcID *int64, svcName *string) ([]*models.ServiceInfo, error) {
	if projID == nil && svcID == nil && svcName == nil {
		return nil, errors.New("empty query, at least one query is reqired")
	}

	filter := sqlm.SelectorFilter{}
	if projID != nil {
		filter["project_id"] = *projID
	}
	if svcID != nil {
		filter["id"] = *svcID
	}
	if svcName != nil {
		filter["name"] = *svcName
	}

	getItems, err := s.Tables.ServiceInfo.List(filter, sqlm.ListOptions{AllColumns: true})
	if err != nil {
		return nil, err
	}

	var services []*models.ServiceInfo
	for _, item := range getItems {
		services = append(services, item.(*models.ServiceInfo))
	}

	return services, nil
}

// stopAndDeleteServiceInCluster call k8s api to stop and delete service in cluster.
func (s *ServiceOperationImpl) stopAndDeleteServiceInCluster(ctx context.Context, svc *models.ServiceInfo) error {
	projectInfo, err := s.getProjectByServiceInfo(svc)
	if err != nil {
		return errors.WithMessage(err, "get project info failed")
	}

	if svc.Running {
		if err := s.K8SManager.DeleteDeployment(ctx, projectInfo.Name, svc.Name); err != nil {
			return errors.WithMessage(err, "delete deployment failed")
		}
	}

	if svc.ExposeType != models.ServiceInfoExposeTypeNone {
		if err := s.K8SManager.DeleteService(ctx, projectInfo.Name, svc.Name); err != nil {
			return errors.WithMessage(err, "delete service failed")
		}
	}

	return nil
}

func (s *ServiceOperationImpl) getProjectByServiceInfo(svc *models.ServiceInfo) (*models.ProjectInfo, error) {
	retProject := &models.ProjectInfo{}
	err := s.Tables.ProjectInfo.Get(sqlm.SelectorFilter{"id": svc.ProjectID}, retProject)
	if err != nil {
		return nil, err
	}
	return retProject, nil
}

func (s *ServiceOperationImpl) getServiceVersionByID(serviceID, serviceVersionID int64) (*models.ServiceVersion, error) {
	retServiceVersion := &models.ServiceVersion{}
	err := s.Tables.ServiceVersion.Get(sqlm.SelectorFilter{"id": serviceVersionID, "service_id": serviceID}, retServiceVersion)
	if err != nil {
		return nil, err
	}
	return retServiceVersion, nil
}

func (s *ServiceOperationImpl) getServiceByID(serviceID int64) (*models.ServiceInfo, error) {
	filter := sqlm.SelectorFilter{"id": serviceID}

	getService := &models.ServiceInfo{}
	if err := s.Tables.ServiceInfo.Get(filter, getService); err != nil {
		return nil, err
	}

	return getService, nil
}

func (s *ServiceOperationImpl) deleteServiceByID(serviceID int64) error {
	filter := sqlm.SelectorFilter{"id": serviceID}
	return s.Tables.ServiceInfo.Delete(filter)
}

func newRBACManager(tables *db.Tables) *rbac.RBACManager {
	return &rbac.RBACManager{
		Tables: tables,
	}
}

func newError(msg ...string) *models.Error {
	return &models.Error{Message: strings.Join(msg, " ")}
}
