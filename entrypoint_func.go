package mbpb

import (
	context "context"
	"database/sql"
	"errors"
	"fmt"
	"mbetl/collector/asynq/common"
	"mbetl/runtime/config"
	"mbetl/runtime/ecode"
	"mbetl/runtime/util"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	flow "github.com/Azure/go-workflow"
	"github.com/Masterminds/semver/v3"
	"github.com/apache/arrow-adbc/go/adbc"
	drvfs "github.com/apache/arrow-adbc/go/adbc/driver/flightsql"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/bytedance/sonic"
	"github.com/cespare/xxhash/v2"
	"github.com/forhsd/logger"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

const (
	Week = 7
	CST8 = "Asia/Shanghai"
)

const (
	TABLE_PREFIX string = `__xfjwi__table__`
	TABLE_SUFFIX string = `__0201`
)

const (
	DorisBatchSize = 4064
)

var (
	plug Plugin
)

var (
	ErrEntrypointIdCannotBeEmpty      = errors.New("enterpriseId不能为空")
	ErrRequestIDCannotBeEmpty         = errors.New("请求ID不能为空")
	ErrRequestSequenceIdCannotBeEmpty = errors.New("序列ID不能为空")
	ErrGraphDataException             = errors.New("Graph数据异常,请勿使用脏数据")
)

const (
	Err_workflow_IsRunning      = "the workflow[%d] is running, please try again later"
	Err_workflow_IsRunning_Skip = "工作流运行中,本次不运行"
	Err_workflow_Node_IsRunning = "工作流节点[%d]运行中,请稍后再试" // the workflow node[%d] is running, please try again later
)

func RegisterPlugin(tool Plugin) {
	plug = tool
}

// 插件接口
type Plugin interface {
	Sum64(b []byte) uint64                     // xxhash.Sum64
	Marshal(val any) ([]byte, error)           // sonic.Marshal
	Unmarshal(buf []byte, val any) error       // sonic.Unmarshal
	FromContextError(err error) *status.Status // ecode.FromContextError
}

func init() {
	plug = &tool{}
}

type tool struct{}

func (*tool) Sum64(b []byte) uint64 {
	return xxhash.Sum64(b)
}

func (*tool) Marshal(val any) ([]byte, error) {
	return sonic.Marshal(val)
}

func (*tool) Unmarshal(buf []byte, val any) error {
	return sonic.Unmarshal(buf, val)
}

func (*tool) FromContextError(err error) *status.Status {
	return ecode.FromContextError(err)
}

type RequestImpi interface {
	EnableRequest | FlowEnableRequest | WebHookRequest
}

func Hash(item ...any) uint64 {

	var arr []string
	for _, val := range item {
		arr = append(arr, fmt.Sprintf("%v", val))
	}

	key := strings.Join(arr, ".")
	return plug.Sum64([]byte(key))
}

func HashString(item ...any) string {
	return strconv.FormatUint(Hash(item...), 10)
}

func (o *Over) IsWorkflowNode(s *Table) bool {

	logger.Info("children", o.GetChildren())

	for _, val := range o.GetChildren() {

		if val.GetBaseId() == s.GetSource().GetBaseId() {
			// logger.Warn("命中来源%+v %+v", val, s.GetSource())
			return true
		}

		// if val.GetBaseId() == s.GetTarget().GetBaseId() {
		// 	logger.Warn("命中目标%+v %+v", val, s.GetTarget())
		// 	return true
		// }
	}

	return false
}

func (x *Over) GetTopic() string {

	switch x.GetRunType() {

	case RunType_Cycle:
		return common.Topic_Workflow_Cycle
	case RunType_Spark:
		return common.Topic_Workflow_Spark
	default:
		return "unknownErr"
	}
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
	hash := plug.Sum64([]byte(key))
	hashsrt := strconv.FormatUint(hash, 10)
	return &hashsrt
}

func (r *EnableRequest) Hash() string {
	key := fmt.Sprintf("%v.%v", r.GetEnterpriseId(), r.GetCardId())
	hash := plug.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

func (r *EnableRequest) SequenceID(runRime *time.Time) string {
	if runRime == nil {
		runRime = &time.Time{}
	}
	key := fmt.Sprintf("%v.%v.%v", r.GetEnterpriseId(), r.GetCardId(), runRime.UnixNano())
	hash := plug.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

type MbSign bool

func (x *MbSign) Visited() {
	*x = true
}

func (x *MbSign) Get() bool {
	if x == nil {
		return false
	}
	return bool(*x)
}

type UKID struct {
	IDs  []string
	Reqs map[string]*MbSign
	mu   sync.Mutex
}

func NewUKID() *UKID {
	return &UKID{
		Reqs: make(map[string]*MbSign),
	}
}

func (w *UKID) Add(id string) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.Reqs[id]; !exists {
		w.IDs = append(w.IDs, id)
		w.Reqs[id] = func() *MbSign {
			sign := false
			return (*MbSign)(&sign)
		}()
	}

}

func (w *UKID) Get(id string) (*MbSign, bool) {
	sign, found := w.Reqs[id]
	return sign, found
}

func (w *UKID) AddPointer(id string, sign *MbSign) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if _, exists := w.Reqs[id]; !exists {
		w.IDs = append(w.IDs, id)
		w.Reqs[id] = sign
	}

}

func (w *UKID) Visited(id string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// w.Reqs[id] = true
	sign, found := w.Reqs[id]
	if found {
		sign.Visited()
	}
}

func (w *UKID) GetVisited(id string) bool {

	visited := w.Reqs[id]
	return visited.Get()
}

func (w *UKID) Next(ids []string) *UKID {

	uk := NewUKID()

	for _, id := range ids {
		if !w.GetVisited(id) {
			// uk.Add(id)
			sign, found := w.Reqs[id]
			if found {
				uk.AddPointer(id, sign)
			}
		}
	}

	return uk
}

func (w *UKID) Empty() bool {

	var lenght int
	for _, sign := range w.Reqs {

		if !sign.Get() {
			lenght++
		}
	}

	return lenght == 0
}

// 从请求Tables获取来源ID
//
//	参数
//	uk := UKID{
//		Reqs: make(map[string]bool),
//	}
func (x *EnableRequest) GetSourceIDWithRequestTables(uk *UKID, opts ...func(table *Table)) (result []string) {

	for _, table := range x.GetTables() {

		if table.GetSource().GetType() == SourceType_MbEtl {

			id := table.GetSource().GetId()

			result = append(result, id)

			uk.Add(id)

			for _, op := range opts {
				op(table)
			}
		}

	}

	return util.Unique(result)

}

func (x *EnableRequest) GetSourceFQNWithRequestTables() (tables []string) {
	for _, table := range x.GetTables() {
		tables = append(tables, fmt.Sprintf(`"%s"."%s"`, table.GetSchema(), table.GetTable()))
	}
	return
}

func FlowSequenceID(id int32, runTime *time.Time) string {

	if runTime == nil {
		runTime = &time.Time{}
	}

	key := fmt.Sprintf("workflow.%v.%v", id, runTime.UnixNano())
	hash := plug.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

func (r *FlowRequest) SequenceID(runTime *time.Time) string {
	return FlowSequenceID(r.GetId(), runTime)
}

func (r *Request) SequenceID(runRime *time.Time) string {
	if runRime == nil {
		runRime = &time.Time{}
	}
	key := fmt.Sprintf("%v.%v.%v", r.GetEnterpriseID(), r.GetCardId(), runRime.UnixNano())
	hash := plug.Sum64([]byte(key))
	return strconv.FormatUint(hash, 10)
}

// 获取输出Schema
func (r *EnableRequest) GetOutputSchema() string {

	if r.GetEnterpriseId() == "" {
		return "result_set_defaultenterprise_db"
	}

	return fmt.Sprintf("result_set_%v_db", r.GetEnterpriseId())
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
		"",
	)
}

func (x *EnableRequest) GetRequestHash() string {

	expend := func(s *DBDetail) []any {

		resp := make([]any, 0, 8)

		resp = append(resp, "dbType:", s.GetDbType())

		switch s.GetDbType() {
		case DBType_DORIS:
			db := s.GetDoris()
			resp = append(resp, "host:", db.GetHost())
			resp = append(resp, "port:", db.GetQueryPort())
			resp = append(resp, "dbName:", db.GetDbName())
		case DBType_MYSQL:
			db := s.GetMysql()
			resp = append(resp, "host:", db.GetHost())
			resp = append(resp, "port:", db.GetPort())
			resp = append(resp, "dbName:", db.GetDbName())
		case DBType_POSTGRES:
			db := s.GetPostgres()
			resp = append(resp, "host:", db.GetHost())
			resp = append(resp, "port:", db.GetPort())
			resp = append(resp, "dbName:", db.GetDbName())
		}

		return resp
	}

	return HashString(
		expend(x.GetSourceDb()),
		expend(x.GetTargetDb()),
		x.GetSqlScript(),
	)
}

func (db *DatabaseRequest) Validate() error {

	if err := db.GetDetails().Validate(); err != nil {
		return err
	}

	if err := db.GetDetails().TryConnect(); err != nil {
		return err
	}

	if db.GetDisplayName() == "" {
		return errors.New("显示名字不能为空")
	}

	return nil
}

func (x *ModelRequest) Validate() error {

	var check []string

	if x.GetDisplayName() == "" {
		check = append(check, "显示名")
	}

	if x.GetScript() == "" {
		check = append(check, "sql语句")
	}

	if x.GetTargetDbId() < 1 {
		check = append(check, "目标数据库ID")
	}

	if len(check) > 0 {
		return errors.New(strings.Join(check, ",") + "不能为空")
	}

	return nil
}

// 验证
func (r *Request) Validate() error {

	if r.GetEnterpriseID() == "" {
		return ErrEntrypointIdCannotBeEmpty
	}

	if r.GetCardId() == 0 {
		return errors.New("CardId不能为空")
	}

	return nil
}

func (req *FlowEnableRequest) Validate() error {

	// 为空则初始化
	if req.Crontab == nil {
		req.Crontab = &Crontab{}
		logger.Info("init Crontab")
	}

	var errarg []string

	if req.FlowID == 0 {
		errarg = append(errarg, "flowID")
	}

	if req.UserId == 0 {
		errarg = append(errarg, "userId")
	}

	// if req.EmitCardId != nil && *req.EmitCardId != 0 {
	// 	errarg = append(errarg, "emitCardId")
	// }

	if req.EnterpriseID == "" {
		errarg = append(errarg, "enterpriseID")
	}

	if len(errarg) > 0 {
		return fmt.Errorf("%v不能为空", strings.Join(errarg, " ,"))
	}

	// // req.FlowMap == nil ||
	// if len(req.FlowMap) == 0 {
	// 	return errors.New("flowMAP不能为空")
	// }

	return nil
}

func (req *FlowRequest) Validate() error {

	if req.GetId() < 1 {
		return ecode.Errorfr(ecode.InvalidParameter, ErrRequestIDCannotBeEmpty)
	}
	return nil
}

func (req *FlowRequest) RunValidate() error {

	if req.GetSequenceId() == "" {
		return ecode.Errorfr(ecode.InvalidParameter, ErrRequestSequenceIdCannotBeEmpty)
	}

	return req.Validate()
}

func (r *EnableRequest) DefaultPostgresDwd() {

	if r.GetTargetDb() == nil {

		connectParams := map[string]string{}

		dwd := config.Read().DataSet
		if dwd.Custom != "" {

			params := strings.Split(dwd.Custom, "&")
			for _, param := range params {
				kv := strings.Split(param, "=")
				if len(kv) == 2 {
					connectParams[kv[0]] = kv[1]
				}
			}
		}

		if r.GetTargetDb() == nil {
			r.TargetDb = &DBDetail{
				DbType: DBType_POSTGRES,
				Payload: &DBDetail_Postgres{
					Postgres: &GenericDB{
						// DbType: dwd.Driver,
						Host:          dwd.Host,
						Port:          int32(dwd.Port),
						User:          dwd.User,
						Pwd:           dwd.Pass,
						DbName:        dwd.DBName,
						ConnectParams: connectParams,
					},
				},
			}
		}

	}

}

// 是否为同类型数据库同步
func (r *EnableRequest) IsHomogeneousSync() bool {
	return r.GetSourceDb().GetDbType() == r.GetTargetDb().GetDbType()
}

func dbTypeVerify(typ DBType) error {
	switch typ {
	case DBType_POSTGRES, DBType_DORIS:
	default:
		return errors.New("未兼容的数据库")
	}
	return nil
}

// 验证
func (r *EnableRequest) Validate() error {

	var errs []error

	// 验证源数据库配置
	if err := r.GetSourceDb().Validate(); err != nil {
		errs = append(errs, fmt.Errorf("源库: %w", err))
	}

	// 验证db类型
	if err := dbTypeVerify(r.GetSourceDb().GetDbType()); err != nil {
		errs = append(errs, err)
	}

	// 验证目标数据库配置
	if err := r.GetTargetDb().Validate(); err != nil {
		errs = append(errs, fmt.Errorf("目标: %w", err))
	}

	// 验证db类型
	if err := dbTypeVerify(r.GetTargetDb().GetDbType()); err != nil {
		errs = append(errs, err)
	}

	check := []string{}

	if r.GetCardId() == 0 {
		check = append(check, "CardId")
	}

	if r.GetEnterpriseId() == "" {
		check = append(check, "EnterpriseID")
	}

	if r.GetSqlScript() == "" {
		check = append(check, "SqlScript")
	}

	if len(check) != 0 {
		errs = append(errs, errors.New(strings.Join(check, ", ")+" 不能为空"))
	}

	// 源和目标数据相同返回nil
	if r.GetSourceDb().GetDbType() == r.GetTargetDb().GetDbType() {
		return errors.Join(errs...)
	}

	if r.GetTargetDb().GetDbType() == DBType_POSTGRES {
		errs = append(errs, fmt.Errorf("暂不支持异构%s->%s", r.GetSourceDb().GetDbType(), r.GetTargetDb().GetDbType()))
	}

	return errors.Join(errs...)

}

var (
	defaultMysqlJdbcMapParams = map[string]string{
		"characterEncoding": "charset",
		"useTimezone":       "parseTime",
		"serverTimezone":    "loc",
		"useSSL":            "tls",
		// "allowPublicKeyRetrieval": "allowPublicKeyRetrieval",
	}
	defaultMysqlMapJdbcparams = map[string]string{
		"charset":   "characterEncoding",
		"parseTime": "useTimezone",
		"loc":       "serverTimezone",
		"tls":       "useSSL",
		// "allowPublicKeyRetrieval": "allowPublicKeyRetrieval",
	}
	defaultMysqlParams = map[string]string{
		"charset":   "utf8",
		"parseTime": "True",
		"loc":       "Local",
	}
)

func (x *Doris) GenerateDSN(sign string) (*string, error) {

	if err := x.Validate(); err != nil {
		return nil, err
	}

	generic := &GenericDB{
		Host:          x.GetHost(),
		Port:          x.GetQueryPort(),
		User:          x.GetUser(),
		Pwd:           x.GetPwd(),
		DbName:        x.GetDbName(),
		ConnectParams: x.GetConnectParams(),
	}

	return GenerateMysqlDSN(generic, sign)
}

func (x *DBDetail) GenerateDSN(sign string) (*string, error) {

	if x == nil {
		return nil, errors.New("DBDetail不能为空")
	}

	dbType := x.GetDbType()

	switch dbType {
	case DBType_POSTGRES:
		return GeneratePostgresDSN(x.GetPostgres(), sign)
	case DBType_DORIS:
		return GenerateDorisDSN(x.GetDoris(), sign)
	case DBType_MYSQL:
		return GenerateMysqlDSN(x.GetMysql(), sign)
	default:
		return nil, ecode.Errorf(ecode.DatabaseIncompatible, dbType.String())
	}

	// return nil, nil
}

func (x *DBDetail) GetDriverName() string {

	switch x.GetDbType() {
	case DBType_POSTGRES:
		return POSTGRES
	case DBType_DORIS, DBType_MYSQL:
		return MYSQL
	default:
		return ""
	}
}

func (x *DBDetail) TryConnect() error {

	dsn, err := x.GenerateDSN("ETLConnectVerify")
	if err != nil {
		return ecode.Errorfr(ecode.InvalidParameter, err)
	}

	// payload := x.GetPayload()
	// switch payload.(type) {
	// case *DBDetail_Doris:
	// 	dbType = MYSQL

	// case *DBDetail_Postgres:
	// 	dbType = POSTGRES
	// }

	var dbType string
	switch x.DbType {
	case DBType_POSTGRES:
		dbType = POSTGRES

	case DBType_DORIS:

		// // dbType = MYSQL
		// if err := x.GetDoris().TryConnect(); err != nil {
		// 	return err
		// }

		return x.GetDoris().TryConnect(dsn)

	case DBType_MYSQL:
		dbType = MYSQL

	default:
		return errors.New("暂不支持的数据库")
	}

	return CheckGenericDB(dbType, dsn)
}

// 检查通用db
func CheckGenericDB(dbType string, dsn *string) error {

	db, err := sql.Open(dbType, *dsn)
	if err != nil {
		logger.Error(dbType, *dsn)
		return ecode.Errorfr(ecode.InvalidParameter, err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Error(dbType, *dsn)
		return ecode.Errorfr(ecode.InvalidParameter, err)
	}

	var p string
	if err := db.QueryRow("select 'pong' as ping").Scan(&p); err != nil {
		return ecode.Errorfr(ecode.InvalidParameter, err)
	}

	return nil
}

func (x *WebHookRequest) GenerateUID() string {
	return HashString("*WebHookRequest", x.GetUrl(), x.GetUniqueId())
}

func (x *WebHookRequest) Validate() error {

	var check []string

	if x.GetUniqueId() == "" {
		check = append(check, "uniqueId")
	}

	if x.GetUrl() == "" {
		check = append(check, "url")
	}

	if len(check) != 0 {
		return errors.New(strings.Join(check, ", ") + " 不能为空")
	}

	client := resty.New()
	client.DisableWarn = true
	client.SetTimeout(time.Second * 60)
	client.SetHeader("User-Agent", "mbetl")

	resp, err := client.R().
		Get(x.GetUrl())
	if err != nil {
		return err
	}

	/*
		{
		    "code": 200,
		    "msg": null,
		    "data": {
		    }
		}
	*/

	if resp.StatusCode() != http.StatusOK {
		return errors.New("回调接口健康验证失败:" + resp.String())
	}

	var wh WebHookGetReply
	if err := plug.Unmarshal(resp.Body(), &wh); err != nil {
		return fmt.Errorf("反序列化%s响应结果%s失败", x.GetUrl(), resp.Body())
	}

	if wh.Code != 200 {
		return fmt.Errorf("回调接口健康验证失败:%+v", wh)
	}

	logger.Info(resp.String())

	return nil
}

type WebHookGetReply struct {
	Code int
	Msg  any
	Data any
}

func (x *DBDetail) Validate() error {

	if x == nil {
		return errors.New("DBDetail不能为空")
	}

	dbType := x.GetDbType()

	var err error
	switch dbType {
	case DBType_POSTGRES:
		err = x.GetPostgres().Validate()
	case DBType_DORIS:
		err = x.GetDoris().Validate()
	case DBType_MYSQL:
		err = x.GetMysql().Validate()
	case DBType_UNSPECIFIED:
		return errors.New("DBType不能为空")
	default:
		return errors.New("未兼容的数据库")
	}

	if err != nil {
		return err
	}

	return nil
}

func (x *DBDetail) GenerateUID() string {

	dbType := x.GetDbType()

	switch dbType {
	case DBType_POSTGRES:
		return x.GetPostgres().GenerateUID()
	case DBType_DORIS:
		return x.GetDoris().GenerateUID()
	case DBType_MYSQL:
		return x.GetMysql().GenerateUID()
	default:
		return ""
	}

}

// 验证流式导入端口
//
//	doris端口验证
//
//	curl http://192.168.99.56:8030/api/fe_version_info -u root:123456
//	curl http://192.168.99.56:8030/api/health
func (x *Doris) VerifyStreamLoadPort() error {

	baseUrl := fmt.Sprintf("http://%s:%d", x.Host, x.HttpPort)

	versionInfoUrl := baseUrl + "/api/fe_version_info"

	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		SetBasicAuth(x.User, x.Pwd).
		Get(versionInfoUrl)
	if err != nil {
		return err
	}

	var fever FeVersionInfoResp
	if err := sonic.Unmarshal(resp.Body(), &fever); err != nil {
		return err
	}

	if fever.Code != 0 ||
		fever.Msg != "success" ||
		fever.Data.FeVersionInfo.DorisBuildVersionPrefix != "doris" {
		return errors.New("doris StreamLoad端口验证失败")
	}

	version := fmt.Sprintf(
		"%d.%d.%d-%s",
		fever.Data.FeVersionInfo.DorisBuildVersionMajor,
		fever.Data.FeVersionInfo.DorisBuildVersionMinor,
		fever.Data.FeVersionInfo.DorisBuildVersionPatch,
		fever.Data.FeVersionInfo.DorisBuildVersionRcVersion,
	)

	return CheckDorisRequiredVersion(version)
}

func CheckDorisRequiredVersion(version string) error {

	requiredVer, _ := semver.NewVersion("3.0.4-rc02")
	currentVer, _ := semver.NewVersion(version)

	if currentVer.LessThan(requiredVer) {
		return errors.New("doris版本太低,需升级至3.0.4-rc02或更高")
	}

	return nil
}

func (x *Doris) VerifyQueryPort(dsn *string) error {

	// dsn, err := x.GenerateDSN("ETLConnectVerify")
	// if err != nil {
	// 	return ecode.Errorfr(ecode.InvalidParameter, err)
	// }

	return CheckGenericDB("mysql", dsn)
}

func (x *Doris) VerifyFlightPort() error {

	db, err := x.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	conn, err := db.Open(context.Background())
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.NewStatement()
	if err != nil {
		return fmt.Errorf("failed to create statement: %w", err)
	}
	defer stmt.Close()

	if err = stmt.SetSqlQuery("select 'pong' as ping"); err != nil {
		return fmt.Errorf("failed to set query: %w", err)
	}

	reader, _, err := stmt.ExecuteQuery(context.Background())
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	if !reader.Next() {
		return fmt.Errorf("failed to receive ping value")
	}

	record := reader.Record()

	fmt.Println(record)

	return nil
}

func (x *Doris) TryConnect(dsn *string) error {

	if err := x.VerifyStreamLoadPort(); err != nil {
		return err
	}

	if err := x.VerifyQueryPort(dsn); err != nil {
		return err
	}

	if err := x.VerifyFlightPort(); err != nil {
		return err
	}

	return nil
}

func (x *Doris) GenerateFlightURL() string {
	return fmt.Sprintf("grpc://%s:%d", x.GetHost(), x.GetFlightPort())
}

func (x *Doris) GenerateFlightOption() map[string]string {
	return map[string]string{
		adbc.OptionKeyURI:        x.GenerateFlightURL(),
		adbc.OptionKeyUsername:   x.GetUser(),
		adbc.OptionKeyPassword:   x.GetPwd(),
		drvfs.OptionTimeoutFetch: fmt.Sprintf("%d", 72*60*60), // 72小时
		drvfs.OptionTimeoutQuery: fmt.Sprintf("%d", 10*60),    // 10分钟
	}
}

func (x *Doris) NewDatabase() (adbc.Database, error) {

	options := x.GenerateFlightOption()

	// drvfs.OptionRPCCallHeaderPrefix
	// SELECT @@session.time_zone;
	var headers = map[string]string{
		"batch_size": fmt.Sprintf("%d", DorisBatchSize),
		"timezone":   "Asia/Shanghai",
	}

	for k, v := range headers {
		options[drvfs.OptionRPCCallHeaderPrefix+k] = v
	}

	for k, v := range x.ConnectParams {
		options[drvfs.OptionRPCCallHeaderPrefix+k] = v
	}

	var alloc memory.Allocator
	drv := drvfs.NewDriver(alloc)
	db, err := drv.NewDatabase(options)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}

func (x *Doris) GenerateUID() string {
	return HashString(
		"db",
		x.GetHost(),
		x.GetHttpPort(),
		x.GetQueryPort(),
		x.GetFlightPort(),
		x.GetUser(),
		x.GetPwd(),
		x.GetDbName(),
	)
}

func (x *Doris) Validate() error {

	if x == nil {
		return errors.New("Payload不能为空")
	}

	check := []string{}

	if x.GetHost() == "" {
		check = append(check, "Host")
	}

	if x.GetHttpPort() == 0 {
		check = append(check, "http_port")
	}

	if x.GetQueryPort() == 0 {
		check = append(check, "query_port")
	}

	if x.GetFlightPort() == 0 {
		check = append(check, "flight_port")
	}

	if x.GetUser() == "" {
		check = append(check, "User")
	}

	if x.GetPwd() == "" {
		check = append(check, "PassWord")
	}

	if x.GetDbName() == "" {
		check = append(check, "DBName")
	}

	if len(check) != 0 {
		return errors.New(strings.Join(check, ", ") + " 不能为空")
	}

	return nil
}

func (x *GenericDB) GenerateUID() string {
	return HashString(
		"db",
		x.GetHost(),
		x.GetPort(),
		x.GetUser(),
		x.GetPwd(),
		x.GetDbName(),
	)
}

func (x *GenericDB) Validate() error {

	if x == nil {
		return errors.New("Payload不能为空")
	}

	check := []string{}

	if x.GetHost() == "" {
		check = append(check, "Host")
	}

	if x.GetPort() == 0 {
		check = append(check, "Port")
	}

	if x.GetUser() == "" {
		check = append(check, "User")
	}

	if x.GetPwd() == "" {
		check = append(check, "PassWord")
	}

	if x.GetDbName() == "" {
		check = append(check, "DBName")
	}

	// if x.GetDbType() == "" {
	// 	check = append(check, "DBType")
	// }

	if len(check) != 0 {
		return errors.New(strings.Join(check, ", ") + " 不能为空")
	}

	return nil
}

// 示例
//
//	_, cancel := context.WithCancelCause(ctx)
//
// cancel(err)
func (x *Overview) GetErrorEnding(ctx context.Context, err error) {
	if x == nil {
		x = &Overview{}
	}

	if x.EndTime == "" {
		x.EndTime = time.Now().Local().Format(time.DateTime)
	}

	// if x.Detail == nil {
	// 	x.Detail = &Detail{}
	// }

	if x.Error == nil {
		x.Error = &Error{}
	}

	errs := plug.FromContextError(err)

	// 用户取消或失败
	switch ctx.Err() {
	case context.Canceled, context.DeadlineExceeded:
		x.RunStatus = RunStatus_Cancel

		if x.Error.Code == 0 {
			x.Error.Code = int32(ecode.UserCancellation) // ecode.UserCancellation
		}

		// err := context.Cause(ctx)
		// if err != nil {
		// 	x.Error.Msg = err.Error()
		// 	return
		// }

		if x.Error.Msg == "" {
			x.Error.Msg = errs.String()
		}

	default:

		x.RunStatus = RunStatus_Fail

		if x.Error.Code == 0 {
			x.Error.Code = int32(errs.Code())
		}
		if x.Error.Msg != "" {
			return
		}
		x.Error.Msg = errs.Message()
	}

}

type Node struct {
	ID       string  `json:"id"`
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

	root := &Node{ID: "0"}
	nodes := map[flow.Steper]*Node{
		work: root,
	}
	getNode := func(s flow.Steper) *Node {
		node, ok := nodes[s]
		if !ok {
			// node = &Node{ID: uuid.NewString()}
			t := flow.String(s)
			// id, _ := strconv.ParseUint(t, 10, 64)
			node = &Node{ID: t}
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

// 判断包含元素
func (x *Depends) container(depend *Depend) bool {

	if x == nil || depend == nil {
		return false
	}

	reqHash := Hash(depend)

	// x.mu.Lock()
	// defer x.mu.Unlock()

	for _, sli := range x.slice {

		depHash := Hash(sli)

		if reqHash == depHash {
			return true
		}

	}

	return false
}

// 查找来源ID
func getSourceID(src *Source, w *UKID, exclude string) {

	if src.GetType() == SourceType_MbEtl {

		// *w = append(*w, src.Id)
		id := src.GetId()

		// w.IDs = append(w.IDs, id)
		// w.Reqs[id] = true

		if id == exclude {
			return
		}

		w.Add(id)
	}
}

// GetSourceIDs查找来源IDs
//
// 入参
//
// w *[]string{}存储结果切片
func (x *Depends) GetSourceIDs(w *UKID, exclude string) {

	for _, depend := range x.Element() {

		getSourceID(depend.GetSource(), w, exclude)
		getSourceID(depend.GetTarget(), w, exclude)

	}

	// 去重
	util.Distinct(&w.IDs)
}

// 元素
func (x *Depends) Element() []*Depend {

	if x != nil {
		return x.slice
	}

	return nil
}

func (x *Depends) Len() int {

	if x != nil {
		return len(x.slice)
	}
	return 0
}

// 添加依赖
func (x *Depends) Add(depend *Depend) {

	if depend == nil {
		return
	}

	x.mu.Lock()
	defer x.mu.Unlock()

	if !x.container(depend) {
		x.slice = append(x.slice, depend)
	}

}

// Adds添加多个依赖
//
// 入参
// deps *Depends依赖结构体
func (x *Depends) Adds(deps *Depends) {

	if deps == nil || deps.slice == nil {
		return
	}

	for _, dep := range deps.slice {

		if dep.GetSource().GetType() == SourceType_MbEtl {

			x.Add(dep)

		}
	}
}

func NewEnableReply() *EnableReply {
	return &EnableReply{
		FlowData: &FlowMetadata{},
	}
}

func NewDepends() *Depends {
	return &Depends{
		slice: []*Depend{},
	}
}

func NewFlowMap() map[string]*FlowMap {
	return make(map[string]*FlowMap)
}

// 初始化BeforeFlow
func NewBeforeFlow() map[int64]*Graph {
	return make(map[int64]*Graph)
}

// LinkChildren连接子节点
//
// dag *DAG有向无环图
func (graph *Graph) LinkChildren(dag *DAG) error {

	// 是否循环依赖
	if dag.HasCycle() {
		return ecode.Errorf(ecode.LoopDependency)
	}

	graph.Children = make([]*Source, len(dag.nodes))

	idx := 0
	for _, item := range dag.nodes {
		graph.Children[idx] = &Source{
			Id:     item.Id,
			BaseId: item.BaseId,
			Type:   item.Type,
		}
		// if p.reply.GetFlowData().GetTaskType() == mbpb.TaskType_flowTask && p.basic.Uid == item.Id {
		// 	graph.Children[idx].TaskType = mbpb.TaskType_flowTask
		// }
		idx++
	}

	graph.Edges = dag.Edges

	return nil
}

// 获取父节点
func (x *Graph) GetParentNode(node *Source) []*Source {

	var (
		srcIds []string
		result []*Source
	)

	for _, edge := range x.GetEdges() {

		if edge.GetTarget() == node.GetId() {
			srcIds = append(srcIds, edge.GetSource())
		}
	}

	filter := func(req *Source) bool {

		for _, src := range srcIds {

			if src == req.GetId() {
				return true
			}

		}

		return false
	}

	for _, children := range x.GetChildren() {

		if filter(children) {
			result = append(result, children)
		}
	}

	return result

}

func (x *Table) GetSourceID() string {

	return x.GetSource().GetId()
}

func (x *Table) GetFQN() string {

	if x != nil {
		if x.GetSchema() != "" && x.GetTable() != "" {
			return fmt.Sprintf(`"%s"."%s"`, x.GetSchema(), x.GetTable())
		}
	}

	return ""
}

func (e *Error) Error() string {
	return e.GetMsg()
}

// 联合关系
type UnionRelation struct {
	Host   string `json:"host,omitempty"`
	Port   int32  `json:"port,omitempty"`
	DBType DBType `json:"dbType,omitempty"`
	DBName string `json:"dbName,omitempty"`
	Schema string `json:"schema,omitempty"`
	Table  string `json:"table,omitempty"`
}

// 获取联合关系
func GetUnionRelation(db *DBDetail, table *Table) (*UnionRelation, error) {
	switch db.GetDbType() {
	case DBType_DORIS:
		c := db.GetDoris()
		return &UnionRelation{
			Host:   c.GetHost(),
			Port:   c.GetQueryPort(),
			DBType: (db.GetDbType()),
			DBName: c.GetDbName(),
			Schema: table.GetSchema(),
			Table:  table.GetTable(),
		}, nil
	case DBType_POSTGRES:
		c := db.GetPostgres()
		return &UnionRelation{
			Host:   c.GetHost(),
			Port:   c.GetPort(),
			DBType: db.GetDbType(),
			DBName: c.GetDbName(),
			Schema: table.GetSchema(),
			Table:  table.GetTable(),
		}, nil
	case DBType_MYSQL:
		c := db.GetMysql()
		return &UnionRelation{
			Host:   c.GetHost(),
			Port:   c.GetPort(),
			DBType: db.GetDbType(),
			DBName: c.GetDbName(),
			Schema: table.GetSchema(),
			Table:  table.GetTable(),
		}, nil
	default:
		return nil, errors.New("不支持的数据库")
	}
}

type ColumnType struct {
	Typtype       byte            // 类型的类型
	Typrelid      uint32          //
	Typname       string          // 类型名
	Nspname       string          // 命名空间
	Spec          bool            // 用户定义
	CompositeType []CompositeType // 复合类型
}

// 复合类型
type CompositeType struct {
	Attname    string // 字段名
	Type       string // 类型
	Attnotnull bool   // 是否为空
}
