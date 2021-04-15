package context_strengthen

import (
    `context`
    `fmt`
    `reflect`
    `testing`
    `time`
)

func TestContext_ManyGoRoutineForOneCancel(t *testing.T){
    waitCtx, cancelFunc := context.WithCancel(context.Background())
    go func() {
        ticker := time.Tick(1 * time.Second)
        exit := false
        for{
            if exit {
                break
            }
            select {
            case <-waitCtx.Done():
                fmt.Println("1")
                exit = true
                break
            case <-ticker:
                fmt.Println("1-0")
                time.Sleep(1 * time.Second)
            }
        }
        fmt.Println("1-exit")
    }()

    go func() {
        ticker := time.Tick(1 * time.Second)
        exit := false
        for  {
            if exit {
                break
            }
            select {
            case <-waitCtx.Done():
                fmt.Println("2")
                exit = true
                break
            case <-ticker:
                fmt.Println("2-0")
                time.Sleep(1 * time.Second)
            }
        }
        fmt.Println("2-exit")

    }()

    time.Sleep(3 * time.Second)
    cancelFunc()
    time.Sleep(20 * time.Second)
}

func TestContext_TestCloseChannel(t *testing.T) {
    channel := make(chan float32, 2)
    channel <- 1
    close(channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
    fmt.Println(<-channel)
}

func TestContext_MultiContext(t *testing.T) {
    baseCtx := context.Background()
    ctx1F := context.WithValue(baseCtx, "1F", "2")
    ctx2F := context.WithValue(ctx1F, "2F", "3")
    ctx1FB := context.WithValue(ctx2F, "1F", "1")
    fmt.Println(ctx1F.Value("1F"))
    fmt.Println(ctx1F.Value("2F"))
    fmt.Println(ctx2F.Value("1F"))
    fmt.Println(ctx2F.Value("2F"))
    fmt.Println(ctx1FB.Value("1F"))
    fmt.Println(ctx1FB.Value("2F"))
}

func TestContext_GetValueFromParent(t *testing.T){
    valueCtx := context.WithValue(context.Background(), "A", "B")
    cancelCtx, _ := context.WithCancel(valueCtx)
    fmt.Println(cancelCtx.Value("A"))
}

type myContext struct {
    context.Context
}

func TestMyContext_TestImplement(t *testing.T) {
    myCtx := &myContext{context.Background()}
    typeForMyContext := reflect.TypeOf(myCtx)
    for i := 1; i <typeForMyContext.NumMethod();i++ {
        fmt.Println(typeForMyContext.Method(i))
    }
    myCtx.Err()
    myCtx.Value("1")
    myCtx.Done()
    myCtx.Deadline()
}

func TestValueContext_TestUnImplementMethod(t *testing.T) {
    valueCtx := context.WithValue(context.Background(), "1", "2")
    valueCtx.Err()
    fmt.Println(valueCtx.Deadline())
}