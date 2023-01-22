/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// StreamInfoTypeAdditionalInfo - struct for StreamInfoTypeAdditionalInfo
type StreamInfoTypeAdditionalInfo struct {
	AnalyticsInfoType *AnalyticsInfoType
	PerformanceInfoType *PerformanceInfoType
	TraceInfoType *TraceInfoType
	VsDataContainerType *VsDataContainerType
}

// AnalyticsInfoTypeAsStreamInfoTypeAdditionalInfo is a convenience function that returns AnalyticsInfoType wrapped in StreamInfoTypeAdditionalInfo
func AnalyticsInfoTypeAsStreamInfoTypeAdditionalInfo(v *AnalyticsInfoType) StreamInfoTypeAdditionalInfo {
	return StreamInfoTypeAdditionalInfo{
		AnalyticsInfoType: v,
	}
}

// PerformanceInfoTypeAsStreamInfoTypeAdditionalInfo is a convenience function that returns PerformanceInfoType wrapped in StreamInfoTypeAdditionalInfo
func PerformanceInfoTypeAsStreamInfoTypeAdditionalInfo(v *PerformanceInfoType) StreamInfoTypeAdditionalInfo {
	return StreamInfoTypeAdditionalInfo{
		PerformanceInfoType: v,
	}
}

// TraceInfoTypeAsStreamInfoTypeAdditionalInfo is a convenience function that returns TraceInfoType wrapped in StreamInfoTypeAdditionalInfo
func TraceInfoTypeAsStreamInfoTypeAdditionalInfo(v *TraceInfoType) StreamInfoTypeAdditionalInfo {
	return StreamInfoTypeAdditionalInfo{
		TraceInfoType: v,
	}
}

// VsDataContainerTypeAsStreamInfoTypeAdditionalInfo is a convenience function that returns VsDataContainerType wrapped in StreamInfoTypeAdditionalInfo
func VsDataContainerTypeAsStreamInfoTypeAdditionalInfo(v *VsDataContainerType) StreamInfoTypeAdditionalInfo {
	return StreamInfoTypeAdditionalInfo{
		VsDataContainerType: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *StreamInfoTypeAdditionalInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnalyticsInfoType
	err = newStrictDecoder(data).Decode(&dst.AnalyticsInfoType)
	if err == nil {
		jsonAnalyticsInfoType, _ := json.Marshal(dst.AnalyticsInfoType)
		if string(jsonAnalyticsInfoType) == "{}" { // empty struct
			dst.AnalyticsInfoType = nil
		} else {
			match++
		}
	} else {
		dst.AnalyticsInfoType = nil
	}

	// try to unmarshal data into PerformanceInfoType
	err = newStrictDecoder(data).Decode(&dst.PerformanceInfoType)
	if err == nil {
		jsonPerformanceInfoType, _ := json.Marshal(dst.PerformanceInfoType)
		if string(jsonPerformanceInfoType) == "{}" { // empty struct
			dst.PerformanceInfoType = nil
		} else {
			match++
		}
	} else {
		dst.PerformanceInfoType = nil
	}

	// try to unmarshal data into TraceInfoType
	err = newStrictDecoder(data).Decode(&dst.TraceInfoType)
	if err == nil {
		jsonTraceInfoType, _ := json.Marshal(dst.TraceInfoType)
		if string(jsonTraceInfoType) == "{}" { // empty struct
			dst.TraceInfoType = nil
		} else {
			match++
		}
	} else {
		dst.TraceInfoType = nil
	}

	// try to unmarshal data into VsDataContainerType
	err = newStrictDecoder(data).Decode(&dst.VsDataContainerType)
	if err == nil {
		jsonVsDataContainerType, _ := json.Marshal(dst.VsDataContainerType)
		if string(jsonVsDataContainerType) == "{}" { // empty struct
			dst.VsDataContainerType = nil
		} else {
			match++
		}
	} else {
		dst.VsDataContainerType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AnalyticsInfoType = nil
		dst.PerformanceInfoType = nil
		dst.TraceInfoType = nil
		dst.VsDataContainerType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(StreamInfoTypeAdditionalInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StreamInfoTypeAdditionalInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StreamInfoTypeAdditionalInfo) MarshalJSON() ([]byte, error) {
	if src.AnalyticsInfoType != nil {
		return json.Marshal(&src.AnalyticsInfoType)
	}

	if src.PerformanceInfoType != nil {
		return json.Marshal(&src.PerformanceInfoType)
	}

	if src.TraceInfoType != nil {
		return json.Marshal(&src.TraceInfoType)
	}

	if src.VsDataContainerType != nil {
		return json.Marshal(&src.VsDataContainerType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StreamInfoTypeAdditionalInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AnalyticsInfoType != nil {
		return obj.AnalyticsInfoType
	}

	if obj.PerformanceInfoType != nil {
		return obj.PerformanceInfoType
	}

	if obj.TraceInfoType != nil {
		return obj.TraceInfoType
	}

	if obj.VsDataContainerType != nil {
		return obj.VsDataContainerType
	}

	// all schemas are nil
	return nil
}

type NullableStreamInfoTypeAdditionalInfo struct {
	value *StreamInfoTypeAdditionalInfo
	isSet bool
}

func (v NullableStreamInfoTypeAdditionalInfo) Get() *StreamInfoTypeAdditionalInfo {
	return v.value
}

func (v *NullableStreamInfoTypeAdditionalInfo) Set(val *StreamInfoTypeAdditionalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamInfoTypeAdditionalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamInfoTypeAdditionalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamInfoTypeAdditionalInfo(val *StreamInfoTypeAdditionalInfo) *NullableStreamInfoTypeAdditionalInfo {
	return &NullableStreamInfoTypeAdditionalInfo{value: val, isSet: true}
}

func (v NullableStreamInfoTypeAdditionalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamInfoTypeAdditionalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

