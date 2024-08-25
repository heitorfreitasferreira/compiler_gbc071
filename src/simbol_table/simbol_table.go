package simboltable

import (
	"sync"

	"github.com/heitorfreitasferreira/compiler/types"
)

type SymbolTable struct {
	mu     *sync.Mutex
	NextId int
	table  types.BinTree
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		mu:     &sync.Mutex{},
		NextId: 0,
		table:  types.BinTree{},
	}
}

func (st *SymbolTable) AddSymbol(symbol string) {
	st.mu.Lock()
	defer st.mu.Unlock()
	inserted := st.table.Add(st.NextId)
	if !inserted {
		return // Já existe na tabela
	}
	st.NextId++
}
