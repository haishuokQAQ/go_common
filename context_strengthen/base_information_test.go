package context_strengthen

import (
    `fmt`
    `testing`
)

func TestCreateBaseInformationContext(t *testing.T) {
    ctx := createBaseInfoContext(func() map[interface{}]interface{} {
        return map[interface{}]interface{}{
            "test":"1",
            "test2":3,
        }
    })
    fmt.Println(ForceGetBasementValue(ctx, "test"))
}
