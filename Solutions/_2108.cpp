#include <iostream>
#include <vector>

using namespace std;

string firstPalindrome(vector<string> &words)
{
    for (string word : words)
    {
        int left = 0;
        int right = word.size() - 1;

        while (word[left] == word[right])
        {
            if (left >= right)
            {
                return word;
            }
            left++;
            right--;
        }
    }
    return "";
}

int main()
{
    vector<string> words = {"abc", "def", "ab", "cba", "aba"};
    cout << firstPalindrome(words) << endl;
    return 0;
}