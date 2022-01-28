public class Solution
{
    public int Reverse(int x)
    {
        long result = 0;
        long limit = 2147483648;
        long limit2 = -2147483648;
        while (x != 0)
        {
            var pop = x % 10;
            var temp = result * 10 + pop;
            result = temp;
            x = x / 10;
        }
        var r = (result >= limit || result <= limit2) ? 0 : (int)result;
        return (result >= limit || result <= limit2) ? 0 : (int)result;
    }
}