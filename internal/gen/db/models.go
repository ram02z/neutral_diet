// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type CfTypes string

const (
	CfTypesTypology    CfTypes = "typology"
	CfTypesSubTypology CfTypes = "sub_typology"
	CfTypesItem        CfTypes = "item"
)

func (e *CfTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CfTypes(s)
	case string:
		*e = CfTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for CfTypes: %T", src)
	}
	return nil
}

type NullCfTypes struct {
	CfTypes CfTypes
	Valid   bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCfTypes) Scan(value interface{}) error {
	if value == nil {
		ns.CfTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CfTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCfTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.CfTypes, nil
}

type AggregateFoodItem struct {
	FoodItemID            int32
	N                     int64
	MedianCarbonFootprint pgtype.Numeric
}

type FoodItem struct {
	ID          int32
	Name        string
	TypologyID  int32
	SuggestedCf CfTypes
}

type FoodItemLog struct {
	ID              int32
	FoodItemID      int32
	Weight          pgtype.Numeric
	CarbonFootprint pgtype.Numeric
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type LifeCycle struct {
	ID              int32
	CarbonFootprint pgtype.Numeric
	FoodItemID      int32
	SourceID        int32
}

type Region struct {
	Name string
}

type Source struct {
	ID         int32
	Reference  string
	Year       int32
	RegionName string
}

type SubTypology struct {
	ID   int32
	Name string
}

type Typology struct {
	ID            int32
	Name          string
	SubTypologyID sql.NullInt32
}

type User struct {
	ID          int32
	FirebaseUid string
	Region      sql.NullString
	CfLimit     pgtype.Numeric
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
