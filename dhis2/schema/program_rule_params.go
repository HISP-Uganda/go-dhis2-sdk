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

// checks if the ProgramRuleParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramRuleParams{}

// ProgramRuleParams struct for ProgramRuleParams
type ProgramRuleParams struct {
	AttributeValues    []AttributeValueParams                     `json:"attributeValues,omitempty"`
	Code               *string                                    `json:"code,omitempty"`
	Condition          *string                                    `json:"condition,omitempty"`
	Created            *time.Time                                 `json:"created,omitempty"`
	CreatedBy          *AggregateDataExchangeParamsCreatedBy      `json:"createdBy,omitempty"`
	Description        *string                                    `json:"description,omitempty"`
	DisplayName        *string                                    `json:"displayName,omitempty"`
	Favorite           *bool                                      `json:"favorite,omitempty"`
	Favorites          []string                                   `json:"favorites,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	LastUpdated        *time.Time                                 `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *AggregateDataExchangeParamsCreatedBy      `json:"lastUpdatedBy,omitempty"`
	Name               *string                                    `json:"name,omitempty"`
	Priority           *int32                                     `json:"priority,omitempty"`
	Program            *ProgramParams                             `json:"program,omitempty"`
	ProgramRuleActions []ProgramRuleParamsProgramRuleActionsInner `json:"programRuleActions,omitempty"`
	ProgramStage       *ProgramStageParams                        `json:"programStage,omitempty"`
	Sharing            *Sharing                                   `json:"sharing,omitempty"`
	Translations       []Translation                              `json:"translations,omitempty"`
}

// NewProgramRuleParams instantiates a new ProgramRuleParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramRuleParams() *ProgramRuleParams {
	this := ProgramRuleParams{}
	return &this
}

// NewProgramRuleParamsWithDefaults instantiates a new ProgramRuleParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramRuleParamsWithDefaults() *ProgramRuleParams {
	this := ProgramRuleParams{}
	return &this
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetAttributeValues() []AttributeValueParams {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValueParams
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetAttributeValuesOk() ([]AttributeValueParams, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValueParams and assigns it to the AttributeValues field.
func (o *ProgramRuleParams) SetAttributeValues(v []AttributeValueParams) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProgramRuleParams) SetCode(v string) {
	o.Code = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetCondition() string {
	if o == nil || IsNil(o.Condition) {
		var ret string
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetConditionOk() (*string, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given string and assigns it to the Condition field.
func (o *ProgramRuleParams) SetCondition(v string) {
	o.Condition = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ProgramRuleParams) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetCreatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetCreatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the CreatedBy field.
func (o *ProgramRuleParams) SetCreatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProgramRuleParams) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ProgramRuleParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ProgramRuleParams) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *ProgramRuleParams) SetFavorites(v []string) {
	o.Favorites = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProgramRuleParams) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ProgramRuleParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetLastUpdatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetLastUpdatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the LastUpdatedBy field.
func (o *ProgramRuleParams) SetLastUpdatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProgramRuleParams) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ProgramRuleParams) SetPriority(v int32) {
	o.Priority = &v
}

// GetProgram returns the Program field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetProgram() ProgramParams {
	if o == nil || IsNil(o.Program) {
		var ret ProgramParams
		return ret
	}
	return *o.Program
}

// GetProgramOk returns a tuple with the Program field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetProgramOk() (*ProgramParams, bool) {
	if o == nil || IsNil(o.Program) {
		return nil, false
	}
	return o.Program, true
}

// HasProgram returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasProgram() bool {
	if o != nil && !IsNil(o.Program) {
		return true
	}

	return false
}

// SetProgram gets a reference to the given ProgramParams and assigns it to the Program field.
func (o *ProgramRuleParams) SetProgram(v ProgramParams) {
	o.Program = &v
}

// GetProgramRuleActions returns the ProgramRuleActions field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetProgramRuleActions() []ProgramRuleParamsProgramRuleActionsInner {
	if o == nil || IsNil(o.ProgramRuleActions) {
		var ret []ProgramRuleParamsProgramRuleActionsInner
		return ret
	}
	return o.ProgramRuleActions
}

// GetProgramRuleActionsOk returns a tuple with the ProgramRuleActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetProgramRuleActionsOk() ([]ProgramRuleParamsProgramRuleActionsInner, bool) {
	if o == nil || IsNil(o.ProgramRuleActions) {
		return nil, false
	}
	return o.ProgramRuleActions, true
}

// HasProgramRuleActions returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasProgramRuleActions() bool {
	if o != nil && !IsNil(o.ProgramRuleActions) {
		return true
	}

	return false
}

// SetProgramRuleActions gets a reference to the given []ProgramRuleParamsProgramRuleActionsInner and assigns it to the ProgramRuleActions field.
func (o *ProgramRuleParams) SetProgramRuleActions(v []ProgramRuleParamsProgramRuleActionsInner) {
	o.ProgramRuleActions = v
}

// GetProgramStage returns the ProgramStage field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetProgramStage() ProgramStageParams {
	if o == nil || IsNil(o.ProgramStage) {
		var ret ProgramStageParams
		return ret
	}
	return *o.ProgramStage
}

// GetProgramStageOk returns a tuple with the ProgramStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetProgramStageOk() (*ProgramStageParams, bool) {
	if o == nil || IsNil(o.ProgramStage) {
		return nil, false
	}
	return o.ProgramStage, true
}

// HasProgramStage returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasProgramStage() bool {
	if o != nil && !IsNil(o.ProgramStage) {
		return true
	}

	return false
}

// SetProgramStage gets a reference to the given ProgramStageParams and assigns it to the ProgramStage field.
func (o *ProgramRuleParams) SetProgramStage(v ProgramStageParams) {
	o.ProgramStage = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *ProgramRuleParams) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *ProgramRuleParams) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleParams) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *ProgramRuleParams) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *ProgramRuleParams) SetTranslations(v []Translation) {
	o.Translations = v
}

func (o ProgramRuleParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramRuleParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeValues) {
		toSerialize["attributeValues"] = o.AttributeValues
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
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
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	if !IsNil(o.Favorites) {
		toSerialize["favorites"] = o.Favorites
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
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Program) {
		toSerialize["program"] = o.Program
	}
	if !IsNil(o.ProgramRuleActions) {
		toSerialize["programRuleActions"] = o.ProgramRuleActions
	}
	if !IsNil(o.ProgramStage) {
		toSerialize["programStage"] = o.ProgramStage
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableProgramRuleParams struct {
	value *ProgramRuleParams
	isSet bool
}

func (v NullableProgramRuleParams) Get() *ProgramRuleParams {
	return v.value
}

func (v *NullableProgramRuleParams) Set(val *ProgramRuleParams) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramRuleParams) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramRuleParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramRuleParams(val *ProgramRuleParams) *NullableProgramRuleParams {
	return &NullableProgramRuleParams{value: val, isSet: true}
}

func (v NullableProgramRuleParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramRuleParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
