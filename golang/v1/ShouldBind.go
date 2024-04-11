package mbpb

import (
	"database/sql/driver"
	"fmt"

	"github.com/bytedance/sonic"
)

// DBDetail
func (db *DBDetail) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return sonic.Unmarshal(b, &db)
}

// Value Valuer
func (db *DBDetail) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	// return json.Marshal(db)
	return sonic.Marshal(db)
}

func (db *Crontab) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return sonic.Unmarshal(b, &db)
}

// Value Valuer
func (db *Crontab) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	// return json.Marshal(db)
	return sonic.Marshal(db)
}

func (db *Table) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return sonic.Unmarshal(b, &db)
}

// Value Valuer
func (db *Table) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	// return json.Marshal(db)
	return sonic.Marshal(db)
}

func (db *Extra) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return sonic.Unmarshal(b, &db)
}

// Value Valuer
func (db *Extra) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	// return json.Marshal(db)
	return sonic.Marshal(db)
}
