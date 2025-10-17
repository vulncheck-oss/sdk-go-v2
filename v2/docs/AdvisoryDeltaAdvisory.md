# AdvisoryDeltaAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**RecommendedAction** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDeltaAdvisory

`func NewAdvisoryDeltaAdvisory() *AdvisoryDeltaAdvisory`

NewAdvisoryDeltaAdvisory instantiates a new AdvisoryDeltaAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDeltaAdvisoryWithDefaults

`func NewAdvisoryDeltaAdvisoryWithDefaults() *AdvisoryDeltaAdvisory`

NewAdvisoryDeltaAdvisoryWithDefaults instantiates a new AdvisoryDeltaAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryDeltaAdvisory) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryDeltaAdvisory) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryDeltaAdvisory) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryDeltaAdvisory) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDeltaAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDeltaAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDeltaAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDeltaAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryDeltaAdvisory) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryDeltaAdvisory) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryDeltaAdvisory) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryDeltaAdvisory) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryDeltaAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryDeltaAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryDeltaAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryDeltaAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryDeltaAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryDeltaAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryDeltaAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryDeltaAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryDeltaAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryDeltaAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryDeltaAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryDeltaAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRecommendedAction

`func (o *AdvisoryDeltaAdvisory) GetRecommendedAction() string`

GetRecommendedAction returns the RecommendedAction field if non-nil, zero value otherwise.

### GetRecommendedActionOk

`func (o *AdvisoryDeltaAdvisory) GetRecommendedActionOk() (*string, bool)`

GetRecommendedActionOk returns a tuple with the RecommendedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedAction

`func (o *AdvisoryDeltaAdvisory) SetRecommendedAction(v string)`

SetRecommendedAction sets RecommendedAction field to given value.

### HasRecommendedAction

`func (o *AdvisoryDeltaAdvisory) HasRecommendedAction() bool`

HasRecommendedAction returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDeltaAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDeltaAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDeltaAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDeltaAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


