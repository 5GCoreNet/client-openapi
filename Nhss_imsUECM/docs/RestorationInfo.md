# RestorationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Contact** | **string** |  | 
**InitialCSeqSequenceNumber** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**CallIdSipHeader** | Pointer to **string** |  | [optional] 
**UesubscriptionInfo** | Pointer to [**UeSubscriptionInfo**](UeSubscriptionInfo.md) |  | [optional] 
**PcscfSubscriptionInfo** | Pointer to [**PcscfSubscriptionInfo**](PcscfSubscriptionInfo.md) |  | [optional] 
**ImsSdmSubscriptions** | Pointer to [**map[string]ImsSdmSubscription**](ImsSdmSubscription.md) | A map (list of key-value pairs where subscriptionId serves as key) of ImsSdmSubscription  | [optional] 

## Methods

### NewRestorationInfo

`func NewRestorationInfo(path string, contact string, ) *RestorationInfo`

NewRestorationInfo instantiates a new RestorationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestorationInfoWithDefaults

`func NewRestorationInfoWithDefaults() *RestorationInfo`

NewRestorationInfoWithDefaults instantiates a new RestorationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RestorationInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RestorationInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RestorationInfo) SetPath(v string)`

SetPath sets Path field to given value.


### GetContact

`func (o *RestorationInfo) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *RestorationInfo) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *RestorationInfo) SetContact(v string)`

SetContact sets Contact field to given value.


### GetInitialCSeqSequenceNumber

`func (o *RestorationInfo) GetInitialCSeqSequenceNumber() int32`

GetInitialCSeqSequenceNumber returns the InitialCSeqSequenceNumber field if non-nil, zero value otherwise.

### GetInitialCSeqSequenceNumberOk

`func (o *RestorationInfo) GetInitialCSeqSequenceNumberOk() (*int32, bool)`

GetInitialCSeqSequenceNumberOk returns a tuple with the InitialCSeqSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCSeqSequenceNumber

`func (o *RestorationInfo) SetInitialCSeqSequenceNumber(v int32)`

SetInitialCSeqSequenceNumber sets InitialCSeqSequenceNumber field to given value.

### HasInitialCSeqSequenceNumber

`func (o *RestorationInfo) HasInitialCSeqSequenceNumber() bool`

HasInitialCSeqSequenceNumber returns a boolean if a field has been set.

### GetCallIdSipHeader

`func (o *RestorationInfo) GetCallIdSipHeader() string`

GetCallIdSipHeader returns the CallIdSipHeader field if non-nil, zero value otherwise.

### GetCallIdSipHeaderOk

`func (o *RestorationInfo) GetCallIdSipHeaderOk() (*string, bool)`

GetCallIdSipHeaderOk returns a tuple with the CallIdSipHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIdSipHeader

`func (o *RestorationInfo) SetCallIdSipHeader(v string)`

SetCallIdSipHeader sets CallIdSipHeader field to given value.

### HasCallIdSipHeader

`func (o *RestorationInfo) HasCallIdSipHeader() bool`

HasCallIdSipHeader returns a boolean if a field has been set.

### GetUesubscriptionInfo

`func (o *RestorationInfo) GetUesubscriptionInfo() UeSubscriptionInfo`

GetUesubscriptionInfo returns the UesubscriptionInfo field if non-nil, zero value otherwise.

### GetUesubscriptionInfoOk

`func (o *RestorationInfo) GetUesubscriptionInfoOk() (*UeSubscriptionInfo, bool)`

GetUesubscriptionInfoOk returns a tuple with the UesubscriptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUesubscriptionInfo

`func (o *RestorationInfo) SetUesubscriptionInfo(v UeSubscriptionInfo)`

SetUesubscriptionInfo sets UesubscriptionInfo field to given value.

### HasUesubscriptionInfo

`func (o *RestorationInfo) HasUesubscriptionInfo() bool`

HasUesubscriptionInfo returns a boolean if a field has been set.

### GetPcscfSubscriptionInfo

`func (o *RestorationInfo) GetPcscfSubscriptionInfo() PcscfSubscriptionInfo`

GetPcscfSubscriptionInfo returns the PcscfSubscriptionInfo field if non-nil, zero value otherwise.

### GetPcscfSubscriptionInfoOk

`func (o *RestorationInfo) GetPcscfSubscriptionInfoOk() (*PcscfSubscriptionInfo, bool)`

GetPcscfSubscriptionInfoOk returns a tuple with the PcscfSubscriptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfSubscriptionInfo

`func (o *RestorationInfo) SetPcscfSubscriptionInfo(v PcscfSubscriptionInfo)`

SetPcscfSubscriptionInfo sets PcscfSubscriptionInfo field to given value.

### HasPcscfSubscriptionInfo

`func (o *RestorationInfo) HasPcscfSubscriptionInfo() bool`

HasPcscfSubscriptionInfo returns a boolean if a field has been set.

### GetImsSdmSubscriptions

`func (o *RestorationInfo) GetImsSdmSubscriptions() map[string]ImsSdmSubscription`

GetImsSdmSubscriptions returns the ImsSdmSubscriptions field if non-nil, zero value otherwise.

### GetImsSdmSubscriptionsOk

`func (o *RestorationInfo) GetImsSdmSubscriptionsOk() (*map[string]ImsSdmSubscription, bool)`

GetImsSdmSubscriptionsOk returns a tuple with the ImsSdmSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsSdmSubscriptions

`func (o *RestorationInfo) SetImsSdmSubscriptions(v map[string]ImsSdmSubscription)`

SetImsSdmSubscriptions sets ImsSdmSubscriptions field to given value.

### HasImsSdmSubscriptions

`func (o *RestorationInfo) HasImsSdmSubscriptions() bool`

HasImsSdmSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


