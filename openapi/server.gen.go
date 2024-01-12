// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Detect Language
	// (POST /v1/papago/detectLangs)
	PostV1PapagoDetectLangs(ctx echo.Context, params PostV1PapagoDetectLangsParams) error
	// Translate Text
	// (POST /v1/papago/n2mt)
	PostV1PapagoN2mt(ctx echo.Context, params PostV1PapagoN2mtParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostV1PapagoDetectLangs converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1PapagoDetectLangs(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostV1PapagoDetectLangsParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Naver-Client-Id" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Naver-Client-Id")]; found {
		var XNaverClientId string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Naver-Client-Id, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Naver-Client-Id", runtime.ParamLocationHeader, valueList[0], &XNaverClientId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Naver-Client-Id: %s", err))
		}

		params.XNaverClientId = &XNaverClientId
	}
	// ------------- Optional header parameter "X-Naver-Client-Secret" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Naver-Client-Secret")]; found {
		var XNaverClientSecret string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Naver-Client-Secret, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Naver-Client-Secret", runtime.ParamLocationHeader, valueList[0], &XNaverClientSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Naver-Client-Secret: %s", err))
		}

		params.XNaverClientSecret = &XNaverClientSecret
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1PapagoDetectLangs(ctx, params)
	return err
}

// PostV1PapagoN2mt converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1PapagoN2mt(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostV1PapagoN2mtParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "X-Naver-Client-Id" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Naver-Client-Id")]; found {
		var XNaverClientId string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Naver-Client-Id, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Naver-Client-Id", runtime.ParamLocationHeader, valueList[0], &XNaverClientId)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Naver-Client-Id: %s", err))
		}

		params.XNaverClientId = &XNaverClientId
	}
	// ------------- Optional header parameter "X-Naver-Client-Secret" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Naver-Client-Secret")]; found {
		var XNaverClientSecret string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Naver-Client-Secret, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Naver-Client-Secret", runtime.ParamLocationHeader, valueList[0], &XNaverClientSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Naver-Client-Secret: %s", err))
		}

		params.XNaverClientSecret = &XNaverClientSecret
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostV1PapagoN2mt(ctx, params)
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

	router.POST(baseURL+"/v1/papago/detectLangs", wrapper.PostV1PapagoDetectLangs)
	router.POST(baseURL+"/v1/papago/n2mt", wrapper.PostV1PapagoN2mt)

}
