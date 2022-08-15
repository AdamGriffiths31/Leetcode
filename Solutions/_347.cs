public class Solution
{
    public int[] TopKFrequent(int[] nums, int k)
    {
        var bucket = new List<int>[nums.Length + 1];
        Dictionary<int, int> map = new Dictionary<int, int>();

        foreach (var item in nums)
        {
            {
                map[item] = 1 + (int)((map.ContainsKey(item)) ? map[item] : 0);
            }
        }

        foreach (var item in map)
        {
            bucket[item.Value] = bucket[item.Value] ?? new List<int>();
            bucket[item.Value].Add(item.Key);
        }

        int[] result = new int[k];
        for (int i = bucket.Length - 1, j = 0; i > 0; i--)
        {
            if (bucket[i] != null)
            {
                for (int n = 0; n < bucket[i].Count && j < k; n++, j++)
                {
                    result[j] = bucket[i][n];
                }
            }
        }

        return result;
    }
}