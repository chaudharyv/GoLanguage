package main

//news, image, finances, sports
import "fmt"
import "time"

var timeouterror = false

func main() {
	news := make(chan string)
	image := make(chan string)
	sport := make(chan string)
	finance := make(chan string)
	go timeout()
	go financeSearch(finance)
	go newsSearch(news)
	go sportsSearch(sport)
	go imageSearch(image)
	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-news:
			fmt.Println("received", msg1)
			break
		case msg2 := <-image:
			fmt.Println("received", msg2)
			break
		case msg3 := <-sport:
			fmt.Println("received", msg3)
			break
		case msg4 := <-finance:
			fmt.Println("received", msg4)
			break
		}
		if timeouterror {
			break
		}
	}
	close(news)
	close(finance)
	close(sport)
	close(image)
}
func timeout() {
	breaktest := <-time.After(2 * time.Second)
	fmt.Println("timeout occured ", breaktest)
	timeouterror = true
}
func newsSearch(news chan<- string) {
	time.Sleep(7 * time.Second)
	news <- "news results 6 sec sleep"
}
func imageSearch(image chan<- string) {
	time.Sleep(6 * time.Second)
	image <- "image results 5 sec wait"
}
func sportsSearch(sports chan<- string) {
	time.Sleep(6 * time.Second)
	sports <- "sports results 6 sec wait"
}

func financeSearch(finance chan<- string) {
	time.Sleep(2 * time.Second)
	finance <- "finance results 2 sec wait"
}
