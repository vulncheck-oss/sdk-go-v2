# AdvisoryCvrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**Notes** | Pointer to [**[]AdvisoryDocumentNote**](AdvisoryDocumentNote.md) |  | [optional] 
**ProductTree** | Pointer to [**AdvisoryProductTree**](AdvisoryProductTree.md) |  | [optional] 
**References** | Pointer to [**[]AdvisoryCVRFReference**](AdvisoryCVRFReference.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Tracking** | Pointer to [**AdvisoryDocumentTracking**](AdvisoryDocumentTracking.md) |  | [optional] 
**Vulnerabilities** | Pointer to [**[]AdvisoryVulnerability**](AdvisoryVulnerability.md) |  | [optional] 

## Methods

### NewAdvisoryCvrf

`func NewAdvisoryCvrf() *AdvisoryCvrf`

NewAdvisoryCvrf instantiates a new AdvisoryCvrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCvrfWithDefaults

`func NewAdvisoryCvrfWithDefaults() *AdvisoryCvrf`

NewAdvisoryCvrfWithDefaults instantiates a new AdvisoryCvrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryCvrf) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCvrf) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCvrf) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCvrf) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetNotes

`func (o *AdvisoryCvrf) GetNotes() []AdvisoryDocumentNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AdvisoryCvrf) GetNotesOk() (*[]AdvisoryDocumentNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AdvisoryCvrf) SetNotes(v []AdvisoryDocumentNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AdvisoryCvrf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProductTree

`func (o *AdvisoryCvrf) GetProductTree() AdvisoryProductTree`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *AdvisoryCvrf) GetProductTreeOk() (*AdvisoryProductTree, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *AdvisoryCvrf) SetProductTree(v AdvisoryProductTree)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *AdvisoryCvrf) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCvrf) GetReferences() []AdvisoryCVRFReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCvrf) GetReferencesOk() (*[]AdvisoryCVRFReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCvrf) SetReferences(v []AdvisoryCVRFReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCvrf) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCvrf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCvrf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCvrf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCvrf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTracking

`func (o *AdvisoryCvrf) GetTracking() AdvisoryDocumentTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *AdvisoryCvrf) GetTrackingOk() (*AdvisoryDocumentTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *AdvisoryCvrf) SetTracking(v AdvisoryDocumentTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *AdvisoryCvrf) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisoryCvrf) GetVulnerabilities() []AdvisoryVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisoryCvrf) GetVulnerabilitiesOk() (*[]AdvisoryVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisoryCvrf) SetVulnerabilities(v []AdvisoryVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisoryCvrf) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


