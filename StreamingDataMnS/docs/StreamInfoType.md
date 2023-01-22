# StreamInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamType** | [**StreamTypeType**](StreamTypeType.md) |  | 
**SerializationFormat** | [**SerializationFormatType**](SerializationFormatType.md) |  | 
**StreamId** | [**StreamInfoTypeStreamId**](StreamInfoTypeStreamId.md) |  | 
**AdditionalInfo** | Pointer to [**StreamInfoTypeAdditionalInfo**](StreamInfoTypeAdditionalInfo.md) |  | [optional] 

## Methods

### NewStreamInfoType

`func NewStreamInfoType(streamType StreamTypeType, serializationFormat SerializationFormatType, streamId StreamInfoTypeStreamId, ) *StreamInfoType`

NewStreamInfoType instantiates a new StreamInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamInfoTypeWithDefaults

`func NewStreamInfoTypeWithDefaults() *StreamInfoType`

NewStreamInfoTypeWithDefaults instantiates a new StreamInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamType

`func (o *StreamInfoType) GetStreamType() StreamTypeType`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *StreamInfoType) GetStreamTypeOk() (*StreamTypeType, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *StreamInfoType) SetStreamType(v StreamTypeType)`

SetStreamType sets StreamType field to given value.


### GetSerializationFormat

`func (o *StreamInfoType) GetSerializationFormat() SerializationFormatType`

GetSerializationFormat returns the SerializationFormat field if non-nil, zero value otherwise.

### GetSerializationFormatOk

`func (o *StreamInfoType) GetSerializationFormatOk() (*SerializationFormatType, bool)`

GetSerializationFormatOk returns a tuple with the SerializationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializationFormat

`func (o *StreamInfoType) SetSerializationFormat(v SerializationFormatType)`

SetSerializationFormat sets SerializationFormat field to given value.


### GetStreamId

`func (o *StreamInfoType) GetStreamId() StreamInfoTypeStreamId`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamInfoType) GetStreamIdOk() (*StreamInfoTypeStreamId, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamInfoType) SetStreamId(v StreamInfoTypeStreamId)`

SetStreamId sets StreamId field to given value.


### GetAdditionalInfo

`func (o *StreamInfoType) GetAdditionalInfo() StreamInfoTypeAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *StreamInfoType) GetAdditionalInfoOk() (*StreamInfoTypeAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *StreamInfoType) SetAdditionalInfo(v StreamInfoTypeAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *StreamInfoType) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


