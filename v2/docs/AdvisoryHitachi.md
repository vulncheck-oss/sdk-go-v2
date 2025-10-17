# AdvisoryHitachi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**FixedProducts** | Pointer to **string** |  | [optional] 
**HitachiId** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHitachi

`func NewAdvisoryHitachi() *AdvisoryHitachi`

NewAdvisoryHitachi instantiates a new AdvisoryHitachi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHitachiWithDefaults

`func NewAdvisoryHitachiWithDefaults() *AdvisoryHitachi`

NewAdvisoryHitachiWithDefaults instantiates a new AdvisoryHitachi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryHitachi) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryHitachi) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryHitachi) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryHitachi) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHitachi) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHitachi) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHitachi) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHitachi) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHitachi) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHitachi) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHitachi) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHitachi) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixedProducts

`func (o *AdvisoryHitachi) GetFixedProducts() string`

GetFixedProducts returns the FixedProducts field if non-nil, zero value otherwise.

### GetFixedProductsOk

`func (o *AdvisoryHitachi) GetFixedProductsOk() (*string, bool)`

GetFixedProductsOk returns a tuple with the FixedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedProducts

`func (o *AdvisoryHitachi) SetFixedProducts(v string)`

SetFixedProducts sets FixedProducts field to given value.

### HasFixedProducts

`func (o *AdvisoryHitachi) HasFixedProducts() bool`

HasFixedProducts returns a boolean if a field has been set.

### GetHitachiId

`func (o *AdvisoryHitachi) GetHitachiId() string`

GetHitachiId returns the HitachiId field if non-nil, zero value otherwise.

### GetHitachiIdOk

`func (o *AdvisoryHitachi) GetHitachiIdOk() (*string, bool)`

GetHitachiIdOk returns a tuple with the HitachiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitachiId

`func (o *AdvisoryHitachi) SetHitachiId(v string)`

SetHitachiId sets HitachiId field to given value.

### HasHitachiId

`func (o *AdvisoryHitachi) HasHitachiId() bool`

HasHitachiId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryHitachi) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryHitachi) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryHitachi) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryHitachi) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryHitachi) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryHitachi) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryHitachi) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryHitachi) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryHitachi) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryHitachi) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryHitachi) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryHitachi) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryHitachi) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryHitachi) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryHitachi) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryHitachi) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


