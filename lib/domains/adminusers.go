package domains

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cam-inc/viron-go/lib/logging"

	"github.com/cam-inc/viron-go/lib/errors"

	"github.com/cam-inc/viron-go/lib/repositories/container"

	"github.com/cam-inc/viron-go/lib/helpers"

	"github.com/cam-inc/viron-go/lib/constant"
	"github.com/cam-inc/viron-go/lib/repositories"
)

type (
	AdminUser struct {
		ID        string    `json:"id"`
		Email     string    `json:"email"`
		Password  *string   `json:"password,omitempty"`
		Salt      *string   `json:"salt,omitempty"`
		RoleIDs   []string  `json:"roleIds"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	AdminUsersWithPager struct {
		Pager
		List []*AdminUser `json:"list"`
	}

	AdminUserConditions struct {
		ID     string
		Email  string
		RoleID string
		Size   int
		Page   int
		Sort   []string
	}
)

// CreateAdminUser adminUser insert
func CreateAdminUser(ctx context.Context, payload *AdminUser, authType string) (*AdminUser, *errors.VironError) {

	adminUser := &repositories.AdminUserEntity{}

	switch authType {
	case constant.AUTH_TYPE_EMAIL:
		adminUser.Email = string(payload.Email)
		if payload.Password == nil {
			return nil, errors.Initialize(http.StatusBadRequest, "password is nil.")
		}
		password := helpers.GenPassword(*payload.Password, "")
		adminUser.Password = &password.Password
		adminUser.Salt = &password.Salt
	case constant.AUTH_TYPE_OIDC:
		adminUser.Email = string(payload.Email)
	default:
		return nil, errors.Initialize(http.StatusBadRequest, fmt.Sprintf("authType %s is not supported.", authType))
	}

	entity, err := container.GetAdminUserRepository().CreateOne(ctx, adminUser)
	if err != nil {
		return nil, errors.Initialize(http.StatusInternalServerError, fmt.Sprintf("adminUser createOne %+v", err))
	}

	if err := entity.Bind(adminUser); err != nil {
		return nil, errors.Initialize(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	}

	payload.ID = adminUser.ID
	payload.Salt = adminUser.Salt
	payload.Password = adminUser.Password

	// Role update
	if len(payload.RoleIDs) > 0 {
		updateRolesForUser(payload.ID, payload.RoleIDs)
	}

	return payload, nil
}

// CountAdminUser adminUserレコード数をカウント
func CountAdminUser(ctx context.Context) int {
	repo := container.GetAdminUserRepository()
	return repo.Count(ctx, nil)
}

func findOne(ctx context.Context, conditions *repositories.AdminUserConditions) *AdminUser {
	repo := container.GetAdminUserRepository()
	result, err := repo.Find(ctx, conditions)
	if err != nil || len(result) == 0 {
		log.Errorf("adminusers.go findOne conditions:%+v err %+v result %+v", conditions, err, result)
		return nil
	}

	user := &repositories.AdminUserEntity{}

	if err := result[0].Bind(user); err != nil {
		log.Errorf("adminusers.go findOne bind failed err:%v", err)
		return nil
	}

	user.RoleIDs = listRoles(user.ID)

	auser := &AdminUser{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      user.Salt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RoleIDs:   user.RoleIDs,
	}
	return auser
}

// FindByEmail emailで1件取得
func FindByEmail(ctx context.Context, email string) *AdminUser {

	conditions := &repositories.AdminUserConditions{
		Email: email,
		Paginate: &repositories.Paginate{
			Size: 1,
			Page: 1,
		},
	}

	return findOne(ctx, conditions)
}

// FindByID IDで1件取得
func FindByID(ctx context.Context, userID string) *AdminUser {

	conditions := &repositories.AdminUserConditions{
		ID: userID,
		Paginate: &repositories.Paginate{
			Size: 1,
			Page: 1,
		},
	}

	return findOne(ctx, conditions)
}

// ListAdminUser 一覧取得
func ListAdminUser(ctx context.Context, opts *AdminUserConditions) (*AdminUsersWithPager, error) {

	repo := container.GetAdminUserRepository()

	conditions := &repositories.AdminUserConditions{}
	if opts != nil {
		conditions.ID = opts.ID
		conditions.Email = opts.Email
		conditions.Paginate = &repositories.Paginate{
			Sort: opts.Sort,
			Page: opts.Page,
			Size: opts.Size,
		}
	}
	if conditions.Page <= 0 {
		conditions.Page = constant.DEFAULT_PAGER_PAGE
	}
	if conditions.Size <= 0 {
		conditions.Size = constant.DEFAULT_PAGER_SIZE
	}

	results, err := repo.Find(ctx, conditions)
	if err != nil {
		return nil, err
	}

	withPager := &AdminUsersWithPager{
		List: []*AdminUser{},
	}

	for _, result := range results {
		entity := &repositories.AdminUserEntity{}
		if err := result.Bind(entity); err != nil {
			return nil, err
		}
		entity.RoleIDs = listRoles(entity.ID)
		adminuser := &AdminUser{
			ID:        entity.ID,
			Email:     entity.Email,
			Password:  entity.Password,
			Salt:      entity.Salt,
			RoleIDs:   entity.RoleIDs,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}

		withPager.List = append(withPager.List, adminuser)
	}
	count := CountAdminUser(ctx)
	pager := Paging(count, conditions.Size, conditions.Page)
	withPager.Pager = pager
	return withPager, nil
}

// UpdateAdminUserByID IDで1件更新
func UpdateAdminUserByID(ctx context.Context, id string, payload *AdminUser) *errors.VironError {
	user := FindByID(ctx, id)
	if user == nil {
		return errors.AdminUserNotfound
	}
	repo := container.GetAdminUserRepository()

	var entity *repositories.AdminUserEntity
	if user.Password != nil && payload.Password != nil {
		pass := helpers.GenPassword(*payload.Password, *user.Salt)
		if pass == nil {
			return errors.Initialize(http.StatusInternalServerError, "password gen failed.")
		}
		entity = &repositories.AdminUserEntity{
			ID:       user.ID,
			Email:    user.Email,
			Password: &pass.Password,
		}
	}

	if entity != nil {
		if err := repo.UpdateByID(ctx, id, entity); err != nil {
			return errors.Initialize(http.StatusInternalServerError, fmt.Sprintf("adminUser update failed. %+v", err))
		}
	}

	log := logging.GetDefaultLogger()
	log.Debugf("roleIds %+v", payload.RoleIDs)

	if len(payload.RoleIDs) > 0 {
		updateRolesForUser(id, payload.RoleIDs)
	}
	return nil
}

func RemoveAdminUserById(ctx context.Context, id string) *errors.VironError {
	// 存在チェック
	user := FindByID(ctx, id)
	if user == nil {
		return errors.AdminUserNotfound
	}

	// userを削除
	repoAdminUser := container.GetAdminUserRepository()
	if err := repoAdminUser.RemoveByID(ctx, id); err != nil {
		return errors.Initialize(http.StatusInternalServerError, fmt.Sprintf("adminUser delete failed. %+v", err))
	}

	// ssotokenをuserIdですべてのトークンを削除
	repoAdminUserSSOToken := container.GetAdminUserSSOTokenRepository()
	if err := repoAdminUserSSOToken.RemoveByID(ctx, id); err != nil {
		return errors.Initialize(http.StatusInternalServerError, fmt.Sprintf("adminUserSSOTokens delete failed. %+v", err))
	}

	// ユーザーからロールを剥奪
	if len(user.RoleIDs) > 0 {
		for _, role := range user.RoleIDs {
			RevokeRoleForUser(id, role)
		}
	}
	return nil
}
