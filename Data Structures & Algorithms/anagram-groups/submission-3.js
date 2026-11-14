class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        /* 
            {
                act: [act, cat]
                aht: [hat]
                opst: [stop, pots, tops]

            }

            OR

            Give act = 0, aht = 1, opst = 2, and keep pushing the strings to the belonging group.
        */

        let res = {};

        for (let str of strs) {
            const sortedStr = this.sortHelper(str);
            if (!res[sortedStr]) {
                res[sortedStr] = [];
            }
            res[sortedStr].push(str);
        }

        return Object.values(res);
    }

    sortHelper(str) {
        return str.split("").sort().join("");
    }
}
