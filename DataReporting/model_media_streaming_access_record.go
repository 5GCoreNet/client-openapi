/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package DataReporting

import (
	"encoding/json"
)

// MediaStreamingAccessRecord struct for MediaStreamingAccessRecord
type MediaStreamingAccessRecord struct {
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
	MediaStreamHandlerEndpointAddress EndpointAddress `json:"mediaStreamHandlerEndpointAddress"`
	ApplicationServerEndpointAddress EndpointAddress `json:"applicationServerEndpointAddress"`
	SessionIdentifier *string `json:"sessionIdentifier,omitempty"`
	RequestMessage MediaStreamingAccessRecordAllOfRequestMessage `json:"requestMessage"`
	CacheStatus *CacheStatus `json:"cacheStatus,omitempty"`
	ResponseMessage MediaStreamingAccessRecordAllOfResponseMessage `json:"responseMessage"`
	// string with format 'float' as defined in OpenAPI.
	ProcessingLatency float32 `json:"processingLatency"`
	ConnectionMetrics *MediaStreamingAccessRecordAllOfConnectionMetrics `json:"connectionMetrics,omitempty"`
}

// NewMediaStreamingAccessRecord instantiates a new MediaStreamingAccessRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaStreamingAccessRecord(timestamp time.Time, mediaStreamHandlerEndpointAddress EndpointAddress, applicationServerEndpointAddress EndpointAddress, requestMessage MediaStreamingAccessRecordAllOfRequestMessage, responseMessage MediaStreamingAccessRecordAllOfResponseMessage, processingLatency float32) *MediaStreamingAccessRecord {
	this := MediaStreamingAccessRecord{}
	this.Timestamp = timestamp
	this.MediaStreamHandlerEndpointAddress = mediaStreamHandlerEndpointAddress
	this.ApplicationServerEndpointAddress = applicationServerEndpointAddress
	this.RequestMessage = requestMessage
	this.ResponseMessage = responseMessage
	this.ProcessingLatency = processingLatency
	return &this
}

// NewMediaStreamingAccessRecordWithDefaults instantiates a new MediaStreamingAccessRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaStreamingAccessRecordWithDefaults() *MediaStreamingAccessRecord {
	this := MediaStreamingAccessRecord{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *MediaStreamingAccessRecord) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MediaStreamingAccessRecord) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetMediaStreamHandlerEndpointAddress returns the MediaStreamHandlerEndpointAddress field value
func (o *MediaStreamingAccessRecord) GetMediaStreamHandlerEndpointAddress() EndpointAddress {
	if o == nil {
		var ret EndpointAddress
		return ret
	}

	return o.MediaStreamHandlerEndpointAddress
}

// GetMediaStreamHandlerEndpointAddressOk returns a tuple with the MediaStreamHandlerEndpointAddress field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetMediaStreamHandlerEndpointAddressOk() (*EndpointAddress, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MediaStreamHandlerEndpointAddress, true
}

// SetMediaStreamHandlerEndpointAddress sets field value
func (o *MediaStreamingAccessRecord) SetMediaStreamHandlerEndpointAddress(v EndpointAddress) {
	o.MediaStreamHandlerEndpointAddress = v
}

// GetApplicationServerEndpointAddress returns the ApplicationServerEndpointAddress field value
func (o *MediaStreamingAccessRecord) GetApplicationServerEndpointAddress() EndpointAddress {
	if o == nil {
		var ret EndpointAddress
		return ret
	}

	return o.ApplicationServerEndpointAddress
}

// GetApplicationServerEndpointAddressOk returns a tuple with the ApplicationServerEndpointAddress field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetApplicationServerEndpointAddressOk() (*EndpointAddress, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ApplicationServerEndpointAddress, true
}

// SetApplicationServerEndpointAddress sets field value
func (o *MediaStreamingAccessRecord) SetApplicationServerEndpointAddress(v EndpointAddress) {
	o.ApplicationServerEndpointAddress = v
}

// GetSessionIdentifier returns the SessionIdentifier field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecord) GetSessionIdentifier() string {
	if o == nil || isNil(o.SessionIdentifier) {
		var ret string
		return ret
	}
	return *o.SessionIdentifier
}

// GetSessionIdentifierOk returns a tuple with the SessionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetSessionIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.SessionIdentifier) {
    return nil, false
	}
	return o.SessionIdentifier, true
}

// HasSessionIdentifier returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecord) HasSessionIdentifier() bool {
	if o != nil && !isNil(o.SessionIdentifier) {
		return true
	}

	return false
}

// SetSessionIdentifier gets a reference to the given string and assigns it to the SessionIdentifier field.
func (o *MediaStreamingAccessRecord) SetSessionIdentifier(v string) {
	o.SessionIdentifier = &v
}

// GetRequestMessage returns the RequestMessage field value
func (o *MediaStreamingAccessRecord) GetRequestMessage() MediaStreamingAccessRecordAllOfRequestMessage {
	if o == nil {
		var ret MediaStreamingAccessRecordAllOfRequestMessage
		return ret
	}

	return o.RequestMessage
}

// GetRequestMessageOk returns a tuple with the RequestMessage field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetRequestMessageOk() (*MediaStreamingAccessRecordAllOfRequestMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RequestMessage, true
}

// SetRequestMessage sets field value
func (o *MediaStreamingAccessRecord) SetRequestMessage(v MediaStreamingAccessRecordAllOfRequestMessage) {
	o.RequestMessage = v
}

// GetCacheStatus returns the CacheStatus field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecord) GetCacheStatus() CacheStatus {
	if o == nil || isNil(o.CacheStatus) {
		var ret CacheStatus
		return ret
	}
	return *o.CacheStatus
}

// GetCacheStatusOk returns a tuple with the CacheStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetCacheStatusOk() (*CacheStatus, bool) {
	if o == nil || isNil(o.CacheStatus) {
    return nil, false
	}
	return o.CacheStatus, true
}

// HasCacheStatus returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecord) HasCacheStatus() bool {
	if o != nil && !isNil(o.CacheStatus) {
		return true
	}

	return false
}

// SetCacheStatus gets a reference to the given CacheStatus and assigns it to the CacheStatus field.
func (o *MediaStreamingAccessRecord) SetCacheStatus(v CacheStatus) {
	o.CacheStatus = &v
}

// GetResponseMessage returns the ResponseMessage field value
func (o *MediaStreamingAccessRecord) GetResponseMessage() MediaStreamingAccessRecordAllOfResponseMessage {
	if o == nil {
		var ret MediaStreamingAccessRecordAllOfResponseMessage
		return ret
	}

	return o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetResponseMessageOk() (*MediaStreamingAccessRecordAllOfResponseMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ResponseMessage, true
}

// SetResponseMessage sets field value
func (o *MediaStreamingAccessRecord) SetResponseMessage(v MediaStreamingAccessRecordAllOfResponseMessage) {
	o.ResponseMessage = v
}

// GetProcessingLatency returns the ProcessingLatency field value
func (o *MediaStreamingAccessRecord) GetProcessingLatency() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProcessingLatency
}

// GetProcessingLatencyOk returns a tuple with the ProcessingLatency field value
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetProcessingLatencyOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ProcessingLatency, true
}

// SetProcessingLatency sets field value
func (o *MediaStreamingAccessRecord) SetProcessingLatency(v float32) {
	o.ProcessingLatency = v
}

// GetConnectionMetrics returns the ConnectionMetrics field value if set, zero value otherwise.
func (o *MediaStreamingAccessRecord) GetConnectionMetrics() MediaStreamingAccessRecordAllOfConnectionMetrics {
	if o == nil || isNil(o.ConnectionMetrics) {
		var ret MediaStreamingAccessRecordAllOfConnectionMetrics
		return ret
	}
	return *o.ConnectionMetrics
}

// GetConnectionMetricsOk returns a tuple with the ConnectionMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaStreamingAccessRecord) GetConnectionMetricsOk() (*MediaStreamingAccessRecordAllOfConnectionMetrics, bool) {
	if o == nil || isNil(o.ConnectionMetrics) {
    return nil, false
	}
	return o.ConnectionMetrics, true
}

// HasConnectionMetrics returns a boolean if a field has been set.
func (o *MediaStreamingAccessRecord) HasConnectionMetrics() bool {
	if o != nil && !isNil(o.ConnectionMetrics) {
		return true
	}

	return false
}

// SetConnectionMetrics gets a reference to the given MediaStreamingAccessRecordAllOfConnectionMetrics and assigns it to the ConnectionMetrics field.
func (o *MediaStreamingAccessRecord) SetConnectionMetrics(v MediaStreamingAccessRecordAllOfConnectionMetrics) {
	o.ConnectionMetrics = &v
}

func (o MediaStreamingAccessRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["mediaStreamHandlerEndpointAddress"] = o.MediaStreamHandlerEndpointAddress
	}
	if true {
		toSerialize["applicationServerEndpointAddress"] = o.ApplicationServerEndpointAddress
	}
	if !isNil(o.SessionIdentifier) {
		toSerialize["sessionIdentifier"] = o.SessionIdentifier
	}
	if true {
		toSerialize["requestMessage"] = o.RequestMessage
	}
	if !isNil(o.CacheStatus) {
		toSerialize["cacheStatus"] = o.CacheStatus
	}
	if true {
		toSerialize["responseMessage"] = o.ResponseMessage
	}
	if true {
		toSerialize["processingLatency"] = o.ProcessingLatency
	}
	if !isNil(o.ConnectionMetrics) {
		toSerialize["connectionMetrics"] = o.ConnectionMetrics
	}
	return json.Marshal(toSerialize)
}

type NullableMediaStreamingAccessRecord struct {
	value *MediaStreamingAccessRecord
	isSet bool
}

func (v NullableMediaStreamingAccessRecord) Get() *MediaStreamingAccessRecord {
	return v.value
}

func (v *NullableMediaStreamingAccessRecord) Set(val *MediaStreamingAccessRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamingAccessRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamingAccessRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamingAccessRecord(val *MediaStreamingAccessRecord) *NullableMediaStreamingAccessRecord {
	return &NullableMediaStreamingAccessRecord{value: val, isSet: true}
}

func (v NullableMediaStreamingAccessRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamingAccessRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


