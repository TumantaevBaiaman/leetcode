class Solution {
    public int findCount(int num){
        int count = 0;
        while (num > 0){
            num = num/10;
            count++;
        };
        return count;
    }
    public int findNumbers(int[] nums) {
        int res = 0;
        for (int i = 0; i < nums.length; i++){
            if (findCount(nums[i])%2==0){
                res++;
            }
        }
        return res;
    }
}