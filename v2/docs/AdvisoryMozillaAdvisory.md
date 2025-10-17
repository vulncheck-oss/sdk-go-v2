# AdvisoryMozillaAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedComponents** | Pointer to [**[]AdvisoryMozillaComponent**](AdvisoryMozillaComponent.md) |  | [optional] 
**Bugzilla** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FixedIn** | Pointer to **[]string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**Reporter** | Pointer to **string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMozillaAdvisory

`func NewAdvisoryMozillaAdvisory() *AdvisoryMozillaAdvisory`

NewAdvisoryMozillaAdvisory instantiates a new AdvisoryMozillaAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMozillaAdvisoryWithDefaults

`func NewAdvisoryMozillaAdvisoryWithDefaults() *AdvisoryMozillaAdvisory`

NewAdvisoryMozillaAdvisoryWithDefaults instantiates a new AdvisoryMozillaAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedComponents

`func (o *AdvisoryMozillaAdvisory) GetAffectedComponents() []AdvisoryMozillaComponent`

GetAffectedComponents returns the AffectedComponents field if non-nil, zero value otherwise.

### GetAffectedComponentsOk

`func (o *AdvisoryMozillaAdvisory) GetAffectedComponentsOk() (*[]AdvisoryMozillaComponent, bool)`

GetAffectedComponentsOk returns a tuple with the AffectedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponents

`func (o *AdvisoryMozillaAdvisory) SetAffectedComponents(v []AdvisoryMozillaComponent)`

SetAffectedComponents sets AffectedComponents field to given value.

### HasAffectedComponents

`func (o *AdvisoryMozillaAdvisory) HasAffectedComponents() bool`

HasAffectedComponents returns a boolean if a field has been set.

### GetBugzilla

`func (o *AdvisoryMozillaAdvisory) GetBugzilla() []string`

GetBugzilla returns the Bugzilla field if non-nil, zero value otherwise.

### GetBugzillaOk

`func (o *AdvisoryMozillaAdvisory) GetBugzillaOk() (*[]string, bool)`

GetBugzillaOk returns a tuple with the Bugzilla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzilla

`func (o *AdvisoryMozillaAdvisory) SetBugzilla(v []string)`

SetBugzilla sets Bugzilla field to given value.

### HasBugzilla

`func (o *AdvisoryMozillaAdvisory) HasBugzilla() bool`

HasBugzilla returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryMozillaAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMozillaAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMozillaAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMozillaAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMozillaAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMozillaAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMozillaAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMozillaAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryMozillaAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMozillaAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMozillaAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMozillaAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedIn

`func (o *AdvisoryMozillaAdvisory) GetFixedIn() []string`

GetFixedIn returns the FixedIn field if non-nil, zero value otherwise.

### GetFixedInOk

`func (o *AdvisoryMozillaAdvisory) GetFixedInOk() (*[]string, bool)`

GetFixedInOk returns a tuple with the FixedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIn

`func (o *AdvisoryMozillaAdvisory) SetFixedIn(v []string)`

SetFixedIn sets FixedIn field to given value.

### HasFixedIn

`func (o *AdvisoryMozillaAdvisory) HasFixedIn() bool`

HasFixedIn returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisoryMozillaAdvisory) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisoryMozillaAdvisory) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisoryMozillaAdvisory) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisoryMozillaAdvisory) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisoryMozillaAdvisory) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisoryMozillaAdvisory) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisoryMozillaAdvisory) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisoryMozillaAdvisory) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetReporter

`func (o *AdvisoryMozillaAdvisory) GetReporter() string`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *AdvisoryMozillaAdvisory) GetReporterOk() (*string, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *AdvisoryMozillaAdvisory) SetReporter(v string)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *AdvisoryMozillaAdvisory) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetRisk

`func (o *AdvisoryMozillaAdvisory) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AdvisoryMozillaAdvisory) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AdvisoryMozillaAdvisory) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AdvisoryMozillaAdvisory) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMozillaAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMozillaAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMozillaAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMozillaAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMozillaAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMozillaAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMozillaAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMozillaAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


