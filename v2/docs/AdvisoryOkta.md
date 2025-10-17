# AdvisoryOkta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOkta

`func NewAdvisoryOkta() *AdvisoryOkta`

NewAdvisoryOkta instantiates a new AdvisoryOkta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOktaWithDefaults

`func NewAdvisoryOktaWithDefaults() *AdvisoryOkta`

NewAdvisoryOktaWithDefaults instantiates a new AdvisoryOkta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryOkta) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryOkta) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryOkta) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryOkta) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryOkta) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOkta) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOkta) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOkta) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryOkta) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryOkta) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryOkta) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryOkta) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryOkta) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryOkta) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryOkta) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryOkta) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryOkta) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryOkta) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryOkta) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryOkta) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryOkta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryOkta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryOkta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryOkta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOkta) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOkta) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOkta) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOkta) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetResolution

`func (o *AdvisoryOkta) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AdvisoryOkta) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AdvisoryOkta) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AdvisoryOkta) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOkta) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOkta) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOkta) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOkta) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryOkta) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryOkta) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryOkta) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryOkta) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


