# BackupListBackupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BackupFeedItem**](BackupFeedItem.md) |  | [optional] 

## Methods

### NewBackupListBackupsResponse

`func NewBackupListBackupsResponse() *BackupListBackupsResponse`

NewBackupListBackupsResponse instantiates a new BackupListBackupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupListBackupsResponseWithDefaults

`func NewBackupListBackupsResponseWithDefaults() *BackupListBackupsResponse`

NewBackupListBackupsResponseWithDefaults instantiates a new BackupListBackupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BackupListBackupsResponse) GetData() []BackupFeedItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BackupListBackupsResponse) GetDataOk() (*[]BackupFeedItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BackupListBackupsResponse) SetData(v []BackupFeedItem)`

SetData sets Data field to given value.

### HasData

`func (o *BackupListBackupsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


