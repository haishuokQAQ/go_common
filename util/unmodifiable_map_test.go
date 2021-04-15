package util

import (
    `github.com/go-playground/assert/v2`
    `reflect`
    `testing`
)

func TestNewUnmodifiableMap(t *testing.T) {
    data := map[interface{}]interface{}{
        "test":"1",
        "test2":3,
    }
    um := NewUnmodifiableMap(data)
    um.Store("test",3)
    result, ok := um.Load("test")
    assert.Equal(t, ok, true)
    assert.Equal(t, reflect.TypeOf(result), reflect.TypeOf(""))
    assert.Equal(t, result, "1")
}
