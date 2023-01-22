# UserPlaneNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**EventReports** | [**[]UserPlaneEventReport**](UserPlaneEventReport.md) | Contains the reported event and applicable information | 

## Methods

### NewUserPlaneNotificationData

`func NewUserPlaneNotificationData(transaction string, eventReports []UserPlaneEventReport, ) *UserPlaneNotificationData`

NewUserPlaneNotificationData instantiates a new UserPlaneNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlaneNotificationDataWithDefaults

`func NewUserPlaneNotificationDataWithDefaults() *UserPlaneNotificationData`

NewUserPlaneNotificationDataWithDefaults instantiates a new UserPlaneNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *UserPlaneNotificationData) GetTransaction() string`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *UserPlaneNotificationData) GetTransactionOk() (*string, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *UserPlaneNotificationData) SetTransaction(v string)`

SetTransaction sets Transaction field to given value.


### GetEventReports

`func (o *UserPlaneNotificationData) GetEventReports() []UserPlaneEventReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *UserPlaneNotificationData) GetEventReportsOk() (*[]UserPlaneEventReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *UserPlaneNotificationData) SetEventReports(v []UserPlaneEventReport)`

SetEventReports sets EventReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


