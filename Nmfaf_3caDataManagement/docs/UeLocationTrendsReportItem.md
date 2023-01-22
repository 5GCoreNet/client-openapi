# UeLocationTrendsReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**N3gaLocation** | Pointer to [**N3gaLocation**](N3gaLocation.md) |  | [optional] 
**Spacing** | **int32** | indicating a time in seconds. | 
**Duration** | **int32** | indicating a time in seconds. | 
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewUeLocationTrendsReportItem

`func NewUeLocationTrendsReportItem(spacing int32, duration int32, timestamp time.Time, ) *UeLocationTrendsReportItem`

NewUeLocationTrendsReportItem instantiates a new UeLocationTrendsReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeLocationTrendsReportItemWithDefaults

`func NewUeLocationTrendsReportItemWithDefaults() *UeLocationTrendsReportItem`

NewUeLocationTrendsReportItemWithDefaults instantiates a new UeLocationTrendsReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *UeLocationTrendsReportItem) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *UeLocationTrendsReportItem) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *UeLocationTrendsReportItem) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *UeLocationTrendsReportItem) HasTai() bool`

HasTai returns a boolean if a field has been set.

### GetNcgi

`func (o *UeLocationTrendsReportItem) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *UeLocationTrendsReportItem) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *UeLocationTrendsReportItem) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *UeLocationTrendsReportItem) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetEcgi

`func (o *UeLocationTrendsReportItem) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *UeLocationTrendsReportItem) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *UeLocationTrendsReportItem) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *UeLocationTrendsReportItem) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetN3gaLocation

`func (o *UeLocationTrendsReportItem) GetN3gaLocation() N3gaLocation`

GetN3gaLocation returns the N3gaLocation field if non-nil, zero value otherwise.

### GetN3gaLocationOk

`func (o *UeLocationTrendsReportItem) GetN3gaLocationOk() (*N3gaLocation, bool)`

GetN3gaLocationOk returns a tuple with the N3gaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gaLocation

`func (o *UeLocationTrendsReportItem) SetN3gaLocation(v N3gaLocation)`

SetN3gaLocation sets N3gaLocation field to given value.

### HasN3gaLocation

`func (o *UeLocationTrendsReportItem) HasN3gaLocation() bool`

HasN3gaLocation returns a boolean if a field has been set.

### GetSpacing

`func (o *UeLocationTrendsReportItem) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *UeLocationTrendsReportItem) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *UeLocationTrendsReportItem) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.


### GetDuration

`func (o *UeLocationTrendsReportItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UeLocationTrendsReportItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UeLocationTrendsReportItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimestamp

`func (o *UeLocationTrendsReportItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UeLocationTrendsReportItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UeLocationTrendsReportItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


