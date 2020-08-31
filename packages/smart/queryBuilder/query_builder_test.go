/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package queryBuilder

import (
	"fmt"
	"testing"

	"github.com/IBAX-io/go-ibax/packages/types"

// whereF="[id]"
// whereV="[-6752330173818123413]"

type TestKeyTableChecker struct {
	Val bool
}

func (tc TestKeyTableChecker) IsKeyTable(tableName string) bool {
	return tc.Val
}
func TestSqlFields(t *testing.T) {
	qb := SQLQueryBuilder{
		Entry:        log.WithFields(log.Fields{"mod": "test"}),
		Table:        "1_keys",
		Fields:       []string{"+amount"},
		FieldValues:  []interface{}{2912910000000},
		Where:        types.LoadMap(map[string]interface{}{`id`: `-6752330173818123413`}),
		KeyTableChkr: TestKeyTableChecker{true},
	}

	fields, err := qb.GetSQLSelectFieldsExpr()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(fields)
}