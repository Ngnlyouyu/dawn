package binding

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"gopkg.in/yaml.v3"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	return decodeYAML(req.Body, obj)
}

func (yamlBinding) BindBody(body []byte, obj any) error {
	return decodeYAML(bytes.NewReader(body), obj)
}

func decodeYAML(r io.Reader, obj any) error {
	if err := yaml.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
