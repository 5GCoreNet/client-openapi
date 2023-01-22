/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// CCOParametersAttrAllOfAttributes struct for CCOParametersAttrAllOfAttributes
type CCOParametersAttrAllOfAttributes struct {
	CoverageShapeList *int32 `json:"coverageShapeList,omitempty"`
	DownlinkTransmitPowerRange *ParameterRange `json:"downlinkTransmitPowerRange,omitempty"`
	AntennaTiltRange *ParameterRange `json:"antennaTiltRange,omitempty"`
	AntennaAzimuthRange *ParameterRange `json:"antennaAzimuthRange,omitempty"`
	DigitalTiltRange *ParameterRange `json:"digitalTiltRange,omitempty"`
	DigitalAzimuthRange *ParameterRange `json:"digitalAzimuthRange,omitempty"`
}

// NewCCOParametersAttrAllOfAttributes instantiates a new CCOParametersAttrAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOParametersAttrAllOfAttributes() *CCOParametersAttrAllOfAttributes {
	this := CCOParametersAttrAllOfAttributes{}
	return &this
}

// NewCCOParametersAttrAllOfAttributesWithDefaults instantiates a new CCOParametersAttrAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOParametersAttrAllOfAttributesWithDefaults() *CCOParametersAttrAllOfAttributes {
	this := CCOParametersAttrAllOfAttributes{}
	return &this
}

// GetCoverageShapeList returns the CoverageShapeList field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetCoverageShapeList() int32 {
	if o == nil || isNil(o.CoverageShapeList) {
		var ret int32
		return ret
	}
	return *o.CoverageShapeList
}

// GetCoverageShapeListOk returns a tuple with the CoverageShapeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetCoverageShapeListOk() (*int32, bool) {
	if o == nil || isNil(o.CoverageShapeList) {
    return nil, false
	}
	return o.CoverageShapeList, true
}

// HasCoverageShapeList returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasCoverageShapeList() bool {
	if o != nil && !isNil(o.CoverageShapeList) {
		return true
	}

	return false
}

// SetCoverageShapeList gets a reference to the given int32 and assigns it to the CoverageShapeList field.
func (o *CCOParametersAttrAllOfAttributes) SetCoverageShapeList(v int32) {
	o.CoverageShapeList = &v
}

// GetDownlinkTransmitPowerRange returns the DownlinkTransmitPowerRange field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetDownlinkTransmitPowerRange() ParameterRange {
	if o == nil || isNil(o.DownlinkTransmitPowerRange) {
		var ret ParameterRange
		return ret
	}
	return *o.DownlinkTransmitPowerRange
}

// GetDownlinkTransmitPowerRangeOk returns a tuple with the DownlinkTransmitPowerRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetDownlinkTransmitPowerRangeOk() (*ParameterRange, bool) {
	if o == nil || isNil(o.DownlinkTransmitPowerRange) {
    return nil, false
	}
	return o.DownlinkTransmitPowerRange, true
}

// HasDownlinkTransmitPowerRange returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasDownlinkTransmitPowerRange() bool {
	if o != nil && !isNil(o.DownlinkTransmitPowerRange) {
		return true
	}

	return false
}

// SetDownlinkTransmitPowerRange gets a reference to the given ParameterRange and assigns it to the DownlinkTransmitPowerRange field.
func (o *CCOParametersAttrAllOfAttributes) SetDownlinkTransmitPowerRange(v ParameterRange) {
	o.DownlinkTransmitPowerRange = &v
}

// GetAntennaTiltRange returns the AntennaTiltRange field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetAntennaTiltRange() ParameterRange {
	if o == nil || isNil(o.AntennaTiltRange) {
		var ret ParameterRange
		return ret
	}
	return *o.AntennaTiltRange
}

// GetAntennaTiltRangeOk returns a tuple with the AntennaTiltRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetAntennaTiltRangeOk() (*ParameterRange, bool) {
	if o == nil || isNil(o.AntennaTiltRange) {
    return nil, false
	}
	return o.AntennaTiltRange, true
}

// HasAntennaTiltRange returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasAntennaTiltRange() bool {
	if o != nil && !isNil(o.AntennaTiltRange) {
		return true
	}

	return false
}

// SetAntennaTiltRange gets a reference to the given ParameterRange and assigns it to the AntennaTiltRange field.
func (o *CCOParametersAttrAllOfAttributes) SetAntennaTiltRange(v ParameterRange) {
	o.AntennaTiltRange = &v
}

// GetAntennaAzimuthRange returns the AntennaAzimuthRange field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetAntennaAzimuthRange() ParameterRange {
	if o == nil || isNil(o.AntennaAzimuthRange) {
		var ret ParameterRange
		return ret
	}
	return *o.AntennaAzimuthRange
}

// GetAntennaAzimuthRangeOk returns a tuple with the AntennaAzimuthRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetAntennaAzimuthRangeOk() (*ParameterRange, bool) {
	if o == nil || isNil(o.AntennaAzimuthRange) {
    return nil, false
	}
	return o.AntennaAzimuthRange, true
}

// HasAntennaAzimuthRange returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasAntennaAzimuthRange() bool {
	if o != nil && !isNil(o.AntennaAzimuthRange) {
		return true
	}

	return false
}

// SetAntennaAzimuthRange gets a reference to the given ParameterRange and assigns it to the AntennaAzimuthRange field.
func (o *CCOParametersAttrAllOfAttributes) SetAntennaAzimuthRange(v ParameterRange) {
	o.AntennaAzimuthRange = &v
}

// GetDigitalTiltRange returns the DigitalTiltRange field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetDigitalTiltRange() ParameterRange {
	if o == nil || isNil(o.DigitalTiltRange) {
		var ret ParameterRange
		return ret
	}
	return *o.DigitalTiltRange
}

// GetDigitalTiltRangeOk returns a tuple with the DigitalTiltRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetDigitalTiltRangeOk() (*ParameterRange, bool) {
	if o == nil || isNil(o.DigitalTiltRange) {
    return nil, false
	}
	return o.DigitalTiltRange, true
}

// HasDigitalTiltRange returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasDigitalTiltRange() bool {
	if o != nil && !isNil(o.DigitalTiltRange) {
		return true
	}

	return false
}

// SetDigitalTiltRange gets a reference to the given ParameterRange and assigns it to the DigitalTiltRange field.
func (o *CCOParametersAttrAllOfAttributes) SetDigitalTiltRange(v ParameterRange) {
	o.DigitalTiltRange = &v
}

// GetDigitalAzimuthRange returns the DigitalAzimuthRange field value if set, zero value otherwise.
func (o *CCOParametersAttrAllOfAttributes) GetDigitalAzimuthRange() ParameterRange {
	if o == nil || isNil(o.DigitalAzimuthRange) {
		var ret ParameterRange
		return ret
	}
	return *o.DigitalAzimuthRange
}

// GetDigitalAzimuthRangeOk returns a tuple with the DigitalAzimuthRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CCOParametersAttrAllOfAttributes) GetDigitalAzimuthRangeOk() (*ParameterRange, bool) {
	if o == nil || isNil(o.DigitalAzimuthRange) {
    return nil, false
	}
	return o.DigitalAzimuthRange, true
}

// HasDigitalAzimuthRange returns a boolean if a field has been set.
func (o *CCOParametersAttrAllOfAttributes) HasDigitalAzimuthRange() bool {
	if o != nil && !isNil(o.DigitalAzimuthRange) {
		return true
	}

	return false
}

// SetDigitalAzimuthRange gets a reference to the given ParameterRange and assigns it to the DigitalAzimuthRange field.
func (o *CCOParametersAttrAllOfAttributes) SetDigitalAzimuthRange(v ParameterRange) {
	o.DigitalAzimuthRange = &v
}

func (o CCOParametersAttrAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoverageShapeList) {
		toSerialize["coverageShapeList"] = o.CoverageShapeList
	}
	if !isNil(o.DownlinkTransmitPowerRange) {
		toSerialize["downlinkTransmitPowerRange"] = o.DownlinkTransmitPowerRange
	}
	if !isNil(o.AntennaTiltRange) {
		toSerialize["antennaTiltRange"] = o.AntennaTiltRange
	}
	if !isNil(o.AntennaAzimuthRange) {
		toSerialize["antennaAzimuthRange"] = o.AntennaAzimuthRange
	}
	if !isNil(o.DigitalTiltRange) {
		toSerialize["digitalTiltRange"] = o.DigitalTiltRange
	}
	if !isNil(o.DigitalAzimuthRange) {
		toSerialize["digitalAzimuthRange"] = o.DigitalAzimuthRange
	}
	return json.Marshal(toSerialize)
}

type NullableCCOParametersAttrAllOfAttributes struct {
	value *CCOParametersAttrAllOfAttributes
	isSet bool
}

func (v NullableCCOParametersAttrAllOfAttributes) Get() *CCOParametersAttrAllOfAttributes {
	return v.value
}

func (v *NullableCCOParametersAttrAllOfAttributes) Set(val *CCOParametersAttrAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOParametersAttrAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOParametersAttrAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOParametersAttrAllOfAttributes(val *CCOParametersAttrAllOfAttributes) *NullableCCOParametersAttrAllOfAttributes {
	return &NullableCCOParametersAttrAllOfAttributes{value: val, isSet: true}
}

func (v NullableCCOParametersAttrAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOParametersAttrAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


