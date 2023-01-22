/*
CAPIF_Auditing_API

API for auditing.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CAPIF_Auditing_API

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiApiInvocationLogsGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	aefId *string
	apiInvokerId *string
	timeRangeStart *time.Time
	timeRangeEnd *time.Time
	apiId *string
	apiName *string
	apiVersion *string
	protocol *Protocol
	operation *Operation
	result *string
	resourceName *string
	srcInterface *InterfaceDescription
	destInterface *InterfaceDescription
	supportedFeatures *string
}

// String identifying the API exposing function.
func (r ApiApiInvocationLogsGetRequest) AefId(aefId string) ApiApiInvocationLogsGetRequest {
	r.aefId = &aefId
	return r
}

// String identifying the API invoker which invoked the service API.
func (r ApiApiInvocationLogsGetRequest) ApiInvokerId(apiInvokerId string) ApiApiInvocationLogsGetRequest {
	r.apiInvokerId = &apiInvokerId
	return r
}

// Start time of the invocation time range.
func (r ApiApiInvocationLogsGetRequest) TimeRangeStart(timeRangeStart time.Time) ApiApiInvocationLogsGetRequest {
	r.timeRangeStart = &timeRangeStart
	return r
}

// End time of the invocation time range.
func (r ApiApiInvocationLogsGetRequest) TimeRangeEnd(timeRangeEnd time.Time) ApiApiInvocationLogsGetRequest {
	r.timeRangeEnd = &timeRangeEnd
	return r
}

// String identifying the API invoked.
func (r ApiApiInvocationLogsGetRequest) ApiId(apiId string) ApiApiInvocationLogsGetRequest {
	r.apiId = &apiId
	return r
}

// API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122. 
func (r ApiApiInvocationLogsGetRequest) ApiName(apiName string) ApiApiInvocationLogsGetRequest {
	r.apiName = &apiName
	return r
}

// Version of the API which was invoked.
func (r ApiApiInvocationLogsGetRequest) ApiVersion(apiVersion string) ApiApiInvocationLogsGetRequest {
	r.apiVersion = &apiVersion
	return r
}

// Protocol invoked.
func (r ApiApiInvocationLogsGetRequest) Protocol(protocol Protocol) ApiApiInvocationLogsGetRequest {
	r.protocol = &protocol
	return r
}

// Operation that was invoked on the API.
func (r ApiApiInvocationLogsGetRequest) Operation(operation Operation) ApiApiInvocationLogsGetRequest {
	r.operation = &operation
	return r
}

// Result or output of the invocation.
func (r ApiApiInvocationLogsGetRequest) Result(result string) ApiApiInvocationLogsGetRequest {
	r.result = &result
	return r
}

// Name of the specific resource invoked.
func (r ApiApiInvocationLogsGetRequest) ResourceName(resourceName string) ApiApiInvocationLogsGetRequest {
	r.resourceName = &resourceName
	return r
}

// Interface description of the API invoker.
func (r ApiApiInvocationLogsGetRequest) SrcInterface(srcInterface InterfaceDescription) ApiApiInvocationLogsGetRequest {
	r.srcInterface = &srcInterface
	return r
}

// Interface description of the API invoked.
func (r ApiApiInvocationLogsGetRequest) DestInterface(destInterface InterfaceDescription) ApiApiInvocationLogsGetRequest {
	r.destInterface = &destInterface
	return r
}

// To filter irrelevant responses related to unsupported features
func (r ApiApiInvocationLogsGetRequest) SupportedFeatures(supportedFeatures string) ApiApiInvocationLogsGetRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiApiInvocationLogsGetRequest) Execute() (*InvocationLog, *http.Response, error) {
	return r.ApiService.ApiInvocationLogsGetExecute(r)
}

/*
ApiInvocationLogsGet Method for ApiInvocationLogsGet

Query and retrieve service API invocation logs stored on the CAPIF core function.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiInvocationLogsGetRequest
*/
func (a *DefaultApiService) ApiInvocationLogsGet(ctx context.Context) ApiApiInvocationLogsGetRequest {
	return ApiApiInvocationLogsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvocationLog
func (a *DefaultApiService) ApiInvocationLogsGetExecute(r ApiApiInvocationLogsGetRequest) (*InvocationLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvocationLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ApiInvocationLogsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/apiInvocationLogs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aefId != nil {
		localVarQueryParams.Add("aef-id", parameterToString(*r.aefId, ""))
	}
	if r.apiInvokerId != nil {
		localVarQueryParams.Add("api-invoker-id", parameterToString(*r.apiInvokerId, ""))
	}
	if r.timeRangeStart != nil {
		localVarQueryParams.Add("time-range-start", parameterToString(*r.timeRangeStart, ""))
	}
	if r.timeRangeEnd != nil {
		localVarQueryParams.Add("time-range-end", parameterToString(*r.timeRangeEnd, ""))
	}
	if r.apiId != nil {
		localVarQueryParams.Add("api-id", parameterToString(*r.apiId, ""))
	}
	if r.apiName != nil {
		localVarQueryParams.Add("api-name", parameterToString(*r.apiName, ""))
	}
	if r.apiVersion != nil {
		localVarQueryParams.Add("api-version", parameterToString(*r.apiVersion, ""))
	}
	if r.protocol != nil {
		localVarQueryParams.Add("protocol", parameterToString(*r.protocol, ""))
	}
	if r.operation != nil {
		localVarQueryParams.Add("operation", parameterToString(*r.operation, ""))
	}
	if r.result != nil {
		localVarQueryParams.Add("result", parameterToString(*r.result, ""))
	}
	if r.resourceName != nil {
		localVarQueryParams.Add("resource-name", parameterToString(*r.resourceName, ""))
	}
	if r.srcInterface != nil {
		localVarQueryParams.Add("src-interface", parameterToString(*r.srcInterface, ""))
	}
	if r.destInterface != nil {
		localVarQueryParams.Add("dest-interface", parameterToString(*r.destInterface, ""))
	}
	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 406 {
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
		if localVarHTTPResponse.StatusCode == 414 {
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
