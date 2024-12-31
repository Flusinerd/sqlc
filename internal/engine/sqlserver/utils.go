package sqlserver

import (
	"github.com/sqlc-dev/sqlc/internal/engine/sqlserver/parser"
	"github.com/sqlc-dev/sqlc/internal/sql/ast"
)

type tableNamer interface {
	Table_name() parser.ITable_nameContext
	Schema_name() parser.IId_Context
}

func parseTableName(c parser.ITable_nameContext) *ast.TableName {
	ids := c.AllId_()
	if len(ids) > 1 {
		name := ast.TableName{
			Name:   identifier(ids[1].GetText()),
			Schema: identifier(ids[0].GetText()),
		}
		return &name
	} else {
		name := ast.TableName{
			Name:   identifier(ids[0].GetText()),
			Schema: "dbo",
		}

		return &name
	}
}

func hasNotNullConstraint(col parser.IColumn_definitionContext) bool {
	for _, def := range col.AllColumn_definition_element() {
		constraint := def.Column_constraint()
		if constraint.Null_notnull() != nil && constraint.Null_notnull().NOT() != nil {
			return true
		}

		if constraint.PRIMARY() != nil && constraint.KEY() != nil {
			return true
		}
	}

	return false
}
