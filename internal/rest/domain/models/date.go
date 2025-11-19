package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Date represents a date in YYYY-MM-DD format
type Date struct {
	time.Time
}

// DateLayout is the expected date format
const DateLayout = "2006-01-02"

// UnmarshalJSON parses dates from JSON in YYYY-MM-DD format
func (d *Date) UnmarshalJSON(data []byte) error {
	// Remove quotes
	str := strings.Trim(string(data), `"`)

	// Handle empty/null values
	if str == "" || str == "null" {
		d.Time = time.Time{}
		return nil
	}

	// Try to parse as date only (YYYY-MM-DD)
	parsedTime, err := time.Parse(DateLayout, str)
	if err == nil {
		d.Time = parsedTime
		return nil
	}

	// Fallback: try to parse as RFC3339 if full datetime is provided
	parsedTime, err = time.Parse(time.RFC3339, str)
	if err != nil {
		return fmt.Errorf("invalid date format: expected YYYY-MM-DD or RFC3339, got %s", str)
	}

	d.Time = parsedTime
	return nil
}

// MarshalJSON formats dates to JSON in YYYY-MM-DD format
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}

	// Format as YYYY-MM-DD
	str := fmt.Sprintf(`"%s"`, d.Time.Format(DateLayout))
	return []byte(str), nil
}

// Scan implements the sql.Scanner interface
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	case nil:
		d.Time = time.Time{}
		return nil
	default:
		return fmt.Errorf("cannot scan %T into Date", value)
	}
}

// Value implements the driver.Valuer interface
func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

// String returns the date as a string in YYYY-MM-DD format
func (d Date) String() string {
	if d.Time.IsZero() {
		return ""
	}
	return d.Time.Format(DateLayout)
}

// NewDate creates a new Date from a time.Time
func NewDate(t time.Time) Date {
	return Date{Time: t}
}

// ParseDate parses a date string in YYYY-MM-DD format
func ParseDate(s string) (Date, error) {
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: t}, nil
}
