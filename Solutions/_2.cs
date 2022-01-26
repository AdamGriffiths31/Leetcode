public class Solution {
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
            ListNode start = null;
            ListNode previous = null;
            int carry = 0;
            int sum = 0;

            while (l1 != null | l2 != null || carry==1)
            {
                if (l1 == null) l1 = new ListNode() { val = 0 };
                if (l2 == null) l2 = new ListNode() { val = 0 };

                sum = l1.val + l2.val + carry;
                if (sum > 9)
                {
                    sum -= 10;
                    carry = 1;
                }
                else
                {
                    carry = 0;
                }
                var current = new ListNode()
                {
                    val = sum                    
                };
                if (previous != null)
                {
                    previous.next = current;
                }
                else
                {
                    start = current;
                }
                previous = current;
                l1 = l1.next;
                l2 = l2.next;
            }
            return start;
    }
}
