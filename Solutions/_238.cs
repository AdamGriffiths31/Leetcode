public class Solution
{
    public int[] ProductExceptSelf(int[] nums)
    {

        var result = Enumerable.Repeat<int>(1, nums.Length).ToArray();
        var prefix = 1;
        var postfix = 1;

        for (int i = 0; i < nums.Length; i++)
        {
            result[i] = prefix;
            prefix *= nums[i];
        }
        for (int i = nums.Length - 1; i >= 0; i--)
        {
            result[i] *= postfix;
            postfix *= nums[i];
        }

        return result;
    }
}