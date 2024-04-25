package mysql

import (
	"bytes"
	"fmt"

	"sniffer-agent/model"
	"sniffer-agent/tidb/util/hack"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/format"
	driver "github.com/pingcap/tidb/types/parser_driver"
)

type FingerprintVisitor struct{}

func (f *FingerprintVisitor) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	if v, ok := n.(*driver.ValueExpr); ok {
		v.SetValue([]byte("?"))
	}
	return n, false
}

func (f *FingerprintVisitor) Leave(n ast.Node) (node ast.Node, ok bool) {
	return n, true
}

func processSQL(mqp *model.PooledMysqlQueryPiece, sql []byte) *model.PooledMysqlQueryPiece {
	p := parser.New()
	stmt, err := p.ParseOneStmt(string(sql), "", "")
	if err != nil {
		fmt.Println("解析错误:" + err.Error())
		return mqp
	}
	stmt.Accept(&FingerprintVisitor{})

	buf := new(bytes.Buffer)
	restoreCtx := format.NewRestoreCtx(format.RestoreKeyWordUppercase|format.RestoreNameBackQuotes, buf)
	err = stmt.Restore(restoreCtx)
	if nil != err {
		fmt.Println("解析错误:" + err.Error())
		return mqp
	}
	fmt.Println(buf.String())
	querySQL := hack.String(buf)
	mqp.QuerySQLFinger = &querySQL
	return mqp

}
