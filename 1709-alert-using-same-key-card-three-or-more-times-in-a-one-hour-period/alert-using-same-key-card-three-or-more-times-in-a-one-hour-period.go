type Record struct {
    time int
    diff int
}

func alertNames(keyName []string, keyTime []string) []string {
    logs := make([][]string, 0)
    for i, name := range keyName {
        logs = append(logs, []string{ name, keyTime[i] }) 
    }
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })
    record := make(map[string]*Record)
    alert := make([]string, 0)
    for _, log := range logs {
        name := log[0]
        time := log[1]
        if _, ok := record[name] ; !ok {
            record[name] = &Record{
                time: timeToNumber(time),
                diff: 0,
            }
        } else {
            var diff int
            if timeToNumber(time) < record[name].time {
                diff = 24 * 60 - record[name].time + timeToNumber(time)
            } else {
                diff = timeToNumber(time) - record[name].time
            }
            fmt.Println(timeToNumber(time), record[name].time)
            fmt.Println(record[name].diff, diff)
            if record[name].diff + diff <= 60 {
                if record[name].diff > 0 {
                    if !slices.Contains(alert, name) {
                        alert = append(alert, name)
                    }
                }
            }
            record[name].time = timeToNumber(time)
            record[name].diff = diff
        }
    }
    slices.Sort(alert)
    return alert
}

// keyName = ["daniel","daniel","daniel","luis","luis","luis","luis"],
// keyTime = ["10:00","10:40","11:00", "13:00","15:00"]
              //    40       20      120      
              // 

func timeToNumber(time string) int {
    timeSplit := strings.Split(time, ":")
    hour, _ := strconv.Atoi(timeSplit[0])
    minute, _ := strconv.Atoi(timeSplit[1])
    return hour * 60 + minute
}