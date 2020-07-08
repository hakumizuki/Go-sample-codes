package main
import(
	"sync"
	"fmt"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg.Done()
			fmt.Println("Process...", i * 1000) // この処理が失敗するかもしれないようなものの時、deferにしておくことでいい感じになる
		}()
	}
	fmt.Println("channelがCloseされました")
}

// Main3 はmain()に実行しています
func Main3() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch) // Wait()後にclose()することで、既定の数の受信をし終えたときconsumer()のforループを終了させることができる しないとfor後の処理を行うことができない
	fmt.Println("Done!")
}