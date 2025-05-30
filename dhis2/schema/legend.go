/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"time"
)

// checks if the Legend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Legend{}

// Legend struct for Legend
type Legend struct {
	Access          *Access          `json:"access,omitempty"`
	AttributeValues []AttributeValue `json:"attributeValues,omitempty"`
	Code            *string          `json:"code,omitempty"`
	Color           *string          `json:"color,omitempty"`
	Created         *time.Time       `json:"created,omitempty"`
	CreatedBy       *UserDto         `json:"createdBy,omitempty"`
	DisplayName     *string          `json:"displayName,omitempty"`
	EndValue        *float64         `json:"endValue,omitempty"`
	Favorite        *bool            `json:"favorite,omitempty"`
	Favorites       []string         `json:"favorites,omitempty"`
	Href            *string          `json:"href,omitempty"`
	Id              *string          `json:"id,omitempty"`
	Image           *string          `json:"image,omitempty"`
	LastUpdated     *time.Time       `json:"lastUpdated,omitempty"`
	LastUpdatedBy   *UserDto         `json:"lastUpdatedBy,omitempty"`
	Name            *string          `json:"name,omitempty"`
	Sharing         *Sharing         `json:"sharing,omitempty"`
	StartValue      *float64         `json:"startValue,omitempty"`
	Translations    []Translation    `json:"translations,omitempty"`
}

// NewLegend instantiates a new Legend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegend() *Legend {
	this := Legend{}
	return &this
}

// NewLegendWithDefaults instantiates a new Legend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegendWithDefaults() *Legend {
	this := Legend{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Legend) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *Legend) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *Legend) SetAccess(v Access) {
	o.Access = &v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *Legend) GetAttributeValues() []AttributeValue {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValue
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetAttributeValuesOk() ([]AttributeValue, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *Legend) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValue and assigns it to the AttributeValues field.
func (o *Legend) SetAttributeValues(v []AttributeValue) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Legend) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Legend) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Legend) SetCode(v string) {
	o.Code = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Legend) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Legend) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Legend) SetColor(v string) {
	o.Color = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Legend) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Legend) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Legend) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Legend) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Legend) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *Legend) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Legend) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Legend) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Legend) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEndValue returns the EndValue field value if set, zero value otherwise.
func (o *Legend) GetEndValue() float64 {
	if o == nil || IsNil(o.EndValue) {
		var ret float64
		return ret
	}
	return *o.EndValue
}

// GetEndValueOk returns a tuple with the EndValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetEndValueOk() (*float64, bool) {
	if o == nil || IsNil(o.EndValue) {
		return nil, false
	}
	return o.EndValue, true
}

// HasEndValue returns a boolean if a field has been set.
func (o *Legend) HasEndValue() bool {
	if o != nil && !IsNil(o.EndValue) {
		return true
	}

	return false
}

// SetEndValue gets a reference to the given float64 and assigns it to the EndValue field.
func (o *Legend) SetEndValue(v float64) {
	o.EndValue = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *Legend) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *Legend) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *Legend) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *Legend) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *Legend) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *Legend) SetFavorites(v []string) {
	o.Favorites = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Legend) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Legend) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Legend) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Legend) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Legend) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Legend) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Legend) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Legend) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Legend) SetImage(v string) {
	o.Image = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Legend) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Legend) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Legend) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *Legend) GetLastUpdatedBy() UserDto {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetLastUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *Legend) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given UserDto and assigns it to the LastUpdatedBy field.
func (o *Legend) SetLastUpdatedBy(v UserDto) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Legend) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Legend) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Legend) SetName(v string) {
	o.Name = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *Legend) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *Legend) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *Legend) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetStartValue returns the StartValue field value if set, zero value otherwise.
func (o *Legend) GetStartValue() float64 {
	if o == nil || IsNil(o.StartValue) {
		var ret float64
		return ret
	}
	return *o.StartValue
}

// GetStartValueOk returns a tuple with the StartValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetStartValueOk() (*float64, bool) {
	if o == nil || IsNil(o.StartValue) {
		return nil, false
	}
	return o.StartValue, true
}

// HasStartValue returns a boolean if a field has been set.
func (o *Legend) HasStartValue() bool {
	if o != nil && !IsNil(o.StartValue) {
		return true
	}

	return false
}

// SetStartValue gets a reference to the given float64 and assigns it to the StartValue field.
func (o *Legend) SetStartValue(v float64) {
	o.StartValue = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *Legend) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Legend) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *Legend) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *Legend) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o Legend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Legend) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
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
	if !IsNil(o.EndValue) {
		toSerialize["endValue"] = o.EndValue
	}
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
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
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
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.StartValue) {
		toSerialize["startValue"] = o.StartValue
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableLegend struct {
	value *Legend
	isSet bool
}

func (v NullableLegend) Get() *Legend {
	return v.value
}

func (v *NullableLegend) Set(val *Legend) {
	v.value = val
	v.isSet = true
}

func (v NullableLegend) IsSet() bool {
	return v.isSet
}

func (v *NullableLegend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegend(val *Legend) *NullableLegend {
	return &NullableLegend{value: val, isSet: true}
}

func (v NullableLegend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
