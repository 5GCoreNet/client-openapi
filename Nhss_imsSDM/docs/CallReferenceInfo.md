# CallReferenceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRefNumber** | **string** | The content is according to CallReferenceNumber type described in 3GPP TS 29.002 [30]. Base64 encoded according to IETF RFC 2045 [28]  | 
**AsNumber** | **string** | The content is according to ISDN-AddressString type described in 3GPP TS 29.002 [30]. Base64 encoded according to IETF RFC 2045 [28]  | 

## Methods

### NewCallReferenceInfo

`func NewCallReferenceInfo(callRefNumber string, asNumber string, ) *CallReferenceInfo`

NewCallReferenceInfo instantiates a new CallReferenceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallReferenceInfoWithDefaults

`func NewCallReferenceInfoWithDefaults() *CallReferenceInfo`

NewCallReferenceInfoWithDefaults instantiates a new CallReferenceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallRefNumber

`func (o *CallReferenceInfo) GetCallRefNumber() string`

GetCallRefNumber returns the CallRefNumber field if non-nil, zero value otherwise.

### GetCallRefNumberOk

`func (o *CallReferenceInfo) GetCallRefNumberOk() (*string, bool)`

GetCallRefNumberOk returns a tuple with the CallRefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRefNumber

`func (o *CallReferenceInfo) SetCallRefNumber(v string)`

SetCallRefNumber sets CallRefNumber field to given value.


### GetAsNumber

`func (o *CallReferenceInfo) GetAsNumber() string`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *CallReferenceInfo) GetAsNumberOk() (*string, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *CallReferenceInfo) SetAsNumber(v string)`

SetAsNumber sets AsNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


