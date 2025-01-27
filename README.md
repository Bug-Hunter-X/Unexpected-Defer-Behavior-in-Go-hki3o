# Unexpected Defer Behavior in Go

This example demonstrates a common pitfall with Go's `defer` statement. The `defer` keyword schedules a function call to be executed after the surrounding function returns. However, the arguments to the deferred function are evaluated *when the defer statement is encountered*, not when the function actually runs.

In this specific case, the `defer fmt.Println(x)` statement is executed at the beginning of `main()`, before `x` is updated to 10. As a result, the output of the program is 0 followed by 10. This might be unexpected for developers unfamiliar with this behavior.

The solution shows how to address this issue by using a closure to capture the current value of x.