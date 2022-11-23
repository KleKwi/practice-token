package unittest

import (
	"token/models/db"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// PrepareDB accept models, and prepare memory database for testing.
func PrepareTestDB() (err error) {
	err = createTestEngine()
	if err != nil {
		return
	}

	err = db.Migrate()
	if err != nil {
		return
	}

	return
}

func createTestEngine() (err error) {
	x, err := xorm.NewEngine("sqlite3", "file::memory:?cache=shared&_txlock=immediate")
	if err != nil {
		return
	}
	x.SetMapper(names.GonicMapper{})
	x.ShowSQL(true)

	db.SetDefaultEngine(x)

	return
}
