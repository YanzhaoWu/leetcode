/**
 * Definition for isBadVersion()
 * 
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 * isBadVersion = function(version) {
 *     ...
 * };
 */

/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function(isBadVersion) {
    /**
     * @param {integer} n Total versions
     * @return {integer} The first bad version
     */
    return function(n) {
        if (isBadVersion(n)) {
            var low = 1;
            var high = n;
            while (low < high) {
                var mid = Math.floor((low + high) / 2);
                if (isBadVersion(mid)) {
                    high = mid;
                }
                else {
                    low = mid + 1;
                }
            }
            return high;
        }
        else {
            return -1;
        }
    };
};