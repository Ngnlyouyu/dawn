package binding

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(m map[string][]string, obj any) error {
	if err := mapURI(obj, m); err != nil {
		return err
	}
	return validate(obj)
}
