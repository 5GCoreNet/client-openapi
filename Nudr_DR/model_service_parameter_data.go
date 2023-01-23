/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// ServiceParameterData Represents the service parameter data.
type ServiceParameterData struct {
	// Identifies an application.
	AppId *string `json:"appId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId *string `json:"interGroupId,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	UeIpv4 *string `json:"ueIpv4,omitempty"`
	// string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used.
	UeIpv6 *string `json:"ueIpv6,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac *string `json:"ueMac,omitempty"`
	AnyUeInd *bool `json:"anyUeInd,omitempty"`
	// Represents configuration parameters for V2X communications over PC5 reference point. 
	ParamOverPc5 *string `json:"paramOverPc5,omitempty"`
	// Represents configuration parameters for V2X communications over Uu reference point. 
	ParamOverUu *string `json:"paramOverUu,omitempty"`
	// Represents the service parameters for 5G ProSe direct discovery.
	ParamForProSeDd *string `json:"paramForProSeDd,omitempty"`
	// Represents the service parameters for 5G ProSe direct communications.
	ParamForProSeDc *string `json:"paramForProSeDc,omitempty"`
	// Represents the service parameters for 5G ProSe UE-to-network relay UE.
	ParamForProSeU2NRelUe *string `json:"paramForProSeU2NRelUe,omitempty"`
	// Represents the service parameters for 5G ProSe Remate UE.
	ParamForProSeRemUe *string `json:"paramForProSeRemUe,omitempty"`
	// Contains the service parameter used to guide the URSP.
	UrspGuidance []UrspRuleRequest `json:"urspGuidance,omitempty"`
	// Contains the outcome of the UE Policy Delivery.
	DeliveryEvents []Event `json:"deliveryEvents,omitempty"`
	// Contains the Notification Correlation Id allocated by the NEF for the notification of UE Policy delivery outcome. 
	PolicDelivNotifCorreId *string `json:"policDelivNotifCorreId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PolicDelivNotifUri *string `json:"policDelivNotifUri,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ResUri *string `json:"resUri,omitempty"`
	Headers []string `json:"headers,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewServiceParameterData instantiates a new ServiceParameterData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceParameterData() *ServiceParameterData {
	this := ServiceParameterData{}
	return &this
}

// NewServiceParameterDataWithDefaults instantiates a new ServiceParameterData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceParameterDataWithDefaults() *ServiceParameterData {
	this := ServiceParameterData{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ServiceParameterData) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
    return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ServiceParameterData) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ServiceParameterData) SetAppId(v string) {
	o.AppId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *ServiceParameterData) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *ServiceParameterData) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *ServiceParameterData) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *ServiceParameterData) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *ServiceParameterData) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *ServiceParameterData) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetInterGroupId returns the InterGroupId field value if set, zero value otherwise.
func (o *ServiceParameterData) GetInterGroupId() string {
	if o == nil || isNil(o.InterGroupId) {
		var ret string
		return ret
	}
	return *o.InterGroupId
}

// GetInterGroupIdOk returns a tuple with the InterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetInterGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.InterGroupId) {
    return nil, false
	}
	return o.InterGroupId, true
}

// HasInterGroupId returns a boolean if a field has been set.
func (o *ServiceParameterData) HasInterGroupId() bool {
	if o != nil && !isNil(o.InterGroupId) {
		return true
	}

	return false
}

// SetInterGroupId gets a reference to the given string and assigns it to the InterGroupId field.
func (o *ServiceParameterData) SetInterGroupId(v string) {
	o.InterGroupId = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *ServiceParameterData) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *ServiceParameterData) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *ServiceParameterData) SetSupi(v string) {
	o.Supi = &v
}

// GetUeIpv4 returns the UeIpv4 field value if set, zero value otherwise.
func (o *ServiceParameterData) GetUeIpv4() string {
	if o == nil || isNil(o.UeIpv4) {
		var ret string
		return ret
	}
	return *o.UeIpv4
}

// GetUeIpv4Ok returns a tuple with the UeIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetUeIpv4Ok() (*string, bool) {
	if o == nil || isNil(o.UeIpv4) {
    return nil, false
	}
	return o.UeIpv4, true
}

// HasUeIpv4 returns a boolean if a field has been set.
func (o *ServiceParameterData) HasUeIpv4() bool {
	if o != nil && !isNil(o.UeIpv4) {
		return true
	}

	return false
}

// SetUeIpv4 gets a reference to the given string and assigns it to the UeIpv4 field.
func (o *ServiceParameterData) SetUeIpv4(v string) {
	o.UeIpv4 = &v
}

// GetUeIpv6 returns the UeIpv6 field value if set, zero value otherwise.
func (o *ServiceParameterData) GetUeIpv6() string {
	if o == nil || isNil(o.UeIpv6) {
		var ret string
		return ret
	}
	return *o.UeIpv6
}

// GetUeIpv6Ok returns a tuple with the UeIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetUeIpv6Ok() (*string, bool) {
	if o == nil || isNil(o.UeIpv6) {
    return nil, false
	}
	return o.UeIpv6, true
}

// HasUeIpv6 returns a boolean if a field has been set.
func (o *ServiceParameterData) HasUeIpv6() bool {
	if o != nil && !isNil(o.UeIpv6) {
		return true
	}

	return false
}

// SetUeIpv6 gets a reference to the given string and assigns it to the UeIpv6 field.
func (o *ServiceParameterData) SetUeIpv6(v string) {
	o.UeIpv6 = &v
}

// GetUeMac returns the UeMac field value if set, zero value otherwise.
func (o *ServiceParameterData) GetUeMac() string {
	if o == nil || isNil(o.UeMac) {
		var ret string
		return ret
	}
	return *o.UeMac
}

// GetUeMacOk returns a tuple with the UeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetUeMacOk() (*string, bool) {
	if o == nil || isNil(o.UeMac) {
    return nil, false
	}
	return o.UeMac, true
}

// HasUeMac returns a boolean if a field has been set.
func (o *ServiceParameterData) HasUeMac() bool {
	if o != nil && !isNil(o.UeMac) {
		return true
	}

	return false
}

// SetUeMac gets a reference to the given string and assigns it to the UeMac field.
func (o *ServiceParameterData) SetUeMac(v string) {
	o.UeMac = &v
}

// GetAnyUeInd returns the AnyUeInd field value if set, zero value otherwise.
func (o *ServiceParameterData) GetAnyUeInd() bool {
	if o == nil || isNil(o.AnyUeInd) {
		var ret bool
		return ret
	}
	return *o.AnyUeInd
}

// GetAnyUeIndOk returns a tuple with the AnyUeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetAnyUeIndOk() (*bool, bool) {
	if o == nil || isNil(o.AnyUeInd) {
    return nil, false
	}
	return o.AnyUeInd, true
}

// HasAnyUeInd returns a boolean if a field has been set.
func (o *ServiceParameterData) HasAnyUeInd() bool {
	if o != nil && !isNil(o.AnyUeInd) {
		return true
	}

	return false
}

// SetAnyUeInd gets a reference to the given bool and assigns it to the AnyUeInd field.
func (o *ServiceParameterData) SetAnyUeInd(v bool) {
	o.AnyUeInd = &v
}

// GetParamOverPc5 returns the ParamOverPc5 field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamOverPc5() string {
	if o == nil || isNil(o.ParamOverPc5) {
		var ret string
		return ret
	}
	return *o.ParamOverPc5
}

// GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamOverPc5Ok() (*string, bool) {
	if o == nil || isNil(o.ParamOverPc5) {
    return nil, false
	}
	return o.ParamOverPc5, true
}

// HasParamOverPc5 returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamOverPc5() bool {
	if o != nil && !isNil(o.ParamOverPc5) {
		return true
	}

	return false
}

// SetParamOverPc5 gets a reference to the given string and assigns it to the ParamOverPc5 field.
func (o *ServiceParameterData) SetParamOverPc5(v string) {
	o.ParamOverPc5 = &v
}

// GetParamOverUu returns the ParamOverUu field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamOverUu() string {
	if o == nil || isNil(o.ParamOverUu) {
		var ret string
		return ret
	}
	return *o.ParamOverUu
}

// GetParamOverUuOk returns a tuple with the ParamOverUu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamOverUuOk() (*string, bool) {
	if o == nil || isNil(o.ParamOverUu) {
    return nil, false
	}
	return o.ParamOverUu, true
}

// HasParamOverUu returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamOverUu() bool {
	if o != nil && !isNil(o.ParamOverUu) {
		return true
	}

	return false
}

// SetParamOverUu gets a reference to the given string and assigns it to the ParamOverUu field.
func (o *ServiceParameterData) SetParamOverUu(v string) {
	o.ParamOverUu = &v
}

// GetParamForProSeDd returns the ParamForProSeDd field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamForProSeDd() string {
	if o == nil || isNil(o.ParamForProSeDd) {
		var ret string
		return ret
	}
	return *o.ParamForProSeDd
}

// GetParamForProSeDdOk returns a tuple with the ParamForProSeDd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamForProSeDdOk() (*string, bool) {
	if o == nil || isNil(o.ParamForProSeDd) {
    return nil, false
	}
	return o.ParamForProSeDd, true
}

// HasParamForProSeDd returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamForProSeDd() bool {
	if o != nil && !isNil(o.ParamForProSeDd) {
		return true
	}

	return false
}

// SetParamForProSeDd gets a reference to the given string and assigns it to the ParamForProSeDd field.
func (o *ServiceParameterData) SetParamForProSeDd(v string) {
	o.ParamForProSeDd = &v
}

// GetParamForProSeDc returns the ParamForProSeDc field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamForProSeDc() string {
	if o == nil || isNil(o.ParamForProSeDc) {
		var ret string
		return ret
	}
	return *o.ParamForProSeDc
}

// GetParamForProSeDcOk returns a tuple with the ParamForProSeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamForProSeDcOk() (*string, bool) {
	if o == nil || isNil(o.ParamForProSeDc) {
    return nil, false
	}
	return o.ParamForProSeDc, true
}

// HasParamForProSeDc returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamForProSeDc() bool {
	if o != nil && !isNil(o.ParamForProSeDc) {
		return true
	}

	return false
}

// SetParamForProSeDc gets a reference to the given string and assigns it to the ParamForProSeDc field.
func (o *ServiceParameterData) SetParamForProSeDc(v string) {
	o.ParamForProSeDc = &v
}

// GetParamForProSeU2NRelUe returns the ParamForProSeU2NRelUe field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamForProSeU2NRelUe() string {
	if o == nil || isNil(o.ParamForProSeU2NRelUe) {
		var ret string
		return ret
	}
	return *o.ParamForProSeU2NRelUe
}

// GetParamForProSeU2NRelUeOk returns a tuple with the ParamForProSeU2NRelUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamForProSeU2NRelUeOk() (*string, bool) {
	if o == nil || isNil(o.ParamForProSeU2NRelUe) {
    return nil, false
	}
	return o.ParamForProSeU2NRelUe, true
}

// HasParamForProSeU2NRelUe returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamForProSeU2NRelUe() bool {
	if o != nil && !isNil(o.ParamForProSeU2NRelUe) {
		return true
	}

	return false
}

// SetParamForProSeU2NRelUe gets a reference to the given string and assigns it to the ParamForProSeU2NRelUe field.
func (o *ServiceParameterData) SetParamForProSeU2NRelUe(v string) {
	o.ParamForProSeU2NRelUe = &v
}

// GetParamForProSeRemUe returns the ParamForProSeRemUe field value if set, zero value otherwise.
func (o *ServiceParameterData) GetParamForProSeRemUe() string {
	if o == nil || isNil(o.ParamForProSeRemUe) {
		var ret string
		return ret
	}
	return *o.ParamForProSeRemUe
}

// GetParamForProSeRemUeOk returns a tuple with the ParamForProSeRemUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetParamForProSeRemUeOk() (*string, bool) {
	if o == nil || isNil(o.ParamForProSeRemUe) {
    return nil, false
	}
	return o.ParamForProSeRemUe, true
}

// HasParamForProSeRemUe returns a boolean if a field has been set.
func (o *ServiceParameterData) HasParamForProSeRemUe() bool {
	if o != nil && !isNil(o.ParamForProSeRemUe) {
		return true
	}

	return false
}

// SetParamForProSeRemUe gets a reference to the given string and assigns it to the ParamForProSeRemUe field.
func (o *ServiceParameterData) SetParamForProSeRemUe(v string) {
	o.ParamForProSeRemUe = &v
}

// GetUrspGuidance returns the UrspGuidance field value if set, zero value otherwise.
func (o *ServiceParameterData) GetUrspGuidance() []UrspRuleRequest {
	if o == nil || isNil(o.UrspGuidance) {
		var ret []UrspRuleRequest
		return ret
	}
	return o.UrspGuidance
}

// GetUrspGuidanceOk returns a tuple with the UrspGuidance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetUrspGuidanceOk() ([]UrspRuleRequest, bool) {
	if o == nil || isNil(o.UrspGuidance) {
    return nil, false
	}
	return o.UrspGuidance, true
}

// HasUrspGuidance returns a boolean if a field has been set.
func (o *ServiceParameterData) HasUrspGuidance() bool {
	if o != nil && !isNil(o.UrspGuidance) {
		return true
	}

	return false
}

// SetUrspGuidance gets a reference to the given []UrspRuleRequest and assigns it to the UrspGuidance field.
func (o *ServiceParameterData) SetUrspGuidance(v []UrspRuleRequest) {
	o.UrspGuidance = v
}

// GetDeliveryEvents returns the DeliveryEvents field value if set, zero value otherwise.
func (o *ServiceParameterData) GetDeliveryEvents() []Event {
	if o == nil || isNil(o.DeliveryEvents) {
		var ret []Event
		return ret
	}
	return o.DeliveryEvents
}

// GetDeliveryEventsOk returns a tuple with the DeliveryEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetDeliveryEventsOk() ([]Event, bool) {
	if o == nil || isNil(o.DeliveryEvents) {
    return nil, false
	}
	return o.DeliveryEvents, true
}

// HasDeliveryEvents returns a boolean if a field has been set.
func (o *ServiceParameterData) HasDeliveryEvents() bool {
	if o != nil && !isNil(o.DeliveryEvents) {
		return true
	}

	return false
}

// SetDeliveryEvents gets a reference to the given []Event and assigns it to the DeliveryEvents field.
func (o *ServiceParameterData) SetDeliveryEvents(v []Event) {
	o.DeliveryEvents = v
}

// GetPolicDelivNotifCorreId returns the PolicDelivNotifCorreId field value if set, zero value otherwise.
func (o *ServiceParameterData) GetPolicDelivNotifCorreId() string {
	if o == nil || isNil(o.PolicDelivNotifCorreId) {
		var ret string
		return ret
	}
	return *o.PolicDelivNotifCorreId
}

// GetPolicDelivNotifCorreIdOk returns a tuple with the PolicDelivNotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetPolicDelivNotifCorreIdOk() (*string, bool) {
	if o == nil || isNil(o.PolicDelivNotifCorreId) {
    return nil, false
	}
	return o.PolicDelivNotifCorreId, true
}

// HasPolicDelivNotifCorreId returns a boolean if a field has been set.
func (o *ServiceParameterData) HasPolicDelivNotifCorreId() bool {
	if o != nil && !isNil(o.PolicDelivNotifCorreId) {
		return true
	}

	return false
}

// SetPolicDelivNotifCorreId gets a reference to the given string and assigns it to the PolicDelivNotifCorreId field.
func (o *ServiceParameterData) SetPolicDelivNotifCorreId(v string) {
	o.PolicDelivNotifCorreId = &v
}

// GetPolicDelivNotifUri returns the PolicDelivNotifUri field value if set, zero value otherwise.
func (o *ServiceParameterData) GetPolicDelivNotifUri() string {
	if o == nil || isNil(o.PolicDelivNotifUri) {
		var ret string
		return ret
	}
	return *o.PolicDelivNotifUri
}

// GetPolicDelivNotifUriOk returns a tuple with the PolicDelivNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetPolicDelivNotifUriOk() (*string, bool) {
	if o == nil || isNil(o.PolicDelivNotifUri) {
    return nil, false
	}
	return o.PolicDelivNotifUri, true
}

// HasPolicDelivNotifUri returns a boolean if a field has been set.
func (o *ServiceParameterData) HasPolicDelivNotifUri() bool {
	if o != nil && !isNil(o.PolicDelivNotifUri) {
		return true
	}

	return false
}

// SetPolicDelivNotifUri gets a reference to the given string and assigns it to the PolicDelivNotifUri field.
func (o *ServiceParameterData) SetPolicDelivNotifUri(v string) {
	o.PolicDelivNotifUri = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *ServiceParameterData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *ServiceParameterData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *ServiceParameterData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResUri returns the ResUri field value if set, zero value otherwise.
func (o *ServiceParameterData) GetResUri() string {
	if o == nil || isNil(o.ResUri) {
		var ret string
		return ret
	}
	return *o.ResUri
}

// GetResUriOk returns a tuple with the ResUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetResUriOk() (*string, bool) {
	if o == nil || isNil(o.ResUri) {
    return nil, false
	}
	return o.ResUri, true
}

// HasResUri returns a boolean if a field has been set.
func (o *ServiceParameterData) HasResUri() bool {
	if o != nil && !isNil(o.ResUri) {
		return true
	}

	return false
}

// SetResUri gets a reference to the given string and assigns it to the ResUri field.
func (o *ServiceParameterData) SetResUri(v string) {
	o.ResUri = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ServiceParameterData) GetHeaders() []string {
	if o == nil || isNil(o.Headers) {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetHeadersOk() ([]string, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ServiceParameterData) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []string and assigns it to the Headers field.
func (o *ServiceParameterData) SetHeaders(v []string) {
	o.Headers = v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *ServiceParameterData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *ServiceParameterData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *ServiceParameterData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o ServiceParameterData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.InterGroupId) {
		toSerialize["interGroupId"] = o.InterGroupId
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.UeIpv4) {
		toSerialize["ueIpv4"] = o.UeIpv4
	}
	if !isNil(o.UeIpv6) {
		toSerialize["ueIpv6"] = o.UeIpv6
	}
	if !isNil(o.UeMac) {
		toSerialize["ueMac"] = o.UeMac
	}
	if !isNil(o.AnyUeInd) {
		toSerialize["anyUeInd"] = o.AnyUeInd
	}
	if !isNil(o.ParamOverPc5) {
		toSerialize["paramOverPc5"] = o.ParamOverPc5
	}
	if !isNil(o.ParamOverUu) {
		toSerialize["paramOverUu"] = o.ParamOverUu
	}
	if !isNil(o.ParamForProSeDd) {
		toSerialize["paramForProSeDd"] = o.ParamForProSeDd
	}
	if !isNil(o.ParamForProSeDc) {
		toSerialize["paramForProSeDc"] = o.ParamForProSeDc
	}
	if !isNil(o.ParamForProSeU2NRelUe) {
		toSerialize["paramForProSeU2NRelUe"] = o.ParamForProSeU2NRelUe
	}
	if !isNil(o.ParamForProSeRemUe) {
		toSerialize["paramForProSeRemUe"] = o.ParamForProSeRemUe
	}
	if !isNil(o.UrspGuidance) {
		toSerialize["urspGuidance"] = o.UrspGuidance
	}
	if !isNil(o.DeliveryEvents) {
		toSerialize["deliveryEvents"] = o.DeliveryEvents
	}
	if !isNil(o.PolicDelivNotifCorreId) {
		toSerialize["policDelivNotifCorreId"] = o.PolicDelivNotifCorreId
	}
	if !isNil(o.PolicDelivNotifUri) {
		toSerialize["policDelivNotifUri"] = o.PolicDelivNotifUri
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.ResUri) {
		toSerialize["resUri"] = o.ResUri
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableServiceParameterData struct {
	value *ServiceParameterData
	isSet bool
}

func (v NullableServiceParameterData) Get() *ServiceParameterData {
	return v.value
}

func (v *NullableServiceParameterData) Set(val *ServiceParameterData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceParameterData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceParameterData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceParameterData(val *ServiceParameterData) *NullableServiceParameterData {
	return &NullableServiceParameterData{value: val, isSet: true}
}

func (v NullableServiceParameterData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceParameterData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


