/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// ResourceReservationStatus An attribute which specifies the resource reservation result for the feasibility check job.
type ResourceReservationStatus string

// List of ResourceReservationStatus
const (
	RESERVED ResourceReservationStatus = "RESERVED"
	UNRESERVED ResourceReservationStatus = "UNRESERVED"
	USED ResourceReservationStatus = "USED"
)

// All allowed values of ResourceReservationStatus enum
var AllowedResourceReservationStatusEnumValues = []ResourceReservationStatus{
	"RESERVED",
	"UNRESERVED",
	"USED",
}

func (v *ResourceReservationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceReservationStatus(value)
	for _, existing := range AllowedResourceReservationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceReservationStatus", value)
}

// NewResourceReservationStatusFromValue returns a pointer to a valid ResourceReservationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceReservationStatusFromValue(v string) (*ResourceReservationStatus, error) {
	ev := ResourceReservationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceReservationStatus: valid values are %v", v, AllowedResourceReservationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceReservationStatus) IsValid() bool {
	for _, existing := range AllowedResourceReservationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ResourceReservationStatus value
func (v ResourceReservationStatus) Ptr() *ResourceReservationStatus {
	return &v
}

type NullableResourceReservationStatus struct {
	value *ResourceReservationStatus
	isSet bool
}

func (v NullableResourceReservationStatus) Get() *ResourceReservationStatus {
	return v.value
}

func (v *NullableResourceReservationStatus) Set(val *ResourceReservationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceReservationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceReservationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceReservationStatus(val *ResourceReservationStatus) *NullableResourceReservationStatus {
	return &NullableResourceReservationStatus{value: val, isSet: true}
}

func (v NullableResourceReservationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceReservationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

