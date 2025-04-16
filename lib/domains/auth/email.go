package auth

import (
	"net/http"

	"github.com/cam-inc/viron-go/helpers"

	"github.com/cam-inc/viron-go/errors"

	"github.com/cam-inc/viron-go/constant"
	"github.com/cam-inc/viron-go/domains"
)

// SigninEmail Emailアドレスでサインイン
func SigninEmail(r *http.Request, email string, password string) (string, *errors.VironError) {
	ctx := r.Context()
	user := domains.FindByEmail(ctx, email)
	log.Debugf("user(%s)", email)
	if user == nil {
		log.Debugf("user(%s) is nil", email)
		payload := &domains.AdminUser{
			Email:    email,
			Password: &password,
		}
		var err error
		user, _, err = createFirstAdminUser(ctx, payload, nil, constant.AUTH_TYPE_EMAIL)
		if err != nil || user == nil {
			return "", errors.SigninFailed
		}
	}

	if !helpers.VerifyPassword(password, *user.Password, *user.Salt) {
		return "", errors.SigninFailed
	}

	token, err := Sign(r, user.ID)
	if err != nil {
		log.Error("SigninEmail sign failed %#v \n", err)
		return "", errors.SigninFailed
	}
	return token, nil
}
