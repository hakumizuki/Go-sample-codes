package main
import(
	"fmt"
	// "time"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done() // こういうときにdeferが使えるんや
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}
func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

// channel 並列処理での値の受け渡し
	func goroutine1(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // channelに送信
	}
	func goroutine2(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
			// ここにc <- sumとすると、毎回の結果が送信される for i := range channel {println(i)} ですべて出力できるが、closeを忘れずに
		}
		c <- sum // channelに送信
	}

// Main2 is in main()
func Main2() {
	// Goroutine 並列処理 goとつけるだけで並列になる goの並列処理よりも早く他の処理が完了してしまった場合はその時点で処理が終了する
	// WaitGroup
	var wg sync.WaitGroup
	wg.Add(1) // 待つDone()の回数
	go goroutine("Goroutine!", &wg) // 
	normal("normal")
	wg.Wait()

	// channel
	ssArr := []int{1, 2, 3, 4, 5}
	ch := make(chan int) // 15, 20 ... とQueueのようにどんどんデータが受信できる(unbuffer 領域制限なし)
	go goroutine1(ssArr, ch)
	go goroutine2(ssArr, ch)
	receivedChannel1 := <-ch // channelから受信する=値を取り出す データが送信されるまで次の処理に進まず待機し、channelに送信されたらそれを受け取る syncのWait()で待つ必要がない
	fmt.Println(receivedChannel1, "12345")
	receivedChannel2 := <-ch
	fmt.Println(receivedChannel2)

	ch2 := make(chan int, 2) // buffered channel
	ch2 <- 100
	fmt.Println(len(ch2))
	ch2 <- 200
	fmt.Println(len(ch2)) // 2個まで入る
	receivedChannel3 := <-ch2 // channelの古いものからひとつ取り出す
	fmt.Println(receivedChannel3)
	ch2 <- 300 // また入れることができる
	fmt.Println(len(ch2))
	close(ch2) // channelに終わりを告げる forループするときにrangeが噛み合わなくなるため

	for c := range ch2 {
		fmt.Println(c, "for rooping")
	}
}