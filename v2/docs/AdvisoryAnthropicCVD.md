# AdvisoryAnthropicCVD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribution** | Pointer to [**AdvisoryCredit**](AdvisoryCredit.md) |  | [optional] 
**BugClass** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateCommitted** | Pointer to **string** |  | [optional] 
**Entry** | Pointer to **[]int32** |  | [optional] 
**Ghsa** | Pointer to **[]string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**RawPreimage** | Pointer to **[]int32** |  | [optional] 
**Severities** | Pointer to [**AdvisoryCVDSeverityCompare**](AdvisoryCVDSeverityCompare.md) |  | [optional] 
**Snapshot** | Pointer to [**AdvisoryCVDSnapshot**](AdvisoryCVDSnapshot.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Timeline** | Pointer to [**[]AdvisoryTimeline**](AdvisoryTimeline.md) | Timeline / Attribution / Severities come from the detail-page HTML (parsed at ingest, not from any JSON document). Revealed-only. | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAnthropicCVD

`func NewAdvisoryAnthropicCVD() *AdvisoryAnthropicCVD`

NewAdvisoryAnthropicCVD instantiates a new AdvisoryAnthropicCVD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAnthropicCVDWithDefaults

`func NewAdvisoryAnthropicCVDWithDefaults() *AdvisoryAnthropicCVD`

NewAdvisoryAnthropicCVDWithDefaults instantiates a new AdvisoryAnthropicCVD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribution

`func (o *AdvisoryAnthropicCVD) GetAttribution() AdvisoryCredit`

GetAttribution returns the Attribution field if non-nil, zero value otherwise.

### GetAttributionOk

`func (o *AdvisoryAnthropicCVD) GetAttributionOk() (*AdvisoryCredit, bool)`

GetAttributionOk returns a tuple with the Attribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribution

`func (o *AdvisoryAnthropicCVD) SetAttribution(v AdvisoryCredit)`

SetAttribution sets Attribution field to given value.

### HasAttribution

`func (o *AdvisoryAnthropicCVD) HasAttribution() bool`

HasAttribution returns a boolean if a field has been set.

### GetBugClass

`func (o *AdvisoryAnthropicCVD) GetBugClass() string`

GetBugClass returns the BugClass field if non-nil, zero value otherwise.

### GetBugClassOk

`func (o *AdvisoryAnthropicCVD) GetBugClassOk() (*string, bool)`

GetBugClassOk returns a tuple with the BugClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugClass

`func (o *AdvisoryAnthropicCVD) SetBugClass(v string)`

SetBugClass sets BugClass field to given value.

### HasBugClass

`func (o *AdvisoryAnthropicCVD) HasBugClass() bool`

HasBugClass returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAnthropicCVD) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAnthropicCVD) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAnthropicCVD) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAnthropicCVD) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAnthropicCVD) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAnthropicCVD) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAnthropicCVD) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAnthropicCVD) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateCommitted

`func (o *AdvisoryAnthropicCVD) GetDateCommitted() string`

GetDateCommitted returns the DateCommitted field if non-nil, zero value otherwise.

### GetDateCommittedOk

`func (o *AdvisoryAnthropicCVD) GetDateCommittedOk() (*string, bool)`

GetDateCommittedOk returns a tuple with the DateCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCommitted

`func (o *AdvisoryAnthropicCVD) SetDateCommitted(v string)`

SetDateCommitted sets DateCommitted field to given value.

### HasDateCommitted

`func (o *AdvisoryAnthropicCVD) HasDateCommitted() bool`

HasDateCommitted returns a boolean if a field has been set.

### GetEntry

`func (o *AdvisoryAnthropicCVD) GetEntry() []int32`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *AdvisoryAnthropicCVD) GetEntryOk() (*[]int32, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *AdvisoryAnthropicCVD) SetEntry(v []int32)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *AdvisoryAnthropicCVD) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetGhsa

`func (o *AdvisoryAnthropicCVD) GetGhsa() []string`

GetGhsa returns the Ghsa field if non-nil, zero value otherwise.

### GetGhsaOk

`func (o *AdvisoryAnthropicCVD) GetGhsaOk() (*[]string, bool)`

GetGhsaOk returns a tuple with the Ghsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhsa

`func (o *AdvisoryAnthropicCVD) SetGhsa(v []string)`

SetGhsa sets Ghsa field to given value.

### HasGhsa

`func (o *AdvisoryAnthropicCVD) HasGhsa() bool`

HasGhsa returns a boolean if a field has been set.

### GetHash

`func (o *AdvisoryAnthropicCVD) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *AdvisoryAnthropicCVD) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *AdvisoryAnthropicCVD) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *AdvisoryAnthropicCVD) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetProject

`func (o *AdvisoryAnthropicCVD) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AdvisoryAnthropicCVD) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AdvisoryAnthropicCVD) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AdvisoryAnthropicCVD) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRawPreimage

`func (o *AdvisoryAnthropicCVD) GetRawPreimage() []int32`

GetRawPreimage returns the RawPreimage field if non-nil, zero value otherwise.

### GetRawPreimageOk

`func (o *AdvisoryAnthropicCVD) GetRawPreimageOk() (*[]int32, bool)`

GetRawPreimageOk returns a tuple with the RawPreimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPreimage

`func (o *AdvisoryAnthropicCVD) SetRawPreimage(v []int32)`

SetRawPreimage sets RawPreimage field to given value.

### HasRawPreimage

`func (o *AdvisoryAnthropicCVD) HasRawPreimage() bool`

HasRawPreimage returns a boolean if a field has been set.

### GetSeverities

`func (o *AdvisoryAnthropicCVD) GetSeverities() AdvisoryCVDSeverityCompare`

GetSeverities returns the Severities field if non-nil, zero value otherwise.

### GetSeveritiesOk

`func (o *AdvisoryAnthropicCVD) GetSeveritiesOk() (*AdvisoryCVDSeverityCompare, bool)`

GetSeveritiesOk returns a tuple with the Severities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverities

`func (o *AdvisoryAnthropicCVD) SetSeverities(v AdvisoryCVDSeverityCompare)`

SetSeverities sets Severities field to given value.

### HasSeverities

`func (o *AdvisoryAnthropicCVD) HasSeverities() bool`

HasSeverities returns a boolean if a field has been set.

### GetSnapshot

`func (o *AdvisoryAnthropicCVD) GetSnapshot() AdvisoryCVDSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *AdvisoryAnthropicCVD) GetSnapshotOk() (*AdvisoryCVDSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *AdvisoryAnthropicCVD) SetSnapshot(v AdvisoryCVDSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *AdvisoryAnthropicCVD) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryAnthropicCVD) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryAnthropicCVD) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryAnthropicCVD) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryAnthropicCVD) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeline

`func (o *AdvisoryAnthropicCVD) GetTimeline() []AdvisoryTimeline`

GetTimeline returns the Timeline field if non-nil, zero value otherwise.

### GetTimelineOk

`func (o *AdvisoryAnthropicCVD) GetTimelineOk() (*[]AdvisoryTimeline, bool)`

GetTimelineOk returns a tuple with the Timeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeline

`func (o *AdvisoryAnthropicCVD) SetTimeline(v []AdvisoryTimeline)`

SetTimeline sets Timeline field to given value.

### HasTimeline

`func (o *AdvisoryAnthropicCVD) HasTimeline() bool`

HasTimeline returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryAnthropicCVD) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryAnthropicCVD) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryAnthropicCVD) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryAnthropicCVD) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryAnthropicCVD) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryAnthropicCVD) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryAnthropicCVD) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryAnthropicCVD) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


