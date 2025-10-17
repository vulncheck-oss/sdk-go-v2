# AdvisorySecurityBulletin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledgement** | Pointer to **string** |  | [optional] 
**BulletinId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvedetails** | Pointer to [**[]AdvisoryCVEDetail**](AdvisoryCVEDetail.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**HardwareUpdates** | Pointer to [**[]AdvisoryHardwareUpdate**](AdvisoryHardwareUpdate.md) |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Revisions** | Pointer to [**[]AdvisoryNvidiaRevision**](AdvisoryNvidiaRevision.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SoftwareUpdates** | Pointer to [**[]AdvisorySoftwareUpdate**](AdvisorySoftwareUpdate.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySecurityBulletin

`func NewAdvisorySecurityBulletin() *AdvisorySecurityBulletin`

NewAdvisorySecurityBulletin instantiates a new AdvisorySecurityBulletin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySecurityBulletinWithDefaults

`func NewAdvisorySecurityBulletinWithDefaults() *AdvisorySecurityBulletin`

NewAdvisorySecurityBulletinWithDefaults instantiates a new AdvisorySecurityBulletin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgement

`func (o *AdvisorySecurityBulletin) GetAcknowledgement() string`

GetAcknowledgement returns the Acknowledgement field if non-nil, zero value otherwise.

### GetAcknowledgementOk

`func (o *AdvisorySecurityBulletin) GetAcknowledgementOk() (*string, bool)`

GetAcknowledgementOk returns a tuple with the Acknowledgement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgement

`func (o *AdvisorySecurityBulletin) SetAcknowledgement(v string)`

SetAcknowledgement sets Acknowledgement field to given value.

### HasAcknowledgement

`func (o *AdvisorySecurityBulletin) HasAcknowledgement() bool`

HasAcknowledgement returns a boolean if a field has been set.

### GetBulletinId

`func (o *AdvisorySecurityBulletin) GetBulletinId() string`

GetBulletinId returns the BulletinId field if non-nil, zero value otherwise.

### GetBulletinIdOk

`func (o *AdvisorySecurityBulletin) GetBulletinIdOk() (*string, bool)`

GetBulletinIdOk returns a tuple with the BulletinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinId

`func (o *AdvisorySecurityBulletin) SetBulletinId(v string)`

SetBulletinId sets BulletinId field to given value.

### HasBulletinId

`func (o *AdvisorySecurityBulletin) HasBulletinId() bool`

HasBulletinId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySecurityBulletin) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySecurityBulletin) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySecurityBulletin) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySecurityBulletin) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvedetails

`func (o *AdvisorySecurityBulletin) GetCvedetails() []AdvisoryCVEDetail`

GetCvedetails returns the Cvedetails field if non-nil, zero value otherwise.

### GetCvedetailsOk

`func (o *AdvisorySecurityBulletin) GetCvedetailsOk() (*[]AdvisoryCVEDetail, bool)`

GetCvedetailsOk returns a tuple with the Cvedetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvedetails

`func (o *AdvisorySecurityBulletin) SetCvedetails(v []AdvisoryCVEDetail)`

SetCvedetails sets Cvedetails field to given value.

### HasCvedetails

`func (o *AdvisorySecurityBulletin) HasCvedetails() bool`

HasCvedetails returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySecurityBulletin) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySecurityBulletin) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySecurityBulletin) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySecurityBulletin) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetHardwareUpdates

`func (o *AdvisorySecurityBulletin) GetHardwareUpdates() []AdvisoryHardwareUpdate`

GetHardwareUpdates returns the HardwareUpdates field if non-nil, zero value otherwise.

### GetHardwareUpdatesOk

`func (o *AdvisorySecurityBulletin) GetHardwareUpdatesOk() (*[]AdvisoryHardwareUpdate, bool)`

GetHardwareUpdatesOk returns a tuple with the HardwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareUpdates

`func (o *AdvisorySecurityBulletin) SetHardwareUpdates(v []AdvisoryHardwareUpdate)`

SetHardwareUpdates sets HardwareUpdates field to given value.

### HasHardwareUpdates

`func (o *AdvisorySecurityBulletin) HasHardwareUpdates() bool`

HasHardwareUpdates returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AdvisorySecurityBulletin) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AdvisorySecurityBulletin) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AdvisorySecurityBulletin) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AdvisorySecurityBulletin) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLink

`func (o *AdvisorySecurityBulletin) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisorySecurityBulletin) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisorySecurityBulletin) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisorySecurityBulletin) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRevisions

`func (o *AdvisorySecurityBulletin) GetRevisions() []AdvisoryNvidiaRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *AdvisorySecurityBulletin) GetRevisionsOk() (*[]AdvisoryNvidiaRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *AdvisorySecurityBulletin) SetRevisions(v []AdvisoryNvidiaRevision)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *AdvisorySecurityBulletin) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisorySecurityBulletin) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisorySecurityBulletin) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisorySecurityBulletin) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisorySecurityBulletin) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSoftwareUpdates

`func (o *AdvisorySecurityBulletin) GetSoftwareUpdates() []AdvisorySoftwareUpdate`

GetSoftwareUpdates returns the SoftwareUpdates field if non-nil, zero value otherwise.

### GetSoftwareUpdatesOk

`func (o *AdvisorySecurityBulletin) GetSoftwareUpdatesOk() (*[]AdvisorySoftwareUpdate, bool)`

GetSoftwareUpdatesOk returns a tuple with the SoftwareUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareUpdates

`func (o *AdvisorySecurityBulletin) SetSoftwareUpdates(v []AdvisorySoftwareUpdate)`

SetSoftwareUpdates sets SoftwareUpdates field to given value.

### HasSoftwareUpdates

`func (o *AdvisorySecurityBulletin) HasSoftwareUpdates() bool`

HasSoftwareUpdates returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySecurityBulletin) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySecurityBulletin) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySecurityBulletin) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySecurityBulletin) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisorySecurityBulletin) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisorySecurityBulletin) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisorySecurityBulletin) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisorySecurityBulletin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


