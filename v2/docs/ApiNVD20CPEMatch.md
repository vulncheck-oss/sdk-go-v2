# ApiNVD20CPEMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpeLastModified** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**MatchCriteriaId** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to [**[]ApiNVD20CPEName**](ApiNVD20CPEName.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VersionEndExcluding** | Pointer to **string** |  | [optional] 
**VersionEndIncluding** | Pointer to **string** |  | [optional] 
**VersionStartExcluding** | Pointer to **string** |  | [optional] 
**VersionStartIncluding** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20CPEMatch

`func NewApiNVD20CPEMatch() *ApiNVD20CPEMatch`

NewApiNVD20CPEMatch instantiates a new ApiNVD20CPEMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CPEMatchWithDefaults

`func NewApiNVD20CPEMatchWithDefaults() *ApiNVD20CPEMatch`

NewApiNVD20CPEMatchWithDefaults instantiates a new ApiNVD20CPEMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpeLastModified

`func (o *ApiNVD20CPEMatch) GetCpeLastModified() string`

GetCpeLastModified returns the CpeLastModified field if non-nil, zero value otherwise.

### GetCpeLastModifiedOk

`func (o *ApiNVD20CPEMatch) GetCpeLastModifiedOk() (*string, bool)`

GetCpeLastModifiedOk returns a tuple with the CpeLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpeLastModified

`func (o *ApiNVD20CPEMatch) SetCpeLastModified(v string)`

SetCpeLastModified sets CpeLastModified field to given value.

### HasCpeLastModified

`func (o *ApiNVD20CPEMatch) HasCpeLastModified() bool`

HasCpeLastModified returns a boolean if a field has been set.

### GetCreated

`func (o *ApiNVD20CPEMatch) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiNVD20CPEMatch) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiNVD20CPEMatch) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiNVD20CPEMatch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCriteria

`func (o *ApiNVD20CPEMatch) GetCriteria() string`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ApiNVD20CPEMatch) GetCriteriaOk() (*string, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ApiNVD20CPEMatch) SetCriteria(v string)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ApiNVD20CPEMatch) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetLastModified

`func (o *ApiNVD20CPEMatch) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApiNVD20CPEMatch) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApiNVD20CPEMatch) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ApiNVD20CPEMatch) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMatchCriteriaId

`func (o *ApiNVD20CPEMatch) GetMatchCriteriaId() string`

GetMatchCriteriaId returns the MatchCriteriaId field if non-nil, zero value otherwise.

### GetMatchCriteriaIdOk

`func (o *ApiNVD20CPEMatch) GetMatchCriteriaIdOk() (*string, bool)`

GetMatchCriteriaIdOk returns a tuple with the MatchCriteriaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCriteriaId

`func (o *ApiNVD20CPEMatch) SetMatchCriteriaId(v string)`

SetMatchCriteriaId sets MatchCriteriaId field to given value.

### HasMatchCriteriaId

`func (o *ApiNVD20CPEMatch) HasMatchCriteriaId() bool`

HasMatchCriteriaId returns a boolean if a field has been set.

### GetMatches

`func (o *ApiNVD20CPEMatch) GetMatches() []ApiNVD20CPEName`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *ApiNVD20CPEMatch) GetMatchesOk() (*[]ApiNVD20CPEName, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *ApiNVD20CPEMatch) SetMatches(v []ApiNVD20CPEName)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *ApiNVD20CPEMatch) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetStatus

`func (o *ApiNVD20CPEMatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiNVD20CPEMatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiNVD20CPEMatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiNVD20CPEMatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersionEndExcluding

`func (o *ApiNVD20CPEMatch) GetVersionEndExcluding() string`

GetVersionEndExcluding returns the VersionEndExcluding field if non-nil, zero value otherwise.

### GetVersionEndExcludingOk

`func (o *ApiNVD20CPEMatch) GetVersionEndExcludingOk() (*string, bool)`

GetVersionEndExcludingOk returns a tuple with the VersionEndExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndExcluding

`func (o *ApiNVD20CPEMatch) SetVersionEndExcluding(v string)`

SetVersionEndExcluding sets VersionEndExcluding field to given value.

### HasVersionEndExcluding

`func (o *ApiNVD20CPEMatch) HasVersionEndExcluding() bool`

HasVersionEndExcluding returns a boolean if a field has been set.

### GetVersionEndIncluding

`func (o *ApiNVD20CPEMatch) GetVersionEndIncluding() string`

GetVersionEndIncluding returns the VersionEndIncluding field if non-nil, zero value otherwise.

### GetVersionEndIncludingOk

`func (o *ApiNVD20CPEMatch) GetVersionEndIncludingOk() (*string, bool)`

GetVersionEndIncludingOk returns a tuple with the VersionEndIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionEndIncluding

`func (o *ApiNVD20CPEMatch) SetVersionEndIncluding(v string)`

SetVersionEndIncluding sets VersionEndIncluding field to given value.

### HasVersionEndIncluding

`func (o *ApiNVD20CPEMatch) HasVersionEndIncluding() bool`

HasVersionEndIncluding returns a boolean if a field has been set.

### GetVersionStartExcluding

`func (o *ApiNVD20CPEMatch) GetVersionStartExcluding() string`

GetVersionStartExcluding returns the VersionStartExcluding field if non-nil, zero value otherwise.

### GetVersionStartExcludingOk

`func (o *ApiNVD20CPEMatch) GetVersionStartExcludingOk() (*string, bool)`

GetVersionStartExcludingOk returns a tuple with the VersionStartExcluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartExcluding

`func (o *ApiNVD20CPEMatch) SetVersionStartExcluding(v string)`

SetVersionStartExcluding sets VersionStartExcluding field to given value.

### HasVersionStartExcluding

`func (o *ApiNVD20CPEMatch) HasVersionStartExcluding() bool`

HasVersionStartExcluding returns a boolean if a field has been set.

### GetVersionStartIncluding

`func (o *ApiNVD20CPEMatch) GetVersionStartIncluding() string`

GetVersionStartIncluding returns the VersionStartIncluding field if non-nil, zero value otherwise.

### GetVersionStartIncludingOk

`func (o *ApiNVD20CPEMatch) GetVersionStartIncludingOk() (*string, bool)`

GetVersionStartIncludingOk returns a tuple with the VersionStartIncluding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStartIncluding

`func (o *ApiNVD20CPEMatch) SetVersionStartIncluding(v string)`

SetVersionStartIncluding sets VersionStartIncluding field to given value.

### HasVersionStartIncluding

`func (o *ApiNVD20CPEMatch) HasVersionStartIncluding() bool`

HasVersionStartIncluding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


