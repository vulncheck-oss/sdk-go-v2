# AdvisoryAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affects** | Pointer to **string** |  | [optional] 
**Announced** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Corrections** | Pointer to [**[]AdvisoryCorrection**](AdvisoryCorrection.md) |  | [optional] 
**Credits** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Module** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAdvisory

`func NewAdvisoryAdvisory() *AdvisoryAdvisory`

NewAdvisoryAdvisory instantiates a new AdvisoryAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAdvisoryWithDefaults

`func NewAdvisoryAdvisoryWithDefaults() *AdvisoryAdvisory`

NewAdvisoryAdvisoryWithDefaults instantiates a new AdvisoryAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffects

`func (o *AdvisoryAdvisory) GetAffects() string`

GetAffects returns the Affects field if non-nil, zero value otherwise.

### GetAffectsOk

`func (o *AdvisoryAdvisory) GetAffectsOk() (*string, bool)`

GetAffectsOk returns a tuple with the Affects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffects

`func (o *AdvisoryAdvisory) SetAffects(v string)`

SetAffects sets Affects field to given value.

### HasAffects

`func (o *AdvisoryAdvisory) HasAffects() bool`

HasAffects returns a boolean if a field has been set.

### GetAnnounced

`func (o *AdvisoryAdvisory) GetAnnounced() string`

GetAnnounced returns the Announced field if non-nil, zero value otherwise.

### GetAnnouncedOk

`func (o *AdvisoryAdvisory) GetAnnouncedOk() (*string, bool)`

GetAnnouncedOk returns a tuple with the Announced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounced

`func (o *AdvisoryAdvisory) SetAnnounced(v string)`

SetAnnounced sets Announced field to given value.

### HasAnnounced

`func (o *AdvisoryAdvisory) HasAnnounced() bool`

HasAnnounced returns a boolean if a field has been set.

### GetCategory

`func (o *AdvisoryAdvisory) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryAdvisory) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryAdvisory) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryAdvisory) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCorrections

`func (o *AdvisoryAdvisory) GetCorrections() []AdvisoryCorrection`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *AdvisoryAdvisory) GetCorrectionsOk() (*[]AdvisoryCorrection, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *AdvisoryAdvisory) SetCorrections(v []AdvisoryCorrection)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *AdvisoryAdvisory) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryAdvisory) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryAdvisory) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryAdvisory) SetCredits(v string)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryAdvisory) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetModule

`func (o *AdvisoryAdvisory) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AdvisoryAdvisory) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AdvisoryAdvisory) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *AdvisoryAdvisory) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryAdvisory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryAdvisory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryAdvisory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryAdvisory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopic

`func (o *AdvisoryAdvisory) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *AdvisoryAdvisory) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *AdvisoryAdvisory) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *AdvisoryAdvisory) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


