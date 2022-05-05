package project

import (
	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/ez-deploy/ezdeploy/restapi/operations/project"
	"github.com/go-openapi/runtime/middleware"
)

// ProjectOperationImpl impl restapi.ProjectOperation interface.
type ProjectOperationImpl struct {
	Tables *db.Tables
}

// CreateProject Create Project.
func (p *ProjectOperationImpl) CreateProject(params project.CreateProjectParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

// ListProject List All Projects.
func (p *ProjectOperationImpl) ListProject(params project.ListProjectParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
