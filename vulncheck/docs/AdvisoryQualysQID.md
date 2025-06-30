# AdvisoryQualysQID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consequence** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssV2** | Pointer to [**[]AdvisoryCvsssV23**](AdvisoryCvsssV23.md) |  | [optional] 
**CvssV3** | Pointer to [**[]AdvisoryCvsssV23**](AdvisoryCvsssV23.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateInsert** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Patches** | Pointer to [**[]AdvisoryPatch**](AdvisoryPatch.md) |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**Qid** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VendorRefs** | Pointer to [**[]AdvisoryVendorRef**](AdvisoryVendorRef.md) |  | [optional] 

## Methods

### NewAdvisoryQualysQID

`func NewAdvisoryQualysQID() *AdvisoryQualysQID`

NewAdvisoryQualysQID instantiates a new AdvisoryQualysQID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryQualysQIDWithDefaults

`func NewAdvisoryQualysQIDWithDefaults() *AdvisoryQualysQID`

NewAdvisoryQualysQIDWithDefaults instantiates a new AdvisoryQualysQID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsequence

`func (o *AdvisoryQualysQID) GetConsequence() string`

GetConsequence returns the Consequence field if non-nil, zero value otherwise.

### GetConsequenceOk

`func (o *AdvisoryQualysQID) GetConsequenceOk() (*string, bool)`

GetConsequenceOk returns a tuple with the Consequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequence

`func (o *AdvisoryQualysQID) SetConsequence(v string)`

SetConsequence sets Consequence field to given value.

### HasConsequence

`func (o *AdvisoryQualysQID) HasConsequence() bool`

HasConsequence returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryQualysQID) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryQualysQID) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryQualysQID) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryQualysQID) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssV2

`func (o *AdvisoryQualysQID) GetCvssV2() []AdvisoryCvsssV23`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *AdvisoryQualysQID) GetCvssV2Ok() (*[]AdvisoryCvsssV23, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *AdvisoryQualysQID) SetCvssV2(v []AdvisoryCvsssV23)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *AdvisoryQualysQID) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetCvssV3

`func (o *AdvisoryQualysQID) GetCvssV3() []AdvisoryCvsssV23`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *AdvisoryQualysQID) GetCvssV3Ok() (*[]AdvisoryCvsssV23, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *AdvisoryQualysQID) SetCvssV3(v []AdvisoryCvsssV23)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *AdvisoryQualysQID) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryQualysQID) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryQualysQID) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryQualysQID) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryQualysQID) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateInsert

`func (o *AdvisoryQualysQID) GetDateInsert() string`

GetDateInsert returns the DateInsert field if non-nil, zero value otherwise.

### GetDateInsertOk

`func (o *AdvisoryQualysQID) GetDateInsertOk() (*string, bool)`

GetDateInsertOk returns a tuple with the DateInsert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInsert

`func (o *AdvisoryQualysQID) SetDateInsert(v string)`

SetDateInsert sets DateInsert field to given value.

### HasDateInsert

`func (o *AdvisoryQualysQID) HasDateInsert() bool`

HasDateInsert returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryQualysQID) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryQualysQID) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryQualysQID) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryQualysQID) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPatches

`func (o *AdvisoryQualysQID) GetPatches() []AdvisoryPatch`

GetPatches returns the Patches field if non-nil, zero value otherwise.

### GetPatchesOk

`func (o *AdvisoryQualysQID) GetPatchesOk() (*[]AdvisoryPatch, bool)`

GetPatchesOk returns a tuple with the Patches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatches

`func (o *AdvisoryQualysQID) SetPatches(v []AdvisoryPatch)`

SetPatches sets Patches field to given value.

### HasPatches

`func (o *AdvisoryQualysQID) HasPatches() bool`

HasPatches returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryQualysQID) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryQualysQID) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryQualysQID) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryQualysQID) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetQid

`func (o *AdvisoryQualysQID) GetQid() string`

GetQid returns the Qid field if non-nil, zero value otherwise.

### GetQidOk

`func (o *AdvisoryQualysQID) GetQidOk() (*string, bool)`

GetQidOk returns a tuple with the Qid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQid

`func (o *AdvisoryQualysQID) SetQid(v string)`

SetQid sets Qid field to given value.

### HasQid

`func (o *AdvisoryQualysQID) HasQid() bool`

HasQid returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryQualysQID) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryQualysQID) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryQualysQID) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryQualysQID) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSolution

`func (o *AdvisoryQualysQID) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *AdvisoryQualysQID) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *AdvisoryQualysQID) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *AdvisoryQualysQID) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryQualysQID) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryQualysQID) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryQualysQID) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryQualysQID) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryQualysQID) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryQualysQID) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryQualysQID) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryQualysQID) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendorRefs

`func (o *AdvisoryQualysQID) GetVendorRefs() []AdvisoryVendorRef`

GetVendorRefs returns the VendorRefs field if non-nil, zero value otherwise.

### GetVendorRefsOk

`func (o *AdvisoryQualysQID) GetVendorRefsOk() (*[]AdvisoryVendorRef, bool)`

GetVendorRefsOk returns a tuple with the VendorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorRefs

`func (o *AdvisoryQualysQID) SetVendorRefs(v []AdvisoryVendorRef)`

SetVendorRefs sets VendorRefs field to given value.

### HasVendorRefs

`func (o *AdvisoryQualysQID) HasVendorRefs() bool`

HasVendorRefs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


