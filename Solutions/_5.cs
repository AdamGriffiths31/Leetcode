public class Solution
{
    public string LongestPalindrome(string s)
    {
        int start = 0;
        int end = 0;
        int len = 0;
        int maxLen = 0;
        if (s == null || s.Length <= 1) return s;
        for (int i = 0; i < s.Length; i++)
        {
            var palindrome1 = Helper(s, i, i);
            var palindrome2 = Helper(s, i, i + 1);
            len = Math.Max(palindrome1, palindrome2);
            if (len > end - start)
            {
                start = i - (len - 1) / 2;
                end = i + len / 2;
                maxLen = len;
            }
        }
        var x = s.Substring(start, maxLen);
        return s.Substring(start, maxLen);
    }
    public int Helper(string s, int left, int right)
    {
        while (left >= 0 && right < s.Length && s[left] == s[right])
        {
            left--;
            right++;
        }
        return right - left - 1;
    }
}