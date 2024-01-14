#include <iostream>
#include <vector>

using namespace std;

class MinStack
{
private:
    vector<int> stack;
    vector<int> minStack;

public:
    MinStack()
    {
    }

    void push(int val)
    {
        stack.push_back(val);
        if (minStack.empty())
        {
            minStack.push_back(val);
            return;
        }
        val = min(val, minStack.back());
        minStack.push_back(val);
    }

    void pop()
    {
        stack.pop_back();
        minStack.pop_back();
    }

    int top()
    {
        return stack.back();
    }

    int getMin()
    {
        return minStack.back();
    }
};

int main()
{

    MinStack *obj = new MinStack();
    obj->push(1);
    obj->pop();
    int param_3 = obj->top();
    int param_4 = obj->getMin();

    std::cout << param_3 << "\n";
    std::cout << param_4 << "\n";
}