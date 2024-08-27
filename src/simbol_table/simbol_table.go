package simboltable

import (
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
