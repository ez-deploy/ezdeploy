package project

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/handle/k8s"
	"github.com/ez-deploy/ezdeploy/handle/rbac"
	"github.com/ez-deploy/ezdeploy/models"
	projectop "github.com/ez-deploy/ezdeploy/restapi/operations/project"
	"github.com/go-openapi/runtime/middleware"
	"github.com/wuhuizuo/sqlm"
)

// ProjectOperationImpl impl restapi.ProjectOperation interface.
type ProjectOperationImpl struct {
	Tables     *db.Tables
	K8SManager *k8s.K8SManager
}

// CreateProject Create Project.
func (p *ProjectOperationImpl) CreateProject(params projectop.CreateProjectParams, principal *models.AuthInfo) middleware.Responder {
	projectInfo := params.Body
	if projectInfo.Name == "" {
		errBody := newError("project name is empty")
		return projectop.NewCreateProjectConflict().WithPayload(errBody)
	}
	projectInfo.OwnerID = principal.UserInfo.ID

	// create project in k8s
	err := p.K8SManager.CreateNamespace(params.HTTPRequest.Context(), params.Body.Name)
	if errors.Is(err, k8s.ERRNamespaceConflict) {
		errBody := newError("namespace", params.Body.Name, "is exist")
		return projectop.NewCreateProjectConflict().WithPayload(errBody)
	}

	// create project in db
	if err := p.createProject(params.Body); err != nil {
		errBody := newError("create project error, get err = ", err.Error())
		return projectop.NewCreateProjectInternalServerError().WithPayload(errBody)
	}

	// get project id.
	rbacRole := newDefaultRABCRoles(params.Body.ID, principal.UserInfo.ID)
	if err := newRBACManager(p.Tables).CreateRole(rbacRole); err != nil {
		errBody := newError("create rbac role error, get err = ", err.Error())
		return projectop.NewCreateProjectInternalServerError().WithPayload(errBody)
	}

	return projectop.NewCreateProjectCreated().WithPayload(params.Body)
}

// GetProject Get project by id.
func (p *ProjectOperationImpl) GetProject(params projectop.GetProjectParams, principal *models.AuthInfo) middleware.Responder {
	params.ID = int64(params.ID)
	getProject, err := p.getProjectByID(params.ID)
	if errors.Is(err, sql.ErrNoRows) {
		errBody := newError("project not found")
		return projectop.NewGetProjectNotFound().WithPayload(errBody)
	}
	if err != nil {
		errBody := newError("get project error, get err = ", err.Error())
		return projectop.NewGetProjectInternalServerError().WithPayload(errBody)
	}

	return projectop.NewGetProjectOK().WithPayload(getProject)
}

func (p *ProjectOperationImpl) getProjectByID(id int64) (*models.ProjectInfo, error) {
	retProject := &models.ProjectInfo{}
	err := p.Tables.ProjectInfo.Get(sqlm.SelectorFilter{"id": id}, retProject)
	if err != nil {
		return nil, err
	}
	return retProject, nil
}

func (p *ProjectOperationImpl) createProject(proj *models.ProjectInfo) (err error) {
	proj.CreateAt = time.Now().Unix()
	proj.UpdateAt = proj.CreateAt
	proj.ID = 0

	proj.ID, err = p.Tables.ProjectInfo.Insert(proj)

	return err
}

func newDefaultRABCRoles(projectID, ownerID int64) *models.RoleView {
	return &models.RoleView{
		Info: &models.RoleInfo{
			Role:      "admin",
			ProjectID: projectID,
		},
		// admin user has all permissions
		Permissions: []*models.RolePermission{
			{Permission: models.RolePermissionPermissionDelete},
			{Permission: models.RolePermissionPermissionDeploy},
			{Permission: models.RolePermissionPermissionRead},
			{Permission: models.RolePermissionPermissionWrite},
		},
		// owner is the admin user.
		Members: []*models.RoleMember{
			{
				UserID: ownerID,
			},
		},
	}
}

func newRBACManager(tables *db.Tables) *rbac.RBACManager {
	return &rbac.RBACManager{
		Tables: tables,
	}
}

func newError(msg ...string) *models.Error {
	return &models.Error{Message: strings.Join(msg, " ")}
}
