# MbsSessionCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**MbsSession** | [**MbsSession**](MbsSession.md) |  | 
**EventList** | Pointer to [**MbsSessionEventReportList**](MbsSessionEventReportList.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsSessionCreateReq

`func NewMbsSessionCreateReq(afId string, mbsSession MbsSession, ) *MbsSessionCreateReq`

NewMbsSessionCreateReq instantiates a new MbsSessionCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionCreateReqWithDefaults

`func NewMbsSessionCreateReqWithDefaults() *MbsSessionCreateReq`

NewMbsSessionCreateReqWithDefaults instantiates a new MbsSessionCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *MbsSessionCreateReq) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *MbsSessionCreateReq) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *MbsSessionCreateReq) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetMbsSession

`func (o *MbsSessionCreateReq) GetMbsSession() MbsSession`

GetMbsSession returns the MbsSession field if non-nil, zero value otherwise.

### GetMbsSessionOk

`func (o *MbsSessionCreateReq) GetMbsSessionOk() (*MbsSession, bool)`

GetMbsSessionOk returns a tuple with the MbsSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSession

`func (o *MbsSessionCreateReq) SetMbsSession(v MbsSession)`

SetMbsSession sets MbsSession field to given value.


### GetEventList

`func (o *MbsSessionCreateReq) GetEventList() MbsSessionEventReportList`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *MbsSessionCreateReq) GetEventListOk() (*MbsSessionEventReportList, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *MbsSessionCreateReq) SetEventList(v MbsSessionEventReportList)`

SetEventList sets EventList field to given value.

### HasEventList

`func (o *MbsSessionCreateReq) HasEventList() bool`

HasEventList returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MbsSessionCreateReq) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MbsSessionCreateReq) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MbsSessionCreateReq) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MbsSessionCreateReq) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


