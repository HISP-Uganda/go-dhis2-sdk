/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the IndicatorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndicatorType{}

// IndicatorType struct for IndicatorType
type IndicatorType struct {
	Access          *Access          `json:"access,omitempty"`
	AttributeValues []AttributeValue `json:"attributeValues,omitempty"`
	Code            *string          `json:"code,omitempty"`
	Created         *time.Time       `json:"created,omitempty"`
	CreatedBy       *UserDto         `json:"createdBy,omitempty"`
	DisplayName     *string          `json:"displayName,omitempty"`
	Factor          int32            `json:"factor"`
	Favorite        *bool            `json:"favorite,omitempty"`
	Favorites       []string         `json:"favorites,omitempty"`
	Href            *string          `json:"href,omitempty"`
	Id              *string          `json:"id,omitempty"`
	LastUpdated     *time.Time       `json:"lastUpdated,omitempty"`
	LastUpdatedBy   *UserDto         `json:"lastUpdatedBy,omitempty"`
	Name            *string          `json:"name,omitempty"`
	Number          *bool            `json:"number,omitempty"`
	Sharing         *Sharing         `json:"sharing,omitempty"`
	Translations    []Translation    `json:"translations,omitempty"`
}

type _IndicatorType IndicatorType

// NewIndicatorType instantiates a new IndicatorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndicatorType(factor int32) *IndicatorType {
	this := IndicatorType{}
	this.Factor = factor
	return &this
}

// NewIndicatorTypeWithDefaults instantiates a new IndicatorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndicatorTypeWithDefaults() *IndicatorType {
	this := IndicatorType{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *IndicatorType) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *IndicatorType) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *IndicatorType) SetAccess(v Access) {
	o.Access = &v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *IndicatorType) GetAttributeValues() []AttributeValue {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValue
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetAttributeValuesOk() ([]AttributeValue, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *IndicatorType) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValue and assigns it to the AttributeValues field.
func (o *IndicatorType) SetAttributeValues(v []AttributeValue) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IndicatorType) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IndicatorType) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *IndicatorType) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IndicatorType) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IndicatorType) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IndicatorType) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IndicatorType) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IndicatorType) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *IndicatorType) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IndicatorType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IndicatorType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IndicatorType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFactor returns the Factor field value
func (o *IndicatorType) GetFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Factor
}

// GetFactorOk returns a tuple with the Factor field value
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Factor, true
}

// SetFactor sets field value
func (o *IndicatorType) SetFactor(v int32) {
	o.Factor = v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *IndicatorType) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *IndicatorType) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *IndicatorType) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *IndicatorType) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *IndicatorType) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *IndicatorType) SetFavorites(v []string) {
	o.Favorites = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IndicatorType) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IndicatorType) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IndicatorType) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IndicatorType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IndicatorType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IndicatorType) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IndicatorType) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IndicatorType) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IndicatorType) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *IndicatorType) GetLastUpdatedBy() UserDto {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetLastUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *IndicatorType) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given UserDto and assigns it to the LastUpdatedBy field.
func (o *IndicatorType) SetLastUpdatedBy(v UserDto) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IndicatorType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndicatorType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IndicatorType) SetName(v string) {
	o.Name = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *IndicatorType) GetNumber() bool {
	if o == nil || IsNil(o.Number) {
		var ret bool
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetNumberOk() (*bool, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *IndicatorType) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given bool and assigns it to the Number field.
func (o *IndicatorType) SetNumber(v bool) {
	o.Number = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *IndicatorType) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *IndicatorType) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *IndicatorType) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *IndicatorType) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndicatorType) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *IndicatorType) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *IndicatorType) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o IndicatorType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndicatorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.AttributeValues) {
		toSerialize["attributeValues"] = o.AttributeValues
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["factor"] = o.Factor
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	if !IsNil(o.Favorites) {
		toSerialize["favorites"] = o.Favorites
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

func (o *IndicatorType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"factor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIndicatorType := _IndicatorType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIndicatorType)

	if err != nil {
		return err
	}

	*o = IndicatorType(varIndicatorType)

	return err
}

type NullableIndicatorType struct {
	value *IndicatorType
	isSet bool
}

func (v NullableIndicatorType) Get() *IndicatorType {
	return v.value
}

func (v *NullableIndicatorType) Set(val *IndicatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableIndicatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableIndicatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndicatorType(val *IndicatorType) *NullableIndicatorType {
	return &NullableIndicatorType{value: val, isSet: true}
}

func (v NullableIndicatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndicatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
