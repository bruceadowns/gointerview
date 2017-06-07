#include <stdio.h>

#define MAXLEN 1024

int mygetline(char line[], int max);
void copy(char src[], char dest[]);

int main() {
  char *line = NULL;
  size_t linecap = 0;
  ssize_t linelen;
  while ((linelen = getline(&line, &linecap, stdin)) > 0)
    fwrite(line, linelen, 1, stdout);

  /*
  int currCount = 0;
  int longestCount = 0;

  char currLine[MAXLEN];
  char longestLine[MAXLEN];

  while( (currCount = mygetline(currLine, MAXLEN)) > 0) {
    if (currCount > longestCount) {
      longestCount = currCount;
      copy(currLine, longestLine);
    }
  }
  */

  /*
  for (char c = getchar(); c != EOF; c = getchar()) {
    if (c == '\n') {
      if (currCount > longestCount) {
        longestCount = currCount;
        copy(currLine, longestLine);
      }

      currCount = 0;
      currLine[0] = '\0';
    } else {
      if (currCount < MAXLEN) {
        currLine[currCount++] = c;
      }
    }
  }
  */

  /*
  printf("%d\n", longestCount);
  printf("%s\n", longestLine);
  */
}

int mygetline(char line[], int max) {
  int i = 0;
  char c;

  for (c = getchar(); c != EOF && c != '\n'; c = getchar()) {
    line[i]=c;
    ++i;
  }
  if (c == '\n') {
    line[i]=c;
    ++i;
  }

  return i;
}

void copy(char src[], char dest[]) {
  for (int i = 0; src[i] != '\0'; i++ ) {
    dest[i] = src[i];
  }
}
