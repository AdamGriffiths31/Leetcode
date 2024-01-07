#include <vector>
#include <map>
#include <iostream>

using namespace std;

bool isValidSudoku(vector<vector<char>> &board)
{
    bool sqaures[9][9] = {false};
    bool rows[9][9] = {false};
    bool cols[9][9] = {false};

    for (size_t r = 0; r < board.size(); r++)
    {
        for (size_t c = 0; c < board[r].size(); c++)
        {
            if (board[r][c] == '.')
            {
                continue;
            }

            int idx = board[r][c] - '0' - 1;
            int area = (r / 3) * 3 + (c / 3);

            if (sqaures[area][idx] || rows[r][idx] || cols[c][idx])
            {
                return false;
            }
            sqaures[area][idx] = true;
            rows[r][idx] = true;
            cols[c][idx] = true;
        }
    }

    return true;
}

int main()
{
    auto input = vector<vector<char>>{
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '.', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '.', '.', '.', '.', '.', '1'},
        {'7', '.', '.', '.', '.', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '.', '9'}};
    auto result = isValidSudoku(input);
    cout << result << "\n";
}