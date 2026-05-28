# AdvisoryMogwaiLabsAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedVersion** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FixedVersion** | Pointer to **string** |  | [optional] 
**Gcve** | Pointer to **string** | From detail page | [optional] 
**Id** | Pointer to **string** | From index table | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMogwaiLabsAdvisory

`func NewAdvisoryMogwaiLabsAdvisory() *AdvisoryMogwaiLabsAdvisory`

NewAdvisoryMogwaiLabsAdvisory instantiates a new AdvisoryMogwaiLabsAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMogwaiLabsAdvisoryWithDefaults

`func NewAdvisoryMogwaiLabsAdvisoryWithDefaults() *AdvisoryMogwaiLabsAdvisory`

NewAdvisoryMogwaiLabsAdvisoryWithDefaults instantiates a new AdvisoryMogwaiLabsAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) GetAffectedVersion() string`

GetAffectedVersion returns the AffectedVersion field if non-nil, zero value otherwise.

### GetAffectedVersionOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetAffectedVersionOk() (*string, bool)`

GetAffectedVersionOk returns a tuple with the AffectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) SetAffectedVersion(v string)`

SetAffectedVersion sets AffectedVersion field to given value.

### HasAffectedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) HasAffectedVersion() bool`

HasAffectedVersion returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryMogwaiLabsAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMogwaiLabsAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMogwaiLabsAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryMogwaiLabsAdvisory) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryMogwaiLabsAdvisory) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryMogwaiLabsAdvisory) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMogwaiLabsAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMogwaiLabsAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMogwaiLabsAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryMogwaiLabsAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMogwaiLabsAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMogwaiLabsAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) GetFixedVersion() string`

GetFixedVersion returns the FixedVersion field if non-nil, zero value otherwise.

### GetFixedVersionOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetFixedVersionOk() (*string, bool)`

GetFixedVersionOk returns a tuple with the FixedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) SetFixedVersion(v string)`

SetFixedVersion sets FixedVersion field to given value.

### HasFixedVersion

`func (o *AdvisoryMogwaiLabsAdvisory) HasFixedVersion() bool`

HasFixedVersion returns a boolean if a field has been set.

### GetGcve

`func (o *AdvisoryMogwaiLabsAdvisory) GetGcve() string`

GetGcve returns the Gcve field if non-nil, zero value otherwise.

### GetGcveOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetGcveOk() (*string, bool)`

GetGcveOk returns a tuple with the Gcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcve

`func (o *AdvisoryMogwaiLabsAdvisory) SetGcve(v string)`

SetGcve sets Gcve field to given value.

### HasGcve

`func (o *AdvisoryMogwaiLabsAdvisory) HasGcve() bool`

HasGcve returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryMogwaiLabsAdvisory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryMogwaiLabsAdvisory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryMogwaiLabsAdvisory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMogwaiLabsAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMogwaiLabsAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMogwaiLabsAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryMogwaiLabsAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryMogwaiLabsAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryMogwaiLabsAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryMogwaiLabsAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryMogwaiLabsAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryMogwaiLabsAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMogwaiLabsAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMogwaiLabsAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMogwaiLabsAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryMogwaiLabsAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryMogwaiLabsAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryMogwaiLabsAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMogwaiLabsAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMogwaiLabsAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMogwaiLabsAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMogwaiLabsAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


