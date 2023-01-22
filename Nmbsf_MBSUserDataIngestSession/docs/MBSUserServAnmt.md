# MBSUserServAnmt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtServiceId** | **[]string** |  | 
**ServClass** | **string** |  | 
**StartTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EndTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**ServNameDescs** | [**[]ServiceNameDescription**](ServiceNameDescription.md) |  | 
**MainServLang** | Pointer to **string** |  | [optional] 
**MbsDistSessAnmt** | Pointer to [**map[string]MBSDistSessionAnmt**](MBSDistSessionAnmt.md) | Represents the set of MBS Distribution Session Announcements currently associated with  this MBS User Service Announcement.  | [optional] 

## Methods

### NewMBSUserServAnmt

`func NewMBSUserServAnmt(extServiceId []string, servClass string, servNameDescs []ServiceNameDescription, ) *MBSUserServAnmt`

NewMBSUserServAnmt instantiates a new MBSUserServAnmt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSUserServAnmtWithDefaults

`func NewMBSUserServAnmtWithDefaults() *MBSUserServAnmt`

NewMBSUserServAnmtWithDefaults instantiates a new MBSUserServAnmt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtServiceId

`func (o *MBSUserServAnmt) GetExtServiceId() []string`

GetExtServiceId returns the ExtServiceId field if non-nil, zero value otherwise.

### GetExtServiceIdOk

`func (o *MBSUserServAnmt) GetExtServiceIdOk() (*[]string, bool)`

GetExtServiceIdOk returns a tuple with the ExtServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtServiceId

`func (o *MBSUserServAnmt) SetExtServiceId(v []string)`

SetExtServiceId sets ExtServiceId field to given value.


### GetServClass

`func (o *MBSUserServAnmt) GetServClass() string`

GetServClass returns the ServClass field if non-nil, zero value otherwise.

### GetServClassOk

`func (o *MBSUserServAnmt) GetServClassOk() (*string, bool)`

GetServClassOk returns a tuple with the ServClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServClass

`func (o *MBSUserServAnmt) SetServClass(v string)`

SetServClass sets ServClass field to given value.


### GetStartTime

`func (o *MBSUserServAnmt) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MBSUserServAnmt) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MBSUserServAnmt) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MBSUserServAnmt) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *MBSUserServAnmt) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MBSUserServAnmt) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MBSUserServAnmt) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MBSUserServAnmt) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetServNameDescs

`func (o *MBSUserServAnmt) GetServNameDescs() []ServiceNameDescription`

GetServNameDescs returns the ServNameDescs field if non-nil, zero value otherwise.

### GetServNameDescsOk

`func (o *MBSUserServAnmt) GetServNameDescsOk() (*[]ServiceNameDescription, bool)`

GetServNameDescsOk returns a tuple with the ServNameDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServNameDescs

`func (o *MBSUserServAnmt) SetServNameDescs(v []ServiceNameDescription)`

SetServNameDescs sets ServNameDescs field to given value.


### GetMainServLang

`func (o *MBSUserServAnmt) GetMainServLang() string`

GetMainServLang returns the MainServLang field if non-nil, zero value otherwise.

### GetMainServLangOk

`func (o *MBSUserServAnmt) GetMainServLangOk() (*string, bool)`

GetMainServLangOk returns a tuple with the MainServLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainServLang

`func (o *MBSUserServAnmt) SetMainServLang(v string)`

SetMainServLang sets MainServLang field to given value.

### HasMainServLang

`func (o *MBSUserServAnmt) HasMainServLang() bool`

HasMainServLang returns a boolean if a field has been set.

### GetMbsDistSessAnmt

`func (o *MBSUserServAnmt) GetMbsDistSessAnmt() map[string]MBSDistSessionAnmt`

GetMbsDistSessAnmt returns the MbsDistSessAnmt field if non-nil, zero value otherwise.

### GetMbsDistSessAnmtOk

`func (o *MBSUserServAnmt) GetMbsDistSessAnmtOk() (*map[string]MBSDistSessionAnmt, bool)`

GetMbsDistSessAnmtOk returns a tuple with the MbsDistSessAnmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDistSessAnmt

`func (o *MBSUserServAnmt) SetMbsDistSessAnmt(v map[string]MBSDistSessionAnmt)`

SetMbsDistSessAnmt sets MbsDistSessAnmt field to given value.

### HasMbsDistSessAnmt

`func (o *MBSUserServAnmt) HasMbsDistSessAnmt() bool`

HasMbsDistSessAnmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


