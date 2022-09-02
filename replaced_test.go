package replaced_test

import (
	"database/sql"
	"testing"

	_ "github.com/mackee/pgx-replaced"
)

func TestRun(t *testing.T) {
	db, err := sql.Open("pgx-replaced", "postgres://postgres:mysecretpassword@localhost:5432/postgres")
	if err != nil {
		t.Errorf("error sql.Open: %s", err)
		t.Fail()
	}
	row := db.QueryRow("SELECT * FROM sample1 WHERE id = ?", 1)
	var (
		id   int64
		name string
	)
	if err := row.Scan(&id, &name); err != nil {
		t.Errorf("error row.Scan: %s", err)
		t.Fail()
	}
	if id != 1 {
		t.Errorf("id is not 1")
	}
	if name != "hogehoge" {
		t.Errorf("name is not hogehoge")
	}
}
