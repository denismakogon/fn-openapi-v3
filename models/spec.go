package models

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io"
)

type Documentation struct {
	Version     string  `json:"version,omitempty" yaml:"version,omitempty"`
	Title       string  `json:"title,omitempty" yaml:"title,omitempty"`
	Description string  `json:"description,omitempty" yaml:"description,omitempty"`
	Models      []Model `json:"models,omitempty" yaml:"models,omitempty"`
}

type Model struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Schema      Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type Fn struct {
	Functions   map[string]Function `json:"functions,omitempty" yaml:"functions,omitempty"`
	Version     string              `json:"version" yaml:"version"`
	Description string              `json:"description" yaml:"description"`
	Application Application         `json:"application" yaml:"application"`
}

type Application struct {
	Name   string                 `json:"name"`
	Config map[string]interface{} `json:"config"`
}

func (fn *Fn) Unmarshal(content []byte, w io.Writer) error {
	err := yaml.Unmarshal(content, &fn)
	if err != nil {
		return err
	}
	return nil
}

func (fn *Fn) Marshal(w io.Writer) error {
	return json.NewEncoder(w).Encode(fn)
}

func (fn *Fn) Validate() bool {
	for _, f := range fn.Functions {
		if !f.ValidateEventTypes() {
			return false
		}
	}
	return true
}
