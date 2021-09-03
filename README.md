# scheduler-go
GoLang version simple scheduler by onion model

### Install

> go get -u github.com/hubvue/scheduler


### Create Handler
```go
type handler struct {}

func (h handler) Start(data interface{}) interface{} {
    num := data.(int)
    return num + num
}

func (h handler) End(data interface{}) interface{} {
    num := data.(int)
    return num * num
}
```

### Add
Add a handler to the schedule

```go
scheduler.Add(handler{})
```

### Run
execute all handlers in the schedule
```go
res := scheduler.Run(2)
fmt.Println(res) // 16
```

### Remove
remove handler from scheduler
```go
h := handler{}
scheduler.Add(h)
scheduler.Remove(h)
```

### Clear
clear all handlers in the air conditioner
```go
scheduler.Clear()
```

### New Install
create schedule instance
```go
schedule := scheduler.New()
```
