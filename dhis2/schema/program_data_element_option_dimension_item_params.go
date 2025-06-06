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

// checks if the ProgramDataElementOptionDimensionItemParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramDataElementOptionDimensionItemParams{}

// ProgramDataElementOptionDimensionItemParams struct for ProgramDataElementOptionDimensionItemParams
type ProgramDataElementOptionDimensionItemParams struct {
	AttributeValues    []AttributeValueParams                `json:"attributeValues,omitempty"`
	Code               *string                               `json:"code,omitempty"`
	Created            *time.Time                            `json:"created,omitempty"`
	CreatedBy          *AggregateDataExchangeParamsCreatedBy `json:"createdBy,omitempty"`
	DataElement        *DataElementParams                    `json:"dataElement,omitempty"`
	Description        *string                               `json:"description,omitempty"`
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
	Option             *OptionParams                         `json:"option,omitempty"`
	Program            *ProgramParams                        `json:"program,omitempty"`
	QueryMods          *QueryModifiers                       `json:"queryMods,omitempty"`
	Sharing            *Sharing                              `json:"sharing,omitempty"`
	Translations       []Translation                         `json:"translations,omitempty"`
	ValueType          ValueType                             `json:"valueType"`
}

type _ProgramDataElementOptionDimensionItemParams ProgramDataElementOptionDimensionItemParams

// NewProgramDataElementOptionDimensionItemParams instantiates a new ProgramDataElementOptionDimensionItemParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramDataElementOptionDimensionItemParams(valueType ValueType) *ProgramDataElementOptionDimensionItemParams {
	this := ProgramDataElementOptionDimensionItemParams{}
	this.ValueType = valueType
	return &this
}

// NewProgramDataElementOptionDimensionItemParamsWithDefaults instantiates a new ProgramDataElementOptionDimensionItemParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramDataElementOptionDimensionItemParamsWithDefaults() *ProgramDataElementOptionDimensionItemParams {
	this := ProgramDataElementOptionDimensionItemParams{}
	return &this
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetAttributeValues() []AttributeValueParams {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValueParams
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetAttributeValuesOk() ([]AttributeValueParams, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValueParams and assigns it to the AttributeValues field.
func (o *ProgramDataElementOptionDimensionItemParams) SetAttributeValues(v []AttributeValueParams) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProgramDataElementOptionDimensionItemParams) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ProgramDataElementOptionDimensionItemParams) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetCreatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetCreatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the CreatedBy field.
func (o *ProgramDataElementOptionDimensionItemParams) SetCreatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.CreatedBy = &v
}

// GetDataElement returns the DataElement field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDataElement() DataElementParams {
	if o == nil || IsNil(o.DataElement) {
		var ret DataElementParams
		return ret
	}
	return *o.DataElement
}

// GetDataElementOk returns a tuple with the DataElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDataElementOk() (*DataElementParams, bool) {
	if o == nil || IsNil(o.DataElement) {
		return nil, false
	}
	return o.DataElement, true
}

// HasDataElement returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDataElement() bool {
	if o != nil && !IsNil(o.DataElement) {
		return true
	}

	return false
}

// SetDataElement gets a reference to the given DataElementParams and assigns it to the DataElement field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDataElement(v DataElementParams) {
	o.DataElement = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetDisplayFormName returns the DisplayFormName field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayFormName() string {
	if o == nil || IsNil(o.DisplayFormName) {
		var ret string
		return ret
	}
	return *o.DisplayFormName
}

// GetDisplayFormNameOk returns a tuple with the DisplayFormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayFormName) {
		return nil, false
	}
	return o.DisplayFormName, true
}

// HasDisplayFormName returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDisplayFormName() bool {
	if o != nil && !IsNil(o.DisplayFormName) {
		return true
	}

	return false
}

// SetDisplayFormName gets a reference to the given string and assigns it to the DisplayFormName field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDisplayFormName(v string) {
	o.DisplayFormName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayShortName returns the DisplayShortName field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayShortName() string {
	if o == nil || IsNil(o.DisplayShortName) {
		var ret string
		return ret
	}
	return *o.DisplayShortName
}

// GetDisplayShortNameOk returns a tuple with the DisplayShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetDisplayShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayShortName) {
		return nil, false
	}
	return o.DisplayShortName, true
}

// HasDisplayShortName returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasDisplayShortName() bool {
	if o != nil && !IsNil(o.DisplayShortName) {
		return true
	}

	return false
}

// SetDisplayShortName gets a reference to the given string and assigns it to the DisplayShortName field.
func (o *ProgramDataElementOptionDimensionItemParams) SetDisplayShortName(v string) {
	o.DisplayShortName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ProgramDataElementOptionDimensionItemParams) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *ProgramDataElementOptionDimensionItemParams) SetFavorites(v []string) {
	o.Favorites = v
}

// GetFormName returns the FormName field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetFormName() string {
	if o == nil || IsNil(o.FormName) {
		var ret string
		return ret
	}
	return *o.FormName
}

// GetFormNameOk returns a tuple with the FormName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetFormNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormName) {
		return nil, false
	}
	return o.FormName, true
}

// HasFormName returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasFormName() bool {
	if o != nil && !IsNil(o.FormName) {
		return true
	}

	return false
}

// SetFormName gets a reference to the given string and assigns it to the FormName field.
func (o *ProgramDataElementOptionDimensionItemParams) SetFormName(v string) {
	o.FormName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProgramDataElementOptionDimensionItemParams) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ProgramDataElementOptionDimensionItemParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetLastUpdatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetLastUpdatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the LastUpdatedBy field.
func (o *ProgramDataElementOptionDimensionItemParams) SetLastUpdatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.LastUpdatedBy = &v
}

// GetLegendSet returns the LegendSet field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetLegendSet() CategoryOptionComboParamsLegendSet {
	if o == nil || IsNil(o.LegendSet) {
		var ret CategoryOptionComboParamsLegendSet
		return ret
	}
	return *o.LegendSet
}

// GetLegendSetOk returns a tuple with the LegendSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetLegendSetOk() (*CategoryOptionComboParamsLegendSet, bool) {
	if o == nil || IsNil(o.LegendSet) {
		return nil, false
	}
	return o.LegendSet, true
}

// HasLegendSet returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasLegendSet() bool {
	if o != nil && !IsNil(o.LegendSet) {
		return true
	}

	return false
}

// SetLegendSet gets a reference to the given CategoryOptionComboParamsLegendSet and assigns it to the LegendSet field.
func (o *ProgramDataElementOptionDimensionItemParams) SetLegendSet(v CategoryOptionComboParamsLegendSet) {
	o.LegendSet = &v
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetOption() OptionParams {
	if o == nil || IsNil(o.Option) {
		var ret OptionParams
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetOptionOk() (*OptionParams, bool) {
	if o == nil || IsNil(o.Option) {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasOption() bool {
	if o != nil && !IsNil(o.Option) {
		return true
	}

	return false
}

// SetOption gets a reference to the given OptionParams and assigns it to the Option field.
func (o *ProgramDataElementOptionDimensionItemParams) SetOption(v OptionParams) {
	o.Option = &v
}

// GetProgram returns the Program field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetProgram() ProgramParams {
	if o == nil || IsNil(o.Program) {
		var ret ProgramParams
		return ret
	}
	return *o.Program
}

// GetProgramOk returns a tuple with the Program field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetProgramOk() (*ProgramParams, bool) {
	if o == nil || IsNil(o.Program) {
		return nil, false
	}
	return o.Program, true
}

// HasProgram returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasProgram() bool {
	if o != nil && !IsNil(o.Program) {
		return true
	}

	return false
}

// SetProgram gets a reference to the given ProgramParams and assigns it to the Program field.
func (o *ProgramDataElementOptionDimensionItemParams) SetProgram(v ProgramParams) {
	o.Program = &v
}

// GetQueryMods returns the QueryMods field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetQueryMods() QueryModifiers {
	if o == nil || IsNil(o.QueryMods) {
		var ret QueryModifiers
		return ret
	}
	return *o.QueryMods
}

// GetQueryModsOk returns a tuple with the QueryMods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetQueryModsOk() (*QueryModifiers, bool) {
	if o == nil || IsNil(o.QueryMods) {
		return nil, false
	}
	return o.QueryMods, true
}

// HasQueryMods returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasQueryMods() bool {
	if o != nil && !IsNil(o.QueryMods) {
		return true
	}

	return false
}

// SetQueryMods gets a reference to the given QueryModifiers and assigns it to the QueryMods field.
func (o *ProgramDataElementOptionDimensionItemParams) SetQueryMods(v QueryModifiers) {
	o.QueryMods = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *ProgramDataElementOptionDimensionItemParams) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *ProgramDataElementOptionDimensionItemParams) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *ProgramDataElementOptionDimensionItemParams) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *ProgramDataElementOptionDimensionItemParams) SetTranslations(v []Translation) {
	o.Translations = v
}

// GetValueType returns the ValueType field value
func (o *ProgramDataElementOptionDimensionItemParams) GetValueType() ValueType {
	if o == nil {
		var ret ValueType
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *ProgramDataElementOptionDimensionItemParams) GetValueTypeOk() (*ValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *ProgramDataElementOptionDimensionItemParams) SetValueType(v ValueType) {
	o.ValueType = v
}

func (o ProgramDataElementOptionDimensionItemParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramDataElementOptionDimensionItemParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.DataElement) {
		toSerialize["dataElement"] = o.DataElement
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
	if !IsNil(o.Option) {
		toSerialize["option"] = o.Option
	}
	if !IsNil(o.Program) {
		toSerialize["program"] = o.Program
	}
	if !IsNil(o.QueryMods) {
		toSerialize["queryMods"] = o.QueryMods
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	toSerialize["valueType"] = o.ValueType
	return toSerialize, nil
}

func (o *ProgramDataElementOptionDimensionItemParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valueType",
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

	varProgramDataElementOptionDimensionItemParams := _ProgramDataElementOptionDimensionItemParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgramDataElementOptionDimensionItemParams)

	if err != nil {
		return err
	}

	*o = ProgramDataElementOptionDimensionItemParams(varProgramDataElementOptionDimensionItemParams)

	return err
}

type NullableProgramDataElementOptionDimensionItemParams struct {
	value *ProgramDataElementOptionDimensionItemParams
	isSet bool
}

func (v NullableProgramDataElementOptionDimensionItemParams) Get() *ProgramDataElementOptionDimensionItemParams {
	return v.value
}

func (v *NullableProgramDataElementOptionDimensionItemParams) Set(val *ProgramDataElementOptionDimensionItemParams) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramDataElementOptionDimensionItemParams) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramDataElementOptionDimensionItemParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramDataElementOptionDimensionItemParams(val *ProgramDataElementOptionDimensionItemParams) *NullableProgramDataElementOptionDimensionItemParams {
	return &NullableProgramDataElementOptionDimensionItemParams{value: val, isSet: true}
}

func (v NullableProgramDataElementOptionDimensionItemParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramDataElementOptionDimensionItemParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
