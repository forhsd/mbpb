package mbpb

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
	"google.golang.org/protobuf/encoding/protojson"
)

var ProtoJsonOP = protojson.MarshalOptions{
	EmitUnpopulated: true,
	// UseEnumNumbers:  true,
}

// 跟踪信息
type Trace struct {
	State     int32  `redis:"state,omitempty" json:"state,omitempty"`
	Owner     string `redis:"owner,omitempty" json:"owner,omitempty"`
	Offset    int64  `redis:"offset,omitempty" json:"offset,omitempty"`
	Partition int32  `redis:"partition,omitempty" json:"partition,omitempty"`
}

/*Trace*/
func (x *Trace) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return plug.Unmarshal(b, x)
}

func (x *Trace) Value() (driver.Value, error) {

	if x == nil {
		return nil, nil
	}

	buff, err := plug.Marshal(x)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Trace*/

// FlowDataStat

func (x *FlowDataStat) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.(int64)
	if !ok {
		return fmt.Errorf("value is not int64, value: %v", value)
	}

	*x = FlowDataStat(b)

	return nil
}

func (x *FlowDataStat) Value() (driver.Value, error) {

	if x == nil {
		return nil, nil
	}

	return int64(*x), nil
}

// FlowDataStat

func (t *TaskType) Scan(value any) error {

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

	return int64(*t), nil
}

// 上次的联合关系
type UpUnionRelation struct {
	Source []*UnionRelation `json:"source,omitempty"`
	Target *UnionRelation   `json:"target,omitempty"`
}

/*UpUnionRelation*/
func (req *UpUnionRelation) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	return sonic.Unmarshal(b, req)
}

func (req *UpUnionRelation) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := sonic.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*UpUnionRelation*/

type Tables []*Table

func (req *Tables) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return plug.Unmarshal(b, req)
}

func (req *Tables) Value() (driver.Value, error) {

	if req == nil || len(*req) == 0 {
		return nil, nil
	}

	buff, err := plug.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Depend*/

type Depends struct {
	mu    sync.Mutex
	slice []*Depend
}

func (x *Depends) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	x.slice = []*Depend{}

	return plug.Unmarshal(b, &x.slice)
}

func (x *Depends) Value() (driver.Value, error) {

	if x == nil {
		return nil, nil
	}

	buff, err := plug.Marshal(x.slice)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Depend*/

/*FlowMetadata*/
func (req *FlowMetadata) Scan(value any) error {

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

func (req *FlowMetadata) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := ProtoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*FlowMetadata*/

/*EnableRequest*/
func (req *EnableRequest) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	return protojson.Unmarshal(b, req)
}

func (req *EnableRequest) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := ProtoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*EnableRequest*/

/*WebHookRequest*/
func (req *WebHookRequest) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	return protojson.Unmarshal(b, req)
}

func (req *WebHookRequest) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := ProtoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*WebHookRequest*/

// DBDetail
func (db *DBDetail) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return protojson.Unmarshal(b, db)
}

func (db *DBDetail) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := ProtoJsonOP.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

func (s *RunStatus) Scan(value any) error {

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

func (db *Crontab) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return plug.Unmarshal(b, db)
}

// Value Valuer
func (db *Crontab) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := plug.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

func (db *Extra) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	// return json.Unmarshal(b, &db)
	return plug.Unmarshal(b, db)
}

// Value Valuer
func (db *Extra) Value() (driver.Value, error) {

	if db == nil {
		return nil, nil
	}

	buff, err := plug.Marshal(db)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*TaskflowRequest*/
func (req *FlowEnableRequest) Scan(value any) error {

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

func (req *FlowEnableRequest) Value() (driver.Value, error) {

	if req == nil {
		return nil, nil
	}

	buff, err := ProtoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*TaskflowRequest*/

/*Graph*/
func (req *Graph) Scan(value any) error {

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

	buff, err := ProtoJsonOP.Marshal(req)
	if err != nil {
		return nil, err
	}
	return string(buff), nil
}

/*Graph*/

/*DBType*/
func (x *DBType) Scan(value any) error {

	if value == nil {
		return nil
	}

	b, ok := value.(int64)
	if !ok {
		return fmt.Errorf("value is not int64, value: %v", value)
	}

	*x = DBType(b)

	return nil
}

func (x *DBType) Value() (driver.Value, error) {

	if x == nil {
		return nil, nil
	}

	return int32(*x), nil
}

/*DBType*/

/*Overview*/

// Value 实现了gorm.Type接口，用于获取存储的值
func (o *Overview) Value() (driver.Value, error) {

	bytes, err := plug.Marshal(o)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Scan 实现了sql.Scanner接口，用于从数据库中扫描值
func (o *Overview) Scan(value any) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for JSONBType")
	}

	err := plug.Unmarshal(bytes, o)
	if err != nil {
		return err
	}

	return nil
}

// GormDataType 实现了gorm.Type接口，用于指定Gorm中的数据类型
func (j *Overview) GormDataType() string {
	return "jsonb"
}

/*Overview*/
