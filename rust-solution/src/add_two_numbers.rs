use std::{borrow::Borrow, ops::Deref};

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    fn add(&mut self, val: i32) {
        let mut cur = self;

        while let Some(ref mut next) = cur.next {
            cur = next;
        }
        cur.next = Some(Box::new(Self { val, next: None }))
    }
}

fn createList(values: Vec<i32>) -> Option<Box<ListNode>> {
    if values.len() == 0 {
        return None;
    }

    let mut root = Box::new(ListNode::new(values[0]));

    for i in 1..values.len() {
        root.add(values[i]);
    }

    return Some(root);
}

pub struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut carry = 0;
        let mut result = Some(Box::new(ListNode::new(0)));
        let mut head = result.as_mut();

        let (mut l1, mut l2) = (l1.as_ref(), l2.as_ref());

        while l1.is_some() || l2.is_some() {
            let mut sum = 0;

            if let Some(node) = l1 {
                sum += node.val;
                l1 = node.next.as_ref();
            }

            if let Some(node) = l2 {
                sum += node.val;
                l2 = node.next.as_ref();
            }

            sum += carry;
            carry = sum / 10;
            sum %= 10;

            head.as_mut().unwrap().next = Some(Box::new(ListNode::new(sum)));
            head = head.unwrap().next.as_mut();
        }

        if carry != 0 {
            head.as_mut().unwrap().next = Some(Box::new(ListNode::new(carry)));
        }

        result.unwrap().next
    }
}

#[cfg(test)]
mod tests {
    use crate::add_two_numbers::{createList, ListNode, Solution};

    #[test]
    fn it_expects_solution() {
        struct TestCase {
            name: String,
            l1: Option<Box<ListNode>>,
            l2: Option<Box<ListNode>>,
            expected: Option<Box<ListNode>>,
        }

        let test_cases: Vec<TestCase> = vec![
            TestCase {
                name: "Case #1".to_string(),
                l1: createList(vec![2, 4, 3]),
                l2: createList(vec![5, 6, 4]),
                expected: createList(vec![7, 0, 8]),
            },
            TestCase {
                name: "Case #2".to_string(),
                l1: createList(vec![9, 9, 9]),
                l2: createList(vec![1]),
                expected: createList(vec![0, 0, 0, 1]),
            },
        ];

        for tc in test_cases {
            let result = Solution::add_two_numbers(tc.l1, tc.l2);
            assert_eq!(result, tc.expected, "{}", tc.name);
        }
    }
}
