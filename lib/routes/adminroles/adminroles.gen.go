// Package adminroles provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package adminroles

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/cam-inc/viron-go/lib/routes/components"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

const (
	JwtScopes = "jwt.Scopes"
)

// Defines values for VironAdminRolePermissionPermission.
const (
	All   VironAdminRolePermissionPermission = "all"
	Deny  VironAdminRolePermissionPermission = "deny"
	Read  VironAdminRolePermissionPermission = "read"
	Write VironAdminRolePermissionPermission = "write"
)

// VironAdminRole defines model for VironAdminRole.
type VironAdminRole struct {
	// Id ロールID
	Id string `json:"id"`

	// Permissions 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// VironAdminRoleCreatePayload defines model for VironAdminRoleCreatePayload.
type VironAdminRoleCreatePayload struct {
	// Id ロールID
	Id string `json:"id"`

	// Permissions 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// VironAdminRoleList defines model for VironAdminRoleList.
type VironAdminRoleList = []VironAdminRole

// VironAdminRoleListWithPager defines model for VironAdminRoleListWithPager.
type VironAdminRoleListWithPager struct {
	CurrentPage int                `json:"currentPage"`
	List        VironAdminRoleList `json:"list"`
	MaxPage     int                `json:"maxPage"`
}

// VironAdminRolePermission defines model for VironAdminRolePermission.
type VironAdminRolePermission struct {
	Permission VironAdminRolePermissionPermission `json:"permission"`
	ResourceId string                             `json:"resourceId"`
}

// VironAdminRolePermissionPermission defines model for VironAdminRolePermission.Permission.
type VironAdminRolePermissionPermission string

// VironAdminRoleUpdatePayload defines model for VironAdminRoleUpdatePayload.
type VironAdminRoleUpdatePayload struct {
	// Permissions 権限
	Permissions []VironAdminRolePermission `json:"permissions"`
}

// ListVironAdminRolesParams defines parameters for ListVironAdminRoles.
type ListVironAdminRolesParams struct {
	// Size Size of list
	Size *externalRef0.VironPagerSizeQueryParam `form:"size,omitempty" json:"size,omitempty"`

	// Page Page number of list
	Page *externalRef0.VironPagerPageQueryParam `form:"page,omitempty" json:"page,omitempty"`
}

// CreateVironAdminRoleJSONRequestBody defines body for CreateVironAdminRole for application/json ContentType.
type CreateVironAdminRoleJSONRequestBody = VironAdminRoleCreatePayload

// UpdateVironAdminRoleJSONRequestBody defines body for UpdateVironAdminRole for application/json ContentType.
type UpdateVironAdminRoleJSONRequestBody = VironAdminRoleUpdatePayload

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// list admin roles
	// (GET /viron/adminroles)
	ListVironAdminRoles(w http.ResponseWriter, r *http.Request, params ListVironAdminRolesParams)
	// create an admin role
	// (POST /viron/adminroles)
	CreateVironAdminRole(w http.ResponseWriter, r *http.Request)
	// delete an admin role
	// (DELETE /viron/adminroles/{id})
	RemoveVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam)
	// update an admin role
	// (PUT /viron/adminroles/{id})
	UpdateVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// list admin roles
// (GET /viron/adminroles)
func (_ Unimplemented) ListVironAdminRoles(w http.ResponseWriter, r *http.Request, params ListVironAdminRolesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// create an admin role
// (POST /viron/adminroles)
func (_ Unimplemented) CreateVironAdminRole(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// delete an admin role
// (DELETE /viron/adminroles/{id})
func (_ Unimplemented) RemoveVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam) {
	w.WriteHeader(http.StatusNotImplemented)
}

// update an admin role
// (PUT /viron/adminroles/{id})
func (_ Unimplemented) UpdateVironAdminRole(w http.ResponseWriter, r *http.Request, id externalRef0.VironIdPathParam) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListVironAdminRoles operation middleware
func (siw *ServerInterfaceWrapper) ListVironAdminRoles(w http.ResponseWriter, r *http.Request) {

	var err error

	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params ListVironAdminRolesParams

	// ------------- Optional query parameter "size" -------------

	err = runtime.BindQueryParameter("form", true, false, "size", r.URL.Query(), &params.Size)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "size", Err: err})
		return
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListVironAdminRoles(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) CreateVironAdminRole(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateVironAdminRole(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RemoveVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) RemoveVironAdminRole(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.VironIdPathParam

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RemoveVironAdminRole(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateVironAdminRole operation middleware
func (siw *ServerInterfaceWrapper) UpdateVironAdminRole(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "id" -------------
	var id externalRef0.VironIdPathParam

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, JwtScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateVironAdminRole(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/viron/adminroles", wrapper.ListVironAdminRoles)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/viron/adminroles", wrapper.CreateVironAdminRole)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/viron/adminroles/{id}", wrapper.RemoveVironAdminRole)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/viron/adminroles/{id}", wrapper.UpdateVironAdminRole)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RXz2skRRT+V4anxzIzu3rqk78uwYUdV1wPYZBK95tJhe6q2qrq7M4ODcaAqCAEDwtB",
	"L4K4iigiHjyo/8xoYv4LeVWT6R81yUwMAcHLsuFVvfe97/ve65oZpKrQSqJ0FpIZaG54gQ6N/6uOvf9Q",
	"GCW3syF3e0M6Q2EhIQHN3R4wkLxASEBkwMDgo1IYzCBxpkQGNt3DgtMNN9V0yjoj5ASqikUlhnyChv55",
	"u0QzXZbK0KZGaCcU1aR4T5bFLpqeGvdyYR2wAOcRXavxaD5BWIFASIcTNJdDeEc8vRICxdfUtuLpmtrV",
	"RdSz7Yu/lhVCPlA5ejWM0micQB8XWQxjfvTD/Oi3+dH3228C69LLQKMphLVCSRvfPf32u/OTYwLvsPDx",
	"Fw2OIYEX+jUl/QXAfhvdcJmYyizqcmP41FNaO2AnWKIJZLS8oHb3MXWUoZ39DYPc4ZBPc8Wz/zMR98hd",
	"1PH1gcVwViV/T9A8kxuTGfA8vz+GZOfqMivHBSrWVSlfYN8csu+2y5pPEzM1itppKBE5RrdiKMuCUhvk",
	"JMljIxwNaoaSppfneaNebSGDVpUmxe1s9SZrgm6cbUq+geLv6uwq6/8XfLzOwqsNErWSlsagdBRdtRwZ",
	"FPzJZcEOoouTrJU0xkb3hByrmDpOPPSMytH2uBY9qzEVY5Fy59tk4ISjnQyvHlBL/Vzs9hp3gMEBmuAv",
	"GGwNtu4Agycv0dfH+oFKlXThG7szA2LBJyYreYO3taB0mk+E5AFd+Io2/QcH7Vlfdur4Lo3+iHX6O/vx",
	"q7Pjj5ZL8s9fP/j7m+fzDz8/f/Yz/efw+fzws/nhH/PDE2AwMarUkARnhptkqrhukxh/uNepE+ZUaZRc",
	"C0jg5a3B1l3fndvzJugHOj2VgclkBhN0sUIrOwDWIfPeJWTWr5pLtlt9pL/Je4DW3b9J03nZkFAGrVbS",
	"htbvDgbh4eXt4pey1vnCh/19GxZY/aS43nKtt70fhTbB99/yc2UxLY1wU8/U/mMHyc6IYNqyKLiZLgzb",
	"Mb/jE6K2a0sSXyu7gZq/f3n68XGkZngHPOx6nUYfrXtdZdNbIqv9AKna+4bGsYp0u3NLUG4kVer76HHZ",
	"0OsquSoWT2R/JrIqKJijw7Va/vXJp+cnX0daPsBCHcRa3mw0mz9IVszSKzHYjakL3W5OHQNdrjf66Re/",
	"nD77KSInfPVvnZzbHpv242WjsbmJRKUvdw13V8vg7OJHWveDVo2qfwIAAP//5iYtIhQPAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "./components.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
