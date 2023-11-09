/*
 * @lc app=leetcode id=41 lang=rust
 *
 * [41] First Missing Positive
 */

// @lc code=start
impl Solution {
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        let mut i = 0;
        let n = nums.len();

        while i < n {
            if nums[i] > 0 && nums[i] <= n as i32 && nums[i] != nums[(nums[i] - 1) as usize] {
                let mut j = (nums[i] - 1) as usize;
                nums.swap(i, j);
            } else {
                i += 1;
            }
        }

        for i in 0..n {
            if (nums[i] != (i + 1) as i32) {
                return i as i32 + 1;
            }
        }

        return n as i32 + 1;
    }
}
// @lc code=end
