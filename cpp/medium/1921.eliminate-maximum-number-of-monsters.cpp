/*
 * @lc app=leetcode id=1921 lang=cpp
 *
 * [1921] Eliminate Maximum Number of Monsters
 */

#include <bits/stdc++.h>

using namespace std;

// @lc code=start
class Solution
{
public:
    int eliminateMaximum(vector<int> &dist, vector<int> &speed)
    {
        int eliminated = 0;
        vector<float> arrivals;

        for (int i = 0; i < dist.size(); i++)
        {
            arrivals.push_back((float)dist[i] / speed[i]);
        }

        sort(arrivals.begin(), arrivals.end());

        for (int i = 0; i < arrivals.size(); i++)
        {
            if (arrivals[i] <= i)
            {
                break;
            }
            eliminated++;
        }

        return eliminated;
    }
};
// @lc code=end
