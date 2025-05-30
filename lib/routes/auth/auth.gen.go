// Package auth provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package auth

import (
	"bytes"
	"compress/gzip"
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
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// OAuth2GoogleCallbackPayload defines model for OAuth2GoogleCallbackPayload.
type OAuth2GoogleCallbackPayload struct {
	// ClientId GoogleOAuth2クライアントID
	ClientId string `json:"clientId"`

	// Code Googleが発行した認可コード
	Code string `json:"code"`

	// RedirectUri GoogleOAuth2コールバックURI
	RedirectUri string `json:"redirectUri"`

	// State CSRF対策用のステートパラメータ
	State string `json:"state"`
}

// OidcCallbackPayload defines model for OidcCallbackPayload.
type OidcCallbackPayload struct {
	// ClientId OIDCクライアントID
	ClientId string `json:"clientId"`

	// Code OIDC Idpが発行した認可コード
	Code string `json:"code"`

	// RedirectUri OIDCコールバックURI
	RedirectUri string `json:"redirectUri"`

	// State CSRF対策用のステートパラメータ
	State string `json:"state"`
}

// SigninEmailPayload defines model for SigninEmailPayload.
type SigninEmailPayload struct {
	// Email Eメールアドレス
	Email openapi_types.Email `json:"email"`

	// Password パスワード
	Password string `json:"password"`
}

// ClientIdQueryParam defines model for ClientIdQueryParam.
type ClientIdQueryParam = string

// RedirectUriQueryParam defines model for RedirectUriQueryParam.
type RedirectUriQueryParam = string

// Oauth2GoogleAuthorizationParams defines parameters for Oauth2GoogleAuthorization.
type Oauth2GoogleAuthorizationParams struct {
	RedirectUri RedirectUriQueryParam `form:"redirectUri" json:"redirectUri"`
	ClientId    ClientIdQueryParam    `form:"clientId" json:"clientId"`
}

// OidcAuthorizationParams defines parameters for OidcAuthorization.
type OidcAuthorizationParams struct {
	RedirectUri RedirectUriQueryParam `form:"redirectUri" json:"redirectUri"`
	ClientId    ClientIdQueryParam    `form:"clientId" json:"clientId"`
}

// SigninEmailJSONRequestBody defines body for SigninEmail for application/json ContentType.
type SigninEmailJSONRequestBody = SigninEmailPayload

// Oauth2GoogleCallbackJSONRequestBody defines body for Oauth2GoogleCallback for application/json ContentType.
type Oauth2GoogleCallbackJSONRequestBody = OAuth2GoogleCallbackPayload

// OidcCallbackJSONRequestBody defines body for OidcCallback for application/json ContentType.
type OidcCallbackJSONRequestBody = OidcCallbackPayload

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// signin to viron with email/password
	// (POST /email/signin)
	SigninEmail(w http.ResponseWriter, r *http.Request)
	// redirect to google oauth
	// (GET /oauth2/google/authorization)
	Oauth2GoogleAuthorization(w http.ResponseWriter, r *http.Request, params Oauth2GoogleAuthorizationParams)
	// callback from google oauth
	// (POST /oauth2/google/callback)
	Oauth2GoogleCallback(w http.ResponseWriter, r *http.Request)
	// redirect to oidc idp authorization
	// (GET /oidc/authorization)
	OidcAuthorization(w http.ResponseWriter, r *http.Request, params OidcAuthorizationParams)
	// callback from oidc idp
	// (POST /oidc/callback)
	OidcCallback(w http.ResponseWriter, r *http.Request)
	// signout of viron
	// (POST /signout)
	Signout(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// signin to viron with email/password
// (POST /email/signin)
func (_ Unimplemented) SigninEmail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// redirect to google oauth
// (GET /oauth2/google/authorization)
func (_ Unimplemented) Oauth2GoogleAuthorization(w http.ResponseWriter, r *http.Request, params Oauth2GoogleAuthorizationParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// callback from google oauth
// (POST /oauth2/google/callback)
func (_ Unimplemented) Oauth2GoogleCallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// redirect to oidc idp authorization
// (GET /oidc/authorization)
func (_ Unimplemented) OidcAuthorization(w http.ResponseWriter, r *http.Request, params OidcAuthorizationParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// callback from oidc idp
// (POST /oidc/callback)
func (_ Unimplemented) OidcCallback(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// signout of viron
// (POST /signout)
func (_ Unimplemented) Signout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// SigninEmail operation middleware
func (siw *ServerInterfaceWrapper) SigninEmail(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SigninEmail(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Oauth2GoogleAuthorization operation middleware
func (siw *ServerInterfaceWrapper) Oauth2GoogleAuthorization(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params Oauth2GoogleAuthorizationParams

	// ------------- Required query parameter "redirectUri" -------------

	if paramValue := r.URL.Query().Get("redirectUri"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "redirectUri"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "redirectUri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redirectUri", Err: err})
		return
	}

	// ------------- Required query parameter "clientId" -------------

	if paramValue := r.URL.Query().Get("clientId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "clientId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "clientId", r.URL.Query(), &params.ClientId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clientId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Oauth2GoogleAuthorization(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Oauth2GoogleCallback operation middleware
func (siw *ServerInterfaceWrapper) Oauth2GoogleCallback(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Oauth2GoogleCallback(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// OidcAuthorization operation middleware
func (siw *ServerInterfaceWrapper) OidcAuthorization(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OidcAuthorizationParams

	// ------------- Required query parameter "redirectUri" -------------

	if paramValue := r.URL.Query().Get("redirectUri"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "redirectUri"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "redirectUri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redirectUri", Err: err})
		return
	}

	// ------------- Required query parameter "clientId" -------------

	if paramValue := r.URL.Query().Get("clientId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "clientId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "clientId", r.URL.Query(), &params.ClientId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "clientId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OidcAuthorization(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// OidcCallback operation middleware
func (siw *ServerInterfaceWrapper) OidcCallback(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OidcCallback(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Signout operation middleware
func (siw *ServerInterfaceWrapper) Signout(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Signout(w, r)
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
		r.Post(options.BaseURL+"/email/signin", wrapper.SigninEmail)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/oauth2/google/authorization", wrapper.Oauth2GoogleAuthorization)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/oauth2/google/callback", wrapper.Oauth2GoogleCallback)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/oidc/authorization", wrapper.OidcAuthorization)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/oidc/callback", wrapper.OidcCallback)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/signout", wrapper.Signout)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xX7W/bRBj/V8ytfBlOnLUIoUhoK9lAERMZqcoH2oCu9iW54fi883kjiyxhW0Kt1okX",
	"TaB9QGgDTaXdGqQKicHY/pijavkv0HN23p2l6ssH9qW1fc89L7/f/e550kEma7nMIY7wULGDXMxxiwjC",
	"1VvJpsQRZesjn/D2NViCr9RBRXQDPiEdObhFUBGZqSXSESc3fMqJhYqC+0RHntkkLQwb64y3sEBF5HOK",
	"dCTaLmz1BKdOAwWBjqrEopyYYpnT2SH5wPgkUYOeraq4suiL5vz7jDVsUsK2vYbNz6/hts2wpeDhzCVc",
	"UKKM+0UXO8ginsmpKyiDVBMHiTMZdWX8q4x+kdFDGe/JeL18eTIPHZnMItM8yXDz4P6fhw82ZfiDDH86",
	"3L67/3VXRnsyfibjjSxvw+jMSC/xsiPjb2Qcy6i7XC1nefQEFhkJlpaq7+13nx88+f7g3pYMd2X0VMZf",
	"KZfrMv4WSo8fwGv0IpPzAXErCQS9SPoYw320a303bO06MQUkV6GWeQK+KuXLpZPxBB60suWeMlNJYq8C",
	"Q0u04VDnSgtTeypBBFYnC7iS5hfvKGo2ZPxYRk+RPpB2sjEDEhd73i3GMyiHygGIbp+Zvrf+JtgvBOFg",
	"/2n+/MUVnLu9mPtkddV67eK5uddX/UJh/q1L598wcp/B84L5jvpHap239WAmlr2k++EmYYMt1KmzyfSx",
	"L5oadqnmucSkdWpiWPAgKBU2+Lh0k3LmGDZd08AY6egm4V6yu5Av5C8APswlDnYpKqKFfCE/n1TcVGQY",
	"Kj/DU7wprpgnMsgBq2FeZLg1hq2Mvvvn7xcy/FmGjz6GpGS4I6PfQWrxXvo3vC+jO0glxFUpINPhQ5Pe",
	"8cQT7zKrreTMHEEclRJ2XTuFwLjuMZXuoAPMcVJHRXTOGPQ6I73zjYxjGYwSBf1EffBc5njJQZ0vvDmJ",
	"xIdMK6Upqa7it1qYt4F/FUMTTFOUaLeoaGoJukNHTeCGB8dCkVUDDwaD53mjoa5rA14Yp7dxEq+DGkRM",
	"bxi76npP7p+De3/9++NDGf4h420Zf6kgf6zuu/UpwFfwoA8ujsTVR2aElWx0ByZGdk8P9JkbM4aPoDbG",
	"w0LhwiQCvYgAeFKBprDQlqtX82PU8CHbBGaNpWoZ5UNHX+TgyWROnTZyFqlj3xa50Xlp5CpHcx3lqg9A",
	"tRwkZ2uMVzPtW9M1llSRcLn/fFPd4ZM94epLWex1xzPS0csGp5MLapy2HmRanbPWsYgbw+DIzFHLPKIQ",
	"VesOd6nlaofbdw+3nh1DhdQy/8/qU0PRSAGzNAj4agAZHiv7NNRILXMapbM1CLUcQ4FDc+lZKS9j9D1r",
	"xfVoOi21ZTMDfZP5Yjon6TBxR0YbMn4io99gDokeTRfUUurxKFhUPsho48wXGqsnfTyjZwf9T53ez1S1",
	"FNSC/wIAAP//A+3S52oPAAA=",
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
