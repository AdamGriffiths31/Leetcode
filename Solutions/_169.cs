public class Solution
{
    public int MajorityElement(int[] nums)
    {
        Hashtable hash = new Hashtable();
        for (int i = 0; i < nums.Length; i++)
        {
            hash[nums[i]] = 1 + (int)(hash.ContainsKey(nums[i]) ? hash[nums[i]] : 0);
        }

        int frequent = 0;
        DictionaryEntry result = new DictionaryEntry();
        foreach (DictionaryEntry item in hash)
        {
            if (frequent < (int)item.Value)
            {
                frequent = (int)item.Value;
                result = item;
            }
        }

        return (int)result.Key;
    }
}