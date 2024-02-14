#include <string>
#include <iostream>
#include <unordered_map>

using namespace std;

bool canConstruct(string ransomNote, string magazine)
{
    unordered_map<char, int> map;

    for (char c : ransomNote)
    {
        map[c]++;
    }
    for (char c : magazine)
    {
        map[c]--;
    }

    for (auto val : map)
    {
        if (val.second > 0)
        {
            return false;
        }
    }
    return true;
}

int main()
{
    string ransomNote = "aa", magazine = "ab";
    cout << canConstruct(ransomNote, magazine) << endl;
    return 0;
}