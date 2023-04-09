// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
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
	Valid   bool // Valid is true if CfTypes is not NULL
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
	return string(ns.CfTypes), nil
}

type AggregateFoodItem struct {
	FoodItemID            int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

type CarbonFootprintGoal struct {
	ID                    int32
	UserID                int32
	Description           string
	StartDate             pgtype.Date
	EndDate               pgtype.Date
	StartCarbonFootprint  decimal.Decimal
	TargetCarbonFootprint decimal.Decimal
	Completed             bool
}

type DailyUserAverage struct {
	AverageCarbonFootprint decimal.Decimal
}

type DailyUserAverageByDietaryRequirement struct {
	AverageCarbonFootprint decimal.Decimal
	DietaryRequirement     int32
}

type Device struct {
	ID        int32
	UserID    int32
	FcmToken  string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
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
	Quantity        decimal.Decimal
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	UserID          int32
	LogDate         pgtype.Date
	Unit            int32
	Region          int32
	CarbonFootprint decimal.Decimal
	Meal            int32
}

type LifeCycle struct {
	ID              int32
	CarbonFootprint decimal.Decimal
	FoodItemID      int32
	SourceID        int32
}

type RegionalAggregateFoodItem struct {
	FoodItemID            int32
	Region                int32
	N                     int64
	MedianCarbonFootprint decimal.Decimal
}

type Source struct {
	ID        int32
	Reference string
	Year      int32
	Region    int32
}

type SubTypology struct {
	ID   int32
	Name string
}

type Typology struct {
	ID            int32
	Name          string
	SubTypologyID *int32
}

type User struct {
	ID                 int32
	FirebaseUid        string
	Region             int32
	CfLimit            decimal.Decimal
	CreatedAt          pgtype.Timestamptz
	UpdatedAt          pgtype.Timestamptz
	DietaryRequirement int32
}
