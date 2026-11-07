class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let keys = {} // {3: 4, 4: 3}

        for(let i = 0; i<nums.length; i++) {
            let comp = target - nums[i]
            if(keys[comp] !== undefined) {
                return [keys[comp], i]
            }
            keys[nums[i]] = i
        }
        return []
    }
}
