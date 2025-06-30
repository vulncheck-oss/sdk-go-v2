# AdvisorySiemensDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledgments** | Pointer to [**[]AdvisorySiemensAcknowledgments**](AdvisorySiemensAcknowledgments.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**CsafVersion** | Pointer to **string** |  | [optional] 
**Distribution** | Pointer to [**AdvisorySiemensDistribution**](AdvisorySiemensDistribution.md) |  | [optional] 
**Notes** | Pointer to [**[]AdvisorySiemensNotes**](AdvisorySiemensNotes.md) |  | [optional] 
**Publisher** | Pointer to [**AdvisorySiemensPublisher**](AdvisorySiemensPublisher.md) |  | [optional] 
**References** | Pointer to [**[]AdvisorySiemensReferences**](AdvisorySiemensReferences.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tracking** | Pointer to [**AdvisorySiemensTracking**](AdvisorySiemensTracking.md) |  | [optional] 

## Methods

### NewAdvisorySiemensDocument

`func NewAdvisorySiemensDocument() *AdvisorySiemensDocument`

NewAdvisorySiemensDocument instantiates a new AdvisorySiemensDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySiemensDocumentWithDefaults

`func NewAdvisorySiemensDocumentWithDefaults() *AdvisorySiemensDocument`

NewAdvisorySiemensDocumentWithDefaults instantiates a new AdvisorySiemensDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgments

`func (o *AdvisorySiemensDocument) GetAcknowledgments() []AdvisorySiemensAcknowledgments`

GetAcknowledgments returns the Acknowledgments field if non-nil, zero value otherwise.

### GetAcknowledgmentsOk

`func (o *AdvisorySiemensDocument) GetAcknowledgmentsOk() (*[]AdvisorySiemensAcknowledgments, bool)`

GetAcknowledgmentsOk returns a tuple with the Acknowledgments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgments

`func (o *AdvisorySiemensDocument) SetAcknowledgments(v []AdvisorySiemensAcknowledgments)`

SetAcknowledgments sets Acknowledgments field to given value.

### HasAcknowledgments

`func (o *AdvisorySiemensDocument) HasAcknowledgments() bool`

HasAcknowledgments returns a boolean if a field has been set.

### GetCategory

`func (o *AdvisorySiemensDocument) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisorySiemensDocument) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisorySiemensDocument) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisorySiemensDocument) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCsafVersion

`func (o *AdvisorySiemensDocument) GetCsafVersion() string`

GetCsafVersion returns the CsafVersion field if non-nil, zero value otherwise.

### GetCsafVersionOk

`func (o *AdvisorySiemensDocument) GetCsafVersionOk() (*string, bool)`

GetCsafVersionOk returns a tuple with the CsafVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsafVersion

`func (o *AdvisorySiemensDocument) SetCsafVersion(v string)`

SetCsafVersion sets CsafVersion field to given value.

### HasCsafVersion

`func (o *AdvisorySiemensDocument) HasCsafVersion() bool`

HasCsafVersion returns a boolean if a field has been set.

### GetDistribution

`func (o *AdvisorySiemensDocument) GetDistribution() AdvisorySiemensDistribution`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *AdvisorySiemensDocument) GetDistributionOk() (*AdvisorySiemensDistribution, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *AdvisorySiemensDocument) SetDistribution(v AdvisorySiemensDistribution)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *AdvisorySiemensDocument) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetNotes

`func (o *AdvisorySiemensDocument) GetNotes() []AdvisorySiemensNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AdvisorySiemensDocument) GetNotesOk() (*[]AdvisorySiemensNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AdvisorySiemensDocument) SetNotes(v []AdvisorySiemensNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AdvisorySiemensDocument) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPublisher

`func (o *AdvisorySiemensDocument) GetPublisher() AdvisorySiemensPublisher`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AdvisorySiemensDocument) GetPublisherOk() (*AdvisorySiemensPublisher, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AdvisorySiemensDocument) SetPublisher(v AdvisorySiemensPublisher)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AdvisorySiemensDocument) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisorySiemensDocument) GetReferences() []AdvisorySiemensReferences`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisorySiemensDocument) GetReferencesOk() (*[]AdvisorySiemensReferences, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisorySiemensDocument) SetReferences(v []AdvisorySiemensReferences)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisorySiemensDocument) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySiemensDocument) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySiemensDocument) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySiemensDocument) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySiemensDocument) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTracking

`func (o *AdvisorySiemensDocument) GetTracking() AdvisorySiemensTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *AdvisorySiemensDocument) GetTrackingOk() (*AdvisorySiemensTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *AdvisorySiemensDocument) SetTracking(v AdvisorySiemensTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *AdvisorySiemensDocument) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


