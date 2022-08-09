public class Solution
{
    public bool IsAnagram(string s, string t)
    {

        if (s.Length != t.Length)
        {
            return false;
        }
        Hashtable hashS = new Hashtable();
        Hashtable hashT = new Hashtable();

        for (int i = 0; i < s.Length; i++)
        {
            hashS[s.Substring(i, 1)] = 1 + (int)(hashS.ContainsKey(s.Substring(i, 1)) ? hashS[s.Substring(i, 1)] : 0);
            hashT[t.Substring(i, 1)] = 1 + (int)(hashT.ContainsKey(t.Substring(i, 1)) ? hashT[t.Substring(i, 1)] : 0);
        }

        foreach (DictionaryEntry hash in hashS)
        {
            if (!hashT.ContainsKey(hash.Key))
            {
                return false;
            }

            if ((int)hash.Value != (int)hashT[hash.Key])
            {
                return false;
            }
        }

        return true;
    }
}