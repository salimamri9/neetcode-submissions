class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix or len(matrix) == 0:
            return False

        columns = len(matrix[0])
        rows = len(matrix)
        high = rows * columns - 1
        low = 0

        while low <= high:
            mid = (low + high) // 2
            i = mid % columns      
            j = mid // columns 
            if target == matrix[j][i]:
                return True
            elif target > matrix[j][i]:
                low = mid + 1
            else:
                high = mid - 1

        return False