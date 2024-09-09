package mbpb

import (
	"database/sql/driver"
	"fmt"

	"github.com/forhsd/sonic"
	"google.golang.org/protobuf/encoding/protojson"
)

var protoJsonOP = protojson.MarshalOptions{
	EmitUnpopulated: true,
	// UseEnumNumbers:  true,
}

func (t *TaskType) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.(int64)
	if !ok {
		return fmt.Errorf("value is not int64, value: %v", value)
	}

	*t = TaskType(b)

	return nil
}

func (t *TaskType) Value() (driver.Value, error) {

	if t == nil {
		return nil, nil
	}

	return int(*t), nil
}

type Tables []*Table

func (req *Tables) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return sonic.Unmarshal(b, req)
}

func (req *Tables) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Depend*/

type Depends []*Depend

func (req *Depends) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	return sonic.Unmarshal(b, req)
}

func (req *Depends) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Depend*/

// DBDetail
func (req *EnableRequest) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return protojson.Unmarshal(b, req)
}

// Value Valuer
func (req *EnableRequest) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := protoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

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
	return sonic.Unmarshal(b, db)
}

// Value Valuer
func (db *DBDetail) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

// DBDetail
func (s *RunStatus) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.(int64)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	*s = RunStatus(b)
	return nil
}

// Value Valuer
func (s *RunStatus) Value() (driver.Value, error) {

	if s == nil {
		return nil, nil
	}

	return int(*s), nil
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
	return sonic.Unmarshal(b, db)
}

// Value Valuer
func (db *Crontab) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
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
	return sonic.Unmarshal(b, db)
}

// Value Valuer
func (db *Extra) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*TaskflowRequest*/
func (req *TaskflowRequest) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return protojson.Unmarshal(b, req)
}

func (req *TaskflowRequest) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := protoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*TaskflowRequest*/

/*Graph*/
func (req *Graph) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return protojson.Unmarshal(b, req)
}

func (req *Graph) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := protoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Graph*/
