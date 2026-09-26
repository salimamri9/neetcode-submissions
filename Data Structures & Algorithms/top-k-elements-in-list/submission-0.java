

class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> freq = new HashMap<>();
for (int num : nums) {
    freq.put(num, freq.getOrDefault(num, 0) + 1);
}

// Max frequency can't exceed nums.length
List<Integer>[] buckets = new List[nums.length + 1];
for (int i = 0; i <= nums.length; i++) {
    buckets[i] = new ArrayList<>();
}

// Put numbers in buckets by frequency
for (Map.Entry<Integer, Integer> entry : freq.entrySet()) {
    int number = entry.getKey();
    int count = entry.getValue();
    buckets[count].add(number);
}

// Collect top k from buckets
List<Integer> result = new ArrayList<>();
for (int i = buckets.length - 1; i >= 0 && result.size() < k; i--) {
    for (int num : buckets[i]) {
        result.add(num);
        if (result.size() == k) break;
    }
}

int[] finalResult = new int[result.size()];
for(int i=0;i<result.size();i++){
    finalResult[i]=result.get(i);

    }
    return finalResult;
}
}
