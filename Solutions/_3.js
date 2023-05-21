/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    const map = new Map();
    let max = 0;
    let start = 0;
    for (let i = 0; i < s.length; i++) {
        const char = s[i];
        if (map.has(char)) {
            start = Math.max(start, map.get(char) + 1);
        }
        max = Math.max(max, i - start + 1);
        map.set(char, i);
    }
    return max;
};