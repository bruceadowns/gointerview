#include <stdio.h>

#define STEP 20
#define START 300
#define END 0

int main(int argc, char ** argv) {
  printf("%d arguments\n", argc-1);
  for (int i = 0; i < argc ; ++i) {
    if (i > 0) printf(" ");
    printf("%s", argv[i]);
  }
  printf("\n");

  for (int f = START; f >= END; f -= STEP) {
    int c = (5.0 / 9.0) * (f - 32.0);
    printf("%df => %dc\n", f, c);
  }
}
