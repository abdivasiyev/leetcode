pub struct Solution;

impl Solution {
    pub fn first_missing_positive(mut nums: Vec<i32>) -> i32 {
        let mut i: usize = 0;
        let n = nums.len();

        while i < n {
            let j = nums[i] - 1;
            if j >= 0 && j < n as i32 && i != j as usize && nums[i] != nums[j as usize] {
                nums.swap(i, j as usize);
            } else {
                i += 1;
            }
        }

        for i in 0..n {
            if nums[i] != (i + 1) as i32 {
                return (i + 1) as i32;
            }
        }

        return n as i32 + 1;
    }
}

#[cfg(test)]
mod tests {
    use std::vec;

    use super::*;

    #[test]
    fn it_expects_solution() {
        struct TestCase {
            name: String,
            nums: Vec<i32>,
            expected: i32,
        }

        let test_cases: Vec<TestCase> = vec![
            TestCase {
                name: "Case #1".to_string(),
                nums: vec![1, 2, 0],
                expected: 3,
            },
            TestCase {
                name: "Case #2".to_string(),
                nums: vec![3, 4, -1, 1],
                expected: 2,
            },
            TestCase {
                name: "Case #3".to_string(),
                nums: vec![7, 8, 9, 11, 12],
                expected: 1,
            },
            TestCase {
                name: "Case #4".to_string(),
                nums: vec![1],
                expected: 2,
            },
            TestCase {
                name: "Case #5".to_string(),
                nums: vec![1, 1],
                expected: 2,
            },
        ];

        for tc in test_cases {
            let result = Solution::first_missing_positive(tc.nums);
            assert_eq!(result, tc.expected, "{}", tc.name);
        }
    }
}
