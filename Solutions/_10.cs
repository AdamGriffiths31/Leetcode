public class Solution
{
    public bool IsMatch(string s, string p)
    {
        Hashtable hash = new Hashtable();
        return BackTracking(0, 0, s, p, hash);
    }

    public bool BackTracking(int i, int j, string s, string p, Hashtable hash)
    {
        if (hash.ContainsKey($"i:{i}&j:{j}"))
        {
            return (hash[$"i:{i}&j:{j}"] as bool?) ?? false;
        }

        if (i >= s.Length && j >= p.Length)
        {
            return true;
        }

        if (j >= p.Length)
        {
            return false;
        }

        var match = i < s.Length && ((s[i] == p[j]) || p[j] == '.');

        if ((j + 1) < p.Length && p[j + 1] == '*')
        {
            hash[$"i:{i}&j:{j}"] = BackTracking(i, j + 2, s, p, hash) || match && BackTracking(i + 1, j, s, p, hash);
            return (hash[$"i:{i}&j:{j}"] as bool?) ?? false;
        }

        if (match)
        {
            hash[$"i:{i}&j:{j}"] = BackTracking(i + 1, j + 1, s, p, hash);
            return (hash[$"i:{i}&j:{j}"] as bool?) ?? false;
        }
        hash[$"i:{i}&j:{j}"] = false;
        return false;
    }
}