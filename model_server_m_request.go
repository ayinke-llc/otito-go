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

// checks if the ServerMRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerMRequest{}

// ServerMRequest struct for ServerMRequest
type ServerMRequest struct {
	// The app reference of the user/organisation
	App *string `json:"app,omitempty"`
	// if empty, the current timestamp when the message was received would be used
	CreatedAt *int32 `json:"created_at,omitempty"`
	// IPAddress is used to identify the iP of the user. this is optional and can be left
	IpAddress *string `json:"ip_address,omitempty"`
	Request *ServerMessageHTTPDefinition `json:"request,omitempty"`
	Response *ServerMessageHTTPDefinition `json:"response,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServerMRequest instantiates a new ServerMRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerMRequest() *ServerMRequest {
	this := ServerMRequest{}
	return &this
}

// NewServerMRequestWithDefaults instantiates a new ServerMRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerMRequestWithDefaults() *ServerMRequest {
	this := ServerMRequest{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ServerMRequest) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ServerMRequest) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *ServerMRequest) SetApp(v string) {
	o.App = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServerMRequest) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServerMRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *ServerMRequest) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ServerMRequest) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ServerMRequest) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ServerMRequest) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ServerMRequest) GetRequest() ServerMessageHTTPDefinition {
	if o == nil || IsNil(o.Request) {
		var ret ServerMessageHTTPDefinition
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetRequestOk() (*ServerMessageHTTPDefinition, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ServerMRequest) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given ServerMessageHTTPDefinition and assigns it to the Request field.
func (o *ServerMRequest) SetRequest(v ServerMessageHTTPDefinition) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ServerMRequest) GetResponse() ServerMessageHTTPDefinition {
	if o == nil || IsNil(o.Response) {
		var ret ServerMessageHTTPDefinition
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetResponseOk() (*ServerMessageHTTPDefinition, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ServerMRequest) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ServerMessageHTTPDefinition and assigns it to the Response field.
func (o *ServerMRequest) SetResponse(v ServerMessageHTTPDefinition) {
	o.Response = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServerMRequest) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerMRequest) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServerMRequest) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServerMRequest) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServerMRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerMRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableServerMRequest struct {
	value *ServerMRequest
	isSet bool
}

func (v NullableServerMRequest) Get() *ServerMRequest {
	return v.value
}

func (v *NullableServerMRequest) Set(val *ServerMRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerMRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerMRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerMRequest(val *ServerMRequest) *NullableServerMRequest {
	return &NullableServerMRequest{value: val, isSet: true}
}

func (v NullableServerMRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerMRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

