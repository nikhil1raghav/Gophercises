# What is it?

Quiet HN is a minimal hackernews clone. It has similar links to HN but no discussions.
Goal of this gophercise is to add concurrency and caching to the website.

Before doing any changes, this page loads in 11 seconds.

### Improvements
- Adding basic concurrency, not checking number of stories strictly. Got it down to 1.3 seconds.
- This can be achieved via blindly firing goroutines and using a waitgroup.

```go
for i:=0;i<numStories;i++{
    results=append(result, <-resultCh)
}
```

- Or using a channel to receive results of goroutines. But how to get exactly 30 stories on the page.
- Three Ask HN threads in top 30 were reducing the number of stories fetched to 27.
- Fetch some extra stories to offset this and select top 30.







### Learnings
- Check for race conditions before implementing cache `go run -race main.go`.
- adding cache
- Using mutex to remove race conditions
- Rotating cache to remove delays when cache expires or first time load
