package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"
)

var once sync.Once
var dbConn *sql.DB

type XSQL struct {
	ctx context.Context
	db  *sql.DB

	table     string   //表名
	primary   string   //表主键
	fields    []string //字段
	values    []any    //字段-值
	where     []string //条件
	order     string   //排序
	group     string   //分组
	have      string   //分组条件
	leftJoin  string   //左关联
	rightJoin string   //右关联
	on        string   //on条件
	sql       string   //sql
}

// NewXSQL
func NewXSQL(ctx context.Context, config *Config) *XSQL {
	xsql := &XSQL{
		ctx:     ctx,
		primary: "id", //默认主键
		fields:  make([]string, 0, 20),
		values:  make([]any, 0, 20),
		where:   make([]string, 0, 5),
	}
	once.Do(func() {
		if err := xsql.connect(config); err != nil {
			panic(fmt.Errorf(`new xsql happened error:%v`, err))
		}
	})
	if dbConn == nil {
		xsql.connect(config)
	}
	xsql.db = dbConn

	//return
	return xsql
}

// connect
func (d *XSQL) connect(config *Config) error {
	var err error
	if config.Charset == "" {
		config.Charset = "utf8"
	}
	if config.ConnMaxLifetime == 0 {
		config.ConnMaxLifetime = 300 * time.Second
	}
	if config.ConnMaxIdleTime == 0 {
		config.ConnMaxIdleTime = 300 * time.Second
	}
	if config.MaxOpenConns == 0 {
		config.MaxOpenConns = 10
	}
	if config.MaxIdleConns == 0 {
		config.MaxIdleConns = 10
	}
	dbConn, err = sql.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=%s`, config.UserName, config.Password, config.Host, config.Port, config.DBName, config.Charset))
	if err != nil {
		return err
	}
	if dbConn == nil {
		return errors.New("sql.open happened error")
	}
	if err = dbConn.Ping(); err != nil {
		return err
	}

	//设置连接可以重用的最长时间
	dbConn.SetConnMaxLifetime(config.ConnMaxLifetime)
	//设置连接可能处于空闲状态的最长时间
	dbConn.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	//设置与数据库的最大打开连接数,如果 n <= 0,则对打开的连接数没有限制(默认值为0,无限制)
	dbConn.SetMaxOpenConns(config.MaxOpenConns)
	//设置空闲连接池的最大连接数,如果 n <= 0，则不保留任何空闲连接
	dbConn.SetMaxIdleConns(config.MaxIdleConns)

	//Finalizer
	runtime.SetFinalizer(dbConn, func(conn *sql.DB) {
		conn.Close()
	})

	//return
	return nil
}

// Table 字段
func (d *XSQL) Table(table string) *XSQL {
	d.table = table

	//return
	return d
}

// Primary 主键
func (d *XSQL) Primary(key string) *XSQL {
	d.primary = key

	//return
	return d
}

// Fields 字段
func (d *XSQL) Fields(fields ...string) *XSQL {
	if len(fields) > 0 {
		d.fields = append(d.fields, fields...)
	}
	if false == slices.Contains(d.fields, d.primary) {
		d.fields = append(d.fields, d.primary)
	}

	//return
	return d
}

// Where 条件
func (d *XSQL) Where(condition string) *XSQL {
	if condition != "" {
		if len(d.where) == 0 {
			d.where = append(d.where, condition)
		} else {
			d.where = append(d.where, " AND ", condition)
		}
	}

	//return
	return d
}

// ORWhere 条件
func (d *XSQL) ORWhere(condition string) *XSQL {
	if condition != "" {
		if len(d.where) == 0 {
			d.where = append(d.where, condition)
		} else {
			d.where = append(d.where, " OR ", condition)
		}
	}

	//return
	return d
}

// GroupBy 分组
func (d *XSQL) GroupBy(group string) *XSQL {
	d.group = group

	//return
	return d
}

// Having 分组条件
func (d *XSQL) Having(having string) *XSQL {
	d.have = having

	//return
	return d
}

// LeftJoin 关联
func (d *XSQL) LeftJoin(join string) *XSQL {
	d.leftJoin = join

	//return
	return d
}

// RightJoin 关联
func (d *XSQL) RightJoin(join string) *XSQL {
	d.rightJoin = join

	//return
	return d
}

// ON 关联
func (d *XSQL) ON(on string) *XSQL {
	d.on = on

	//return
	return d
}

// OrderBy 排序
func (d *XSQL) OrderBy(order string) *XSQL {
	d.order = order

	//return
	return d
}

// QueryRow 查询单条数据
func (d *XSQL) QueryRow() (map[string]any, error) {
	data, err := d.Query()
	if err != nil {
		return nil, err
	}

	return data[0], nil
}

// Query 查询数据
func (d *XSQL) Query() ([]map[string]any, error) {
	defer d.RestSQL()

	//生成SQL
	d.sql = d.GenRawSQL()

	//QUERY
	rows, err := d.db.Query(d.sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//字段 - 数据
	data := make([]map[string]any, 0, 20)
	entity := genEntity(len(d.fields))
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, err
		}
		data = append(data, genRecord(entity, d.fields))
	}
	if len(data) < 1 {
		return nil, sql.ErrNoRows
	}

	//return
	return data, nil
}

// QueryMap 查询数据
func (d *XSQL) QueryMap() (map[int]map[string]any, error) {
	defer d.RestSQL()

	//生成SQL
	d.sql = d.GenRawSQL()

	//QUERY
	rows, err := d.db.Query(d.sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//字段 - 数据
	data := make(map[int]map[string]any, 50)
	entity := genEntity(len(d.fields))
	for rows.Next() {
		if err := rows.Scan(entity...); err != nil {
			return nil, err
		}
		record := genRecord(entity, d.fields)
		data[cast.ToInt(record[d.primary])] = record
	}
	if len(data) < 1 {
		return nil, sql.ErrNoRows
	}

	//return
	return data, nil
}

// GenRawSQL 生成查询SQL
func (d *XSQL) GenRawSQL() string {
	var rawsql strings.Builder
	if len(d.fields) > 0 && d.table != "" {
		rawsql.WriteString(fmt.Sprintf("SELECT %v FROM %v", strings.Join(d.fields, ","), d.table))
	}
	if len(d.where) > 0 {
		rawsql.WriteString(fmt.Sprintf(` WHERE %v`, strings.Join(d.where, "")))
	}
	if d.group != "" {
		rawsql.WriteString(fmt.Sprintf(` GROUP BY %v`, d.order))
	}
	if d.have != "" {
		rawsql.WriteString(fmt.Sprintf(` HAVING %v`, d.have))
	}
	if d.leftJoin != "" {
		rawsql.WriteString(fmt.Sprintf(` LEFT JOIN %v`, d.leftJoin))
	}
	if d.rightJoin != "" {
		rawsql.WriteString(fmt.Sprintf(` RIGHT JOIN %v`, d.rightJoin))
	}
	if d.on != "" {
		rawsql.WriteString(fmt.Sprintf(` ON %v`, d.on))
	}
	if d.order != "" {
		rawsql.WriteString(fmt.Sprintf(` ORDER BY %v`, d.order))
	}

	//return
	return rawsql.String()
}

// Insert 插入数据
func (d *XSQL) Insert(params map[string]any) *XSQL {
	for k, v := range params {
		d.fields = append(d.fields, k)
		d.values = append(d.values, v)
	}
	if len(params) > 0 {
		d.sql = fmt.Sprintf("INSERT INTO %v(%v) VALUES (%v)", d.table, strings.Join(d.fields, ","), strings.Repeat(",?", len(d.fields))[1:])
	}

	//return
	return d
}

// Modify 修改数据
func (d *XSQL) Modify(params map[string]any) *XSQL {
	for k, v := range params {
		d.fields = append(d.fields, fmt.Sprintf(`%v=?`, k))
		d.values = append(d.values, v)
	}
	if len(params) > 0 {
		d.sql = fmt.Sprintf(`UPDATE %v SET %v WHERE %v`, d.table, strings.Join(d.fields, ","), strings.Join(d.where, ""))
	}

	//return
	return d
}

// Delete 删除数据
func (d *XSQL) Delete() *XSQL {
	d.sql = fmt.Sprintf(`DELETE FROM %v WHERE %v`, d.table, strings.Join(d.where, ""))

	//return
	return d
}

// Exec 执行SQL
func (d *XSQL) Exec() (sql.Result, error) {
	defer d.RestSQL()

	stmt, err := d.db.Prepare(d.sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(d.values...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RestSQL
func (d *XSQL) RestSQL() {
	d.table = ""
	d.fields = make([]string, 0, 20)
	d.values = make([]any, 0, 20)
	d.where = make([]string, 0, 5)
	d.group, d.have = "", ""
	d.order = ""
	d.on, d.leftJoin, d.rightJoin = "", "", ""

	d.sql = ""
}
