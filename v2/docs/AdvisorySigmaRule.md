# AdvisorySigmaRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MitreAttackTechniques** | Pointer to **[]string** |  | [optional] 
**SigmaRule** | Pointer to [**AdvisorySigmaRuleRule**](AdvisorySigmaRuleRule.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySigmaRule

`func NewAdvisorySigmaRule() *AdvisorySigmaRule`

NewAdvisorySigmaRule instantiates a new AdvisorySigmaRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySigmaRuleWithDefaults

`func NewAdvisorySigmaRuleWithDefaults() *AdvisorySigmaRule`

NewAdvisorySigmaRuleWithDefaults instantiates a new AdvisorySigmaRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisorySigmaRule) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySigmaRule) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySigmaRule) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySigmaRule) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySigmaRule) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySigmaRule) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySigmaRule) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySigmaRule) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMitreAttackTechniques

`func (o *AdvisorySigmaRule) GetMitreAttackTechniques() []string`

GetMitreAttackTechniques returns the MitreAttackTechniques field if non-nil, zero value otherwise.

### GetMitreAttackTechniquesOk

`func (o *AdvisorySigmaRule) GetMitreAttackTechniquesOk() (*[]string, bool)`

GetMitreAttackTechniquesOk returns a tuple with the MitreAttackTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreAttackTechniques

`func (o *AdvisorySigmaRule) SetMitreAttackTechniques(v []string)`

SetMitreAttackTechniques sets MitreAttackTechniques field to given value.

### HasMitreAttackTechniques

`func (o *AdvisorySigmaRule) HasMitreAttackTechniques() bool`

HasMitreAttackTechniques returns a boolean if a field has been set.

### GetSigmaRule

`func (o *AdvisorySigmaRule) GetSigmaRule() AdvisorySigmaRuleRule`

GetSigmaRule returns the SigmaRule field if non-nil, zero value otherwise.

### GetSigmaRuleOk

`func (o *AdvisorySigmaRule) GetSigmaRuleOk() (*AdvisorySigmaRuleRule, bool)`

GetSigmaRuleOk returns a tuple with the SigmaRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigmaRule

`func (o *AdvisorySigmaRule) SetSigmaRule(v AdvisorySigmaRuleRule)`

SetSigmaRule sets SigmaRule field to given value.

### HasSigmaRule

`func (o *AdvisorySigmaRule) HasSigmaRule() bool`

HasSigmaRule returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisorySigmaRule) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisorySigmaRule) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisorySigmaRule) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisorySigmaRule) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySigmaRule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySigmaRule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySigmaRule) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySigmaRule) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisorySigmaRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisorySigmaRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisorySigmaRule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisorySigmaRule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


