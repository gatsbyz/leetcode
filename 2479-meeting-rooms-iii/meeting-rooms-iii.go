func mostBooked(n int, meetings [][]int) int {
    m := len(meetings)
    // rooms holds end time of ongoing meeting (if any) in i room
    rooms := make([]int, n) 
    // meetingsCount holds total number of held meetings in each room
    meetingsCount := make([]int, n) 
    // sort meetings by start time
    slices.SortFunc(meetings, func (i, j []int) int {
        return i[0] - j[0]
    })
    // loop through meetings
    for i := 0; i < m; i++ {
        minEnd := 10000000000
        roomNum := -1
        start := meetings[i][0]
        // loop through rooms to find available one with lowest number
        for j := 0; j < n; j++ {
            rooms[j] = max(rooms[j], start)
            if rooms[j] < minEnd {
                minEnd = rooms[j]
                roomNum = j
            }
        }
        // set new end time for a room and increase number of held
        // meetings in this room
        rooms[roomNum] += meetings[i][1] - meetings[i][0]
        meetingsCount[roomNum]++
    }
    maxMeetings := -1
    roomNum := -1
    // find room with most held meetings
    for i := 0; i < n; i++ {
        if meetingsCount[i] > maxMeetings {
            maxMeetings = meetingsCount[i]
            roomNum = i
        }
    }
    return roomNum
}