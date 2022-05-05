package rbac

import (
	"github.com/ez-deploy/ezdeploy/models"
	rbacop "github.com/ez-deploy/ezdeploy/restapi/operations/r_b_a_c"
	"github.com/go-openapi/runtime/middleware"
)

// RBACOperationImpl Handle rbac related requests.
type RBACOperationImpl struct {
	RBACManager RBACManager
}

// GetProjectRBAC Get Project RBAC.
func (r *RBACOperationImpl) GetProjectRBAC(params rbacop.GetProjectRBACParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

// GetUserRBAC Get User RBAC.
func (r *RBACOperationImpl) GetUserRBAC(params rbacop.GetUserRBACParams, principal *models.AuthInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
