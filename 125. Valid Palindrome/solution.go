func isPalindrome(s string) bool {
    pattern := `[a-z0-9]`
    re := regexp.MustCompile(pattern)
    s = strings.ToLower(s)
    matches := re.FindAllString(s, -1)
    size := len(matches)

    for i := range size/2 {
        if condition := matches[i] != matches[size - 1 - i]; condition {
            return false
        }
    }
    return true
}
