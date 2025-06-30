# AdvisoryDocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CsafVersion** | Pointer to **string** |  | [optional] 
**Distribution** | Pointer to **map[string]interface{}** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**[]AdvisoryCSAFNote**](AdvisoryCSAFNote.md) | used by ncsc | [optional] 
**Publisher** | Pointer to [**AdvisoryPublisher**](AdvisoryPublisher.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryCSAFReference**](AdvisoryCSAFReference.md) |  | [optional] 
**Title** | Pointer to **string** | Aggregate severity is a vehicle that is provided by the document producer to convey the urgency and criticality with which the one or more vulnerabilities reported should be addressed. | [optional] 
**Tracking** | Pointer to [**AdvisoryTracking**](AdvisoryTracking.md) |  | [optional] 

## Methods

### NewAdvisoryDocumentMetadata

`func NewAdvisoryDocumentMetadata() *AdvisoryDocumentMetadata`

NewAdvisoryDocumentMetadata instantiates a new AdvisoryDocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDocumentMetadataWithDefaults

`func NewAdvisoryDocumentMetadataWithDefaults() *AdvisoryDocumentMetadata`

NewAdvisoryDocumentMetadataWithDefaults instantiates a new AdvisoryDocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AdvisoryDocumentMetadata) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryDocumentMetadata) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryDocumentMetadata) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryDocumentMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCsafVersion

`func (o *AdvisoryDocumentMetadata) GetCsafVersion() string`

GetCsafVersion returns the CsafVersion field if non-nil, zero value otherwise.

### GetCsafVersionOk

`func (o *AdvisoryDocumentMetadata) GetCsafVersionOk() (*string, bool)`

GetCsafVersionOk returns a tuple with the CsafVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsafVersion

`func (o *AdvisoryDocumentMetadata) SetCsafVersion(v string)`

SetCsafVersion sets CsafVersion field to given value.

### HasCsafVersion

`func (o *AdvisoryDocumentMetadata) HasCsafVersion() bool`

HasCsafVersion returns a boolean if a field has been set.

### GetDistribution

`func (o *AdvisoryDocumentMetadata) GetDistribution() map[string]interface{}`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *AdvisoryDocumentMetadata) GetDistributionOk() (*map[string]interface{}, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *AdvisoryDocumentMetadata) SetDistribution(v map[string]interface{})`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *AdvisoryDocumentMetadata) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetLang

`func (o *AdvisoryDocumentMetadata) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *AdvisoryDocumentMetadata) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *AdvisoryDocumentMetadata) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *AdvisoryDocumentMetadata) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetNotes

`func (o *AdvisoryDocumentMetadata) GetNotes() []AdvisoryCSAFNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AdvisoryDocumentMetadata) GetNotesOk() (*[]AdvisoryCSAFNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AdvisoryDocumentMetadata) SetNotes(v []AdvisoryCSAFNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AdvisoryDocumentMetadata) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPublisher

`func (o *AdvisoryDocumentMetadata) GetPublisher() AdvisoryPublisher`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AdvisoryDocumentMetadata) GetPublisherOk() (*AdvisoryPublisher, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AdvisoryDocumentMetadata) SetPublisher(v AdvisoryPublisher)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AdvisoryDocumentMetadata) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryDocumentMetadata) GetReferences() []AdvisoryCSAFReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryDocumentMetadata) GetReferencesOk() (*[]AdvisoryCSAFReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryDocumentMetadata) SetReferences(v []AdvisoryCSAFReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryDocumentMetadata) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDocumentMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDocumentMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDocumentMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDocumentMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTracking

`func (o *AdvisoryDocumentMetadata) GetTracking() AdvisoryTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *AdvisoryDocumentMetadata) GetTrackingOk() (*AdvisoryTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *AdvisoryDocumentMetadata) SetTracking(v AdvisoryTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *AdvisoryDocumentMetadata) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


