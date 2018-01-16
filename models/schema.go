package models

import (
	"github.com/go-openapi/spec"
	"github.com/mitchellh/mapstructure"
)

type Schema struct {
	spec.Schema
}

func (schema *Schema) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s map[string]interface{}

	err := unmarshal(&s)
	if err != nil {
		return unmarshalReferenceFile(unmarshal, schema)
	}
	err = mapstructure.Decode(s, schema)

	return nil
}
