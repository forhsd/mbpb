package mbpb

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cespare/xxhash/v2"
)

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

func (r *EnableRequest) Hash() string {
	key := fmt.Sprintf("%v.%v", r.GetEnterpriseID(), r.GetCardId())
	hash := xxhash.Sum64([]byte(key))
	return "etl_" + strconv.FormatUint(hash, 10)
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
