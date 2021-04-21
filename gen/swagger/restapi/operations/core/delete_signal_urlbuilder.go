// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteSignalURL generates an URL for the delete signal operation
type DeleteSignalURL struct {
	SignalID int64

	Datetime *strfmt.DateTime

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteSignalURL) WithBasePath(bp string) *DeleteSignalURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteSignalURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteSignalURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/core/signals/{signal_id}"

	signalID := swag.FormatInt64(o.SignalID)
	if signalID != "" {
		_path = strings.Replace(_path, "{signal_id}", signalID, -1)
	} else {
		return nil, errors.New("signalId is required on DeleteSignalURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var datetimeQ string
	if o.Datetime != nil {
		datetimeQ = o.Datetime.String()
	}
	if datetimeQ != "" {
		qs.Set("datetime", datetimeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteSignalURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteSignalURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteSignalURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteSignalURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteSignalURL")
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
func (o *DeleteSignalURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
