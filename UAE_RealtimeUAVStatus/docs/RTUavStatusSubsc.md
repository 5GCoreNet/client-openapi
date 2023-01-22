# RTUavStatusSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UassId** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**UavIds** | [**[]UavId**](UavId.md) |  | 
**NotificationUri** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewRTUavStatusSubsc

`func NewRTUavStatusSubsc(uassId string, uavIds []UavId, notificationUri string, ) *RTUavStatusSubsc`

NewRTUavStatusSubsc instantiates a new RTUavStatusSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRTUavStatusSubscWithDefaults

`func NewRTUavStatusSubscWithDefaults() *RTUavStatusSubsc`

NewRTUavStatusSubscWithDefaults instantiates a new RTUavStatusSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUassId

`func (o *RTUavStatusSubsc) GetUassId() string`

GetUassId returns the UassId field if non-nil, zero value otherwise.

### GetUassIdOk

`func (o *RTUavStatusSubsc) GetUassIdOk() (*string, bool)`

GetUassIdOk returns a tuple with the UassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUassId

`func (o *RTUavStatusSubsc) SetUassId(v string)`

SetUassId sets UassId field to given value.


### GetUavIds

`func (o *RTUavStatusSubsc) GetUavIds() []UavId`

GetUavIds returns the UavIds field if non-nil, zero value otherwise.

### GetUavIdsOk

`func (o *RTUavStatusSubsc) GetUavIdsOk() (*[]UavId, bool)`

GetUavIdsOk returns a tuple with the UavIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavIds

`func (o *RTUavStatusSubsc) SetUavIds(v []UavId)`

SetUavIds sets UavIds field to given value.


### GetNotificationUri

`func (o *RTUavStatusSubsc) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *RTUavStatusSubsc) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *RTUavStatusSubsc) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetSuppFeat

`func (o *RTUavStatusSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *RTUavStatusSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *RTUavStatusSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *RTUavStatusSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


