package util

// UniqArray は配列中の要素が一意になるようにフィルタリングする。
func UniqArray(arr []string) []string {
    uniqArray := []string{}

    for _, v := range arr {
        if !contains(uniqArray, v) {
            uniqArray = append(uniqArray, v)
        }
    }

    return uniqArray
}

func contains(arr []string, v string) bool {
    for _, arrValue := range arr {
        if arrValue == v {
            return true
        }
    }
    return false
}
