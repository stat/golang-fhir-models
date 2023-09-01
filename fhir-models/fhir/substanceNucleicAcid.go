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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/stat/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SubstanceNucleicAcid is documented here http://hl7.org/fhir/StructureDefinition/SubstanceNucleicAcid
type SubstanceNucleicAcid struct {
	Id                  *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceType        *CodeableConcept              `bson:"sequenceType,omitempty" json:"sequenceType,omitempty"`
	NumberOfSubunits    *int                          `bson:"numberOfSubunits,omitempty" json:"numberOfSubunits,omitempty"`
	AreaOfHybridisation *string                       `bson:"areaOfHybridisation,omitempty" json:"areaOfHybridisation,omitempty"`
	OligoNucleotideType *CodeableConcept              `bson:"oligoNucleotideType,omitempty" json:"oligoNucleotideType,omitempty"`
	Subunit             []SubstanceNucleicAcidSubunit `bson:"subunit,omitempty" json:"subunit,omitempty"`
}
type SubstanceNucleicAcidSubunit struct {
	Id                 *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Subunit            *int                                 `bson:"subunit,omitempty" json:"subunit,omitempty"`
	Sequence           *string                              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Length             *int                                 `bson:"length,omitempty" json:"length,omitempty"`
	SequenceAttachment *Attachment                          `bson:"sequenceAttachment,omitempty" json:"sequenceAttachment,omitempty"`
	FivePrime          *CodeableConcept                     `bson:"fivePrime,omitempty" json:"fivePrime,omitempty"`
	ThreePrime         *CodeableConcept                     `bson:"threePrime,omitempty" json:"threePrime,omitempty"`
	Linkage            []SubstanceNucleicAcidSubunitLinkage `bson:"linkage,omitempty" json:"linkage,omitempty"`
	Sugar              []SubstanceNucleicAcidSubunitSugar   `bson:"sugar,omitempty" json:"sugar,omitempty"`
}
type SubstanceNucleicAcidSubunitLinkage struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Connectivity      *string     `bson:"connectivity,omitempty" json:"connectivity,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	ResidueSite       *string     `bson:"residueSite,omitempty" json:"residueSite,omitempty"`
}
type SubstanceNucleicAcidSubunitSugar struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	ResidueSite       *string     `bson:"residueSite,omitempty" json:"residueSite,omitempty"`
}
type OtherSubstanceNucleicAcid SubstanceNucleicAcid

// MarshalJSON marshals the given SubstanceNucleicAcid as JSON into a byte slice
 func (r *SubstanceNucleicAcid)MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceNucleicAcid
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceNucleicAcid: OtherSubstanceNucleicAcid(*r),
		ResourceType:              "SubstanceNucleicAcid",
	})
}

// UnmarshalSubstanceNucleicAcid unmarshals a SubstanceNucleicAcid.
func UnmarshalSubstanceNucleicAcid(b []byte) (SubstanceNucleicAcid, error) {
	var substanceNucleicAcid SubstanceNucleicAcid
	if err := json.Unmarshal(b, &substanceNucleicAcid); err != nil {
		return substanceNucleicAcid, err
	}
	return substanceNucleicAcid, nil
}
