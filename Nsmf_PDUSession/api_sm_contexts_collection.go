/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)


type SMContextsCollectionApi interface {

	/*
	PostSmContexts Create SM Context

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostSmContextsRequest
	*/
	PostSmContexts(ctx context.Context) ApiPostSmContextsRequest

	// PostSmContextsExecute executes the request
	//  @return SmContextCreatedData
	PostSmContextsExecute(r ApiPostSmContextsRequest) (*SmContextCreatedData, *http.Response, error)
}

// SMContextsCollectionApiService SMContextsCollectionApi service
type SMContextsCollectionApiService service

type ApiPostSmContextsRequest struct {
	ctx context.Context
	ApiService SMContextsCollectionApi
	jsonData *SmContextCreateData
	binaryDataN1SmMessage **os.File
	binaryDataN2SmInformation **os.File
	binaryDataN2SmInformationExt1 **os.File
}

func (r ApiPostSmContextsRequest) JsonData(jsonData SmContextCreateData) ApiPostSmContextsRequest {
	r.jsonData = &jsonData
	return r
}

func (r ApiPostSmContextsRequest) BinaryDataN1SmMessage(binaryDataN1SmMessage *os.File) ApiPostSmContextsRequest {
	r.binaryDataN1SmMessage = &binaryDataN1SmMessage
	return r
}

func (r ApiPostSmContextsRequest) BinaryDataN2SmInformation(binaryDataN2SmInformation *os.File) ApiPostSmContextsRequest {
	r.binaryDataN2SmInformation = &binaryDataN2SmInformation
	return r
}

func (r ApiPostSmContextsRequest) BinaryDataN2SmInformationExt1(binaryDataN2SmInformationExt1 *os.File) ApiPostSmContextsRequest {
	r.binaryDataN2SmInformationExt1 = &binaryDataN2SmInformationExt1
	return r
}

func (r ApiPostSmContextsRequest) Execute() (*SmContextCreatedData, *http.Response, error) {
	return r.ApiService.PostSmContextsExecute(r)
}

/*
PostSmContexts Create SM Context

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostSmContextsRequest
*/
func (a *SMContextsCollectionApiService) PostSmContexts(ctx context.Context) ApiPostSmContextsRequest {
	return ApiPostSmContextsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SmContextCreatedData
func (a *SMContextsCollectionApiService) PostSmContextsExecute(r ApiPostSmContextsRequest) (*SmContextCreatedData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmContextCreatedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMContextsCollectionApiService.PostSmContexts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sm-contexts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/related"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "multipart/related", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.jsonData != nil {
		paramJson, err := parameterToJson(*r.jsonData)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("jsonData", paramJson)
	}
	var binaryDataN1SmMessageLocalVarFormFileName string
	var binaryDataN1SmMessageLocalVarFileName     string
	var binaryDataN1SmMessageLocalVarFileBytes    []byte

	binaryDataN1SmMessageLocalVarFormFileName = "binaryDataN1SmMessage"

	var binaryDataN1SmMessageLocalVarFile *os.File
	if r.binaryDataN1SmMessage != nil {
		binaryDataN1SmMessageLocalVarFile = *r.binaryDataN1SmMessage
	}
	if binaryDataN1SmMessageLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataN1SmMessageLocalVarFile)
		binaryDataN1SmMessageLocalVarFileBytes = fbs
		binaryDataN1SmMessageLocalVarFileName = binaryDataN1SmMessageLocalVarFile.Name()
		binaryDataN1SmMessageLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataN1SmMessageLocalVarFileBytes, fileName: binaryDataN1SmMessageLocalVarFileName, formFileName: binaryDataN1SmMessageLocalVarFormFileName})
	var binaryDataN2SmInformationLocalVarFormFileName string
	var binaryDataN2SmInformationLocalVarFileName     string
	var binaryDataN2SmInformationLocalVarFileBytes    []byte

	binaryDataN2SmInformationLocalVarFormFileName = "binaryDataN2SmInformation"

	var binaryDataN2SmInformationLocalVarFile *os.File
	if r.binaryDataN2SmInformation != nil {
		binaryDataN2SmInformationLocalVarFile = *r.binaryDataN2SmInformation
	}
	if binaryDataN2SmInformationLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataN2SmInformationLocalVarFile)
		binaryDataN2SmInformationLocalVarFileBytes = fbs
		binaryDataN2SmInformationLocalVarFileName = binaryDataN2SmInformationLocalVarFile.Name()
		binaryDataN2SmInformationLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataN2SmInformationLocalVarFileBytes, fileName: binaryDataN2SmInformationLocalVarFileName, formFileName: binaryDataN2SmInformationLocalVarFormFileName})
	var binaryDataN2SmInformationExt1LocalVarFormFileName string
	var binaryDataN2SmInformationExt1LocalVarFileName     string
	var binaryDataN2SmInformationExt1LocalVarFileBytes    []byte

	binaryDataN2SmInformationExt1LocalVarFormFileName = "binaryDataN2SmInformationExt1"

	var binaryDataN2SmInformationExt1LocalVarFile *os.File
	if r.binaryDataN2SmInformationExt1 != nil {
		binaryDataN2SmInformationExt1LocalVarFile = *r.binaryDataN2SmInformationExt1
	}
	if binaryDataN2SmInformationExt1LocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataN2SmInformationExt1LocalVarFile)
		binaryDataN2SmInformationExt1LocalVarFileBytes = fbs
		binaryDataN2SmInformationExt1LocalVarFileName = binaryDataN2SmInformationExt1LocalVarFile.Name()
		binaryDataN2SmInformationExt1LocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataN2SmInformationExt1LocalVarFileBytes, fileName: binaryDataN2SmInformationExt1LocalVarFileName, formFileName: binaryDataN2SmInformationExt1LocalVarFormFileName})
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
		if localVarHTTPResponse.StatusCode == 307 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 308 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v SmContextCreateError
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
			var v SmContextCreateError
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
			var v SmContextCreateError
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
			var v ExtProblemDetails
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
			var v ExtProblemDetails
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
			var v ExtProblemDetails
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
			var v SmContextCreateError
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
			var v SmContextCreateError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 504 {
			var v SmContextCreateError
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
