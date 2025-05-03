class Solution(object):
    def convertToTitle(self, columnNumber):
        result = ""
        while columnNumber > 0:
            columnNumber -= 1
            result = chr(65 + columnNumber % 26) + result
            columnNumber //= 26
        return result