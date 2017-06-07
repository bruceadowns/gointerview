#include <stdio.h>

int main(int argc, char ** argv) {
  /*
  char c = getchar();
  while (c != EOF) {
    putchar(c);
    c = getchar();
  }
  */

  /*
  char c;
  while ((c = getchar()) != EOF) {
    putchar(c);
  }
  */

  // putchar(c);

  if (argc == 1) {
    for (char c = getchar(); c != EOF; c = getchar())
      fputc(c, stdout);
  } else {
    FILE * fp = fopen(*++argv, "r");
    if (fp == NULL) {
      fprintf(stderr, "Error opening file: %s", *argv);
      return 1;
    } else {
      for (char c = fgetc(fp); c != EOF; c = fgetc(fp))
        fputc(c, stdout);
      fclose(fp);
    }
  }
}
