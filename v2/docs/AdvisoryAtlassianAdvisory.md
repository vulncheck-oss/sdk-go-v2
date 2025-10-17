# AdvisoryAtlassianAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedVersion** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DetailedSummary** | Pointer to **string** | overloading in places with &#39;RiskAssessment&#39; and other places with &#39;Description&#39; | [optional] 
**FixedVersion** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAtlassianAdvisory

`func NewAdvisoryAtlassianAdvisory() *AdvisoryAtlassianAdvisory`

NewAdvisoryAtlassianAdvisory instantiates a new AdvisoryAtlassianAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAtlassianAdvisoryWithDefaults

`func NewAdvisoryAtlassianAdvisoryWithDefaults() *AdvisoryAtlassianAdvisory`

NewAdvisoryAtlassianAdvisoryWithDefaults instantiates a new AdvisoryAtlassianAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedVersion

`func (o *AdvisoryAtlassianAdvisory) GetAffectedVersion() []string`

GetAffectedVersion returns the AffectedVersion field if non-nil, zero value otherwise.

### GetAffectedVersionOk

`func (o *AdvisoryAtlassianAdvisory) GetAffectedVersionOk() (*[]string, bool)`

GetAffectedVersionOk returns a tuple with the AffectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersion

`func (o *AdvisoryAtlassianAdvisory) SetAffectedVersion(v []string)`

SetAffectedVersion sets AffectedVersion field to given value.

### HasAffectedVersion

`func (o *AdvisoryAtlassianAdvisory) HasAffectedVersion() bool`

HasAffectedVersion returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAtlassianAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAtlassianAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAtlassianAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAtlassianAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAtlassianAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAtlassianAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAtlassianAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAtlassianAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetailedSummary

`func (o *AdvisoryAtlassianAdvisory) GetDetailedSummary() string`

GetDetailedSummary returns the DetailedSummary field if non-nil, zero value otherwise.

### GetDetailedSummaryOk

`func (o *AdvisoryAtlassianAdvisory) GetDetailedSummaryOk() (*string, bool)`

GetDetailedSummaryOk returns a tuple with the DetailedSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedSummary

`func (o *AdvisoryAtlassianAdvisory) SetDetailedSummary(v string)`

SetDetailedSummary sets DetailedSummary field to given value.

### HasDetailedSummary

`func (o *AdvisoryAtlassianAdvisory) HasDetailedSummary() bool`

HasDetailedSummary returns a boolean if a field has been set.

### GetFixedVersion

`func (o *AdvisoryAtlassianAdvisory) GetFixedVersion() string`

GetFixedVersion returns the FixedVersion field if non-nil, zero value otherwise.

### GetFixedVersionOk

`func (o *AdvisoryAtlassianAdvisory) GetFixedVersionOk() (*string, bool)`

GetFixedVersionOk returns a tuple with the FixedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersion

`func (o *AdvisoryAtlassianAdvisory) SetFixedVersion(v string)`

SetFixedVersion sets FixedVersion field to given value.

### HasFixedVersion

`func (o *AdvisoryAtlassianAdvisory) HasFixedVersion() bool`

HasFixedVersion returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryAtlassianAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryAtlassianAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryAtlassianAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryAtlassianAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisoryAtlassianAdvisory) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisoryAtlassianAdvisory) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisoryAtlassianAdvisory) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisoryAtlassianAdvisory) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryAtlassianAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryAtlassianAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryAtlassianAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryAtlassianAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetReleaseDate

`func (o *AdvisoryAtlassianAdvisory) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AdvisoryAtlassianAdvisory) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AdvisoryAtlassianAdvisory) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AdvisoryAtlassianAdvisory) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryAtlassianAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryAtlassianAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryAtlassianAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryAtlassianAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryAtlassianAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryAtlassianAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryAtlassianAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryAtlassianAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryAtlassianAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryAtlassianAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryAtlassianAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryAtlassianAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryAtlassianAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryAtlassianAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryAtlassianAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryAtlassianAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


