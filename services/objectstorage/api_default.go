/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage.  **Disclaimer**: The API is still under development.

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	oapiError "github.com/stackitcloud/stackit-sdk-go/core/oapierror"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateBucketRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	bucketName string
}

func (r ApiCreateBucketRequest) Execute() (*CreateBucketResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateBucketResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateBucket")
	if err != nil {
		return localVarReturnValue, &oapiError.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/project/{projectId}/bucket/{bucketName}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bucketName) < 3 {
		return localVarReturnValue, fmt.Errorf("bucketName must have at least 3 elements")
	}
	if strlen(r.bucketName) > 63 {
		return localVarReturnValue, fmt.Errorf("bucketName must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
CreateBucket Create Bucket

Create a bucket for the given project. Bucket with the same name
cannot already exists in the object storage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId STACKIT project ID
	@param bucketName The name has to be dns-conform.
	@return ApiCreateBucketRequest
*/
func (a *APIClient) CreateBucket(ctx context.Context, projectId string, bucketName string) ApiCreateBucketRequest {
	return ApiCreateBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
}

func (a *APIClient) CreateBucketExecute(ctx context.Context, projectId string, bucketName string) (*CreateBucketResponse, error) {
	r := ApiCreateBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
	return r.Execute()
}

type ApiDeleteBucketRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	bucketName string
}

func (r ApiDeleteBucketRequest) Execute() (*DeleteBucketResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeleteBucketResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteBucket")
	if err != nil {
		return localVarReturnValue, &oapiError.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/project/{projectId}/bucket/{bucketName}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bucketName) < 3 {
		return localVarReturnValue, fmt.Errorf("bucketName must have at least 3 elements")
	}
	if strlen(r.bucketName) > 63 {
		return localVarReturnValue, fmt.Errorf("bucketName must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
DeleteBucket Delete Bucket

Delete a bucket from the given project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId STACKIT project ID
	@param bucketName The name has to be dns-conform.
	@return ApiDeleteBucketRequest
*/
func (a *APIClient) DeleteBucket(ctx context.Context, projectId string, bucketName string) ApiDeleteBucketRequest {
	return ApiDeleteBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
}

func (a *APIClient) DeleteBucketExecute(ctx context.Context, projectId string, bucketName string) (*DeleteBucketResponse, error) {
	r := ApiDeleteBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
	return r.Execute()
}

type ApiGetBucketRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
	bucketName string
}

func (r ApiGetBucketRequest) Execute() (*GetBucketResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBucketResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetBucket")
	if err != nil {
		return localVarReturnValue, &oapiError.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/project/{projectId}/bucket/{bucketName}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"bucketName"+"}", url.PathEscape(ParameterValueToString(r.bucketName, "bucketName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bucketName) < 3 {
		return localVarReturnValue, fmt.Errorf("bucketName must have at least 3 elements")
	}
	if strlen(r.bucketName) > 63 {
		return localVarReturnValue, fmt.Errorf("bucketName must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
GetBucket Get Bucket

Get information for the given bucket in the project.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId STACKIT project ID
	@param bucketName The name has to be dns-conform.
	@return ApiGetBucketRequest
*/
func (a *APIClient) GetBucket(ctx context.Context, projectId string, bucketName string) ApiGetBucketRequest {
	return ApiGetBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
}

func (a *APIClient) GetBucketExecute(ctx context.Context, projectId string, bucketName string) (*GetBucketResponse, error) {
	r := ApiGetBucketRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
		bucketName: bucketName,
	}
	return r.Execute()
}

type ApiGetBucketsRequest struct {
	ctx        context.Context
	apiService *DefaultApiService
	projectId  string
}

func (r ApiGetBucketsRequest) Execute() (*GetBucketsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetBucketsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetBuckets")
	if err != nil {
		return localVarReturnValue, &oapiError.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/v1/project/{projectId}/buckets"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(ParameterValueToString(r.projectId, "projectId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.ErrorMessage = oapiError.FormatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.Model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &oapiError.GenericOpenAPIError{
			StatusCode:   localVarHTTPResponse.StatusCode,
			Body:         localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
GetBuckets Get Buckets

Get the list of all buckets in the given project

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId STACKIT project ID
	@return ApiGetBucketsRequest
*/
func (a *APIClient) GetBuckets(ctx context.Context, projectId string) ApiGetBucketsRequest {
	return ApiGetBucketsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
}

func (a *APIClient) GetBucketsExecute(ctx context.Context, projectId string) (*GetBucketsResponse, error) {
	r := ApiGetBucketsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
		projectId:  projectId,
	}
	return r.Execute()
}
