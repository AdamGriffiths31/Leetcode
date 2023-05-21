/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if (x < 0 ) return false; 
    if(x == 0) return true;

    let num = x;
    let pointerLeft = 0;
    let pointerRight = x.toString().length - 1;

    while(pointerLeft < pointerRight){
        if(num.toString()[pointerLeft] != num.toString()[pointerRight]) return false;
        pointerLeft++;
        pointerRight--;
    } 
    return true; 
};