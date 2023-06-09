package binding

import (
	"bytes"
	"io"
	"net/http"

	"github.com/BurntSushi/toml"
)

type tomlBinding struct{}

func (tomlBinding) Name() string {
	return "toml"
}

func (tomlBinding) Bind(req *http.Request, obj any) error {
	return decodeToml(req.Body, obj)
}

func (tomlBinding) BindBody(body []byte, obj any) error {
	return decodeToml(bytes.NewReader(body), obj)
}

func decodeToml(r io.Reader, obj any) error {
	if _, err := toml.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
