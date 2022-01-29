public class Solution
{
    public int MyAtoi(string s)
    {
        s = s.TrimStart();
        if (string.IsNullOrEmpty(s))
        {
            return 0;
        }
        var result = 0;
        bool isNegative = false;
        if (s[0] == '-')
        {
            isNegative = true;
        }
        else if (s[0] == '+')
        {
            isNegative = false;
        }
        else if (!char.IsDigit(s[0]))
        {
            return 0;
        }
        else
        {
            result = s[0] - '0';
        }
        for (int i = 1; i < s.Length; i++)
        {
            if (char.IsDigit(s[i]))
            {
                if (result > Int32.MaxValue / 10 || (result == Int32.MaxValue / 10 && s[i] - '0' > (Int32.MaxValue % 10)))
                {
                    if (isNegative)
                    {
                        return int.MinValue;
                    }
                    else
                    {
                        return int.MaxValue;
                    }
                }
                result = result * 10 + s[i] - '0';
            }
            else
            {
                break;
            }
        }
        return (isNegative) ? -result : result;
    }
}