class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        ptr1 = 0
        ptr2 = 0
        maxLength = 0
        while ptr2 < len(s) and ptr1 < len(s):
            if ptr1 >= ptr2:
                ptr2 += 1
                tmp = ptr2 - ptr1
                maxLength = max(maxLength,tmp)
                continue
            else:
                sameIndex = s[ptr1 : ptr2].find(s[ptr2])
                if sameIndex >= 0:
                    tmp = ptr2 - ptr1
                    ptr1 = ptr1 + sameIndex + 1
                else:
                    ptr2 += 1
                    tmp = ptr2 - ptr1
                maxLength = max(maxLength,tmp)
        return maxLength

