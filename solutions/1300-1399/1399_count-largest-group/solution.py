class Solution(object):
    def countLargestGroup(self, n):
        """
        :type n: int
        :rtype: int
        """
        def digit_sum(num):
            total = 0
            while num > 0:
                total += num % 10
                num = num // 10
            return total

        groups = {}
        for i in range(1, n + 1):
            s = digit_sum(i)
            if s not in groups:
                groups[s] = 1
            else:
                groups[s] += 1

        max_size = max(groups.values())
        return sum(1 for count in groups.values() if count == max_size)