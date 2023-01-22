# AfNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ReportEvent** | Pointer to [**Event**](Event.md) |  | [optional] 
**AuthResult** | Pointer to [**AuthorizationResult**](AuthorizationResult.md) |  | [optional] 
**Gpsis** | Pointer to **[]string** |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**EventInfo** | Pointer to [**EventInfo**](EventInfo.md) |  | [optional] 

## Methods

### NewAfNotification

`func NewAfNotification(subscription string, ) *AfNotification`

NewAfNotification instantiates a new AfNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfNotificationWithDefaults

`func NewAfNotificationWithDefaults() *AfNotification`

NewAfNotificationWithDefaults instantiates a new AfNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *AfNotification) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AfNotification) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AfNotification) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetReportEvent

`func (o *AfNotification) GetReportEvent() Event`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *AfNotification) GetReportEventOk() (*Event, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *AfNotification) SetReportEvent(v Event)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *AfNotification) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.

### GetAuthResult

`func (o *AfNotification) GetAuthResult() AuthorizationResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *AfNotification) GetAuthResultOk() (*AuthorizationResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *AfNotification) SetAuthResult(v AuthorizationResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *AfNotification) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetGpsis

`func (o *AfNotification) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *AfNotification) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *AfNotification) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *AfNotification) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetDnn

`func (o *AfNotification) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *AfNotification) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *AfNotification) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *AfNotification) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *AfNotification) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AfNotification) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AfNotification) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *AfNotification) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetEventInfo

`func (o *AfNotification) GetEventInfo() EventInfo`

GetEventInfo returns the EventInfo field if non-nil, zero value otherwise.

### GetEventInfoOk

`func (o *AfNotification) GetEventInfoOk() (*EventInfo, bool)`

GetEventInfoOk returns a tuple with the EventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventInfo

`func (o *AfNotification) SetEventInfo(v EventInfo)`

SetEventInfo sets EventInfo field to given value.

### HasEventInfo

`func (o *AfNotification) HasEventInfo() bool`

HasEventInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


