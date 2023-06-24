package binding

import "net/http"

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "aplication/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEYAML              = "application/x-yaml"
	MIMETOML              = "application/toml"
)

// Binding describes the interface which needs to be implemented for binding the data present
// in the request such as JSON request body, query parameters or the form POST.
type Binding interface {
	Name() string
	Bind(*http.Request, any) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with Bind, but it reads
// the body from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, any) error
}

// BindingUri adds BindUri method to Binding. BindUri is similar with Bind,
// but it reads the Params.
type BindingUri interface {
	Name() string
	BindUri(map[string][]string, any) error
}

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Form          = formBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	ProtoBuf      = protobufBinding{}
	YAML          = yamlBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
	TOML          = tomlBinding{}
)

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method, contentType string) Binding {
	if method == http.MethodGet {
		return Form
	}

	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML, MIMEXML2:
		return XML
	case MIMEPROTOBUF:
		return ProtoBuf
	case MIMEYAML:
		return YAML
	case MIMEMultipartPOSTForm:
		return FormMultipart
	case MIMETOML:
		return TOML
	default: // case MIMEPOSTForm:
		return Form
	}
}
