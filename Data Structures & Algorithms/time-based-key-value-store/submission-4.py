class TimeMap:

    def __init__(self):
       self.timeMap = {}
        

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.timeMap:
            self.timeMap[key] = []
        self.timeMap[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.timeMap:
            return ""
        
        values = self.timeMap[key]

        l = 0
        r = len(values) - 1
        
        while l <= r :
            mid = (l + r) // 2
            if values[mid][0] == timestamp:
                return values[mid][1]
            elif values[mid][0] > timestamp:
                r = mid - 1
            else:
                l = mid + 1 
        
        if values[r][0] > timestamp:
            return ""
        else:
            return values[r][1]