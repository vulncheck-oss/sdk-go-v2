# AdvisoryUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | sort // key | [optional] 
**Issued** | Pointer to [**AdvisoryDateTime**](AdvisoryDateTime.md) |  | [optional] 
**OsArch** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryPackage**](AdvisoryPackage.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryReference**](AdvisoryReference.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to [**AdvisoryDateTime**](AdvisoryDateTime.md) |  | [optional] 

## Methods

### NewAdvisoryUpdate

`func NewAdvisoryUpdate() *AdvisoryUpdate`

NewAdvisoryUpdate instantiates a new AdvisoryUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryUpdateWithDefaults

`func NewAdvisoryUpdateWithDefaults() *AdvisoryUpdate`

NewAdvisoryUpdateWithDefaults instantiates a new AdvisoryUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryUpdate) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryUpdate) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryUpdate) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryUpdate) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryUpdate) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryUpdate) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryUpdate) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryUpdate) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssued

`func (o *AdvisoryUpdate) GetIssued() AdvisoryDateTime`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *AdvisoryUpdate) GetIssuedOk() (*AdvisoryDateTime, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *AdvisoryUpdate) SetIssued(v AdvisoryDateTime)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *AdvisoryUpdate) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetOsArch

`func (o *AdvisoryUpdate) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *AdvisoryUpdate) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *AdvisoryUpdate) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *AdvisoryUpdate) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### GetOsVersion

`func (o *AdvisoryUpdate) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *AdvisoryUpdate) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *AdvisoryUpdate) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *AdvisoryUpdate) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryUpdate) GetPackages() []AdvisoryPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryUpdate) GetPackagesOk() (*[]AdvisoryPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryUpdate) SetPackages(v []AdvisoryPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryUpdate) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryUpdate) GetReferences() []AdvisoryReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryUpdate) GetReferencesOk() (*[]AdvisoryReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryUpdate) SetReferences(v []AdvisoryReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryUpdate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryUpdate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryUpdate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryUpdate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *AdvisoryUpdate) GetUpdated() AdvisoryDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AdvisoryUpdate) GetUpdatedOk() (*AdvisoryDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AdvisoryUpdate) SetUpdated(v AdvisoryDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AdvisoryUpdate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


