// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package jsonquery

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	ProcessorConfigField     = "field"
	ProcessorConfigThreshold = "threshold"
)

func (ProcessorConfig) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		ProcessorConfigField: {
			Default:     "",
			Description: "Field is the target field that will be set.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
				config.ValidationExclusion{List: []string{".Position"}},
			},
		},
		ProcessorConfigThreshold: {
			Default:     "",
			Description: "Threshold is the threshold for filtering the record.",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{
				config.ValidationRequired{},
				config.ValidationGreaterThan{V: 0},
			},
		},
	}
}
