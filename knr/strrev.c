#include <stdio.h>
#include <string.h>

void strrev(char * s);

int main() {
  char s[] = "Hello World";
  printf("s: %s\n", s);
  strrev(s);
  printf("s: %s\n", s);
}

void strrev(char * s) {
  int i, j;
  for (i=0, j=strlen(s)-1; i<j; ++i, --j ) {
    char temp = s[i];
    s[i] = s[j];
    s[j] = temp;
  }
}
