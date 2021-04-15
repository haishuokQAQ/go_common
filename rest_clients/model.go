package clients

type Meta struct {
    Code    int64  `json:"code"`
    Message string `json:"message"`
}

// BaseResp is the response entity
type BaseResp struct {
    Meta *Meta       `json:"meta"`
    Data interface{} `json:"data"`
}
