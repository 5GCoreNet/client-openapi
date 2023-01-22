# SliceInfoForPDUSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssai** | [**Snssai**](Snssai.md) |  | 
**RoamingIndication** | [**RoamingIndication**](RoamingIndication.md) |  | 
**HomeSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewSliceInfoForPDUSession

`func NewSliceInfoForPDUSession(sNssai Snssai, roamingIndication RoamingIndication, ) *SliceInfoForPDUSession`

NewSliceInfoForPDUSession instantiates a new SliceInfoForPDUSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceInfoForPDUSessionWithDefaults

`func NewSliceInfoForPDUSessionWithDefaults() *SliceInfoForPDUSession`

NewSliceInfoForPDUSessionWithDefaults instantiates a new SliceInfoForPDUSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssai

`func (o *SliceInfoForPDUSession) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SliceInfoForPDUSession) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SliceInfoForPDUSession) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetRoamingIndication

`func (o *SliceInfoForPDUSession) GetRoamingIndication() RoamingIndication`

GetRoamingIndication returns the RoamingIndication field if non-nil, zero value otherwise.

### GetRoamingIndicationOk

`func (o *SliceInfoForPDUSession) GetRoamingIndicationOk() (*RoamingIndication, bool)`

GetRoamingIndicationOk returns a tuple with the RoamingIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingIndication

`func (o *SliceInfoForPDUSession) SetRoamingIndication(v RoamingIndication)`

SetRoamingIndication sets RoamingIndication field to given value.


### GetHomeSnssai

`func (o *SliceInfoForPDUSession) GetHomeSnssai() Snssai`

GetHomeSnssai returns the HomeSnssai field if non-nil, zero value otherwise.

### GetHomeSnssaiOk

`func (o *SliceInfoForPDUSession) GetHomeSnssaiOk() (*Snssai, bool)`

GetHomeSnssaiOk returns a tuple with the HomeSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeSnssai

`func (o *SliceInfoForPDUSession) SetHomeSnssai(v Snssai)`

SetHomeSnssai sets HomeSnssai field to given value.

### HasHomeSnssai

`func (o *SliceInfoForPDUSession) HasHomeSnssai() bool`

HasHomeSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


