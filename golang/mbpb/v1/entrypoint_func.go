package mbpb

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cespare/xxhash/v2"
)

const Week = 7

var (
	TABLE_PREFIX string
	TABLE_SUFFIX string
)

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

// 获取输出表
func (r *EnableRequest) GetOutputTable() string {

	if TABLE_PREFIX == "" {
		TABLE_PREFIX = os.Getenv("MB_TABLE_PREFIX")
	}

	if TABLE_SUFFIX == "" {
		TABLE_SUFFIX = os.Getenv("MB_TABLE_SUFFIX")
	}

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
