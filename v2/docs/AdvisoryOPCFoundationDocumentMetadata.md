# AdvisoryOPCFoundationDocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CsafVersion** | Pointer to **string** |  | [optional] 
**Distribution** | Pointer to [**AdvisoryOPCFoundationDistribution**](AdvisoryOPCFoundationDistribution.md) |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**[]AdvisoryCSAFNote**](AdvisoryCSAFNote.md) | used by ncsc | [optional] 
**Publisher** | Pointer to [**AdvisoryPublisher**](AdvisoryPublisher.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryCSAFReference**](AdvisoryCSAFReference.md) |  | [optional] 
**Title** | Pointer to **string** | Aggregate severity is a vehicle that is provided by the document producer to convey the urgency and criticality with which the one or more vulnerabilities reported should be addressed. | [optional] 
**Tracking** | Pointer to [**AdvisoryTracking**](AdvisoryTracking.md) |  | [optional] 

## Methods

### NewAdvisoryOPCFoundationDocumentMetadata

`func NewAdvisoryOPCFoundationDocumentMetadata() *AdvisoryOPCFoundationDocumentMetadata`

NewAdvisoryOPCFoundationDocumentMetadata instantiates a new AdvisoryOPCFoundationDocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOPCFoundationDocumentMetadataWithDefaults

`func NewAdvisoryOPCFoundationDocumentMetadataWithDefaults() *AdvisoryOPCFoundationDocumentMetadata`

NewAdvisoryOPCFoundationDocumentMetadataWithDefaults instantiates a new AdvisoryOPCFoundationDocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCsafVersion

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetCsafVersion() string`

GetCsafVersion returns the CsafVersion field if non-nil, zero value otherwise.

### GetCsafVersionOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetCsafVersionOk() (*string, bool)`

GetCsafVersionOk returns a tuple with the CsafVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsafVersion

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetCsafVersion(v string)`

SetCsafVersion sets CsafVersion field to given value.

### HasCsafVersion

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasCsafVersion() bool`

HasCsafVersion returns a boolean if a field has been set.

### GetDistribution

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetDistribution() AdvisoryOPCFoundationDistribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetDistributionOk() (*AdvisoryOPCFoundationDistribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetDistribution(v AdvisoryOPCFoundationDistribution)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetLang

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetNotes

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetNotes() []AdvisoryCSAFNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetNotesOk() (*[]AdvisoryCSAFNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetNotes(v []AdvisoryCSAFNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPublisher

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetPublisher() AdvisoryPublisher`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetPublisherOk() (*AdvisoryPublisher, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetPublisher(v AdvisoryPublisher)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetReferences() []AdvisoryCSAFReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetReferencesOk() (*[]AdvisoryCSAFReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetReferences(v []AdvisoryCSAFReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTracking

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetTracking() AdvisoryTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *AdvisoryOPCFoundationDocumentMetadata) GetTrackingOk() (*AdvisoryTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *AdvisoryOPCFoundationDocumentMetadata) SetTracking(v AdvisoryTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *AdvisoryOPCFoundationDocumentMetadata) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


