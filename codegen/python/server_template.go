package python

import (
	"fmt"
)

type serverTemplate struct {
	apiFile     string
	apiName     string
	handlerFile string
	handlerName string
}

func templates(kind string) serverTemplate {
	switch kind {
	case serverKindFlask, serverKindGeventFlask:
		return templatesFlask()
	case serverKindSanic:
		return templatesSanic()
	default:
		panic(fmt.Errorf("invalid server kind:%v", kind))
	}
}

func templatesFlask() serverTemplate {
	return serverTemplate{
		apiFile:     "./templates/python/server_resource_api_flask.tmpl",
		apiName:     "server_resource_api_flask",
		handlerFile: "./templates/python/server_resource_handler_flask.tmpl",
		handlerName: "server_resource_handler_flask",
	}
}

func templatesSanic() serverTemplate {
	return serverTemplate{
		apiFile:     "./templates/python/server_resource_api_sanic.tmpl",
		apiName:     "server_resource_api_sanic",
		handlerFile: "./templates/python/server_resource_handler_sanic.tmpl",
		handlerName: "server_resource_handler_sanic",
	}
}
