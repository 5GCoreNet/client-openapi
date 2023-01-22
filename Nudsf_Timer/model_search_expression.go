/*
Nudsf_Timer

Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudsf_Timer

import (
	"encoding/json"
	"fmt"
)

// SearchExpression - A logical expression element
type SearchExpression struct {
	RecordIdList *RecordIdList
	SearchComparison *SearchComparison
	SearchCondition *SearchCondition
}

// RecordIdListAsSearchExpression is a convenience function that returns RecordIdList wrapped in SearchExpression
func RecordIdListAsSearchExpression(v *RecordIdList) SearchExpression {
	return SearchExpression{
		RecordIdList: v,
	}
}

// SearchComparisonAsSearchExpression is a convenience function that returns SearchComparison wrapped in SearchExpression
func SearchComparisonAsSearchExpression(v *SearchComparison) SearchExpression {
	return SearchExpression{
		SearchComparison: v,
	}
}

// SearchConditionAsSearchExpression is a convenience function that returns SearchCondition wrapped in SearchExpression
func SearchConditionAsSearchExpression(v *SearchCondition) SearchExpression {
	return SearchExpression{
		SearchCondition: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SearchExpression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RecordIdList
	err = newStrictDecoder(data).Decode(&dst.RecordIdList)
	if err == nil {
		jsonRecordIdList, _ := json.Marshal(dst.RecordIdList)
		if string(jsonRecordIdList) == "{}" { // empty struct
			dst.RecordIdList = nil
		} else {
			match++
		}
	} else {
		dst.RecordIdList = nil
	}

	// try to unmarshal data into SearchComparison
	err = newStrictDecoder(data).Decode(&dst.SearchComparison)
	if err == nil {
		jsonSearchComparison, _ := json.Marshal(dst.SearchComparison)
		if string(jsonSearchComparison) == "{}" { // empty struct
			dst.SearchComparison = nil
		} else {
			match++
		}
	} else {
		dst.SearchComparison = nil
	}

	// try to unmarshal data into SearchCondition
	err = newStrictDecoder(data).Decode(&dst.SearchCondition)
	if err == nil {
		jsonSearchCondition, _ := json.Marshal(dst.SearchCondition)
		if string(jsonSearchCondition) == "{}" { // empty struct
			dst.SearchCondition = nil
		} else {
			match++
		}
	} else {
		dst.SearchCondition = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RecordIdList = nil
		dst.SearchComparison = nil
		dst.SearchCondition = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SearchExpression)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SearchExpression)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SearchExpression) MarshalJSON() ([]byte, error) {
	if src.RecordIdList != nil {
		return json.Marshal(&src.RecordIdList)
	}

	if src.SearchComparison != nil {
		return json.Marshal(&src.SearchComparison)
	}

	if src.SearchCondition != nil {
		return json.Marshal(&src.SearchCondition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SearchExpression) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RecordIdList != nil {
		return obj.RecordIdList
	}

	if obj.SearchComparison != nil {
		return obj.SearchComparison
	}

	if obj.SearchCondition != nil {
		return obj.SearchCondition
	}

	// all schemas are nil
	return nil
}

type NullableSearchExpression struct {
	value *SearchExpression
	isSet bool
}

func (v NullableSearchExpression) Get() *SearchExpression {
	return v.value
}

func (v *NullableSearchExpression) Set(val *SearchExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchExpression(val *SearchExpression) *NullableSearchExpression {
	return &NullableSearchExpression{value: val, isSet: true}
}

func (v NullableSearchExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


