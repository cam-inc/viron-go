package auditlogs

import (
	"net/http"

	"github.com/cam-inc/viron-go/lib/constant"

	"github.com/cam-inc/viron-go/lib/helpers"

	"github.com/cam-inc/viron-go/lib/domains"
)

type (
	auditlogsImpl struct{}
)

//	func (user *VironUserIdQueryParam) stringPtr() *string {
//		if user == nil {
//			return nil
//		}
//		s := string(*user)
//		return &s
//	}
//
//	func (uri *VironRequestUriQueryParam) stringPtr() *string {
//		if uri == nil {
//			return nil
//		}
//		s := string(*uri)
//		return &s
//	}
//
//	func (method *VironRequestMethodQueryParam) stringPtr() *string {
//		if method == nil {
//			return nil
//		}
//		s := string(*method)
//		return &s
//	}
//
//	func (ip *VironSourceIpQueryParam) stringPtr() *string {
//		if ip == nil {
//			return nil
//		}
//		s := string(*ip)
//		return &s
//	}
//
//	func (status *VironStatusCodeQueryParam) intPtr() *int {
//		if status == nil {
//			return nil
//		}
//		i := int(*status)
//		return &i
//	}
func (params ListVironAuditlogsParams) convertToDomainsAuditLog() *domains.AuditLog {
	return &domains.AuditLog{
		UserId:        params.UserId,
		RequestUri:    params.RequestUri,
		RequestMethod: params.RequestMethod,
		SourceIp:      params.SourceIp,
		StatusCode:    params.StatusCode,
	}
}
func (params ListVironAuditlogsParams) page() int {
	if params.Page == nil {
		return constant.DEFAULT_PAGER_PAGE
	}
	return *params.Page
}
func (params ListVironAuditlogsParams) size() int {
	if params.Size == nil {
		return constant.DEFAULT_PAGER_SIZE
	}
	return *params.Size
}
func (params ListVironAuditlogsParams) sort() []string {
	if params.Sort == nil {
		return []string{}
	}
	return *params.Sort
}

func (a auditlogsImpl) ListVironAuditlogs(w http.ResponseWriter, r *http.Request, params ListVironAuditlogsParams) {
	pager := domains.ListAuditLog(r.Context(), params.convertToDomainsAuditLog(), params.page(), params.size(), params.sort())
	helpers.Send(w, http.StatusOK, pager)
}

func New() ServerInterface {
	return &auditlogsImpl{}
}
