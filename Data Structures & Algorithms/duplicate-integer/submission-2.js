class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        const keys = {}
        for(let item of nums) {
            if(keys[item]) {
                return true
            } else {
                keys[item] = true
            }
        }
        return false
    }
}
