# AdvisoryJVN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedEn** | Pointer to **string** |  | [optional] 
**AffectedJa** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DescriptionEn** | Pointer to **string** |  | [optional] 
**DescriptionJa** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**SolutionEn** | Pointer to **string** |  | [optional] 
**SolutionJa** | Pointer to **string** |  | [optional] 
**SummaryEn** | Pointer to **string** |  | [optional] 
**SummaryJa** | Pointer to **string** |  | [optional] 
**TitleEn** | Pointer to **string** |  | [optional] 
**TitleJa** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryJVN

`func NewAdvisoryJVN() *AdvisoryJVN`

NewAdvisoryJVN instantiates a new AdvisoryJVN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryJVNWithDefaults

`func NewAdvisoryJVNWithDefaults() *AdvisoryJVN`

NewAdvisoryJVNWithDefaults instantiates a new AdvisoryJVN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedEn

`func (o *AdvisoryJVN) GetAffectedEn() string`

GetAffectedEn returns the AffectedEn field if non-nil, zero value otherwise.

### GetAffectedEnOk

`func (o *AdvisoryJVN) GetAffectedEnOk() (*string, bool)`

GetAffectedEnOk returns a tuple with the AffectedEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEn

`func (o *AdvisoryJVN) SetAffectedEn(v string)`

SetAffectedEn sets AffectedEn field to given value.

### HasAffectedEn

`func (o *AdvisoryJVN) HasAffectedEn() bool`

HasAffectedEn returns a boolean if a field has been set.

### GetAffectedJa

`func (o *AdvisoryJVN) GetAffectedJa() string`

GetAffectedJa returns the AffectedJa field if non-nil, zero value otherwise.

### GetAffectedJaOk

`func (o *AdvisoryJVN) GetAffectedJaOk() (*string, bool)`

GetAffectedJaOk returns a tuple with the AffectedJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedJa

`func (o *AdvisoryJVN) SetAffectedJa(v string)`

SetAffectedJa sets AffectedJa field to given value.

### HasAffectedJa

`func (o *AdvisoryJVN) HasAffectedJa() bool`

HasAffectedJa returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryJVN) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryJVN) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryJVN) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryJVN) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryJVN) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryJVN) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryJVN) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryJVN) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescriptionEn

`func (o *AdvisoryJVN) GetDescriptionEn() string`

GetDescriptionEn returns the DescriptionEn field if non-nil, zero value otherwise.

### GetDescriptionEnOk

`func (o *AdvisoryJVN) GetDescriptionEnOk() (*string, bool)`

GetDescriptionEnOk returns a tuple with the DescriptionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionEn

`func (o *AdvisoryJVN) SetDescriptionEn(v string)`

SetDescriptionEn sets DescriptionEn field to given value.

### HasDescriptionEn

`func (o *AdvisoryJVN) HasDescriptionEn() bool`

HasDescriptionEn returns a boolean if a field has been set.

### GetDescriptionJa

`func (o *AdvisoryJVN) GetDescriptionJa() string`

GetDescriptionJa returns the DescriptionJa field if non-nil, zero value otherwise.

### GetDescriptionJaOk

`func (o *AdvisoryJVN) GetDescriptionJaOk() (*string, bool)`

GetDescriptionJaOk returns a tuple with the DescriptionJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionJa

`func (o *AdvisoryJVN) SetDescriptionJa(v string)`

SetDescriptionJa sets DescriptionJa field to given value.

### HasDescriptionJa

`func (o *AdvisoryJVN) HasDescriptionJa() bool`

HasDescriptionJa returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryJVN) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryJVN) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryJVN) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryJVN) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryJVN) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryJVN) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryJVN) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryJVN) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSolutionEn

`func (o *AdvisoryJVN) GetSolutionEn() string`

GetSolutionEn returns the SolutionEn field if non-nil, zero value otherwise.

### GetSolutionEnOk

`func (o *AdvisoryJVN) GetSolutionEnOk() (*string, bool)`

GetSolutionEnOk returns a tuple with the SolutionEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionEn

`func (o *AdvisoryJVN) SetSolutionEn(v string)`

SetSolutionEn sets SolutionEn field to given value.

### HasSolutionEn

`func (o *AdvisoryJVN) HasSolutionEn() bool`

HasSolutionEn returns a boolean if a field has been set.

### GetSolutionJa

`func (o *AdvisoryJVN) GetSolutionJa() string`

GetSolutionJa returns the SolutionJa field if non-nil, zero value otherwise.

### GetSolutionJaOk

`func (o *AdvisoryJVN) GetSolutionJaOk() (*string, bool)`

GetSolutionJaOk returns a tuple with the SolutionJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionJa

`func (o *AdvisoryJVN) SetSolutionJa(v string)`

SetSolutionJa sets SolutionJa field to given value.

### HasSolutionJa

`func (o *AdvisoryJVN) HasSolutionJa() bool`

HasSolutionJa returns a boolean if a field has been set.

### GetSummaryEn

`func (o *AdvisoryJVN) GetSummaryEn() string`

GetSummaryEn returns the SummaryEn field if non-nil, zero value otherwise.

### GetSummaryEnOk

`func (o *AdvisoryJVN) GetSummaryEnOk() (*string, bool)`

GetSummaryEnOk returns a tuple with the SummaryEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryEn

`func (o *AdvisoryJVN) SetSummaryEn(v string)`

SetSummaryEn sets SummaryEn field to given value.

### HasSummaryEn

`func (o *AdvisoryJVN) HasSummaryEn() bool`

HasSummaryEn returns a boolean if a field has been set.

### GetSummaryJa

`func (o *AdvisoryJVN) GetSummaryJa() string`

GetSummaryJa returns the SummaryJa field if non-nil, zero value otherwise.

### GetSummaryJaOk

`func (o *AdvisoryJVN) GetSummaryJaOk() (*string, bool)`

GetSummaryJaOk returns a tuple with the SummaryJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryJa

`func (o *AdvisoryJVN) SetSummaryJa(v string)`

SetSummaryJa sets SummaryJa field to given value.

### HasSummaryJa

`func (o *AdvisoryJVN) HasSummaryJa() bool`

HasSummaryJa returns a boolean if a field has been set.

### GetTitleEn

`func (o *AdvisoryJVN) GetTitleEn() string`

GetTitleEn returns the TitleEn field if non-nil, zero value otherwise.

### GetTitleEnOk

`func (o *AdvisoryJVN) GetTitleEnOk() (*string, bool)`

GetTitleEnOk returns a tuple with the TitleEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEn

`func (o *AdvisoryJVN) SetTitleEn(v string)`

SetTitleEn sets TitleEn field to given value.

### HasTitleEn

`func (o *AdvisoryJVN) HasTitleEn() bool`

HasTitleEn returns a boolean if a field has been set.

### GetTitleJa

`func (o *AdvisoryJVN) GetTitleJa() string`

GetTitleJa returns the TitleJa field if non-nil, zero value otherwise.

### GetTitleJaOk

`func (o *AdvisoryJVN) GetTitleJaOk() (*string, bool)`

GetTitleJaOk returns a tuple with the TitleJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleJa

`func (o *AdvisoryJVN) SetTitleJa(v string)`

SetTitleJa sets TitleJa field to given value.

### HasTitleJa

`func (o *AdvisoryJVN) HasTitleJa() bool`

HasTitleJa returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryJVN) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryJVN) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryJVN) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryJVN) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryJVN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryJVN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryJVN) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryJVN) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


