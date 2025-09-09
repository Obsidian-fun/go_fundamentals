package main

import (
    "fmt"
    "time"
)

func fetchUserDataOne(userId int) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        time.Sleep(80 * time.Millisecond)
        out <- "user data"
    }()
    return out
}

func fetchUserRecommendationsOne(userId int) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        time.Sleep(120 * time.Millisecond)
        out <- "user recommendations"
    }()
    return out
}

func fetchUsersLikesOne(userId int) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        time.Sleep(50 * time.Millisecond)
        out <- "user likes"
    }()
    return out
}

func main() {
    userId := 8
    futures := make(chan string)

    go func() {
        defer close(futures)
        futures <- <-fetchUserDataOne(userId)
        futures <- <-fetchUserRecommendationsOne(userId)
        futures <- <-fetchUsersLikesOne(userId)
    }()

    for future := range futures {
        fmt.Println(future)
    }
}
