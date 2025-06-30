# AdvisoryCommVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**CveDetails** | Pointer to [**[]AdvisoryCommVaultCVEDetails**](AdvisoryCommVaultCVEDetails.md) |  | [optional] 
**CvssRange** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImpactedProduct** | Pointer to [**AdvisoryCommVaultImpactedProduct**](AdvisoryCommVaultImpactedProduct.md) |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Resolution** | Pointer to [**AdvisoryCommVaultResolution**](AdvisoryCommVaultResolution.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCommVault

`func NewAdvisoryCommVault() *AdvisoryCommVault`

NewAdvisoryCommVault instantiates a new AdvisoryCommVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCommVaultWithDefaults

`func NewAdvisoryCommVaultWithDefaults() *AdvisoryCommVault`

NewAdvisoryCommVaultWithDefaults instantiates a new AdvisoryCommVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryCommVault) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCommVault) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCommVault) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCommVault) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveDetails

`func (o *AdvisoryCommVault) GetCveDetails() []AdvisoryCommVaultCVEDetails`

GetCveDetails returns the CveDetails field if non-nil, zero value otherwise.

### GetCveDetailsOk

`func (o *AdvisoryCommVault) GetCveDetailsOk() (*[]AdvisoryCommVaultCVEDetails, bool)`

GetCveDetailsOk returns a tuple with the CveDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveDetails

`func (o *AdvisoryCommVault) SetCveDetails(v []AdvisoryCommVaultCVEDetails)`

SetCveDetails sets CveDetails field to given value.

### HasCveDetails

`func (o *AdvisoryCommVault) HasCveDetails() bool`

HasCveDetails returns a boolean if a field has been set.

### GetCvssRange

`func (o *AdvisoryCommVault) GetCvssRange() string`

GetCvssRange returns the CvssRange field if non-nil, zero value otherwise.

### GetCvssRangeOk

`func (o *AdvisoryCommVault) GetCvssRangeOk() (*string, bool)`

GetCvssRangeOk returns a tuple with the CvssRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssRange

`func (o *AdvisoryCommVault) SetCvssRange(v string)`

SetCvssRange sets CvssRange field to given value.

### HasCvssRange

`func (o *AdvisoryCommVault) HasCvssRange() bool`

HasCvssRange returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCommVault) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCommVault) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCommVault) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCommVault) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryCommVault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryCommVault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryCommVault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryCommVault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryCommVault) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryCommVault) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryCommVault) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryCommVault) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpactedProduct

`func (o *AdvisoryCommVault) GetImpactedProduct() AdvisoryCommVaultImpactedProduct`

GetImpactedProduct returns the ImpactedProduct field if non-nil, zero value otherwise.

### GetImpactedProductOk

`func (o *AdvisoryCommVault) GetImpactedProductOk() (*AdvisoryCommVaultImpactedProduct, bool)`

GetImpactedProductOk returns a tuple with the ImpactedProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedProduct

`func (o *AdvisoryCommVault) SetImpactedProduct(v AdvisoryCommVaultImpactedProduct)`

SetImpactedProduct sets ImpactedProduct field to given value.

### HasImpactedProduct

`func (o *AdvisoryCommVault) HasImpactedProduct() bool`

HasImpactedProduct returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCommVault) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCommVault) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCommVault) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCommVault) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetResolution

`func (o *AdvisoryCommVault) GetResolution() AdvisoryCommVaultResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AdvisoryCommVault) GetResolutionOk() (*AdvisoryCommVaultResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AdvisoryCommVault) SetResolution(v AdvisoryCommVaultResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AdvisoryCommVault) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryCommVault) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryCommVault) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryCommVault) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryCommVault) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCommVault) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCommVault) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCommVault) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCommVault) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryCommVault) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryCommVault) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryCommVault) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryCommVault) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCommVault) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCommVault) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCommVault) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCommVault) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


