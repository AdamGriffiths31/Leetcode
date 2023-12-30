#include <iostream> 
#include <vector>
#include <string> 
#include <unordered_map>
#include <algorithm>

std::vector<std::vector<std::string>> groupAnagrams(std::vector<std::string>& strs) {
    std::unordered_map<std::string, std::vector<std::string>> result;

    for (const std::string& s : strs) {
        std::string key = s;
        std::sort(key.begin(), key.end()); 
        result[key].push_back(s);
    }

    std::vector<std::vector<std::string>> output;
    for (const auto& pair : result) {
        output.push_back(pair.second);
    }

    return output;
}

int main(){
    std::vector<std::string> input = {"eat", "tea", "tan", "ate", "nat", "bat"}; 
    std::vector<std::vector<std::string>> output = groupAnagrams(input);
    for (int i = 0; i < output.size(); i++){
        for (int j = 0; j < output[i].size(); j++){
            std::cout << output[i][j] << " ";
        }
        std::cout << std::endl;
    } 
    return 0;
}