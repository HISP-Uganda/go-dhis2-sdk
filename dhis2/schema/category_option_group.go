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

// checks if the CategoryOptionGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CategoryOptionGroup{}

// CategoryOptionGroup struct for CategoryOptionGroup
type CategoryOptionGroup struct {
	Access             *Access                  `json:"access,omitempty"`
	AggregationType    AggregationType          `json:"aggregationType"`
	AttributeValues    []AttributeValue         `json:"attributeValues,omitempty"`
	CategoryOptions    []BaseIdentifiableObject `json:"categoryOptions,omitempty"`
	Code               *string                  `json:"code,omitempty"`
	Created            *time.Time               `json:"created,omitempty"`
	CreatedBy          *UserDto                 `json:"createdBy,omitempty"`
	DataDimensionType  DataDimensionType        `json:"dataDimensionType"`
	Description        *string                  `json:"description,omitempty"`
	DimensionItem      *string                  `json:"dimensionItem,omitempty"`
	DisplayDescription *string                  `json:"displayDescription,omitempty"`
	DisplayFormName    *string                  `json:"displayFormName,omitempty"`
	DisplayName        *string                  `json:"displayName,omitempty"`
	DisplayShortName   *string                  `json:"displayShortName,omitempty"`
	Favorite           *bool                    `json:"favorite,omitempty"`
	Favorites          []string                 `json:"favorites,omitempty"`
	FormName           *string                  `json:"formName,omitempty"`
	GroupSets          []BaseIdentifiableObject `json:"groupSets,omitempty"`
	Href               *string                  `json:"href,omitempty"`
	Id                 *string                  `json:"id,omitempty"`
	LastUpdated        *time.Time               `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *UserDto                 `json:"lastUpdatedBy,omitempty"`
	LegendSet          *LegendSet               `json:"legendSet,omitempty"`
	LegendSets         []LegendSet              `json:"legendSets,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	QueryMods          *QueryModifiers          `json:"queryMods,omitempty"`
	Sharing            *Sharing                 `json:"sharing,omitempty"`
	ShortName          *string                  `json:"shortName,omitempty"`
	Translations       []Translation            `json:"translations,omitempty"`
}

type _CategoryOptionGroup CategoryOptionGroup

// NewCategoryOptionGroup instantiates a new CategoryOptionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryOptionGroup(aggregationType AggregationType, dataDimensionType DataDimensionType) *CategoryOptionGroup {
	this := CategoryOptionGroup{}
	this.AggregationType = aggregationType
	this.DataDimensionType = dataDimensionType
	return &this
}

// NewCategoryOptionGroupWithDefaults instantiates a new CategoryOptionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryOptionGroupWithDefaults() *CategoryOptionGroup {
	this := CategoryOptionGroup{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetAccess() Access {
	if o == nil || IsNil(o.Access) {
		var ret Access
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetAccessOk() (*Access, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given Access and assigns it to the Access field.
func (o *CategoryOptionGroup) SetAccess(v Access) {
	o.Access = &v
}

// GetAggregationType returns the AggregationType field value
func (o *CategoryOptionGroup) GetAggregationType() AggregationType {
	if o == nil {
		var ret AggregationType
		return ret
	}

	return o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetAggregationTypeOk() (*AggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationType, true
}

// SetAggregationType sets field value
func (o *CategoryOptionGroup) SetAggregationType(v AggregationType) {
	o.AggregationType = v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetAttributeValues() []AttributeValue {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValue
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetAttributeValuesOk() ([]AttributeValue, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValue and assigns it to the AttributeValues field.
func (o *CategoryOptionGroup) SetAttributeValues(v []AttributeValue) {
	o.AttributeValues = v
}

// GetCategoryOptions returns the CategoryOptions field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetCategoryOptions() []BaseIdentifiableObject {
	if o == nil || IsNil(o.CategoryOptions) {
		var ret []BaseIdentifiableObject
		return ret
	}
	return o.CategoryOptions
}

// GetCategoryOptionsOk returns a tuple with the CategoryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetCategoryOptionsOk() ([]BaseIdentifiableObject, bool) {
	if o == nil || IsNil(o.CategoryOptions) {
		return nil, false
	}
	return o.CategoryOptions, true
}

// HasCategoryOptions returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasCategoryOptions() bool {
	if o != nil && !IsNil(o.CategoryOptions) {
		return true
	}

	return false
}

// SetCategoryOptions gets a reference to the given []BaseIdentifiableObject and assigns it to the CategoryOptions field.
func (o *CategoryOptionGroup) SetCategoryOptions(v []BaseIdentifiableObject) {
	o.CategoryOptions = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CategoryOptionGroup) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CategoryOptionGroup) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetCreatedBy() UserDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret UserDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetCreatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given UserDto and assigns it to the CreatedBy field.
func (o *CategoryOptionGroup) SetCreatedBy(v UserDto) {
	o.CreatedBy = &v
}

// GetDataDimensionType returns the DataDimensionType field value
func (o *CategoryOptionGroup) GetDataDimensionType() DataDimensionType {
	if o == nil {
		var ret DataDimensionType
		return ret
	}

	return o.DataDimensionType
}

// GetDataDimensionTypeOk returns a tuple with the DataDimensionType field value
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDataDimensionTypeOk() (*DataDimensionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDimensionType, true
}

// SetDataDimensionType sets field value
func (o *CategoryOptionGroup) SetDataDimensionType(v DataDimensionType) {
	o.DataDimensionType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CategoryOptionGroup) SetDescription(v string) {
	o.Description = &v
}

// GetDimensionItem returns the DimensionItem field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDimensionItem() string {
	if o == nil || IsNil(o.DimensionItem) {
		var ret string
		return ret
	}
	return *o.DimensionItem
}

// GetDimensionItemOk returns a tuple with the DimensionItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDimensionItemOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionItem) {
		return nil, false
	}
	return o.DimensionItem, true
}

// HasDimensionItem returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDimensionItem() bool {
	if o != nil && !IsNil(o.DimensionItem) {
		return true
	}

	return false
}

// SetDimensionItem gets a reference to the given string and assigns it to the DimensionItem field.
func (o *CategoryOptionGroup) SetDimensionItem(v string) {
	o.DimensionItem = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *CategoryOptionGroup) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetDisplayFormName returns the DisplayFormName field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDisplayFormName() string {
	if o == nil || IsNil(o.DisplayFormName) {
		var ret string
		return ret
	}
	return *o.DisplayFormName
}

// GetDisplayFormNameOk returns a tuple with the DisplayFormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDisplayFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayFormName) {
		return nil, false
	}
	return o.DisplayFormName, true
}

// HasDisplayFormName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDisplayFormName() bool {
	if o != nil && !IsNil(o.DisplayFormName) {
		return true
	}

	return false
}

// SetDisplayFormName gets a reference to the given string and assigns it to the DisplayFormName field.
func (o *CategoryOptionGroup) SetDisplayFormName(v string) {
	o.DisplayFormName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CategoryOptionGroup) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayShortName returns the DisplayShortName field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetDisplayShortName() string {
	if o == nil || IsNil(o.DisplayShortName) {
		var ret string
		return ret
	}
	return *o.DisplayShortName
}

// GetDisplayShortNameOk returns a tuple with the DisplayShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetDisplayShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayShortName) {
		return nil, false
	}
	return o.DisplayShortName, true
}

// HasDisplayShortName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasDisplayShortName() bool {
	if o != nil && !IsNil(o.DisplayShortName) {
		return true
	}

	return false
}

// SetDisplayShortName gets a reference to the given string and assigns it to the DisplayShortName field.
func (o *CategoryOptionGroup) SetDisplayShortName(v string) {
	o.DisplayShortName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *CategoryOptionGroup) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *CategoryOptionGroup) SetFavorites(v []string) {
	o.Favorites = v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *CategoryOptionGroup) SetFormName(v string) {
	o.FormName = &v
}

// GetGroupSets returns the GroupSets field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetGroupSets() []BaseIdentifiableObject {
	if o == nil || IsNil(o.GroupSets) {
		var ret []BaseIdentifiableObject
		return ret
	}
	return o.GroupSets
}

// GetGroupSetsOk returns a tuple with the GroupSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetGroupSetsOk() ([]BaseIdentifiableObject, bool) {
	if o == nil || IsNil(o.GroupSets) {
		return nil, false
	}
	return o.GroupSets, true
}

// HasGroupSets returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasGroupSets() bool {
	if o != nil && !IsNil(o.GroupSets) {
		return true
	}

	return false
}

// SetGroupSets gets a reference to the given []BaseIdentifiableObject and assigns it to the GroupSets field.
func (o *CategoryOptionGroup) SetGroupSets(v []BaseIdentifiableObject) {
	o.GroupSets = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CategoryOptionGroup) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CategoryOptionGroup) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CategoryOptionGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetLastUpdatedBy() UserDto {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret UserDto
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetLastUpdatedByOk() (*UserDto, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given UserDto and assigns it to the LastUpdatedBy field.
func (o *CategoryOptionGroup) SetLastUpdatedBy(v UserDto) {
	o.LastUpdatedBy = &v
}

// GetLegendSet returns the LegendSet field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetLegendSet() LegendSet {
	if o == nil || IsNil(o.LegendSet) {
		var ret LegendSet
		return ret
	}
	return *o.LegendSet
}

// GetLegendSetOk returns a tuple with the LegendSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetLegendSetOk() (*LegendSet, bool) {
	if o == nil || IsNil(o.LegendSet) {
		return nil, false
	}
	return o.LegendSet, true
}

// HasLegendSet returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasLegendSet() bool {
	if o != nil && !IsNil(o.LegendSet) {
		return true
	}

	return false
}

// SetLegendSet gets a reference to the given LegendSet and assigns it to the LegendSet field.
func (o *CategoryOptionGroup) SetLegendSet(v LegendSet) {
	o.LegendSet = &v
}

// GetLegendSets returns the LegendSets field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetLegendSets() []LegendSet {
	if o == nil || IsNil(o.LegendSets) {
		var ret []LegendSet
		return ret
	}
	return o.LegendSets
}

// GetLegendSetsOk returns a tuple with the LegendSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetLegendSetsOk() ([]LegendSet, bool) {
	if o == nil || IsNil(o.LegendSets) {
		return nil, false
	}
	return o.LegendSets, true
}

// HasLegendSets returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasLegendSets() bool {
	if o != nil && !IsNil(o.LegendSets) {
		return true
	}

	return false
}

// SetLegendSets gets a reference to the given []LegendSet and assigns it to the LegendSets field.
func (o *CategoryOptionGroup) SetLegendSets(v []LegendSet) {
	o.LegendSets = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CategoryOptionGroup) SetName(v string) {
	o.Name = &v
}

// GetQueryMods returns the QueryMods field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetQueryMods() QueryModifiers {
	if o == nil || IsNil(o.QueryMods) {
		var ret QueryModifiers
		return ret
	}
	return *o.QueryMods
}

// GetQueryModsOk returns a tuple with the QueryMods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetQueryModsOk() (*QueryModifiers, bool) {
	if o == nil || IsNil(o.QueryMods) {
		return nil, false
	}
	return o.QueryMods, true
}

// HasQueryMods returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasQueryMods() bool {
	if o != nil && !IsNil(o.QueryMods) {
		return true
	}

	return false
}

// SetQueryMods gets a reference to the given QueryModifiers and assigns it to the QueryMods field.
func (o *CategoryOptionGroup) SetQueryMods(v QueryModifiers) {
	o.QueryMods = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *CategoryOptionGroup) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *CategoryOptionGroup) SetShortName(v string) {
	o.ShortName = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *CategoryOptionGroup) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryOptionGroup) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *CategoryOptionGroup) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *CategoryOptionGroup) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o CategoryOptionGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CategoryOptionGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	toSerialize["aggregationType"] = o.AggregationType
	if !IsNil(o.AttributeValues) {
		toSerialize["attributeValues"] = o.AttributeValues
	}
	if !IsNil(o.CategoryOptions) {
		toSerialize["categoryOptions"] = o.CategoryOptions
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
	toSerialize["dataDimensionType"] = o.DataDimensionType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DimensionItem) {
		toSerialize["dimensionItem"] = o.DimensionItem
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
	if !IsNil(o.GroupSets) {
		toSerialize["groupSets"] = o.GroupSets
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
	if !IsNil(o.LegendSet) {
		toSerialize["legendSet"] = o.LegendSet
	}
	if !IsNil(o.LegendSets) {
		toSerialize["legendSets"] = o.LegendSets
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.QueryMods) {
		toSerialize["queryMods"] = o.QueryMods
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

func (o *CategoryOptionGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aggregationType",
		"dataDimensionType",
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

	varCategoryOptionGroup := _CategoryOptionGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCategoryOptionGroup)

	if err != nil {
		return err
	}

	*o = CategoryOptionGroup(varCategoryOptionGroup)

	return err
}

type NullableCategoryOptionGroup struct {
	value *CategoryOptionGroup
	isSet bool
}

func (v NullableCategoryOptionGroup) Get() *CategoryOptionGroup {
	return v.value
}

func (v *NullableCategoryOptionGroup) Set(val *CategoryOptionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryOptionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryOptionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryOptionGroup(val *CategoryOptionGroup) *NullableCategoryOptionGroup {
	return &NullableCategoryOptionGroup{value: val, isSet: true}
}

func (v NullableCategoryOptionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryOptionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
