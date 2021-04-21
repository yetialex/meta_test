// Code generated by go-swagger; DO NOT EDIT.

package iba

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetIBAGateByNameURL generates an URL for the get i b a gate by name operation
type GetIBAGateByNameURL struct {
	GateName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetIBAGateByNameURL) WithBasePath(bp string) *GetIBAGateByNameURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetIBAGateByNameURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetIBAGateByNameURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/ibas/gates/${gate_name}"

	gateName := o.GateName
	if gateName != "" {
		_path = strings.Replace(_path, "{gate_name}", gateName, -1)
	} else {
		return nil, errors.New("gateName is required on GetIBAGateByNameURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetIBAGateByNameURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetIBAGateByNameURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetIBAGateByNameURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetIBAGateByNameURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetIBAGateByNameURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetIBAGateByNameURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
