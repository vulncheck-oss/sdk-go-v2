# IndexCveSearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the unique document identifier within the index. | [optional] 
**Index** | Pointer to **string** | Index is the name of the index where this document was found. | [optional] 
**Score** | Pointer to **float32** | Score is the relevance score from OpenSearch (higher &#x3D; more relevant). | [optional] 
**Source** | Pointer to **map[string]interface{}** | Source is the full document content. Structure varies by index. | [optional] 

## Methods

### NewIndexCveSearchHit

`func NewIndexCveSearchHit() *IndexCveSearchHit`

NewIndexCveSearchHit instantiates a new IndexCveSearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexCveSearchHitWithDefaults

`func NewIndexCveSearchHitWithDefaults() *IndexCveSearchHit`

NewIndexCveSearchHitWithDefaults instantiates a new IndexCveSearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IndexCveSearchHit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IndexCveSearchHit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IndexCveSearchHit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IndexCveSearchHit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *IndexCveSearchHit) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *IndexCveSearchHit) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *IndexCveSearchHit) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *IndexCveSearchHit) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetScore

`func (o *IndexCveSearchHit) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IndexCveSearchHit) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IndexCveSearchHit) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *IndexCveSearchHit) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSource

`func (o *IndexCveSearchHit) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IndexCveSearchHit) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IndexCveSearchHit) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *IndexCveSearchHit) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


