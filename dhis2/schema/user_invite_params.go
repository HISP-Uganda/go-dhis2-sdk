/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// checks if the UserInviteParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInviteParams{}

// UserInviteParams struct for UserInviteParams
type UserInviteParams struct {
	Email              *string `json:"email,omitempty"`
	FirstName          *string `json:"firstName,omitempty"`
	GRecaptchaResponse *string `json:"g-recaptcha-response,omitempty"`
	Password           *string `json:"password,omitempty"`
	Surname            *string `json:"surname,omitempty"`
	Token              *string `json:"token,omitempty"`
	Username           *string `json:"username,omitempty"`
}

// NewUserInviteParams instantiates a new UserInviteParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInviteParams() *UserInviteParams {
	this := UserInviteParams{}
	return &this
}

// NewUserInviteParamsWithDefaults instantiates a new UserInviteParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInviteParamsWithDefaults() *UserInviteParams {
	this := UserInviteParams{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserInviteParams) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserInviteParams) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserInviteParams) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserInviteParams) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserInviteParams) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserInviteParams) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGRecaptchaResponse returns the GRecaptchaResponse field value if set, zero value otherwise.
func (o *UserInviteParams) GetGRecaptchaResponse() string {
	if o == nil || IsNil(o.GRecaptchaResponse) {
		var ret string
		return ret
	}
	return *o.GRecaptchaResponse
}

// GetGRecaptchaResponseOk returns a tuple with the GRecaptchaResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetGRecaptchaResponseOk() (*string, bool) {
	if o == nil || IsNil(o.GRecaptchaResponse) {
		return nil, false
	}
	return o.GRecaptchaResponse, true
}

// HasGRecaptchaResponse returns a boolean if a field has been set.
func (o *UserInviteParams) HasGRecaptchaResponse() bool {
	if o != nil && !IsNil(o.GRecaptchaResponse) {
		return true
	}

	return false
}

// SetGRecaptchaResponse gets a reference to the given string and assigns it to the GRecaptchaResponse field.
func (o *UserInviteParams) SetGRecaptchaResponse(v string) {
	o.GRecaptchaResponse = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserInviteParams) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserInviteParams) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserInviteParams) SetPassword(v string) {
	o.Password = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *UserInviteParams) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *UserInviteParams) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *UserInviteParams) SetSurname(v string) {
	o.Surname = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserInviteParams) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserInviteParams) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserInviteParams) SetToken(v string) {
	o.Token = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserInviteParams) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInviteParams) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserInviteParams) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserInviteParams) SetUsername(v string) {
	o.Username = &v
}

func (o UserInviteParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInviteParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.GRecaptchaResponse) {
		toSerialize["g-recaptcha-response"] = o.GRecaptchaResponse
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUserInviteParams struct {
	value *UserInviteParams
	isSet bool
}

func (v NullableUserInviteParams) Get() *UserInviteParams {
	return v.value
}

func (v *NullableUserInviteParams) Set(val *UserInviteParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInviteParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInviteParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInviteParams(val *UserInviteParams) *NullableUserInviteParams {
	return &NullableUserInviteParams{value: val, isSet: true}
}

func (v NullableUserInviteParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInviteParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
