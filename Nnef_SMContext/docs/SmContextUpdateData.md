# SmContextUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DlNiddEndPoint** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmContextConfig** | Pointer to [**SmContextConfiguration**](SmContextConfiguration.md) |  | [optional] 

## Methods

### NewSmContextUpdateData

`func NewSmContextUpdateData() *SmContextUpdateData`

NewSmContextUpdateData instantiates a new SmContextUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextUpdateDataWithDefaults

`func NewSmContextUpdateDataWithDefaults() *SmContextUpdateData`

NewSmContextUpdateDataWithDefaults instantiates a new SmContextUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDlNiddEndPoint

`func (o *SmContextUpdateData) GetDlNiddEndPoint() string`

GetDlNiddEndPoint returns the DlNiddEndPoint field if non-nil, zero value otherwise.

### GetDlNiddEndPointOk

`func (o *SmContextUpdateData) GetDlNiddEndPointOk() (*string, bool)`

GetDlNiddEndPointOk returns a tuple with the DlNiddEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlNiddEndPoint

`func (o *SmContextUpdateData) SetDlNiddEndPoint(v string)`

SetDlNiddEndPoint sets DlNiddEndPoint field to given value.

### HasDlNiddEndPoint

`func (o *SmContextUpdateData) HasDlNiddEndPoint() bool`

HasDlNiddEndPoint returns a boolean if a field has been set.

### GetNotificationUri

`func (o *SmContextUpdateData) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *SmContextUpdateData) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *SmContextUpdateData) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.

### HasNotificationUri

`func (o *SmContextUpdateData) HasNotificationUri() bool`

HasNotificationUri returns a boolean if a field has been set.

### GetSmContextConfig

`func (o *SmContextUpdateData) GetSmContextConfig() SmContextConfiguration`

GetSmContextConfig returns the SmContextConfig field if non-nil, zero value otherwise.

### GetSmContextConfigOk

`func (o *SmContextUpdateData) GetSmContextConfigOk() (*SmContextConfiguration, bool)`

GetSmContextConfigOk returns a tuple with the SmContextConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextConfig

`func (o *SmContextUpdateData) SetSmContextConfig(v SmContextConfiguration)`

SetSmContextConfig sets SmContextConfig field to given value.

### HasSmContextConfig

`func (o *SmContextUpdateData) HasSmContextConfig() bool`

HasSmContextConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


