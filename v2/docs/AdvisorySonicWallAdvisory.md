# AdvisorySonicWallAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryId** | Pointer to **string** |  | [optional] 
**AffectedProducts** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**CvssVersion** | Pointer to **float32** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**IsWorkaroundAvailable** | Pointer to **bool** |  | [optional] 
**LastUpdatedWhen** | Pointer to **string** |  | [optional] 
**PublishedWhen** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VulnStatus** | Pointer to **string** |  | [optional] 
**VulnerableProductsList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisorySonicWallAdvisory

`func NewAdvisorySonicWallAdvisory() *AdvisorySonicWallAdvisory`

NewAdvisorySonicWallAdvisory instantiates a new AdvisorySonicWallAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySonicWallAdvisoryWithDefaults

`func NewAdvisorySonicWallAdvisoryWithDefaults() *AdvisorySonicWallAdvisory`

NewAdvisorySonicWallAdvisoryWithDefaults instantiates a new AdvisorySonicWallAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryId

`func (o *AdvisorySonicWallAdvisory) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *AdvisorySonicWallAdvisory) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *AdvisorySonicWallAdvisory) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *AdvisorySonicWallAdvisory) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetAffectedProducts

`func (o *AdvisorySonicWallAdvisory) GetAffectedProducts() []string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisorySonicWallAdvisory) GetAffectedProductsOk() (*[]string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisorySonicWallAdvisory) SetAffectedProducts(v []string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisorySonicWallAdvisory) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySonicWallAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySonicWallAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySonicWallAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySonicWallAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisorySonicWallAdvisory) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisorySonicWallAdvisory) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisorySonicWallAdvisory) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisorySonicWallAdvisory) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisorySonicWallAdvisory) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisorySonicWallAdvisory) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisorySonicWallAdvisory) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisorySonicWallAdvisory) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCvssVersion

`func (o *AdvisorySonicWallAdvisory) GetCvssVersion() float32`

GetCvssVersion returns the CvssVersion field if non-nil, zero value otherwise.

### GetCvssVersionOk

`func (o *AdvisorySonicWallAdvisory) GetCvssVersionOk() (*float32, bool)`

GetCvssVersionOk returns a tuple with the CvssVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVersion

`func (o *AdvisorySonicWallAdvisory) SetCvssVersion(v float32)`

SetCvssVersion sets CvssVersion field to given value.

### HasCvssVersion

`func (o *AdvisorySonicWallAdvisory) HasCvssVersion() bool`

HasCvssVersion returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisorySonicWallAdvisory) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisorySonicWallAdvisory) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisorySonicWallAdvisory) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisorySonicWallAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySonicWallAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySonicWallAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySonicWallAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySonicWallAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisorySonicWallAdvisory) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisorySonicWallAdvisory) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisorySonicWallAdvisory) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisorySonicWallAdvisory) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetIsWorkaroundAvailable

`func (o *AdvisorySonicWallAdvisory) GetIsWorkaroundAvailable() bool`

GetIsWorkaroundAvailable returns the IsWorkaroundAvailable field if non-nil, zero value otherwise.

### GetIsWorkaroundAvailableOk

`func (o *AdvisorySonicWallAdvisory) GetIsWorkaroundAvailableOk() (*bool, bool)`

GetIsWorkaroundAvailableOk returns a tuple with the IsWorkaroundAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWorkaroundAvailable

`func (o *AdvisorySonicWallAdvisory) SetIsWorkaroundAvailable(v bool)`

SetIsWorkaroundAvailable sets IsWorkaroundAvailable field to given value.

### HasIsWorkaroundAvailable

`func (o *AdvisorySonicWallAdvisory) HasIsWorkaroundAvailable() bool`

HasIsWorkaroundAvailable returns a boolean if a field has been set.

### GetLastUpdatedWhen

`func (o *AdvisorySonicWallAdvisory) GetLastUpdatedWhen() string`

GetLastUpdatedWhen returns the LastUpdatedWhen field if non-nil, zero value otherwise.

### GetLastUpdatedWhenOk

`func (o *AdvisorySonicWallAdvisory) GetLastUpdatedWhenOk() (*string, bool)`

GetLastUpdatedWhenOk returns a tuple with the LastUpdatedWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedWhen

`func (o *AdvisorySonicWallAdvisory) SetLastUpdatedWhen(v string)`

SetLastUpdatedWhen sets LastUpdatedWhen field to given value.

### HasLastUpdatedWhen

`func (o *AdvisorySonicWallAdvisory) HasLastUpdatedWhen() bool`

HasLastUpdatedWhen returns a boolean if a field has been set.

### GetPublishedWhen

`func (o *AdvisorySonicWallAdvisory) GetPublishedWhen() string`

GetPublishedWhen returns the PublishedWhen field if non-nil, zero value otherwise.

### GetPublishedWhenOk

`func (o *AdvisorySonicWallAdvisory) GetPublishedWhenOk() (*string, bool)`

GetPublishedWhenOk returns a tuple with the PublishedWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedWhen

`func (o *AdvisorySonicWallAdvisory) SetPublishedWhen(v string)`

SetPublishedWhen sets PublishedWhen field to given value.

### HasPublishedWhen

`func (o *AdvisorySonicWallAdvisory) HasPublishedWhen() bool`

HasPublishedWhen returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisorySonicWallAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisorySonicWallAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisorySonicWallAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisorySonicWallAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySonicWallAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySonicWallAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySonicWallAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySonicWallAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisorySonicWallAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisorySonicWallAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisorySonicWallAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisorySonicWallAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnStatus

`func (o *AdvisorySonicWallAdvisory) GetVulnStatus() string`

GetVulnStatus returns the VulnStatus field if non-nil, zero value otherwise.

### GetVulnStatusOk

`func (o *AdvisorySonicWallAdvisory) GetVulnStatusOk() (*string, bool)`

GetVulnStatusOk returns a tuple with the VulnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnStatus

`func (o *AdvisorySonicWallAdvisory) SetVulnStatus(v string)`

SetVulnStatus sets VulnStatus field to given value.

### HasVulnStatus

`func (o *AdvisorySonicWallAdvisory) HasVulnStatus() bool`

HasVulnStatus returns a boolean if a field has been set.

### GetVulnerableProductsList

`func (o *AdvisorySonicWallAdvisory) GetVulnerableProductsList() []string`

GetVulnerableProductsList returns the VulnerableProductsList field if non-nil, zero value otherwise.

### GetVulnerableProductsListOk

`func (o *AdvisorySonicWallAdvisory) GetVulnerableProductsListOk() (*[]string, bool)`

GetVulnerableProductsListOk returns a tuple with the VulnerableProductsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableProductsList

`func (o *AdvisorySonicWallAdvisory) SetVulnerableProductsList(v []string)`

SetVulnerableProductsList sets VulnerableProductsList field to given value.

### HasVulnerableProductsList

`func (o *AdvisorySonicWallAdvisory) HasVulnerableProductsList() bool`

HasVulnerableProductsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


