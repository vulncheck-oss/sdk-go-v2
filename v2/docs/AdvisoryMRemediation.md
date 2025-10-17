# AdvisoryMRemediation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedFiles** | Pointer to [**[]AdvisoryAffectedFile**](AdvisoryAffectedFile.md) |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**DateSpecified** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to [**AdvisoryIVal**](AdvisoryIVal.md) |  | [optional] 
**FixedBuild** | Pointer to **string** |  | [optional] 
**ProductID** | Pointer to **[]string** |  | [optional] 
**RestartRequired** | Pointer to [**AdvisoryIVal**](AdvisoryIVal.md) |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | diff | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Supercedence** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMRemediation

`func NewAdvisoryMRemediation() *AdvisoryMRemediation`

NewAdvisoryMRemediation instantiates a new AdvisoryMRemediation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMRemediationWithDefaults

`func NewAdvisoryMRemediationWithDefaults() *AdvisoryMRemediation`

NewAdvisoryMRemediationWithDefaults instantiates a new AdvisoryMRemediation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedFiles

`func (o *AdvisoryMRemediation) GetAffectedFiles() []AdvisoryAffectedFile`

GetAffectedFiles returns the AffectedFiles field if non-nil, zero value otherwise.

### GetAffectedFilesOk

`func (o *AdvisoryMRemediation) GetAffectedFilesOk() (*[]AdvisoryAffectedFile, bool)`

GetAffectedFilesOk returns a tuple with the AffectedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedFiles

`func (o *AdvisoryMRemediation) SetAffectedFiles(v []AdvisoryAffectedFile)`

SetAffectedFiles sets AffectedFiles field to given value.

### HasAffectedFiles

`func (o *AdvisoryMRemediation) HasAffectedFiles() bool`

HasAffectedFiles returns a boolean if a field has been set.

### GetDate

`func (o *AdvisoryMRemediation) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryMRemediation) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryMRemediation) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryMRemediation) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateSpecified

`func (o *AdvisoryMRemediation) GetDateSpecified() bool`

GetDateSpecified returns the DateSpecified field if non-nil, zero value otherwise.

### GetDateSpecifiedOk

`func (o *AdvisoryMRemediation) GetDateSpecifiedOk() (*bool, bool)`

GetDateSpecifiedOk returns a tuple with the DateSpecified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSpecified

`func (o *AdvisoryMRemediation) SetDateSpecified(v bool)`

SetDateSpecified sets DateSpecified field to given value.

### HasDateSpecified

`func (o *AdvisoryMRemediation) HasDateSpecified() bool`

HasDateSpecified returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryMRemediation) GetDescription() AdvisoryIVal`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMRemediation) GetDescriptionOk() (*AdvisoryIVal, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMRemediation) SetDescription(v AdvisoryIVal)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMRemediation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedBuild

`func (o *AdvisoryMRemediation) GetFixedBuild() string`

GetFixedBuild returns the FixedBuild field if non-nil, zero value otherwise.

### GetFixedBuildOk

`func (o *AdvisoryMRemediation) GetFixedBuildOk() (*string, bool)`

GetFixedBuildOk returns a tuple with the FixedBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedBuild

`func (o *AdvisoryMRemediation) SetFixedBuild(v string)`

SetFixedBuild sets FixedBuild field to given value.

### HasFixedBuild

`func (o *AdvisoryMRemediation) HasFixedBuild() bool`

HasFixedBuild returns a boolean if a field has been set.

### GetProductID

`func (o *AdvisoryMRemediation) GetProductID() []string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *AdvisoryMRemediation) GetProductIDOk() (*[]string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *AdvisoryMRemediation) SetProductID(v []string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *AdvisoryMRemediation) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetRestartRequired

`func (o *AdvisoryMRemediation) GetRestartRequired() AdvisoryIVal`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *AdvisoryMRemediation) GetRestartRequiredOk() (*AdvisoryIVal, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *AdvisoryMRemediation) SetRestartRequired(v AdvisoryIVal)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *AdvisoryMRemediation) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.

### GetSubType

`func (o *AdvisoryMRemediation) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *AdvisoryMRemediation) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *AdvisoryMRemediation) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *AdvisoryMRemediation) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryMRemediation) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryMRemediation) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryMRemediation) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryMRemediation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMRemediation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMRemediation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMRemediation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMRemediation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSupercedence

`func (o *AdvisoryMRemediation) GetSupercedence() string`

GetSupercedence returns the Supercedence field if non-nil, zero value otherwise.

### GetSupercedenceOk

`func (o *AdvisoryMRemediation) GetSupercedenceOk() (*string, bool)`

GetSupercedenceOk returns a tuple with the Supercedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupercedence

`func (o *AdvisoryMRemediation) SetSupercedence(v string)`

SetSupercedence sets Supercedence field to given value.

### HasSupercedence

`func (o *AdvisoryMRemediation) HasSupercedence() bool`

HasSupercedence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


