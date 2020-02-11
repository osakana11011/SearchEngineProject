package valueobject_test

import (
    "testing"
    "strconv"

    "search_engine_project/search_engine/src/domain/valueobject"
)

func TestToken(t *testing.T) {
    testPatterns := []struct {
        givenID int64
        givenStr string
        expectedID int64
        expectedStr string
    } {
        {1, "星", 1, "星"},
        {10, "蟹", 10, "蟹"},
    }

    for i, data := range testPatterns {
        t.Run("test"+strconv.Itoa(i), func(t *testing.T) {
            res := valueobject.NewToken(data.givenID, data.givenStr)

            // IDの一致検証
            if res.ID() != data.expectedID {
                t.Fatalf("[IDの不一致] expected: %d, actual: %d", data.expectedID, res.ID())
            }

            // Strの一致検証
            if res.Str() != data.expectedStr {
                t.Fatalf("Strの不一致 expected: %s, actual: %s", data.expectedStr, res.Str())
            }
        })
    }
}
