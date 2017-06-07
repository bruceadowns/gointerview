#include <stdio.h>
#include <string.h>

int mystrlen(const char *);

int main() {
  char * str = "Hello World";
  const int len = mystrlen(str);

  printf("str: '%s' mystrlen: %d\n", str, len);
}

int mystrlen(const char * s) {
  int i = 0;
  for ( ; s[i] != '\0'; ++i ) ;
  return i;
}
