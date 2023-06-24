package binding

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj any) error {
	if req == nil || req.Body == nil {
		return errors.New("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj any) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj any) error {
	if err := json.NewDecoder(r).Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
