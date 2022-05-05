package rbac

import (
	"strings"

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
	allowd, err := r.RBACManager.Check(principal.UserInfo.ID, params.ID, models.RolePermissionPermissionRead)
	if err != nil {
		errBody := newError("get permission info failed", err.Error())
		return rbacop.NewGetProjectRBACInternalServerError().WithPayload(errBody)
	}
	if !allowd {
		errBody := newError("permission denied")
		rbacop.NewGetProjectRBACForbidden().WithPayload(errBody)

		return rbacop.NewGetProjectRBACForbidden().WithPayload(errBody)
	}

	projectRole, err := r.RBACManager.GetProjectRoleByProjectID(params.ID)
	if err != nil {
		errBody := newError("get project role error, get err = ", err.Error())
		return rbacop.NewGetProjectRBACInternalServerError().WithPayload(errBody)
	}

	return rbacop.NewGetProjectRBACOK().WithPayload(projectRole)
}

// GetUserRBAC Get User RBAC.
func (r *RBACOperationImpl) GetUserRBAC(params rbacop.GetUserRBACParams, principal *models.AuthInfo) middleware.Responder {
	userRole, err := r.RBACManager.GetUserRoleByUserID(principal.UserInfo.ID)
	if err != nil {
		errBody := newError("get user role error, get err = ", err.Error())
		return rbacop.NewGetUserRBACInternalServerError().WithPayload(errBody)
	}

	return rbacop.NewGetUserRBACOK().WithPayload(userRole)
}

func newError(msg ...string) *models.Error {
	return &models.Error{Message: strings.Join(msg, " ")}
}
