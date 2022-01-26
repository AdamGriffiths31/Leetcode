public int[] TwoSum(int[] nums, int target)
{
    Hashtable hash = new Hashtable();
    List<int> result = new List<int>();
    for (int i = 0; i < nums.Count(); i++)
    {
        var b = target - nums[i];
        if (hash.ContainsKey(b))
        {
            result.Add(Array.IndexOf(nums, b));
            result.Add(i);
            return result.ToArray();
        }
        hash[nums[i]] = i;
    }
    return null;
}
