package models

import "github.com/go-openapi/spec"

type Function struct {
	Handler string             `json:"handler,omitempty" yaml:"handler,omitempty"`
	Events  []map[string]Event `json:"events,omitempty" yaml:"events,omitempty"`
}

func isValidKey(key string, listOfValidKeys []string) bool {
	for _, validKey := range listOfValidKeys {
		if key == validKey {
			return true
		}
	}
	return false
}

func (f *Function) ValidateEventTypes() bool {
	validEvents := []string{"http", "json", "default"}
	for _, event := range f.Events {
		for eventType := range event {
			if !isValidKey(eventType, validEvents) {
				return false
			}
		}
	}
	return true
}

type FnFunction struct {
	Name        string                 `json:"name" yaml:"name,omitempty"`
	Version     string                 `json:"version" yaml:"version,omitempty"`
	Path        string                 `json:"path,omitempty" yaml:"path,omitempty"`
	Image       string                 `json:"image,omitempty" yaml:"image,omitempty"`
	Timeout     int32                  `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	IdleTimeout int32                  `json:"idle_timeout,omitempty" yaml:"idle_timeout,omitempty"`
	Memory      int32                  `json:"memory,omitempty" yaml:"memory,omitempty"`
	Type        string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Config      map[string]interface{} `json:"config" yaml:"config"`
}

func (fn *FnFunction) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var fnInterface map[string]interface{}

	err := unmarshal(&fnInterface)
	if err != nil {
		return unmarshalReferenceFile(unmarshal, fn)
	}

	return nil
}

type Event struct {
	Method        string                `json:"method,omitempty" yaml:"method,omitempty"`
	Fn            FnFunction            `json:"fn" yaml:"fn"`
	Documentation FunctionDocumentation `json:"documentation,omitempty" yaml:"documentation,omitempty"`
}

type FunctionDocumentation struct {
	Summary     string            `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	RequestBody PayloadSchema     `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Parameters  []spec.ParamProps `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[int]Response  `json:"responses,omitempty" yaml:"responses,omitempty"`
}

type Response struct {
	Description string                   `json:"description,omitempty"`
	Content     map[string]PayloadSchema `json:"content"`
}

type PayloadSchema struct {
	Schema Schema `json:"schema"`
}

type RequestResponseBody struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
