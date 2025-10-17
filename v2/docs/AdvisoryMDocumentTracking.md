# AdvisoryMDocumentTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentReleaseDate** | Pointer to **string** |  | [optional] 
**InitialReleaseDate** | Pointer to **string** |  | [optional] 
**Identification** | Pointer to [**AdvisoryMIdentification**](AdvisoryMIdentification.md) |  | [optional] 
**Revisionhistory** | Pointer to [**[]AdvisoryRRevision**](AdvisoryRRevision.md) | diff in xml/json | [optional] 
**Status** | Pointer to **int32** | again - change in json/xml | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryMDocumentTracking

`func NewAdvisoryMDocumentTracking() *AdvisoryMDocumentTracking`

NewAdvisoryMDocumentTracking instantiates a new AdvisoryMDocumentTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMDocumentTrackingWithDefaults

`func NewAdvisoryMDocumentTrackingWithDefaults() *AdvisoryMDocumentTracking`

NewAdvisoryMDocumentTrackingWithDefaults instantiates a new AdvisoryMDocumentTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentReleaseDate

`func (o *AdvisoryMDocumentTracking) GetCurrentReleaseDate() string`

GetCurrentReleaseDate returns the CurrentReleaseDate field if non-nil, zero value otherwise.

### GetCurrentReleaseDateOk

`func (o *AdvisoryMDocumentTracking) GetCurrentReleaseDateOk() (*string, bool)`

GetCurrentReleaseDateOk returns a tuple with the CurrentReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReleaseDate

`func (o *AdvisoryMDocumentTracking) SetCurrentReleaseDate(v string)`

SetCurrentReleaseDate sets CurrentReleaseDate field to given value.

### HasCurrentReleaseDate

`func (o *AdvisoryMDocumentTracking) HasCurrentReleaseDate() bool`

HasCurrentReleaseDate returns a boolean if a field has been set.

### GetInitialReleaseDate

`func (o *AdvisoryMDocumentTracking) GetInitialReleaseDate() string`

GetInitialReleaseDate returns the InitialReleaseDate field if non-nil, zero value otherwise.

### GetInitialReleaseDateOk

`func (o *AdvisoryMDocumentTracking) GetInitialReleaseDateOk() (*string, bool)`

GetInitialReleaseDateOk returns a tuple with the InitialReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialReleaseDate

`func (o *AdvisoryMDocumentTracking) SetInitialReleaseDate(v string)`

SetInitialReleaseDate sets InitialReleaseDate field to given value.

### HasInitialReleaseDate

`func (o *AdvisoryMDocumentTracking) HasInitialReleaseDate() bool`

HasInitialReleaseDate returns a boolean if a field has been set.

### GetIdentification

`func (o *AdvisoryMDocumentTracking) GetIdentification() AdvisoryMIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *AdvisoryMDocumentTracking) GetIdentificationOk() (*AdvisoryMIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *AdvisoryMDocumentTracking) SetIdentification(v AdvisoryMIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *AdvisoryMDocumentTracking) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetRevisionhistory

`func (o *AdvisoryMDocumentTracking) GetRevisionhistory() []AdvisoryRRevision`

GetRevisionhistory returns the Revisionhistory field if non-nil, zero value otherwise.

### GetRevisionhistoryOk

`func (o *AdvisoryMDocumentTracking) GetRevisionhistoryOk() (*[]AdvisoryRRevision, bool)`

GetRevisionhistoryOk returns a tuple with the Revisionhistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionhistory

`func (o *AdvisoryMDocumentTracking) SetRevisionhistory(v []AdvisoryRRevision)`

SetRevisionhistory sets Revisionhistory field to given value.

### HasRevisionhistory

`func (o *AdvisoryMDocumentTracking) HasRevisionhistory() bool`

HasRevisionhistory returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryMDocumentTracking) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryMDocumentTracking) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryMDocumentTracking) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryMDocumentTracking) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryMDocumentTracking) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryMDocumentTracking) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryMDocumentTracking) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryMDocumentTracking) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


