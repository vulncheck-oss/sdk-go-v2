# AdvisoryMicrosoftDriverBlockList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateAdded** | Pointer to **string** |  | [optional] 
**FileId** | Pointer to **string** | From FileAttrib or Deny | [optional] 
**FileMetadata** | Pointer to [**AdvisoryMicrosoftFileMetadata**](AdvisoryMicrosoftFileMetadata.md) | File-level metadata | [optional] 

## Methods

### NewAdvisoryMicrosoftDriverBlockList

`func NewAdvisoryMicrosoftDriverBlockList() *AdvisoryMicrosoftDriverBlockList`

NewAdvisoryMicrosoftDriverBlockList instantiates a new AdvisoryMicrosoftDriverBlockList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMicrosoftDriverBlockListWithDefaults

`func NewAdvisoryMicrosoftDriverBlockListWithDefaults() *AdvisoryMicrosoftDriverBlockList`

NewAdvisoryMicrosoftDriverBlockListWithDefaults instantiates a new AdvisoryMicrosoftDriverBlockList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateAdded

`func (o *AdvisoryMicrosoftDriverBlockList) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryMicrosoftDriverBlockList) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryMicrosoftDriverBlockList) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryMicrosoftDriverBlockList) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFileId

`func (o *AdvisoryMicrosoftDriverBlockList) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AdvisoryMicrosoftDriverBlockList) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AdvisoryMicrosoftDriverBlockList) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AdvisoryMicrosoftDriverBlockList) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetFileMetadata

`func (o *AdvisoryMicrosoftDriverBlockList) GetFileMetadata() AdvisoryMicrosoftFileMetadata`

GetFileMetadata returns the FileMetadata field if non-nil, zero value otherwise.

### GetFileMetadataOk

`func (o *AdvisoryMicrosoftDriverBlockList) GetFileMetadataOk() (*AdvisoryMicrosoftFileMetadata, bool)`

GetFileMetadataOk returns a tuple with the FileMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMetadata

`func (o *AdvisoryMicrosoftDriverBlockList) SetFileMetadata(v AdvisoryMicrosoftFileMetadata)`

SetFileMetadata sets FileMetadata field to given value.

### HasFileMetadata

`func (o *AdvisoryMicrosoftDriverBlockList) HasFileMetadata() bool`

HasFileMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


