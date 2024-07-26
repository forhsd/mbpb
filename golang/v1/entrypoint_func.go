package mbpb

import (
	context "context"
	"database/sql/driver"
	"errors"
	"fmt"
	"mbetl/ecode"
	"strconv"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cespare/xxhash/v2"
	flow "github.com/forhsd/go-workflow"
	"github.com/google/uuid"
)

const (
	Week = 7
	CST8 = "Asia/Shanghai"
)

const (
	TABLE_PREFIX string = `__xfjwi__table`
	TABLE_SUFFIX string = `0201`
)

func Hash(item ...interface{}) uint64 {

	var arr []string
	for _, val := range item {
		arr = append(arr, fmt.Sprintf("%v", val))
	}

	key := strings.Join(arr, ".")
	return xxhash.Sum64([]byte(key))
}

func HashString(item ...interface{}) string {
	return strconv.FormatUint(Hash(item...), 10)
}

// 使用按位或设置状态
func (s *RunStatus) Set(status RunStatus) {
	// *s |= status
	*s = status
}

// 返回当前状态
func (s RunStatus) Get() RunStatus {
	return s
}

// 使用按位与测试状态是否存在
func (s RunStatus) Has(status RunStatus) bool {
	return s&status != 0
}

func (r *Request) GetUniqueId() *string {
	key := fmt.Sprintf("%v.%v", r.GetEnterpriseID(), r.GetCardId())
	hash := xxhash.Sum64([]byte(key))
	hashsrt := strconv.FormatUint(hash, 10)
	return &hashsrt
}

func (r *EnableRequest) Hash() string {
	key := fmt.Sprintf("%v.%v", r.GetEnterpriseID(), r.GetCardId())
	hash := xxhash.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

func (r *EnableRequest) SequenceID(runRime *time.Time) string {
	if runRime == nil {
		runRime = &time.Time{}
	}
	key := fmt.Sprintf("%v.%v.%v", r.GetEnterpriseID(), r.GetCardId(), runRime.UnixNano())
	hash := xxhash.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

func (r *Request) SequenceID(runRime *time.Time) string {
	if runRime == nil {
		runRime = &time.Time{}
	}
	key := fmt.Sprintf("%v.%v.%v", r.GetEnterpriseID(), r.GetCardId(), runRime.UnixNano())
	hash := xxhash.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

// 获取输出Schema
func (r *EnableRequest) GetOutputSchema() string {

	if r.GetEnterpriseID() == "" {
		return "result_set_defaultenterprise_db"
	}

	return fmt.Sprintf("result_set_%v_db", r.GetEnterpriseID())
}

// 获取时区
func (r *EnableRequest) GetZone() string {

	zone := r.GetCrontab().GetLifeCycle().GetZone()
	if zone == "" {
		zone = CST8
	}
	return zone
}

func (r *EnableRequest) GetLocation() (*time.Location, error) {
	return time.LoadLocation(r.GetZone())
}

// 获取输出表
func (r *EnableRequest) GetOutputTable() string {

	return strings.Join(
		[]string{
			TABLE_PREFIX,
			// strconv.Itoa(table),
			strconv.FormatInt(r.GetCardId(), 10),
			TABLE_SUFFIX,
		},
		"__",
	)
}

// 验证
func (r *Request) Validate() error {

	if r.GetEnterpriseID() == "" {
		return errors.New("EnterpriseID不能为空")
	}

	if r.GetCardId() == 0 {
		return errors.New("CardId不能为空")
	}

	return nil
}

// 验证
func (r *EnableRequest) Validate() error {

	check := []string{}

	if r.GetCardId() == 0 {
		check = append(check, "CardId")
	}

	if r.GetSqlScript() == "" {
		check = append(check, "SqlScript")
	}

	if r.GetDBInfo() == nil {
		check = append(check, "DBInfo")
	}

	if r.GetDBInfo().GetHost() == "" {
		check = append(check, "Host")
	}

	if r.GetDBInfo().GetPort() == 0 {
		check = append(check, "Port")
	}

	if r.GetDBInfo().GetUser() == "" {
		check = append(check, "User")
	}

	if r.GetDBInfo().GetPwd() == "" {
		check = append(check, "PassWord")
	}

	if r.GetDBInfo().GetDBName() == "" {
		check = append(check, "DBName")
	}

	if r.GetDBInfo().GetDBType() == "" {
		check = append(check, "DBType")
	}

	if len(check) != 0 {
		return errors.New(strings.Join(check, ", ") + " 不能为空")
	}

	return nil

}

func (x *Overview) GetErrorEnding(ctx context.Context, err error) {

	if x == nil {
		x = &Overview{}
	}

	if x.EndTime == "" {
		x.EndTime = time.Now().Local().Format(time.DateTime)
	}

	if x.Detail == nil {
		x.Detail = &Detail{}
	}
	if x.Detail.Error == nil {
		x.Detail.Error = &Error{}
	}

	errs := ecode.FromContextError(err)

	// 用户取消或失败
	if ctx.Err() == context.Canceled {
		x.RunStatus = RunStatus_Cancel

		if x.Detail.Error.Code == 0 {
			x.Detail.Error.Code = int32(ecode.UserCancellation)
		}
		if x.Detail.Error.Msg == "" {
			x.Detail.Error.Msg = ctx.Err().Error()
		}
	} else {
		x.RunStatus = RunStatus_Fail

		if x.Detail.Error.Code == 0 {
			x.Detail.Error.Code = int32(errs.Code())
		}
		if x.Detail.Error.Msg != "" {
			return
		}
		x.Detail.Error.Msg = errs.Message()
	}
}

// Value 实现了gorm.Type接口，用于获取存储的值
func (o *Overview) Value() (driver.Value, error) {

	bytes, err := sonic.Marshal(o)
	if err != nil {
		return nil, err
	}

	return string(bytes), nil
}

// Scan 实现了sql.Scanner接口，用于从数据库中扫描值
func (o *Overview) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for JSONBType")
	}

	err := sonic.Unmarshal(bytes, o)
	if err != nil {
		return err
	}

	return nil
}

// GormDataType 实现了gorm.Type接口，用于指定Gorm中的数据类型
func (j *Overview) GormDataType() string {
	return "jsonb"
}

type Node struct {
	ID       uint64  `json:"id"`
	Children []*Node `json:"children,omitempty"`
	Edges    []*Edge `json:"edges,omitempty"`
	X        int     `json:"x,omitempty"`
	Y        int     `json:"y,omitempty"`
	Width    int     `json:"width,omitempty"`
	Height   int     `json:"height,omitempty"`
	Labels   []Label `json:"labels,omitempty"`
}

type Label struct {
	Text string `json:"text,omitempty"`
}

// 血亲
func Relatives(work *flow.Workflow) *Node {

	root := &Node{ID: 0}
	nodes := map[flow.Steper]*Node{
		work: root,
	}
	getNode := func(s flow.Steper) *Node {
		node, ok := nodes[s]
		if !ok {
			// node = &Node{ID: uuid.NewString()}
			t := flow.String(s)
			id, _ := strconv.ParseUint(t, 10, 64)
			node = &Node{ID: id}
			nodes[s] = node
		}
		return node
	}
	flow.Traverse(work, func(s flow.Steper, walked []flow.Steper) flow.TraverseDecision {
		if w, ok := s.(interface {
			Unwrap() []flow.Steper
			UpstreamOf(flow.Steper) map[flow.Steper]flow.StepResult
		}); ok {
			for _, r := range w.Unwrap() {
				n := getNode(r)
				n.Labels = append(n.Labels, Label{flow.String(r)})
				parent := s
				for i := len(walked) - 1; i >= 0; i-- {
					if _, ok := walked[i].(interface{ Unwrap() []flow.Steper }); ok {
						if i < len(walked)-1 {
							parent = walked[i+1]
							break
						}
					}
				}
				getNode(parent).Children = append(getNode(parent).Children, n)

				for up := range w.UpstreamOf(r) {
					eid := uuid.NewString()
					// eid := flow.String(r)
					getNode(parent).Edges = append(getNode(parent).Edges, &Edge{
						Id:     eid,
						Source: getNode(up).ID,
						Target: getNode(r).ID,
					})
				}
			}
		}
		return flow.TraverseContinue
	})

	return root
}
