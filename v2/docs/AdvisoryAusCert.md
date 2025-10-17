# AdvisoryAusCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**BulletinId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Publisher** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAusCert

`func NewAdvisoryAusCert() *AdvisoryAusCert`

NewAdvisoryAusCert instantiates a new AdvisoryAusCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAusCertWithDefaults

`func NewAdvisoryAusCertWithDefaults() *AdvisoryAusCert`

NewAdvisoryAusCertWithDefaults instantiates a new AdvisoryAusCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *AdvisoryAusCert) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AdvisoryAusCert) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AdvisoryAusCert) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AdvisoryAusCert) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBulletinId

`func (o *AdvisoryAusCert) GetBulletinId() string`

GetBulletinId returns the BulletinId field if non-nil, zero value otherwise.

### GetBulletinIdOk

`func (o *AdvisoryAusCert) GetBulletinIdOk() (*string, bool)`

GetBulletinIdOk returns a tuple with the BulletinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinId

`func (o *AdvisoryAusCert) SetBulletinId(v string)`

SetBulletinId sets BulletinId field to given value.

### HasBulletinId

`func (o *AdvisoryAusCert) HasBulletinId() bool`

HasBulletinId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAusCert) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAusCert) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAusCert) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAusCert) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryAusCert) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryAusCert) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryAusCert) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryAusCert) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAusCert) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAusCert) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAusCert) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAusCert) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryAusCert) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryAusCert) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryAusCert) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryAusCert) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *AdvisoryAusCert) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *AdvisoryAusCert) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *AdvisoryAusCert) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *AdvisoryAusCert) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryAusCert) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryAusCert) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryAusCert) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryAusCert) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetPublisher

`func (o *AdvisoryAusCert) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AdvisoryAusCert) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AdvisoryAusCert) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AdvisoryAusCert) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetResolution

`func (o *AdvisoryAusCert) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AdvisoryAusCert) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AdvisoryAusCert) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AdvisoryAusCert) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


