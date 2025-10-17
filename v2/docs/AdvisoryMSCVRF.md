# AdvisoryMSCVRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTitle** | Pointer to [**AdvisoryMSDocumentTitle**](AdvisoryMSDocumentTitle.md) |  | [optional] 
**DocumentTracking** | Pointer to [**AdvisoryMDocumentTracking**](AdvisoryMDocumentTracking.md) |  | [optional] 
**ProductTree** | Pointer to [**AdvisoryMProductTree**](AdvisoryMProductTree.md) |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**Documentnotes** | Pointer to [**[]AdvisoryRNote**](AdvisoryRNote.md) | diff | [optional] 
**Documentpublisher** | Pointer to [**AdvisoryDocumentPublisher**](AdvisoryDocumentPublisher.md) |  | [optional] 
**Vulnerability** | Pointer to [**[]AdvisoryMVulnerability**](AdvisoryMVulnerability.md) |  | [optional] 

## Methods

### NewAdvisoryMSCVRF

`func NewAdvisoryMSCVRF() *AdvisoryMSCVRF`

NewAdvisoryMSCVRF instantiates a new AdvisoryMSCVRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMSCVRFWithDefaults

`func NewAdvisoryMSCVRFWithDefaults() *AdvisoryMSCVRF`

NewAdvisoryMSCVRFWithDefaults instantiates a new AdvisoryMSCVRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTitle

`func (o *AdvisoryMSCVRF) GetDocumentTitle() AdvisoryMSDocumentTitle`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *AdvisoryMSCVRF) GetDocumentTitleOk() (*AdvisoryMSDocumentTitle, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *AdvisoryMSCVRF) SetDocumentTitle(v AdvisoryMSDocumentTitle)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *AdvisoryMSCVRF) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetDocumentTracking

`func (o *AdvisoryMSCVRF) GetDocumentTracking() AdvisoryMDocumentTracking`

GetDocumentTracking returns the DocumentTracking field if non-nil, zero value otherwise.

### GetDocumentTrackingOk

`func (o *AdvisoryMSCVRF) GetDocumentTrackingOk() (*AdvisoryMDocumentTracking, bool)`

GetDocumentTrackingOk returns a tuple with the DocumentTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTracking

`func (o *AdvisoryMSCVRF) SetDocumentTracking(v AdvisoryMDocumentTracking)`

SetDocumentTracking sets DocumentTracking field to given value.

### HasDocumentTracking

`func (o *AdvisoryMSCVRF) HasDocumentTracking() bool`

HasDocumentTracking returns a boolean if a field has been set.

### GetProductTree

`func (o *AdvisoryMSCVRF) GetProductTree() AdvisoryMProductTree`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *AdvisoryMSCVRF) GetProductTreeOk() (*AdvisoryMProductTree, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *AdvisoryMSCVRF) SetProductTree(v AdvisoryMProductTree)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *AdvisoryMSCVRF) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetDocumentType

`func (o *AdvisoryMSCVRF) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *AdvisoryMSCVRF) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *AdvisoryMSCVRF) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *AdvisoryMSCVRF) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentnotes

`func (o *AdvisoryMSCVRF) GetDocumentnotes() []AdvisoryRNote`

GetDocumentnotes returns the Documentnotes field if non-nil, zero value otherwise.

### GetDocumentnotesOk

`func (o *AdvisoryMSCVRF) GetDocumentnotesOk() (*[]AdvisoryRNote, bool)`

GetDocumentnotesOk returns a tuple with the Documentnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentnotes

`func (o *AdvisoryMSCVRF) SetDocumentnotes(v []AdvisoryRNote)`

SetDocumentnotes sets Documentnotes field to given value.

### HasDocumentnotes

`func (o *AdvisoryMSCVRF) HasDocumentnotes() bool`

HasDocumentnotes returns a boolean if a field has been set.

### GetDocumentpublisher

`func (o *AdvisoryMSCVRF) GetDocumentpublisher() AdvisoryDocumentPublisher`

GetDocumentpublisher returns the Documentpublisher field if non-nil, zero value otherwise.

### GetDocumentpublisherOk

`func (o *AdvisoryMSCVRF) GetDocumentpublisherOk() (*AdvisoryDocumentPublisher, bool)`

GetDocumentpublisherOk returns a tuple with the Documentpublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentpublisher

`func (o *AdvisoryMSCVRF) SetDocumentpublisher(v AdvisoryDocumentPublisher)`

SetDocumentpublisher sets Documentpublisher field to given value.

### HasDocumentpublisher

`func (o *AdvisoryMSCVRF) HasDocumentpublisher() bool`

HasDocumentpublisher returns a boolean if a field has been set.

### GetVulnerability

`func (o *AdvisoryMSCVRF) GetVulnerability() []AdvisoryMVulnerability`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *AdvisoryMSCVRF) GetVulnerabilityOk() (*[]AdvisoryMVulnerability, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *AdvisoryMSCVRF) SetVulnerability(v []AdvisoryMVulnerability)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *AdvisoryMSCVRF) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


