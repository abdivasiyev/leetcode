/*
 * @lc app=leetcode id=16 lang=cpp
 *
 * [16] 3Sum Closest
 */

#include <bits/stdc++.h>

using namespace std;

// @lc code=start
class Solution
{
public:
    int threeSumClosest(vector<int> &nums, int target)
    {
        sort(nums.begin(), nums.end());

        int closestSum = 0;
        int minDiff = INT_MAX;

        for (int i = 0; i < nums.size() - 2; i++)
        {
            int start = i + 1;
            int end = nums.size() - 1;

            while (start < end)
            {
                int sum = nums[i] + nums[start] + nums[end];

                if (minDiff > abs(target - sum))
                {
                    minDiff = abs(target - sum);
                    closestSum = sum;
                }

                if (sum == target)
                {
                    return sum;
                }
                else if (sum < target)
                {
                    start++;
                }
                else
                {
                    end--;
                }
            }
        }

        return closestSum;
    }
};
// @lc code=end
