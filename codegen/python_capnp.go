package codegen

import (
	"github.com/feeltheajf/go-raml/codegen/python"
	"github.com/feeltheajf/go-raml/raml"
)

func GeneratePythonCapnp(apiDef *raml.APIDefinition, dir string) error {
	return python.GeneratePythonCapnpClasses(apiDef, dir)
}
