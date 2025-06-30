# AdvisoryZDIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to [**AdvisoryZDIResponseVendor**](AdvisoryZDIResponseVendor.md) |  | [optional] 

## Methods

### NewAdvisoryZDIResponse

`func NewAdvisoryZDIResponse() *AdvisoryZDIResponse`

NewAdvisoryZDIResponse instantiates a new AdvisoryZDIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryZDIResponseWithDefaults

`func NewAdvisoryZDIResponseWithDefaults() *AdvisoryZDIResponse`

NewAdvisoryZDIResponseWithDefaults instantiates a new AdvisoryZDIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *AdvisoryZDIResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AdvisoryZDIResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AdvisoryZDIResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AdvisoryZDIResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUri

`func (o *AdvisoryZDIResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AdvisoryZDIResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AdvisoryZDIResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AdvisoryZDIResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryZDIResponse) GetVendor() AdvisoryZDIResponseVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryZDIResponse) GetVendorOk() (*AdvisoryZDIResponseVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryZDIResponse) SetVendor(v AdvisoryZDIResponseVendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryZDIResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


