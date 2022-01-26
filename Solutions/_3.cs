public class Solution
{
    public int LengthOfLongestSubstring(string s)
    {
        if (s.Length == 1)
        {
            return 1;
        }
        Hashtable hashtable = new Hashtable();
        int max = 0;
        int start = 0;
        for (int i = 0; i < s.Length; i++)
        {
            if (hashtable.ContainsKey(s[i]))
            {
                start = Math.Max(start, (int)hashtable[s[i]] + 1);
            }
            max = Math.Max(max, i - start + 1);
            hashtable[s[i]] = i;
        }
        return max;
    }
}