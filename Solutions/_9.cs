public class Solution {
    public bool IsPalindrome(int x) {
        
        if (x < 0) return false;
        if (x == 0) return true;
        
        var intergerChar = x.ToString().ToCharArray();
        var pointerRight = intergerChar.Length - 1;
        var pointerLeft = 0;
        
        while(pointerLeft < pointerRight)
        {
            if(intergerChar[pointerLeft] != intergerChar[pointerRight] )
            {
                return false;
            }
            
            pointerRight--;
            pointerLeft++;
        }
        
        return true; 
    }
}