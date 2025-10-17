# AdvisoryOverrideAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CveId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to **string** |  | [optional] 
**TriageNotes** | Pointer to [**AdvisoryTriageNotes**](AdvisoryTriageNotes.md) |  | [optional] 

## Methods

### NewAdvisoryOverrideAnnotation

`func NewAdvisoryOverrideAnnotation() *AdvisoryOverrideAnnotation`

NewAdvisoryOverrideAnnotation instantiates a new AdvisoryOverrideAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOverrideAnnotationWithDefaults

`func NewAdvisoryOverrideAnnotationWithDefaults() *AdvisoryOverrideAnnotation`

NewAdvisoryOverrideAnnotationWithDefaults instantiates a new AdvisoryOverrideAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCveId

`func (o *AdvisoryOverrideAnnotation) GetCveId() string`

GetCveId returns the CveId field if non-nil, zero value otherwise.

### GetCveIdOk

`func (o *AdvisoryOverrideAnnotation) GetCveIdOk() (*string, bool)`

GetCveIdOk returns a tuple with the CveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveId

`func (o *AdvisoryOverrideAnnotation) SetCveId(v string)`

SetCveId sets CveId field to given value.

### HasCveId

`func (o *AdvisoryOverrideAnnotation) HasCveId() bool`

HasCveId returns a boolean if a field has been set.

### GetReason

`func (o *AdvisoryOverrideAnnotation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdvisoryOverrideAnnotation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdvisoryOverrideAnnotation) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AdvisoryOverrideAnnotation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSnapshot

`func (o *AdvisoryOverrideAnnotation) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *AdvisoryOverrideAnnotation) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *AdvisoryOverrideAnnotation) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *AdvisoryOverrideAnnotation) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetTriageNotes

`func (o *AdvisoryOverrideAnnotation) GetTriageNotes() AdvisoryTriageNotes`

GetTriageNotes returns the TriageNotes field if non-nil, zero value otherwise.

### GetTriageNotesOk

`func (o *AdvisoryOverrideAnnotation) GetTriageNotesOk() (*AdvisoryTriageNotes, bool)`

GetTriageNotesOk returns a tuple with the TriageNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriageNotes

`func (o *AdvisoryOverrideAnnotation) SetTriageNotes(v AdvisoryTriageNotes)`

SetTriageNotes sets TriageNotes field to given value.

### HasTriageNotes

`func (o *AdvisoryOverrideAnnotation) HasTriageNotes() bool`

HasTriageNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


