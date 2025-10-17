# ApiUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | sort // key | [optional] 
**Issued** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**OsArch** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]ApiPackage**](ApiPackage.md) |  | [optional] 
**References** | Pointer to [**[]ApiReference**](ApiReference.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 

## Methods

### NewApiUpdate

`func NewApiUpdate() *ApiUpdate`

NewApiUpdate instantiates a new ApiUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUpdateWithDefaults

`func NewApiUpdateWithDefaults() *ApiUpdate`

NewApiUpdateWithDefaults instantiates a new ApiUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *ApiUpdate) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiUpdate) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiUpdate) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiUpdate) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *ApiUpdate) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ApiUpdate) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ApiUpdate) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ApiUpdate) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *ApiUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssued

`func (o *ApiUpdate) GetIssued() ApiDateTime`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *ApiUpdate) GetIssuedOk() (*ApiDateTime, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *ApiUpdate) SetIssued(v ApiDateTime)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *ApiUpdate) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetOsArch

`func (o *ApiUpdate) GetOsArch() string`

GetOsArch returns the OsArch field if non-nil, zero value otherwise.

### GetOsArchOk

`func (o *ApiUpdate) GetOsArchOk() (*string, bool)`

GetOsArchOk returns a tuple with the OsArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsArch

`func (o *ApiUpdate) SetOsArch(v string)`

SetOsArch sets OsArch field to given value.

### HasOsArch

`func (o *ApiUpdate) HasOsArch() bool`

HasOsArch returns a boolean if a field has been set.

### GetOsVersion

`func (o *ApiUpdate) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ApiUpdate) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ApiUpdate) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ApiUpdate) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPackages

`func (o *ApiUpdate) GetPackages() []ApiPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *ApiUpdate) GetPackagesOk() (*[]ApiPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *ApiUpdate) SetPackages(v []ApiPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *ApiUpdate) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReferences

`func (o *ApiUpdate) GetReferences() []ApiReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiUpdate) GetReferencesOk() (*[]ApiReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiUpdate) SetReferences(v []ApiReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiUpdate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiUpdate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiUpdate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiUpdate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *ApiUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ApiUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *ApiUpdate) GetUpdated() ApiDateTime`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ApiUpdate) GetUpdatedOk() (*ApiDateTime, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ApiUpdate) SetUpdated(v ApiDateTime)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ApiUpdate) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


