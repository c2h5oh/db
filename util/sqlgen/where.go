package sqlgen

import (
	"strings"
)

type (
	Or    []interface{}
	And   []interface{}
	Where []interface{}
)

type conds struct {
	Conds string
}

func (self Or) String() string {
	return groupCondition(self, mustParse(Layout.ClauseOperator, Layout.OrKeyword))
}

func (self And) String() string {
	return groupCondition(self, mustParse(Layout.ClauseOperator, Layout.AndKeyword))
}

func (self Where) String() string {
	grouped := groupCondition(self, mustParse(Layout.ClauseOperator, Layout.AndKeyword))
	if grouped != "" {
		return mustParse(Layout.WhereLayout, conds{grouped})
	}
	return ""
}

func groupCondition(terms []interface{}, joinKeyword string) string {
	l := len(terms)

	chunks := make([]string, 0, l)

	if l > 0 {
		var i int
		for i = 0; i < l; i++ {
			switch v := terms[i].(type) {
			case ColumnValue:
				chunks = append(chunks, v.String())
			case Or:
				chunks = append(chunks, v.String())
			case And:
				chunks = append(chunks, v.String())
			case Raw:
				chunks = append(chunks, v.String())
			}
		}
	}

	if len(chunks) > 0 {
		return mustParse(Layout.ClauseGroup, strings.Join(chunks, joinKeyword))
	}

	return ""
}
