package mbpb

import (
	"database/sql"
	"fmt"
	"maps"
	"mbetl/runtime/ecode"
	"mbetl/runtime/util"
	"mbetl/worker/consts"
	"net/url"
	"sort"
	"strings"
	"sync"

	"github.com/forhsd/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
)

const (
	POSTGRES   string = "postgres"
	POSTGRESQL string = "postgresql"
	MYSQL      string = "mysql"
)

// PG DSN
func GeneratePostgresDSN(req *GenericDB, sign string) (*string, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	// ?sslmode=disable&TimeZone=Asia/Shanghai&Application_Name=etl_reader&connect_timeout=1
	dorp := ""
	keys := sync.Map{}
	for key, val := range req.GetConnectParams() {

		_, ok := consts.PGSQL_OPTIONS[key]
		if ok {
			// keys[key] = val
			keys.Store(key, val)
		} else {
			dorp += "&" + strings.Join([]string{key, val}, "=")
		}

	}

	// appName := keys["Application_Name"]
	appName, ok := keys.Load("Application_Name")
	// _ = ok
	_ = appName

	// (appName == "" || appName == nil)
	if sign != "" && !ok {
		// keys["Application_Name"] = sign
		keys.Store("Application_Name", sign)
	}

	for key, val := range consts.PGSQL_OPTIONS_DEFAULT {

		// _, ok := keys[key]
		_, ok := keys.Load(key)
		if !ok {
			// keys[key] = val
			keys.Store(key, val)
		}

	}

	n := 0
	param := make([]string, util.MapLenght(&keys))
	keys.Range(func(key, value any) bool {
		param[n] = fmt.Sprintf("%v=%v", key, value)
		n++
		return true
	})

	sort.Strings(param)

	params := strings.Join(param, "&")

	dsn := fmt.Sprintf(
		consts.PGSQL_DSN_TEMPLATE,
		POSTGRES,
		req.GetUser(),
		req.GetPwd(),
		req.GetHost(),
		req.GetPort(),
		req.GetDbName(),
		params,
	)

	if dorp != "" {
		logger.Warn("丢弃的参数", dorp)
	}

	return &dsn, nil
}

func GenerateMysqlDSN(x *GenericDB, sign string) (*string, error) {

	// root:mysqlHolder@tcp(192.168.100.138:3306)/hst_takeaway_887f2181-eb06-4d77-b914-7c37c884952c_db?charset=utf8&parseTime=True&loc=Local

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", x.User, x.Pwd, x.Host, x.Port, x.DbName)

	// useUnicode=true
	// characterEncoding=utf8          = charset=utf8
	// useTimezone=true                = parseTime=true
	// serverTimezone=Asia/Shanghai    = loc=Asia%2FShanghai
	// useSSL=false                    = tls=false
	// allowPublicKeyRetrieval=true"   = allowPublicKeyRetrieval=true

	var ConnectParams = map[string]string{}
	maps.Copy(ConnectParams, x.GetConnectParams())

	// ConnectParams := x.GetConnectParams()

	var param = map[string]string{}

	for key, val := range defaultMysqlParams {

		jdbckey, exists := defaultMysqlMapJdbcparams[key]
		if exists {

			connectVal, ok := ConnectParams[jdbckey]
			if ok {
				// param = append(param, fmt.Sprintf("%s=%s", key, goconnectVal))

				param[key] = fmt.Sprintf("%s=%s", key, url.QueryEscape(connectVal))
			}
			delete(ConnectParams, jdbckey)

		} else {
			// param = append(param, fmt.Sprintf("%s=%s", key, val))
			param[key] = fmt.Sprintf("%s=%s", key, val)
		}
	}

	for key, val := range ConnectParams {
		gokey, exists := defaultMysqlJdbcMapParams[key]
		if exists {
			// param = append(param, fmt.Sprintf("%s=%s", gokey, val))
			param[gokey] = fmt.Sprintf("%s=%s", gokey, url.QueryEscape(val))
			delete(ConnectParams, key)
		}

	}

	vals := make([]string, 0, len(param))
	for _, val := range param {
		vals = append(vals, val)
	}

	sort.Strings(vals)

	params := strings.Join(vals, "&")

	if params != "" {
		dsn = strings.Join([]string{dsn, params}, "?")
	}

	return &dsn, nil
}

func GenerateDorisDSN(req *Doris, sign string) (*string, error) {

	return req.GenerateDSN(sign)
}

// 获取驱动类型
func GetDriverType(db *sql.DB) (string, error) {

	driver := db.Driver()

	switch driver.(type) {
	case *pq.Driver:
		return POSTGRES, nil
	case *mysql.MySQLDriver:
		return MYSQL, nil
	default:
		return "",
			ecode.Errorf(ecode.DatabaseIncompatible)
	}
}

type FeVersionInfoResp struct {
	Code  int    `json:"code"`
	Count int    `json:"count"`
	Data  FeData `json:"data"`
	Msg   string `json:"msg"`
}

type FeVersionInfo struct {
	DorisBuildHash             string `json:"dorisBuildHash"`
	DorisBuildInfo             string `json:"dorisBuildInfo"`
	DorisBuildShortHash        string `json:"dorisBuildShortHash"`
	DorisBuildTime             string `json:"dorisBuildTime"`
	DorisBuildVersion          string `json:"dorisBuildVersion"`
	DorisBuildVersionMajor     int    `json:"dorisBuildVersionMajor"`
	DorisBuildVersionMinor     int    `json:"dorisBuildVersionMinor"`
	DorisBuildVersionPatch     int    `json:"dorisBuildVersionPatch"`
	DorisBuildVersionPrefix    string `json:"dorisBuildVersionPrefix"`
	DorisBuildVersionRcVersion string `json:"dorisBuildVersionRcVersion"`
	DorisJavaCompileVersion    string `json:"dorisJavaCompileVersion"`
}

type FeData struct {
	FeVersionInfo FeVersionInfo `json:"feVersionInfo"`
}
