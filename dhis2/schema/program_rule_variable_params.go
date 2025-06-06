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

// checks if the ProgramRuleVariableParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramRuleVariableParams{}

// ProgramRuleVariableParams struct for ProgramRuleVariableParams
type ProgramRuleVariableParams struct {
	AttributeValues               []AttributeValueParams                `json:"attributeValues,omitempty"`
	Code                          *string                               `json:"code,omitempty"`
	Created                       *time.Time                            `json:"created,omitempty"`
	CreatedBy                     *AggregateDataExchangeParamsCreatedBy `json:"createdBy,omitempty"`
	DataElement                   *DataElementParams                    `json:"dataElement,omitempty"`
	DisplayName                   *string                               `json:"displayName,omitempty"`
	Favorite                      *bool                                 `json:"favorite,omitempty"`
	Favorites                     []string                              `json:"favorites,omitempty"`
	Id                            *string                               `json:"id,omitempty"`
	LastUpdated                   *time.Time                            `json:"lastUpdated,omitempty"`
	LastUpdatedBy                 *AggregateDataExchangeParamsCreatedBy `json:"lastUpdatedBy,omitempty"`
	Name                          *string                               `json:"name,omitempty"`
	Program                       *ProgramParams                        `json:"program,omitempty"`
	ProgramRuleVariableSourceType ProgramRuleVariableSourceType         `json:"programRuleVariableSourceType"`
	ProgramStage                  *ProgramStageParams                   `json:"programStage,omitempty"`
	Sharing                       *Sharing                              `json:"sharing,omitempty"`
	TrackedEntityAttribute        *TrackedEntityAttributeParams         `json:"trackedEntityAttribute,omitempty"`
	Translations                  []Translation                         `json:"translations,omitempty"`
	UseCodeForOptionSet           *bool                                 `json:"useCodeForOptionSet,omitempty"`
	ValueType                     ValueType                             `json:"valueType"`
}

type _ProgramRuleVariableParams ProgramRuleVariableParams

// NewProgramRuleVariableParams instantiates a new ProgramRuleVariableParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramRuleVariableParams(programRuleVariableSourceType ProgramRuleVariableSourceType, valueType ValueType) *ProgramRuleVariableParams {
	this := ProgramRuleVariableParams{}
	this.ProgramRuleVariableSourceType = programRuleVariableSourceType
	this.ValueType = valueType
	return &this
}

// NewProgramRuleVariableParamsWithDefaults instantiates a new ProgramRuleVariableParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramRuleVariableParamsWithDefaults() *ProgramRuleVariableParams {
	this := ProgramRuleVariableParams{}
	return &this
}

// GetAttributeValues returns the AttributeValues field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetAttributeValues() []AttributeValueParams {
	if o == nil || IsNil(o.AttributeValues) {
		var ret []AttributeValueParams
		return ret
	}
	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetAttributeValuesOk() ([]AttributeValueParams, bool) {
	if o == nil || IsNil(o.AttributeValues) {
		return nil, false
	}
	return o.AttributeValues, true
}

// HasAttributeValues returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasAttributeValues() bool {
	if o != nil && !IsNil(o.AttributeValues) {
		return true
	}

	return false
}

// SetAttributeValues gets a reference to the given []AttributeValueParams and assigns it to the AttributeValues field.
func (o *ProgramRuleVariableParams) SetAttributeValues(v []AttributeValueParams) {
	o.AttributeValues = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProgramRuleVariableParams) SetCode(v string) {
	o.Code = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ProgramRuleVariableParams) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetCreatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetCreatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the CreatedBy field.
func (o *ProgramRuleVariableParams) SetCreatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.CreatedBy = &v
}

// GetDataElement returns the DataElement field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetDataElement() DataElementParams {
	if o == nil || IsNil(o.DataElement) {
		var ret DataElementParams
		return ret
	}
	return *o.DataElement
}

// GetDataElementOk returns a tuple with the DataElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetDataElementOk() (*DataElementParams, bool) {
	if o == nil || IsNil(o.DataElement) {
		return nil, false
	}
	return o.DataElement, true
}

// HasDataElement returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasDataElement() bool {
	if o != nil && !IsNil(o.DataElement) {
		return true
	}

	return false
}

// SetDataElement gets a reference to the given DataElementParams and assigns it to the DataElement field.
func (o *ProgramRuleVariableParams) SetDataElement(v DataElementParams) {
	o.DataElement = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ProgramRuleVariableParams) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ProgramRuleVariableParams) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetFavorites() []string {
	if o == nil || IsNil(o.Favorites) {
		var ret []string
		return ret
	}
	return o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetFavoritesOk() ([]string, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given []string and assigns it to the Favorites field.
func (o *ProgramRuleVariableParams) SetFavorites(v []string) {
	o.Favorites = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProgramRuleVariableParams) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ProgramRuleVariableParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetLastUpdatedBy() AggregateDataExchangeParamsCreatedBy {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret AggregateDataExchangeParamsCreatedBy
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetLastUpdatedByOk() (*AggregateDataExchangeParamsCreatedBy, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given AggregateDataExchangeParamsCreatedBy and assigns it to the LastUpdatedBy field.
func (o *ProgramRuleVariableParams) SetLastUpdatedBy(v AggregateDataExchangeParamsCreatedBy) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProgramRuleVariableParams) SetName(v string) {
	o.Name = &v
}

// GetProgram returns the Program field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetProgram() ProgramParams {
	if o == nil || IsNil(o.Program) {
		var ret ProgramParams
		return ret
	}
	return *o.Program
}

// GetProgramOk returns a tuple with the Program field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetProgramOk() (*ProgramParams, bool) {
	if o == nil || IsNil(o.Program) {
		return nil, false
	}
	return o.Program, true
}

// HasProgram returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasProgram() bool {
	if o != nil && !IsNil(o.Program) {
		return true
	}

	return false
}

// SetProgram gets a reference to the given ProgramParams and assigns it to the Program field.
func (o *ProgramRuleVariableParams) SetProgram(v ProgramParams) {
	o.Program = &v
}

// GetProgramRuleVariableSourceType returns the ProgramRuleVariableSourceType field value
func (o *ProgramRuleVariableParams) GetProgramRuleVariableSourceType() ProgramRuleVariableSourceType {
	if o == nil {
		var ret ProgramRuleVariableSourceType
		return ret
	}

	return o.ProgramRuleVariableSourceType
}

// GetProgramRuleVariableSourceTypeOk returns a tuple with the ProgramRuleVariableSourceType field value
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetProgramRuleVariableSourceTypeOk() (*ProgramRuleVariableSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramRuleVariableSourceType, true
}

// SetProgramRuleVariableSourceType sets field value
func (o *ProgramRuleVariableParams) SetProgramRuleVariableSourceType(v ProgramRuleVariableSourceType) {
	o.ProgramRuleVariableSourceType = v
}

// GetProgramStage returns the ProgramStage field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetProgramStage() ProgramStageParams {
	if o == nil || IsNil(o.ProgramStage) {
		var ret ProgramStageParams
		return ret
	}
	return *o.ProgramStage
}

// GetProgramStageOk returns a tuple with the ProgramStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetProgramStageOk() (*ProgramStageParams, bool) {
	if o == nil || IsNil(o.ProgramStage) {
		return nil, false
	}
	return o.ProgramStage, true
}

// HasProgramStage returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasProgramStage() bool {
	if o != nil && !IsNil(o.ProgramStage) {
		return true
	}

	return false
}

// SetProgramStage gets a reference to the given ProgramStageParams and assigns it to the ProgramStage field.
func (o *ProgramRuleVariableParams) SetProgramStage(v ProgramStageParams) {
	o.ProgramStage = &v
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetSharing() Sharing {
	if o == nil || IsNil(o.Sharing) {
		var ret Sharing
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetSharingOk() (*Sharing, bool) {
	if o == nil || IsNil(o.Sharing) {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasSharing() bool {
	if o != nil && !IsNil(o.Sharing) {
		return true
	}

	return false
}

// SetSharing gets a reference to the given Sharing and assigns it to the Sharing field.
func (o *ProgramRuleVariableParams) SetSharing(v Sharing) {
	o.Sharing = &v
}

// GetTrackedEntityAttribute returns the TrackedEntityAttribute field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetTrackedEntityAttribute() TrackedEntityAttributeParams {
	if o == nil || IsNil(o.TrackedEntityAttribute) {
		var ret TrackedEntityAttributeParams
		return ret
	}
	return *o.TrackedEntityAttribute
}

// GetTrackedEntityAttributeOk returns a tuple with the TrackedEntityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetTrackedEntityAttributeOk() (*TrackedEntityAttributeParams, bool) {
	if o == nil || IsNil(o.TrackedEntityAttribute) {
		return nil, false
	}
	return o.TrackedEntityAttribute, true
}

// HasTrackedEntityAttribute returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasTrackedEntityAttribute() bool {
	if o != nil && !IsNil(o.TrackedEntityAttribute) {
		return true
	}

	return false
}

// SetTrackedEntityAttribute gets a reference to the given TrackedEntityAttributeParams and assigns it to the TrackedEntityAttribute field.
func (o *ProgramRuleVariableParams) SetTrackedEntityAttribute(v TrackedEntityAttributeParams) {
	o.TrackedEntityAttribute = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetTranslations() []Translation {
	if o == nil || IsNil(o.Translations) {
		var ret []Translation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetTranslationsOk() ([]Translation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []Translation and assigns it to the Translations field.
func (o *ProgramRuleVariableParams) SetTranslations(v []Translation) {
	o.Translations = v
}

// GetUseCodeForOptionSet returns the UseCodeForOptionSet field value if set, zero value otherwise.
func (o *ProgramRuleVariableParams) GetUseCodeForOptionSet() bool {
	if o == nil || IsNil(o.UseCodeForOptionSet) {
		var ret bool
		return ret
	}
	return *o.UseCodeForOptionSet
}

// GetUseCodeForOptionSetOk returns a tuple with the UseCodeForOptionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetUseCodeForOptionSetOk() (*bool, bool) {
	if o == nil || IsNil(o.UseCodeForOptionSet) {
		return nil, false
	}
	return o.UseCodeForOptionSet, true
}

// HasUseCodeForOptionSet returns a boolean if a field has been set.
func (o *ProgramRuleVariableParams) HasUseCodeForOptionSet() bool {
	if o != nil && !IsNil(o.UseCodeForOptionSet) {
		return true
	}

	return false
}

// SetUseCodeForOptionSet gets a reference to the given bool and assigns it to the UseCodeForOptionSet field.
func (o *ProgramRuleVariableParams) SetUseCodeForOptionSet(v bool) {
	o.UseCodeForOptionSet = &v
}

// GetValueType returns the ValueType field value
func (o *ProgramRuleVariableParams) GetValueType() ValueType {
	if o == nil {
		var ret ValueType
		return ret
	}

	return o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value
// and a boolean to check if the value has been set.
func (o *ProgramRuleVariableParams) GetValueTypeOk() (*ValueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueType, true
}

// SetValueType sets field value
func (o *ProgramRuleVariableParams) SetValueType(v ValueType) {
	o.ValueType = v
}

func (o ProgramRuleVariableParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramRuleVariableParams) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Program) {
		toSerialize["program"] = o.Program
	}
	toSerialize["programRuleVariableSourceType"] = o.ProgramRuleVariableSourceType
	if !IsNil(o.ProgramStage) {
		toSerialize["programStage"] = o.ProgramStage
	}
	if !IsNil(o.Sharing) {
		toSerialize["sharing"] = o.Sharing
	}
	if !IsNil(o.TrackedEntityAttribute) {
		toSerialize["trackedEntityAttribute"] = o.TrackedEntityAttribute
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	if !IsNil(o.UseCodeForOptionSet) {
		toSerialize["useCodeForOptionSet"] = o.UseCodeForOptionSet
	}
	toSerialize["valueType"] = o.ValueType
	return toSerialize, nil
}

func (o *ProgramRuleVariableParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"programRuleVariableSourceType",
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

	varProgramRuleVariableParams := _ProgramRuleVariableParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgramRuleVariableParams)

	if err != nil {
		return err
	}

	*o = ProgramRuleVariableParams(varProgramRuleVariableParams)

	return err
}

type NullableProgramRuleVariableParams struct {
	value *ProgramRuleVariableParams
	isSet bool
}

func (v NullableProgramRuleVariableParams) Get() *ProgramRuleVariableParams {
	return v.value
}

func (v *NullableProgramRuleVariableParams) Set(val *ProgramRuleVariableParams) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramRuleVariableParams) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramRuleVariableParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramRuleVariableParams(val *ProgramRuleVariableParams) *NullableProgramRuleVariableParams {
	return &NullableProgramRuleVariableParams{value: val, isSet: true}
}

func (v NullableProgramRuleVariableParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramRuleVariableParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
