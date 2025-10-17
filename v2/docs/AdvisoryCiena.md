# AdvisoryCiena

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**IssueNo** | Pointer to **int32** |  | [optional] 
**SecurityAdvisoryNumber** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VulnerableProducts** | Pointer to [**[]AdvisoryVulnerableProduct**](AdvisoryVulnerableProduct.md) |  | [optional] 

## Methods

### NewAdvisoryCiena

`func NewAdvisoryCiena() *AdvisoryCiena`

NewAdvisoryCiena instantiates a new AdvisoryCiena object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCienaWithDefaults

`func NewAdvisoryCienaWithDefaults() *AdvisoryCiena`

NewAdvisoryCienaWithDefaults instantiates a new AdvisoryCiena object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *AdvisoryCiena) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *AdvisoryCiena) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *AdvisoryCiena) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *AdvisoryCiena) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCiena) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCiena) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCiena) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCiena) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetIssueNo

`func (o *AdvisoryCiena) GetIssueNo() int32`

GetIssueNo returns the IssueNo field if non-nil, zero value otherwise.

### GetIssueNoOk

`func (o *AdvisoryCiena) GetIssueNoOk() (*int32, bool)`

GetIssueNoOk returns a tuple with the IssueNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueNo

`func (o *AdvisoryCiena) SetIssueNo(v int32)`

SetIssueNo sets IssueNo field to given value.

### HasIssueNo

`func (o *AdvisoryCiena) HasIssueNo() bool`

HasIssueNo returns a boolean if a field has been set.

### GetSecurityAdvisoryNumber

`func (o *AdvisoryCiena) GetSecurityAdvisoryNumber() string`

GetSecurityAdvisoryNumber returns the SecurityAdvisoryNumber field if non-nil, zero value otherwise.

### GetSecurityAdvisoryNumberOk

`func (o *AdvisoryCiena) GetSecurityAdvisoryNumberOk() (*string, bool)`

GetSecurityAdvisoryNumberOk returns a tuple with the SecurityAdvisoryNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAdvisoryNumber

`func (o *AdvisoryCiena) SetSecurityAdvisoryNumber(v string)`

SetSecurityAdvisoryNumber sets SecurityAdvisoryNumber field to given value.

### HasSecurityAdvisoryNumber

`func (o *AdvisoryCiena) HasSecurityAdvisoryNumber() bool`

HasSecurityAdvisoryNumber returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryCiena) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryCiena) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryCiena) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryCiena) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCiena) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCiena) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCiena) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCiena) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryCiena) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryCiena) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryCiena) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryCiena) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerableProducts

`func (o *AdvisoryCiena) GetVulnerableProducts() []AdvisoryVulnerableProduct`

GetVulnerableProducts returns the VulnerableProducts field if non-nil, zero value otherwise.

### GetVulnerableProductsOk

`func (o *AdvisoryCiena) GetVulnerableProductsOk() (*[]AdvisoryVulnerableProduct, bool)`

GetVulnerableProductsOk returns a tuple with the VulnerableProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableProducts

`func (o *AdvisoryCiena) SetVulnerableProducts(v []AdvisoryVulnerableProduct)`

SetVulnerableProducts sets VulnerableProducts field to given value.

### HasVulnerableProducts

`func (o *AdvisoryCiena) HasVulnerableProducts() bool`

HasVulnerableProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


