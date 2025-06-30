# AdvisoryMetaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | Pointer to [**AdvisoryAdvisoryDetails**](AdvisoryAdvisoryDetails.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryVulnCheckPackage**](AdvisoryVulnCheckPackage.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryOvalReference**](AdvisoryOvalReference.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMetaData

`func NewAdvisoryMetaData() *AdvisoryMetaData`

NewAdvisoryMetaData instantiates a new AdvisoryMetaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMetaDataWithDefaults

`func NewAdvisoryMetaDataWithDefaults() *AdvisoryMetaData`

NewAdvisoryMetaDataWithDefaults instantiates a new AdvisoryMetaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisory

`func (o *AdvisoryMetaData) GetAdvisory() AdvisoryAdvisoryDetails`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *AdvisoryMetaData) GetAdvisoryOk() (*AdvisoryAdvisoryDetails, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *AdvisoryMetaData) SetAdvisory(v AdvisoryAdvisoryDetails)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *AdvisoryMetaData) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryMetaData) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryMetaData) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryMetaData) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryMetaData) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryMetaData) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMetaData) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMetaData) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMetaData) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryMetaData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryMetaData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryMetaData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryMetaData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryMetaData) GetPackages() []AdvisoryVulnCheckPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryMetaData) GetPackagesOk() (*[]AdvisoryVulnCheckPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryMetaData) SetPackages(v []AdvisoryVulnCheckPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryMetaData) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMetaData) GetReferences() []AdvisoryOvalReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMetaData) GetReferencesOk() (*[]AdvisoryOvalReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMetaData) SetReferences(v []AdvisoryOvalReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMetaData) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryMetaData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryMetaData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryMetaData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryMetaData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


