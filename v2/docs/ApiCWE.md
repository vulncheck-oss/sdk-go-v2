# ApiCWE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstraction** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**KevCount** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Structure** | Pointer to **string** |  | [optional] 
**VulncheckNvdCount** | Pointer to **int32** |  | [optional] 
**WeaknessId** | Pointer to **string** |  | [optional] 
**WeaknessName** | Pointer to **string** |  | [optional] 
**WeightedScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewApiCWE

`func NewApiCWE() *ApiCWE`

NewApiCWE instantiates a new ApiCWE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCWEWithDefaults

`func NewApiCWEWithDefaults() *ApiCWE`

NewApiCWEWithDefaults instantiates a new ApiCWE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstraction

`func (o *ApiCWE) GetAbstraction() string`

GetAbstraction returns the Abstraction field if non-nil, zero value otherwise.

### GetAbstractionOk

`func (o *ApiCWE) GetAbstractionOk() (*string, bool)`

GetAbstractionOk returns a tuple with the Abstraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstraction

`func (o *ApiCWE) SetAbstraction(v string)`

SetAbstraction sets Abstraction field to given value.

### HasAbstraction

`func (o *ApiCWE) HasAbstraction() bool`

HasAbstraction returns a boolean if a field has been set.

### GetDescription

`func (o *ApiCWE) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiCWE) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiCWE) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiCWE) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKevCount

`func (o *ApiCWE) GetKevCount() int32`

GetKevCount returns the KevCount field if non-nil, zero value otherwise.

### GetKevCountOk

`func (o *ApiCWE) GetKevCountOk() (*int32, bool)`

GetKevCountOk returns a tuple with the KevCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKevCount

`func (o *ApiCWE) SetKevCount(v int32)`

SetKevCount sets KevCount field to given value.

### HasKevCount

`func (o *ApiCWE) HasKevCount() bool`

HasKevCount returns a boolean if a field has been set.

### GetStatus

`func (o *ApiCWE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiCWE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiCWE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiCWE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStructure

`func (o *ApiCWE) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *ApiCWE) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *ApiCWE) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *ApiCWE) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetVulncheckNvdCount

`func (o *ApiCWE) GetVulncheckNvdCount() int32`

GetVulncheckNvdCount returns the VulncheckNvdCount field if non-nil, zero value otherwise.

### GetVulncheckNvdCountOk

`func (o *ApiCWE) GetVulncheckNvdCountOk() (*int32, bool)`

GetVulncheckNvdCountOk returns a tuple with the VulncheckNvdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulncheckNvdCount

`func (o *ApiCWE) SetVulncheckNvdCount(v int32)`

SetVulncheckNvdCount sets VulncheckNvdCount field to given value.

### HasVulncheckNvdCount

`func (o *ApiCWE) HasVulncheckNvdCount() bool`

HasVulncheckNvdCount returns a boolean if a field has been set.

### GetWeaknessId

`func (o *ApiCWE) GetWeaknessId() string`

GetWeaknessId returns the WeaknessId field if non-nil, zero value otherwise.

### GetWeaknessIdOk

`func (o *ApiCWE) GetWeaknessIdOk() (*string, bool)`

GetWeaknessIdOk returns a tuple with the WeaknessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaknessId

`func (o *ApiCWE) SetWeaknessId(v string)`

SetWeaknessId sets WeaknessId field to given value.

### HasWeaknessId

`func (o *ApiCWE) HasWeaknessId() bool`

HasWeaknessId returns a boolean if a field has been set.

### GetWeaknessName

`func (o *ApiCWE) GetWeaknessName() string`

GetWeaknessName returns the WeaknessName field if non-nil, zero value otherwise.

### GetWeaknessNameOk

`func (o *ApiCWE) GetWeaknessNameOk() (*string, bool)`

GetWeaknessNameOk returns a tuple with the WeaknessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaknessName

`func (o *ApiCWE) SetWeaknessName(v string)`

SetWeaknessName sets WeaknessName field to given value.

### HasWeaknessName

`func (o *ApiCWE) HasWeaknessName() bool`

HasWeaknessName returns a boolean if a field has been set.

### GetWeightedScore

`func (o *ApiCWE) GetWeightedScore() float32`

GetWeightedScore returns the WeightedScore field if non-nil, zero value otherwise.

### GetWeightedScoreOk

`func (o *ApiCWE) GetWeightedScoreOk() (*float32, bool)`

GetWeightedScoreOk returns a tuple with the WeightedScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedScore

`func (o *ApiCWE) SetWeightedScore(v float32)`

SetWeightedScore sets WeightedScore field to given value.

### HasWeightedScore

`func (o *ApiCWE) HasWeightedScore() bool`

HasWeightedScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


