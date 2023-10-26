#include <iostream>

using namespace std;

bool isAnagram(string s, string t) {
    if (s.length() != t.length()){
        return false;
    }

    int count[26] ={0};

    for(char c:s){
        count[c-'a']++;
    }
    for(char c:t){
        count[c-'a']--;
    }

    for(int x:count){
        if(x != 0){
            return false;
        }
    }
    return true;
}

int main(){
    bool result = isAnagram("rat", "car");
    cout << "result = " << result << "\n";

    result = isAnagram("anagram", "nagaram");
    cout << "result = " << result << "\n";
    return 0;
}