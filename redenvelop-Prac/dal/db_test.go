package dal

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()

	if rdb == nil {
		t.Errorf("rdb is nil")
	}
}
