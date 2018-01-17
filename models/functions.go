package models

import (
	"github.com/go-openapi/spec"
	"github.com/mitchellh/mapstructure"
)

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
	Application Application            `json:"application" yaml:"application"`
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

	return mapstructure.Decode(fnInterface, fn)
}

type Event struct {
	Method        string                `json:"method,omitempty" yaml:"method,omitempty"`
	Fn            FnFunction            `json:"fn" yaml:"fn"`
	Documentation FunctionDocumentation `json:"documentation,omitempty" yaml:"documentation,omitempty"`
}

type SimpleSchema struct {
	spec.SimpleSchema
	Maximum          *float64      `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	Minimum          *float64      `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	MaxLength        *int64        `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        *int64        `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          string        `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	MaxItems         *int64        `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
	MinItems         *int64        `json:"minItems,omitempty" yaml:"minItems,omitempty"`
	UniqueItems      bool          `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
	MultipleOf       *float64      `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
	Enum             []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
}

type Parameter struct {
	Description      string       `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string       `json:"name,omitempty" yaml:"name,omitempty"`
	In               string       `json:"in,omitempty" yaml:"in,omitempty"`
	Required         bool         `json:"required,omitempty" yaml:"required,omitempty"`
	Schema           SimpleSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
	AllowEmptyValue  bool         `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	Type             string       `json:"type,omitempty" yaml:"type,omitempty"`
	Format           string       `json:"format,omitempty" yaml:"format,omitempty"`
	Items            *spec.Items  `json:"items,omitempty" yaml:"items,omitempty"`
	CollectionFormat string       `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          interface{}  `json:"default,omitempty" yaml:"default,omitempty"`
}

type FunctionDocumentation struct {
	Summary     string           `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	RequestBody PayloadSchema    `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Parameters  []Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[int]Response `json:"responses,omitempty" yaml:"responses,omitempty"`
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
