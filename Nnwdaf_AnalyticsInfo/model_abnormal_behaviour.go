/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// AbnormalBehaviour Represents the abnormal behaviour information.
type AbnormalBehaviour struct {
	Supis []string `json:"supis,omitempty"`
	Excep Exception `json:"excep"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio *int32 `json:"ratio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence *int32 `json:"confidence,omitempty"`
	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty"`
}

// NewAbnormalBehaviour instantiates a new AbnormalBehaviour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbnormalBehaviour(excep Exception) *AbnormalBehaviour {
	this := AbnormalBehaviour{}
	this.Excep = excep
	return &this
}

// NewAbnormalBehaviourWithDefaults instantiates a new AbnormalBehaviour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbnormalBehaviourWithDefaults() *AbnormalBehaviour {
	this := AbnormalBehaviour{}
	return &this
}

// GetSupis returns the Supis field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetSupis() []string {
	if o == nil || isNil(o.Supis) {
		var ret []string
		return ret
	}
	return o.Supis
}

// GetSupisOk returns a tuple with the Supis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetSupisOk() ([]string, bool) {
	if o == nil || isNil(o.Supis) {
    return nil, false
	}
	return o.Supis, true
}

// HasSupis returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasSupis() bool {
	if o != nil && !isNil(o.Supis) {
		return true
	}

	return false
}

// SetSupis gets a reference to the given []string and assigns it to the Supis field.
func (o *AbnormalBehaviour) SetSupis(v []string) {
	o.Supis = v
}

// GetExcep returns the Excep field value
func (o *AbnormalBehaviour) GetExcep() Exception {
	if o == nil {
		var ret Exception
		return ret
	}

	return o.Excep
}

// GetExcepOk returns a tuple with the Excep field value
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetExcepOk() (*Exception, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Excep, true
}

// SetExcep sets field value
func (o *AbnormalBehaviour) SetExcep(v Exception) {
	o.Excep = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *AbnormalBehaviour) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *AbnormalBehaviour) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetRatio() int32 {
	if o == nil || isNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetRatioOk() (*int32, bool) {
	if o == nil || isNil(o.Ratio) {
    return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasRatio() bool {
	if o != nil && !isNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *AbnormalBehaviour) SetRatio(v int32) {
	o.Ratio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetConfidence() int32 {
	if o == nil || isNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetConfidenceOk() (*int32, bool) {
	if o == nil || isNil(o.Confidence) {
    return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasConfidence() bool {
	if o != nil && !isNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *AbnormalBehaviour) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetAddtMeasInfo returns the AddtMeasInfo field value if set, zero value otherwise.
func (o *AbnormalBehaviour) GetAddtMeasInfo() AdditionalMeasurement {
	if o == nil || isNil(o.AddtMeasInfo) {
		var ret AdditionalMeasurement
		return ret
	}
	return *o.AddtMeasInfo
}

// GetAddtMeasInfoOk returns a tuple with the AddtMeasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalBehaviour) GetAddtMeasInfoOk() (*AdditionalMeasurement, bool) {
	if o == nil || isNil(o.AddtMeasInfo) {
    return nil, false
	}
	return o.AddtMeasInfo, true
}

// HasAddtMeasInfo returns a boolean if a field has been set.
func (o *AbnormalBehaviour) HasAddtMeasInfo() bool {
	if o != nil && !isNil(o.AddtMeasInfo) {
		return true
	}

	return false
}

// SetAddtMeasInfo gets a reference to the given AdditionalMeasurement and assigns it to the AddtMeasInfo field.
func (o *AbnormalBehaviour) SetAddtMeasInfo(v AdditionalMeasurement) {
	o.AddtMeasInfo = &v
}

func (o AbnormalBehaviour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Supis) {
		toSerialize["supis"] = o.Supis
	}
	if true {
		toSerialize["excep"] = o.Excep
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !isNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !isNil(o.AddtMeasInfo) {
		toSerialize["addtMeasInfo"] = o.AddtMeasInfo
	}
	return json.Marshal(toSerialize)
}

type NullableAbnormalBehaviour struct {
	value *AbnormalBehaviour
	isSet bool
}

func (v NullableAbnormalBehaviour) Get() *AbnormalBehaviour {
	return v.value
}

func (v *NullableAbnormalBehaviour) Set(val *AbnormalBehaviour) {
	v.value = val
	v.isSet = true
}

func (v NullableAbnormalBehaviour) IsSet() bool {
	return v.isSet
}

func (v *NullableAbnormalBehaviour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbnormalBehaviour(val *AbnormalBehaviour) *NullableAbnormalBehaviour {
	return &NullableAbnormalBehaviour{value: val, isSet: true}
}

func (v NullableAbnormalBehaviour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbnormalBehaviour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


