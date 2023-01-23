/*
3gpp-time-sync-exposure

API for time synchronization exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TimeSyncExposure

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type IndividualTimeSynchronizationExposureConfigurationApi interface {

	/*
	DeleteAnConfiguration Deletes an already existing configuration

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param afId Identifier of the AF
	@param subscriptionId Identifier of the subscription resource
	@param instanceReference Identifier of the configuration resource
	@return ApiDeleteAnConfigurationRequest
	*/
	DeleteAnConfiguration(ctx context.Context, afId string, subscriptionId string, instanceReference string) ApiDeleteAnConfigurationRequest

	// DeleteAnConfigurationExecute executes the request
	DeleteAnConfigurationExecute(r ApiDeleteAnConfigurationRequest) (*http.Response, error)

	/*
	FullyUpdateAnConfiguration Fully updates/replaces an existing configuration resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param afId Identifier of the AF
	@param subscriptionId Identifier of the subscription resource
	@param instanceReference Identifier of the configuration resource
	@return ApiFullyUpdateAnConfigurationRequest
	*/
	FullyUpdateAnConfiguration(ctx context.Context, afId string, subscriptionId string, instanceReference string) ApiFullyUpdateAnConfigurationRequest

	// FullyUpdateAnConfigurationExecute executes the request
	//  @return TimeSyncExposureConfig
	FullyUpdateAnConfigurationExecute(r ApiFullyUpdateAnConfigurationRequest) (*TimeSyncExposureConfig, *http.Response, error)
}

// IndividualTimeSynchronizationExposureConfigurationApiService IndividualTimeSynchronizationExposureConfigurationApi service
type IndividualTimeSynchronizationExposureConfigurationApiService service

type ApiDeleteAnConfigurationRequest struct {
	ctx context.Context
	ApiService IndividualTimeSynchronizationExposureConfigurationApi
	afId string
	subscriptionId string
	instanceReference string
}

func (r ApiDeleteAnConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAnConfigurationExecute(r)
}

/*
DeleteAnConfiguration Deletes an already existing configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param afId Identifier of the AF
 @param subscriptionId Identifier of the subscription resource
 @param instanceReference Identifier of the configuration resource
 @return ApiDeleteAnConfigurationRequest
*/
func (a *IndividualTimeSynchronizationExposureConfigurationApiService) DeleteAnConfiguration(ctx context.Context, afId string, subscriptionId string, instanceReference string) ApiDeleteAnConfigurationRequest {
	return ApiDeleteAnConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		afId: afId,
		subscriptionId: subscriptionId,
		instanceReference: instanceReference,
	}
}

// Execute executes the request
func (a *IndividualTimeSynchronizationExposureConfigurationApiService) DeleteAnConfigurationExecute(r ApiDeleteAnConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualTimeSynchronizationExposureConfigurationApiService.DeleteAnConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{afId}/subscriptions/{subscriptionId}/configurations/{instanceReference}"
	localVarPath = strings.Replace(localVarPath, "{"+"afId"+"}", url.PathEscape(parameterToString(r.afId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceReference"+"}", url.PathEscape(parameterToString(r.instanceReference, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiFullyUpdateAnConfigurationRequest struct {
	ctx context.Context
	ApiService IndividualTimeSynchronizationExposureConfigurationApi
	afId string
	subscriptionId string
	instanceReference string
	timeSyncExposureConfig *TimeSyncExposureConfig
}

// Parameters to update/replace the existing configuration
func (r ApiFullyUpdateAnConfigurationRequest) TimeSyncExposureConfig(timeSyncExposureConfig TimeSyncExposureConfig) ApiFullyUpdateAnConfigurationRequest {
	r.timeSyncExposureConfig = &timeSyncExposureConfig
	return r
}

func (r ApiFullyUpdateAnConfigurationRequest) Execute() (*TimeSyncExposureConfig, *http.Response, error) {
	return r.ApiService.FullyUpdateAnConfigurationExecute(r)
}

/*
FullyUpdateAnConfiguration Fully updates/replaces an existing configuration resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param afId Identifier of the AF
 @param subscriptionId Identifier of the subscription resource
 @param instanceReference Identifier of the configuration resource
 @return ApiFullyUpdateAnConfigurationRequest
*/
func (a *IndividualTimeSynchronizationExposureConfigurationApiService) FullyUpdateAnConfiguration(ctx context.Context, afId string, subscriptionId string, instanceReference string) ApiFullyUpdateAnConfigurationRequest {
	return ApiFullyUpdateAnConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		afId: afId,
		subscriptionId: subscriptionId,
		instanceReference: instanceReference,
	}
}

// Execute executes the request
//  @return TimeSyncExposureConfig
func (a *IndividualTimeSynchronizationExposureConfigurationApiService) FullyUpdateAnConfigurationExecute(r ApiFullyUpdateAnConfigurationRequest) (*TimeSyncExposureConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimeSyncExposureConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualTimeSynchronizationExposureConfigurationApiService.FullyUpdateAnConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{afId}/subscriptions/{subscriptionId}/configurations/{instanceReference}"
	localVarPath = strings.Replace(localVarPath, "{"+"afId"+"}", url.PathEscape(parameterToString(r.afId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceReference"+"}", url.PathEscape(parameterToString(r.instanceReference, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timeSyncExposureConfig == nil {
		return localVarReturnValue, nil, reportError("timeSyncExposureConfig is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.timeSyncExposureConfig
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 411 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
