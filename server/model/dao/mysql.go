/*
 * error code: 30001000 ` 30001999
 */

package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server/config"
	"server/core/logger"
	"strings"
	"time"
)

// SchemaMeta struct
type SchemaMeta struct {
	DBName    string //数据库名
	TableName string //表
	Field     string //字段名
	Type      string //字段类型
	Comment   string //字段备注
}

// DBBase struct
type DBBase struct {
	db  *sql.DB
	ctx context.Context
}

// db connect struct
var dbConn *sql.DB

func init() {
	if err := createConnDB(context.TODO()); err != nil {
		panic("[create db connect error]")
	} else {
		fmt.Println("[database init successfully] host:", config.Config.Mysql.Host)
	}
}

// create db connect
func createConnDB(ctx context.Context) error {
	host := config.Config.Mysql.Host
	port := config.Config.Mysql.Port
	user := config.Config.Mysql.Username
	pass := config.Config.Mysql.Password
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/test?charset=utf8", user, pass, host, port)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf(`mysql connect is error: %v`, err))
		return err
	}
	if err := db.Ping(); err != nil {
		logger.Error(ctx, fmt.Sprintf(`mysql ping is error: %v`, err))
		return err
	}
	//设置连接可以重用的最长时间
	db.SetConnMaxLifetime(5 * time.Minute)
	//设置与数据库的最大打开连接数,如果 n <= 0,则对打开的连接数没有限制。默认值为 0（无限制）。
	db.SetMaxOpenConns(10)
	//设置空闲连接池的最大连接数,如果 n <= 0，则不保留任何空闲连接。
	db.SetMaxIdleConns(10)
	//设置连接可能处于空闲状态的最长时间
	db.SetConnMaxIdleTime(5 * time.Minute)
	//dbConn
	dbConn = db

	//return
	return nil
}

// NewDBBase
func NewDBBase(ctx context.Context) *DBBase {
	if dbConn == nil {
		if err := createConnDB(ctx); err != nil {
			logger.Error(ctx, "[1000100] create db connect err")
			return nil
		}
	}

	//return
	return &DBBase{
		db:  dbConn,
		ctx: ctx,
	}
}

// get
func (d *DBBase) Get(table string, uid int) (any, error) {
	sqlText := fmt.Sprintf(`SELECT content FROM %v WHERE uid = ? LIMIT 1`, getTableName(uid, table))
	stmt, err := d.db.Prepare(sqlText)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000110] get error: %v`, err))
		return nil, err
	}
	defer stmt.Close()

	var data any
	if err := stmt.QueryRow(uid).Scan(&data); err == sql.ErrNoRows {
		defaultData, err := d.initData(table, uid)
		if err != nil {
			logger.Error(d.ctx, fmt.Sprintf(`[1000112] get error: %v`, err))
			return nil, err
		}
		return defaultData, nil
	} else if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000114] get error: %v`, err))
		return nil, err
	}
	if data != nil {
		return data, nil
	}

	logger.Error(d.ctx, "[1000119] get unknow error")
	return nil, errors.New("unknow error")
}

// modify
func (d *DBBase) Modify(table string, uid int, data []byte) error {
	sql := fmt.Sprintf(`UPDATE %v SET content=?, stime=? WHERE uid=? LIMIT 1`, getTableName(uid, table))
	stmt, err := d.db.Prepare(sql)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000120] modify error: %v`, err))
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(data, time.Now().Unix(), uid)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000122] modify error: %v`, err))
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000124] modify error: %v`, err))
		return err
	}
	if num == 1 {
		return nil
	}
	if num == 0 {
		logger.Error(d.ctx, "[1000128] modify error")
		return ErrorNoRows
	}

	logger.Error(d.ctx, `[1000129] modify unknow error`)
	return errors.New("unknow error")
}

// init
func (d *DBBase) initData(table string, uid int) (any, error) {
	defaultData := config.GetDefaultDBValue(table)
	sqltext := fmt.Sprintf(`INSERT INTO %v(uid, content, stime) VALUES(?, ?, ?)`, getTableName(uid, table))
	stmt, err := d.db.Prepare(sqltext)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000130] init data error: %v`, err))
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(uid, defaultData, time.Now().Unix())
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000132] init data error: %v`, err))
		return nil, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000134] init data error: %v`, err))
		return nil, err
	}
	if num == 0 {
		logger.Error(d.ctx, fmt.Sprintf(`[1000138] init data error: %v`, err))
		return nil, ErrorNoRows
	} else if num == 1 {
		return defaultData, nil
	}

	logger.Error(d.ctx, "[1000139] nuknow error")
	return nil, errors.New("nuknow error")
}

/* ----------------------------------run raw sql---------------------------------- */

// 查询单条数据并返回error
func (d *DBBase) QueryRow(rawsql string, scanArgs ...any) error {
	if err := d.db.QueryRow(rawsql).Scan(scanArgs...); err == sql.ErrNoRows {
		logger.Error(d.ctx, fmt.Sprintf(`[1000140] query row error: %v`, err))
		return ErrorNoRows
	} else if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000142] query row error: %v`, err))
		return err
	}

	//return
	return nil
}

// 查询多条数据并返回*sql.Row、error
func (d *DBBase) Query(rawsql string) (*sql.Rows, error) {
	rows, err := d.db.Query(rawsql)
	if err == sql.ErrNoRows { //这里不会被触发,通常会在Scan时触发error
		//logger.Error(d.ctx, fmt.Sprintf(`[1000150] query error: %v`, err))
		return nil, ErrorNoRows
	}
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000154] query error: %v`, err))
		return nil, err
	}
	//defer rows.Close()

	return rows, nil
}

// 执行SQL并返回是否成功
func (d *DBBase) Exec(rawsql string, args ...any) error {
	result, err := d.db.Exec(rawsql, args...)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000155] exec error: %v`, err))
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000156] exec error: %v`, err))
		return err
	}
	if n == 0 {
		logger.Error(d.ctx, "[1000158] exec the affected rows is the zero.")
		return ErrorNoRows
	}
	if n > 0 {
		return nil
	}

	logger.Error(d.ctx, "[1000159] exec unknow error.")
	return errors.New("exec unknow error.")
}

// 插入并返回error
func (d *DBBase) Inert(rawsql string, args ...any) error {
	result, err := d.db.Exec(rawsql, args...)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000170] insert error: %v`, err))
		return err
	}
	if n, err := result.RowsAffected(); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000175] insert error: %v`, err))
		return err
	} else if n == 0 {
		logger.Error(d.ctx, fmt.Sprintf(`[1000178] insert error: %v`, "the affected rows is the zero"))
		return ErrorNoRows
	} else if n > 0 {
		return nil
	}

	logger.Error(d.ctx, "[1000169] insert unknow error")
	return errors.New("insert unknow error")
}

// 插入并返回ID
func (d *DBBase) InertID(rawsql string, args ...any) (int64, error) {
	result, err := d.db.Exec(rawsql, args...)
	if err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000180] insert id error: %v`, err))
		return 0, err
	}
	if n, err := result.RowsAffected(); err != nil {
		logger.Error(d.ctx, fmt.Sprintf(`[1000182] insert id error: %v`, err))
		return 0, err
	} else if n == 0 {
		logger.Error(d.ctx, fmt.Sprintf(`[1000184] insert id error: %v`, "affected row is the zero"))
		return 0, ErrorNoRows
	} else if n == 1 {
		newId, err := result.LastInsertId()
		if err != nil {
			logger.Error(d.ctx, fmt.Sprintf(`[1000188] insert id error: %v`, err))
			return 0, err
		}
		return newId, nil
	}

	logger.Error(d.ctx, `[1000189] insert id unknow error`)
	return 0, errors.New("insert id unknow error")
}

/* ----------------------------------schema meta---------------------------------- */

//获取表结构
func (d *DBBase) GetTableSchemaMeta(tableName string) ([]SchemaMeta, error) {
	//list, _ := db.Query(fmt.Sprintf(`show columns from %s`, tableName))
	list, err := d.db.Query(fmt.Sprintf("SELECT `TABLE_SCHEMA`,`TABLE_NAME`,`COLUMN_NAME`,`DATA_TYPE`,`COLUMN_COMMENT` FROM `COLUMNS` WHERE TABLE_NAME='%v'", tableName))
	if err != nil {
		return nil, err
	}
	defer list.Close()

	metas := make([]SchemaMeta, 0, 50)
	for list.Next() {
		var data SchemaMeta
		err := list.Scan(&data.DBName, &data.TableName, &data.Field, &data.Type, &data.Comment)
		if err != nil {
			return nil, err
		}
		metas = append(metas, data)
	}

	return metas, nil
}

//生成表表结构Struct
func (d *DBBase) GenTableStruct(tableName string, metas []SchemaMeta) string {
	var fieldValue string

	//字段处理
	for _, v := range metas {
		ftype := "any"
		if strings.Contains(v.Type, "int") {
			ftype = "int"
		} else if strings.Contains(v.Type, "char") {
			ftype = "string"
		} else if strings.Contains(v.Type, "datetime") {
			ftype = "time.Time"
		}

		field := v.Field
		if strings.Contains(field, "_") {
			fields := strings.Split(field, "_")
			for k, v := range fields {
				fields[k] = fmt.Sprintf(`%s%s`, strings.ToUpper(v[:1]), v[1:])
			}
			field = fmt.Sprintf(`%s`, strings.Join(fields, ""))
		} else {
			field = strings.ToUpper(field[:1]) + field[1:]
		}

		comment := ""
		if v.Comment != "" {
			comment = "//" + v.Comment
		}
		fieldValue += fmt.Sprintf("%s %s	`json:\"%v\"` %s \n", field, ftype, strings.ToLower(v.Field), comment)
	}

	//表名处理
	if strings.Contains(tableName, "_") {
		tblName := strings.Split(tableName, "_")
		for k, v := range tblName {
			tblName[k] = fmt.Sprintf(`%s%s`, strings.ToUpper(v[:1]), v[1:])
		}
		tableName = fmt.Sprintf(`%s%s`, strings.Join(tblName, ""), "Table")
	}

	//备注
	structComment := fmt.Sprintf("//%v Struct \n", tableName)

	//return
	return fmt.Sprintf("%stype %s struct {\n%s}", structComment, tableName, fieldValue)
}

//程序结束之后的清理工作
func FinishClear() {
	dbConn.Close()
	dbConn = nil
}
