# AdvisoryUbuntuCVE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPackages** | Pointer to [**[]AdvisoryAffectedUbuntuPackage**](AdvisoryAffectedUbuntuPackage.md) |  | [optional] 
**Cve** | Pointer to **[]string** | Candidate | [optional] 
**DateAdded** | Pointer to **string** | PublicDate | [optional] 
**ReferenceUrls** | Pointer to **[]string** | References | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | active || retired | [optional] 
**UbuntuUrl** | Pointer to **string** |  | [optional] 
**Usn** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryUbuntuCVE

`func NewAdvisoryUbuntuCVE() *AdvisoryUbuntuCVE`

NewAdvisoryUbuntuCVE instantiates a new AdvisoryUbuntuCVE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryUbuntuCVEWithDefaults

`func NewAdvisoryUbuntuCVEWithDefaults() *AdvisoryUbuntuCVE`

NewAdvisoryUbuntuCVEWithDefaults instantiates a new AdvisoryUbuntuCVE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPackages

`func (o *AdvisoryUbuntuCVE) GetAffectedPackages() []AdvisoryAffectedUbuntuPackage`

GetAffectedPackages returns the AffectedPackages field if non-nil, zero value otherwise.

### GetAffectedPackagesOk

`func (o *AdvisoryUbuntuCVE) GetAffectedPackagesOk() (*[]AdvisoryAffectedUbuntuPackage, bool)`

GetAffectedPackagesOk returns a tuple with the AffectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPackages

`func (o *AdvisoryUbuntuCVE) SetAffectedPackages(v []AdvisoryAffectedUbuntuPackage)`

SetAffectedPackages sets AffectedPackages field to given value.

### HasAffectedPackages

`func (o *AdvisoryUbuntuCVE) HasAffectedPackages() bool`

HasAffectedPackages returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryUbuntuCVE) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryUbuntuCVE) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryUbuntuCVE) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryUbuntuCVE) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryUbuntuCVE) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryUbuntuCVE) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryUbuntuCVE) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryUbuntuCVE) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetReferenceUrls

`func (o *AdvisoryUbuntuCVE) GetReferenceUrls() []string`

GetReferenceUrls returns the ReferenceUrls field if non-nil, zero value otherwise.

### GetReferenceUrlsOk

`func (o *AdvisoryUbuntuCVE) GetReferenceUrlsOk() (*[]string, bool)`

GetReferenceUrlsOk returns a tuple with the ReferenceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceUrls

`func (o *AdvisoryUbuntuCVE) SetReferenceUrls(v []string)`

SetReferenceUrls sets ReferenceUrls field to given value.

### HasReferenceUrls

`func (o *AdvisoryUbuntuCVE) HasReferenceUrls() bool`

HasReferenceUrls returns a boolean if a field has been set.

### GetSourceUrl

`func (o *AdvisoryUbuntuCVE) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *AdvisoryUbuntuCVE) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *AdvisoryUbuntuCVE) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *AdvisoryUbuntuCVE) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryUbuntuCVE) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryUbuntuCVE) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryUbuntuCVE) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryUbuntuCVE) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUbuntuUrl

`func (o *AdvisoryUbuntuCVE) GetUbuntuUrl() string`

GetUbuntuUrl returns the UbuntuUrl field if non-nil, zero value otherwise.

### GetUbuntuUrlOk

`func (o *AdvisoryUbuntuCVE) GetUbuntuUrlOk() (*string, bool)`

GetUbuntuUrlOk returns a tuple with the UbuntuUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUbuntuUrl

`func (o *AdvisoryUbuntuCVE) SetUbuntuUrl(v string)`

SetUbuntuUrl sets UbuntuUrl field to given value.

### HasUbuntuUrl

`func (o *AdvisoryUbuntuCVE) HasUbuntuUrl() bool`

HasUbuntuUrl returns a boolean if a field has been set.

### GetUsn

`func (o *AdvisoryUbuntuCVE) GetUsn() []string`

GetUsn returns the Usn field if non-nil, zero value otherwise.

### GetUsnOk

`func (o *AdvisoryUbuntuCVE) GetUsnOk() (*[]string, bool)`

GetUsnOk returns a tuple with the Usn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsn

`func (o *AdvisoryUbuntuCVE) SetUsn(v []string)`

SetUsn sets Usn field to given value.

### HasUsn

`func (o *AdvisoryUbuntuCVE) HasUsn() bool`

HasUsn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


