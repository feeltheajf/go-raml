package commands

import (
	"github.com/feeltheajf/go-raml/raml"

	log "github.com/Sirupsen/logrus"
)

//ParseCommand is executed to generate client from a RAML specification
type ParseCommand struct {
	Language    string
	Dir         string //target dir
	RamlFile    string //raml file
	PackageName string //package name in the generated go source files
	ImportPath  string
	Kind        string

	// Root URL of the libraries.
	// Usefull if we want to use remote libraries.
	// Example:
	//   root url     = http://localhost.com/lib
	//   library file = http://localhost.com/lib/libraries/security.raml
	//	 the library file is going to treated the same as local : libraries/security.raml
	LibRootURLs string

	// If true, python client will unmarshall the response
	// Other languages already unmarshall the response
	PythonUnmarshallResponse bool
}

//Execute generates a client from a RAML specification
func (command *ParseCommand) Execute() error {
	log.Debug("Generating a rest client for ", command.Language)
	apiDef := new(raml.APIDefinition)
	err := raml.ParseFile(command.RamlFile, apiDef)
	if err != nil {
		return err
	}

	// TODO: do something here

	return nil
}
