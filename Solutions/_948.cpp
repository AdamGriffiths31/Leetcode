#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int bagOfTokensScore(vector<int> &tokens, int power)
{
    int score = 0;
    int result = 0;
    int leftPointer = 0;
    int rightPointer = tokens.size() - 1;
    sort(tokens.begin(), tokens.end());

    while(leftPointer <= rightPointer)
    {
        if (power >= tokens[leftPointer])
        {
            power -= tokens[leftPointer];
            leftPointer++;
            score++;
            result = max(result, score);
        }
        else if (score > 0)
        {
            power += tokens[rightPointer];
            rightPointer--;
            score--;
        }
        else{
            return result;
        }
    }
    return result;
}

int main()
{
    vector<int> tokens = {100, 200, 300, 400};
    int power = 200;
    int res = bagOfTokensScore(tokens, power);
    cout << res << endl;
    return 0;
}