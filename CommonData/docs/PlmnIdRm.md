# PlmnIdRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mcc** | **string** | Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.   | 
**Mnc** | **string** | Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413. | 

## Methods

### NewPlmnIdRm

`func NewPlmnIdRm(mcc string, mnc string, ) *PlmnIdRm`

NewPlmnIdRm instantiates a new PlmnIdRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlmnIdRmWithDefaults

`func NewPlmnIdRmWithDefaults() *PlmnIdRm`

NewPlmnIdRmWithDefaults instantiates a new PlmnIdRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcc

`func (o *PlmnIdRm) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *PlmnIdRm) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *PlmnIdRm) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetMnc

`func (o *PlmnIdRm) GetMnc() string`

GetMnc returns the Mnc field if non-nil, zero value otherwise.

### GetMncOk

`func (o *PlmnIdRm) GetMncOk() (*string, bool)`

GetMncOk returns a tuple with the Mnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnc

`func (o *PlmnIdRm) SetMnc(v string)`

SetMnc sets Mnc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


