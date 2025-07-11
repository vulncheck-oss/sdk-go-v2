/*
VulnCheck API

Version 3 of the VulnCheck API

API version: 3.0
Contact: support@vulncheck.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vulncheck

import (
	"encoding/json"
)

// checks if the AdvisoryManageEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvisoryManageEngine{}

// AdvisoryManageEngine struct for AdvisoryManageEngine
type AdvisoryManageEngine struct {
	ADVISORY *string `json:"ADVISORY,omitempty"`
	AddedTime *string `json:"Added_Time,omitempty"`
	CVEDetailsLink *AdvisoryCVEDetailsLink `json:"CVE_Details_Link,omitempty"`
	CVE_ID *string `json:"CVE_ID,omitempty"`
	CVSSSeverityRating *string `json:"CVSS_Severity_Rating,omitempty"`
	Fixed *string `json:"Fixed,omitempty"`
	ForProductSearch *string `json:"For_product_search,omitempty"`
	ID *string `json:"ID,omitempty"`
	Product *AdvisoryMEProduct `json:"Product,omitempty"`
	ProductList []AdvisoryMEProduct `json:"Product_list,omitempty"`
	ProductSpecificDetails []AdvisoryProductSpecificDetail `json:"Product_specific_details,omitempty"`
	Summary *string `json:"Summary,omitempty"`
	Version *string `json:"Version,omitempty"`
	IndexField *string `json:"index_field,omitempty"`
}

// NewAdvisoryManageEngine instantiates a new AdvisoryManageEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvisoryManageEngine() *AdvisoryManageEngine {
	this := AdvisoryManageEngine{}
	return &this
}

// NewAdvisoryManageEngineWithDefaults instantiates a new AdvisoryManageEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvisoryManageEngineWithDefaults() *AdvisoryManageEngine {
	this := AdvisoryManageEngine{}
	return &this
}

// GetADVISORY returns the ADVISORY field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetADVISORY() string {
	if o == nil || IsNil(o.ADVISORY) {
		var ret string
		return ret
	}
	return *o.ADVISORY
}

// GetADVISORYOk returns a tuple with the ADVISORY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetADVISORYOk() (*string, bool) {
	if o == nil || IsNil(o.ADVISORY) {
		return nil, false
	}
	return o.ADVISORY, true
}

// HasADVISORY returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasADVISORY() bool {
	if o != nil && !IsNil(o.ADVISORY) {
		return true
	}

	return false
}

// SetADVISORY gets a reference to the given string and assigns it to the ADVISORY field.
func (o *AdvisoryManageEngine) SetADVISORY(v string) {
	o.ADVISORY = &v
}

// GetAddedTime returns the AddedTime field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetAddedTime() string {
	if o == nil || IsNil(o.AddedTime) {
		var ret string
		return ret
	}
	return *o.AddedTime
}

// GetAddedTimeOk returns a tuple with the AddedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetAddedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddedTime) {
		return nil, false
	}
	return o.AddedTime, true
}

// HasAddedTime returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasAddedTime() bool {
	if o != nil && !IsNil(o.AddedTime) {
		return true
	}

	return false
}

// SetAddedTime gets a reference to the given string and assigns it to the AddedTime field.
func (o *AdvisoryManageEngine) SetAddedTime(v string) {
	o.AddedTime = &v
}

// GetCVEDetailsLink returns the CVEDetailsLink field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetCVEDetailsLink() AdvisoryCVEDetailsLink {
	if o == nil || IsNil(o.CVEDetailsLink) {
		var ret AdvisoryCVEDetailsLink
		return ret
	}
	return *o.CVEDetailsLink
}

// GetCVEDetailsLinkOk returns a tuple with the CVEDetailsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetCVEDetailsLinkOk() (*AdvisoryCVEDetailsLink, bool) {
	if o == nil || IsNil(o.CVEDetailsLink) {
		return nil, false
	}
	return o.CVEDetailsLink, true
}

// HasCVEDetailsLink returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasCVEDetailsLink() bool {
	if o != nil && !IsNil(o.CVEDetailsLink) {
		return true
	}

	return false
}

// SetCVEDetailsLink gets a reference to the given AdvisoryCVEDetailsLink and assigns it to the CVEDetailsLink field.
func (o *AdvisoryManageEngine) SetCVEDetailsLink(v AdvisoryCVEDetailsLink) {
	o.CVEDetailsLink = &v
}

// GetCVE_ID returns the CVE_ID field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetCVE_ID() string {
	if o == nil || IsNil(o.CVE_ID) {
		var ret string
		return ret
	}
	return *o.CVE_ID
}

// GetCVE_IDOk returns a tuple with the CVE_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetCVE_IDOk() (*string, bool) {
	if o == nil || IsNil(o.CVE_ID) {
		return nil, false
	}
	return o.CVE_ID, true
}

// HasCVE_ID returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasCVE_ID() bool {
	if o != nil && !IsNil(o.CVE_ID) {
		return true
	}

	return false
}

// SetCVE_ID gets a reference to the given string and assigns it to the CVE_ID field.
func (o *AdvisoryManageEngine) SetCVE_ID(v string) {
	o.CVE_ID = &v
}

// GetCVSSSeverityRating returns the CVSSSeverityRating field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetCVSSSeverityRating() string {
	if o == nil || IsNil(o.CVSSSeverityRating) {
		var ret string
		return ret
	}
	return *o.CVSSSeverityRating
}

// GetCVSSSeverityRatingOk returns a tuple with the CVSSSeverityRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetCVSSSeverityRatingOk() (*string, bool) {
	if o == nil || IsNil(o.CVSSSeverityRating) {
		return nil, false
	}
	return o.CVSSSeverityRating, true
}

// HasCVSSSeverityRating returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasCVSSSeverityRating() bool {
	if o != nil && !IsNil(o.CVSSSeverityRating) {
		return true
	}

	return false
}

// SetCVSSSeverityRating gets a reference to the given string and assigns it to the CVSSSeverityRating field.
func (o *AdvisoryManageEngine) SetCVSSSeverityRating(v string) {
	o.CVSSSeverityRating = &v
}

// GetFixed returns the Fixed field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetFixed() string {
	if o == nil || IsNil(o.Fixed) {
		var ret string
		return ret
	}
	return *o.Fixed
}

// GetFixedOk returns a tuple with the Fixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetFixedOk() (*string, bool) {
	if o == nil || IsNil(o.Fixed) {
		return nil, false
	}
	return o.Fixed, true
}

// HasFixed returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasFixed() bool {
	if o != nil && !IsNil(o.Fixed) {
		return true
	}

	return false
}

// SetFixed gets a reference to the given string and assigns it to the Fixed field.
func (o *AdvisoryManageEngine) SetFixed(v string) {
	o.Fixed = &v
}

// GetForProductSearch returns the ForProductSearch field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetForProductSearch() string {
	if o == nil || IsNil(o.ForProductSearch) {
		var ret string
		return ret
	}
	return *o.ForProductSearch
}

// GetForProductSearchOk returns a tuple with the ForProductSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetForProductSearchOk() (*string, bool) {
	if o == nil || IsNil(o.ForProductSearch) {
		return nil, false
	}
	return o.ForProductSearch, true
}

// HasForProductSearch returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasForProductSearch() bool {
	if o != nil && !IsNil(o.ForProductSearch) {
		return true
	}

	return false
}

// SetForProductSearch gets a reference to the given string and assigns it to the ForProductSearch field.
func (o *AdvisoryManageEngine) SetForProductSearch(v string) {
	o.ForProductSearch = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetID() string {
	if o == nil || IsNil(o.ID) {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetIDOk() (*string, bool) {
	if o == nil || IsNil(o.ID) {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasID() bool {
	if o != nil && !IsNil(o.ID) {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *AdvisoryManageEngine) SetID(v string) {
	o.ID = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetProduct() AdvisoryMEProduct {
	if o == nil || IsNil(o.Product) {
		var ret AdvisoryMEProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetProductOk() (*AdvisoryMEProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AdvisoryMEProduct and assigns it to the Product field.
func (o *AdvisoryManageEngine) SetProduct(v AdvisoryMEProduct) {
	o.Product = &v
}

// GetProductList returns the ProductList field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetProductList() []AdvisoryMEProduct {
	if o == nil || IsNil(o.ProductList) {
		var ret []AdvisoryMEProduct
		return ret
	}
	return o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetProductListOk() ([]AdvisoryMEProduct, bool) {
	if o == nil || IsNil(o.ProductList) {
		return nil, false
	}
	return o.ProductList, true
}

// HasProductList returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasProductList() bool {
	if o != nil && !IsNil(o.ProductList) {
		return true
	}

	return false
}

// SetProductList gets a reference to the given []AdvisoryMEProduct and assigns it to the ProductList field.
func (o *AdvisoryManageEngine) SetProductList(v []AdvisoryMEProduct) {
	o.ProductList = v
}

// GetProductSpecificDetails returns the ProductSpecificDetails field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetProductSpecificDetails() []AdvisoryProductSpecificDetail {
	if o == nil || IsNil(o.ProductSpecificDetails) {
		var ret []AdvisoryProductSpecificDetail
		return ret
	}
	return o.ProductSpecificDetails
}

// GetProductSpecificDetailsOk returns a tuple with the ProductSpecificDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetProductSpecificDetailsOk() ([]AdvisoryProductSpecificDetail, bool) {
	if o == nil || IsNil(o.ProductSpecificDetails) {
		return nil, false
	}
	return o.ProductSpecificDetails, true
}

// HasProductSpecificDetails returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasProductSpecificDetails() bool {
	if o != nil && !IsNil(o.ProductSpecificDetails) {
		return true
	}

	return false
}

// SetProductSpecificDetails gets a reference to the given []AdvisoryProductSpecificDetail and assigns it to the ProductSpecificDetails field.
func (o *AdvisoryManageEngine) SetProductSpecificDetails(v []AdvisoryProductSpecificDetail) {
	o.ProductSpecificDetails = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AdvisoryManageEngine) SetSummary(v string) {
	o.Summary = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AdvisoryManageEngine) SetVersion(v string) {
	o.Version = &v
}

// GetIndexField returns the IndexField field value if set, zero value otherwise.
func (o *AdvisoryManageEngine) GetIndexField() string {
	if o == nil || IsNil(o.IndexField) {
		var ret string
		return ret
	}
	return *o.IndexField
}

// GetIndexFieldOk returns a tuple with the IndexField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvisoryManageEngine) GetIndexFieldOk() (*string, bool) {
	if o == nil || IsNil(o.IndexField) {
		return nil, false
	}
	return o.IndexField, true
}

// HasIndexField returns a boolean if a field has been set.
func (o *AdvisoryManageEngine) HasIndexField() bool {
	if o != nil && !IsNil(o.IndexField) {
		return true
	}

	return false
}

// SetIndexField gets a reference to the given string and assigns it to the IndexField field.
func (o *AdvisoryManageEngine) SetIndexField(v string) {
	o.IndexField = &v
}

func (o AdvisoryManageEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvisoryManageEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ADVISORY) {
		toSerialize["ADVISORY"] = o.ADVISORY
	}
	if !IsNil(o.AddedTime) {
		toSerialize["Added_Time"] = o.AddedTime
	}
	if !IsNil(o.CVEDetailsLink) {
		toSerialize["CVE_Details_Link"] = o.CVEDetailsLink
	}
	if !IsNil(o.CVE_ID) {
		toSerialize["CVE_ID"] = o.CVE_ID
	}
	if !IsNil(o.CVSSSeverityRating) {
		toSerialize["CVSS_Severity_Rating"] = o.CVSSSeverityRating
	}
	if !IsNil(o.Fixed) {
		toSerialize["Fixed"] = o.Fixed
	}
	if !IsNil(o.ForProductSearch) {
		toSerialize["For_product_search"] = o.ForProductSearch
	}
	if !IsNil(o.ID) {
		toSerialize["ID"] = o.ID
	}
	if !IsNil(o.Product) {
		toSerialize["Product"] = o.Product
	}
	if !IsNil(o.ProductList) {
		toSerialize["Product_list"] = o.ProductList
	}
	if !IsNil(o.ProductSpecificDetails) {
		toSerialize["Product_specific_details"] = o.ProductSpecificDetails
	}
	if !IsNil(o.Summary) {
		toSerialize["Summary"] = o.Summary
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.IndexField) {
		toSerialize["index_field"] = o.IndexField
	}
	return toSerialize, nil
}

type NullableAdvisoryManageEngine struct {
	value *AdvisoryManageEngine
	isSet bool
}

func (v NullableAdvisoryManageEngine) Get() *AdvisoryManageEngine {
	return v.value
}

func (v *NullableAdvisoryManageEngine) Set(val *AdvisoryManageEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvisoryManageEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvisoryManageEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvisoryManageEngine(val *AdvisoryManageEngine) *NullableAdvisoryManageEngine {
	return &NullableAdvisoryManageEngine{value: val, isSet: true}
}

func (v NullableAdvisoryManageEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvisoryManageEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


