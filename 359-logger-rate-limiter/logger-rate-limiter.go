type Logger struct {
    logs map[string]int
}


func Constructor() Logger {
    return Logger{
        map[string]int{},
    }
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    if _, exists := this.logs[message]; !exists {
        this.logs[message] = timestamp
        return true
    }
    if this.logs[message] + 10 <= timestamp {
        this.logs[message] = timestamp
        return true
    }
    return false
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */