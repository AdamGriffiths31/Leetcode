public class Solution
{
    public int MaxArea(int[] height)
    {
        int left = 0;
        int right = height.Count() - 1;
        int maxWater = 0;
        while (left < right)
        {
            int leftValue = height[left];
            int rightValue = height[right];
            int currentWater = Math.Min(height[left], height[right]) * (right - left);
            maxWater = Math.Max(maxWater, currentWater);
            if (height[left] > height[right])
            {
                right--;
            }
            else
            {
                left++;
            }
        }
        return maxWater;
    }
}