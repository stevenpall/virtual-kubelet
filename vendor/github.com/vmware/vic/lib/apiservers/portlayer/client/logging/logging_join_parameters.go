package logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewLoggingJoinParams creates a new LoggingJoinParams object
// with the default values initialized.
func NewLoggingJoinParams() *LoggingJoinParams {
	var ()
	return &LoggingJoinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoggingJoinParamsWithTimeout creates a new LoggingJoinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoggingJoinParamsWithTimeout(timeout time.Duration) *LoggingJoinParams {
	var ()
	return &LoggingJoinParams{

		timeout: timeout,
	}
}

// NewLoggingJoinParamsWithContext creates a new LoggingJoinParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoggingJoinParamsWithContext(ctx context.Context) *LoggingJoinParams {
	var ()
	return &LoggingJoinParams{

		Context: ctx,
	}
}

// NewLoggingJoinParamsWithHTTPClient creates a new LoggingJoinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoggingJoinParamsWithHTTPClient(client *http.Client) *LoggingJoinParams {
	var ()
	return &LoggingJoinParams{
		HTTPClient: client,
	}
}

/*LoggingJoinParams contains all the parameters to send to the API endpoint
for the logging join operation typically these are written to a http.Request
*/
type LoggingJoinParams struct {

	/*Config*/
	Config *models.LoggingJoinConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the logging join params
func (o *LoggingJoinParams) WithTimeout(timeout time.Duration) *LoggingJoinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logging join params
func (o *LoggingJoinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logging join params
func (o *LoggingJoinParams) WithContext(ctx context.Context) *LoggingJoinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logging join params
func (o *LoggingJoinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logging join params
func (o *LoggingJoinParams) WithHTTPClient(client *http.Client) *LoggingJoinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logging join params
func (o *LoggingJoinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the logging join params
func (o *LoggingJoinParams) WithConfig(config *models.LoggingJoinConfig) *LoggingJoinParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the logging join params
func (o *LoggingJoinParams) SetConfig(config *models.LoggingJoinConfig) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *LoggingJoinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Config == nil {
		o.Config = new(models.LoggingJoinConfig)
	}

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
