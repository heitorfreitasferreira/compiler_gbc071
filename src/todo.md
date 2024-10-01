# Tarefas do projeto de compiladores

Insira o @githubusername no final de cada tarefa para atribuir a tarefa a você.

## Etapa 1

- [x] Definir a GLC com as estruturas @yurihnrq
- [x] Identificar tokens @yurihnrq
  - [x] tabela com nome + atributo (se aplicável) @heitorfreitasferreira
- [x] Definir regex para cada token @Eduardorib

## Etapa 2

- [x] Diagrama de transição @todos
  - [x] Individuais por token @yuurihnrq
  - [x] Unificar em não determinístico @Eduardorib
  - [x] Determinizar @Eduardorib
- [x] Implementacão dirigida por tabela @heitorfreitasferreira
  - [x] Devolver um token por vez @heitorfreitasferreira
  - [x] Retornar o tipo do token, valor (se aplicável) e posição (linha e coluna) @heitorfreitasferreira
  - [x] Preencher tabela de símbolos com identificadores e constantes @heitorfreitasferreira
  - [x] Tratar comentários e separadores @heitorfreitasferreira
  - [x] Emitor erro útil quando necessário @heitorfreitasferreira

## Etapa 3

- [x] Ajustes na GLC @yurihnrq
  - [x] Remover recursão a esquerda @yurihnrq
  - [x] Tratar ambiguidades (associatividade, precedência, fatoração, etc) @yurihnrq
- [x] Calcular o First @yurihnrq
- [ ] Calcular o Follow
- [ ] Construir os grafos sintáticos
- [ ] Implementar o analisador sintático baseada em descida recursiva @heitorfreitasferreira
- [ ] Construir a árvore sintática concreta @heitorfreitasferreira
  - [ ] Implementar estrutura de árvore genérica @heitorfreitasferreira
- [ ] Emitir erros úteis quando necessário @heitorfreitasferreira
