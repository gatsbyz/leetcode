func intersection(nums1 []int, nums2 []int) []int {
    var count = map[int]bool{}
    var result = []int{}
    for _,num := range nums1{
        count[num]=true
    }
    for _, num := range nums2{
        if(count[num]==true){
            result=append(result,num)
            count[num]=false
        }
    } 
    return result    
}