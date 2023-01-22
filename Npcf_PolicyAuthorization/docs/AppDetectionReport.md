# AppDetectionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdNotifType** | [**AppDetectionNotifType**](AppDetectionNotifType.md) |  | 
**AfAppId** | **string** | Contains an AF application identifier. | 

## Methods

### NewAppDetectionReport

`func NewAppDetectionReport(adNotifType AppDetectionNotifType, afAppId string, ) *AppDetectionReport`

NewAppDetectionReport instantiates a new AppDetectionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDetectionReportWithDefaults

`func NewAppDetectionReportWithDefaults() *AppDetectionReport`

NewAppDetectionReportWithDefaults instantiates a new AppDetectionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdNotifType

`func (o *AppDetectionReport) GetAdNotifType() AppDetectionNotifType`

GetAdNotifType returns the AdNotifType field if non-nil, zero value otherwise.

### GetAdNotifTypeOk

`func (o *AppDetectionReport) GetAdNotifTypeOk() (*AppDetectionNotifType, bool)`

GetAdNotifTypeOk returns a tuple with the AdNotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdNotifType

`func (o *AppDetectionReport) SetAdNotifType(v AppDetectionNotifType)`

SetAdNotifType sets AdNotifType field to given value.


### GetAfAppId

`func (o *AppDetectionReport) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *AppDetectionReport) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *AppDetectionReport) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


