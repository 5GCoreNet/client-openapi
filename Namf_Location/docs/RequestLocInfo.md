# RequestLocInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Req5gsLoc** | Pointer to **bool** |  | [optional] [default to false]
**ReqCurrentLoc** | Pointer to **bool** |  | [optional] [default to false]
**ReqRatType** | Pointer to **bool** |  | [optional] [default to false]
**ReqTimeZone** | Pointer to **bool** |  | [optional] [default to false]
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewRequestLocInfo

`func NewRequestLocInfo() *RequestLocInfo`

NewRequestLocInfo instantiates a new RequestLocInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestLocInfoWithDefaults

`func NewRequestLocInfoWithDefaults() *RequestLocInfo`

NewRequestLocInfoWithDefaults instantiates a new RequestLocInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReq5gsLoc

`func (o *RequestLocInfo) GetReq5gsLoc() bool`

GetReq5gsLoc returns the Req5gsLoc field if non-nil, zero value otherwise.

### GetReq5gsLocOk

`func (o *RequestLocInfo) GetReq5gsLocOk() (*bool, bool)`

GetReq5gsLocOk returns a tuple with the Req5gsLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReq5gsLoc

`func (o *RequestLocInfo) SetReq5gsLoc(v bool)`

SetReq5gsLoc sets Req5gsLoc field to given value.

### HasReq5gsLoc

`func (o *RequestLocInfo) HasReq5gsLoc() bool`

HasReq5gsLoc returns a boolean if a field has been set.

### GetReqCurrentLoc

`func (o *RequestLocInfo) GetReqCurrentLoc() bool`

GetReqCurrentLoc returns the ReqCurrentLoc field if non-nil, zero value otherwise.

### GetReqCurrentLocOk

`func (o *RequestLocInfo) GetReqCurrentLocOk() (*bool, bool)`

GetReqCurrentLocOk returns a tuple with the ReqCurrentLoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCurrentLoc

`func (o *RequestLocInfo) SetReqCurrentLoc(v bool)`

SetReqCurrentLoc sets ReqCurrentLoc field to given value.

### HasReqCurrentLoc

`func (o *RequestLocInfo) HasReqCurrentLoc() bool`

HasReqCurrentLoc returns a boolean if a field has been set.

### GetReqRatType

`func (o *RequestLocInfo) GetReqRatType() bool`

GetReqRatType returns the ReqRatType field if non-nil, zero value otherwise.

### GetReqRatTypeOk

`func (o *RequestLocInfo) GetReqRatTypeOk() (*bool, bool)`

GetReqRatTypeOk returns a tuple with the ReqRatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqRatType

`func (o *RequestLocInfo) SetReqRatType(v bool)`

SetReqRatType sets ReqRatType field to given value.

### HasReqRatType

`func (o *RequestLocInfo) HasReqRatType() bool`

HasReqRatType returns a boolean if a field has been set.

### GetReqTimeZone

`func (o *RequestLocInfo) GetReqTimeZone() bool`

GetReqTimeZone returns the ReqTimeZone field if non-nil, zero value otherwise.

### GetReqTimeZoneOk

`func (o *RequestLocInfo) GetReqTimeZoneOk() (*bool, bool)`

GetReqTimeZoneOk returns a tuple with the ReqTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTimeZone

`func (o *RequestLocInfo) SetReqTimeZone(v bool)`

SetReqTimeZone sets ReqTimeZone field to given value.

### HasReqTimeZone

`func (o *RequestLocInfo) HasReqTimeZone() bool`

HasReqTimeZone returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *RequestLocInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RequestLocInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RequestLocInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RequestLocInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


