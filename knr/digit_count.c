#include <stdio.h>
#include <ctype.h>

int main() {
  int dc = 0;

  for (char c = getchar(); c != EOF; c = getchar()) {
    dc += isdigit(c) ? 1 : 0;
  }

  printf("%d\n", dc);
}
