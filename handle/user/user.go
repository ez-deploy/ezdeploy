package user

import (
	"database/sql"
	"errors"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/ez-deploy/ezdeploy/restapi/operations/identity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/wuhuizuo/sqlm"
)

// UserOperationImpl impl restapi.UserOperation and restapi.Authable interface.
type UserOperationImpl struct {
	Tables db.Tables
}

// CreateUser 新建用户.
func (u *UserOperationImpl) CreateUser(params identity.CreateUserParams) middleware.Responder {
	return nil
}

// Login 用户登录.
func (u *UserOperationImpl) Login(params identity.LoginParams) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

// Applies when the "Cookie" header is set
func (u *UserOperationImpl) KeyAuth(token string) (*models.UserInfo, error) {
	return nil, errors.New("not impl")
}

// Logout 用户登出.
func (u *UserOperationImpl) Logout(params identity.LogoutParams, principal *models.UserInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

// Whoami 获取当前登录用户信息.
func (*UserOperationImpl) Whoami(params identity.WhoamiParams, principal *models.UserInfo) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

// isUserExist check if user exist by email.
func (u *UserOperationImpl) isUserExist(email string) (bool, error) {
	userInfo := &models.UserInfo{}

	err := u.Tables.User.Get(sqlm.SelectorFilter{"email": email}, userInfo)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func newError(msg string) *models.Error {
	return &models.Error{Message: msg}
}
