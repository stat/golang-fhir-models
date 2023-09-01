// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// BiologicallyDerivedProductCategory is documented here http://hl7.org/fhir/ValueSet/product-category
type BiologicallyDerivedProductCategory int

const (
	BiologicallyDerivedProductCategoryOrgan BiologicallyDerivedProductCategory = iota
	BiologicallyDerivedProductCategoryTissue
	BiologicallyDerivedProductCategoryFluid
	BiologicallyDerivedProductCategoryCells
	BiologicallyDerivedProductCategoryBiologicalAgent
)

 func (code *BiologicallyDerivedProductCategory)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "organ":
		*code = BiologicallyDerivedProductCategoryOrgan
	case "tissue":
		*code = BiologicallyDerivedProductCategoryTissue
	case "fluid":
		*code = BiologicallyDerivedProductCategoryFluid
	case "cells":
		*code = BiologicallyDerivedProductCategoryCells
	case "biologicalAgent":
		*code = BiologicallyDerivedProductCategoryBiologicalAgent
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductCategory code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductCategory) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductCategory) Code() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "organ"
	case BiologicallyDerivedProductCategoryTissue:
		return "tissue"
	case BiologicallyDerivedProductCategoryFluid:
		return "fluid"
	case BiologicallyDerivedProductCategoryCells:
		return "cells"
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "biologicalAgent"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductCategory) Display() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "Organ"
	case BiologicallyDerivedProductCategoryTissue:
		return "Tissue"
	case BiologicallyDerivedProductCategoryFluid:
		return "Fluid"
	case BiologicallyDerivedProductCategoryCells:
		return "Cells"
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "BiologicalAgent"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductCategory) Definition() string {
	switch code {
	case BiologicallyDerivedProductCategoryOrgan:
		return "A collection of tissues joined in a structural unit to serve a common function."
	case BiologicallyDerivedProductCategoryTissue:
		return "An ensemble of similar cells and their extracellular matrix from the same origin that together carry out a specific function."
	case BiologicallyDerivedProductCategoryFluid:
		return "Body fluid."
	case BiologicallyDerivedProductCategoryCells:
		return "Collection of cells."
	case BiologicallyDerivedProductCategoryBiologicalAgent:
		return "Biological agent of unspecified type."
	}
	return "<unknown>"
}
