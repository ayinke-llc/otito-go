/*
Otito Ingester documentation

A suite of APIs to store and manage request and response logs

API version: 0.0.1
Contact: lanre@ayinke.ventures
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServerAppresp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerAppresp{}

// ServerAppresp struct for ServerAppresp
type ServerAppresp struct {
	Application *ServerApprespApplication `json:"application,omitempty"`
	// Generic message that tells you the status of the operation
	Message *string `json:"message,omitempty"`
	Status *bool `json:"status,omitempty"`
}

// NewServerAppresp instantiates a new ServerAppresp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerAppresp() *ServerAppresp {
	this := ServerAppresp{}
	return &this
}

// NewServerApprespWithDefaults instantiates a new ServerAppresp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerApprespWithDefaults() *ServerAppresp {
	this := ServerAppresp{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *ServerAppresp) GetApplication() ServerApprespApplication {
	if o == nil || IsNil(o.Application) {
		var ret ServerApprespApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAppresp) GetApplicationOk() (*ServerApprespApplication, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *ServerAppresp) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ServerApprespApplication and assigns it to the Application field.
func (o *ServerAppresp) SetApplication(v ServerApprespApplication) {
	o.Application = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServerAppresp) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAppresp) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServerAppresp) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServerAppresp) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServerAppresp) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAppresp) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServerAppresp) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ServerAppresp) SetStatus(v bool) {
	o.Status = &v
}

func (o ServerAppresp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerAppresp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableServerAppresp struct {
	value *ServerAppresp
	isSet bool
}

func (v NullableServerAppresp) Get() *ServerAppresp {
	return v.value
}

func (v *NullableServerAppresp) Set(val *ServerAppresp) {
	v.value = val
	v.isSet = true
}

func (v NullableServerAppresp) IsSet() bool {
	return v.isSet
}

func (v *NullableServerAppresp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerAppresp(val *ServerAppresp) *NullableServerAppresp {
	return &NullableServerAppresp{value: val, isSet: true}
}

func (v NullableServerAppresp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerAppresp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


