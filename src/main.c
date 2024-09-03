#include <stdio.h>
#include <stdlib.h>

void print_read_char(char c, int next_state) {
  switch (c) {
  case '\n':
    printf("Read char: \\n ");
    break;
  case '\t':
    printf("Read char: \\t ");
    break;
  case ' ':
    printf("Read char: \\s ");
    break;
  default:
    printf("Read char: %c ", c);
  }
  printf("- Moved to: %d\n", next_state);
}

typedef struct {
  char read_char;
  int next_state;
  int is_final;
} Transition;

typedef struct {
  int id;
  int is_final;
  Transition *transitions;
} State;

typedef struct {
  int curr_state;
  int input_pos;
  State *states;
} Lexer;

Lexer *create_lexer(State *states, int initial_state) {
  Lexer *lexer = (Lexer *)malloc(sizeof(Lexer));
  if (lexer == NULL) {
    fprintf(stderr, "Error: Memory allocation failed\n");
    return NULL;
  }
  lexer->states = states;
  lexer->curr_state = initial_state;
  lexer->input_pos = 0;

  return lexer;
}

void move(Lexer *lexer, char read_char) {
  // Transitions for current state
  Transition *possible_transitions =
      lexer->states[lexer->curr_state].transitions;

  int i = 0;
  while (possible_transitions[i].read_char != '\0') {
    if (possible_transitions[i].read_char == read_char) {
      lexer->curr_state = possible_transitions[i].next_state;

      print_read_char(read_char, lexer->curr_state);

      return;
    }
    i++;
  }

  // Verify if there is a transition for '\0'
  lexer->curr_state = possible_transitions[i].next_state;
  print_read_char(read_char, lexer->curr_state);
}

void handle_lookahead(Lexer *lexer) { lexer->input_pos--; }

void act(Lexer *lexer) {
  if (lexer->states[lexer->curr_state].is_final) {
    if (lexer->curr_state == 4) {
      printf("DIV\n");
      handle_lookahead(lexer);
    }
    if (lexer->curr_state == 7) {
      printf("COMMENT MULTI LINE\n");
    }
    if (lexer->curr_state == 6) {
      printf("COMMENT SINGLE LINE\n");
    }
    if (lexer->curr_state == 8) {
      printf("ERROR\n");
    }

  } else {
    printf("Rejected\n");
  }
}

int main() {
  // Example usage
  int error_state = 8;
  Transition state0[] = {{'/', 1, 0}, {'\0', error_state}};
  Transition state1[] = {{'/', 3}, {'*', 2}, {'\n', 4}, {'\0', 4}};
  Transition state2[] = {{'/', 2}, {'*', 5}, {'\n', 2}, {'\0', 2}};
  Transition state3[] = {{'\n', 6}, {'\0', 3}};
  Transition state4[] = {{'\0', error_state}};
  Transition state5[] = {{'/', 7}, {'\0', 2}};
  Transition state6[] = {{'\0', error_state}};
  Transition state7[] = {{'\0', error_state}};
  Transition state8[] = {{'\0', error_state}};
  State *states = (State[]){{0, 0, state0}, {1, 0, state1}, {2, 0, state2},
                            {3, 0, state3}, {4, 1, state4}, {5, 0, state5},
                            {6, 1, state6}, {7, 1, state7}, {8, 1, state8}};

  Lexer *lexer = create_lexer(states, 0);
  if (lexer == NULL) {
    return 1;
  }

  // Read input from keyboard incuding \n
  char input[100];
  fgets(input, 100, stdin);

  while (input[lexer->input_pos] != '\0') {

    while (!lexer->states[lexer->curr_state].is_final) {
      move(lexer, input[lexer->input_pos++]);
    }

    act(lexer);
    // Go back to initial state
    lexer->curr_state = 0;
  }
  free(lexer);
  return 0;
}
