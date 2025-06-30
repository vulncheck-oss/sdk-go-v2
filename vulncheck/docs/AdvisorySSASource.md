# AdvisorySSASource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**AdvisorySiemensDocument**](AdvisorySiemensDocument.md) |  | [optional] 
**ProductTree** | Pointer to [**AdvisorySiemensProductTree**](AdvisorySiemensProductTree.md) |  | [optional] 
**Vulnerabilities** | Pointer to [**[]AdvisorySiemensVulnerability**](AdvisorySiemensVulnerability.md) |  | [optional] 

## Methods

### NewAdvisorySSASource

`func NewAdvisorySSASource() *AdvisorySSASource`

NewAdvisorySSASource instantiates a new AdvisorySSASource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySSASourceWithDefaults

`func NewAdvisorySSASourceWithDefaults() *AdvisorySSASource`

NewAdvisorySSASourceWithDefaults instantiates a new AdvisorySSASource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *AdvisorySSASource) GetDocument() AdvisorySiemensDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AdvisorySSASource) GetDocumentOk() (*AdvisorySiemensDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AdvisorySSASource) SetDocument(v AdvisorySiemensDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AdvisorySSASource) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetProductTree

`func (o *AdvisorySSASource) GetProductTree() AdvisorySiemensProductTree`

GetProductTree returns the ProductTree field if non-nil, zero value otherwise.

### GetProductTreeOk

`func (o *AdvisorySSASource) GetProductTreeOk() (*AdvisorySiemensProductTree, bool)`

GetProductTreeOk returns a tuple with the ProductTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTree

`func (o *AdvisorySSASource) SetProductTree(v AdvisorySiemensProductTree)`

SetProductTree sets ProductTree field to given value.

### HasProductTree

`func (o *AdvisorySSASource) HasProductTree() bool`

HasProductTree returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisorySSASource) GetVulnerabilities() []AdvisorySiemensVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisorySSASource) GetVulnerabilitiesOk() (*[]AdvisorySiemensVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisorySSASource) SetVulnerabilities(v []AdvisorySiemensVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisorySSASource) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


