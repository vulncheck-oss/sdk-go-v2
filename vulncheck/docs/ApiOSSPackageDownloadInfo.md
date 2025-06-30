# ApiOSSPackageDownloadInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hashes** | Pointer to [**[]ApiOSSPackageHashInfo**](ApiOSSPackageHashInfo.md) |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | See OSSPackageDownloadInfoType* consts | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewApiOSSPackageDownloadInfo

`func NewApiOSSPackageDownloadInfo() *ApiOSSPackageDownloadInfo`

NewApiOSSPackageDownloadInfo instantiates a new ApiOSSPackageDownloadInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOSSPackageDownloadInfoWithDefaults

`func NewApiOSSPackageDownloadInfoWithDefaults() *ApiOSSPackageDownloadInfo`

NewApiOSSPackageDownloadInfoWithDefaults instantiates a new ApiOSSPackageDownloadInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashes

`func (o *ApiOSSPackageDownloadInfo) GetHashes() []ApiOSSPackageHashInfo`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *ApiOSSPackageDownloadInfo) GetHashesOk() (*[]ApiOSSPackageHashInfo, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *ApiOSSPackageDownloadInfo) SetHashes(v []ApiOSSPackageHashInfo)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *ApiOSSPackageDownloadInfo) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetReference

`func (o *ApiOSSPackageDownloadInfo) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApiOSSPackageDownloadInfo) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApiOSSPackageDownloadInfo) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApiOSSPackageDownloadInfo) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *ApiOSSPackageDownloadInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiOSSPackageDownloadInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiOSSPackageDownloadInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiOSSPackageDownloadInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ApiOSSPackageDownloadInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiOSSPackageDownloadInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiOSSPackageDownloadInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiOSSPackageDownloadInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


