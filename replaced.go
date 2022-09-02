package replaced

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/stdlib"
	proxy "github.com/shogo82148/go-sql-proxy"
)

var hooks = &proxy.HooksContext{
	PreExec: func(_ context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
		stmt.QueryString = replaceQuery(stmt.QueryString)
		return nil, nil
	},
	PreQuery: func(_ context.Context, stmt *proxy.Stmt, args []driver.NamedValue) (interface{}, error) {
		stmt.QueryString = replaceQuery(stmt.QueryString)
		return nil, nil
	},
}

func replaceQuery(q string) string {
	count := 1
	for {
		n := strings.Replace(q, "?", fmt.Sprintf("$%d", count), 1)
		if n == q {
			break
		}
		count++
		q = n
	}
	return q
}

func init() {
	sql.Register("pgx-replaced", proxy.NewProxyContext(&stdlib.Driver{}, hooks))
}
