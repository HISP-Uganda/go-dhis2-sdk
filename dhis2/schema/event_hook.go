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

// checks if the EventHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventHook{}

// EventHook struct for EventHook
type EventHook struct {
	Access          *Access                 `json:"access,omitempty"`
	AttributeValues []AttributeValue        `json:"attributeValues,omitempty"`
	Code            *string                 `json:"code,omitempty"`
	Created         *time.Time              `json:"created,omitempty"`
	CreatedBy       *UserDto                `json:"createdBy,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	Disabled        bool                    `json:"disabled"`
	DisplayName     *string                 `json:"displayName,omitempty"`
	Favorite        *bool                   `json:"favorite,omitempty"`
	Favorites       []string                `json:"favorites,omitempty"`
	Href            *string                 `json:"href,omitempty"`
	Id              *string                 `json:"id,omitempty"`
	LastUpdated     *time.Time              `json:"lastUpdated,omitempty"`
	LastUpdatedBy   *UserDto                `json:"lastUpdatedBy,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Sharing         *Sharing                `json:"sharing,omitempty"`
	Source          Source                  `json:"source"`
	Targets         []EventHookTargetsInner `json:"targets"`
	Translations    []Translation           `json:"translations,omitempty"`
}

type _EventHook EventHook

// NewEventHook instantiates a new EventHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHook(disabled bool, source Source, targets []EventHookTargetsInner) *EventHook {
	this := EventHook{}
	this.Disabled = disabled
	this.Source = source
	this.Targets = targets
	return &this
}

// NewEventHookWithDefaults instantiates a new EventHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookWithDefaults() *EventHook {
	this := EventHook{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *EventHook) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *EventHook) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *EventHook) SetAccess(v Access) {
	o.Access = &v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *EventHook) GetAttributeValues() []AttributeValue {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValue
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetAttributeValuesOk() ([]AttributeValue, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *EventHook) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValue and assigns it to the AttributeValues field.
func (o *EventHook) SetAttributeValues(v []AttributeValue) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EventHook) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EventHook) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EventHook) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EventHook) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EventHook) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EventHook) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *EventHook) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EventHook) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *EventHook) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventHook) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventHook) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventHook) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value
func (o *EventHook) GetDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *EventHook) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *EventHook) SetDisabled(v bool) {
	o.Disabled = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *EventHook) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EventHook) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *EventHook) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *EventHook) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *EventHook) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *EventHook) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *EventHook) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *EventHook) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *EventHook) SetFavorites(v []string) {
	o.Favorites = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *EventHook) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *EventHook) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *EventHook) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventHook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventHook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventHook) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EventHook) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EventHook) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EventHook) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *EventHook) GetLastUpdatedBy() UserDto {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetLastUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *EventHook) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given UserDto and assigns it to the LastUpdatedBy field.
func (o *EventHook) SetLastUpdatedBy(v UserDto) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventHook) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventHook) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventHook) SetName(v string) {
	o.Name = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *EventHook) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *EventHook) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *EventHook) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetSource returns the Source field value
func (o *EventHook) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *EventHook) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *EventHook) SetSource(v Source) {
	o.Source = v
}

// GetTargets returns the Targets field value
func (o *EventHook) GetTargets() []EventHookTargetsInner {
	if o == nil {
		var ret []EventHookTargetsInner
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *EventHook) GetTargetsOk() ([]EventHookTargetsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *EventHook) SetTargets(v []EventHookTargetsInner) {
	o.Targets = v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *EventHook) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHook) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *EventHook) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *EventHook) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o EventHook) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventHook) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["disabled"] = o.Disabled
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
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
	toSerialize["source"] = o.Source
	toSerialize["targets"] = o.Targets
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

func (o *EventHook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"disabled",
		"source",
		"targets",
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

	varEventHook := _EventHook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventHook)

	if err != nil {
		return err
	}

	*o = EventHook(varEventHook)

	return err
}

type NullableEventHook struct {
	value *EventHook
	isSet bool
}

func (v NullableEventHook) Get() *EventHook {
	return v.value
}

func (v *NullableEventHook) Set(val *EventHook) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHook) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHook(val *EventHook) *NullableEventHook {
	return &NullableEventHook{value: val, isSet: true}
}

func (v NullableEventHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
