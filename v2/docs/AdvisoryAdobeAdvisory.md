# AdvisoryAdobeAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdobeCves** | Pointer to [**[]AdvisoryAdobeCVE**](AdvisoryAdobeCVE.md) |  | [optional] 
**Affected** | Pointer to [**[]AdvisoryAdobeAffected**](AdvisoryAdobeAffected.md) |  | [optional] 
**BulletinId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Solutions** | Pointer to [**[]AdvisoryAdobeSolution**](AdvisoryAdobeSolution.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAdobeAdvisory

`func NewAdvisoryAdobeAdvisory() *AdvisoryAdobeAdvisory`

NewAdvisoryAdobeAdvisory instantiates a new AdvisoryAdobeAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAdobeAdvisoryWithDefaults

`func NewAdvisoryAdobeAdvisoryWithDefaults() *AdvisoryAdobeAdvisory`

NewAdvisoryAdobeAdvisoryWithDefaults instantiates a new AdvisoryAdobeAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdobeCves

`func (o *AdvisoryAdobeAdvisory) GetAdobeCves() []AdvisoryAdobeCVE`

GetAdobeCves returns the AdobeCves field if non-nil, zero value otherwise.

### GetAdobeCvesOk

`func (o *AdvisoryAdobeAdvisory) GetAdobeCvesOk() (*[]AdvisoryAdobeCVE, bool)`

GetAdobeCvesOk returns a tuple with the AdobeCves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdobeCves

`func (o *AdvisoryAdobeAdvisory) SetAdobeCves(v []AdvisoryAdobeCVE)`

SetAdobeCves sets AdobeCves field to given value.

### HasAdobeCves

`func (o *AdvisoryAdobeAdvisory) HasAdobeCves() bool`

HasAdobeCves returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryAdobeAdvisory) GetAffected() []AdvisoryAdobeAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryAdobeAdvisory) GetAffectedOk() (*[]AdvisoryAdobeAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryAdobeAdvisory) SetAffected(v []AdvisoryAdobeAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryAdobeAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetBulletinId

`func (o *AdvisoryAdobeAdvisory) GetBulletinId() string`

GetBulletinId returns the BulletinId field if non-nil, zero value otherwise.

### GetBulletinIdOk

`func (o *AdvisoryAdobeAdvisory) GetBulletinIdOk() (*string, bool)`

GetBulletinIdOk returns a tuple with the BulletinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinId

`func (o *AdvisoryAdobeAdvisory) SetBulletinId(v string)`

SetBulletinId sets BulletinId field to given value.

### HasBulletinId

`func (o *AdvisoryAdobeAdvisory) HasBulletinId() bool`

HasBulletinId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAdobeAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAdobeAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAdobeAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAdobeAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAdobeAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAdobeAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAdobeAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAdobeAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryAdobeAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryAdobeAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryAdobeAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryAdobeAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryAdobeAdvisory) GetSolutions() []AdvisoryAdobeSolution`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryAdobeAdvisory) GetSolutionsOk() (*[]AdvisoryAdobeSolution, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryAdobeAdvisory) SetSolutions(v []AdvisoryAdobeSolution)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryAdobeAdvisory) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryAdobeAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryAdobeAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryAdobeAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryAdobeAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


