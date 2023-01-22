/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// GeographicArea Geographic area specified by different shape.
type GeographicArea struct {
	EllipsoidArc *EllipsoidArc
	Point *Point
	PointAltitude *PointAltitude
	PointAltitudeUncertainty *PointAltitudeUncertainty
	PointUncertaintyCircle *PointUncertaintyCircle
	PointUncertaintyEllipse *PointUncertaintyEllipse
	Polygon *Polygon
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GeographicArea) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ELLIPSOID_ARC'
	if jsonDict["shape"] == "ELLIPSOID_ARC" {
		// try to unmarshal JSON data into EllipsoidArc
		err = json.Unmarshal(data, &dst.EllipsoidArc);
		if err == nil {
			jsonEllipsoidArc, _ := json.Marshal(dst.EllipsoidArc)
			if string(jsonEllipsoidArc) == "{}" { // empty struct
				dst.EllipsoidArc = nil
			} else {
				return nil // data stored in dst.EllipsoidArc, return on the first match
			}
		} else {
			dst.EllipsoidArc = nil
		}
	}

	// check if the discriminator value is 'EllipsoidArc'
	if jsonDict["shape"] == "EllipsoidArc" {
		// try to unmarshal JSON data into EllipsoidArc
		err = json.Unmarshal(data, &dst.EllipsoidArc);
		if err == nil {
			jsonEllipsoidArc, _ := json.Marshal(dst.EllipsoidArc)
			if string(jsonEllipsoidArc) == "{}" { // empty struct
				dst.EllipsoidArc = nil
			} else {
				return nil // data stored in dst.EllipsoidArc, return on the first match
			}
		} else {
			dst.EllipsoidArc = nil
		}
	}

	// check if the discriminator value is 'LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE'
	if jsonDict["shape"] == "LOCAL_2D_POINT_UNCERTAINTY_ELLIPSE" {
		// try to unmarshal JSON data into Local2dPointUncertaintyEllipse
		err = json.Unmarshal(data, &dst.Local2dPointUncertaintyEllipse);
		if err == nil {
			jsonLocal2dPointUncertaintyEllipse, _ := json.Marshal(dst.Local2dPointUncertaintyEllipse)
			if string(jsonLocal2dPointUncertaintyEllipse) == "{}" { // empty struct
				dst.Local2dPointUncertaintyEllipse = nil
			} else {
				return nil // data stored in dst.Local2dPointUncertaintyEllipse, return on the first match
			}
		} else {
			dst.Local2dPointUncertaintyEllipse = nil
		}
	}

	// check if the discriminator value is 'LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID'
	if jsonDict["shape"] == "LOCAL_3D_POINT_UNCERTAINTY_ELLIPSOID" {
		// try to unmarshal JSON data into Local3dPointUncertaintyEllipsoid
		err = json.Unmarshal(data, &dst.Local3dPointUncertaintyEllipsoid);
		if err == nil {
			jsonLocal3dPointUncertaintyEllipsoid, _ := json.Marshal(dst.Local3dPointUncertaintyEllipsoid)
			if string(jsonLocal3dPointUncertaintyEllipsoid) == "{}" { // empty struct
				dst.Local3dPointUncertaintyEllipsoid = nil
			} else {
				return nil // data stored in dst.Local3dPointUncertaintyEllipsoid, return on the first match
			}
		} else {
			dst.Local3dPointUncertaintyEllipsoid = nil
		}
	}

	// check if the discriminator value is 'POINT'
	if jsonDict["shape"] == "POINT" {
		// try to unmarshal JSON data into Point
		err = json.Unmarshal(data, &dst.Point);
		if err == nil {
			jsonPoint, _ := json.Marshal(dst.Point)
			if string(jsonPoint) == "{}" { // empty struct
				dst.Point = nil
			} else {
				return nil // data stored in dst.Point, return on the first match
			}
		} else {
			dst.Point = nil
		}
	}

	// check if the discriminator value is 'POINT_ALTITUDE'
	if jsonDict["shape"] == "POINT_ALTITUDE" {
		// try to unmarshal JSON data into PointAltitude
		err = json.Unmarshal(data, &dst.PointAltitude);
		if err == nil {
			jsonPointAltitude, _ := json.Marshal(dst.PointAltitude)
			if string(jsonPointAltitude) == "{}" { // empty struct
				dst.PointAltitude = nil
			} else {
				return nil // data stored in dst.PointAltitude, return on the first match
			}
		} else {
			dst.PointAltitude = nil
		}
	}

	// check if the discriminator value is 'POINT_ALTITUDE_UNCERTAINTY'
	if jsonDict["shape"] == "POINT_ALTITUDE_UNCERTAINTY" {
		// try to unmarshal JSON data into PointAltitudeUncertainty
		err = json.Unmarshal(data, &dst.PointAltitudeUncertainty);
		if err == nil {
			jsonPointAltitudeUncertainty, _ := json.Marshal(dst.PointAltitudeUncertainty)
			if string(jsonPointAltitudeUncertainty) == "{}" { // empty struct
				dst.PointAltitudeUncertainty = nil
			} else {
				return nil // data stored in dst.PointAltitudeUncertainty, return on the first match
			}
		} else {
			dst.PointAltitudeUncertainty = nil
		}
	}

	// check if the discriminator value is 'POINT_UNCERTAINTY_CIRCLE'
	if jsonDict["shape"] == "POINT_UNCERTAINTY_CIRCLE" {
		// try to unmarshal JSON data into PointUncertaintyCircle
		err = json.Unmarshal(data, &dst.PointUncertaintyCircle);
		if err == nil {
			jsonPointUncertaintyCircle, _ := json.Marshal(dst.PointUncertaintyCircle)
			if string(jsonPointUncertaintyCircle) == "{}" { // empty struct
				dst.PointUncertaintyCircle = nil
			} else {
				return nil // data stored in dst.PointUncertaintyCircle, return on the first match
			}
		} else {
			dst.PointUncertaintyCircle = nil
		}
	}

	// check if the discriminator value is 'POINT_UNCERTAINTY_ELLIPSE'
	if jsonDict["shape"] == "POINT_UNCERTAINTY_ELLIPSE" {
		// try to unmarshal JSON data into PointUncertaintyEllipse
		err = json.Unmarshal(data, &dst.PointUncertaintyEllipse);
		if err == nil {
			jsonPointUncertaintyEllipse, _ := json.Marshal(dst.PointUncertaintyEllipse)
			if string(jsonPointUncertaintyEllipse) == "{}" { // empty struct
				dst.PointUncertaintyEllipse = nil
			} else {
				return nil // data stored in dst.PointUncertaintyEllipse, return on the first match
			}
		} else {
			dst.PointUncertaintyEllipse = nil
		}
	}

	// check if the discriminator value is 'POLYGON'
	if jsonDict["shape"] == "POLYGON" {
		// try to unmarshal JSON data into Polygon
		err = json.Unmarshal(data, &dst.Polygon);
		if err == nil {
			jsonPolygon, _ := json.Marshal(dst.Polygon)
			if string(jsonPolygon) == "{}" { // empty struct
				dst.Polygon = nil
			} else {
				return nil // data stored in dst.Polygon, return on the first match
			}
		} else {
			dst.Polygon = nil
		}
	}

	// check if the discriminator value is 'Point'
	if jsonDict["shape"] == "Point" {
		// try to unmarshal JSON data into Point
		err = json.Unmarshal(data, &dst.Point);
		if err == nil {
			jsonPoint, _ := json.Marshal(dst.Point)
			if string(jsonPoint) == "{}" { // empty struct
				dst.Point = nil
			} else {
				return nil // data stored in dst.Point, return on the first match
			}
		} else {
			dst.Point = nil
		}
	}

	// check if the discriminator value is 'PointAltitude'
	if jsonDict["shape"] == "PointAltitude" {
		// try to unmarshal JSON data into PointAltitude
		err = json.Unmarshal(data, &dst.PointAltitude);
		if err == nil {
			jsonPointAltitude, _ := json.Marshal(dst.PointAltitude)
			if string(jsonPointAltitude) == "{}" { // empty struct
				dst.PointAltitude = nil
			} else {
				return nil // data stored in dst.PointAltitude, return on the first match
			}
		} else {
			dst.PointAltitude = nil
		}
	}

	// check if the discriminator value is 'PointAltitudeUncertainty'
	if jsonDict["shape"] == "PointAltitudeUncertainty" {
		// try to unmarshal JSON data into PointAltitudeUncertainty
		err = json.Unmarshal(data, &dst.PointAltitudeUncertainty);
		if err == nil {
			jsonPointAltitudeUncertainty, _ := json.Marshal(dst.PointAltitudeUncertainty)
			if string(jsonPointAltitudeUncertainty) == "{}" { // empty struct
				dst.PointAltitudeUncertainty = nil
			} else {
				return nil // data stored in dst.PointAltitudeUncertainty, return on the first match
			}
		} else {
			dst.PointAltitudeUncertainty = nil
		}
	}

	// check if the discriminator value is 'PointUncertaintyCircle'
	if jsonDict["shape"] == "PointUncertaintyCircle" {
		// try to unmarshal JSON data into PointUncertaintyCircle
		err = json.Unmarshal(data, &dst.PointUncertaintyCircle);
		if err == nil {
			jsonPointUncertaintyCircle, _ := json.Marshal(dst.PointUncertaintyCircle)
			if string(jsonPointUncertaintyCircle) == "{}" { // empty struct
				dst.PointUncertaintyCircle = nil
			} else {
				return nil // data stored in dst.PointUncertaintyCircle, return on the first match
			}
		} else {
			dst.PointUncertaintyCircle = nil
		}
	}

	// check if the discriminator value is 'PointUncertaintyEllipse'
	if jsonDict["shape"] == "PointUncertaintyEllipse" {
		// try to unmarshal JSON data into PointUncertaintyEllipse
		err = json.Unmarshal(data, &dst.PointUncertaintyEllipse);
		if err == nil {
			jsonPointUncertaintyEllipse, _ := json.Marshal(dst.PointUncertaintyEllipse)
			if string(jsonPointUncertaintyEllipse) == "{}" { // empty struct
				dst.PointUncertaintyEllipse = nil
			} else {
				return nil // data stored in dst.PointUncertaintyEllipse, return on the first match
			}
		} else {
			dst.PointUncertaintyEllipse = nil
		}
	}

	// check if the discriminator value is 'Polygon'
	if jsonDict["shape"] == "Polygon" {
		// try to unmarshal JSON data into Polygon
		err = json.Unmarshal(data, &dst.Polygon);
		if err == nil {
			jsonPolygon, _ := json.Marshal(dst.Polygon)
			if string(jsonPolygon) == "{}" { // empty struct
				dst.Polygon = nil
			} else {
				return nil // data stored in dst.Polygon, return on the first match
			}
		} else {
			dst.Polygon = nil
		}
	}

	// try to unmarshal JSON data into EllipsoidArc
	err = json.Unmarshal(data, &dst.EllipsoidArc);
	if err == nil {
		jsonEllipsoidArc, _ := json.Marshal(dst.EllipsoidArc)
		if string(jsonEllipsoidArc) == "{}" { // empty struct
			dst.EllipsoidArc = nil
		} else {
			return nil // data stored in dst.EllipsoidArc, return on the first match
		}
	} else {
		dst.EllipsoidArc = nil
	}

	// try to unmarshal JSON data into Point
	err = json.Unmarshal(data, &dst.Point);
	if err == nil {
		jsonPoint, _ := json.Marshal(dst.Point)
		if string(jsonPoint) == "{}" { // empty struct
			dst.Point = nil
		} else {
			return nil // data stored in dst.Point, return on the first match
		}
	} else {
		dst.Point = nil
	}

	// try to unmarshal JSON data into PointAltitude
	err = json.Unmarshal(data, &dst.PointAltitude);
	if err == nil {
		jsonPointAltitude, _ := json.Marshal(dst.PointAltitude)
		if string(jsonPointAltitude) == "{}" { // empty struct
			dst.PointAltitude = nil
		} else {
			return nil // data stored in dst.PointAltitude, return on the first match
		}
	} else {
		dst.PointAltitude = nil
	}

	// try to unmarshal JSON data into PointAltitudeUncertainty
	err = json.Unmarshal(data, &dst.PointAltitudeUncertainty);
	if err == nil {
		jsonPointAltitudeUncertainty, _ := json.Marshal(dst.PointAltitudeUncertainty)
		if string(jsonPointAltitudeUncertainty) == "{}" { // empty struct
			dst.PointAltitudeUncertainty = nil
		} else {
			return nil // data stored in dst.PointAltitudeUncertainty, return on the first match
		}
	} else {
		dst.PointAltitudeUncertainty = nil
	}

	// try to unmarshal JSON data into PointUncertaintyCircle
	err = json.Unmarshal(data, &dst.PointUncertaintyCircle);
	if err == nil {
		jsonPointUncertaintyCircle, _ := json.Marshal(dst.PointUncertaintyCircle)
		if string(jsonPointUncertaintyCircle) == "{}" { // empty struct
			dst.PointUncertaintyCircle = nil
		} else {
			return nil // data stored in dst.PointUncertaintyCircle, return on the first match
		}
	} else {
		dst.PointUncertaintyCircle = nil
	}

	// try to unmarshal JSON data into PointUncertaintyEllipse
	err = json.Unmarshal(data, &dst.PointUncertaintyEllipse);
	if err == nil {
		jsonPointUncertaintyEllipse, _ := json.Marshal(dst.PointUncertaintyEllipse)
		if string(jsonPointUncertaintyEllipse) == "{}" { // empty struct
			dst.PointUncertaintyEllipse = nil
		} else {
			return nil // data stored in dst.PointUncertaintyEllipse, return on the first match
		}
	} else {
		dst.PointUncertaintyEllipse = nil
	}

	// try to unmarshal JSON data into Polygon
	err = json.Unmarshal(data, &dst.Polygon);
	if err == nil {
		jsonPolygon, _ := json.Marshal(dst.Polygon)
		if string(jsonPolygon) == "{}" { // empty struct
			dst.Polygon = nil
		} else {
			return nil // data stored in dst.Polygon, return on the first match
		}
	} else {
		dst.Polygon = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GeographicArea)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GeographicArea) MarshalJSON() ([]byte, error) {
	if src.EllipsoidArc != nil {
		return json.Marshal(&src.EllipsoidArc)
	}

	if src.Point != nil {
		return json.Marshal(&src.Point)
	}

	if src.PointAltitude != nil {
		return json.Marshal(&src.PointAltitude)
	}

	if src.PointAltitudeUncertainty != nil {
		return json.Marshal(&src.PointAltitudeUncertainty)
	}

	if src.PointUncertaintyCircle != nil {
		return json.Marshal(&src.PointUncertaintyCircle)
	}

	if src.PointUncertaintyEllipse != nil {
		return json.Marshal(&src.PointUncertaintyEllipse)
	}

	if src.Polygon != nil {
		return json.Marshal(&src.Polygon)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGeographicArea struct {
	value *GeographicArea
	isSet bool
}

func (v NullableGeographicArea) Get() *GeographicArea {
	return v.value
}

func (v *NullableGeographicArea) Set(val *GeographicArea) {
	v.value = val
	v.isSet = true
}

func (v NullableGeographicArea) IsSet() bool {
	return v.isSet
}

func (v *NullableGeographicArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeographicArea(val *GeographicArea) *NullableGeographicArea {
	return &NullableGeographicArea{value: val, isSet: true}
}

func (v NullableGeographicArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeographicArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


