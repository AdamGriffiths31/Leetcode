#include <vector>
#include <string>
#include <iostream>

using namespace std;

int evalRPN(vector<string> &tokens)
{
    vector<int> stack;
    for (int i = 0; i < tokens.size(); i++)
    {
        string token = tokens[i];
        if (token == "+")
        {
            int a = stack.back();
            stack.pop_back();
            int b = stack.back();
            stack.pop_back();
            stack.push_back(a + b);
        }
        else if (token == "-")
        {
            int a = stack.back();
            stack.pop_back();
            int b = stack.back();
            stack.pop_back();
            stack.push_back(b - a);
        }
        else if (token == "*")
        {
            int a = stack.back();
            stack.pop_back();
            int b = stack.back();
            stack.pop_back();
            stack.push_back(a * b);
        }
        else if (token == "/")
        {
            int a = stack.back();
            stack.pop_back();
            int b = stack.back();
            stack.pop_back();
            stack.push_back(b / a);
        }
        else
        {
            int value = stoi(token);
            stack.push_back(value);
        }
    }
    return stack.back();
}

int main()
{
    auto input = vector<string>{"2", "1", "+", "3", "*"};
    auto output = evalRPN(input);
    cout << output << endl;
}