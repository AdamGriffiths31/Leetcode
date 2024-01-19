#include <iostream>
#include <string>

using namespace std;

bool isAlphaNumeric(char c) {
  if (c >= 'a' && c <= 'z')
    return true;
  if (c >= 'A' && c <= 'Z')
    return true;
  if (c >= '0' && c <= '9')
    return true;
  return false;
}

bool isPalindrome(string s) {
  int lPointer = 0;
  int rPointer = s.length() - 1;

  while (lPointer < rPointer) {
    if (!isAlphaNumeric(s[lPointer])) {
      lPointer++;
      continue;
    }
    if (!isAlphaNumeric(s[rPointer])) {
      rPointer--;
      continue;
    }
    if (tolower(s[lPointer]) != tolower(s[rPointer]))
      return false;
    lPointer++;
    rPointer--;
  }
  return true;
}

int main() {
  string input = "A man, a plan, a canal: Panama";
  cout << isPalindrome(input) << endl;
  input = "race a car";
  cout << isPalindrome(input) << endl;
  input = " ";
  cout << isPalindrome(input) << endl;
  return 0;
}
