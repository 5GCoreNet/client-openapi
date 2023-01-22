# UserPlaneEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | **string** | String identifying the individual data session information for which the QoS event notification is delivered.  | 
**EventReports** | [**[]UserPlaneEventReport**](UserPlaneEventReport.md) | Contains the flow description for the Uplink and/or Downlink IP flows.  | 

## Methods

### NewUserPlaneEventNotification

`func NewUserPlaneEventNotification(sessionId string, eventReports []UserPlaneEventReport, ) *UserPlaneEventNotification`

NewUserPlaneEventNotification instantiates a new UserPlaneEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlaneEventNotificationWithDefaults

`func NewUserPlaneEventNotificationWithDefaults() *UserPlaneEventNotification`

NewUserPlaneEventNotificationWithDefaults instantiates a new UserPlaneEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *UserPlaneEventNotification) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *UserPlaneEventNotification) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *UserPlaneEventNotification) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetEventReports

`func (o *UserPlaneEventNotification) GetEventReports() []UserPlaneEventReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *UserPlaneEventNotification) GetEventReportsOk() (*[]UserPlaneEventReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *UserPlaneEventNotification) SetEventReports(v []UserPlaneEventReport)`

SetEventReports sets EventReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


