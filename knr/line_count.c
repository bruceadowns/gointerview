#include <stdio.h>

enum bool { NO, YES };

int main() {
  int lc=0;

  /*
  for (char c = getchar(); c != EOF; c = getchar()) {
    if (c == '\n') ++lc;
  }
  */

  enum bool il = NO;

  for (char c = getchar(); c != EOF; c = getchar()) {
    if (il == NO) {
      il = YES;
      ++lc;
    }

    if (c == '\n') il = NO;
  }

  printf("%d\n", lc);
}
