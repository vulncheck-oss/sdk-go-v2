# AdvisoryGreyNoiseDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Intention** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RecommendBlock** | Pointer to **bool** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**RelatedTags** | Pointer to [**[]AdvisoryGreyNoiseTags**](AdvisoryGreyNoiseTags.md) |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGreyNoiseDetection

`func NewAdvisoryGreyNoiseDetection() *AdvisoryGreyNoiseDetection`

NewAdvisoryGreyNoiseDetection instantiates a new AdvisoryGreyNoiseDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGreyNoiseDetectionWithDefaults

`func NewAdvisoryGreyNoiseDetectionWithDefaults() *AdvisoryGreyNoiseDetection`

NewAdvisoryGreyNoiseDetectionWithDefaults instantiates a new AdvisoryGreyNoiseDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AdvisoryGreyNoiseDetection) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AdvisoryGreyNoiseDetection) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AdvisoryGreyNoiseDetection) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AdvisoryGreyNoiseDetection) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryGreyNoiseDetection) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGreyNoiseDetection) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGreyNoiseDetection) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGreyNoiseDetection) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGreyNoiseDetection) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGreyNoiseDetection) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGreyNoiseDetection) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGreyNoiseDetection) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryGreyNoiseDetection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryGreyNoiseDetection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryGreyNoiseDetection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryGreyNoiseDetection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryGreyNoiseDetection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryGreyNoiseDetection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryGreyNoiseDetection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryGreyNoiseDetection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntention

`func (o *AdvisoryGreyNoiseDetection) GetIntention() string`

GetIntention returns the Intention field if non-nil, zero value otherwise.

### GetIntentionOk

`func (o *AdvisoryGreyNoiseDetection) GetIntentionOk() (*string, bool)`

GetIntentionOk returns a tuple with the Intention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntention

`func (o *AdvisoryGreyNoiseDetection) SetIntention(v string)`

SetIntention sets Intention field to given value.

### HasIntention

`func (o *AdvisoryGreyNoiseDetection) HasIntention() bool`

HasIntention returns a boolean if a field has been set.

### GetLabel

`func (o *AdvisoryGreyNoiseDetection) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AdvisoryGreyNoiseDetection) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AdvisoryGreyNoiseDetection) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AdvisoryGreyNoiseDetection) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryGreyNoiseDetection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryGreyNoiseDetection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryGreyNoiseDetection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryGreyNoiseDetection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecommendBlock

`func (o *AdvisoryGreyNoiseDetection) GetRecommendBlock() bool`

GetRecommendBlock returns the RecommendBlock field if non-nil, zero value otherwise.

### GetRecommendBlockOk

`func (o *AdvisoryGreyNoiseDetection) GetRecommendBlockOk() (*bool, bool)`

GetRecommendBlockOk returns a tuple with the RecommendBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendBlock

`func (o *AdvisoryGreyNoiseDetection) SetRecommendBlock(v bool)`

SetRecommendBlock sets RecommendBlock field to given value.

### HasRecommendBlock

`func (o *AdvisoryGreyNoiseDetection) HasRecommendBlock() bool`

HasRecommendBlock returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryGreyNoiseDetection) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryGreyNoiseDetection) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryGreyNoiseDetection) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryGreyNoiseDetection) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedTags

`func (o *AdvisoryGreyNoiseDetection) GetRelatedTags() []AdvisoryGreyNoiseTags`

GetRelatedTags returns the RelatedTags field if non-nil, zero value otherwise.

### GetRelatedTagsOk

`func (o *AdvisoryGreyNoiseDetection) GetRelatedTagsOk() (*[]AdvisoryGreyNoiseTags, bool)`

GetRelatedTagsOk returns a tuple with the RelatedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTags

`func (o *AdvisoryGreyNoiseDetection) SetRelatedTags(v []AdvisoryGreyNoiseTags)`

SetRelatedTags sets RelatedTags field to given value.

### HasRelatedTags

`func (o *AdvisoryGreyNoiseDetection) HasRelatedTags() bool`

HasRelatedTags returns a boolean if a field has been set.

### GetSlug

`func (o *AdvisoryGreyNoiseDetection) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *AdvisoryGreyNoiseDetection) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *AdvisoryGreyNoiseDetection) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *AdvisoryGreyNoiseDetection) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryGreyNoiseDetection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryGreyNoiseDetection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryGreyNoiseDetection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryGreyNoiseDetection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


