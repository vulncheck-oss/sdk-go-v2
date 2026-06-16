# AdvisoryIOCRansomwareAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MaliciousDomains** | Pointer to **[]string** |  | [optional] 
**MaliciousFiles** | Pointer to [**[]AdvisoryIOCFile**](AdvisoryIOCFile.md) |  | [optional] 
**MaliciousIp** | Pointer to **[]string** |  | [optional] 
**RansomwareName** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryIOCRansomwareAdvisory

`func NewAdvisoryIOCRansomwareAdvisory() *AdvisoryIOCRansomwareAdvisory`

NewAdvisoryIOCRansomwareAdvisory instantiates a new AdvisoryIOCRansomwareAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryIOCRansomwareAdvisoryWithDefaults

`func NewAdvisoryIOCRansomwareAdvisoryWithDefaults() *AdvisoryIOCRansomwareAdvisory`

NewAdvisoryIOCRansomwareAdvisoryWithDefaults instantiates a new AdvisoryIOCRansomwareAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryIOCRansomwareAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryIOCRansomwareAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryIOCRansomwareAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryIOCRansomwareAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryIOCRansomwareAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryIOCRansomwareAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMaliciousDomains

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousDomains() []string`

GetMaliciousDomains returns the MaliciousDomains field if non-nil, zero value otherwise.

### GetMaliciousDomainsOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousDomainsOk() (*[]string, bool)`

GetMaliciousDomainsOk returns a tuple with the MaliciousDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousDomains

`func (o *AdvisoryIOCRansomwareAdvisory) SetMaliciousDomains(v []string)`

SetMaliciousDomains sets MaliciousDomains field to given value.

### HasMaliciousDomains

`func (o *AdvisoryIOCRansomwareAdvisory) HasMaliciousDomains() bool`

HasMaliciousDomains returns a boolean if a field has been set.

### GetMaliciousFiles

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousFiles() []AdvisoryIOCFile`

GetMaliciousFiles returns the MaliciousFiles field if non-nil, zero value otherwise.

### GetMaliciousFilesOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousFilesOk() (*[]AdvisoryIOCFile, bool)`

GetMaliciousFilesOk returns a tuple with the MaliciousFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousFiles

`func (o *AdvisoryIOCRansomwareAdvisory) SetMaliciousFiles(v []AdvisoryIOCFile)`

SetMaliciousFiles sets MaliciousFiles field to given value.

### HasMaliciousFiles

`func (o *AdvisoryIOCRansomwareAdvisory) HasMaliciousFiles() bool`

HasMaliciousFiles returns a boolean if a field has been set.

### GetMaliciousIp

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousIp() []string`

GetMaliciousIp returns the MaliciousIp field if non-nil, zero value otherwise.

### GetMaliciousIpOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetMaliciousIpOk() (*[]string, bool)`

GetMaliciousIpOk returns a tuple with the MaliciousIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaliciousIp

`func (o *AdvisoryIOCRansomwareAdvisory) SetMaliciousIp(v []string)`

SetMaliciousIp sets MaliciousIp field to given value.

### HasMaliciousIp

`func (o *AdvisoryIOCRansomwareAdvisory) HasMaliciousIp() bool`

HasMaliciousIp returns a boolean if a field has been set.

### GetRansomwareName

`func (o *AdvisoryIOCRansomwareAdvisory) GetRansomwareName() string`

GetRansomwareName returns the RansomwareName field if non-nil, zero value otherwise.

### GetRansomwareNameOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetRansomwareNameOk() (*string, bool)`

GetRansomwareNameOk returns a tuple with the RansomwareName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRansomwareName

`func (o *AdvisoryIOCRansomwareAdvisory) SetRansomwareName(v string)`

SetRansomwareName sets RansomwareName field to given value.

### HasRansomwareName

`func (o *AdvisoryIOCRansomwareAdvisory) HasRansomwareName() bool`

HasRansomwareName returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryIOCRansomwareAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryIOCRansomwareAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryIOCRansomwareAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryIOCRansomwareAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryIOCRansomwareAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryIOCRansomwareAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryIOCRansomwareAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryIOCRansomwareAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryIOCRansomwareAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryIOCRansomwareAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryIOCRansomwareAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryIOCRansomwareAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryIOCRansomwareAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


