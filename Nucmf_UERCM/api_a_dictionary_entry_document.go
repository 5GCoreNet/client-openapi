/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nucmf_UERCM

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)


// ADictionaryEntryDocumentApiService ADictionaryEntryDocumentApi service
type ADictionaryEntryDocumentApiService service

type ApiCreateDictionaryEntryRequest struct {
	ctx context.Context
	ApiService *ADictionaryEntryDocumentApiService
	jsonData *DicEntryCreateData
	binaryDataUeRadioCapability5GS **os.File
	binaryDataUeRadioCapabilityEPS **os.File
	binaryDataUeRadioCap5GSForPaging **os.File
	binaryDataUeRadioCapEPSForPaging **os.File
}

func (r ApiCreateDictionaryEntryRequest) JsonData(jsonData DicEntryCreateData) ApiCreateDictionaryEntryRequest {
	r.jsonData = &jsonData
	return r
}

func (r ApiCreateDictionaryEntryRequest) BinaryDataUeRadioCapability5GS(binaryDataUeRadioCapability5GS *os.File) ApiCreateDictionaryEntryRequest {
	r.binaryDataUeRadioCapability5GS = &binaryDataUeRadioCapability5GS
	return r
}

func (r ApiCreateDictionaryEntryRequest) BinaryDataUeRadioCapabilityEPS(binaryDataUeRadioCapabilityEPS *os.File) ApiCreateDictionaryEntryRequest {
	r.binaryDataUeRadioCapabilityEPS = &binaryDataUeRadioCapabilityEPS
	return r
}

func (r ApiCreateDictionaryEntryRequest) BinaryDataUeRadioCap5GSForPaging(binaryDataUeRadioCap5GSForPaging *os.File) ApiCreateDictionaryEntryRequest {
	r.binaryDataUeRadioCap5GSForPaging = &binaryDataUeRadioCap5GSForPaging
	return r
}

func (r ApiCreateDictionaryEntryRequest) BinaryDataUeRadioCapEPSForPaging(binaryDataUeRadioCapEPSForPaging *os.File) ApiCreateDictionaryEntryRequest {
	r.binaryDataUeRadioCapEPSForPaging = &binaryDataUeRadioCapEPSForPaging
	return r
}

func (r ApiCreateDictionaryEntryRequest) Execute() (*DicEntryCreatedData, *http.Response, error) {
	return r.ApiService.CreateDictionaryEntryExecute(r)
}

/*
CreateDictionaryEntry Create a dictionary entry in the UCMF

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDictionaryEntryRequest
*/
func (a *ADictionaryEntryDocumentApiService) CreateDictionaryEntry(ctx context.Context) ApiCreateDictionaryEntryRequest {
	return ApiCreateDictionaryEntryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DicEntryCreatedData
func (a *ADictionaryEntryDocumentApiService) CreateDictionaryEntryExecute(r ApiCreateDictionaryEntryRequest) (*DicEntryCreatedData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DicEntryCreatedData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ADictionaryEntryDocumentApiService.CreateDictionaryEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dic-entries"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

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
	var binaryDataUeRadioCapability5GSLocalVarFormFileName string
	var binaryDataUeRadioCapability5GSLocalVarFileName     string
	var binaryDataUeRadioCapability5GSLocalVarFileBytes    []byte

	binaryDataUeRadioCapability5GSLocalVarFormFileName = "binaryDataUeRadioCapability5GS"

	var binaryDataUeRadioCapability5GSLocalVarFile *os.File
	if r.binaryDataUeRadioCapability5GS != nil {
		binaryDataUeRadioCapability5GSLocalVarFile = *r.binaryDataUeRadioCapability5GS
	}
	if binaryDataUeRadioCapability5GSLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataUeRadioCapability5GSLocalVarFile)
		binaryDataUeRadioCapability5GSLocalVarFileBytes = fbs
		binaryDataUeRadioCapability5GSLocalVarFileName = binaryDataUeRadioCapability5GSLocalVarFile.Name()
		binaryDataUeRadioCapability5GSLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataUeRadioCapability5GSLocalVarFileBytes, fileName: binaryDataUeRadioCapability5GSLocalVarFileName, formFileName: binaryDataUeRadioCapability5GSLocalVarFormFileName})
	var binaryDataUeRadioCapabilityEPSLocalVarFormFileName string
	var binaryDataUeRadioCapabilityEPSLocalVarFileName     string
	var binaryDataUeRadioCapabilityEPSLocalVarFileBytes    []byte

	binaryDataUeRadioCapabilityEPSLocalVarFormFileName = "binaryDataUeRadioCapabilityEPS"

	var binaryDataUeRadioCapabilityEPSLocalVarFile *os.File
	if r.binaryDataUeRadioCapabilityEPS != nil {
		binaryDataUeRadioCapabilityEPSLocalVarFile = *r.binaryDataUeRadioCapabilityEPS
	}
	if binaryDataUeRadioCapabilityEPSLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataUeRadioCapabilityEPSLocalVarFile)
		binaryDataUeRadioCapabilityEPSLocalVarFileBytes = fbs
		binaryDataUeRadioCapabilityEPSLocalVarFileName = binaryDataUeRadioCapabilityEPSLocalVarFile.Name()
		binaryDataUeRadioCapabilityEPSLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataUeRadioCapabilityEPSLocalVarFileBytes, fileName: binaryDataUeRadioCapabilityEPSLocalVarFileName, formFileName: binaryDataUeRadioCapabilityEPSLocalVarFormFileName})
	var binaryDataUeRadioCap5GSForPagingLocalVarFormFileName string
	var binaryDataUeRadioCap5GSForPagingLocalVarFileName     string
	var binaryDataUeRadioCap5GSForPagingLocalVarFileBytes    []byte

	binaryDataUeRadioCap5GSForPagingLocalVarFormFileName = "binaryDataUeRadioCap5GSForPaging"

	var binaryDataUeRadioCap5GSForPagingLocalVarFile *os.File
	if r.binaryDataUeRadioCap5GSForPaging != nil {
		binaryDataUeRadioCap5GSForPagingLocalVarFile = *r.binaryDataUeRadioCap5GSForPaging
	}
	if binaryDataUeRadioCap5GSForPagingLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataUeRadioCap5GSForPagingLocalVarFile)
		binaryDataUeRadioCap5GSForPagingLocalVarFileBytes = fbs
		binaryDataUeRadioCap5GSForPagingLocalVarFileName = binaryDataUeRadioCap5GSForPagingLocalVarFile.Name()
		binaryDataUeRadioCap5GSForPagingLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataUeRadioCap5GSForPagingLocalVarFileBytes, fileName: binaryDataUeRadioCap5GSForPagingLocalVarFileName, formFileName: binaryDataUeRadioCap5GSForPagingLocalVarFormFileName})
	var binaryDataUeRadioCapEPSForPagingLocalVarFormFileName string
	var binaryDataUeRadioCapEPSForPagingLocalVarFileName     string
	var binaryDataUeRadioCapEPSForPagingLocalVarFileBytes    []byte

	binaryDataUeRadioCapEPSForPagingLocalVarFormFileName = "binaryDataUeRadioCapEPSForPaging"

	var binaryDataUeRadioCapEPSForPagingLocalVarFile *os.File
	if r.binaryDataUeRadioCapEPSForPaging != nil {
		binaryDataUeRadioCapEPSForPagingLocalVarFile = *r.binaryDataUeRadioCapEPSForPaging
	}
	if binaryDataUeRadioCapEPSForPagingLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(binaryDataUeRadioCapEPSForPagingLocalVarFile)
		binaryDataUeRadioCapEPSForPagingLocalVarFileBytes = fbs
		binaryDataUeRadioCapEPSForPagingLocalVarFileName = binaryDataUeRadioCapEPSForPagingLocalVarFile.Name()
		binaryDataUeRadioCapEPSForPagingLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: binaryDataUeRadioCapEPSForPagingLocalVarFileBytes, fileName: binaryDataUeRadioCapEPSForPagingLocalVarFileName, formFileName: binaryDataUeRadioCapEPSForPagingLocalVarFormFileName})
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
		if localVarHTTPResponse.StatusCode == 501 {
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
