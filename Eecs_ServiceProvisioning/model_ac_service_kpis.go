/*
Eecs_ServiceProvisioning

API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_ServiceProvisioning

import (
	"encoding/json"
)

// ACServiceKPIs EAS details.
type ACServiceKPIs struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConnBand *string `json:"connBand,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	ReqRate *int32 `json:"reqRate,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	RespTime *int32 `json:"respTime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Avail *int32 `json:"avail,omitempty"`
	// The compute resources required by the AC.
	ReqComp *string `json:"reqComp,omitempty"`
	// The graphical compute resources required by the AC.
	ReqGrapComp *string `json:"reqGrapComp,omitempty"`
	// The memory resources required by the AC.
	ReqMem *string `json:"reqMem,omitempty"`
	// The storage resources required by the AC.
	ReqStrg *string `json:"reqStrg,omitempty"`
}

// NewACServiceKPIs instantiates a new ACServiceKPIs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACServiceKPIs() *ACServiceKPIs {
	this := ACServiceKPIs{}
	return &this
}

// NewACServiceKPIsWithDefaults instantiates a new ACServiceKPIs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACServiceKPIsWithDefaults() *ACServiceKPIs {
	this := ACServiceKPIs{}
	return &this
}

// GetConnBand returns the ConnBand field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetConnBand() string {
	if o == nil || isNil(o.ConnBand) {
		var ret string
		return ret
	}
	return *o.ConnBand
}

// GetConnBandOk returns a tuple with the ConnBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetConnBandOk() (*string, bool) {
	if o == nil || isNil(o.ConnBand) {
    return nil, false
	}
	return o.ConnBand, true
}

// HasConnBand returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasConnBand() bool {
	if o != nil && !isNil(o.ConnBand) {
		return true
	}

	return false
}

// SetConnBand gets a reference to the given string and assigns it to the ConnBand field.
func (o *ACServiceKPIs) SetConnBand(v string) {
	o.ConnBand = &v
}

// GetReqRate returns the ReqRate field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetReqRate() int32 {
	if o == nil || isNil(o.ReqRate) {
		var ret int32
		return ret
	}
	return *o.ReqRate
}

// GetReqRateOk returns a tuple with the ReqRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetReqRateOk() (*int32, bool) {
	if o == nil || isNil(o.ReqRate) {
    return nil, false
	}
	return o.ReqRate, true
}

// HasReqRate returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasReqRate() bool {
	if o != nil && !isNil(o.ReqRate) {
		return true
	}

	return false
}

// SetReqRate gets a reference to the given int32 and assigns it to the ReqRate field.
func (o *ACServiceKPIs) SetReqRate(v int32) {
	o.ReqRate = &v
}

// GetRespTime returns the RespTime field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetRespTime() int32 {
	if o == nil || isNil(o.RespTime) {
		var ret int32
		return ret
	}
	return *o.RespTime
}

// GetRespTimeOk returns a tuple with the RespTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetRespTimeOk() (*int32, bool) {
	if o == nil || isNil(o.RespTime) {
    return nil, false
	}
	return o.RespTime, true
}

// HasRespTime returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasRespTime() bool {
	if o != nil && !isNil(o.RespTime) {
		return true
	}

	return false
}

// SetRespTime gets a reference to the given int32 and assigns it to the RespTime field.
func (o *ACServiceKPIs) SetRespTime(v int32) {
	o.RespTime = &v
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetAvail() int32 {
	if o == nil || isNil(o.Avail) {
		var ret int32
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetAvailOk() (*int32, bool) {
	if o == nil || isNil(o.Avail) {
    return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasAvail() bool {
	if o != nil && !isNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given int32 and assigns it to the Avail field.
func (o *ACServiceKPIs) SetAvail(v int32) {
	o.Avail = &v
}

// GetReqComp returns the ReqComp field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetReqComp() string {
	if o == nil || isNil(o.ReqComp) {
		var ret string
		return ret
	}
	return *o.ReqComp
}

// GetReqCompOk returns a tuple with the ReqComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetReqCompOk() (*string, bool) {
	if o == nil || isNil(o.ReqComp) {
    return nil, false
	}
	return o.ReqComp, true
}

// HasReqComp returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasReqComp() bool {
	if o != nil && !isNil(o.ReqComp) {
		return true
	}

	return false
}

// SetReqComp gets a reference to the given string and assigns it to the ReqComp field.
func (o *ACServiceKPIs) SetReqComp(v string) {
	o.ReqComp = &v
}

// GetReqGrapComp returns the ReqGrapComp field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetReqGrapComp() string {
	if o == nil || isNil(o.ReqGrapComp) {
		var ret string
		return ret
	}
	return *o.ReqGrapComp
}

// GetReqGrapCompOk returns a tuple with the ReqGrapComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetReqGrapCompOk() (*string, bool) {
	if o == nil || isNil(o.ReqGrapComp) {
    return nil, false
	}
	return o.ReqGrapComp, true
}

// HasReqGrapComp returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasReqGrapComp() bool {
	if o != nil && !isNil(o.ReqGrapComp) {
		return true
	}

	return false
}

// SetReqGrapComp gets a reference to the given string and assigns it to the ReqGrapComp field.
func (o *ACServiceKPIs) SetReqGrapComp(v string) {
	o.ReqGrapComp = &v
}

// GetReqMem returns the ReqMem field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetReqMem() string {
	if o == nil || isNil(o.ReqMem) {
		var ret string
		return ret
	}
	return *o.ReqMem
}

// GetReqMemOk returns a tuple with the ReqMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetReqMemOk() (*string, bool) {
	if o == nil || isNil(o.ReqMem) {
    return nil, false
	}
	return o.ReqMem, true
}

// HasReqMem returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasReqMem() bool {
	if o != nil && !isNil(o.ReqMem) {
		return true
	}

	return false
}

// SetReqMem gets a reference to the given string and assigns it to the ReqMem field.
func (o *ACServiceKPIs) SetReqMem(v string) {
	o.ReqMem = &v
}

// GetReqStrg returns the ReqStrg field value if set, zero value otherwise.
func (o *ACServiceKPIs) GetReqStrg() string {
	if o == nil || isNil(o.ReqStrg) {
		var ret string
		return ret
	}
	return *o.ReqStrg
}

// GetReqStrgOk returns a tuple with the ReqStrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACServiceKPIs) GetReqStrgOk() (*string, bool) {
	if o == nil || isNil(o.ReqStrg) {
    return nil, false
	}
	return o.ReqStrg, true
}

// HasReqStrg returns a boolean if a field has been set.
func (o *ACServiceKPIs) HasReqStrg() bool {
	if o != nil && !isNil(o.ReqStrg) {
		return true
	}

	return false
}

// SetReqStrg gets a reference to the given string and assigns it to the ReqStrg field.
func (o *ACServiceKPIs) SetReqStrg(v string) {
	o.ReqStrg = &v
}

func (o ACServiceKPIs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConnBand) {
		toSerialize["connBand"] = o.ConnBand
	}
	if !isNil(o.ReqRate) {
		toSerialize["reqRate"] = o.ReqRate
	}
	if !isNil(o.RespTime) {
		toSerialize["respTime"] = o.RespTime
	}
	if !isNil(o.Avail) {
		toSerialize["avail"] = o.Avail
	}
	if !isNil(o.ReqComp) {
		toSerialize["reqComp"] = o.ReqComp
	}
	if !isNil(o.ReqGrapComp) {
		toSerialize["reqGrapComp"] = o.ReqGrapComp
	}
	if !isNil(o.ReqMem) {
		toSerialize["reqMem"] = o.ReqMem
	}
	if !isNil(o.ReqStrg) {
		toSerialize["reqStrg"] = o.ReqStrg
	}
	return json.Marshal(toSerialize)
}

type NullableACServiceKPIs struct {
	value *ACServiceKPIs
	isSet bool
}

func (v NullableACServiceKPIs) Get() *ACServiceKPIs {
	return v.value
}

func (v *NullableACServiceKPIs) Set(val *ACServiceKPIs) {
	v.value = val
	v.isSet = true
}

func (v NullableACServiceKPIs) IsSet() bool {
	return v.isSet
}

func (v *NullableACServiceKPIs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACServiceKPIs(val *ACServiceKPIs) *NullableACServiceKPIs {
	return &NullableACServiceKPIs{value: val, isSet: true}
}

func (v NullableACServiceKPIs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACServiceKPIs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


