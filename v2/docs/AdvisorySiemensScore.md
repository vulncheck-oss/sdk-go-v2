# AdvisorySiemensScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssV3** | Pointer to [**AdvisorySiemensCVSSV3**](AdvisorySiemensCVSSV3.md) |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisorySiemensScore

`func NewAdvisorySiemensScore() *AdvisorySiemensScore`

NewAdvisorySiemensScore instantiates a new AdvisorySiemensScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySiemensScoreWithDefaults

`func NewAdvisorySiemensScoreWithDefaults() *AdvisorySiemensScore`

NewAdvisorySiemensScoreWithDefaults instantiates a new AdvisorySiemensScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssV3

`func (o *AdvisorySiemensScore) GetCvssV3() AdvisorySiemensCVSSV3`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *AdvisorySiemensScore) GetCvssV3Ok() (*AdvisorySiemensCVSSV3, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *AdvisorySiemensScore) SetCvssV3(v AdvisorySiemensCVSSV3)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *AdvisorySiemensScore) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.

### GetProducts

`func (o *AdvisorySiemensScore) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *AdvisorySiemensScore) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *AdvisorySiemensScore) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *AdvisorySiemensScore) HasProducts() bool`

HasProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


