#include <stdio.h>

int main() {
  // 0 => 48
  // 9 => 57

  int digits[10];
  for (int i = 0; i < 10; ++i ) digits[i] = 0;

  for (char c = getchar(); c != EOF; c = getchar()) {
    if (c >= '0' && c <= '9') ++digits[c - '0'];
  }

  for (int i = 0; i < 10; ++i ) {
    printf("%d %d\n", i, digits[i]);
  }
}
