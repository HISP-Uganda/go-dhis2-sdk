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

// checks if the DashboardParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardParams{}

// DashboardParams struct for DashboardParams
type DashboardParams struct {
	AllowedFilters     []string                              `json:"allowedFilters,omitempty"`
	AttributeValues    []AttributeValueParams                `json:"attributeValues,omitempty"`
	Code               *string                               `json:"code,omitempty"`
	Created            *time.Time                            `json:"created,omitempty"`
	CreatedBy          *AggregateDataExchangeParamsCreatedBy `json:"createdBy,omitempty"`
	DashboardItems     []DashboardItemParams                 `json:"dashboardItems,omitempty"`
	Description        *string                               `json:"description,omitempty"`
	DisplayDescription *string                               `json:"displayDescription,omitempty"`
	DisplayFormName    *string                               `json:"displayFormName,omitempty"`
	DisplayName        *string                               `json:"displayName,omitempty"`
	DisplayShortName   *string                               `json:"displayShortName,omitempty"`
	Embedded           *EmbeddedDashboard                    `json:"embedded,omitempty"`
	Favorite           *bool                                 `json:"favorite,omitempty"`
	Favorites          []string                              `json:"favorites,omitempty"`
	FormName           *string                               `json:"formName,omitempty"`
	Id                 *string                               `json:"id,omitempty"`
	ItemConfig         *ItemConfig                           `json:"itemConfig,omitempty"`
	ItemCount          int32                                 `json:"itemCount"`
	LastUpdated        *time.Time                            `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *AggregateDataExchangeParamsCreatedBy `json:"lastUpdatedBy,omitempty"`
	Layout             *Layout                               `json:"layout,omitempty"`
	Name               *string                               `json:"name,omitempty"`
	RestrictFilters    *bool                                 `json:"restrictFilters,omitempty"`
	Sharing            *Sharing                              `json:"sharing,omitempty"`
	ShortName          *string                               `json:"shortName,omitempty"`
	Translations       []Translation                         `json:"translations,omitempty"`
}

type _DashboardParams DashboardParams

// NewDashboardParams instantiates a new DashboardParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardParams(itemCount int32) *DashboardParams {
	this := DashboardParams{}
	this.ItemCount = itemCount
	return &this
}

// NewDashboardParamsWithDefaults instantiates a new DashboardParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardParamsWithDefaults() *DashboardParams {
	this := DashboardParams{}
	return &this
}

// GetAllowedFilters returns the AllowedFilters field value if set, zero value otherwise.
func (o *DashboardParams) GetAllowedFilters() []string {
	if o == nil || IsNil(o.AllowedFilters) {
		var ret []string
		return ret
	}
	return o.AllowedFilters
}

// GetAllowedFiltersOk returns a tuple with the AllowedFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetAllowedFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedFilters) {
		return nil, false
	}
	return o.AllowedFilters, true
}

// HasAllowedFilters returns a boolean if a field has been set.
func (o *DashboardParams) HasAllowedFilters() bool {
	if o != nil && !IsNil(o.AllowedFilters) {
		return true
	}

	return false
}

// SetAllowedFilters gets a reference to the given []string and assigns it to the AllowedFilters field.
func (o *DashboardParams) SetAllowedFilters(v []string) {
	o.AllowedFilters = v
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *DashboardParams) GetAttributeValues() []AttributeValueParams {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValueParams
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetAttributeValuesOk() ([]AttributeValueParams, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *DashboardParams) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValueParams and assigns it to the AttributeValues field.
func (o *DashboardParams) SetAttributeValues(v []AttributeValueParams) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DashboardParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DashboardParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DashboardParams) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DashboardParams) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DashboardParams) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DashboardParams) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DashboardParams) GetCreatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetCreatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DashboardParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the CreatedBy field.
func (o *DashboardParams) SetCreatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.CreatedBy = &v
}

// GetDashboardItems returns the DashboardItems field value if set, zero value otherwise.
func (o *DashboardParams) GetDashboardItems() []DashboardItemParams {
	if o == nil || IsNil(o.DashboardItems) {
		var ret []DashboardItemParams
		return ret
	}
	return o.DashboardItems
}

// GetDashboardItemsOk returns a tuple with the DashboardItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDashboardItemsOk() ([]DashboardItemParams, bool) {
	if o == nil || IsNil(o.DashboardItems) {
		return nil, false
	}
	return o.DashboardItems, true
}

// HasDashboardItems returns a boolean if a field has been set.
func (o *DashboardParams) HasDashboardItems() bool {
	if o != nil && !IsNil(o.DashboardItems) {
		return true
	}

	return false
}

// SetDashboardItems gets a reference to the given []DashboardItemParams and assigns it to the DashboardItems field.
func (o *DashboardParams) SetDashboardItems(v []DashboardItemParams) {
	o.DashboardItems = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DashboardParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DashboardParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DashboardParams) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *DashboardParams) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *DashboardParams) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *DashboardParams) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetDisplayFormName returns the DisplayFormName field value if set, zero value otherwise.
func (o *DashboardParams) GetDisplayFormName() string {
	if o == nil || IsNil(o.DisplayFormName) {
		var ret string
		return ret
	}
	return *o.DisplayFormName
}

// GetDisplayFormNameOk returns a tuple with the DisplayFormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDisplayFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayFormName) {
		return nil, false
	}
	return o.DisplayFormName, true
}

// HasDisplayFormName returns a boolean if a field has been set.
func (o *DashboardParams) HasDisplayFormName() bool {
	if o != nil && !IsNil(o.DisplayFormName) {
		return true
	}

	return false
}

// SetDisplayFormName gets a reference to the given string and assigns it to the DisplayFormName field.
func (o *DashboardParams) SetDisplayFormName(v string) {
	o.DisplayFormName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DashboardParams) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DashboardParams) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DashboardParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayShortName returns the DisplayShortName field value if set, zero value otherwise.
func (o *DashboardParams) GetDisplayShortName() string {
	if o == nil || IsNil(o.DisplayShortName) {
		var ret string
		return ret
	}
	return *o.DisplayShortName
}

// GetDisplayShortNameOk returns a tuple with the DisplayShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetDisplayShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayShortName) {
		return nil, false
	}
	return o.DisplayShortName, true
}

// HasDisplayShortName returns a boolean if a field has been set.
func (o *DashboardParams) HasDisplayShortName() bool {
	if o != nil && !IsNil(o.DisplayShortName) {
		return true
	}

	return false
}

// SetDisplayShortName gets a reference to the given string and assigns it to the DisplayShortName field.
func (o *DashboardParams) SetDisplayShortName(v string) {
	o.DisplayShortName = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *DashboardParams) GetEmbedded() EmbeddedDashboard {
	if o == nil || IsNil(o.Embedded) {
		var ret EmbeddedDashboard
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetEmbeddedOk() (*EmbeddedDashboard, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *DashboardParams) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedDashboard and assigns it to the Embedded field.
func (o *DashboardParams) SetEmbedded(v EmbeddedDashboard) {
	o.Embedded = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *DashboardParams) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *DashboardParams) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *DashboardParams) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *DashboardParams) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *DashboardParams) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *DashboardParams) SetFavorites(v []string) {
	o.Favorites = v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *DashboardParams) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *DashboardParams) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *DashboardParams) SetFormName(v string) {
	o.FormName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DashboardParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DashboardParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DashboardParams) SetId(v string) {
	o.Id = &v
}

// GetItemConfig returns the ItemConfig field value if set, zero value otherwise.
func (o *DashboardParams) GetItemConfig() ItemConfig {
	if o == nil || IsNil(o.ItemConfig) {
		var ret ItemConfig
		return ret
	}
	return *o.ItemConfig
}

// GetItemConfigOk returns a tuple with the ItemConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetItemConfigOk() (*ItemConfig, bool) {
	if o == nil || IsNil(o.ItemConfig) {
		return nil, false
	}
	return o.ItemConfig, true
}

// HasItemConfig returns a boolean if a field has been set.
func (o *DashboardParams) HasItemConfig() bool {
	if o != nil && !IsNil(o.ItemConfig) {
		return true
	}

	return false
}

// SetItemConfig gets a reference to the given ItemConfig and assigns it to the ItemConfig field.
func (o *DashboardParams) SetItemConfig(v ItemConfig) {
	o.ItemConfig = &v
}

// GetItemCount returns the ItemCount field value
func (o *DashboardParams) GetItemCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemCount, true
}

// SetItemCount sets field value
func (o *DashboardParams) SetItemCount(v int32) {
	o.ItemCount = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DashboardParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DashboardParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DashboardParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *DashboardParams) GetLastUpdatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetLastUpdatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *DashboardParams) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the LastUpdatedBy field.
func (o *DashboardParams) SetLastUpdatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.LastUpdatedBy = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *DashboardParams) GetLayout() Layout {
	if o == nil || IsNil(o.Layout) {
		var ret Layout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetLayoutOk() (*Layout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *DashboardParams) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given Layout and assigns it to the Layout field.
func (o *DashboardParams) SetLayout(v Layout) {
	o.Layout = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DashboardParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DashboardParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DashboardParams) SetName(v string) {
	o.Name = &v
}

// GetRestrictFilters returns the RestrictFilters field value if set, zero value otherwise.
func (o *DashboardParams) GetRestrictFilters() bool {
	if o == nil || IsNil(o.RestrictFilters) {
		var ret bool
		return ret
	}
	return *o.RestrictFilters
}

// GetRestrictFiltersOk returns a tuple with the RestrictFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetRestrictFiltersOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictFilters) {
		return nil, false
	}
	return o.RestrictFilters, true
}

// HasRestrictFilters returns a boolean if a field has been set.
func (o *DashboardParams) HasRestrictFilters() bool {
	if o != nil && !IsNil(o.RestrictFilters) {
		return true
	}

	return false
}

// SetRestrictFilters gets a reference to the given bool and assigns it to the RestrictFilters field.
func (o *DashboardParams) SetRestrictFilters(v bool) {
	o.RestrictFilters = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *DashboardParams) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *DashboardParams) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *DashboardParams) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *DashboardParams) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *DashboardParams) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *DashboardParams) SetShortName(v string) {
	o.ShortName = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *DashboardParams) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardParams) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *DashboardParams) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *DashboardParams) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o DashboardParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedFilters) {
		toSerialize["allowedFilters"] = o.AllowedFilters
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
	if !IsNil(o.DashboardItems) {
		toSerialize["dashboardItems"] = o.DashboardItems
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
	if !IsNil(o.Embedded) {
		toSerialize["embedded"] = o.Embedded
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
	if !IsNil(o.ItemConfig) {
		toSerialize["itemConfig"] = o.ItemConfig
	}
	toSerialize["itemCount"] = o.ItemCount
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RestrictFilters) {
		toSerialize["restrictFilters"] = o.RestrictFilters
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

func (o *DashboardParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"itemCount",
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

	varDashboardParams := _DashboardParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDashboardParams)

	if err != nil {
		return err
	}

	*o = DashboardParams(varDashboardParams)

	return err
}

type NullableDashboardParams struct {
	value *DashboardParams
	isSet bool
}

func (v NullableDashboardParams) Get() *DashboardParams {
	return v.value
}

func (v *NullableDashboardParams) Set(val *DashboardParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardParams(val *DashboardParams) *NullableDashboardParams {
	return &NullableDashboardParams{value: val, isSet: true}
}

func (v NullableDashboardParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
