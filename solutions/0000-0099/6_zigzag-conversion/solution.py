class Solution(object):
    def convert(self, s, numRows):
        if numRows == 1:
            return s

        zigzag_list = ['' for _ in range(numRows)]
        currentRow = 0
        goingDown = False

        for symbol in s:
            zigzag_list[currentRow] += symbol
            if currentRow == 0 or currentRow == numRows - 1:
                goingDown = not goingDown
            currentRow += 1 if goingDown else -1

        return ''.join(zigzag_list)