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

// checks if the Option type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Option{}

// Option struct for Option
type Option struct {
	Access             *Access                 `json:"access,omitempty"`
	AttributeValues    []AttributeValue        `json:"attributeValues,omitempty"`
	Code               *string                 `json:"code,omitempty"`
	Created            *time.Time              `json:"created,omitempty"`
	CreatedBy          *UserDto                `json:"createdBy,omitempty"`
	Description        *string                 `json:"description,omitempty"`
	DisplayDescription *string                 `json:"displayDescription,omitempty"`
	DisplayFormName    *string                 `json:"displayFormName,omitempty"`
	DisplayName        *string                 `json:"displayName,omitempty"`
	DisplayShortName   *string                 `json:"displayShortName,omitempty"`
	Favorite           *bool                   `json:"favorite,omitempty"`
	Favorites          []string                `json:"favorites,omitempty"`
	FormName           *string                 `json:"formName,omitempty"`
	Href               *string                 `json:"href,omitempty"`
	Id                 *string                 `json:"id,omitempty"`
	LastUpdated        *time.Time              `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *UserDto                `json:"lastUpdatedBy,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	OptionSet          *BaseIdentifiableObject `json:"optionSet,omitempty"`
	Sharing            *Sharing                `json:"sharing,omitempty"`
	ShortName          *string                 `json:"shortName,omitempty"`
	SortOrder          *int32                  `json:"sortOrder,omitempty"`
	Style              *ObjectStyle            `json:"style,omitempty"`
	Translations       []Translation           `json:"translations,omitempty"`
}

// NewOption instantiates a new Option object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOption() *Option {
	this := Option{}
	return &this
}

// NewOptionWithDefaults instantiates a new Option object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionWithDefaults() *Option {
	this := Option{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *Option) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *Option) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *Option) SetAccess(v Access) {
	o.Access = &v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *Option) GetAttributeValues() []AttributeValue {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValue
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetAttributeValuesOk() ([]AttributeValue, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *Option) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValue and assigns it to the AttributeValues field.
func (o *Option) SetAttributeValues(v []AttributeValue) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Option) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Option) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Option) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Option) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Option) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Option) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Option) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Option) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *Option) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Option) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Option) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Option) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *Option) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *Option) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *Option) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetDisplayFormName returns the DisplayFormName field value if set, zero value otherwise.
func (o *Option) GetDisplayFormName() string {
	if o == nil || IsNil(o.DisplayFormName) {
		var ret string
		return ret
	}
	return *o.DisplayFormName
}

// GetDisplayFormNameOk returns a tuple with the DisplayFormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayFormName) {
		return nil, false
	}
	return o.DisplayFormName, true
}

// HasDisplayFormName returns a boolean if a field has been set.
func (o *Option) HasDisplayFormName() bool {
	if o != nil && !IsNil(o.DisplayFormName) {
		return true
	}

	return false
}

// SetDisplayFormName gets a reference to the given string and assigns it to the DisplayFormName field.
func (o *Option) SetDisplayFormName(v string) {
	o.DisplayFormName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Option) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Option) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Option) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayShortName returns the DisplayShortName field value if set, zero value otherwise.
func (o *Option) GetDisplayShortName() string {
	if o == nil || IsNil(o.DisplayShortName) {
		var ret string
		return ret
	}
	return *o.DisplayShortName
}

// GetDisplayShortNameOk returns a tuple with the DisplayShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayShortName) {
		return nil, false
	}
	return o.DisplayShortName, true
}

// HasDisplayShortName returns a boolean if a field has been set.
func (o *Option) HasDisplayShortName() bool {
	if o != nil && !IsNil(o.DisplayShortName) {
		return true
	}

	return false
}

// SetDisplayShortName gets a reference to the given string and assigns it to the DisplayShortName field.
func (o *Option) SetDisplayShortName(v string) {
	o.DisplayShortName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *Option) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *Option) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *Option) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *Option) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *Option) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *Option) SetFavorites(v []string) {
	o.Favorites = v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *Option) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *Option) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *Option) SetFormName(v string) {
	o.FormName = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Option) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Option) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Option) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Option) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Option) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Option) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Option) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Option) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Option) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *Option) GetLastUpdatedBy() UserDto {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetLastUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *Option) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given UserDto and assigns it to the LastUpdatedBy field.
func (o *Option) SetLastUpdatedBy(v UserDto) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Option) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Option) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Option) SetName(v string) {
	o.Name = &v
}

// GetOptionSet returns the OptionSet field value if set, zero value otherwise.
func (o *Option) GetOptionSet() BaseIdentifiableObject {
	if o == nil || IsNil(o.OptionSet) {
		var ret BaseIdentifiableObject
		return ret
	}
	return *o.OptionSet
}

// GetOptionSetOk returns a tuple with the OptionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetOptionSetOk() (*BaseIdentifiableObject, bool) {
	if o == nil || IsNil(o.OptionSet) {
		return nil, false
	}
	return o.OptionSet, true
}

// HasOptionSet returns a boolean if a field has been set.
func (o *Option) HasOptionSet() bool {
	if o != nil && !IsNil(o.OptionSet) {
		return true
	}

	return false
}

// SetOptionSet gets a reference to the given BaseIdentifiableObject and assigns it to the OptionSet field.
func (o *Option) SetOptionSet(v BaseIdentifiableObject) {
	o.OptionSet = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *Option) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *Option) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *Option) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *Option) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *Option) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *Option) SetShortName(v string) {
	o.ShortName = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *Option) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *Option) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *Option) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *Option) GetStyle() ObjectStyle {
	if o == nil || IsNil(o.Style) {
		var ret ObjectStyle
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetStyleOk() (*ObjectStyle, bool) {
	if o == nil || IsNil(o.Style) {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *Option) HasStyle() bool {
	if o != nil && !IsNil(o.Style) {
		return true
	}

	return false
}

// SetStyle gets a reference to the given ObjectStyle and assigns it to the Style field.
func (o *Option) SetStyle(v ObjectStyle) {
	o.Style = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *Option) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Option) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *Option) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *Option) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o Option) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Option) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DisplayDescription) {
		toSerialize["displayDescription"] = o.DisplayDescription
	}
	if !IsNil(o.DisplayFormName) {
		toSerialize["displayFormName"] = o.DisplayFormName
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DisplayShortName) {
		toSerialize["displayShortName"] = o.DisplayShortName
	}
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	if !IsNil(o.Favorites) {
		toSerialize["favorites"] = o.Favorites
	}
	if !IsNil(o.FormName) {
		toSerialize["formName"] = o.FormName
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
	if !IsNil(o.OptionSet) {
		toSerialize["optionSet"] = o.OptionSet
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.Style) {
		toSerialize["style"] = o.Style
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableOption struct {
	value *Option
	isSet bool
}

func (v NullableOption) Get() *Option {
	return v.value
}

func (v *NullableOption) Set(val *Option) {
	v.value = val
	v.isSet = true
}

func (v NullableOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOption(val *Option) *NullableOption {
	return &NullableOption{value: val, isSet: true}
}

func (v NullableOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
