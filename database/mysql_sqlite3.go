package database

import (
	"database/sql"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	_ "github.com/mattn/go-sqlite3"
)

var (
	W = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
	})
	_ InterfaceDataBase = (*ObjDataBase)(nil)
)

type (
	InterfaceDataBase interface {
		DataBaseInit(driverName, dataSourceName string) (ok bool)
		DataBaseCreatTables(DDL string) bool
		DataBaseQuery(query string) (ok bool)
		DataBaseQueryResult() interface{}
		DataBaseUpdate(query string, args ...interface{}) (ok bool)
		DataBaseInsert(query string, args ...interface{}) (ok bool)
	}
	ObjDataBase struct {
		//Client *redis.Client //写到另外的文件，移除工程的爬虫工程的全局变量
		db          *sql.DB
		stmt        *sql.Stmt
		rows        *sql.Rows
		result      sql.Result
		queryResult interface{}
		err         error
	}
)

func (p *ObjDataBase) DataBaseInit(driverName, dataSourceName string) (ok bool) {
	p.db, p.err = sql.Open(driverName, dataSourceName)
	if !W.CheckErr(p.err) {
		return
	}
	p.db.SetMaxOpenConns(1000)
	p.db.SetMaxIdleConns(30000)
	if !W.CheckErr(p.db.Ping()) {
		return
	}
	return true
}

func (p *ObjDataBase) DataBaseCreatTables(DDL string) bool {
	return W.CheckErr2(p.db.Exec(DDL))
}

func (p *ObjDataBase) DataBaseQueryResult() interface{} {
	return p.queryResult
}

func (p *ObjDataBase) DataBaseQuery(query string) (ok bool) {
	p.rows, p.err = p.db.Query(query)
	if !W.CheckErr(p.err) {
		return
	}
	defer func() { W.CheckErr(p.rows.Close()) }()
	for p.rows.Next() {
		if !(W.CheckErr(p.rows.Scan(&p.queryResult))) {
			return
		}
	}
	return true
}

func (p *ObjDataBase) DataBaseUpdate(query string, args ...interface{}) (ok bool) {
	return p.dataBaseStmtExec(query, args)
}

func (p *ObjDataBase) DataBaseInsert(query string, args ...interface{}) (ok bool) {
	return p.dataBaseStmtExec(query, args)
}

func (p *ObjDataBase) dataBaseStmtExec(query string, args ...interface{}) (ok bool) {
	p.stmt, p.err = p.db.Prepare(query)
	if !W.CheckErr(p.err) {
		return
	}
	defer func() { W.CheckErr(p.stmt.Close()) }()
	p.result, p.err = p.stmt.Exec(args)
	if !W.CheckErr(p.err) {
		return
	}
	return true
}
