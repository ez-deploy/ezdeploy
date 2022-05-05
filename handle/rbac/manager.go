package rbac

import (
	"github.com/ez-deploy/ezdeploy/handle/db"
	"github.com/ez-deploy/ezdeploy/models"
	"github.com/wuhuizuo/sqlm"
)

// RBACManager manage RBAC roles and permissions.
type RBACManager struct {
	Tables *db.Tables
}

// Check if user can do operation on project.
func (m *RBACManager) Check(userID, projectID int64, operation string) (bool, error) {
	projectRoles, err := m.GetProjectRoleByProjectID(projectID)
	if err != nil {
		return false, err
	}

	allow := false
	for _, role := range projectRoles.Roles {
		permisssionInRole := false
		for _, permisssion := range role.Permissions {
			if operation == permisssion.Permission {
				permisssionInRole = true
				break
			}
		}

		if !permisssionInRole {
			continue
		}

		// permisssionInRole.
		userInRole := false
		for _, member := range role.Members {
			if userID == member.UserID {
				userInRole = true
				break
			}
		}

		if !userInRole {
			continue
		}

		// user in this role, and permission is allowed.
		allow = true
		break
	}

	return allow, nil
}

// CreateRole and set ids.
func (m *RBACManager) CreateRole(role *models.RoleView) (err error) {
	if role.Info.ID, err = m.Tables.RoleInfo.Insert(role.Info); err != nil {
		return err
	}

	roleID := role.Info.ID

	for _, permission := range role.Permissions {
		permission.RoleID = roleID
		if permission.ID, err = m.Tables.RolePermission.Insert(permission); err != nil {
			return err
		}
	}

	for _, members := range role.Members {
		members.RoleID = roleID
		if members.ID, err = m.Tables.RoleMember.Insert(members); err != nil {
			return err
		}
	}

	return nil
}

func (m *RBACManager) GetProjectRoleByProjectID(projectID int64) (*models.ProjectRole, error) {
	listRoleFilter := &sqlm.SelectorFilter{"project_id": projectID}
	listRoleOptions := sqlm.ListOptions{AllColumns: true}

	listResults, err := m.Tables.RoleInfo.List(listRoleFilter, listRoleOptions)
	if err != nil {
		return nil, err
	}

	// all role ids.
	roleInfos := []*models.RoleInfo{}
	for _, result := range listResults {
		roleInfos = append(roleInfos, result.(*models.RoleInfo))
	}

	// all role views.
	roleViews := []*models.RoleView{}
	for _, roleMember := range roleInfos {
		roleView, err := m.GetRoleViewByRoleID(roleMember.ID)
		if err != nil {
			return nil, err
		}
		roleViews = append(roleViews, roleView)
	}

	return &models.ProjectRole{
		Roles: roleViews,
	}, nil
}

// GetRolesByUserID returns roles for a user.
func (m *RBACManager) GetUserRoleByUserID(userID int64) (*models.UserRole, error) {
	listRoleFilter := &sqlm.SelectorFilter{"user_id": userID}
	listRoleOptions := sqlm.ListOptions{AllColumns: true}

	listResults, err := m.Tables.RoleMember.List(listRoleFilter, listRoleOptions)
	if err != nil {
		return nil, err
	}

	// all role ids.
	roleMemberInfos := []*models.RoleMember{}
	for _, result := range listResults {
		roleMemberInfos = append(roleMemberInfos, result.(*models.RoleMember))
	}

	// all role views.
	roleViews := []*models.RoleView{}
	for _, roleMember := range roleMemberInfos {
		roleView, err := m.GetRoleViewByRoleID(roleMember.RoleID)
		if err != nil {
			return nil, err
		}
		roleViews = append(roleViews, roleView)

		// pass member infos.
		roleView.Members = []*models.RoleMember{}
	}

	return &models.UserRole{
		UserID: userID,
		Roles:  roleViews,
	}, nil
}

func (m *RBACManager) GetRoleViewByRoleID(roleID int64) (*models.RoleView, error) {
	getRole := &models.RoleInfo{}
	if err := m.Tables.RoleInfo.Get(&sqlm.SelectorFilter{"id": roleID}, getRole); err != nil {
		return nil, err
	}

	getPermissions, err := m.getPermissionsByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	getMembers, err := m.getMembersByRoleID(roleID)
	if err != nil {
		return nil, err
	}

	return &models.RoleView{
		Info:        getRole,
		Permissions: getPermissions,
		Members:     getMembers,
	}, nil
}

func (m *RBACManager) getPermissionsByRoleID(roleID int64) ([]*models.RolePermission, error) {
	getPermissionSelector := &sqlm.SelectorFilter{"role_id": roleID}
	getPermissionOpt := sqlm.ListOptions{AllColumns: true}

	getRolePermissions, err := m.Tables.RolePermission.List(getPermissionSelector, getPermissionOpt)
	if err != nil {
		return nil, err
	}

	retRolePermissions := []*models.RolePermission{}
	for _, result := range getRolePermissions {
		retRolePermissions = append(retRolePermissions, result.(*models.RolePermission))
	}

	return retRolePermissions, nil
}

func (m *RBACManager) getMembersByRoleID(roleID int64) ([]*models.RoleMember, error) {
	getMemberSelector := &sqlm.SelectorFilter{"role_id": roleID}
	getMemberOpt := sqlm.ListOptions{AllColumns: true}

	getMembers, err := m.Tables.RoleMember.List(getMemberSelector, getMemberOpt)
	if err != nil {
		return nil, err
	}

	retMembers := []*models.RoleMember{}
	for _, result := range getMembers {
		retMembers = append(retMembers, result.(*models.RoleMember))
	}

	return retMembers, nil
}
