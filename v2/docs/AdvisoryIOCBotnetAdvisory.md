# AdvisoryIOCBotnetAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotnetName** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MaliciousDomains** | Pointer to **[]string** |  | [optional] 
**MaliciousFiles** | Pointer to [**[]AdvisoryIOCFile**](AdvisoryIOCFile.md) |  | [optional] 
**MaliciousIp** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryIOCBotnetAdvisory

`func NewAdvisoryIOCBotnetAdvisory() *AdvisoryIOCBotnetAdvisory`

NewAdvisoryIOCBotnetAdvisory instantiates a new AdvisoryIOCBotnetAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryIOCBotnetAdvisoryWithDefaults

`func NewAdvisoryIOCBotnetAdvisoryWithDefaults() *AdvisoryIOCBotnetAdvisory`

NewAdvisoryIOCBotnetAdvisoryWithDefaults instantiates a new AdvisoryIOCBotnetAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotnetName

`func (o *AdvisoryIOCBotnetAdvisory) GetBotnetName() string`

GetBotnetName returns the BotnetName field if non-nil, zero value otherwise.

### GetBotnetNameOk

`func (o *AdvisoryIOCBotnetAdvisory) GetBotnetNameOk() (*string, bool)`

GetBotnetNameOk returns a tuple with the BotnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotnetName

`func (o *AdvisoryIOCBotnetAdvisory) SetBotnetName(v string)`

SetBotnetName sets BotnetName field to given value.

### HasBotnetName

`func (o *AdvisoryIOCBotnetAdvisory) HasBotnetName() bool`

HasBotnetName returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryIOCBotnetAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryIOCBotnetAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryIOCBotnetAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryIOCBotnetAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryIOCBotnetAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryIOCBotnetAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryIOCBotnetAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryIOCBotnetAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMaliciousDomains

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousDomains() []string`

GetMaliciousDomains returns the MaliciousDomains field if non-nil, zero value otherwise.

### GetMaliciousDomainsOk

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousDomainsOk() (*[]string, bool)`

GetMaliciousDomainsOk returns a tuple with the MaliciousDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousDomains

`func (o *AdvisoryIOCBotnetAdvisory) SetMaliciousDomains(v []string)`

SetMaliciousDomains sets MaliciousDomains field to given value.

### HasMaliciousDomains

`func (o *AdvisoryIOCBotnetAdvisory) HasMaliciousDomains() bool`

HasMaliciousDomains returns a boolean if a field has been set.

### GetMaliciousFiles

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousFiles() []AdvisoryIOCFile`

GetMaliciousFiles returns the MaliciousFiles field if non-nil, zero value otherwise.

### GetMaliciousFilesOk

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousFilesOk() (*[]AdvisoryIOCFile, bool)`

GetMaliciousFilesOk returns a tuple with the MaliciousFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousFiles

`func (o *AdvisoryIOCBotnetAdvisory) SetMaliciousFiles(v []AdvisoryIOCFile)`

SetMaliciousFiles sets MaliciousFiles field to given value.

### HasMaliciousFiles

`func (o *AdvisoryIOCBotnetAdvisory) HasMaliciousFiles() bool`

HasMaliciousFiles returns a boolean if a field has been set.

### GetMaliciousIp

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousIp() []string`

GetMaliciousIp returns the MaliciousIp field if non-nil, zero value otherwise.

### GetMaliciousIpOk

`func (o *AdvisoryIOCBotnetAdvisory) GetMaliciousIpOk() (*[]string, bool)`

GetMaliciousIpOk returns a tuple with the MaliciousIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousIp

`func (o *AdvisoryIOCBotnetAdvisory) SetMaliciousIp(v []string)`

SetMaliciousIp sets MaliciousIp field to given value.

### HasMaliciousIp

`func (o *AdvisoryIOCBotnetAdvisory) HasMaliciousIp() bool`

HasMaliciousIp returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryIOCBotnetAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryIOCBotnetAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryIOCBotnetAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryIOCBotnetAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryIOCBotnetAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryIOCBotnetAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryIOCBotnetAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryIOCBotnetAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryIOCBotnetAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryIOCBotnetAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryIOCBotnetAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryIOCBotnetAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryIOCBotnetAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryIOCBotnetAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryIOCBotnetAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryIOCBotnetAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


