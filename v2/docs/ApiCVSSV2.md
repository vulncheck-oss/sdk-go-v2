# ApiCVSSV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessComplexity** | Pointer to **string** |  | [optional] 
**AccessVector** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**VectorString** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCVSSV2

`func NewApiCVSSV2() *ApiCVSSV2`

NewApiCVSSV2 instantiates a new ApiCVSSV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCVSSV2WithDefaults

`func NewApiCVSSV2WithDefaults() *ApiCVSSV2`

NewApiCVSSV2WithDefaults instantiates a new ApiCVSSV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessComplexity

`func (o *ApiCVSSV2) GetAccessComplexity() string`

GetAccessComplexity returns the AccessComplexity field if non-nil, zero value otherwise.

### GetAccessComplexityOk

`func (o *ApiCVSSV2) GetAccessComplexityOk() (*string, bool)`

GetAccessComplexityOk returns a tuple with the AccessComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessComplexity

`func (o *ApiCVSSV2) SetAccessComplexity(v string)`

SetAccessComplexity sets AccessComplexity field to given value.

### HasAccessComplexity

`func (o *ApiCVSSV2) HasAccessComplexity() bool`

HasAccessComplexity returns a boolean if a field has been set.

### GetAccessVector

`func (o *ApiCVSSV2) GetAccessVector() string`

GetAccessVector returns the AccessVector field if non-nil, zero value otherwise.

### GetAccessVectorOk

`func (o *ApiCVSSV2) GetAccessVectorOk() (*string, bool)`

GetAccessVectorOk returns a tuple with the AccessVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVector

`func (o *ApiCVSSV2) SetAccessVector(v string)`

SetAccessVector sets AccessVector field to given value.

### HasAccessVector

`func (o *ApiCVSSV2) HasAccessVector() bool`

HasAccessVector returns a boolean if a field has been set.

### GetAuthentication

`func (o *ApiCVSSV2) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ApiCVSSV2) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ApiCVSSV2) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *ApiCVSSV2) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *ApiCVSSV2) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *ApiCVSSV2) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *ApiCVSSV2) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *ApiCVSSV2) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetBaseScore

`func (o *ApiCVSSV2) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *ApiCVSSV2) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *ApiCVSSV2) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *ApiCVSSV2) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *ApiCVSSV2) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *ApiCVSSV2) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *ApiCVSSV2) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *ApiCVSSV2) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *ApiCVSSV2) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *ApiCVSSV2) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *ApiCVSSV2) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *ApiCVSSV2) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetVectorString

`func (o *ApiCVSSV2) GetVectorString() string`

GetVectorString returns the VectorString field if non-nil, zero value otherwise.

### GetVectorStringOk

`func (o *ApiCVSSV2) GetVectorStringOk() (*string, bool)`

GetVectorStringOk returns a tuple with the VectorString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorString

`func (o *ApiCVSSV2) SetVectorString(v string)`

SetVectorString sets VectorString field to given value.

### HasVectorString

`func (o *ApiCVSSV2) HasVectorString() bool`

HasVectorString returns a boolean if a field has been set.

### GetVersion

`func (o *ApiCVSSV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiCVSSV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiCVSSV2) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiCVSSV2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


