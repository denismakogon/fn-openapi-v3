package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type unmarshalFunc func(data []byte, v interface{}) error

func processFileWithAttribute(includeFilePath, attributeReference string, unmarshalType interface{}, unmarshal unmarshalFunc) error {
	b, err := ioutil.ReadFile(includeFilePath)
	if attributeReference != "" {
		var withReference map[string]interface{}
		err = unmarshal(b, &withReference)
		if err != nil {
			return err
		}
		err = mapstructure.Decode(withReference[attributeReference], unmarshalType)
		if err != nil {
			return err
		}
		return nil
	}
	err = unmarshal(b, unmarshalType)
	if err != nil {
		return err
	}
	return nil
}

func unmarshalReferenceFile(unmarshal func(interface{}) error, unmarshalType interface{}) error {
	var v string
	err := unmarshal(&v)
	if err != nil {
		return err
	}

	attributeReference := ""
	includeFile := v[len("${file(") : len(v)-len("}")]
	if strings.Contains(includeFile, ":") {
		parts := strings.Split(includeFile, ":")
		includeFile, attributeReference = parts[0], parts[1]
	}
	includeFile = includeFile[:len(includeFile)-1]
	_, err = os.Stat(includeFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("unable to include reference file %v, err: %v", includeFile, err.Error())
	}
	switch filepath.Ext(includeFile) {
	case ".json":
		err = processFileWithAttribute(includeFile, attributeReference, unmarshalType, json.Unmarshal)
		if err != nil {
			return err
		}
	case ".yml", ".yaml":
		err = processFileWithAttribute(includeFile, attributeReference, unmarshalType, yaml.Unmarshal)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported include file format (not *.json or *.y(a)ml)")
	}

	return nil
}
