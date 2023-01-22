/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Eees_ACREvents

import (
	"encoding/json"
)

// CivicAddress Indicates a Civic address.
type CivicAddress struct {
	Country *string `json:"country,omitempty"`
	A1 *string `json:"A1,omitempty"`
	A2 *string `json:"A2,omitempty"`
	A3 *string `json:"A3,omitempty"`
	A4 *string `json:"A4,omitempty"`
	A5 *string `json:"A5,omitempty"`
	A6 *string `json:"A6,omitempty"`
	PRD *string `json:"PRD,omitempty"`
	POD *string `json:"POD,omitempty"`
	STS *string `json:"STS,omitempty"`
	HNO *string `json:"HNO,omitempty"`
	HNS *string `json:"HNS,omitempty"`
	LMK *string `json:"LMK,omitempty"`
	LOC *string `json:"LOC,omitempty"`
	NAM *string `json:"NAM,omitempty"`
	PC *string `json:"PC,omitempty"`
	BLD *string `json:"BLD,omitempty"`
	UNIT *string `json:"UNIT,omitempty"`
	FLR *string `json:"FLR,omitempty"`
	ROOM *string `json:"ROOM,omitempty"`
	PLC *string `json:"PLC,omitempty"`
	PCN *string `json:"PCN,omitempty"`
	POBOX *string `json:"POBOX,omitempty"`
	ADDCODE *string `json:"ADDCODE,omitempty"`
	SEAT *string `json:"SEAT,omitempty"`
	RD *string `json:"RD,omitempty"`
	RDSEC *string `json:"RDSEC,omitempty"`
	RDBR *string `json:"RDBR,omitempty"`
	RDSUBBR *string `json:"RDSUBBR,omitempty"`
	PRM *string `json:"PRM,omitempty"`
	POM *string `json:"POM,omitempty"`
	UsageRules *string `json:"usageRules,omitempty"`
	Method *string `json:"method,omitempty"`
	ProvidedBy *string `json:"providedBy,omitempty"`
}

// NewCivicAddress instantiates a new CivicAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCivicAddress() *CivicAddress {
	this := CivicAddress{}
	return &this
}

// NewCivicAddressWithDefaults instantiates a new CivicAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCivicAddressWithDefaults() *CivicAddress {
	this := CivicAddress{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CivicAddress) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CivicAddress) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CivicAddress) SetCountry(v string) {
	o.Country = &v
}

// GetA1 returns the A1 field value if set, zero value otherwise.
func (o *CivicAddress) GetA1() string {
	if o == nil || isNil(o.A1) {
		var ret string
		return ret
	}
	return *o.A1
}

// GetA1Ok returns a tuple with the A1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA1Ok() (*string, bool) {
	if o == nil || isNil(o.A1) {
    return nil, false
	}
	return o.A1, true
}

// HasA1 returns a boolean if a field has been set.
func (o *CivicAddress) HasA1() bool {
	if o != nil && !isNil(o.A1) {
		return true
	}

	return false
}

// SetA1 gets a reference to the given string and assigns it to the A1 field.
func (o *CivicAddress) SetA1(v string) {
	o.A1 = &v
}

// GetA2 returns the A2 field value if set, zero value otherwise.
func (o *CivicAddress) GetA2() string {
	if o == nil || isNil(o.A2) {
		var ret string
		return ret
	}
	return *o.A2
}

// GetA2Ok returns a tuple with the A2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA2Ok() (*string, bool) {
	if o == nil || isNil(o.A2) {
    return nil, false
	}
	return o.A2, true
}

// HasA2 returns a boolean if a field has been set.
func (o *CivicAddress) HasA2() bool {
	if o != nil && !isNil(o.A2) {
		return true
	}

	return false
}

// SetA2 gets a reference to the given string and assigns it to the A2 field.
func (o *CivicAddress) SetA2(v string) {
	o.A2 = &v
}

// GetA3 returns the A3 field value if set, zero value otherwise.
func (o *CivicAddress) GetA3() string {
	if o == nil || isNil(o.A3) {
		var ret string
		return ret
	}
	return *o.A3
}

// GetA3Ok returns a tuple with the A3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA3Ok() (*string, bool) {
	if o == nil || isNil(o.A3) {
    return nil, false
	}
	return o.A3, true
}

// HasA3 returns a boolean if a field has been set.
func (o *CivicAddress) HasA3() bool {
	if o != nil && !isNil(o.A3) {
		return true
	}

	return false
}

// SetA3 gets a reference to the given string and assigns it to the A3 field.
func (o *CivicAddress) SetA3(v string) {
	o.A3 = &v
}

// GetA4 returns the A4 field value if set, zero value otherwise.
func (o *CivicAddress) GetA4() string {
	if o == nil || isNil(o.A4) {
		var ret string
		return ret
	}
	return *o.A4
}

// GetA4Ok returns a tuple with the A4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA4Ok() (*string, bool) {
	if o == nil || isNil(o.A4) {
    return nil, false
	}
	return o.A4, true
}

// HasA4 returns a boolean if a field has been set.
func (o *CivicAddress) HasA4() bool {
	if o != nil && !isNil(o.A4) {
		return true
	}

	return false
}

// SetA4 gets a reference to the given string and assigns it to the A4 field.
func (o *CivicAddress) SetA4(v string) {
	o.A4 = &v
}

// GetA5 returns the A5 field value if set, zero value otherwise.
func (o *CivicAddress) GetA5() string {
	if o == nil || isNil(o.A5) {
		var ret string
		return ret
	}
	return *o.A5
}

// GetA5Ok returns a tuple with the A5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA5Ok() (*string, bool) {
	if o == nil || isNil(o.A5) {
    return nil, false
	}
	return o.A5, true
}

// HasA5 returns a boolean if a field has been set.
func (o *CivicAddress) HasA5() bool {
	if o != nil && !isNil(o.A5) {
		return true
	}

	return false
}

// SetA5 gets a reference to the given string and assigns it to the A5 field.
func (o *CivicAddress) SetA5(v string) {
	o.A5 = &v
}

// GetA6 returns the A6 field value if set, zero value otherwise.
func (o *CivicAddress) GetA6() string {
	if o == nil || isNil(o.A6) {
		var ret string
		return ret
	}
	return *o.A6
}

// GetA6Ok returns a tuple with the A6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetA6Ok() (*string, bool) {
	if o == nil || isNil(o.A6) {
    return nil, false
	}
	return o.A6, true
}

// HasA6 returns a boolean if a field has been set.
func (o *CivicAddress) HasA6() bool {
	if o != nil && !isNil(o.A6) {
		return true
	}

	return false
}

// SetA6 gets a reference to the given string and assigns it to the A6 field.
func (o *CivicAddress) SetA6(v string) {
	o.A6 = &v
}

// GetPRD returns the PRD field value if set, zero value otherwise.
func (o *CivicAddress) GetPRD() string {
	if o == nil || isNil(o.PRD) {
		var ret string
		return ret
	}
	return *o.PRD
}

// GetPRDOk returns a tuple with the PRD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPRDOk() (*string, bool) {
	if o == nil || isNil(o.PRD) {
    return nil, false
	}
	return o.PRD, true
}

// HasPRD returns a boolean if a field has been set.
func (o *CivicAddress) HasPRD() bool {
	if o != nil && !isNil(o.PRD) {
		return true
	}

	return false
}

// SetPRD gets a reference to the given string and assigns it to the PRD field.
func (o *CivicAddress) SetPRD(v string) {
	o.PRD = &v
}

// GetPOD returns the POD field value if set, zero value otherwise.
func (o *CivicAddress) GetPOD() string {
	if o == nil || isNil(o.POD) {
		var ret string
		return ret
	}
	return *o.POD
}

// GetPODOk returns a tuple with the POD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPODOk() (*string, bool) {
	if o == nil || isNil(o.POD) {
    return nil, false
	}
	return o.POD, true
}

// HasPOD returns a boolean if a field has been set.
func (o *CivicAddress) HasPOD() bool {
	if o != nil && !isNil(o.POD) {
		return true
	}

	return false
}

// SetPOD gets a reference to the given string and assigns it to the POD field.
func (o *CivicAddress) SetPOD(v string) {
	o.POD = &v
}

// GetSTS returns the STS field value if set, zero value otherwise.
func (o *CivicAddress) GetSTS() string {
	if o == nil || isNil(o.STS) {
		var ret string
		return ret
	}
	return *o.STS
}

// GetSTSOk returns a tuple with the STS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetSTSOk() (*string, bool) {
	if o == nil || isNil(o.STS) {
    return nil, false
	}
	return o.STS, true
}

// HasSTS returns a boolean if a field has been set.
func (o *CivicAddress) HasSTS() bool {
	if o != nil && !isNil(o.STS) {
		return true
	}

	return false
}

// SetSTS gets a reference to the given string and assigns it to the STS field.
func (o *CivicAddress) SetSTS(v string) {
	o.STS = &v
}

// GetHNO returns the HNO field value if set, zero value otherwise.
func (o *CivicAddress) GetHNO() string {
	if o == nil || isNil(o.HNO) {
		var ret string
		return ret
	}
	return *o.HNO
}

// GetHNOOk returns a tuple with the HNO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetHNOOk() (*string, bool) {
	if o == nil || isNil(o.HNO) {
    return nil, false
	}
	return o.HNO, true
}

// HasHNO returns a boolean if a field has been set.
func (o *CivicAddress) HasHNO() bool {
	if o != nil && !isNil(o.HNO) {
		return true
	}

	return false
}

// SetHNO gets a reference to the given string and assigns it to the HNO field.
func (o *CivicAddress) SetHNO(v string) {
	o.HNO = &v
}

// GetHNS returns the HNS field value if set, zero value otherwise.
func (o *CivicAddress) GetHNS() string {
	if o == nil || isNil(o.HNS) {
		var ret string
		return ret
	}
	return *o.HNS
}

// GetHNSOk returns a tuple with the HNS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetHNSOk() (*string, bool) {
	if o == nil || isNil(o.HNS) {
    return nil, false
	}
	return o.HNS, true
}

// HasHNS returns a boolean if a field has been set.
func (o *CivicAddress) HasHNS() bool {
	if o != nil && !isNil(o.HNS) {
		return true
	}

	return false
}

// SetHNS gets a reference to the given string and assigns it to the HNS field.
func (o *CivicAddress) SetHNS(v string) {
	o.HNS = &v
}

// GetLMK returns the LMK field value if set, zero value otherwise.
func (o *CivicAddress) GetLMK() string {
	if o == nil || isNil(o.LMK) {
		var ret string
		return ret
	}
	return *o.LMK
}

// GetLMKOk returns a tuple with the LMK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetLMKOk() (*string, bool) {
	if o == nil || isNil(o.LMK) {
    return nil, false
	}
	return o.LMK, true
}

// HasLMK returns a boolean if a field has been set.
func (o *CivicAddress) HasLMK() bool {
	if o != nil && !isNil(o.LMK) {
		return true
	}

	return false
}

// SetLMK gets a reference to the given string and assigns it to the LMK field.
func (o *CivicAddress) SetLMK(v string) {
	o.LMK = &v
}

// GetLOC returns the LOC field value if set, zero value otherwise.
func (o *CivicAddress) GetLOC() string {
	if o == nil || isNil(o.LOC) {
		var ret string
		return ret
	}
	return *o.LOC
}

// GetLOCOk returns a tuple with the LOC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetLOCOk() (*string, bool) {
	if o == nil || isNil(o.LOC) {
    return nil, false
	}
	return o.LOC, true
}

// HasLOC returns a boolean if a field has been set.
func (o *CivicAddress) HasLOC() bool {
	if o != nil && !isNil(o.LOC) {
		return true
	}

	return false
}

// SetLOC gets a reference to the given string and assigns it to the LOC field.
func (o *CivicAddress) SetLOC(v string) {
	o.LOC = &v
}

// GetNAM returns the NAM field value if set, zero value otherwise.
func (o *CivicAddress) GetNAM() string {
	if o == nil || isNil(o.NAM) {
		var ret string
		return ret
	}
	return *o.NAM
}

// GetNAMOk returns a tuple with the NAM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetNAMOk() (*string, bool) {
	if o == nil || isNil(o.NAM) {
    return nil, false
	}
	return o.NAM, true
}

// HasNAM returns a boolean if a field has been set.
func (o *CivicAddress) HasNAM() bool {
	if o != nil && !isNil(o.NAM) {
		return true
	}

	return false
}

// SetNAM gets a reference to the given string and assigns it to the NAM field.
func (o *CivicAddress) SetNAM(v string) {
	o.NAM = &v
}

// GetPC returns the PC field value if set, zero value otherwise.
func (o *CivicAddress) GetPC() string {
	if o == nil || isNil(o.PC) {
		var ret string
		return ret
	}
	return *o.PC
}

// GetPCOk returns a tuple with the PC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPCOk() (*string, bool) {
	if o == nil || isNil(o.PC) {
    return nil, false
	}
	return o.PC, true
}

// HasPC returns a boolean if a field has been set.
func (o *CivicAddress) HasPC() bool {
	if o != nil && !isNil(o.PC) {
		return true
	}

	return false
}

// SetPC gets a reference to the given string and assigns it to the PC field.
func (o *CivicAddress) SetPC(v string) {
	o.PC = &v
}

// GetBLD returns the BLD field value if set, zero value otherwise.
func (o *CivicAddress) GetBLD() string {
	if o == nil || isNil(o.BLD) {
		var ret string
		return ret
	}
	return *o.BLD
}

// GetBLDOk returns a tuple with the BLD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetBLDOk() (*string, bool) {
	if o == nil || isNil(o.BLD) {
    return nil, false
	}
	return o.BLD, true
}

// HasBLD returns a boolean if a field has been set.
func (o *CivicAddress) HasBLD() bool {
	if o != nil && !isNil(o.BLD) {
		return true
	}

	return false
}

// SetBLD gets a reference to the given string and assigns it to the BLD field.
func (o *CivicAddress) SetBLD(v string) {
	o.BLD = &v
}

// GetUNIT returns the UNIT field value if set, zero value otherwise.
func (o *CivicAddress) GetUNIT() string {
	if o == nil || isNil(o.UNIT) {
		var ret string
		return ret
	}
	return *o.UNIT
}

// GetUNITOk returns a tuple with the UNIT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetUNITOk() (*string, bool) {
	if o == nil || isNil(o.UNIT) {
    return nil, false
	}
	return o.UNIT, true
}

// HasUNIT returns a boolean if a field has been set.
func (o *CivicAddress) HasUNIT() bool {
	if o != nil && !isNil(o.UNIT) {
		return true
	}

	return false
}

// SetUNIT gets a reference to the given string and assigns it to the UNIT field.
func (o *CivicAddress) SetUNIT(v string) {
	o.UNIT = &v
}

// GetFLR returns the FLR field value if set, zero value otherwise.
func (o *CivicAddress) GetFLR() string {
	if o == nil || isNil(o.FLR) {
		var ret string
		return ret
	}
	return *o.FLR
}

// GetFLROk returns a tuple with the FLR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetFLROk() (*string, bool) {
	if o == nil || isNil(o.FLR) {
    return nil, false
	}
	return o.FLR, true
}

// HasFLR returns a boolean if a field has been set.
func (o *CivicAddress) HasFLR() bool {
	if o != nil && !isNil(o.FLR) {
		return true
	}

	return false
}

// SetFLR gets a reference to the given string and assigns it to the FLR field.
func (o *CivicAddress) SetFLR(v string) {
	o.FLR = &v
}

// GetROOM returns the ROOM field value if set, zero value otherwise.
func (o *CivicAddress) GetROOM() string {
	if o == nil || isNil(o.ROOM) {
		var ret string
		return ret
	}
	return *o.ROOM
}

// GetROOMOk returns a tuple with the ROOM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetROOMOk() (*string, bool) {
	if o == nil || isNil(o.ROOM) {
    return nil, false
	}
	return o.ROOM, true
}

// HasROOM returns a boolean if a field has been set.
func (o *CivicAddress) HasROOM() bool {
	if o != nil && !isNil(o.ROOM) {
		return true
	}

	return false
}

// SetROOM gets a reference to the given string and assigns it to the ROOM field.
func (o *CivicAddress) SetROOM(v string) {
	o.ROOM = &v
}

// GetPLC returns the PLC field value if set, zero value otherwise.
func (o *CivicAddress) GetPLC() string {
	if o == nil || isNil(o.PLC) {
		var ret string
		return ret
	}
	return *o.PLC
}

// GetPLCOk returns a tuple with the PLC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPLCOk() (*string, bool) {
	if o == nil || isNil(o.PLC) {
    return nil, false
	}
	return o.PLC, true
}

// HasPLC returns a boolean if a field has been set.
func (o *CivicAddress) HasPLC() bool {
	if o != nil && !isNil(o.PLC) {
		return true
	}

	return false
}

// SetPLC gets a reference to the given string and assigns it to the PLC field.
func (o *CivicAddress) SetPLC(v string) {
	o.PLC = &v
}

// GetPCN returns the PCN field value if set, zero value otherwise.
func (o *CivicAddress) GetPCN() string {
	if o == nil || isNil(o.PCN) {
		var ret string
		return ret
	}
	return *o.PCN
}

// GetPCNOk returns a tuple with the PCN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPCNOk() (*string, bool) {
	if o == nil || isNil(o.PCN) {
    return nil, false
	}
	return o.PCN, true
}

// HasPCN returns a boolean if a field has been set.
func (o *CivicAddress) HasPCN() bool {
	if o != nil && !isNil(o.PCN) {
		return true
	}

	return false
}

// SetPCN gets a reference to the given string and assigns it to the PCN field.
func (o *CivicAddress) SetPCN(v string) {
	o.PCN = &v
}

// GetPOBOX returns the POBOX field value if set, zero value otherwise.
func (o *CivicAddress) GetPOBOX() string {
	if o == nil || isNil(o.POBOX) {
		var ret string
		return ret
	}
	return *o.POBOX
}

// GetPOBOXOk returns a tuple with the POBOX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPOBOXOk() (*string, bool) {
	if o == nil || isNil(o.POBOX) {
    return nil, false
	}
	return o.POBOX, true
}

// HasPOBOX returns a boolean if a field has been set.
func (o *CivicAddress) HasPOBOX() bool {
	if o != nil && !isNil(o.POBOX) {
		return true
	}

	return false
}

// SetPOBOX gets a reference to the given string and assigns it to the POBOX field.
func (o *CivicAddress) SetPOBOX(v string) {
	o.POBOX = &v
}

// GetADDCODE returns the ADDCODE field value if set, zero value otherwise.
func (o *CivicAddress) GetADDCODE() string {
	if o == nil || isNil(o.ADDCODE) {
		var ret string
		return ret
	}
	return *o.ADDCODE
}

// GetADDCODEOk returns a tuple with the ADDCODE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetADDCODEOk() (*string, bool) {
	if o == nil || isNil(o.ADDCODE) {
    return nil, false
	}
	return o.ADDCODE, true
}

// HasADDCODE returns a boolean if a field has been set.
func (o *CivicAddress) HasADDCODE() bool {
	if o != nil && !isNil(o.ADDCODE) {
		return true
	}

	return false
}

// SetADDCODE gets a reference to the given string and assigns it to the ADDCODE field.
func (o *CivicAddress) SetADDCODE(v string) {
	o.ADDCODE = &v
}

// GetSEAT returns the SEAT field value if set, zero value otherwise.
func (o *CivicAddress) GetSEAT() string {
	if o == nil || isNil(o.SEAT) {
		var ret string
		return ret
	}
	return *o.SEAT
}

// GetSEATOk returns a tuple with the SEAT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetSEATOk() (*string, bool) {
	if o == nil || isNil(o.SEAT) {
    return nil, false
	}
	return o.SEAT, true
}

// HasSEAT returns a boolean if a field has been set.
func (o *CivicAddress) HasSEAT() bool {
	if o != nil && !isNil(o.SEAT) {
		return true
	}

	return false
}

// SetSEAT gets a reference to the given string and assigns it to the SEAT field.
func (o *CivicAddress) SetSEAT(v string) {
	o.SEAT = &v
}

// GetRD returns the RD field value if set, zero value otherwise.
func (o *CivicAddress) GetRD() string {
	if o == nil || isNil(o.RD) {
		var ret string
		return ret
	}
	return *o.RD
}

// GetRDOk returns a tuple with the RD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetRDOk() (*string, bool) {
	if o == nil || isNil(o.RD) {
    return nil, false
	}
	return o.RD, true
}

// HasRD returns a boolean if a field has been set.
func (o *CivicAddress) HasRD() bool {
	if o != nil && !isNil(o.RD) {
		return true
	}

	return false
}

// SetRD gets a reference to the given string and assigns it to the RD field.
func (o *CivicAddress) SetRD(v string) {
	o.RD = &v
}

// GetRDSEC returns the RDSEC field value if set, zero value otherwise.
func (o *CivicAddress) GetRDSEC() string {
	if o == nil || isNil(o.RDSEC) {
		var ret string
		return ret
	}
	return *o.RDSEC
}

// GetRDSECOk returns a tuple with the RDSEC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetRDSECOk() (*string, bool) {
	if o == nil || isNil(o.RDSEC) {
    return nil, false
	}
	return o.RDSEC, true
}

// HasRDSEC returns a boolean if a field has been set.
func (o *CivicAddress) HasRDSEC() bool {
	if o != nil && !isNil(o.RDSEC) {
		return true
	}

	return false
}

// SetRDSEC gets a reference to the given string and assigns it to the RDSEC field.
func (o *CivicAddress) SetRDSEC(v string) {
	o.RDSEC = &v
}

// GetRDBR returns the RDBR field value if set, zero value otherwise.
func (o *CivicAddress) GetRDBR() string {
	if o == nil || isNil(o.RDBR) {
		var ret string
		return ret
	}
	return *o.RDBR
}

// GetRDBROk returns a tuple with the RDBR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetRDBROk() (*string, bool) {
	if o == nil || isNil(o.RDBR) {
    return nil, false
	}
	return o.RDBR, true
}

// HasRDBR returns a boolean if a field has been set.
func (o *CivicAddress) HasRDBR() bool {
	if o != nil && !isNil(o.RDBR) {
		return true
	}

	return false
}

// SetRDBR gets a reference to the given string and assigns it to the RDBR field.
func (o *CivicAddress) SetRDBR(v string) {
	o.RDBR = &v
}

// GetRDSUBBR returns the RDSUBBR field value if set, zero value otherwise.
func (o *CivicAddress) GetRDSUBBR() string {
	if o == nil || isNil(o.RDSUBBR) {
		var ret string
		return ret
	}
	return *o.RDSUBBR
}

// GetRDSUBBROk returns a tuple with the RDSUBBR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetRDSUBBROk() (*string, bool) {
	if o == nil || isNil(o.RDSUBBR) {
    return nil, false
	}
	return o.RDSUBBR, true
}

// HasRDSUBBR returns a boolean if a field has been set.
func (o *CivicAddress) HasRDSUBBR() bool {
	if o != nil && !isNil(o.RDSUBBR) {
		return true
	}

	return false
}

// SetRDSUBBR gets a reference to the given string and assigns it to the RDSUBBR field.
func (o *CivicAddress) SetRDSUBBR(v string) {
	o.RDSUBBR = &v
}

// GetPRM returns the PRM field value if set, zero value otherwise.
func (o *CivicAddress) GetPRM() string {
	if o == nil || isNil(o.PRM) {
		var ret string
		return ret
	}
	return *o.PRM
}

// GetPRMOk returns a tuple with the PRM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPRMOk() (*string, bool) {
	if o == nil || isNil(o.PRM) {
    return nil, false
	}
	return o.PRM, true
}

// HasPRM returns a boolean if a field has been set.
func (o *CivicAddress) HasPRM() bool {
	if o != nil && !isNil(o.PRM) {
		return true
	}

	return false
}

// SetPRM gets a reference to the given string and assigns it to the PRM field.
func (o *CivicAddress) SetPRM(v string) {
	o.PRM = &v
}

// GetPOM returns the POM field value if set, zero value otherwise.
func (o *CivicAddress) GetPOM() string {
	if o == nil || isNil(o.POM) {
		var ret string
		return ret
	}
	return *o.POM
}

// GetPOMOk returns a tuple with the POM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetPOMOk() (*string, bool) {
	if o == nil || isNil(o.POM) {
    return nil, false
	}
	return o.POM, true
}

// HasPOM returns a boolean if a field has been set.
func (o *CivicAddress) HasPOM() bool {
	if o != nil && !isNil(o.POM) {
		return true
	}

	return false
}

// SetPOM gets a reference to the given string and assigns it to the POM field.
func (o *CivicAddress) SetPOM(v string) {
	o.POM = &v
}

// GetUsageRules returns the UsageRules field value if set, zero value otherwise.
func (o *CivicAddress) GetUsageRules() string {
	if o == nil || isNil(o.UsageRules) {
		var ret string
		return ret
	}
	return *o.UsageRules
}

// GetUsageRulesOk returns a tuple with the UsageRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetUsageRulesOk() (*string, bool) {
	if o == nil || isNil(o.UsageRules) {
    return nil, false
	}
	return o.UsageRules, true
}

// HasUsageRules returns a boolean if a field has been set.
func (o *CivicAddress) HasUsageRules() bool {
	if o != nil && !isNil(o.UsageRules) {
		return true
	}

	return false
}

// SetUsageRules gets a reference to the given string and assigns it to the UsageRules field.
func (o *CivicAddress) SetUsageRules(v string) {
	o.UsageRules = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *CivicAddress) GetMethod() string {
	if o == nil || isNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetMethodOk() (*string, bool) {
	if o == nil || isNil(o.Method) {
    return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *CivicAddress) HasMethod() bool {
	if o != nil && !isNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *CivicAddress) SetMethod(v string) {
	o.Method = &v
}

// GetProvidedBy returns the ProvidedBy field value if set, zero value otherwise.
func (o *CivicAddress) GetProvidedBy() string {
	if o == nil || isNil(o.ProvidedBy) {
		var ret string
		return ret
	}
	return *o.ProvidedBy
}

// GetProvidedByOk returns a tuple with the ProvidedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CivicAddress) GetProvidedByOk() (*string, bool) {
	if o == nil || isNil(o.ProvidedBy) {
    return nil, false
	}
	return o.ProvidedBy, true
}

// HasProvidedBy returns a boolean if a field has been set.
func (o *CivicAddress) HasProvidedBy() bool {
	if o != nil && !isNil(o.ProvidedBy) {
		return true
	}

	return false
}

// SetProvidedBy gets a reference to the given string and assigns it to the ProvidedBy field.
func (o *CivicAddress) SetProvidedBy(v string) {
	o.ProvidedBy = &v
}

func (o CivicAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.A1) {
		toSerialize["A1"] = o.A1
	}
	if !isNil(o.A2) {
		toSerialize["A2"] = o.A2
	}
	if !isNil(o.A3) {
		toSerialize["A3"] = o.A3
	}
	if !isNil(o.A4) {
		toSerialize["A4"] = o.A4
	}
	if !isNil(o.A5) {
		toSerialize["A5"] = o.A5
	}
	if !isNil(o.A6) {
		toSerialize["A6"] = o.A6
	}
	if !isNil(o.PRD) {
		toSerialize["PRD"] = o.PRD
	}
	if !isNil(o.POD) {
		toSerialize["POD"] = o.POD
	}
	if !isNil(o.STS) {
		toSerialize["STS"] = o.STS
	}
	if !isNil(o.HNO) {
		toSerialize["HNO"] = o.HNO
	}
	if !isNil(o.HNS) {
		toSerialize["HNS"] = o.HNS
	}
	if !isNil(o.LMK) {
		toSerialize["LMK"] = o.LMK
	}
	if !isNil(o.LOC) {
		toSerialize["LOC"] = o.LOC
	}
	if !isNil(o.NAM) {
		toSerialize["NAM"] = o.NAM
	}
	if !isNil(o.PC) {
		toSerialize["PC"] = o.PC
	}
	if !isNil(o.BLD) {
		toSerialize["BLD"] = o.BLD
	}
	if !isNil(o.UNIT) {
		toSerialize["UNIT"] = o.UNIT
	}
	if !isNil(o.FLR) {
		toSerialize["FLR"] = o.FLR
	}
	if !isNil(o.ROOM) {
		toSerialize["ROOM"] = o.ROOM
	}
	if !isNil(o.PLC) {
		toSerialize["PLC"] = o.PLC
	}
	if !isNil(o.PCN) {
		toSerialize["PCN"] = o.PCN
	}
	if !isNil(o.POBOX) {
		toSerialize["POBOX"] = o.POBOX
	}
	if !isNil(o.ADDCODE) {
		toSerialize["ADDCODE"] = o.ADDCODE
	}
	if !isNil(o.SEAT) {
		toSerialize["SEAT"] = o.SEAT
	}
	if !isNil(o.RD) {
		toSerialize["RD"] = o.RD
	}
	if !isNil(o.RDSEC) {
		toSerialize["RDSEC"] = o.RDSEC
	}
	if !isNil(o.RDBR) {
		toSerialize["RDBR"] = o.RDBR
	}
	if !isNil(o.RDSUBBR) {
		toSerialize["RDSUBBR"] = o.RDSUBBR
	}
	if !isNil(o.PRM) {
		toSerialize["PRM"] = o.PRM
	}
	if !isNil(o.POM) {
		toSerialize["POM"] = o.POM
	}
	if !isNil(o.UsageRules) {
		toSerialize["usageRules"] = o.UsageRules
	}
	if !isNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !isNil(o.ProvidedBy) {
		toSerialize["providedBy"] = o.ProvidedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCivicAddress struct {
	value *CivicAddress
	isSet bool
}

func (v NullableCivicAddress) Get() *CivicAddress {
	return v.value
}

func (v *NullableCivicAddress) Set(val *CivicAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCivicAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCivicAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCivicAddress(val *CivicAddress) *NullableCivicAddress {
	return &NullableCivicAddress{value: val, isSet: true}
}

func (v NullableCivicAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCivicAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


