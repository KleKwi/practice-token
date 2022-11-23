package db

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var (
	engine    *xorm.Engine
	tables    []interface{}
	initFuncs []func() error
)

func GetEngine() *xorm.Engine {
	return engine
}

func SetDefaultEngine(eng *xorm.Engine) {
	engine = eng
}

// RegisterModel registers model, if initfunc provided, it will be invoked after data model sync
func RegisterModel(bean interface{}, initFunc ...func() error) {
	tables = append(tables, bean)
	if initFunc[0] != nil {
		initFuncs = append(initFuncs, initFunc[0])
	}
}

// SyncAllTables sync the schemas of all tables, is required by unit test code
func SyncAllTables() error {
	return engine.StoreEngine("InnoDB").Sync2(tables...)
}

// InitDB initialize the database
func InitDB() (err error) {
	if err := os.MkdirAll("./data", 0755); err != nil {
		logrus.Fatalln(err)
	}
	x, err := xorm.NewEngine("sqlite3", "./data/token.db")
	if err != nil {
		return
	}

	x.SetMapper(names.GonicMapper{})

	SetDefaultEngine(x)

	return
}

// Migrate migrate table and data
func Migrate() (err error) {
	err = SyncAllTables()
	if err != nil {
		return
	}

	for _, initFunc := range initFuncs {
		err = initFunc()
		if err != nil {
			return
		}
	}

	return
}

func FirstOrCreate(bean interface{}) (has bool, err error) {
	has, err = GetEngine().Get(bean)
	if err != nil {
		return
	}

	if !has {
		_, err = GetEngine().Insert(bean)
		if err != nil {
			return
		}
		has = true
	}

	return
}
