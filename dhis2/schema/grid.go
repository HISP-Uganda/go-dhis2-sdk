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

// checks if the Grid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Grid{}

// Grid struct for Grid
type Grid struct {
	HeaderWidth        int32               `json:"headerWidth"`
	Headers            []GridHeader        `json:"headers,omitempty"`
	Height             int32               `json:"height"`
	InternalMetaData   *map[string]any     `json:"internalMetaData,omitempty"`
	LastDataRow        *bool               `json:"lastDataRow,omitempty"`
	MetaColumnIndexes  []int32             `json:"metaColumnIndexes,omitempty"`
	MetaData           *map[string]any     `json:"metaData,omitempty"`
	MetadataHeaders    []GridHeader        `json:"metadataHeaders,omitempty"`
	PerformanceMetrics *PerformanceMetrics `json:"performanceMetrics,omitempty"`
	Refs               []Reference         `json:"refs,omitempty"`
	// keys are class java.lang.Integer
	RowContext     *map[string]map[string]any `json:"rowContext,omitempty"`
	Rows           [][]any                    `json:"rows,omitempty"`
	Subtitle       *string                    `json:"subtitle,omitempty"`
	Table          *string                    `json:"table,omitempty"`
	Title          *string                    `json:"title,omitempty"`
	VisibleHeaders []GridHeader               `json:"visibleHeaders,omitempty"`
	VisibleRows    [][]any                    `json:"visibleRows,omitempty"`
	VisibleWidth   int32                      `json:"visibleWidth"`
	Width          int32                      `json:"width"`
}

type _Grid Grid

// NewGrid instantiates a new Grid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrid(headerWidth int32, height int32, visibleWidth int32, width int32) *Grid {
	this := Grid{}
	this.HeaderWidth = headerWidth
	this.Height = height
	this.VisibleWidth = visibleWidth
	this.Width = width
	return &this
}

// NewGridWithDefaults instantiates a new Grid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridWithDefaults() *Grid {
	this := Grid{}
	return &this
}

// GetHeaderWidth returns the HeaderWidth field value
func (o *Grid) GetHeaderWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HeaderWidth
}

// GetHeaderWidthOk returns a tuple with the HeaderWidth field value
// and a boolean to check if the value has been set.
func (o *Grid) GetHeaderWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderWidth, true
}

// SetHeaderWidth sets field value
func (o *Grid) SetHeaderWidth(v int32) {
	o.HeaderWidth = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Grid) GetHeaders() []GridHeader {
	if o == nil || IsNil(o.Headers) {
		var ret []GridHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetHeadersOk() ([]GridHeader, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Grid) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []GridHeader and assigns it to the Headers field.
func (o *Grid) SetHeaders(v []GridHeader) {
	o.Headers = v
}

// GetHeight returns the Height field value
func (o *Grid) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *Grid) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *Grid) SetHeight(v int32) {
	o.Height = v
}

// GetInternalMetaData returns the InternalMetaData field value if set, zero value otherwise.
func (o *Grid) GetInternalMetaData() map[string]any {
	if o == nil || IsNil(o.InternalMetaData) {
		var ret map[string]any
		return ret
	}
	return *o.InternalMetaData
}

// GetInternalMetaDataOk returns a tuple with the InternalMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetInternalMetaDataOk() (*map[string]any, bool) {
	if o == nil || IsNil(o.InternalMetaData) {
		return nil, false
	}
	return o.InternalMetaData, true
}

// HasInternalMetaData returns a boolean if a field has been set.
func (o *Grid) HasInternalMetaData() bool {
	if o != nil && !IsNil(o.InternalMetaData) {
		return true
	}

	return false
}

// SetInternalMetaData gets a reference to the given map[string]any and assigns it to the InternalMetaData field.
func (o *Grid) SetInternalMetaData(v map[string]any) {
	o.InternalMetaData = &v
}

// GetLastDataRow returns the LastDataRow field value if set, zero value otherwise.
func (o *Grid) GetLastDataRow() bool {
	if o == nil || IsNil(o.LastDataRow) {
		var ret bool
		return ret
	}
	return *o.LastDataRow
}

// GetLastDataRowOk returns a tuple with the LastDataRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetLastDataRowOk() (*bool, bool) {
	if o == nil || IsNil(o.LastDataRow) {
		return nil, false
	}
	return o.LastDataRow, true
}

// HasLastDataRow returns a boolean if a field has been set.
func (o *Grid) HasLastDataRow() bool {
	if o != nil && !IsNil(o.LastDataRow) {
		return true
	}

	return false
}

// SetLastDataRow gets a reference to the given bool and assigns it to the LastDataRow field.
func (o *Grid) SetLastDataRow(v bool) {
	o.LastDataRow = &v
}

// GetMetaColumnIndexes returns the MetaColumnIndexes field value if set, zero value otherwise.
func (o *Grid) GetMetaColumnIndexes() []int32 {
	if o == nil || IsNil(o.MetaColumnIndexes) {
		var ret []int32
		return ret
	}
	return o.MetaColumnIndexes
}

// GetMetaColumnIndexesOk returns a tuple with the MetaColumnIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetMetaColumnIndexesOk() ([]int32, bool) {
	if o == nil || IsNil(o.MetaColumnIndexes) {
		return nil, false
	}
	return o.MetaColumnIndexes, true
}

// HasMetaColumnIndexes returns a boolean if a field has been set.
func (o *Grid) HasMetaColumnIndexes() bool {
	if o != nil && !IsNil(o.MetaColumnIndexes) {
		return true
	}

	return false
}

// SetMetaColumnIndexes gets a reference to the given []int32 and assigns it to the MetaColumnIndexes field.
func (o *Grid) SetMetaColumnIndexes(v []int32) {
	o.MetaColumnIndexes = v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *Grid) GetMetaData() map[string]any {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]any
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetMetaDataOk() (*map[string]any, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *Grid) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]any and assigns it to the MetaData field.
func (o *Grid) SetMetaData(v map[string]any) {
	o.MetaData = &v
}

// GetMetadataHeaders returns the MetadataHeaders field value if set, zero value otherwise.
func (o *Grid) GetMetadataHeaders() []GridHeader {
	if o == nil || IsNil(o.MetadataHeaders) {
		var ret []GridHeader
		return ret
	}
	return o.MetadataHeaders
}

// GetMetadataHeadersOk returns a tuple with the MetadataHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetMetadataHeadersOk() ([]GridHeader, bool) {
	if o == nil || IsNil(o.MetadataHeaders) {
		return nil, false
	}
	return o.MetadataHeaders, true
}

// HasMetadataHeaders returns a boolean if a field has been set.
func (o *Grid) HasMetadataHeaders() bool {
	if o != nil && !IsNil(o.MetadataHeaders) {
		return true
	}

	return false
}

// SetMetadataHeaders gets a reference to the given []GridHeader and assigns it to the MetadataHeaders field.
func (o *Grid) SetMetadataHeaders(v []GridHeader) {
	o.MetadataHeaders = v
}

// GetPerformanceMetrics returns the PerformanceMetrics field value if set, zero value otherwise.
func (o *Grid) GetPerformanceMetrics() PerformanceMetrics {
	if o == nil || IsNil(o.PerformanceMetrics) {
		var ret PerformanceMetrics
		return ret
	}
	return *o.PerformanceMetrics
}

// GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetPerformanceMetricsOk() (*PerformanceMetrics, bool) {
	if o == nil || IsNil(o.PerformanceMetrics) {
		return nil, false
	}
	return o.PerformanceMetrics, true
}

// HasPerformanceMetrics returns a boolean if a field has been set.
func (o *Grid) HasPerformanceMetrics() bool {
	if o != nil && !IsNil(o.PerformanceMetrics) {
		return true
	}

	return false
}

// SetPerformanceMetrics gets a reference to the given PerformanceMetrics and assigns it to the PerformanceMetrics field.
func (o *Grid) SetPerformanceMetrics(v PerformanceMetrics) {
	o.PerformanceMetrics = &v
}

// GetRefs returns the Refs field value if set, zero value otherwise.
func (o *Grid) GetRefs() []Reference {
	if o == nil || IsNil(o.Refs) {
		var ret []Reference
		return ret
	}
	return o.Refs
}

// GetRefsOk returns a tuple with the Refs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetRefsOk() ([]Reference, bool) {
	if o == nil || IsNil(o.Refs) {
		return nil, false
	}
	return o.Refs, true
}

// HasRefs returns a boolean if a field has been set.
func (o *Grid) HasRefs() bool {
	if o != nil && !IsNil(o.Refs) {
		return true
	}

	return false
}

// SetRefs gets a reference to the given []Reference and assigns it to the Refs field.
func (o *Grid) SetRefs(v []Reference) {
	o.Refs = v
}

// GetRowContext returns the RowContext field value if set, zero value otherwise.
func (o *Grid) GetRowContext() map[string]map[string]any {
	if o == nil || IsNil(o.RowContext) {
		var ret map[string]map[string]any
		return ret
	}
	return *o.RowContext
}

// GetRowContextOk returns a tuple with the RowContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetRowContextOk() (*map[string]map[string]any, bool) {
	if o == nil || IsNil(o.RowContext) {
		return nil, false
	}
	return o.RowContext, true
}

// HasRowContext returns a boolean if a field has been set.
func (o *Grid) HasRowContext() bool {
	if o != nil && !IsNil(o.RowContext) {
		return true
	}

	return false
}

// SetRowContext gets a reference to the given map[string]map[string]any and assigns it to the RowContext field.
func (o *Grid) SetRowContext(v map[string]map[string]any) {
	o.RowContext = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *Grid) GetRows() [][]any {
	if o == nil || IsNil(o.Rows) {
		var ret [][]any
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetRowsOk() ([][]any, bool) {
	if o == nil || IsNil(o.Rows) {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *Grid) HasRows() bool {
	if o != nil && !IsNil(o.Rows) {
		return true
	}

	return false
}

// SetRows gets a reference to the given [][]any and assigns it to the Rows field.
func (o *Grid) SetRows(v [][]any) {
	o.Rows = v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *Grid) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *Grid) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *Grid) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *Grid) GetTable() string {
	if o == nil || IsNil(o.Table) {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetTableOk() (*string, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *Grid) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *Grid) SetTable(v string) {
	o.Table = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Grid) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Grid) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Grid) SetTitle(v string) {
	o.Title = &v
}

// GetVisibleHeaders returns the VisibleHeaders field value if set, zero value otherwise.
func (o *Grid) GetVisibleHeaders() []GridHeader {
	if o == nil || IsNil(o.VisibleHeaders) {
		var ret []GridHeader
		return ret
	}
	return o.VisibleHeaders
}

// GetVisibleHeadersOk returns a tuple with the VisibleHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetVisibleHeadersOk() ([]GridHeader, bool) {
	if o == nil || IsNil(o.VisibleHeaders) {
		return nil, false
	}
	return o.VisibleHeaders, true
}

// HasVisibleHeaders returns a boolean if a field has been set.
func (o *Grid) HasVisibleHeaders() bool {
	if o != nil && !IsNil(o.VisibleHeaders) {
		return true
	}

	return false
}

// SetVisibleHeaders gets a reference to the given []GridHeader and assigns it to the VisibleHeaders field.
func (o *Grid) SetVisibleHeaders(v []GridHeader) {
	o.VisibleHeaders = v
}

// GetVisibleRows returns the VisibleRows field value if set, zero value otherwise.
func (o *Grid) GetVisibleRows() [][]any {
	if o == nil || IsNil(o.VisibleRows) {
		var ret [][]any
		return ret
	}
	return o.VisibleRows
}

// GetVisibleRowsOk returns a tuple with the VisibleRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetVisibleRowsOk() ([][]any, bool) {
	if o == nil || IsNil(o.VisibleRows) {
		return nil, false
	}
	return o.VisibleRows, true
}

// HasVisibleRows returns a boolean if a field has been set.
func (o *Grid) HasVisibleRows() bool {
	if o != nil && !IsNil(o.VisibleRows) {
		return true
	}

	return false
}

// SetVisibleRows gets a reference to the given [][]any and assigns it to the VisibleRows field.
func (o *Grid) SetVisibleRows(v [][]any) {
	o.VisibleRows = v
}

// GetVisibleWidth returns the VisibleWidth field value
func (o *Grid) GetVisibleWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VisibleWidth
}

// GetVisibleWidthOk returns a tuple with the VisibleWidth field value
// and a boolean to check if the value has been set.
func (o *Grid) GetVisibleWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisibleWidth, true
}

// SetVisibleWidth sets field value
func (o *Grid) SetVisibleWidth(v int32) {
	o.VisibleWidth = v
}

// GetWidth returns the Width field value
func (o *Grid) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *Grid) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *Grid) SetWidth(v int32) {
	o.Width = v
}

func (o Grid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Grid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headerWidth"] = o.HeaderWidth
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["height"] = o.Height
	if !IsNil(o.InternalMetaData) {
		toSerialize["internalMetaData"] = o.InternalMetaData
	}
	if !IsNil(o.LastDataRow) {
		toSerialize["lastDataRow"] = o.LastDataRow
	}
	if !IsNil(o.MetaColumnIndexes) {
		toSerialize["metaColumnIndexes"] = o.MetaColumnIndexes
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	if !IsNil(o.MetadataHeaders) {
		toSerialize["metadataHeaders"] = o.MetadataHeaders
	}
	if !IsNil(o.PerformanceMetrics) {
		toSerialize["performanceMetrics"] = o.PerformanceMetrics
	}
	if !IsNil(o.Refs) {
		toSerialize["refs"] = o.Refs
	}
	if !IsNil(o.RowContext) {
		toSerialize["rowContext"] = o.RowContext
	}
	if !IsNil(o.Rows) {
		toSerialize["rows"] = o.Rows
	}
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.VisibleHeaders) {
		toSerialize["visibleHeaders"] = o.VisibleHeaders
	}
	if !IsNil(o.VisibleRows) {
		toSerialize["visibleRows"] = o.VisibleRows
	}
	toSerialize["visibleWidth"] = o.VisibleWidth
	toSerialize["width"] = o.Width
	return toSerialize, nil
}

func (o *Grid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headerWidth",
		"height",
		"visibleWidth",
		"width",
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

	varGrid := _Grid{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGrid)

	if err != nil {
		return err
	}

	*o = Grid(varGrid)

	return err
}

type NullableGrid struct {
	value *Grid
	isSet bool
}

func (v NullableGrid) Get() *Grid {
	return v.value
}

func (v *NullableGrid) Set(val *Grid) {
	v.value = val
	v.isSet = true
}

func (v NullableGrid) IsSet() bool {
	return v.isSet
}

func (v *NullableGrid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrid(val *Grid) *NullableGrid {
	return &NullableGrid{value: val, isSet: true}
}

func (v NullableGrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
