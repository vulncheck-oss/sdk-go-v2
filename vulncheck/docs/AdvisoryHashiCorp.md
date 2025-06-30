# AdvisoryHashiCorp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**Background** | Pointer to **string** |  | [optional] 
**BulletinId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Remediation** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryHashiCorp

`func NewAdvisoryHashiCorp() *AdvisoryHashiCorp`

NewAdvisoryHashiCorp instantiates a new AdvisoryHashiCorp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHashiCorpWithDefaults

`func NewAdvisoryHashiCorpWithDefaults() *AdvisoryHashiCorp`

NewAdvisoryHashiCorpWithDefaults instantiates a new AdvisoryHashiCorp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryHashiCorp) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryHashiCorp) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryHashiCorp) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryHashiCorp) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetBackground

`func (o *AdvisoryHashiCorp) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *AdvisoryHashiCorp) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *AdvisoryHashiCorp) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *AdvisoryHashiCorp) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBulletinId

`func (o *AdvisoryHashiCorp) GetBulletinId() string`

GetBulletinId returns the BulletinId field if non-nil, zero value otherwise.

### GetBulletinIdOk

`func (o *AdvisoryHashiCorp) GetBulletinIdOk() (*string, bool)`

GetBulletinIdOk returns a tuple with the BulletinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinId

`func (o *AdvisoryHashiCorp) SetBulletinId(v string)`

SetBulletinId sets BulletinId field to given value.

### HasBulletinId

`func (o *AdvisoryHashiCorp) HasBulletinId() bool`

HasBulletinId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHashiCorp) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHashiCorp) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHashiCorp) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHashiCorp) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHashiCorp) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHashiCorp) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHashiCorp) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHashiCorp) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryHashiCorp) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryHashiCorp) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryHashiCorp) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryHashiCorp) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRemediation

`func (o *AdvisoryHashiCorp) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *AdvisoryHashiCorp) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *AdvisoryHashiCorp) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *AdvisoryHashiCorp) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryHashiCorp) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryHashiCorp) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryHashiCorp) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryHashiCorp) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryHashiCorp) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryHashiCorp) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryHashiCorp) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryHashiCorp) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryHashiCorp) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryHashiCorp) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryHashiCorp) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryHashiCorp) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


