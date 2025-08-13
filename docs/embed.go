package docs

import (
	_ "embed"

	"github.com/swaggo/swag"
)

//go:embed swagger.json
var swaggerJSON []byte

type reader struct{}

func (r *reader) ReadDoc() string {
	return string(swaggerJSON)
}

func init() {
	swag.Register(swag.Name, &reader{})
}
