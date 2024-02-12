#include <iostream>
#include <string>
#include <unordered_map>

using namespace std;

bool isIsomorphic(string s, string t)
{
    unordered_map<char, char> mapST;
    unordered_map<char, char> mapTS;

    for (int i = 0; i < s.size(); i++)
    {
        if ((mapST.count(s[i]) > 0 && mapST[s[i]] != t[i]) || (mapTS.count(t[i]) > 0 && mapTS[t[i]] != s[i]))
        {
            return false;
        }

        mapST[s[i]] = t[i];
        mapTS[t[i]] = s[i];
    }

    return true;
}

int main()
{
    string s = "paper";
    string t = "title";

    bool result = isIsomorphic(s, t);
    cout << result << "\n";
}