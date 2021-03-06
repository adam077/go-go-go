package scheduler

import (
	"go-go-go/src/data"
	"testing"
)

func TestRunner(t *testing.T) {
	t1()
}

func t1() {
	dbMeta := data.GetDataDB("default")
	//dbMeta.Exec("select 1;")
	asdf := []data.ModelInterface{
		&data.TableTest{},
	}
	for x := range asdf {
		dbMeta.DropTable(asdf[x])
		data.MigrateTable(dbMeta, asdf[x])
	}
}
