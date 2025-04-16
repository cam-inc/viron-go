package adminusers

import (
	"net/http"

	"github.com/cam-inc/viron-go/logging"

	"github.com/cam-inc/viron-go/constant"

	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/cam-inc/viron-go/helpers"

	"github.com/cam-inc/viron-go/domains"

	externalRef0 "github.com/cam-inc/viron-go/routes/components"
)

type (
	adminuser struct{}
)

func (a *adminuser) ListVironAdminUsers(w http.ResponseWriter, r *http.Request, params ListVironAdminUsersParams) {

	conditions := &domains.AdminUserConditions{}
	if params.Size != nil {
		conditions.Size = int(*params.Size)
	}
	if params.Page != nil {
		conditions.Page = int(*params.Page)
	}
	if params.Id != nil {
		conditions.ID = string(*params.Id)
	}
	if params.Email != nil {
		conditions.Email = string(*params.Email)
	}
	if params.RoleId != nil {
		conditions.RoleID = string(*params.RoleId)
	}

	results, err := domains.ListAdminUser(r.Context(), conditions)

	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, err)
		return
	}

	pager := PagerToVironAdminUserListWithPager(results.CurrentPage, results.MaxPage, results.List)

	helpers.Send(w, http.StatusOK, pager)

}

func (a *adminuser) CreateVironAdminUser(w http.ResponseWriter, r *http.Request) {

	payload := VironAdminUserCreatePayload{}
	if err := helpers.BodyDecode(r, &payload); err != nil {
		helpers.SendError(w, err.StatusCode(), err)
		return
	}
	user := &domains.AdminUser{
		Email:    string(payload.Email),
		Password: &payload.Password,
	}
	created, err := domains.CreateAdminUser(r.Context(), user, constant.AUTH_TYPE_EMAIL)
	if err != nil {
		helpers.SendError(w, http.StatusInternalServerError, err)
		return
	}
	helpers.Send(w, http.StatusCreated, created)

}

func (a *adminuser) RemoveVironAdminUser(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam) {
	if err := domains.RemoveAdminUserById(r.Context(), string(id)); err != nil {
		helpers.SendError(w, err.StatusCode(), err)
		return
	}
	helpers.Send(w, http.StatusNoContent, nil)
}

func (a *adminuser) UpdateVironAdminUser(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam) {

	log := logging.GetDefaultLogger()

	payload := VironAdminUserUpdatePayload{}
	if err := helpers.BodyDecode(r, &payload); err != nil {
		helpers.SendError(w, err.StatusCode(), err)
		return
	}

	user := &domains.AdminUser{
		Password: payload.Password,
	}
	if payload.RoleIds != nil {
		user.RoleIDs = *payload.RoleIds
	}

	log.Debugf("payload %+v", payload)

	if err := domains.UpdateAdminUserByID(r.Context(), string(id), user); err != nil {
		helpers.SendError(w, err.StatusCode(), err)
		return
	}
	helpers.Send(w, http.StatusNoContent, nil)
}

func New() ServerInterface {
	return &adminuser{}
}

func PagerToVironAdminUserListWithPager(currentPage, maxPage int, users []*domains.AdminUser) VironAdminUserListWithPager {
	vironPager := VironAdminUserListWithPager{
		CurrentPage: currentPage,
		MaxPage:     maxPage,
		List:        VironAdminUserList{},
	}

	for _, adminUser := range users {
		createdAt := externalRef0.VironCreatedAt(adminUser.CreatedAt)
		updatedAt := externalRef0.VironUpdatedAt(adminUser.UpdatedAt)
		vironPager.List = append(vironPager.List, VironAdminUser{
			Email:     openapi_types.Email(adminUser.Email),
			Id:        adminUser.ID,
			RoleIds:   &adminUser.RoleIDs,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		})
	}

	return vironPager
}
