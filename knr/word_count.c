#include <stdio.h>

enum boolean { NO, YES };

int main() {

  int wc = 0;
  enum boolean inword = NO;

  char c = getchar();
  while ( c != EOF ) {
    if (c == ' ' || c == '\n') {
      inword = NO;
    } else if ( inword == NO ) {
      inword = YES;
      ++wc;
    }

    c = getchar();
  }

  printf("%d\n", wc);

}
