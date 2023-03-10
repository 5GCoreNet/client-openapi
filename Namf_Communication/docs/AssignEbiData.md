# AssignEbiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**ArpList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 
**ReleasedEbiList** | Pointer to **[]int32** |  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ModifiedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 

## Methods

### NewAssignEbiData

`func NewAssignEbiData(pduSessionId int32, ) *AssignEbiData`

NewAssignEbiData instantiates a new AssignEbiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignEbiDataWithDefaults

`func NewAssignEbiDataWithDefaults() *AssignEbiData`

NewAssignEbiDataWithDefaults instantiates a new AssignEbiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *AssignEbiData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *AssignEbiData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *AssignEbiData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetArpList

`func (o *AssignEbiData) GetArpList() []Arp`

GetArpList returns the ArpList field if non-nil, zero value otherwise.

### GetArpListOk

`func (o *AssignEbiData) GetArpListOk() (*[]Arp, bool)`

GetArpListOk returns a tuple with the ArpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpList

`func (o *AssignEbiData) SetArpList(v []Arp)`

SetArpList sets ArpList field to given value.

### HasArpList

`func (o *AssignEbiData) HasArpList() bool`

HasArpList returns a boolean if a field has been set.

### GetReleasedEbiList

`func (o *AssignEbiData) GetReleasedEbiList() []int32`

GetReleasedEbiList returns the ReleasedEbiList field if non-nil, zero value otherwise.

### GetReleasedEbiListOk

`func (o *AssignEbiData) GetReleasedEbiListOk() (*[]int32, bool)`

GetReleasedEbiListOk returns a tuple with the ReleasedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedEbiList

`func (o *AssignEbiData) SetReleasedEbiList(v []int32)`

SetReleasedEbiList sets ReleasedEbiList field to given value.

### HasReleasedEbiList

`func (o *AssignEbiData) HasReleasedEbiList() bool`

HasReleasedEbiList returns a boolean if a field has been set.

### GetOldGuami

`func (o *AssignEbiData) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *AssignEbiData) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *AssignEbiData) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *AssignEbiData) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.

### GetModifiedEbiList

`func (o *AssignEbiData) GetModifiedEbiList() []EbiArpMapping`

GetModifiedEbiList returns the ModifiedEbiList field if non-nil, zero value otherwise.

### GetModifiedEbiListOk

`func (o *AssignEbiData) GetModifiedEbiListOk() (*[]EbiArpMapping, bool)`

GetModifiedEbiListOk returns a tuple with the ModifiedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedEbiList

`func (o *AssignEbiData) SetModifiedEbiList(v []EbiArpMapping)`

SetModifiedEbiList sets ModifiedEbiList field to given value.

### HasModifiedEbiList

`func (o *AssignEbiData) HasModifiedEbiList() bool`

HasModifiedEbiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


