public class Solution
{
    public bool IsValidSudoku(char[][] board)
    {
        HashSet<char> row = new HashSet<char>();
        List<HashSet<char>> cols = new List<HashSet<char>>();
        List<HashSet<char>> squares = new List<HashSet<char>>();

        for (int r = 0; r < 9; r++)
        {
            row = new HashSet<char>();

            for (int c = 0; c < 9; c++)
            {
                if (r == 0)
                {
                    squares.Add(new HashSet<char>());
                    cols.Add(new HashSet<char>());
                }

                if (board[r][c] == '.')
                {
                    continue;
                }

                if (!row.Add(board[r][c]) || !cols[c].Add(board[r][c]) || !squares[r / 3 * 3 + c / 3].Add(board[r][c]))
                {
                    return false;
                }
            }
        }

        return true;
    }
}