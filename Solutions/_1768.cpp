#include <iostream>
#include <string>

std::string solve(std::string word1, std::string word2){
    std::string result = "";
    int longest = word1.length() > word2.length() ? word1.length() : word2.length();
    for (int i = 0; i < longest; i++){
        if (i < word1.length()){
            result += word1[i];
        }
        if (i < word2.length()){
            result += word2[i];
        }
    }

    return result; 
}

int main(){
    std::string word1 = "abcd";
    std::string word2 = "pq";

    std::string result = solve(word1, word2);
    std::cout << result << "\n";
}