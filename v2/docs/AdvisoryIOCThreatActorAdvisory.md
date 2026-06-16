# AdvisoryIOCThreatActorAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MaliciousDomains** | Pointer to **[]string** |  | [optional] 
**MaliciousFiles** | Pointer to [**[]AdvisoryIOCFile**](AdvisoryIOCFile.md) |  | [optional] 
**MaliciousIp** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**ThreatActorName** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryIOCThreatActorAdvisory

`func NewAdvisoryIOCThreatActorAdvisory() *AdvisoryIOCThreatActorAdvisory`

NewAdvisoryIOCThreatActorAdvisory instantiates a new AdvisoryIOCThreatActorAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryIOCThreatActorAdvisoryWithDefaults

`func NewAdvisoryIOCThreatActorAdvisoryWithDefaults() *AdvisoryIOCThreatActorAdvisory`

NewAdvisoryIOCThreatActorAdvisoryWithDefaults instantiates a new AdvisoryIOCThreatActorAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryIOCThreatActorAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryIOCThreatActorAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryIOCThreatActorAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryIOCThreatActorAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryIOCThreatActorAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryIOCThreatActorAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMaliciousDomains

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousDomains() []string`

GetMaliciousDomains returns the MaliciousDomains field if non-nil, zero value otherwise.

### GetMaliciousDomainsOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousDomainsOk() (*[]string, bool)`

GetMaliciousDomainsOk returns a tuple with the MaliciousDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousDomains

`func (o *AdvisoryIOCThreatActorAdvisory) SetMaliciousDomains(v []string)`

SetMaliciousDomains sets MaliciousDomains field to given value.

### HasMaliciousDomains

`func (o *AdvisoryIOCThreatActorAdvisory) HasMaliciousDomains() bool`

HasMaliciousDomains returns a boolean if a field has been set.

### GetMaliciousFiles

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousFiles() []AdvisoryIOCFile`

GetMaliciousFiles returns the MaliciousFiles field if non-nil, zero value otherwise.

### GetMaliciousFilesOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousFilesOk() (*[]AdvisoryIOCFile, bool)`

GetMaliciousFilesOk returns a tuple with the MaliciousFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousFiles

`func (o *AdvisoryIOCThreatActorAdvisory) SetMaliciousFiles(v []AdvisoryIOCFile)`

SetMaliciousFiles sets MaliciousFiles field to given value.

### HasMaliciousFiles

`func (o *AdvisoryIOCThreatActorAdvisory) HasMaliciousFiles() bool`

HasMaliciousFiles returns a boolean if a field has been set.

### GetMaliciousIp

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousIp() []string`

GetMaliciousIp returns the MaliciousIp field if non-nil, zero value otherwise.

### GetMaliciousIpOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetMaliciousIpOk() (*[]string, bool)`

GetMaliciousIpOk returns a tuple with the MaliciousIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousIp

`func (o *AdvisoryIOCThreatActorAdvisory) SetMaliciousIp(v []string)`

SetMaliciousIp sets MaliciousIp field to given value.

### HasMaliciousIp

`func (o *AdvisoryIOCThreatActorAdvisory) HasMaliciousIp() bool`

HasMaliciousIp returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryIOCThreatActorAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryIOCThreatActorAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryIOCThreatActorAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetThreatActorName

`func (o *AdvisoryIOCThreatActorAdvisory) GetThreatActorName() string`

GetThreatActorName returns the ThreatActorName field if non-nil, zero value otherwise.

### GetThreatActorNameOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetThreatActorNameOk() (*string, bool)`

GetThreatActorNameOk returns a tuple with the ThreatActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorName

`func (o *AdvisoryIOCThreatActorAdvisory) SetThreatActorName(v string)`

SetThreatActorName sets ThreatActorName field to given value.

### HasThreatActorName

`func (o *AdvisoryIOCThreatActorAdvisory) HasThreatActorName() bool`

HasThreatActorName returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryIOCThreatActorAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryIOCThreatActorAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryIOCThreatActorAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryIOCThreatActorAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryIOCThreatActorAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryIOCThreatActorAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryIOCThreatActorAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryIOCThreatActorAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryIOCThreatActorAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryIOCThreatActorAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


