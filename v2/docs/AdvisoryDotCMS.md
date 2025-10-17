# AdvisoryDotCMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credit** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FixedVersion** | Pointer to **string** |  | [optional] 
**IssueId** | Pointer to **string** |  | [optional] 
**Mitigation** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDotCMS

`func NewAdvisoryDotCMS() *AdvisoryDotCMS`

NewAdvisoryDotCMS instantiates a new AdvisoryDotCMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDotCMSWithDefaults

`func NewAdvisoryDotCMSWithDefaults() *AdvisoryDotCMS`

NewAdvisoryDotCMSWithDefaults instantiates a new AdvisoryDotCMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredit

`func (o *AdvisoryDotCMS) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AdvisoryDotCMS) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AdvisoryDotCMS) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AdvisoryDotCMS) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDotCMS) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDotCMS) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDotCMS) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDotCMS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryDotCMS) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryDotCMS) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryDotCMS) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryDotCMS) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryDotCMS) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryDotCMS) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryDotCMS) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryDotCMS) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedVersion

`func (o *AdvisoryDotCMS) GetFixedVersion() string`

GetFixedVersion returns the FixedVersion field if non-nil, zero value otherwise.

### GetFixedVersionOk

`func (o *AdvisoryDotCMS) GetFixedVersionOk() (*string, bool)`

GetFixedVersionOk returns a tuple with the FixedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersion

`func (o *AdvisoryDotCMS) SetFixedVersion(v string)`

SetFixedVersion sets FixedVersion field to given value.

### HasFixedVersion

`func (o *AdvisoryDotCMS) HasFixedVersion() bool`

HasFixedVersion returns a boolean if a field has been set.

### GetIssueId

`func (o *AdvisoryDotCMS) GetIssueId() string`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *AdvisoryDotCMS) GetIssueIdOk() (*string, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *AdvisoryDotCMS) SetIssueId(v string)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *AdvisoryDotCMS) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetMitigation

`func (o *AdvisoryDotCMS) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *AdvisoryDotCMS) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *AdvisoryDotCMS) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *AdvisoryDotCMS) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryDotCMS) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryDotCMS) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryDotCMS) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryDotCMS) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryDotCMS) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryDotCMS) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryDotCMS) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryDotCMS) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDotCMS) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDotCMS) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDotCMS) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDotCMS) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryDotCMS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryDotCMS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryDotCMS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryDotCMS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


