public class Solution
{
    public string Convert(string s, int numRows)
    {
        if (numRows == 1)
        {
            return s;
        }
        string result = "";
        for (int r = 0; r < numRows; r++)
        {
            var increment = 2 * (numRows - 1);
            for (int i = r; i < s.Length; i += increment)
            {
                result += s[i];
                if (r > 0 && r < numRows - 1 && i + increment - 2 * r < s.Length)
                {
                    result += s[i + increment - 2 * r];
                }
            }
        }
        return result;
    }
}