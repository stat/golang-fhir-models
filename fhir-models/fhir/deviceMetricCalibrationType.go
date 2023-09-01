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

// DeviceMetricCalibrationType is documented here http://hl7.org/fhir/ValueSet/metric-calibration-type
type DeviceMetricCalibrationType int

const (
	DeviceMetricCalibrationTypeUnspecified DeviceMetricCalibrationType = iota
	DeviceMetricCalibrationTypeOffset
	DeviceMetricCalibrationTypeGain
	DeviceMetricCalibrationTypeTwoPoint
)

 func (code *DeviceMetricCalibrationType)MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricCalibrationType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "unspecified":
		*code = DeviceMetricCalibrationTypeUnspecified
	case "offset":
		*code = DeviceMetricCalibrationTypeOffset
	case "gain":
		*code = DeviceMetricCalibrationTypeGain
	case "two-point":
		*code = DeviceMetricCalibrationTypeTwoPoint
	default:
		return fmt.Errorf("unknown DeviceMetricCalibrationType code `%s`", s)
	}
	return nil
}
func (code DeviceMetricCalibrationType) String() string {
	return code.Code()
}
func (code DeviceMetricCalibrationType) Code() string {
	switch code {
	case DeviceMetricCalibrationTypeUnspecified:
		return "unspecified"
	case DeviceMetricCalibrationTypeOffset:
		return "offset"
	case DeviceMetricCalibrationTypeGain:
		return "gain"
	case DeviceMetricCalibrationTypeTwoPoint:
		return "two-point"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationType) Display() string {
	switch code {
	case DeviceMetricCalibrationTypeUnspecified:
		return "Unspecified"
	case DeviceMetricCalibrationTypeOffset:
		return "Offset"
	case DeviceMetricCalibrationTypeGain:
		return "Gain"
	case DeviceMetricCalibrationTypeTwoPoint:
		return "Two Point"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationType) Definition() string {
	switch code {
	case DeviceMetricCalibrationTypeUnspecified:
		return "Metric calibration method has not been identified."
	case DeviceMetricCalibrationTypeOffset:
		return "Offset metric calibration method."
	case DeviceMetricCalibrationTypeGain:
		return "Gain metric calibration method."
	case DeviceMetricCalibrationTypeTwoPoint:
		return "Two-point metric calibration method."
	}
	return "<unknown>"
}
