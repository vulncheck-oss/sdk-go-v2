# AdvisorySynology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySynology

`func NewAdvisorySynology() *AdvisorySynology`

NewAdvisorySynology instantiates a new AdvisorySynology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySynologyWithDefaults

`func NewAdvisorySynologyWithDefaults() *AdvisorySynology`

NewAdvisorySynologyWithDefaults instantiates a new AdvisorySynology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisorySynology) GetAffectedProducts() []string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisorySynology) GetAffectedProductsOk() (*[]string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisorySynology) SetAffectedProducts(v []string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisorySynology) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySynology) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySynology) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySynology) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySynology) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySynology) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySynology) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySynology) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySynology) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisorySynology) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisorySynology) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisorySynology) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisorySynology) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AdvisorySynology) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AdvisorySynology) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AdvisorySynology) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AdvisorySynology) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisorySynology) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisorySynology) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisorySynology) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisorySynology) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisorySynology) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisorySynology) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisorySynology) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisorySynology) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisorySynology) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisorySynology) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisorySynology) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisorySynology) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisorySynology) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisorySynology) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisorySynology) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisorySynology) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySynology) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySynology) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySynology) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySynology) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisorySynology) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisorySynology) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisorySynology) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisorySynology) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


