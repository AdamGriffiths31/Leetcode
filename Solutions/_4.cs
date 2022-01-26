public class Solution
{
    public double FindMedianSortedArrays(int[] nums1, int[] nums2)
    {
        List<int> results = new List<int>();
        int value = 0;
        bool looping = true;
        int aCounter = 0;
        int bCounter = 0;
        while (looping)
        {
            if (nums1.Length - 1 >= aCounter && nums2.Length - 1 >= bCounter)
            {
                value = (nums1[aCounter] < nums2[bCounter] ? nums1[aCounter] : nums2[bCounter]);
                if (nums1[aCounter] < nums2[bCounter])
                {
                    aCounter++;
                }
                else
                {
                    bCounter++;
                }
                results.Add(value);
            }
            else if (nums1.Length - 1 >= aCounter)
            {
                results.Add(nums1[aCounter]);
                aCounter++;
            }
            else if (nums2.Length - 1 >= bCounter)
            {
                results.Add(nums2[bCounter]);
                bCounter++;
            }
            else
            {
                looping = false;
            }
        }
        int middle = results.Count / 2;
        var x = results.Count % 2;
        double median = (results.Count % 2 != 0) ? results[middle] : (double)((double)results[middle] + (double)results[middle - 1]) / 2;
        return median;
    }
}