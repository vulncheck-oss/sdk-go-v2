# AdvisoryHKCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**RelatedLinks** | Pointer to **[]string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**Solutions** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHKCert

`func NewAdvisoryHKCert() *AdvisoryHKCert`

NewAdvisoryHKCert instantiates a new AdvisoryHKCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHKCertWithDefaults

`func NewAdvisoryHKCertWithDefaults() *AdvisoryHKCert`

NewAdvisoryHKCertWithDefaults instantiates a new AdvisoryHKCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryHKCert) GetAffected() []string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryHKCert) GetAffectedOk() (*[]string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryHKCert) SetAffected(v []string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryHKCert) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHKCert) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHKCert) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHKCert) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHKCert) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHKCert) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHKCert) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHKCert) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHKCert) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisoryHKCert) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisoryHKCert) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisoryHKCert) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisoryHKCert) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryHKCert) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryHKCert) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryHKCert) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryHKCert) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRelatedLinks

`func (o *AdvisoryHKCert) GetRelatedLinks() []string`

GetRelatedLinks returns the RelatedLinks field if non-nil, zero value otherwise.

### GetRelatedLinksOk

`func (o *AdvisoryHKCert) GetRelatedLinksOk() (*[]string, bool)`

GetRelatedLinksOk returns a tuple with the RelatedLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedLinks

`func (o *AdvisoryHKCert) SetRelatedLinks(v []string)`

SetRelatedLinks sets RelatedLinks field to given value.

### HasRelatedLinks

`func (o *AdvisoryHKCert) HasRelatedLinks() bool`

HasRelatedLinks returns a boolean if a field has been set.

### GetRisk

`func (o *AdvisoryHKCert) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *AdvisoryHKCert) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *AdvisoryHKCert) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *AdvisoryHKCert) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryHKCert) GetSolutions() string`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryHKCert) GetSolutionsOk() (*string, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryHKCert) SetSolutions(v string)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryHKCert) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryHKCert) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryHKCert) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryHKCert) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryHKCert) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryHKCert) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryHKCert) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryHKCert) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryHKCert) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryHKCert) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryHKCert) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryHKCert) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryHKCert) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


