func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"

		// In Go lang, a fallghrough basically forcefully execute the program to next condition even if it's not meet the condition. 
	// Go doesn't have this by default as C and C++ but if we use fallthrough, then it will execute the next case ignoring the state of condition.
    // all three of these cases will set creator to "A Steve"
    case "macOS":
        fallthrough
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve"

    default:
        creator = "Unknown"
    }
    return creator
}
