/*
DHIS2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.42
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
	"fmt"
)

// ProgramPropertyNames the model 'ProgramPropertyNames'
type ProgramPropertyNames string

// List of ProgramPropertyNames
const (
	PROGRAMPROPERTYNAMES_ACCESS                                 ProgramPropertyNames = "access"
	PROGRAMPROPERTYNAMES_ACCESS_LEVEL                           ProgramPropertyNames = "accessLevel"
	PROGRAMPROPERTYNAMES_ATTRIBUTE_VALUES                       ProgramPropertyNames = "attributeValues"
	PROGRAMPROPERTYNAMES_CATEGORY_COMBO                         ProgramPropertyNames = "categoryCombo"
	PROGRAMPROPERTYNAMES_CATEGORY_MAPPINGS                      ProgramPropertyNames = "categoryMappings"
	PROGRAMPROPERTYNAMES_CODE                                   ProgramPropertyNames = "code"
	PROGRAMPROPERTYNAMES_COMPLETE_EVENTS_EXPIRY_DAYS            ProgramPropertyNames = "completeEventsExpiryDays"
	PROGRAMPROPERTYNAMES_CREATED                                ProgramPropertyNames = "created"
	PROGRAMPROPERTYNAMES_CREATED_BY                             ProgramPropertyNames = "createdBy"
	PROGRAMPROPERTYNAMES_DATA_ENTRY_FORM                        ProgramPropertyNames = "dataEntryForm"
	PROGRAMPROPERTYNAMES_DESCRIPTION                            ProgramPropertyNames = "description"
	PROGRAMPROPERTYNAMES_DISPLAY_DESCRIPTION                    ProgramPropertyNames = "displayDescription"
	PROGRAMPROPERTYNAMES_DISPLAY_ENROLLMENT_DATE_LABEL          ProgramPropertyNames = "displayEnrollmentDateLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_ENROLLMENT_LABEL               ProgramPropertyNames = "displayEnrollmentLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_EVENT_LABEL                    ProgramPropertyNames = "displayEventLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_FOLLOW_UP_LABEL                ProgramPropertyNames = "displayFollowUpLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_FORM_NAME                      ProgramPropertyNames = "displayFormName"
	PROGRAMPROPERTYNAMES_DISPLAY_FRONT_PAGE_LIST                ProgramPropertyNames = "displayFrontPageList"
	PROGRAMPROPERTYNAMES_DISPLAY_INCIDENT_DATE                  ProgramPropertyNames = "displayIncidentDate"
	PROGRAMPROPERTYNAMES_DISPLAY_INCIDENT_DATE_LABEL            ProgramPropertyNames = "displayIncidentDateLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_NAME                           ProgramPropertyNames = "displayName"
	PROGRAMPROPERTYNAMES_DISPLAY_NOTE_LABEL                     ProgramPropertyNames = "displayNoteLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_ORG_UNIT_LABEL                 ProgramPropertyNames = "displayOrgUnitLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_PROGRAM_STAGE_LABEL            ProgramPropertyNames = "displayProgramStageLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_RELATIONSHIP_LABEL             ProgramPropertyNames = "displayRelationshipLabel"
	PROGRAMPROPERTYNAMES_DISPLAY_SHORT_NAME                     ProgramPropertyNames = "displayShortName"
	PROGRAMPROPERTYNAMES_DISPLAY_TRACKED_ENTITY_ATTRIBUTE_LABEL ProgramPropertyNames = "displayTrackedEntityAttributeLabel"
	PROGRAMPROPERTYNAMES_ENROLLMENT_DATE_LABEL                  ProgramPropertyNames = "enrollmentDateLabel"
	PROGRAMPROPERTYNAMES_ENROLLMENT_LABEL                       ProgramPropertyNames = "enrollmentLabel"
	PROGRAMPROPERTYNAMES_EVENT_LABEL                            ProgramPropertyNames = "eventLabel"
	PROGRAMPROPERTYNAMES_EXPIRY_DAYS                            ProgramPropertyNames = "expiryDays"
	PROGRAMPROPERTYNAMES_EXPIRY_PERIOD_TYPE                     ProgramPropertyNames = "expiryPeriodType"
	PROGRAMPROPERTYNAMES_FAVORITE                               ProgramPropertyNames = "favorite"
	PROGRAMPROPERTYNAMES_FAVORITES                              ProgramPropertyNames = "favorites"
	PROGRAMPROPERTYNAMES_FEATURE_TYPE                           ProgramPropertyNames = "featureType"
	PROGRAMPROPERTYNAMES_FOLLOW_UP_LABEL                        ProgramPropertyNames = "followUpLabel"
	PROGRAMPROPERTYNAMES_FORM_NAME                              ProgramPropertyNames = "formName"
	PROGRAMPROPERTYNAMES_HREF                                   ProgramPropertyNames = "href"
	PROGRAMPROPERTYNAMES_ID                                     ProgramPropertyNames = "id"
	PROGRAMPROPERTYNAMES_IGNORE_OVERDUE_EVENTS                  ProgramPropertyNames = "ignoreOverdueEvents"
	PROGRAMPROPERTYNAMES_INCIDENT_DATE_LABEL                    ProgramPropertyNames = "incidentDateLabel"
	PROGRAMPROPERTYNAMES_LAST_UPDATED                           ProgramPropertyNames = "lastUpdated"
	PROGRAMPROPERTYNAMES_LAST_UPDATED_BY                        ProgramPropertyNames = "lastUpdatedBy"
	PROGRAMPROPERTYNAMES_MAX_TEI_COUNT_TO_RETURN                ProgramPropertyNames = "maxTeiCountToReturn"
	PROGRAMPROPERTYNAMES_MIN_ATTRIBUTES_REQUIRED_TO_SEARCH      ProgramPropertyNames = "minAttributesRequiredToSearch"
	PROGRAMPROPERTYNAMES_NAME                                   ProgramPropertyNames = "name"
	PROGRAMPROPERTYNAMES_NOTE_LABEL                             ProgramPropertyNames = "noteLabel"
	PROGRAMPROPERTYNAMES_NOTIFICATION_TEMPLATES                 ProgramPropertyNames = "notificationTemplates"
	PROGRAMPROPERTYNAMES_ONLY_ENROLL_ONCE                       ProgramPropertyNames = "onlyEnrollOnce"
	PROGRAMPROPERTYNAMES_OPEN_DAYS_AFTER_CO_END_DATE            ProgramPropertyNames = "openDaysAfterCoEndDate"
	PROGRAMPROPERTYNAMES_ORG_UNIT_LABEL                         ProgramPropertyNames = "orgUnitLabel"
	PROGRAMPROPERTYNAMES_ORGANISATION_UNITS                     ProgramPropertyNames = "organisationUnits"
	PROGRAMPROPERTYNAMES_PROGRAM_INDICATORS                     ProgramPropertyNames = "programIndicators"
	PROGRAMPROPERTYNAMES_PROGRAM_RULE_VARIABLES                 ProgramPropertyNames = "programRuleVariables"
	PROGRAMPROPERTYNAMES_PROGRAM_SECTIONS                       ProgramPropertyNames = "programSections"
	PROGRAMPROPERTYNAMES_PROGRAM_STAGE_LABEL                    ProgramPropertyNames = "programStageLabel"
	PROGRAMPROPERTYNAMES_PROGRAM_STAGES                         ProgramPropertyNames = "programStages"
	PROGRAMPROPERTYNAMES_PROGRAM_TRACKED_ENTITY_ATTRIBUTES      ProgramPropertyNames = "programTrackedEntityAttributes"
	PROGRAMPROPERTYNAMES_PROGRAM_TYPE                           ProgramPropertyNames = "programType"
	PROGRAMPROPERTYNAMES_REGISTRATION                           ProgramPropertyNames = "registration"
	PROGRAMPROPERTYNAMES_RELATED_PROGRAM                        ProgramPropertyNames = "relatedProgram"
	PROGRAMPROPERTYNAMES_RELATIONSHIP_LABEL                     ProgramPropertyNames = "relationshipLabel"
	PROGRAMPROPERTYNAMES_SELECT_ENROLLMENT_DATES_IN_FUTURE      ProgramPropertyNames = "selectEnrollmentDatesInFuture"
	PROGRAMPROPERTYNAMES_SELECT_INCIDENT_DATES_IN_FUTURE        ProgramPropertyNames = "selectIncidentDatesInFuture"
	PROGRAMPROPERTYNAMES_SHARING                                ProgramPropertyNames = "sharing"
	PROGRAMPROPERTYNAMES_SHORT_NAME                             ProgramPropertyNames = "shortName"
	PROGRAMPROPERTYNAMES_SKIP_OFFLINE                           ProgramPropertyNames = "skipOffline"
	PROGRAMPROPERTYNAMES_STYLE                                  ProgramPropertyNames = "style"
	PROGRAMPROPERTYNAMES_TRACKED_ENTITY_ATTRIBUTE_LABEL         ProgramPropertyNames = "trackedEntityAttributeLabel"
	PROGRAMPROPERTYNAMES_TRACKED_ENTITY_TYPE                    ProgramPropertyNames = "trackedEntityType"
	PROGRAMPROPERTYNAMES_TRANSLATIONS                           ProgramPropertyNames = "translations"
	PROGRAMPROPERTYNAMES_USE_FIRST_STAGE_DURING_REGISTRATION    ProgramPropertyNames = "useFirstStageDuringRegistration"
	PROGRAMPROPERTYNAMES_USER_ROLES                             ProgramPropertyNames = "userRoles"
	PROGRAMPROPERTYNAMES_VERSION                                ProgramPropertyNames = "version"
	PROGRAMPROPERTYNAMES_WITHOUT_REGISTRATION                   ProgramPropertyNames = "withoutRegistration"
)

// All allowed values of ProgramPropertyNames enum
var AllowedProgramPropertyNamesEnumValues = []ProgramPropertyNames{
	"access",
	"accessLevel",
	"attributeValues",
	"categoryCombo",
	"categoryMappings",
	"code",
	"completeEventsExpiryDays",
	"created",
	"createdBy",
	"dataEntryForm",
	"description",
	"displayDescription",
	"displayEnrollmentDateLabel",
	"displayEnrollmentLabel",
	"displayEventLabel",
	"displayFollowUpLabel",
	"displayFormName",
	"displayFrontPageList",
	"displayIncidentDate",
	"displayIncidentDateLabel",
	"displayName",
	"displayNoteLabel",
	"displayOrgUnitLabel",
	"displayProgramStageLabel",
	"displayRelationshipLabel",
	"displayShortName",
	"displayTrackedEntityAttributeLabel",
	"enrollmentDateLabel",
	"enrollmentLabel",
	"eventLabel",
	"expiryDays",
	"expiryPeriodType",
	"favorite",
	"favorites",
	"featureType",
	"followUpLabel",
	"formName",
	"href",
	"id",
	"ignoreOverdueEvents",
	"incidentDateLabel",
	"lastUpdated",
	"lastUpdatedBy",
	"maxTeiCountToReturn",
	"minAttributesRequiredToSearch",
	"name",
	"noteLabel",
	"notificationTemplates",
	"onlyEnrollOnce",
	"openDaysAfterCoEndDate",
	"orgUnitLabel",
	"organisationUnits",
	"programIndicators",
	"programRuleVariables",
	"programSections",
	"programStageLabel",
	"programStages",
	"programTrackedEntityAttributes",
	"programType",
	"registration",
	"relatedProgram",
	"relationshipLabel",
	"selectEnrollmentDatesInFuture",
	"selectIncidentDatesInFuture",
	"sharing",
	"shortName",
	"skipOffline",
	"style",
	"trackedEntityAttributeLabel",
	"trackedEntityType",
	"translations",
	"useFirstStageDuringRegistration",
	"userRoles",
	"version",
	"withoutRegistration",
}

func (v *ProgramPropertyNames) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProgramPropertyNames(value)
	for _, existing := range AllowedProgramPropertyNamesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProgramPropertyNames", value)
}

// NewProgramPropertyNamesFromValue returns a pointer to a valid ProgramPropertyNames
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgramPropertyNamesFromValue(v string) (*ProgramPropertyNames, error) {
	ev := ProgramPropertyNames(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProgramPropertyNames: valid values are %v", v, AllowedProgramPropertyNamesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgramPropertyNames) IsValid() bool {
	for _, existing := range AllowedProgramPropertyNamesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgramPropertyNames value
func (v ProgramPropertyNames) Ptr() *ProgramPropertyNames {
	return &v
}

type NullableProgramPropertyNames struct {
	value *ProgramPropertyNames
	isSet bool
}

func (v NullableProgramPropertyNames) Get() *ProgramPropertyNames {
	return v.value
}

func (v *NullableProgramPropertyNames) Set(val *ProgramPropertyNames) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramPropertyNames) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramPropertyNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramPropertyNames(val *ProgramPropertyNames) *NullableProgramPropertyNames {
	return &NullableProgramPropertyNames{value: val, isSet: true}
}

func (v NullableProgramPropertyNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramPropertyNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
