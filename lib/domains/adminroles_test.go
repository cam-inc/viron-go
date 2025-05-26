package domains

import (
	"net/http"
	"testing"

	"github.com/cam-inc/viron-go/lib/errors"
	"github.com/stretchr/testify/assert"
)

func setUpRole() {

	if err := NewFile("./testdata/casbin.conf", "./testdata/casbin.csv"); err != nil {
		panic(err)
	}
}

func TestValidateRoleAndPermissions(t *testing.T) {
	tests := []struct {
		Title       string
		RoleId      string
		Permissions []*AdminRolePermission
		Expected    bool
		Err         *errors.VironError
	}{
		{
			Title:  "Role id with commas.",
			RoleId: "role_00001_,",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00001",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Role ID can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Resource id with commas.",
			RoleId: "role_00002",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00002_,",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Resource ID in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Permission with commas.",
			RoleId: "role_00003",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00003",
					Permission: "read_,",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Permission in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "No commas.",
			RoleId: "role_00004",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00004",
					Permission: "read",
				},
			},
			Err: nil,
		},
		{
			Title:  "Role id with double quotes.",
			RoleId: "role_00005\"",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00005",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Role ID can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Role id with single quotes.",
			RoleId: "role_00006'",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00006",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Role ID can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Resource id with double quotes.",
			RoleId: "role_00007",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00007\"",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Resource ID in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Resource id with single quotes.",
			RoleId: "role_00008",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00008'",
					Permission: "read",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Resource ID in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Permission with double quotes.",
			RoleId: "role_00009",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00009",
					Permission: "read\"",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Permission in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Permission with single quotes.",
			RoleId: "role_00010",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource_00010",
					Permission: "read'",
				},
			},
			Err: errors.Initialize(http.StatusBadRequest, "Permission in policy can only contain alphanumeric characters, hyphens, and underscores."),
		},
		{
			Title:  "Role id with hyphen and underscore and colon.",
			RoleId: "role-00011_",
			Permissions: []*AdminRolePermission{
				{
					ResourceID: "resource:-00011_",
					Permission: "read",
				},
			},
			Err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			if err := ValidateRoleAndPermissions(tt.RoleId, tt.Permissions); err != nil {
				assert.EqualError(t, err, tt.Err.Error())
			} else {
				assert.Equal(t, tt.Err, err)
			}
		})
	}
}
