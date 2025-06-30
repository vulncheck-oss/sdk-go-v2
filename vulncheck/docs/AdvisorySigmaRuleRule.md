# AdvisorySigmaRuleRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Detection** | Pointer to **map[string]interface{}** |  | [optional] 
**FalsePositives** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Logsource** | Pointer to [**AdvisoryLogSource**](AdvisoryLogSource.md) |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Related** | Pointer to [**[]AdvisoryRelatedRule**](AdvisoryRelatedRule.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySigmaRuleRule

`func NewAdvisorySigmaRuleRule() *AdvisorySigmaRuleRule`

NewAdvisorySigmaRuleRule instantiates a new AdvisorySigmaRuleRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySigmaRuleRuleWithDefaults

`func NewAdvisorySigmaRuleRuleWithDefaults() *AdvisorySigmaRuleRule`

NewAdvisorySigmaRuleRuleWithDefaults instantiates a new AdvisorySigmaRuleRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AdvisorySigmaRuleRule) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AdvisorySigmaRuleRule) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AdvisorySigmaRuleRule) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AdvisorySigmaRuleRule) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDate

`func (o *AdvisorySigmaRuleRule) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisorySigmaRuleRule) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisorySigmaRuleRule) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisorySigmaRuleRule) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisorySigmaRuleRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisorySigmaRuleRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisorySigmaRuleRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisorySigmaRuleRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetection

`func (o *AdvisorySigmaRuleRule) GetDetection() map[string]interface{}`

GetDetection returns the Detection field if non-nil, zero value otherwise.

### GetDetectionOk

`func (o *AdvisorySigmaRuleRule) GetDetectionOk() (*map[string]interface{}, bool)`

GetDetectionOk returns a tuple with the Detection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetection

`func (o *AdvisorySigmaRuleRule) SetDetection(v map[string]interface{})`

SetDetection sets Detection field to given value.

### HasDetection

`func (o *AdvisorySigmaRuleRule) HasDetection() bool`

HasDetection returns a boolean if a field has been set.

### GetFalsePositives

`func (o *AdvisorySigmaRuleRule) GetFalsePositives() []string`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *AdvisorySigmaRuleRule) GetFalsePositivesOk() (*[]string, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *AdvisorySigmaRuleRule) SetFalsePositives(v []string)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *AdvisorySigmaRuleRule) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetFields

`func (o *AdvisorySigmaRuleRule) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AdvisorySigmaRuleRule) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AdvisorySigmaRuleRule) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AdvisorySigmaRuleRule) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *AdvisorySigmaRuleRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisorySigmaRuleRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisorySigmaRuleRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisorySigmaRuleRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *AdvisorySigmaRuleRule) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AdvisorySigmaRuleRule) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AdvisorySigmaRuleRule) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AdvisorySigmaRuleRule) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLogsource

`func (o *AdvisorySigmaRuleRule) GetLogsource() AdvisoryLogSource`

GetLogsource returns the Logsource field if non-nil, zero value otherwise.

### GetLogsourceOk

`func (o *AdvisorySigmaRuleRule) GetLogsourceOk() (*AdvisoryLogSource, bool)`

GetLogsourceOk returns a tuple with the Logsource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsource

`func (o *AdvisorySigmaRuleRule) SetLogsource(v AdvisoryLogSource)`

SetLogsource sets Logsource field to given value.

### HasLogsource

`func (o *AdvisorySigmaRuleRule) HasLogsource() bool`

HasLogsource returns a boolean if a field has been set.

### GetModified

`func (o *AdvisorySigmaRuleRule) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisorySigmaRuleRule) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisorySigmaRuleRule) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisorySigmaRuleRule) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisorySigmaRuleRule) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisorySigmaRuleRule) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisorySigmaRuleRule) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisorySigmaRuleRule) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelated

`func (o *AdvisorySigmaRuleRule) GetRelated() []AdvisoryRelatedRule`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *AdvisorySigmaRuleRule) GetRelatedOk() (*[]AdvisoryRelatedRule, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *AdvisorySigmaRuleRule) SetRelated(v []AdvisoryRelatedRule)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *AdvisorySigmaRuleRule) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisorySigmaRuleRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisorySigmaRuleRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisorySigmaRuleRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisorySigmaRuleRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *AdvisorySigmaRuleRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdvisorySigmaRuleRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdvisorySigmaRuleRule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdvisorySigmaRuleRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySigmaRuleRule) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySigmaRuleRule) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySigmaRuleRule) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySigmaRuleRule) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


