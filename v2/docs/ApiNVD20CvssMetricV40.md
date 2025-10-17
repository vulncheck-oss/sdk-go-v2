# ApiNVD20CvssMetricV40

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssData** | Pointer to [**AdvisoryCVSSV40**](AdvisoryCVSSV40.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiNVD20CvssMetricV40

`func NewApiNVD20CvssMetricV40() *ApiNVD20CvssMetricV40`

NewApiNVD20CvssMetricV40 instantiates a new ApiNVD20CvssMetricV40 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CvssMetricV40WithDefaults

`func NewApiNVD20CvssMetricV40WithDefaults() *ApiNVD20CvssMetricV40`

NewApiNVD20CvssMetricV40WithDefaults instantiates a new ApiNVD20CvssMetricV40 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssData

`func (o *ApiNVD20CvssMetricV40) GetCvssData() AdvisoryCVSSV40`

GetCvssData returns the CvssData field if non-nil, zero value otherwise.

### GetCvssDataOk

`func (o *ApiNVD20CvssMetricV40) GetCvssDataOk() (*AdvisoryCVSSV40, bool)`

GetCvssDataOk returns a tuple with the CvssData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssData

`func (o *ApiNVD20CvssMetricV40) SetCvssData(v AdvisoryCVSSV40)`

SetCvssData sets CvssData field to given value.

### HasCvssData

`func (o *ApiNVD20CvssMetricV40) HasCvssData() bool`

HasCvssData returns a boolean if a field has been set.

### GetSource

`func (o *ApiNVD20CvssMetricV40) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiNVD20CvssMetricV40) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiNVD20CvssMetricV40) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiNVD20CvssMetricV40) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *ApiNVD20CvssMetricV40) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiNVD20CvssMetricV40) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiNVD20CvssMetricV40) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiNVD20CvssMetricV40) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


