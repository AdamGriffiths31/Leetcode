public class Solution
{
    public bool ContainsDuplicate(int[] nums)
    {
        Hashtable hash = new Hashtable();

        for (int i = 0; i < nums.Length; i++)
        {
            if (hash.ContainsKey(nums[i]))
            {
                return true;
            }

            hash[nums[i]] = nums[i];
        }

        return false;

        //if (nums.Distinct().Count() == nums.Length)
        //{
        //    return false;
        //}
        //else
        //{
        //    return true;
        //}
    }
}