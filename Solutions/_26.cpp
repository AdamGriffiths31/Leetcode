#include <iostream>
#include <vector>

using namespace std;

int RemoveDuplicates(vector<int> &nums) {
  int n = nums.size();
  int j = 0;
  for (int i = 1; i < n; i++) {
    if (nums[i] != nums[j]) {
      j++;
      nums[j] = nums[i];
    }
  }
  return j + 1;
}

int main() {
  vector<int> nums = {1, 1, 2};
  int len = RemoveDuplicates(nums);
  cout << len << endl;
  nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
  len = RemoveDuplicates(nums);
  cout << len << endl;
  return 0;
}
