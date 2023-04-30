package easy

func twoSum(nums []int, target int) []int {
    for i, val := range nums {
        //fmt.Println(i, val) 
        for j := i+1; j < len(nums); j++{
            //fmt.Println("add idx i", i, "value", nums[i], "to idx j", j, "value", nums[j],": ", nums[j] + val)
            if val + nums[j] == target {
                return []int{i,j}
            }
        }   
    }
    return nil
}