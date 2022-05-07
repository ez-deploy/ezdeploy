package user

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/ez-deploy/ezdeploy/restapi/operations/identity"
	"github.com/go-openapi/runtime/middleware"
	"github.com/wuhuizuo/sqlm"
)

const tokenFieldName = "token"

// UserOperationImpl impl restapi.UserOperation and restapi.Authable interface.
type UserOperationImpl struct {
	Tables *db.Tables
}

/* GetUser Get User Info by id */
func (u *UserOperationImpl) GetUser(params identity.GetUserParams) middleware.Responder {
	userInfo := &models.UserInfo{}
	if err := u.Tables.User.Get(sqlm.SelectorFilter{"id": params.ID}, userInfo); err != nil {
		errorBody := newError("get user error, get err = ", err.Error())
		return identity.NewGetUserInternalServerError().WithPayload(errorBody)
	}

	return identity.NewGetUserOK().WithPayload(userInfo)
}

// CreateUser Create user.
func (u *UserOperationImpl) CreateUser(params identity.CreateUserParams) middleware.Responder {
	userInfo := params.Body

	if !isUserInfoValid(userInfo) {
		errBody := newError("user info invalid")
		return identity.NewCreateUserConflict().WithPayload(errBody)
	}

	isExists, err := u.isUserExist(userInfo.Email)
	if err != nil {
		errorBody := newError("check user exist error, get err = ", err.Error())
		return identity.NewCreateUserInternalServerError().WithPayload(errorBody)
	}
	if isExists {
		errBody := newError("user's email", userInfo.Email, "is exist")
		return identity.NewCreateUserConflict().WithPayload(errBody)
	}

	userInfo.ID = 0
	if _, err := u.Tables.User.Insert(userInfo); err != nil {
		errorBody := newError("insert user error, get err = ", err.Error())
		return identity.NewCreateUserInternalServerError().WithPayload(errorBody)
	}

	return identity.NewCreateUserCreated()
}

// Login 用户登录.
func (u *UserOperationImpl) Login(params identity.LoginParams) middleware.Responder {
	userInfo := params.Body

	selector := sqlm.SelectorFilter{"email": userInfo.Email, "password": userInfo.Password}
	err := u.Tables.User.Get(selector, userInfo)
	if err == sql.ErrNoRows {
		errBody := newError("username or password error")
		return identity.NewLoginUnauthorized().WithPayload(errBody)
	}
	if err != nil {
		errBody := newError("get user error, get err = ", err.Error())
		return identity.NewLoginInternalServerError().WithPayload(errBody)
	}

	userInfo.Password = ""

	newToken, err := newTokenManager(u.Tables).createToken(userInfo.ID, models.TokenTypeSession)
	if err != nil {
		errBody := newError("create token error, get err = ", err.Error())
		return identity.NewLoginInternalServerError().WithPayload(errBody)
	}

	setCookie := http.Cookie{
		Name:  tokenFieldName,
		Value: newToken.Value,
		Path:  "/",
	}

	return identity.NewLoginOK().WithSetCookie(setCookie.String()).WithPayload(userInfo)
}

// Applies when the "Cookie" header is set
func (u *UserOperationImpl) KeyAuth(value string) (*models.AuthInfo, error) {
	request := http.Request{Header: http.Header{}}
	request.Header.Add("Cookie", value)

	getTokenValue := false
	tokenValue := ""
	for _, cookie := range request.Cookies() {
		if cookie.Name != tokenFieldName {
			continue
		}

		getTokenValue = true
		tokenValue = cookie.Value
		break
	}

	if !getTokenValue {
		return nil, errors.New("token not found in cookie")
	}

	token, err := newTokenManager(u.Tables).getTokenByValue(tokenValue)
	if err != nil {
		return nil, err
	}

	if time.Now().Unix() > token.ExpiredAt {
		return nil, errors.New("token expired")
	}

	userInfo, err := u.getUserByID(token.UserID)
	if err != nil {
		return nil, err
	}
	userInfo.Password = ""

	return &models.AuthInfo{UserInfo: userInfo, Token: token}, nil
}

// Logout 用户登出.
func (u *UserOperationImpl) Logout(params identity.LogoutParams, principal *models.AuthInfo) middleware.Responder {
	if err := newTokenManager(u.Tables).deleteTokenByID(principal.Token.ID); err != nil {
		errBody := newError("delete token error, get err = ", err.Error())
		return identity.NewLogoutInternalServerError().WithPayload(errBody)
	}

	return identity.NewLogoutOK()
}

// Whoami 获取当前登录用户信息.
func (*UserOperationImpl) Whoami(params identity.WhoamiParams, principal *models.AuthInfo) middleware.Responder {
	principal.UserInfo.Password = ""
	return identity.NewWhoamiOK().WithPayload(principal.UserInfo)
}

func (u *UserOperationImpl) getUserByID(id int64) (*models.UserInfo, error) {
	userInfo := &models.UserInfo{}

	err := u.Tables.User.Get(sqlm.SelectorFilter{"id": id}, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
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

// isUserInfoValid check if user info is valid.
func isUserInfoValid(user *models.UserInfo) bool {
	fmt.Println(user, isPasswordValid(user.Password), isNameValid(user.UserName))
	return isEmailValid(user.Email) && isPasswordValid(user.Password) && isNameValid(user.UserName)
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	fmt.Println(email, err)
	return err == nil
}

func isPasswordValid(password string) bool {
	return len(password) > 6
}

func isNameValid(name string) bool {
	return len(name) > 0
}

func newTokenManager(tables *db.Tables) *TokenOperationImpl {
	return &TokenOperationImpl{Tables: tables}
}

func newError(msg ...string) *models.Error {
	return &models.Error{Message: strings.Join(msg, " ")}
}
