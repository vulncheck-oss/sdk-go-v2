# AdvisoryTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentReleaseDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InitialReleaseDate** | Pointer to **string** |  | [optional] 
**RevisionHistory** | Pointer to [**[]AdvisoryRevisionHistory**](AdvisoryRevisionHistory.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | should match last &#39;number&#39; in []RevisionHistory | [optional] 

## Methods

### NewAdvisoryTracking

`func NewAdvisoryTracking() *AdvisoryTracking`

NewAdvisoryTracking instantiates a new AdvisoryTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryTrackingWithDefaults

`func NewAdvisoryTrackingWithDefaults() *AdvisoryTracking`

NewAdvisoryTrackingWithDefaults instantiates a new AdvisoryTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentReleaseDate

`func (o *AdvisoryTracking) GetCurrentReleaseDate() string`

GetCurrentReleaseDate returns the CurrentReleaseDate field if non-nil, zero value otherwise.

### GetCurrentReleaseDateOk

`func (o *AdvisoryTracking) GetCurrentReleaseDateOk() (*string, bool)`

GetCurrentReleaseDateOk returns a tuple with the CurrentReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReleaseDate

`func (o *AdvisoryTracking) SetCurrentReleaseDate(v string)`

SetCurrentReleaseDate sets CurrentReleaseDate field to given value.

### HasCurrentReleaseDate

`func (o *AdvisoryTracking) HasCurrentReleaseDate() bool`

HasCurrentReleaseDate returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryTracking) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryTracking) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryTracking) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryTracking) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialReleaseDate

`func (o *AdvisoryTracking) GetInitialReleaseDate() string`

GetInitialReleaseDate returns the InitialReleaseDate field if non-nil, zero value otherwise.

### GetInitialReleaseDateOk

`func (o *AdvisoryTracking) GetInitialReleaseDateOk() (*string, bool)`

GetInitialReleaseDateOk returns a tuple with the InitialReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialReleaseDate

`func (o *AdvisoryTracking) SetInitialReleaseDate(v string)`

SetInitialReleaseDate sets InitialReleaseDate field to given value.

### HasInitialReleaseDate

`func (o *AdvisoryTracking) HasInitialReleaseDate() bool`

HasInitialReleaseDate returns a boolean if a field has been set.

### GetRevisionHistory

`func (o *AdvisoryTracking) GetRevisionHistory() []AdvisoryRevisionHistory`

GetRevisionHistory returns the RevisionHistory field if non-nil, zero value otherwise.

### GetRevisionHistoryOk

`func (o *AdvisoryTracking) GetRevisionHistoryOk() (*[]AdvisoryRevisionHistory, bool)`

GetRevisionHistoryOk returns a tuple with the RevisionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistory

`func (o *AdvisoryTracking) SetRevisionHistory(v []AdvisoryRevisionHistory)`

SetRevisionHistory sets RevisionHistory field to given value.

### HasRevisionHistory

`func (o *AdvisoryTracking) HasRevisionHistory() bool`

HasRevisionHistory returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryTracking) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryTracking) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryTracking) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryTracking) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryTracking) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryTracking) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryTracking) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryTracking) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


