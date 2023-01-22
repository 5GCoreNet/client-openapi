# WebsockNotifConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebsocketUri** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestWebsocketUri** | Pointer to **bool** | Set by the SCS/AS to indicate that the Websocket delivery is requested. | [optional] 

## Methods

### NewWebsockNotifConfig

`func NewWebsockNotifConfig() *WebsockNotifConfig`

NewWebsockNotifConfig instantiates a new WebsockNotifConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsockNotifConfigWithDefaults

`func NewWebsockNotifConfigWithDefaults() *WebsockNotifConfig`

NewWebsockNotifConfigWithDefaults instantiates a new WebsockNotifConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebsocketUri

`func (o *WebsockNotifConfig) GetWebsocketUri() string`

GetWebsocketUri returns the WebsocketUri field if non-nil, zero value otherwise.

### GetWebsocketUriOk

`func (o *WebsockNotifConfig) GetWebsocketUriOk() (*string, bool)`

GetWebsocketUriOk returns a tuple with the WebsocketUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketUri

`func (o *WebsockNotifConfig) SetWebsocketUri(v string)`

SetWebsocketUri sets WebsocketUri field to given value.

### HasWebsocketUri

`func (o *WebsockNotifConfig) HasWebsocketUri() bool`

HasWebsocketUri returns a boolean if a field has been set.

### GetRequestWebsocketUri

`func (o *WebsockNotifConfig) GetRequestWebsocketUri() bool`

GetRequestWebsocketUri returns the RequestWebsocketUri field if non-nil, zero value otherwise.

### GetRequestWebsocketUriOk

`func (o *WebsockNotifConfig) GetRequestWebsocketUriOk() (*bool, bool)`

GetRequestWebsocketUriOk returns a tuple with the RequestWebsocketUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestWebsocketUri

`func (o *WebsockNotifConfig) SetRequestWebsocketUri(v bool)`

SetRequestWebsocketUri sets RequestWebsocketUri field to given value.

### HasRequestWebsocketUri

`func (o *WebsockNotifConfig) HasRequestWebsocketUri() bool`

HasRequestWebsocketUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


