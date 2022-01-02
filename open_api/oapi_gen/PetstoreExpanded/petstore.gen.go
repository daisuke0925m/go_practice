// Package PetstoreExpanded provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package PetstoreExpanded

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Error defines model for Error.
type Error struct {
	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// Unique id of the pet
	Id int64 `json:"id"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody NewPet

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// FindPets request
	FindPets(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddPet request with any body
	AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeletePet request
	DeletePet(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FindPetByID request
	FindPetByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) FindPets(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindPetsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPetWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddPet(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddPetRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePet(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeletePetRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FindPetByID(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFindPetByIDRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewFindPetsRequest generates requests for FindPets
func NewFindPetsRequest(server string, params *FindPetsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if params.Tags != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "tags", runtime.ParamLocationQuery, *params.Tags); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Limit != nil {

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddPetRequest calls the generic AddPet builder with application/json body
func NewAddPetRequest(server string, body AddPetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddPetRequestWithBody(server, "application/json", bodyReader)
}

// NewAddPetRequestWithBody generates requests for AddPet with any type of body
func NewAddPetRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeletePetRequest generates requests for DeletePet
func NewDeletePetRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewFindPetByIDRequest generates requests for FindPetByID
func NewFindPetByIDRequest(server string, id int64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// FindPets request
	FindPetsWithResponse(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*FindPetsResponse, error)

	// AddPet request with any body
	AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddPetResponse, error)

	AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddPetResponse, error)

	// DeletePet request
	DeletePetWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeletePetResponse, error)

	// FindPetByID request
	FindPetByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindPetByIDResponse, error)
}

type FindPetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r FindPetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindPetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r AddPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeletePetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r DeletePetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeletePetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FindPetByIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Pet
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r FindPetByIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FindPetByIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// FindPetsWithResponse request returning *FindPetsResponse
func (c *ClientWithResponses) FindPetsWithResponse(ctx context.Context, params *FindPetsParams, reqEditors ...RequestEditorFn) (*FindPetsResponse, error) {
	rsp, err := c.FindPets(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFindPetsResponse(rsp)
}

// AddPetWithBodyWithResponse request with arbitrary body returning *AddPetResponse
func (c *ClientWithResponses) AddPetWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddPetResponse, error) {
	rsp, err := c.AddPetWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

func (c *ClientWithResponses) AddPetWithResponse(ctx context.Context, body AddPetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddPetResponse, error) {
	rsp, err := c.AddPet(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddPetResponse(rsp)
}

// DeletePetWithResponse request returning *DeletePetResponse
func (c *ClientWithResponses) DeletePetWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*DeletePetResponse, error) {
	rsp, err := c.DeletePet(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeletePetResponse(rsp)
}

// FindPetByIDWithResponse request returning *FindPetByIDResponse
func (c *ClientWithResponses) FindPetByIDWithResponse(ctx context.Context, id int64, reqEditors ...RequestEditorFn) (*FindPetByIDResponse, error) {
	rsp, err := c.FindPetByID(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFindPetByIDResponse(rsp)
}

// ParseFindPetsResponse parses an HTTP response from a FindPetsWithResponse call
func ParseFindPetsResponse(rsp *http.Response) (*FindPetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FindPetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseAddPetResponse parses an HTTP response from a AddPetWithResponse call
func ParseAddPetResponse(rsp *http.Response) (*AddPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AddPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDeletePetResponse parses an HTTP response from a DeletePetWithResponse call
func ParseDeletePetResponse(rsp *http.Response) (*DeletePetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeletePetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseFindPetByIDResponse parses an HTTP response from a FindPetByIDWithResponse call
func ParseFindPetByIDResponse(rsp *http.Response) (*FindPetByIDResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FindPetByIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Pet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx echo.Context) error
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id int64) error
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"3d8CAAD//xmGt4NVEgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
