# NgranFailureEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NgranId** | [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | 
**NgranFailureIndication** | [**NgranFailureIndication**](NgranFailureIndication.md) |  | 

## Methods

### NewNgranFailureEvent

`func NewNgranFailureEvent(ngranId GlobalRanNodeId, ngranFailureIndication NgranFailureIndication, ) *NgranFailureEvent`

NewNgranFailureEvent instantiates a new NgranFailureEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgranFailureEventWithDefaults

`func NewNgranFailureEventWithDefaults() *NgranFailureEvent`

NewNgranFailureEventWithDefaults instantiates a new NgranFailureEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNgranId

`func (o *NgranFailureEvent) GetNgranId() GlobalRanNodeId`

GetNgranId returns the NgranId field if non-nil, zero value otherwise.

### GetNgranIdOk

`func (o *NgranFailureEvent) GetNgranIdOk() (*GlobalRanNodeId, bool)`

GetNgranIdOk returns a tuple with the NgranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgranId

`func (o *NgranFailureEvent) SetNgranId(v GlobalRanNodeId)`

SetNgranId sets NgranId field to given value.


### GetNgranFailureIndication

`func (o *NgranFailureEvent) GetNgranFailureIndication() NgranFailureIndication`

GetNgranFailureIndication returns the NgranFailureIndication field if non-nil, zero value otherwise.

### GetNgranFailureIndicationOk

`func (o *NgranFailureEvent) GetNgranFailureIndicationOk() (*NgranFailureIndication, bool)`

GetNgranFailureIndicationOk returns a tuple with the NgranFailureIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgranFailureIndication

`func (o *NgranFailureEvent) SetNgranFailureIndication(v NgranFailureIndication)`

SetNgranFailureIndication sets NgranFailureIndication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


