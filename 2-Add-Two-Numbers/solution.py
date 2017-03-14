# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

def toLinkedList(num):
    tmp = result = ListNode(num % 10)
    tmpNext = ListNode(0)
    num /= 10
    while num > 0:
        tmpNext.val = num % 10
        tmp.next = tmpNext
        tmp = tmpNext
        num /= 10
        tmpNext = ListNode(0)
    tmp.next = None
    return result
    
def linkedListToNum(list):
    order = 1
    result = list.val * order
    order *= 10
    while list.next != None:
        list = list.next
        result += list.val * order
        order *= 10
    return result


class Solution(object):
    
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        return toLinkedList(linkedListToNum(l1)+linkedListToNum(l2))