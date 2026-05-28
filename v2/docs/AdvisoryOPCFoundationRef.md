# AdvisoryOPCFoundationRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**AdvisoryOPCFoundationDocumentMetadata**](AdvisoryOPCFoundationDocumentMetadata.md) |  | [optional] 
**ProductTree** | Pointer to [**AdvisoryProductBranch**](AdvisoryProductBranch.md) |  | [optional] 
**Vulnerabilities** | Pointer to [**[]AdvisoryOPCFoundationVulnerability**](AdvisoryOPCFoundationVulnerability.md) |  | [optional] 

## Methods

### NewAdvisoryOPCFoundationRef

`func NewAdvisoryOPCFoundationRef() *AdvisoryOPCFoundationRef`

NewAdvisoryOPCFoundationRef instantiates a new AdvisoryOPCFoundationRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOPCFoundationRefWithDefaults

`func NewAdvisoryOPCFoundationRefWithDefaults() *AdvisoryOPCFoundationRef`

NewAdvisoryOPCFoundationRefWithDefaults instantiates a new AdvisoryOPCFoundationRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *AdvisoryOPCFoundationRef) GetDocument() AdvisoryOPCFoundationDocumentMetadata`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AdvisoryOPCFoundationRef) GetDocumentOk() (*AdvisoryOPCFoundationDocumentMetadata, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AdvisoryOPCFoundationRef) SetDocument(v AdvisoryOPCFoundationDocumentMetadata)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AdvisoryOPCFoundationRef) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetProductTree

`func (o *AdvisoryOPCFoundationRef) GetProductTree() AdvisoryProductBranch`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *AdvisoryOPCFoundationRef) GetProductTreeOk() (*AdvisoryProductBranch, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *AdvisoryOPCFoundationRef) SetProductTree(v AdvisoryProductBranch)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *AdvisoryOPCFoundationRef) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisoryOPCFoundationRef) GetVulnerabilities() []AdvisoryOPCFoundationVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisoryOPCFoundationRef) GetVulnerabilitiesOk() (*[]AdvisoryOPCFoundationVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisoryOPCFoundationRef) SetVulnerabilities(v []AdvisoryOPCFoundationVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisoryOPCFoundationRef) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


