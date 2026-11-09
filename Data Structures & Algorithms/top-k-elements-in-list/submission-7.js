class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        /**
         * value: occurence
         * {
         *  1: 1
         *  2: 2
         *  3: 3
         * }
         */

        const hash = {}

        for (let n of nums) {
            if(hash[n]) {
                hash[n] += 1
            } else {
                hash[n] = 1
            }
        }
        return Object.keys(hash).sort((a, b) => hash[b] - hash[a]).slice(0, k)
    }   
}
