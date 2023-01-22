# PfdManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PfdDatas** | [**map[string]PfdData**](PfdData.md) | Each element uniquely identifies the PFDs for an external application identifier. Each element is identified in the map via an external application identifier as key. The response shall include successfully provisioned PFD data of application(s). | 
**PfdReports** | Pointer to [**map[string]PfdReport**](PfdReport.md) | Supplied by the SCEF and contains the external application identifiers for which PFD(s) are not added or modified successfully. The failure reason is also included. Each element provides the related information for one or more external application identifier(s) and is identified in the map via the failure identifier as key. | [optional] [readonly] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 

## Methods

### NewPfdManagement

`func NewPfdManagement(pfdDatas map[string]PfdData, ) *PfdManagement`

NewPfdManagement instantiates a new PfdManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfdManagementWithDefaults

`func NewPfdManagementWithDefaults() *PfdManagement`

NewPfdManagementWithDefaults instantiates a new PfdManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *PfdManagement) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PfdManagement) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PfdManagement) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PfdManagement) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PfdManagement) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PfdManagement) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PfdManagement) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PfdManagement) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetPfdDatas

`func (o *PfdManagement) GetPfdDatas() map[string]PfdData`

GetPfdDatas returns the PfdDatas field if non-nil, zero value otherwise.

### GetPfdDatasOk

`func (o *PfdManagement) GetPfdDatasOk() (*map[string]PfdData, bool)`

GetPfdDatasOk returns a tuple with the PfdDatas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdDatas

`func (o *PfdManagement) SetPfdDatas(v map[string]PfdData)`

SetPfdDatas sets PfdDatas field to given value.


### GetPfdReports

`func (o *PfdManagement) GetPfdReports() map[string]PfdReport`

GetPfdReports returns the PfdReports field if non-nil, zero value otherwise.

### GetPfdReportsOk

`func (o *PfdManagement) GetPfdReportsOk() (*map[string]PfdReport, bool)`

GetPfdReportsOk returns a tuple with the PfdReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdReports

`func (o *PfdManagement) SetPfdReports(v map[string]PfdReport)`

SetPfdReports sets PfdReports field to given value.

### HasPfdReports

`func (o *PfdManagement) HasPfdReports() bool`

HasPfdReports returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *PfdManagement) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *PfdManagement) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *PfdManagement) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *PfdManagement) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *PfdManagement) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *PfdManagement) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *PfdManagement) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *PfdManagement) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *PfdManagement) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *PfdManagement) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *PfdManagement) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *PfdManagement) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


