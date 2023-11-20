/*
 * @lc app=leetcode id=41 lang=cpp
 *
 * [41] First Missing Positive
 */

#include <bits/stdc++.h>

using namespace std;

// @lc code=start
class Solution
{
public:
    int firstMissingPositive(vector<int> &nums)
    {
        int n = nums.size();
        int i = 0;

        while (i < n)
        {
            if (nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i] - 1])
            {
                swap(nums[i], nums[nums[i] - 1]);
            }
            else
            {
                i++;
            }
        }

        for (int i = 0; i < n; i++)
        {
            if (nums[i] != i + 1)
            {
                return i + 1;
            }
        }

        return n + 1;
    }
};
// @lc code=end
