#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int carFleet(int target, vector<int> &position, vector<int> &speed) {
  vector<double> fleet;
  vector<pair<int, int>> cars;
  for (int i = 0; i < position.size(); i++) {
    cars.push_back({position[i], speed[i]});
  }
  sort(cars.begin(), cars.end());

  for (int i = cars.size() - 1; i >= 0; i--) {
    double x = (double)(target - cars[i].first) / cars[i].second;
    fleet.push_back(x);
    if (fleet.size() >= 2) {
      double temp1 = fleet[fleet.size() - 1];
      double temp2 = fleet[fleet.size() - 2];

      if (temp1 <= temp2) {
        fleet.pop_back();
      }
    }
  }

  return fleet.size();
}

int main() {
  vector<int> position = {10, 8, 0, 5, 3};
  vector<int> speed = {2, 4, 1, 1, 3};
  int target = 12;
  cout << carFleet(target, position, speed) << endl;
  position = {6, 8};
  speed = {3, 2};
  target = 10;
  cout << carFleet(target, position, speed) << endl;
  return 0;
}
