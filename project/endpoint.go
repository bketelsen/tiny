package project

import "github.com/iancoleman/strcase"

type Endpoint struct {
	Name      string
	MethodMap map[string]*Method
}

func (e *Endpoint) GetMethod(name string) (*Method, bool) {
	method, ok := e.MethodMap[name]
	return method, ok
}

func (e *Endpoint) GetAllMethods() []*Method {
	methods := make([]*Method, 0, len(e.MethodMap))
	for _, method := range e.MethodMap {
		methods = append(methods, method)
	}
	return methods
}

func (e *Endpoint) ClientStructName() string {
	return strcase.ToLowerCamel(e.Name)
}
