* `Go` is using, what's called a `green thread`
* which is Abstraction of an actual thread, creating `go routines` which is managed by go runtinme and we are only interacting with these high level go routiunes
* cheaper and lightweight
* we can run 100s, 1000s and millions of `goroutines` without affecting the performance
* Go routines have concept of  channels to talk with one another