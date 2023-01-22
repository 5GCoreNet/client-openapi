# ContentProtocols

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownlinkIngestProtocols** | Pointer to [**[]ContentProtocolDescriptor**](ContentProtocolDescriptor.md) |  | [optional] 
**UplinkEgestProtocols** | Pointer to [**[]ContentProtocolDescriptor**](ContentProtocolDescriptor.md) |  | [optional] 
**GeoFencingLocatorTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewContentProtocols

`func NewContentProtocols() *ContentProtocols`

NewContentProtocols instantiates a new ContentProtocols object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentProtocolsWithDefaults

`func NewContentProtocolsWithDefaults() *ContentProtocols`

NewContentProtocolsWithDefaults instantiates a new ContentProtocols object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownlinkIngestProtocols

`func (o *ContentProtocols) GetDownlinkIngestProtocols() []ContentProtocolDescriptor`

GetDownlinkIngestProtocols returns the DownlinkIngestProtocols field if non-nil, zero value otherwise.

### GetDownlinkIngestProtocolsOk

`func (o *ContentProtocols) GetDownlinkIngestProtocolsOk() (*[]ContentProtocolDescriptor, bool)`

GetDownlinkIngestProtocolsOk returns a tuple with the DownlinkIngestProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkIngestProtocols

`func (o *ContentProtocols) SetDownlinkIngestProtocols(v []ContentProtocolDescriptor)`

SetDownlinkIngestProtocols sets DownlinkIngestProtocols field to given value.

### HasDownlinkIngestProtocols

`func (o *ContentProtocols) HasDownlinkIngestProtocols() bool`

HasDownlinkIngestProtocols returns a boolean if a field has been set.

### GetUplinkEgestProtocols

`func (o *ContentProtocols) GetUplinkEgestProtocols() []ContentProtocolDescriptor`

GetUplinkEgestProtocols returns the UplinkEgestProtocols field if non-nil, zero value otherwise.

### GetUplinkEgestProtocolsOk

`func (o *ContentProtocols) GetUplinkEgestProtocolsOk() (*[]ContentProtocolDescriptor, bool)`

GetUplinkEgestProtocolsOk returns a tuple with the UplinkEgestProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkEgestProtocols

`func (o *ContentProtocols) SetUplinkEgestProtocols(v []ContentProtocolDescriptor)`

SetUplinkEgestProtocols sets UplinkEgestProtocols field to given value.

### HasUplinkEgestProtocols

`func (o *ContentProtocols) HasUplinkEgestProtocols() bool`

HasUplinkEgestProtocols returns a boolean if a field has been set.

### GetGeoFencingLocatorTypes

`func (o *ContentProtocols) GetGeoFencingLocatorTypes() []string`

GetGeoFencingLocatorTypes returns the GeoFencingLocatorTypes field if non-nil, zero value otherwise.

### GetGeoFencingLocatorTypesOk

`func (o *ContentProtocols) GetGeoFencingLocatorTypesOk() (*[]string, bool)`

GetGeoFencingLocatorTypesOk returns a tuple with the GeoFencingLocatorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoFencingLocatorTypes

`func (o *ContentProtocols) SetGeoFencingLocatorTypes(v []string)`

SetGeoFencingLocatorTypes sets GeoFencingLocatorTypes field to given value.

### HasGeoFencingLocatorTypes

`func (o *ContentProtocols) HasGeoFencingLocatorTypes() bool`

HasGeoFencingLocatorTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


