public class Solution
{
    public bool UniqueOccurrences(int[] arr)
    {
        Hashtable hash = new Hashtable();
        HashSet<int> hashset = new HashSet<int>();

        for (int i = 0; i < arr.Length; i++)
        {

            hash[arr[i]] = 1 + (int)((hash.ContainsKey(arr[i])) ? hash[arr[i]] : 0);
        }

        foreach (DictionaryEntry item in hash)
        {
            if (!hashset.Add((int)item.Value))
            {
                return false;
            }
        }

        return true;
    }
}