# ApplicationSpecificRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**RecordType** | **string** | String providing an URI formatted according to RFC 3986. | 
**RecordContainer** | **interface{}** |  | 

## Methods

### NewApplicationSpecificRecord

`func NewApplicationSpecificRecord(timestamp time.Time, recordType string, recordContainer interface{}, ) *ApplicationSpecificRecord`

NewApplicationSpecificRecord instantiates a new ApplicationSpecificRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSpecificRecordWithDefaults

`func NewApplicationSpecificRecordWithDefaults() *ApplicationSpecificRecord`

NewApplicationSpecificRecordWithDefaults instantiates a new ApplicationSpecificRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApplicationSpecificRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApplicationSpecificRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApplicationSpecificRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetRecordType

`func (o *ApplicationSpecificRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ApplicationSpecificRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ApplicationSpecificRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetRecordContainer

`func (o *ApplicationSpecificRecord) GetRecordContainer() interface{}`

GetRecordContainer returns the RecordContainer field if non-nil, zero value otherwise.

### GetRecordContainerOk

`func (o *ApplicationSpecificRecord) GetRecordContainerOk() (*interface{}, bool)`

GetRecordContainerOk returns a tuple with the RecordContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordContainer

`func (o *ApplicationSpecificRecord) SetRecordContainer(v interface{})`

SetRecordContainer sets RecordContainer field to given value.


### SetRecordContainerNil

`func (o *ApplicationSpecificRecord) SetRecordContainerNil(b bool)`

 SetRecordContainerNil sets the value for RecordContainer to be an explicit nil

### UnsetRecordContainer
`func (o *ApplicationSpecificRecord) UnsetRecordContainer()`

UnsetRecordContainer ensures that no value is present for RecordContainer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


