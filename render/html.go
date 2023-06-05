package render

// Delims represents a set of Left and Right delimeters for HTML template rendering.
type Delims struct {
	// Left delimeter
	Left string
	// Right delimeter
	Right string
}

type HTMLRender interface {
	Instance(string, any) Render
}
