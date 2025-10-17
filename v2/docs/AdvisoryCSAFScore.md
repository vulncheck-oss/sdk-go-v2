# AdvisoryCSAFScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssV2** | Pointer to [**AdvisoryCVSSV2**](AdvisoryCVSSV2.md) |  | [optional] 
**CvssV3** | Pointer to [**AdvisoryCVSSV3**](AdvisoryCVSSV3.md) |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryCSAFScore

`func NewAdvisoryCSAFScore() *AdvisoryCSAFScore`

NewAdvisoryCSAFScore instantiates a new AdvisoryCSAFScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCSAFScoreWithDefaults

`func NewAdvisoryCSAFScoreWithDefaults() *AdvisoryCSAFScore`

NewAdvisoryCSAFScoreWithDefaults instantiates a new AdvisoryCSAFScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssV2

`func (o *AdvisoryCSAFScore) GetCvssV2() AdvisoryCVSSV2`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *AdvisoryCSAFScore) GetCvssV2Ok() (*AdvisoryCVSSV2, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *AdvisoryCSAFScore) SetCvssV2(v AdvisoryCVSSV2)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *AdvisoryCSAFScore) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetCvssV3

`func (o *AdvisoryCSAFScore) GetCvssV3() AdvisoryCVSSV3`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *AdvisoryCSAFScore) GetCvssV3Ok() (*AdvisoryCVSSV3, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *AdvisoryCSAFScore) SetCvssV3(v AdvisoryCVSSV3)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *AdvisoryCSAFScore) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisoryCSAFScore) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisoryCSAFScore) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisoryCSAFScore) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisoryCSAFScore) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


