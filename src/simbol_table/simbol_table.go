package simboltable

import (
	"fmt"
	"strings"
	"sync"
)

var ST = NewSymbolTable()
var ST_KEY = "simbol_table_index"

type SymbolTable struct {
	mu     *sync.Mutex
	NextId int
	table  map[string]int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		mu:     &sync.Mutex{},
		NextId: 0,
		table:  make(map[string]int),
	}
}

func (st *SymbolTable) AddSymbol(symbol string) int {
	st.mu.Lock()
	defer st.mu.Unlock()

	if _, ok := st.table[symbol]; !ok {
		st.table[symbol] = st.NextId
		st.NextId++
	}

	return st.table[symbol]
}

func (st *SymbolTable) String() string {
	sb := strings.Builder{}
	sb.WriteString("{\n")
	for k, v := range st.table {
		sb.WriteString(fmt.Sprintf("\t\"%v\": \"%v\"\n", k, v))
	}
	sb.WriteString("}\n")
	return sb.String()
}
