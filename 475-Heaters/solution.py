class Solution(object):
    def findRadius(self, houses, heaters):
        """
        :type houses: List[int]
        :type heaters: List[int]
        :rtype: int
        """
        heaters.sort()
        result = 0
        for i in range(len(houses)):
            lower_index = self.lower_bound(heaters, houses[i])
            if lower_index < 0:
                tmp = abs(houses[i] - heaters[0])
            elif lower_index == len(heaters) - 1:
                tmp = abs(houses[i] - heaters[len(heaters) - 1])
            else:
                tmp = min(abs(houses[i] - heaters[lower_index]),abs(houses[i] - heaters[lower_index+1]))
            result = max(result, tmp)
        return result

    def lower_bound(self, heaters, target):
        low = 0
        high = len(heaters) - 1
        while low <= high:
            mid = (low + high) / 2
            mid_val = heaters[mid]
            if mid_val < target:
                low = mid + 1
            elif mid_val > target:
                high = mid - 1
            else:
                return mid
        return high
