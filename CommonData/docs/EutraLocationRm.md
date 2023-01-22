# EutraLocationRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | [**Tai**](Tai.md) |  | 
**IgnoreTai** | Pointer to **bool** |  | [optional] [default to false]
**Ecgi** | [**Ecgi**](Ecgi.md) |  | 
**IgnoreEcgi** | Pointer to **bool** | This flag when present shall indicate that the Ecgi shall be ignored When present, it shall be set as follows: - true: ecgi shall be ignored. - false (default): ecgi shall not be ignored.  | [optional] [default to false]
**AgeOfLocationInformation** | Pointer to **int32** | The value represents the elapsed time in minutes since the last network contact of the mobile station.  Value \&quot;0\&quot; indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful NG-RAN location reporting procedure with the eNB when the UE is in connected mode.  Any other value than \&quot;0\&quot; indicates that the location information is the last known one.  See 3GPP TS 29.002 clause 17.7.8.  | [optional] 
**UeLocationTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**GeographicalInformation** | Pointer to **string** | Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.  | [optional] 
**GeodeticInformation** | Pointer to **string** | Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used.  | [optional] 
**GlobalNgenbId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**GlobalENbId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 

## Methods

### NewEutraLocationRm

`func NewEutraLocationRm(tai Tai, ecgi Ecgi, ) *EutraLocationRm`

NewEutraLocationRm instantiates a new EutraLocationRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEutraLocationRmWithDefaults

`func NewEutraLocationRmWithDefaults() *EutraLocationRm`

NewEutraLocationRmWithDefaults instantiates a new EutraLocationRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *EutraLocationRm) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *EutraLocationRm) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *EutraLocationRm) SetTai(v Tai)`

SetTai sets Tai field to given value.


### GetIgnoreTai

`func (o *EutraLocationRm) GetIgnoreTai() bool`

GetIgnoreTai returns the IgnoreTai field if non-nil, zero value otherwise.

### GetIgnoreTaiOk

`func (o *EutraLocationRm) GetIgnoreTaiOk() (*bool, bool)`

GetIgnoreTaiOk returns a tuple with the IgnoreTai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTai

`func (o *EutraLocationRm) SetIgnoreTai(v bool)`

SetIgnoreTai sets IgnoreTai field to given value.

### HasIgnoreTai

`func (o *EutraLocationRm) HasIgnoreTai() bool`

HasIgnoreTai returns a boolean if a field has been set.

### GetEcgi

`func (o *EutraLocationRm) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *EutraLocationRm) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *EutraLocationRm) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.


### GetIgnoreEcgi

`func (o *EutraLocationRm) GetIgnoreEcgi() bool`

GetIgnoreEcgi returns the IgnoreEcgi field if non-nil, zero value otherwise.

### GetIgnoreEcgiOk

`func (o *EutraLocationRm) GetIgnoreEcgiOk() (*bool, bool)`

GetIgnoreEcgiOk returns a tuple with the IgnoreEcgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreEcgi

`func (o *EutraLocationRm) SetIgnoreEcgi(v bool)`

SetIgnoreEcgi sets IgnoreEcgi field to given value.

### HasIgnoreEcgi

`func (o *EutraLocationRm) HasIgnoreEcgi() bool`

HasIgnoreEcgi returns a boolean if a field has been set.

### GetAgeOfLocationInformation

`func (o *EutraLocationRm) GetAgeOfLocationInformation() int32`

GetAgeOfLocationInformation returns the AgeOfLocationInformation field if non-nil, zero value otherwise.

### GetAgeOfLocationInformationOk

`func (o *EutraLocationRm) GetAgeOfLocationInformationOk() (*int32, bool)`

GetAgeOfLocationInformationOk returns a tuple with the AgeOfLocationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInformation

`func (o *EutraLocationRm) SetAgeOfLocationInformation(v int32)`

SetAgeOfLocationInformation sets AgeOfLocationInformation field to given value.

### HasAgeOfLocationInformation

`func (o *EutraLocationRm) HasAgeOfLocationInformation() bool`

HasAgeOfLocationInformation returns a boolean if a field has been set.

### GetUeLocationTimestamp

`func (o *EutraLocationRm) GetUeLocationTimestamp() time.Time`

GetUeLocationTimestamp returns the UeLocationTimestamp field if non-nil, zero value otherwise.

### GetUeLocationTimestampOk

`func (o *EutraLocationRm) GetUeLocationTimestampOk() (*time.Time, bool)`

GetUeLocationTimestampOk returns a tuple with the UeLocationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationTimestamp

`func (o *EutraLocationRm) SetUeLocationTimestamp(v time.Time)`

SetUeLocationTimestamp sets UeLocationTimestamp field to given value.

### HasUeLocationTimestamp

`func (o *EutraLocationRm) HasUeLocationTimestamp() bool`

HasUeLocationTimestamp returns a boolean if a field has been set.

### GetGeographicalInformation

`func (o *EutraLocationRm) GetGeographicalInformation() string`

GetGeographicalInformation returns the GeographicalInformation field if non-nil, zero value otherwise.

### GetGeographicalInformationOk

`func (o *EutraLocationRm) GetGeographicalInformationOk() (*string, bool)`

GetGeographicalInformationOk returns a tuple with the GeographicalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicalInformation

`func (o *EutraLocationRm) SetGeographicalInformation(v string)`

SetGeographicalInformation sets GeographicalInformation field to given value.

### HasGeographicalInformation

`func (o *EutraLocationRm) HasGeographicalInformation() bool`

HasGeographicalInformation returns a boolean if a field has been set.

### GetGeodeticInformation

`func (o *EutraLocationRm) GetGeodeticInformation() string`

GetGeodeticInformation returns the GeodeticInformation field if non-nil, zero value otherwise.

### GetGeodeticInformationOk

`func (o *EutraLocationRm) GetGeodeticInformationOk() (*string, bool)`

GetGeodeticInformationOk returns a tuple with the GeodeticInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeodeticInformation

`func (o *EutraLocationRm) SetGeodeticInformation(v string)`

SetGeodeticInformation sets GeodeticInformation field to given value.

### HasGeodeticInformation

`func (o *EutraLocationRm) HasGeodeticInformation() bool`

HasGeodeticInformation returns a boolean if a field has been set.

### GetGlobalNgenbId

`func (o *EutraLocationRm) GetGlobalNgenbId() GlobalRanNodeId`

GetGlobalNgenbId returns the GlobalNgenbId field if non-nil, zero value otherwise.

### GetGlobalNgenbIdOk

`func (o *EutraLocationRm) GetGlobalNgenbIdOk() (*GlobalRanNodeId, bool)`

GetGlobalNgenbIdOk returns a tuple with the GlobalNgenbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalNgenbId

`func (o *EutraLocationRm) SetGlobalNgenbId(v GlobalRanNodeId)`

SetGlobalNgenbId sets GlobalNgenbId field to given value.

### HasGlobalNgenbId

`func (o *EutraLocationRm) HasGlobalNgenbId() bool`

HasGlobalNgenbId returns a boolean if a field has been set.

### GetGlobalENbId

`func (o *EutraLocationRm) GetGlobalENbId() GlobalRanNodeId`

GetGlobalENbId returns the GlobalENbId field if non-nil, zero value otherwise.

### GetGlobalENbIdOk

`func (o *EutraLocationRm) GetGlobalENbIdOk() (*GlobalRanNodeId, bool)`

GetGlobalENbIdOk returns a tuple with the GlobalENbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalENbId

`func (o *EutraLocationRm) SetGlobalENbId(v GlobalRanNodeId)`

SetGlobalENbId sets GlobalENbId field to given value.

### HasGlobalENbId

`func (o *EutraLocationRm) HasGlobalENbId() bool`

HasGlobalENbId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


