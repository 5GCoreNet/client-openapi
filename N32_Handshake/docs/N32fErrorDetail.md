# N32fErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** |  | 
**MsgReconstructFailReason** | [**FailureReason**](FailureReason.md) |  | 

## Methods

### NewN32fErrorDetail

`func NewN32fErrorDetail(attribute string, msgReconstructFailReason FailureReason, ) *N32fErrorDetail`

NewN32fErrorDetail instantiates a new N32fErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN32fErrorDetailWithDefaults

`func NewN32fErrorDetailWithDefaults() *N32fErrorDetail`

NewN32fErrorDetailWithDefaults instantiates a new N32fErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *N32fErrorDetail) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *N32fErrorDetail) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *N32fErrorDetail) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetMsgReconstructFailReason

`func (o *N32fErrorDetail) GetMsgReconstructFailReason() FailureReason`

GetMsgReconstructFailReason returns the MsgReconstructFailReason field if non-nil, zero value otherwise.

### GetMsgReconstructFailReasonOk

`func (o *N32fErrorDetail) GetMsgReconstructFailReasonOk() (*FailureReason, bool)`

GetMsgReconstructFailReasonOk returns a tuple with the MsgReconstructFailReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgReconstructFailReason

`func (o *N32fErrorDetail) SetMsgReconstructFailReason(v FailureReason)`

SetMsgReconstructFailReason sets MsgReconstructFailReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


