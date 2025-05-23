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
)

// checks if the MetadataImportParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataImportParams{}

// MetadataImportParams struct for MetadataImportParams
type MetadataImportParams struct {
	Async              *bool             `json:"async,omitempty"`
	AtomicMode         AtomicMode        `json:"atomicMode"`
	FlushMode          FlushMode         `json:"flushMode"`
	Identifier         PreheatIdentifier `json:"identifier"`
	ImportMode         ObjectBundleMode  `json:"importMode"`
	ImportReportMode   ImportReportMode  `json:"importReportMode"`
	ImportStrategy     ImportStrategy    `json:"importStrategy"`
	MetadataSyncImport *bool             `json:"metadataSyncImport,omitempty"`
	// A UID for an User object   (Java name `org.hisp.dhis.user.User`)
	OverrideUser    *string     `json:"overrideUser,omitempty"`
	PreheatMode     PreheatMode `json:"preheatMode"`
	SkipSharing     *bool       `json:"skipSharing,omitempty"`
	SkipTranslation *bool       `json:"skipTranslation,omitempty"`
	SkipValidation  *bool       `json:"skipValidation,omitempty"`
	// A UID for an User object   (Java name `org.hisp.dhis.user.User`)
	User             *string          `json:"user,omitempty"`
	UserOverrideMode UserOverrideMode `json:"userOverrideMode"`
}

type _MetadataImportParams MetadataImportParams

// NewMetadataImportParams instantiates a new MetadataImportParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataImportParams(atomicMode AtomicMode, flushMode FlushMode, identifier PreheatIdentifier, importMode ObjectBundleMode, importReportMode ImportReportMode, importStrategy ImportStrategy, preheatMode PreheatMode, userOverrideMode UserOverrideMode) *MetadataImportParams {
	this := MetadataImportParams{}
	this.AtomicMode = atomicMode
	this.FlushMode = flushMode
	this.Identifier = identifier
	this.ImportMode = importMode
	this.ImportReportMode = importReportMode
	this.ImportStrategy = importStrategy
	this.PreheatMode = preheatMode
	this.UserOverrideMode = userOverrideMode
	return &this
}

// NewMetadataImportParamsWithDefaults instantiates a new MetadataImportParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataImportParamsWithDefaults() *MetadataImportParams {
	this := MetadataImportParams{}
	return &this
}

// GetAsync returns the Async field value if set, zero value otherwise.
func (o *MetadataImportParams) GetAsync() bool {
	if o == nil || IsNil(o.Async) {
		var ret bool
		return ret
	}
	return *o.Async
}

// GetAsyncOk returns a tuple with the Async field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetAsyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Async) {
		return nil, false
	}
	return o.Async, true
}

// HasAsync returns a boolean if a field has been set.
func (o *MetadataImportParams) HasAsync() bool {
	if o != nil && !IsNil(o.Async) {
		return true
	}

	return false
}

// SetAsync gets a reference to the given bool and assigns it to the Async field.
func (o *MetadataImportParams) SetAsync(v bool) {
	o.Async = &v
}

// GetAtomicMode returns the AtomicMode field value
func (o *MetadataImportParams) GetAtomicMode() AtomicMode {
	if o == nil {
		var ret AtomicMode
		return ret
	}

	return o.AtomicMode
}

// GetAtomicModeOk returns a tuple with the AtomicMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetAtomicModeOk() (*AtomicMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtomicMode, true
}

// SetAtomicMode sets field value
func (o *MetadataImportParams) SetAtomicMode(v AtomicMode) {
	o.AtomicMode = v
}

// GetFlushMode returns the FlushMode field value
func (o *MetadataImportParams) GetFlushMode() FlushMode {
	if o == nil {
		var ret FlushMode
		return ret
	}

	return o.FlushMode
}

// GetFlushModeOk returns a tuple with the FlushMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetFlushModeOk() (*FlushMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlushMode, true
}

// SetFlushMode sets field value
func (o *MetadataImportParams) SetFlushMode(v FlushMode) {
	o.FlushMode = v
}

// GetIdentifier returns the Identifier field value
func (o *MetadataImportParams) GetIdentifier() PreheatIdentifier {
	if o == nil {
		var ret PreheatIdentifier
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetIdentifierOk() (*PreheatIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *MetadataImportParams) SetIdentifier(v PreheatIdentifier) {
	o.Identifier = v
}

// GetImportMode returns the ImportMode field value
func (o *MetadataImportParams) GetImportMode() ObjectBundleMode {
	if o == nil {
		var ret ObjectBundleMode
		return ret
	}

	return o.ImportMode
}

// GetImportModeOk returns a tuple with the ImportMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetImportModeOk() (*ObjectBundleMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportMode, true
}

// SetImportMode sets field value
func (o *MetadataImportParams) SetImportMode(v ObjectBundleMode) {
	o.ImportMode = v
}

// GetImportReportMode returns the ImportReportMode field value
func (o *MetadataImportParams) GetImportReportMode() ImportReportMode {
	if o == nil {
		var ret ImportReportMode
		return ret
	}

	return o.ImportReportMode
}

// GetImportReportModeOk returns a tuple with the ImportReportMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetImportReportModeOk() (*ImportReportMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportReportMode, true
}

// SetImportReportMode sets field value
func (o *MetadataImportParams) SetImportReportMode(v ImportReportMode) {
	o.ImportReportMode = v
}

// GetImportStrategy returns the ImportStrategy field value
func (o *MetadataImportParams) GetImportStrategy() ImportStrategy {
	if o == nil {
		var ret ImportStrategy
		return ret
	}

	return o.ImportStrategy
}

// GetImportStrategyOk returns a tuple with the ImportStrategy field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetImportStrategyOk() (*ImportStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportStrategy, true
}

// SetImportStrategy sets field value
func (o *MetadataImportParams) SetImportStrategy(v ImportStrategy) {
	o.ImportStrategy = v
}

// GetMetadataSyncImport returns the MetadataSyncImport field value if set, zero value otherwise.
func (o *MetadataImportParams) GetMetadataSyncImport() bool {
	if o == nil || IsNil(o.MetadataSyncImport) {
		var ret bool
		return ret
	}
	return *o.MetadataSyncImport
}

// GetMetadataSyncImportOk returns a tuple with the MetadataSyncImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetMetadataSyncImportOk() (*bool, bool) {
	if o == nil || IsNil(o.MetadataSyncImport) {
		return nil, false
	}
	return o.MetadataSyncImport, true
}

// HasMetadataSyncImport returns a boolean if a field has been set.
func (o *MetadataImportParams) HasMetadataSyncImport() bool {
	if o != nil && !IsNil(o.MetadataSyncImport) {
		return true
	}

	return false
}

// SetMetadataSyncImport gets a reference to the given bool and assigns it to the MetadataSyncImport field.
func (o *MetadataImportParams) SetMetadataSyncImport(v bool) {
	o.MetadataSyncImport = &v
}

// GetOverrideUser returns the OverrideUser field value if set, zero value otherwise.
func (o *MetadataImportParams) GetOverrideUser() string {
	if o == nil || IsNil(o.OverrideUser) {
		var ret string
		return ret
	}
	return *o.OverrideUser
}

// GetOverrideUserOk returns a tuple with the OverrideUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetOverrideUserOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideUser) {
		return nil, false
	}
	return o.OverrideUser, true
}

// HasOverrideUser returns a boolean if a field has been set.
func (o *MetadataImportParams) HasOverrideUser() bool {
	if o != nil && !IsNil(o.OverrideUser) {
		return true
	}

	return false
}

// SetOverrideUser gets a reference to the given string and assigns it to the OverrideUser field.
func (o *MetadataImportParams) SetOverrideUser(v string) {
	o.OverrideUser = &v
}

// GetPreheatMode returns the PreheatMode field value
func (o *MetadataImportParams) GetPreheatMode() PreheatMode {
	if o == nil {
		var ret PreheatMode
		return ret
	}

	return o.PreheatMode
}

// GetPreheatModeOk returns a tuple with the PreheatMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetPreheatModeOk() (*PreheatMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreheatMode, true
}

// SetPreheatMode sets field value
func (o *MetadataImportParams) SetPreheatMode(v PreheatMode) {
	o.PreheatMode = v
}

// GetSkipSharing returns the SkipSharing field value if set, zero value otherwise.
func (o *MetadataImportParams) GetSkipSharing() bool {
	if o == nil || IsNil(o.SkipSharing) {
		var ret bool
		return ret
	}
	return *o.SkipSharing
}

// GetSkipSharingOk returns a tuple with the SkipSharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetSkipSharingOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipSharing) {
		return nil, false
	}
	return o.SkipSharing, true
}

// HasSkipSharing returns a boolean if a field has been set.
func (o *MetadataImportParams) HasSkipSharing() bool {
	if o != nil && !IsNil(o.SkipSharing) {
		return true
	}

	return false
}

// SetSkipSharing gets a reference to the given bool and assigns it to the SkipSharing field.
func (o *MetadataImportParams) SetSkipSharing(v bool) {
	o.SkipSharing = &v
}

// GetSkipTranslation returns the SkipTranslation field value if set, zero value otherwise.
func (o *MetadataImportParams) GetSkipTranslation() bool {
	if o == nil || IsNil(o.SkipTranslation) {
		var ret bool
		return ret
	}
	return *o.SkipTranslation
}

// GetSkipTranslationOk returns a tuple with the SkipTranslation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetSkipTranslationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipTranslation) {
		return nil, false
	}
	return o.SkipTranslation, true
}

// HasSkipTranslation returns a boolean if a field has been set.
func (o *MetadataImportParams) HasSkipTranslation() bool {
	if o != nil && !IsNil(o.SkipTranslation) {
		return true
	}

	return false
}

// SetSkipTranslation gets a reference to the given bool and assigns it to the SkipTranslation field.
func (o *MetadataImportParams) SetSkipTranslation(v bool) {
	o.SkipTranslation = &v
}

// GetSkipValidation returns the SkipValidation field value if set, zero value otherwise.
func (o *MetadataImportParams) GetSkipValidation() bool {
	if o == nil || IsNil(o.SkipValidation) {
		var ret bool
		return ret
	}
	return *o.SkipValidation
}

// GetSkipValidationOk returns a tuple with the SkipValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetSkipValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipValidation) {
		return nil, false
	}
	return o.SkipValidation, true
}

// HasSkipValidation returns a boolean if a field has been set.
func (o *MetadataImportParams) HasSkipValidation() bool {
	if o != nil && !IsNil(o.SkipValidation) {
		return true
	}

	return false
}

// SetSkipValidation gets a reference to the given bool and assigns it to the SkipValidation field.
func (o *MetadataImportParams) SetSkipValidation(v bool) {
	o.SkipValidation = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MetadataImportParams) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MetadataImportParams) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *MetadataImportParams) SetUser(v string) {
	o.User = &v
}

// GetUserOverrideMode returns the UserOverrideMode field value
func (o *MetadataImportParams) GetUserOverrideMode() UserOverrideMode {
	if o == nil {
		var ret UserOverrideMode
		return ret
	}

	return o.UserOverrideMode
}

// GetUserOverrideModeOk returns a tuple with the UserOverrideMode field value
// and a boolean to check if the value has been set.
func (o *MetadataImportParams) GetUserOverrideModeOk() (*UserOverrideMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserOverrideMode, true
}

// SetUserOverrideMode sets field value
func (o *MetadataImportParams) SetUserOverrideMode(v UserOverrideMode) {
	o.UserOverrideMode = v
}

func (o MetadataImportParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataImportParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Async) {
		toSerialize["async"] = o.Async
	}
	toSerialize["atomicMode"] = o.AtomicMode
	toSerialize["flushMode"] = o.FlushMode
	toSerialize["identifier"] = o.Identifier
	toSerialize["importMode"] = o.ImportMode
	toSerialize["importReportMode"] = o.ImportReportMode
	toSerialize["importStrategy"] = o.ImportStrategy
	if !IsNil(o.MetadataSyncImport) {
		toSerialize["metadataSyncImport"] = o.MetadataSyncImport
	}
	if !IsNil(o.OverrideUser) {
		toSerialize["overrideUser"] = o.OverrideUser
	}
	toSerialize["preheatMode"] = o.PreheatMode
	if !IsNil(o.SkipSharing) {
		toSerialize["skipSharing"] = o.SkipSharing
	}
	if !IsNil(o.SkipTranslation) {
		toSerialize["skipTranslation"] = o.SkipTranslation
	}
	if !IsNil(o.SkipValidation) {
		toSerialize["skipValidation"] = o.SkipValidation
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	toSerialize["userOverrideMode"] = o.UserOverrideMode
	return toSerialize, nil
}

func (o *MetadataImportParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"atomicMode",
		"flushMode",
		"identifier",
		"importMode",
		"importReportMode",
		"importStrategy",
		"preheatMode",
		"userOverrideMode",
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

	varMetadataImportParams := _MetadataImportParams{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataImportParams)

	if err != nil {
		return err
	}

	*o = MetadataImportParams(varMetadataImportParams)

	return err
}

type NullableMetadataImportParams struct {
	value *MetadataImportParams
	isSet bool
}

func (v NullableMetadataImportParams) Get() *MetadataImportParams {
	return v.value
}

func (v *NullableMetadataImportParams) Set(val *MetadataImportParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataImportParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataImportParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataImportParams(val *MetadataImportParams) *NullableMetadataImportParams {
	return &NullableMetadataImportParams{value: val, isSet: true}
}

func (v NullableMetadataImportParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataImportParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
