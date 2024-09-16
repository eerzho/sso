package main

import (
	"context"
	"sso/internal/app"
	"sso/internal/model"
	"sso/internal/repo/mongo_repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	app := app.MustNew()

	ctx := context.TODO()
	permissionRepo := mongo_repo.NewPermission(app.Mng)
	roleRepo := mongo_repo.NewRole(app.Mng)
	userRepo := mongo_repo.NewUser(app.Mng)

	permissions := getAllPermissions()
	adminPermissions := make([]primitive.ObjectID, len(permissions))
	for index, permission := range permissions {
		err := permissionRepo.Create(ctx, &permission)
		if err != nil {
			panic(err)
		}
		adminPermissions[index] = permission.ID
	}

	adminRole := getAdminRole()
	adminRole.PermissionIDs = adminPermissions
	err := roleRepo.Create(ctx, adminRole)
	if err != nil {
		panic(err)
	}

	adminUser := getAdminUser()
	adminUser.RoleIDs = []primitive.ObjectID{adminRole.ID}
	err = userRepo.Create(ctx, adminUser)
	if err != nil {
		panic(err)
	}

	defaultUser := getDefaultUser()
	err = userRepo.Create(ctx, defaultUser)
	if err != nil {
		panic(err)
	}
}

func getDefaultUser() *model.User {
	return &model.User{
		Email:    "default@test.com",
		Name:     "default",
		Password: "password",
	}
}

func getAdminUser() *model.User {
	return &model.User{
		Email:    "admin@test.com",
		Name:     "admin",
		Password: "password",
	}
}

func getAdminRole() *model.Role {
	return &model.Role{
		Name: "Admin",
		Slug: "admin",
	}
}

func getAllPermissions() []model.Permission {
	permissionsReadPermission := model.Permission{
		Name: "Permission Read",
		Slug: "permission-read",
	}
	permissionsCreatePermission := model.Permission{
		Name: "Permission Create",
		Slug: "permission-create",
	}
	permissionsDeletePermission := model.Permission{
		Name: "Permission Delete",
		Slug: "permission-delete",
	}
	permissionsUpdatePermission := model.Permission{
		Name: "Permission Update",
		Slug: "permission-update",
	}

	usersReadPermission := model.Permission{
		Name: "User Read",
		Slug: "user-read",
	}
	usersCreatePermission := model.Permission{
		Name: "User Create",
		Slug: "user-create",
	}
	usersDeletePermission := model.Permission{
		Name: "User Delete",
		Slug: "user-delete",
	}
	usersUpdatePermission := model.Permission{
		Name: "User Update",
		Slug: "user-update",
	}

	rolesReadPermission := model.Permission{
		Name: "Role Read",
		Slug: "role-read",
	}
	rolesCreatePermission := model.Permission{
		Name: "Role Create",
		Slug: "role-create",
	}
	rolesDeletePermission := model.Permission{
		Name: "Role Delete",
		Slug: "role-delete",
	}
	rolesUpdatePermission := model.Permission{
		Name: "Role Update",
		Slug: "role-update",
	}

	return []model.Permission{
		permissionsReadPermission,
		permissionsCreatePermission,
		permissionsDeletePermission,
		permissionsUpdatePermission,
		usersReadPermission,
		usersCreatePermission,
		usersDeletePermission,
		usersUpdatePermission,
		rolesReadPermission,
		rolesCreatePermission,
		rolesDeletePermission,
		rolesUpdatePermission,
	}
}
