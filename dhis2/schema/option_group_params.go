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

// checks if the OptionGroupParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionGroupParams{}

// OptionGroupParams struct for OptionGroupParams
type OptionGroupParams struct {
	AggregationType    AggregationType                       `json:"aggregationType"`
	AttributeValues    []AttributeValueParams                `json:"attributeValues,omitempty"`
	Code               *string                               `json:"code,omitempty"`
	Created            *time.Time                            `json:"created,omitempty"`
	CreatedBy          *AggregateDataExchangeParamsCreatedBy `json:"createdBy,omitempty"`
	Description        *string                               `json:"description,omitempty"`
	DimensionItem      *string                               `json:"dimensionItem,omitempty"`
	DisplayDescription *string                               `json:"displayDescription,omitempty"`
	DisplayFormName    *string                               `json:"displayFormName,omitempty"`
	DisplayName        *string                               `json:"displayName,omitempty"`
	DisplayShortName   *string                               `json:"displayShortName,omitempty"`
	Favorite           *bool                                 `json:"favorite,omitempty"`
	Favorites          []string                              `json:"favorites,omitempty"`
	FormName           *string                               `json:"formName,omitempty"`
	Id                 *string                               `json:"id,omitempty"`
	LastUpdated        *time.Time                            `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *AggregateDataExchangeParamsCreatedBy `json:"lastUpdatedBy,omitempty"`
	LegendSet          *CategoryOptionComboParamsLegendSet   `json:"legendSet,omitempty"`
	LegendSets         []CategoryOptionComboParamsLegendSet  `json:"legendSets,omitempty"`
	Name               *string                               `json:"name,omitempty"`
	OptionSet          *OptionSetParams                      `json:"optionSet,omitempty"`
	Options            []OptionGroupParamsOptionsInner       `json:"options,omitempty"`
	QueryMods          *QueryModifiers                       `json:"queryMods,omitempty"`
	Sharing            *Sharing                              `json:"sharing,omitempty"`
	ShortName          *string                               `json:"shortName,omitempty"`
	Translations       []Translation                         `json:"translations,omitempty"`
}

type _OptionGroupParams OptionGroupParams

// NewOptionGroupParams instantiates a new OptionGroupParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionGroupParams(aggregationType AggregationType) *OptionGroupParams {
	this := OptionGroupParams{}
	this.AggregationType = aggregationType
	return &this
}

// NewOptionGroupParamsWithDefaults instantiates a new OptionGroupParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionGroupParamsWithDefaults() *OptionGroupParams {
	this := OptionGroupParams{}
	return &this
}

// GetAggregationType returns the AggregationType field value
func (o *OptionGroupParams) GetAggregationType() AggregationType {
	if o == nil {
		var ret AggregationType
		return ret
	}

	return o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetAggregationTypeOk() (*AggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationType, true
}

// SetAggregationType sets field value
func (o *OptionGroupParams) SetAggregationType(v AggregationType) {
	o.AggregationType = v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *OptionGroupParams) GetAttributeValues() []AttributeValueParams {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValueParams
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetAttributeValuesOk() ([]AttributeValueParams, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *OptionGroupParams) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValueParams and assigns it to the AttributeValues field.
func (o *OptionGroupParams) SetAttributeValues(v []AttributeValueParams) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptionGroupParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptionGroupParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptionGroupParams) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OptionGroupParams) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OptionGroupParams) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OptionGroupParams) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OptionGroupParams) GetCreatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetCreatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OptionGroupParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the CreatedBy field.
func (o *OptionGroupParams) SetCreatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OptionGroupParams) SetDescription(v string) {
	o.Description = &v
}

// GetDimensionItem returns the DimensionItem field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDimensionItem() string {
	if o == nil || IsNil(o.DimensionItem) {
		var ret string
		return ret
	}
	return *o.DimensionItem
}

// GetDimensionItemOk returns a tuple with the DimensionItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDimensionItemOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionItem) {
		return nil, false
	}
	return o.DimensionItem, true
}

// HasDimensionItem returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDimensionItem() bool {
	if o != nil && !IsNil(o.DimensionItem) {
		return true
	}

	return false
}

// SetDimensionItem gets a reference to the given string and assigns it to the DimensionItem field.
func (o *OptionGroupParams) SetDimensionItem(v string) {
	o.DimensionItem = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *OptionGroupParams) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetDisplayFormName returns the DisplayFormName field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDisplayFormName() string {
	if o == nil || IsNil(o.DisplayFormName) {
		var ret string
		return ret
	}
	return *o.DisplayFormName
}

// GetDisplayFormNameOk returns a tuple with the DisplayFormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDisplayFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayFormName) {
		return nil, false
	}
	return o.DisplayFormName, true
}

// HasDisplayFormName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDisplayFormName() bool {
	if o != nil && !IsNil(o.DisplayFormName) {
		return true
	}

	return false
}

// SetDisplayFormName gets a reference to the given string and assigns it to the DisplayFormName field.
func (o *OptionGroupParams) SetDisplayFormName(v string) {
	o.DisplayFormName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OptionGroupParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayShortName returns the DisplayShortName field value if set, zero value otherwise.
func (o *OptionGroupParams) GetDisplayShortName() string {
	if o == nil || IsNil(o.DisplayShortName) {
		var ret string
		return ret
	}
	return *o.DisplayShortName
}

// GetDisplayShortNameOk returns a tuple with the DisplayShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetDisplayShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayShortName) {
		return nil, false
	}
	return o.DisplayShortName, true
}

// HasDisplayShortName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasDisplayShortName() bool {
	if o != nil && !IsNil(o.DisplayShortName) {
		return true
	}

	return false
}

// SetDisplayShortName gets a reference to the given string and assigns it to the DisplayShortName field.
func (o *OptionGroupParams) SetDisplayShortName(v string) {
	o.DisplayShortName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *OptionGroupParams) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *OptionGroupParams) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *OptionGroupParams) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *OptionGroupParams) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *OptionGroupParams) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *OptionGroupParams) SetFavorites(v []string) {
	o.Favorites = v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *OptionGroupParams) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *OptionGroupParams) SetFormName(v string) {
	o.FormName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionGroupParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionGroupParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionGroupParams) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OptionGroupParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OptionGroupParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OptionGroupParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *OptionGroupParams) GetLastUpdatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetLastUpdatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *OptionGroupParams) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the LastUpdatedBy field.
func (o *OptionGroupParams) SetLastUpdatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.LastUpdatedBy = &v
}

// GetLegendSet returns the LegendSet field value if set, zero value otherwise.
func (o *OptionGroupParams) GetLegendSet() CategoryOptionComboParamsLegendSet {
	if o == nil || IsNil(o.LegendSet) {
		var ret CategoryOptionComboParamsLegendSet
		return ret
	}
	return *o.LegendSet
}

// GetLegendSetOk returns a tuple with the LegendSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetLegendSetOk() (*CategoryOptionComboParamsLegendSet, bool) {
	if o == nil || IsNil(o.LegendSet) {
		return nil, false
	}
	return o.LegendSet, true
}

// HasLegendSet returns a boolean if a field has been set.
func (o *OptionGroupParams) HasLegendSet() bool {
	if o != nil && !IsNil(o.LegendSet) {
		return true
	}

	return false
}

// SetLegendSet gets a reference to the given CategoryOptionComboParamsLegendSet and assigns it to the LegendSet field.
func (o *OptionGroupParams) SetLegendSet(v CategoryOptionComboParamsLegendSet) {
	o.LegendSet = &v
}

// GetLegendSets returns the LegendSets field value if set, zero value otherwise.
func (o *OptionGroupParams) GetLegendSets() []CategoryOptionComboParamsLegendSet {
	if o == nil || IsNil(o.LegendSets) {
		var ret []CategoryOptionComboParamsLegendSet
		return ret
	}
	return o.LegendSets
}

// GetLegendSetsOk returns a tuple with the LegendSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetLegendSetsOk() ([]CategoryOptionComboParamsLegendSet, bool) {
	if o == nil || IsNil(o.LegendSets) {
		return nil, false
	}
	return o.LegendSets, true
}

// HasLegendSets returns a boolean if a field has been set.
func (o *OptionGroupParams) HasLegendSets() bool {
	if o != nil && !IsNil(o.LegendSets) {
		return true
	}

	return false
}

// SetLegendSets gets a reference to the given []CategoryOptionComboParamsLegendSet and assigns it to the LegendSets field.
func (o *OptionGroupParams) SetLegendSets(v []CategoryOptionComboParamsLegendSet) {
	o.LegendSets = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OptionGroupParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OptionGroupParams) SetName(v string) {
	o.Name = &v
}

// GetOptionSet returns the OptionSet field value if set, zero value otherwise.
func (o *OptionGroupParams) GetOptionSet() OptionSetParams {
	if o == nil || IsNil(o.OptionSet) {
		var ret OptionSetParams
		return ret
	}
	return *o.OptionSet
}

// GetOptionSetOk returns a tuple with the OptionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetOptionSetOk() (*OptionSetParams, bool) {
	if o == nil || IsNil(o.OptionSet) {
		return nil, false
	}
	return o.OptionSet, true
}

// HasOptionSet returns a boolean if a field has been set.
func (o *OptionGroupParams) HasOptionSet() bool {
	if o != nil && !IsNil(o.OptionSet) {
		return true
	}

	return false
}

// SetOptionSet gets a reference to the given OptionSetParams and assigns it to the OptionSet field.
func (o *OptionGroupParams) SetOptionSet(v OptionSetParams) {
	o.OptionSet = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *OptionGroupParams) GetOptions() []OptionGroupParamsOptionsInner {
	if o == nil || IsNil(o.Options) {
		var ret []OptionGroupParamsOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetOptionsOk() ([]OptionGroupParamsOptionsInner, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *OptionGroupParams) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []OptionGroupParamsOptionsInner and assigns it to the Options field.
func (o *OptionGroupParams) SetOptions(v []OptionGroupParamsOptionsInner) {
	o.Options = v
}

// GetQueryMods returns the QueryMods field value if set, zero value otherwise.
func (o *OptionGroupParams) GetQueryMods() QueryModifiers {
	if o == nil || IsNil(o.QueryMods) {
		var ret QueryModifiers
		return ret
	}
	return *o.QueryMods
}

// GetQueryModsOk returns a tuple with the QueryMods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetQueryModsOk() (*QueryModifiers, bool) {
	if o == nil || IsNil(o.QueryMods) {
		return nil, false
	}
	return o.QueryMods, true
}

// HasQueryMods returns a boolean if a field has been set.
func (o *OptionGroupParams) HasQueryMods() bool {
	if o != nil && !IsNil(o.QueryMods) {
		return true
	}

	return false
}

// SetQueryMods gets a reference to the given QueryModifiers and assigns it to the QueryMods field.
func (o *OptionGroupParams) SetQueryMods(v QueryModifiers) {
	o.QueryMods = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *OptionGroupParams) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *OptionGroupParams) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *OptionGroupParams) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *OptionGroupParams) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *OptionGroupParams) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *OptionGroupParams) SetShortName(v string) {
	o.ShortName = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *OptionGroupParams) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionGroupParams) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *OptionGroupParams) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *OptionGroupParams) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o OptionGroupParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionGroupParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aggregationType"] = o.AggregationType
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
	if !IsNil(o.OptionSet) {
		toSerialize["optionSet"] = o.OptionSet
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
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

func (o *OptionGroupParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aggregationType",
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

	varOptionGroupParams := _OptionGroupParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptionGroupParams)

	if err != nil {
		return err
	}

	*o = OptionGroupParams(varOptionGroupParams)

	return err
}

type NullableOptionGroupParams struct {
	value *OptionGroupParams
	isSet bool
}

func (v NullableOptionGroupParams) Get() *OptionGroupParams {
	return v.value
}

func (v *NullableOptionGroupParams) Set(val *OptionGroupParams) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionGroupParams) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionGroupParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionGroupParams(val *OptionGroupParams) *NullableOptionGroupParams {
	return &NullableOptionGroupParams{value: val, isSet: true}
}

func (v NullableOptionGroupParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionGroupParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
