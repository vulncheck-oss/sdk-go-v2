# AdvisoryDocumentPublisher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactDetails** | Pointer to **string** |  | [optional] 
**IssuingAuthority** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | the json for this is missing/broke | [optional] 

## Methods

### NewAdvisoryDocumentPublisher

`func NewAdvisoryDocumentPublisher() *AdvisoryDocumentPublisher`

NewAdvisoryDocumentPublisher instantiates a new AdvisoryDocumentPublisher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDocumentPublisherWithDefaults

`func NewAdvisoryDocumentPublisherWithDefaults() *AdvisoryDocumentPublisher`

NewAdvisoryDocumentPublisherWithDefaults instantiates a new AdvisoryDocumentPublisher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactDetails

`func (o *AdvisoryDocumentPublisher) GetContactDetails() string`

GetContactDetails returns the ContactDetails field if non-nil, zero value otherwise.

### GetContactDetailsOk

`func (o *AdvisoryDocumentPublisher) GetContactDetailsOk() (*string, bool)`

GetContactDetailsOk returns a tuple with the ContactDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactDetails

`func (o *AdvisoryDocumentPublisher) SetContactDetails(v string)`

SetContactDetails sets ContactDetails field to given value.

### HasContactDetails

`func (o *AdvisoryDocumentPublisher) HasContactDetails() bool`

HasContactDetails returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *AdvisoryDocumentPublisher) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *AdvisoryDocumentPublisher) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *AdvisoryDocumentPublisher) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *AdvisoryDocumentPublisher) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryDocumentPublisher) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryDocumentPublisher) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryDocumentPublisher) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryDocumentPublisher) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


