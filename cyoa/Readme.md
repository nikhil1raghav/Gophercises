# Choose your own adventure

### What does it do?
Render chapter of a story on a webpage, at the end give user an option to choose where to take story next.
Non-linear storyline. 



### Learnings

- You can use [json-to-go](https://mholt.github.io/json-to-go/) to quickly convert a json response or file to a struct for using in the code. Less gruntwork.
- Nesting a main.go in a directory named "X" will give the binary the name "X". 
- While reading a file (that is json), better to use `json.NewDecoder` than reading all data into a byte slice and then unmarshalling it. As decoder will directly work on `io.reader`.
- A map doesn't maintain the order of its keys. When you print content of map multiple times, they can be in different order.
- Functional options
```go
handler:=cyoa.NewHandler(story, cyoa.WithTemplate("someTemplate"), cyoa.WithTimeout(300))
```
- I like this better, and have implemented this one
```go
handler:=cyoa.NewHandler(story).WithTimeOut(300).WithTemplate("someTemplate")
```
