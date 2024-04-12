package mbpb

import (
	"fmt"
	"strconv"

	"github.com/cespare/xxhash/v2"
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

func (td *EnableRequest) Hash() string {
	key := fmt.Sprintf("%v.%v", td.EnterpriseID, td.CardId)
	hash := xxhash.Sum64([]byte(key))
	return "etl_" + strconv.FormatUint(hash, 10)
}
