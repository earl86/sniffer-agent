package mysql

import (
	"bytes"
	"fmt"

	"sniffer-agent/model"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/format"
	driver "github.com/pingcap/tidb/types/parser_driver"
	"github.com/pingcap/tidb/util/hack"
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
	stmt, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		fmt.Println("解析错误:" + err.Error())
		return nil
	}
	stmt.Accept(&FingerprintVisitor{})

	buf := new(bytes.Buffer)
	restoreCtx := format.NewRestoreCtx(format.RestoreKeyWordUppercase|format.RestoreNameBackQuotes, buf)
	err = stmt.Restore(restoreCtx)
	if nil != err {
		fmt.Println("解析错误:" + err.Error())
		return nil
	}
	//fmt.Println(buf.String())
	mqp.QuerySQLFinger = &hack.String(buf.Bytes())
	return mqp

}
