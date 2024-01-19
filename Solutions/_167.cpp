#include <iostream>
#include <vector>

using namespace std;

vector<int> twoSum(vector<int> &numbers, int target) {
  int lPointer = 0;
  int rPointer = numbers.size() - 1;

  while (lPointer < rPointer) {
    int currentSum = numbers[lPointer] + numbers[rPointer];
    if (currentSum == target)
      return {lPointer + 1, rPointer + 1};
    else if (currentSum < target)
      lPointer++;
    else
      rPointer--;
  }
  return {};
}

int main() {
  vector<int> numbers = {2, 7, 11, 15};
  int target = 9;
  vector<int> result = twoSum(numbers, target);
  cout << result[0] << " " << result[1] << endl;
  return 0;
}
