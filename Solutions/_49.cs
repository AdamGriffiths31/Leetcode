public class Solution
{
    public IList<IList<string>> GroupAnagrams(string[] strs)
    {
        Dictionary<string, IList<string>> hashMap = new Dictionary<string, IList<string>>();

        foreach (var s in strs)
        {
            var count = Enumerable.Repeat<int>(0, 26).ToArray();

            foreach (var c in s)
            {
                count[(int)c - (int)'a'] += 1;
            }

            var key = string.Join(",", count);
            if (!hashMap.ContainsKey(key))
            {
                hashMap.Add(key, new List<string> { s });
            }
            else
            {
                hashMap[key].Add(s);
            }

        }

        return hashMap.Values.ToList();
    }
}