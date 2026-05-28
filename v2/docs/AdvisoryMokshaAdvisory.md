# AdvisoryMokshaAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**CveRef** | Pointer to [**AdvisoryMitreCVEListV5Ref**](AdvisoryMitreCVEListV5Ref.md) |  | [optional] 
**Cvss3Score** | Pointer to **float32** |  | [optional] 
**Cvss3Severity** | Pointer to **string** |  | [optional] 
**Cvss4Score** | Pointer to **float32** |  | [optional] 
**Cvss4Severity** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Gcve** | Pointer to **string** |  | [optional] 
**MokshaId** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMokshaAdvisory

`func NewAdvisoryMokshaAdvisory() *AdvisoryMokshaAdvisory`

NewAdvisoryMokshaAdvisory instantiates a new AdvisoryMokshaAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMokshaAdvisoryWithDefaults

`func NewAdvisoryMokshaAdvisoryWithDefaults() *AdvisoryMokshaAdvisory`

NewAdvisoryMokshaAdvisoryWithDefaults instantiates a new AdvisoryMokshaAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryMokshaAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMokshaAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMokshaAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMokshaAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveRef

`func (o *AdvisoryMokshaAdvisory) GetCveRef() AdvisoryMitreCVEListV5Ref`

GetCveRef returns the CveRef field if non-nil, zero value otherwise.

### GetCveRefOk

`func (o *AdvisoryMokshaAdvisory) GetCveRefOk() (*AdvisoryMitreCVEListV5Ref, bool)`

GetCveRefOk returns a tuple with the CveRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveRef

`func (o *AdvisoryMokshaAdvisory) SetCveRef(v AdvisoryMitreCVEListV5Ref)`

SetCveRef sets CveRef field to given value.

### HasCveRef

`func (o *AdvisoryMokshaAdvisory) HasCveRef() bool`

HasCveRef returns a boolean if a field has been set.

### GetCvss3Score

`func (o *AdvisoryMokshaAdvisory) GetCvss3Score() float32`

GetCvss3Score returns the Cvss3Score field if non-nil, zero value otherwise.

### GetCvss3ScoreOk

`func (o *AdvisoryMokshaAdvisory) GetCvss3ScoreOk() (*float32, bool)`

GetCvss3ScoreOk returns a tuple with the Cvss3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Score

`func (o *AdvisoryMokshaAdvisory) SetCvss3Score(v float32)`

SetCvss3Score sets Cvss3Score field to given value.

### HasCvss3Score

`func (o *AdvisoryMokshaAdvisory) HasCvss3Score() bool`

HasCvss3Score returns a boolean if a field has been set.

### GetCvss3Severity

`func (o *AdvisoryMokshaAdvisory) GetCvss3Severity() string`

GetCvss3Severity returns the Cvss3Severity field if non-nil, zero value otherwise.

### GetCvss3SeverityOk

`func (o *AdvisoryMokshaAdvisory) GetCvss3SeverityOk() (*string, bool)`

GetCvss3SeverityOk returns a tuple with the Cvss3Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Severity

`func (o *AdvisoryMokshaAdvisory) SetCvss3Severity(v string)`

SetCvss3Severity sets Cvss3Severity field to given value.

### HasCvss3Severity

`func (o *AdvisoryMokshaAdvisory) HasCvss3Severity() bool`

HasCvss3Severity returns a boolean if a field has been set.

### GetCvss4Score

`func (o *AdvisoryMokshaAdvisory) GetCvss4Score() float32`

GetCvss4Score returns the Cvss4Score field if non-nil, zero value otherwise.

### GetCvss4ScoreOk

`func (o *AdvisoryMokshaAdvisory) GetCvss4ScoreOk() (*float32, bool)`

GetCvss4ScoreOk returns a tuple with the Cvss4Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss4Score

`func (o *AdvisoryMokshaAdvisory) SetCvss4Score(v float32)`

SetCvss4Score sets Cvss4Score field to given value.

### HasCvss4Score

`func (o *AdvisoryMokshaAdvisory) HasCvss4Score() bool`

HasCvss4Score returns a boolean if a field has been set.

### GetCvss4Severity

`func (o *AdvisoryMokshaAdvisory) GetCvss4Severity() string`

GetCvss4Severity returns the Cvss4Severity field if non-nil, zero value otherwise.

### GetCvss4SeverityOk

`func (o *AdvisoryMokshaAdvisory) GetCvss4SeverityOk() (*string, bool)`

GetCvss4SeverityOk returns a tuple with the Cvss4Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss4Severity

`func (o *AdvisoryMokshaAdvisory) SetCvss4Severity(v string)`

SetCvss4Severity sets Cvss4Severity field to given value.

### HasCvss4Severity

`func (o *AdvisoryMokshaAdvisory) HasCvss4Severity() bool`

HasCvss4Severity returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMokshaAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMokshaAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMokshaAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMokshaAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetGcve

`func (o *AdvisoryMokshaAdvisory) GetGcve() string`

GetGcve returns the Gcve field if non-nil, zero value otherwise.

### GetGcveOk

`func (o *AdvisoryMokshaAdvisory) GetGcveOk() (*string, bool)`

GetGcveOk returns a tuple with the Gcve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcve

`func (o *AdvisoryMokshaAdvisory) SetGcve(v string)`

SetGcve sets Gcve field to given value.

### HasGcve

`func (o *AdvisoryMokshaAdvisory) HasGcve() bool`

HasGcve returns a boolean if a field has been set.

### GetMokshaId

`func (o *AdvisoryMokshaAdvisory) GetMokshaId() string`

GetMokshaId returns the MokshaId field if non-nil, zero value otherwise.

### GetMokshaIdOk

`func (o *AdvisoryMokshaAdvisory) GetMokshaIdOk() (*string, bool)`

GetMokshaIdOk returns a tuple with the MokshaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMokshaId

`func (o *AdvisoryMokshaAdvisory) SetMokshaId(v string)`

SetMokshaId sets MokshaId field to given value.

### HasMokshaId

`func (o *AdvisoryMokshaAdvisory) HasMokshaId() bool`

HasMokshaId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMokshaAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMokshaAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMokshaAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMokshaAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMokshaAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMokshaAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMokshaAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMokshaAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryMokshaAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryMokshaAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryMokshaAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryMokshaAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryMokshaAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryMokshaAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryMokshaAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryMokshaAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


