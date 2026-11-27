// Copyright (c) 2026 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package richresponse

import (
	"encoding/json"
	"reflect"
	"strings"

	"go.mau.fi/util/exslices"
)

type RichResponse struct {
	ResponseID      string               `json:"response_id"`
	Sections        []ViewModelContainer `json:"sections"`
	FooterSections  []ViewModelContainer `json:"footer_sections,omitempty"`
	EmbeddedScreens []json.RawMessage    `json:"embedded_screens,omitempty"`
}

// UnmarshalJSON accepts the stable sections even when optional entries use a newer schema.
func (r *RichResponse) UnmarshalJSON(data []byte) error {
	type responseFields RichResponse
	next := responseFields(*r)
	fields := struct {
		*responseFields
		Footer  json.RawMessage `json:"footer_sections"`
		Screens json.RawMessage `json:"embedded_screens"`
	}{responseFields: &next}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if fields.Footer != nil {
		var footer []ViewModelContainer
		if json.Unmarshal(fields.Footer, &footer) != nil {
			footer = nil
		}
		next.FooterSections = footer
	}
	if fields.Screens != nil {
		var screens []json.RawMessage
		if json.Unmarshal(fields.Screens, &screens) != nil {
			screens = nil
		}
		next.EmbeddedScreens = screens
	}
	*r = RichResponse(next)
	return nil
}

type ViewModelContainer struct {
	Model ViewModel `json:"view_model"`
}

type viewModelWrapper struct {
	Model json.RawMessage `json:"view_model"`
}

func (vmc *ViewModelContainer) UnmarshalJSON(data []byte) error {
	var wrapper viewModelWrapper
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}
	val, err := unmarshalWithTypeName[UnknownViewModel](wrapper.Model, viewModelTypes)
	if err != nil {
		return err
	}
	vmc.Model = val.(ViewModel)
	return nil
}

type ViewModel interface {
	isViewModel()
	GetPrimitives() []Primitive
	String() string
}

var viewModelTypes = map[string]reflect.Type{
	"GenAISingleLayoutViewModel":  reflect.TypeFor[GenAISingleLayoutViewModel](),
	"GenAIGridLayoutViewModel":    reflect.TypeFor[GenAIGridLayoutViewModel](),
	"GenAIHScrollLayoutViewModel": reflect.TypeFor[GenAIHScrollLayoutViewModel](),
	"GenAIVStackLayoutViewModel":  reflect.TypeFor[GenAIVStackLayoutViewModel](),
}

var (
	_ ViewModel = (*GenAISingleLayoutViewModel)(nil)
	_ ViewModel = (*MultiLayoutViewModel)(nil)
	_ ViewModel = (UnknownViewModel)(nil)
)

func (*GenAISingleLayoutViewModel) isViewModel() {}
func (*MultiLayoutViewModel) isViewModel()       {}
func (UnknownViewModel) isViewModel()            {}

type GenAISingleLayoutViewModel struct {
	TypeName  string             `json:"__typename"`
	Primitive PrimitiveContainer `json:"primitive"`
}

func (vm *GenAISingleLayoutViewModel) String() string {
	return vm.Primitive.Value.String()
}

func (vm *GenAISingleLayoutViewModel) GetPrimitives() []Primitive {
	return []Primitive{vm.Primitive.Value}
}

type MultiLayoutViewModel struct {
	TypeName   string               `json:"__typename"`
	Primitives []PrimitiveContainer `json:"primitives"`
}

func (vm *MultiLayoutViewModel) String() string {
	return strings.Join(exslices.CastFunc(vm.Primitives, func(from PrimitiveContainer) string {
		return from.Value.String()
	}), "\n")
}

func (vm *MultiLayoutViewModel) GetPrimitives() []Primitive {
	return exslices.CastFunc(vm.Primitives, func(from PrimitiveContainer) Primitive {
		return from.Value
	})
}

type GenAIGridLayoutViewModel = MultiLayoutViewModel
type GenAIHScrollLayoutViewModel = MultiLayoutViewModel
type GenAIVStackLayoutViewModel = MultiLayoutViewModel

type UnknownViewModel json.RawMessage

func (vm UnknownViewModel) GetPrimitives() []Primitive {
	return nil
}

func (vm UnknownViewModel) String() string {
	return ""
}

// MarshalJSON preserves unrecognized schema objects as raw JSON.
func (v UnknownViewModel) MarshalJSON() ([]byte, error) {
	return json.RawMessage(v).MarshalJSON()
}
