/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// SubscrCond - Condition to determine the set of NFs to monitor under a certain subscription in NRF 
type SubscrCond struct {
	AmfCond *AmfCond
	DccfCond *DccfCond
	GuamiListCond *GuamiListCond
	NefCond *NefCond
	NetworkSliceCond *NetworkSliceCond
	NfGroupCond *NfGroupCond
	NfGroupListCond *NfGroupListCond
	NfInstanceIdCond *NfInstanceIdCond
	NfInstanceIdListCond *NfInstanceIdListCond
	NfServiceSetCond *NfServiceSetCond
	NfSetCond *NfSetCond
	NfTypeCond *NfTypeCond
	NwdafCond *NwdafCond
	ScpDomainCond *ScpDomainCond
	ServiceNameCond *ServiceNameCond
	ServiceNameListCond *ServiceNameListCond
	UpfCond *UpfCond
}

// AmfCondAsSubscrCond is a convenience function that returns AmfCond wrapped in SubscrCond
func AmfCondAsSubscrCond(v *AmfCond) SubscrCond {
	return SubscrCond{
		AmfCond: v,
	}
}

// DccfCondAsSubscrCond is a convenience function that returns DccfCond wrapped in SubscrCond
func DccfCondAsSubscrCond(v *DccfCond) SubscrCond {
	return SubscrCond{
		DccfCond: v,
	}
}

// GuamiListCondAsSubscrCond is a convenience function that returns GuamiListCond wrapped in SubscrCond
func GuamiListCondAsSubscrCond(v *GuamiListCond) SubscrCond {
	return SubscrCond{
		GuamiListCond: v,
	}
}

// NefCondAsSubscrCond is a convenience function that returns NefCond wrapped in SubscrCond
func NefCondAsSubscrCond(v *NefCond) SubscrCond {
	return SubscrCond{
		NefCond: v,
	}
}

// NetworkSliceCondAsSubscrCond is a convenience function that returns NetworkSliceCond wrapped in SubscrCond
func NetworkSliceCondAsSubscrCond(v *NetworkSliceCond) SubscrCond {
	return SubscrCond{
		NetworkSliceCond: v,
	}
}

// NfGroupCondAsSubscrCond is a convenience function that returns NfGroupCond wrapped in SubscrCond
func NfGroupCondAsSubscrCond(v *NfGroupCond) SubscrCond {
	return SubscrCond{
		NfGroupCond: v,
	}
}

// NfGroupListCondAsSubscrCond is a convenience function that returns NfGroupListCond wrapped in SubscrCond
func NfGroupListCondAsSubscrCond(v *NfGroupListCond) SubscrCond {
	return SubscrCond{
		NfGroupListCond: v,
	}
}

// NfInstanceIdCondAsSubscrCond is a convenience function that returns NfInstanceIdCond wrapped in SubscrCond
func NfInstanceIdCondAsSubscrCond(v *NfInstanceIdCond) SubscrCond {
	return SubscrCond{
		NfInstanceIdCond: v,
	}
}

// NfInstanceIdListCondAsSubscrCond is a convenience function that returns NfInstanceIdListCond wrapped in SubscrCond
func NfInstanceIdListCondAsSubscrCond(v *NfInstanceIdListCond) SubscrCond {
	return SubscrCond{
		NfInstanceIdListCond: v,
	}
}

// NfServiceSetCondAsSubscrCond is a convenience function that returns NfServiceSetCond wrapped in SubscrCond
func NfServiceSetCondAsSubscrCond(v *NfServiceSetCond) SubscrCond {
	return SubscrCond{
		NfServiceSetCond: v,
	}
}

// NfSetCondAsSubscrCond is a convenience function that returns NfSetCond wrapped in SubscrCond
func NfSetCondAsSubscrCond(v *NfSetCond) SubscrCond {
	return SubscrCond{
		NfSetCond: v,
	}
}

// NfTypeCondAsSubscrCond is a convenience function that returns NfTypeCond wrapped in SubscrCond
func NfTypeCondAsSubscrCond(v *NfTypeCond) SubscrCond {
	return SubscrCond{
		NfTypeCond: v,
	}
}

// NwdafCondAsSubscrCond is a convenience function that returns NwdafCond wrapped in SubscrCond
func NwdafCondAsSubscrCond(v *NwdafCond) SubscrCond {
	return SubscrCond{
		NwdafCond: v,
	}
}

// ScpDomainCondAsSubscrCond is a convenience function that returns ScpDomainCond wrapped in SubscrCond
func ScpDomainCondAsSubscrCond(v *ScpDomainCond) SubscrCond {
	return SubscrCond{
		ScpDomainCond: v,
	}
}

// ServiceNameCondAsSubscrCond is a convenience function that returns ServiceNameCond wrapped in SubscrCond
func ServiceNameCondAsSubscrCond(v *ServiceNameCond) SubscrCond {
	return SubscrCond{
		ServiceNameCond: v,
	}
}

// ServiceNameListCondAsSubscrCond is a convenience function that returns ServiceNameListCond wrapped in SubscrCond
func ServiceNameListCondAsSubscrCond(v *ServiceNameListCond) SubscrCond {
	return SubscrCond{
		ServiceNameListCond: v,
	}
}

// UpfCondAsSubscrCond is a convenience function that returns UpfCond wrapped in SubscrCond
func UpfCondAsSubscrCond(v *UpfCond) SubscrCond {
	return SubscrCond{
		UpfCond: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscrCond) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AmfCond
	err = newStrictDecoder(data).Decode(&dst.AmfCond)
	if err == nil {
		jsonAmfCond, _ := json.Marshal(dst.AmfCond)
		if string(jsonAmfCond) == "{}" { // empty struct
			dst.AmfCond = nil
		} else {
			match++
		}
	} else {
		dst.AmfCond = nil
	}

	// try to unmarshal data into DccfCond
	err = newStrictDecoder(data).Decode(&dst.DccfCond)
	if err == nil {
		jsonDccfCond, _ := json.Marshal(dst.DccfCond)
		if string(jsonDccfCond) == "{}" { // empty struct
			dst.DccfCond = nil
		} else {
			match++
		}
	} else {
		dst.DccfCond = nil
	}

	// try to unmarshal data into GuamiListCond
	err = newStrictDecoder(data).Decode(&dst.GuamiListCond)
	if err == nil {
		jsonGuamiListCond, _ := json.Marshal(dst.GuamiListCond)
		if string(jsonGuamiListCond) == "{}" { // empty struct
			dst.GuamiListCond = nil
		} else {
			match++
		}
	} else {
		dst.GuamiListCond = nil
	}

	// try to unmarshal data into NefCond
	err = newStrictDecoder(data).Decode(&dst.NefCond)
	if err == nil {
		jsonNefCond, _ := json.Marshal(dst.NefCond)
		if string(jsonNefCond) == "{}" { // empty struct
			dst.NefCond = nil
		} else {
			match++
		}
	} else {
		dst.NefCond = nil
	}

	// try to unmarshal data into NetworkSliceCond
	err = newStrictDecoder(data).Decode(&dst.NetworkSliceCond)
	if err == nil {
		jsonNetworkSliceCond, _ := json.Marshal(dst.NetworkSliceCond)
		if string(jsonNetworkSliceCond) == "{}" { // empty struct
			dst.NetworkSliceCond = nil
		} else {
			match++
		}
	} else {
		dst.NetworkSliceCond = nil
	}

	// try to unmarshal data into NfGroupCond
	err = newStrictDecoder(data).Decode(&dst.NfGroupCond)
	if err == nil {
		jsonNfGroupCond, _ := json.Marshal(dst.NfGroupCond)
		if string(jsonNfGroupCond) == "{}" { // empty struct
			dst.NfGroupCond = nil
		} else {
			match++
		}
	} else {
		dst.NfGroupCond = nil
	}

	// try to unmarshal data into NfGroupListCond
	err = newStrictDecoder(data).Decode(&dst.NfGroupListCond)
	if err == nil {
		jsonNfGroupListCond, _ := json.Marshal(dst.NfGroupListCond)
		if string(jsonNfGroupListCond) == "{}" { // empty struct
			dst.NfGroupListCond = nil
		} else {
			match++
		}
	} else {
		dst.NfGroupListCond = nil
	}

	// try to unmarshal data into NfInstanceIdCond
	err = newStrictDecoder(data).Decode(&dst.NfInstanceIdCond)
	if err == nil {
		jsonNfInstanceIdCond, _ := json.Marshal(dst.NfInstanceIdCond)
		if string(jsonNfInstanceIdCond) == "{}" { // empty struct
			dst.NfInstanceIdCond = nil
		} else {
			match++
		}
	} else {
		dst.NfInstanceIdCond = nil
	}

	// try to unmarshal data into NfInstanceIdListCond
	err = newStrictDecoder(data).Decode(&dst.NfInstanceIdListCond)
	if err == nil {
		jsonNfInstanceIdListCond, _ := json.Marshal(dst.NfInstanceIdListCond)
		if string(jsonNfInstanceIdListCond) == "{}" { // empty struct
			dst.NfInstanceIdListCond = nil
		} else {
			match++
		}
	} else {
		dst.NfInstanceIdListCond = nil
	}

	// try to unmarshal data into NfServiceSetCond
	err = newStrictDecoder(data).Decode(&dst.NfServiceSetCond)
	if err == nil {
		jsonNfServiceSetCond, _ := json.Marshal(dst.NfServiceSetCond)
		if string(jsonNfServiceSetCond) == "{}" { // empty struct
			dst.NfServiceSetCond = nil
		} else {
			match++
		}
	} else {
		dst.NfServiceSetCond = nil
	}

	// try to unmarshal data into NfSetCond
	err = newStrictDecoder(data).Decode(&dst.NfSetCond)
	if err == nil {
		jsonNfSetCond, _ := json.Marshal(dst.NfSetCond)
		if string(jsonNfSetCond) == "{}" { // empty struct
			dst.NfSetCond = nil
		} else {
			match++
		}
	} else {
		dst.NfSetCond = nil
	}

	// try to unmarshal data into NfTypeCond
	err = newStrictDecoder(data).Decode(&dst.NfTypeCond)
	if err == nil {
		jsonNfTypeCond, _ := json.Marshal(dst.NfTypeCond)
		if string(jsonNfTypeCond) == "{}" { // empty struct
			dst.NfTypeCond = nil
		} else {
			match++
		}
	} else {
		dst.NfTypeCond = nil
	}

	// try to unmarshal data into NwdafCond
	err = newStrictDecoder(data).Decode(&dst.NwdafCond)
	if err == nil {
		jsonNwdafCond, _ := json.Marshal(dst.NwdafCond)
		if string(jsonNwdafCond) == "{}" { // empty struct
			dst.NwdafCond = nil
		} else {
			match++
		}
	} else {
		dst.NwdafCond = nil
	}

	// try to unmarshal data into ScpDomainCond
	err = newStrictDecoder(data).Decode(&dst.ScpDomainCond)
	if err == nil {
		jsonScpDomainCond, _ := json.Marshal(dst.ScpDomainCond)
		if string(jsonScpDomainCond) == "{}" { // empty struct
			dst.ScpDomainCond = nil
		} else {
			match++
		}
	} else {
		dst.ScpDomainCond = nil
	}

	// try to unmarshal data into ServiceNameCond
	err = newStrictDecoder(data).Decode(&dst.ServiceNameCond)
	if err == nil {
		jsonServiceNameCond, _ := json.Marshal(dst.ServiceNameCond)
		if string(jsonServiceNameCond) == "{}" { // empty struct
			dst.ServiceNameCond = nil
		} else {
			match++
		}
	} else {
		dst.ServiceNameCond = nil
	}

	// try to unmarshal data into ServiceNameListCond
	err = newStrictDecoder(data).Decode(&dst.ServiceNameListCond)
	if err == nil {
		jsonServiceNameListCond, _ := json.Marshal(dst.ServiceNameListCond)
		if string(jsonServiceNameListCond) == "{}" { // empty struct
			dst.ServiceNameListCond = nil
		} else {
			match++
		}
	} else {
		dst.ServiceNameListCond = nil
	}

	// try to unmarshal data into UpfCond
	err = newStrictDecoder(data).Decode(&dst.UpfCond)
	if err == nil {
		jsonUpfCond, _ := json.Marshal(dst.UpfCond)
		if string(jsonUpfCond) == "{}" { // empty struct
			dst.UpfCond = nil
		} else {
			match++
		}
	} else {
		dst.UpfCond = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AmfCond = nil
		dst.DccfCond = nil
		dst.GuamiListCond = nil
		dst.NefCond = nil
		dst.NetworkSliceCond = nil
		dst.NfGroupCond = nil
		dst.NfGroupListCond = nil
		dst.NfInstanceIdCond = nil
		dst.NfInstanceIdListCond = nil
		dst.NfServiceSetCond = nil
		dst.NfSetCond = nil
		dst.NfTypeCond = nil
		dst.NwdafCond = nil
		dst.ScpDomainCond = nil
		dst.ServiceNameCond = nil
		dst.ServiceNameListCond = nil
		dst.UpfCond = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscrCond)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscrCond)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscrCond) MarshalJSON() ([]byte, error) {
	if src.AmfCond != nil {
		return json.Marshal(&src.AmfCond)
	}

	if src.DccfCond != nil {
		return json.Marshal(&src.DccfCond)
	}

	if src.GuamiListCond != nil {
		return json.Marshal(&src.GuamiListCond)
	}

	if src.NefCond != nil {
		return json.Marshal(&src.NefCond)
	}

	if src.NetworkSliceCond != nil {
		return json.Marshal(&src.NetworkSliceCond)
	}

	if src.NfGroupCond != nil {
		return json.Marshal(&src.NfGroupCond)
	}

	if src.NfGroupListCond != nil {
		return json.Marshal(&src.NfGroupListCond)
	}

	if src.NfInstanceIdCond != nil {
		return json.Marshal(&src.NfInstanceIdCond)
	}

	if src.NfInstanceIdListCond != nil {
		return json.Marshal(&src.NfInstanceIdListCond)
	}

	if src.NfServiceSetCond != nil {
		return json.Marshal(&src.NfServiceSetCond)
	}

	if src.NfSetCond != nil {
		return json.Marshal(&src.NfSetCond)
	}

	if src.NfTypeCond != nil {
		return json.Marshal(&src.NfTypeCond)
	}

	if src.NwdafCond != nil {
		return json.Marshal(&src.NwdafCond)
	}

	if src.ScpDomainCond != nil {
		return json.Marshal(&src.ScpDomainCond)
	}

	if src.ServiceNameCond != nil {
		return json.Marshal(&src.ServiceNameCond)
	}

	if src.ServiceNameListCond != nil {
		return json.Marshal(&src.ServiceNameListCond)
	}

	if src.UpfCond != nil {
		return json.Marshal(&src.UpfCond)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscrCond) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AmfCond != nil {
		return obj.AmfCond
	}

	if obj.DccfCond != nil {
		return obj.DccfCond
	}

	if obj.GuamiListCond != nil {
		return obj.GuamiListCond
	}

	if obj.NefCond != nil {
		return obj.NefCond
	}

	if obj.NetworkSliceCond != nil {
		return obj.NetworkSliceCond
	}

	if obj.NfGroupCond != nil {
		return obj.NfGroupCond
	}

	if obj.NfGroupListCond != nil {
		return obj.NfGroupListCond
	}

	if obj.NfInstanceIdCond != nil {
		return obj.NfInstanceIdCond
	}

	if obj.NfInstanceIdListCond != nil {
		return obj.NfInstanceIdListCond
	}

	if obj.NfServiceSetCond != nil {
		return obj.NfServiceSetCond
	}

	if obj.NfSetCond != nil {
		return obj.NfSetCond
	}

	if obj.NfTypeCond != nil {
		return obj.NfTypeCond
	}

	if obj.NwdafCond != nil {
		return obj.NwdafCond
	}

	if obj.ScpDomainCond != nil {
		return obj.ScpDomainCond
	}

	if obj.ServiceNameCond != nil {
		return obj.ServiceNameCond
	}

	if obj.ServiceNameListCond != nil {
		return obj.ServiceNameListCond
	}

	if obj.UpfCond != nil {
		return obj.UpfCond
	}

	// all schemas are nil
	return nil
}

type NullableSubscrCond struct {
	value *SubscrCond
	isSet bool
}

func (v NullableSubscrCond) Get() *SubscrCond {
	return v.value
}

func (v *NullableSubscrCond) Set(val *SubscrCond) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscrCond) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscrCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscrCond(val *SubscrCond) *NullableSubscrCond {
	return &NullableSubscrCond{value: val, isSet: true}
}

func (v NullableSubscrCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscrCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


