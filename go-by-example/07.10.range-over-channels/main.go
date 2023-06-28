package main

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for item := range queue {
        println(item)
        // values are accessible even
        // after channel is closed
    }
}
