/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1OrganismQueryRequest struct for V1OrganismQueryRequest
type V1OrganismQueryRequest struct {
	OrganismQuery *string `json:"organism_query,omitempty"`
	TaxonQuery *string `json:"taxon_query,omitempty"`
	TaxRankFilter *V1OrganismQueryRequestTaxRankFilter `json:"tax_rank_filter,omitempty"`
}

// NewV1OrganismQueryRequest instantiates a new V1OrganismQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1OrganismQueryRequest() *V1OrganismQueryRequest {
	this := V1OrganismQueryRequest{}
	var taxRankFilter V1OrganismQueryRequestTaxRankFilter = V1ORGANISMQUERYREQUESTTAXRANKFILTER_SPECIES
	this.TaxRankFilter = &taxRankFilter
	return &this
}

// NewV1OrganismQueryRequestWithDefaults instantiates a new V1OrganismQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1OrganismQueryRequestWithDefaults() *V1OrganismQueryRequest {
	this := V1OrganismQueryRequest{}
	var taxRankFilter V1OrganismQueryRequestTaxRankFilter = V1ORGANISMQUERYREQUESTTAXRANKFILTER_SPECIES
	this.TaxRankFilter = &taxRankFilter
	return &this
}

// GetOrganismQuery returns the OrganismQuery field value if set, zero value otherwise.
func (o *V1OrganismQueryRequest) GetOrganismQuery() string {
	if o == nil || o.OrganismQuery == nil {
		var ret string
		return ret
	}
	return *o.OrganismQuery
}

// GetOrganismQueryOk returns a tuple with the OrganismQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1OrganismQueryRequest) GetOrganismQueryOk() (*string, bool) {
	if o == nil || o.OrganismQuery == nil {
		return nil, false
	}
	return o.OrganismQuery, true
}

// HasOrganismQuery returns a boolean if a field has been set.
func (o *V1OrganismQueryRequest) HasOrganismQuery() bool {
	if o != nil && o.OrganismQuery != nil {
		return true
	}

	return false
}

// SetOrganismQuery gets a reference to the given string and assigns it to the OrganismQuery field.
func (o *V1OrganismQueryRequest) SetOrganismQuery(v string) {
	o.OrganismQuery = &v
}

// GetTaxonQuery returns the TaxonQuery field value if set, zero value otherwise.
func (o *V1OrganismQueryRequest) GetTaxonQuery() string {
	if o == nil || o.TaxonQuery == nil {
		var ret string
		return ret
	}
	return *o.TaxonQuery
}

// GetTaxonQueryOk returns a tuple with the TaxonQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1OrganismQueryRequest) GetTaxonQueryOk() (*string, bool) {
	if o == nil || o.TaxonQuery == nil {
		return nil, false
	}
	return o.TaxonQuery, true
}

// HasTaxonQuery returns a boolean if a field has been set.
func (o *V1OrganismQueryRequest) HasTaxonQuery() bool {
	if o != nil && o.TaxonQuery != nil {
		return true
	}

	return false
}

// SetTaxonQuery gets a reference to the given string and assigns it to the TaxonQuery field.
func (o *V1OrganismQueryRequest) SetTaxonQuery(v string) {
	o.TaxonQuery = &v
}

// GetTaxRankFilter returns the TaxRankFilter field value if set, zero value otherwise.
func (o *V1OrganismQueryRequest) GetTaxRankFilter() V1OrganismQueryRequestTaxRankFilter {
	if o == nil || o.TaxRankFilter == nil {
		var ret V1OrganismQueryRequestTaxRankFilter
		return ret
	}
	return *o.TaxRankFilter
}

// GetTaxRankFilterOk returns a tuple with the TaxRankFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1OrganismQueryRequest) GetTaxRankFilterOk() (*V1OrganismQueryRequestTaxRankFilter, bool) {
	if o == nil || o.TaxRankFilter == nil {
		return nil, false
	}
	return o.TaxRankFilter, true
}

// HasTaxRankFilter returns a boolean if a field has been set.
func (o *V1OrganismQueryRequest) HasTaxRankFilter() bool {
	if o != nil && o.TaxRankFilter != nil {
		return true
	}

	return false
}

// SetTaxRankFilter gets a reference to the given V1OrganismQueryRequestTaxRankFilter and assigns it to the TaxRankFilter field.
func (o *V1OrganismQueryRequest) SetTaxRankFilter(v V1OrganismQueryRequestTaxRankFilter) {
	o.TaxRankFilter = &v
}

func (o V1OrganismQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganismQuery != nil  {
		toSerialize["organism_query"] = o.OrganismQuery
	}
	if o.TaxonQuery != nil  {
		toSerialize["taxon_query"] = o.TaxonQuery
	}
	if o.TaxRankFilter != nil  {
		toSerialize["tax_rank_filter"] = o.TaxRankFilter
	}
	return json.Marshal(toSerialize)
}

type NullableV1OrganismQueryRequest struct {
	value *V1OrganismQueryRequest
	isSet bool
}

func (v NullableV1OrganismQueryRequest) Get() *V1OrganismQueryRequest {
	return v.value
}

func (v *NullableV1OrganismQueryRequest) Set(val *V1OrganismQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1OrganismQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1OrganismQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1OrganismQueryRequest(val *V1OrganismQueryRequest) *NullableV1OrganismQueryRequest {
	return &NullableV1OrganismQueryRequest{value: val, isSet: true}
}

func (v NullableV1OrganismQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1OrganismQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

