# AdvisoryCSAF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**AdvisoryDocumentMetadata**](AdvisoryDocumentMetadata.md) | Document contains metadata about the CSAF document itself.  https://docs.oasis-open.org/csaf/csaf/v2.0/os/csaf-v2.0-os.html#321-document-property | [optional] 
**Notes** | Pointer to [**[]AdvisoryCSAFNote**](AdvisoryCSAFNote.md) | Notes holds notes associated with the whole document. https://docs.oasis-open.org/csaf/csaf/v2.0/os/csaf-v2.0-os.html#3217-document-property---notes | [optional] 
**ProductTree** | Pointer to [**AdvisoryProductBranch**](AdvisoryProductBranch.md) | ProductTree contains information about the product tree (branches only).  https://docs.oasis-open.org/csaf/csaf/v2.0/os/csaf-v2.0-os.html#322-product-tree-property | [optional] 
**Vulnerabilities** | Pointer to [**[]AdvisoryCSAFVulnerability**](AdvisoryCSAFVulnerability.md) | Vulnerabilities contains information about the vulnerabilities, (i.e. CVEs), associated threats, and product status.  https://docs.oasis-open.org/csaf/csaf/v2.0/os/csaf-v2.0-os.html#323-vulnerabilities-property | [optional] 

## Methods

### NewAdvisoryCSAF

`func NewAdvisoryCSAF() *AdvisoryCSAF`

NewAdvisoryCSAF instantiates a new AdvisoryCSAF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCSAFWithDefaults

`func NewAdvisoryCSAFWithDefaults() *AdvisoryCSAF`

NewAdvisoryCSAFWithDefaults instantiates a new AdvisoryCSAF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *AdvisoryCSAF) GetDocument() AdvisoryDocumentMetadata`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AdvisoryCSAF) GetDocumentOk() (*AdvisoryDocumentMetadata, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AdvisoryCSAF) SetDocument(v AdvisoryDocumentMetadata)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AdvisoryCSAF) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetNotes

`func (o *AdvisoryCSAF) GetNotes() []AdvisoryCSAFNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AdvisoryCSAF) GetNotesOk() (*[]AdvisoryCSAFNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AdvisoryCSAF) SetNotes(v []AdvisoryCSAFNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AdvisoryCSAF) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProductTree

`func (o *AdvisoryCSAF) GetProductTree() AdvisoryProductBranch`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *AdvisoryCSAF) GetProductTreeOk() (*AdvisoryProductBranch, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *AdvisoryCSAF) SetProductTree(v AdvisoryProductBranch)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *AdvisoryCSAF) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisoryCSAF) GetVulnerabilities() []AdvisoryCSAFVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisoryCSAF) GetVulnerabilitiesOk() (*[]AdvisoryCSAFVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisoryCSAF) SetVulnerabilities(v []AdvisoryCSAFVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisoryCSAF) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


