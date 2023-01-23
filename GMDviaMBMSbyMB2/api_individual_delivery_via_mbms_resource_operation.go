/*
GMDviaMBMSbyMB2

API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GMDviaMBMSbyMB2

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type IndividualDeliveryViaMBMSResourceOperationApi interface {

	/*
	DeleteIndDeliveryViaMBMS Deletes a delivery via MBMS resource for a given SCS/AS, a TMGI and a transcation Id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scsAsId Identifier of SCS/AS
	@param tmgi TMGI
	@param transactionId Identifier of transaction
	@return ApiDeleteIndDeliveryViaMBMSRequest
	*/
	DeleteIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiDeleteIndDeliveryViaMBMSRequest

	// DeleteIndDeliveryViaMBMSExecute executes the request
	DeleteIndDeliveryViaMBMSExecute(r ApiDeleteIndDeliveryViaMBMSRequest) (*http.Response, error)

	/*
	FetchIndDeliveryViaMBMS Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scsAsId Identifier of SCS/AS
	@param tmgi TMGI
	@param transactionId Identifier of transaction
	@return ApiFetchIndDeliveryViaMBMSRequest
	*/
	FetchIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiFetchIndDeliveryViaMBMSRequest

	// FetchIndDeliveryViaMBMSExecute executes the request
	//  @return GMDViaMBMSByMb2
	FetchIndDeliveryViaMBMSExecute(r ApiFetchIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error)

	/*
	ModifyIndDeliveryViaMBMS Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scsAsId Identifier of SCS/AS
	@param tmgi TMGI
	@param transactionId Identifier of transaction
	@return ApiModifyIndDeliveryViaMBMSRequest
	*/
	ModifyIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiModifyIndDeliveryViaMBMSRequest

	// ModifyIndDeliveryViaMBMSExecute executes the request
	//  @return GMDViaMBMSByMb2
	ModifyIndDeliveryViaMBMSExecute(r ApiModifyIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error)

	/*
	UpdateIndDeliveryViaMBMS Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param scsAsId Identifier of SCS/AS
	@param tmgi TMGI
	@param transactionId Identifier of transaction
	@return ApiUpdateIndDeliveryViaMBMSRequest
	*/
	UpdateIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiUpdateIndDeliveryViaMBMSRequest

	// UpdateIndDeliveryViaMBMSExecute executes the request
	//  @return GMDViaMBMSByMb2
	UpdateIndDeliveryViaMBMSExecute(r ApiUpdateIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error)
}

// IndividualDeliveryViaMBMSResourceOperationApiService IndividualDeliveryViaMBMSResourceOperationApi service
type IndividualDeliveryViaMBMSResourceOperationApiService service

type ApiDeleteIndDeliveryViaMBMSRequest struct {
	ctx context.Context
	ApiService IndividualDeliveryViaMBMSResourceOperationApi
	scsAsId string
	tmgi string
	transactionId string
}

func (r ApiDeleteIndDeliveryViaMBMSRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIndDeliveryViaMBMSExecute(r)
}

/*
DeleteIndDeliveryViaMBMS Deletes a delivery via MBMS resource for a given SCS/AS, a TMGI and a transcation Id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param scsAsId Identifier of SCS/AS
 @param tmgi TMGI
 @param transactionId Identifier of transaction
 @return ApiDeleteIndDeliveryViaMBMSRequest
*/
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) DeleteIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiDeleteIndDeliveryViaMBMSRequest {
	return ApiDeleteIndDeliveryViaMBMSRequest{
		ApiService: a,
		ctx: ctx,
		scsAsId: scsAsId,
		tmgi: tmgi,
		transactionId: transactionId,
	}
}

// Execute executes the request
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) DeleteIndDeliveryViaMBMSExecute(r ApiDeleteIndDeliveryViaMBMSRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualDeliveryViaMBMSResourceOperationApiService.DeleteIndDeliveryViaMBMS")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"scsAsId"+"}", url.PathEscape(parameterToString(r.scsAsId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tmgi"+"}", url.PathEscape(parameterToString(r.tmgi, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterToString(r.transactionId, "")), -1)

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

type ApiFetchIndDeliveryViaMBMSRequest struct {
	ctx context.Context
	ApiService IndividualDeliveryViaMBMSResourceOperationApi
	scsAsId string
	tmgi string
	transactionId string
}

func (r ApiFetchIndDeliveryViaMBMSRequest) Execute() (*GMDViaMBMSByMb2, *http.Response, error) {
	return r.ApiService.FetchIndDeliveryViaMBMSExecute(r)
}

/*
FetchIndDeliveryViaMBMS Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param scsAsId Identifier of SCS/AS
 @param tmgi TMGI
 @param transactionId Identifier of transaction
 @return ApiFetchIndDeliveryViaMBMSRequest
*/
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) FetchIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiFetchIndDeliveryViaMBMSRequest {
	return ApiFetchIndDeliveryViaMBMSRequest{
		ApiService: a,
		ctx: ctx,
		scsAsId: scsAsId,
		tmgi: tmgi,
		transactionId: transactionId,
	}
}

// Execute executes the request
//  @return GMDViaMBMSByMb2
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) FetchIndDeliveryViaMBMSExecute(r ApiFetchIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GMDViaMBMSByMb2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualDeliveryViaMBMSResourceOperationApiService.FetchIndDeliveryViaMBMS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"scsAsId"+"}", url.PathEscape(parameterToString(r.scsAsId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tmgi"+"}", url.PathEscape(parameterToString(r.tmgi, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterToString(r.transactionId, "")), -1)

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

type ApiModifyIndDeliveryViaMBMSRequest struct {
	ctx context.Context
	ApiService IndividualDeliveryViaMBMSResourceOperationApi
	scsAsId string
	tmgi string
	transactionId string
	gMDViaMBMSByMb2Patch *GMDViaMBMSByMb2Patch
}

// representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF
func (r ApiModifyIndDeliveryViaMBMSRequest) GMDViaMBMSByMb2Patch(gMDViaMBMSByMb2Patch GMDViaMBMSByMb2Patch) ApiModifyIndDeliveryViaMBMSRequest {
	r.gMDViaMBMSByMb2Patch = &gMDViaMBMSByMb2Patch
	return r
}

func (r ApiModifyIndDeliveryViaMBMSRequest) Execute() (*GMDViaMBMSByMb2, *http.Response, error) {
	return r.ApiService.ModifyIndDeliveryViaMBMSExecute(r)
}

/*
ModifyIndDeliveryViaMBMS Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param scsAsId Identifier of SCS/AS
 @param tmgi TMGI
 @param transactionId Identifier of transaction
 @return ApiModifyIndDeliveryViaMBMSRequest
*/
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) ModifyIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiModifyIndDeliveryViaMBMSRequest {
	return ApiModifyIndDeliveryViaMBMSRequest{
		ApiService: a,
		ctx: ctx,
		scsAsId: scsAsId,
		tmgi: tmgi,
		transactionId: transactionId,
	}
}

// Execute executes the request
//  @return GMDViaMBMSByMb2
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) ModifyIndDeliveryViaMBMSExecute(r ApiModifyIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GMDViaMBMSByMb2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualDeliveryViaMBMSResourceOperationApiService.ModifyIndDeliveryViaMBMS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"scsAsId"+"}", url.PathEscape(parameterToString(r.scsAsId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tmgi"+"}", url.PathEscape(parameterToString(r.tmgi, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterToString(r.transactionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gMDViaMBMSByMb2Patch == nil {
		return localVarReturnValue, nil, reportError("gMDViaMBMSByMb2Patch is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json"}

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
	localVarPostBody = r.gMDViaMBMSByMb2Patch
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

type ApiUpdateIndDeliveryViaMBMSRequest struct {
	ctx context.Context
	ApiService IndividualDeliveryViaMBMSResourceOperationApi
	scsAsId string
	tmgi string
	transactionId string
	gMDViaMBMSByMb2 *GMDViaMBMSByMb2
}

// representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF
func (r ApiUpdateIndDeliveryViaMBMSRequest) GMDViaMBMSByMb2(gMDViaMBMSByMb2 GMDViaMBMSByMb2) ApiUpdateIndDeliveryViaMBMSRequest {
	r.gMDViaMBMSByMb2 = &gMDViaMBMSByMb2
	return r
}

func (r ApiUpdateIndDeliveryViaMBMSRequest) Execute() (*GMDViaMBMSByMb2, *http.Response, error) {
	return r.ApiService.UpdateIndDeliveryViaMBMSExecute(r)
}

/*
UpdateIndDeliveryViaMBMS Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param scsAsId Identifier of SCS/AS
 @param tmgi TMGI
 @param transactionId Identifier of transaction
 @return ApiUpdateIndDeliveryViaMBMSRequest
*/
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) UpdateIndDeliveryViaMBMS(ctx context.Context, scsAsId string, tmgi string, transactionId string) ApiUpdateIndDeliveryViaMBMSRequest {
	return ApiUpdateIndDeliveryViaMBMSRequest{
		ApiService: a,
		ctx: ctx,
		scsAsId: scsAsId,
		tmgi: tmgi,
		transactionId: transactionId,
	}
}

// Execute executes the request
//  @return GMDViaMBMSByMb2
func (a *IndividualDeliveryViaMBMSResourceOperationApiService) UpdateIndDeliveryViaMBMSExecute(r ApiUpdateIndDeliveryViaMBMSRequest) (*GMDViaMBMSByMb2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GMDViaMBMSByMb2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualDeliveryViaMBMSResourceOperationApiService.UpdateIndDeliveryViaMBMS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"scsAsId"+"}", url.PathEscape(parameterToString(r.scsAsId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tmgi"+"}", url.PathEscape(parameterToString(r.tmgi, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transactionId"+"}", url.PathEscape(parameterToString(r.transactionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gMDViaMBMSByMb2 == nil {
		return localVarReturnValue, nil, reportError("gMDViaMBMSByMb2 is required and must be specified")
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
	localVarPostBody = r.gMDViaMBMSByMb2
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
