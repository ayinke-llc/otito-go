/*
Otito API documentation

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
Contact: contact@ayinke.ventures
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServerCreateApplicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerCreateApplicationRequest{}

// ServerCreateApplicationRequest struct for ServerCreateApplicationRequest
type ServerCreateApplicationRequest struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name"`
}

// NewServerCreateApplicationRequest instantiates a new ServerCreateApplicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCreateApplicationRequest(name string) *ServerCreateApplicationRequest {
	this := ServerCreateApplicationRequest{}
	this.Name = name
	return &this
}

// NewServerCreateApplicationRequestWithDefaults instantiates a new ServerCreateApplicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCreateApplicationRequestWithDefaults() *ServerCreateApplicationRequest {
	this := ServerCreateApplicationRequest{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServerCreateApplicationRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCreateApplicationRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServerCreateApplicationRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ServerCreateApplicationRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *ServerCreateApplicationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServerCreateApplicationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServerCreateApplicationRequest) SetName(v string) {
	o.Name = v
}

func (o ServerCreateApplicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerCreateApplicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableServerCreateApplicationRequest struct {
	value *ServerCreateApplicationRequest
	isSet bool
}

func (v NullableServerCreateApplicationRequest) Get() *ServerCreateApplicationRequest {
	return v.value
}

func (v *NullableServerCreateApplicationRequest) Set(val *ServerCreateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCreateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCreateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCreateApplicationRequest(val *ServerCreateApplicationRequest) *NullableServerCreateApplicationRequest {
	return &NullableServerCreateApplicationRequest{value: val, isSet: true}
}

func (v NullableServerCreateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCreateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


