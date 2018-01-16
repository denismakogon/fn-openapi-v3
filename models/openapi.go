package models

import (
	"encoding/json"
	"github.com/go-openapi/spec"
	"io"
)

type OpenAPISpec struct {
	Version     string                   `json:"openapi" yaml:"openapi"`
	Servers     []Server                 `json:"servers" yaml:"servers"`
	APISpecInfo APISpecInfo              `json:"info" yaml:"info"`
	Paths       map[string]PathItemProps `json:"paths" yaml:"paths"`
}

type APISpecInfo struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	APIVersion  string `json:"version" yaml:"version"`
}

type Server struct {
	URL string `json:"url"`
}

type OperationProperties struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Consumes    []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Schemes     []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"` // the scheme, when present must be from [http, https, ws, wss]
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	ID          string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Security    []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Parameters  []spec.Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[int]Response      `json:"responses,omitempty" yaml:"responses,omitempty"`
}

type PathItemProps struct {
	Get        *OperationProperties `json:"get,omitempty" yaml:"get,omitempty"`
	Put        *OperationProperties `json:"put,omitempty" yaml:"put,omitempty"`
	Post       *OperationProperties `json:"post,omitempty" yaml:"post,omitempty"`
	Delete     *OperationProperties `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options    *OperationProperties `json:"options,omitempty" yaml:"options,omitempty"`
	Head       *OperationProperties `json:"head,omitempty" yaml:"head,omitempty"`
	Patch      *OperationProperties `json:"patch,omitempty" yaml:"patch,omitempty"`
	Parameters []spec.Parameter     `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

func (oai *OpenAPISpec) FromFnSpec(fnAPIURL string, fn *Fn) error {
	apiSpec := APISpecInfo{
		Title:       "Fn serverless application API spec",
		Description: fn.Description,
		APIVersion:  fn.Version,
	}
	fnServer := Server{URL: fnAPIURL}

	oai.Version = "3.0.0"
	oai.Servers = []Server{fnServer}
	oai.APISpecInfo = apiSpec
	oai.Paths = map[string]PathItemProps{}

	// fnName
	for fnName, function := range fn.Functions {
		for _, typedEvents := range function.Events {
			// fnFormat
			for _, event := range typedEvents {
				pathProps := PathItemProps{}
				var params []spec.Parameter
				for _, p := range event.Documentation.Parameters {
					params = append(params, spec.Parameter{ParamProps: p})
				}
				//event.Documentation.MethodResponses
				op := OperationProperties{
					ID:          fnName,
					Description: event.Documentation.Description,
					Summary:     event.Documentation.Summary,
					Parameters:  params,
					Responses:   event.Documentation.Responses,
				}
				switch event.Method {
				case "get":
					pathProps.Get = &op
				case "post":
					pathProps.Post = &op
				case "delete":
					pathProps.Delete = &op
				case "patch":
					pathProps.Patch = &op
				case "put":
					pathProps.Put = &op
				case "options":
					pathProps.Options = &op
				case "head":
					pathProps.Head = &op
				}
				//TODO(denimakogon): join /r appName and route
				oai.Paths[event.Fn.Path] = pathProps
			}
		}
	}
	return nil
}

func (oia *OpenAPISpec) Marshal(w io.Writer) error {
	return json.NewEncoder(w).Encode(oia)
}
