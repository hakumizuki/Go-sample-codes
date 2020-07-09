// Producer, Consumer
// package main
// import(
// 	"sync"
// 	"fmt"
// )

// func producer(ch chan int, i int) {
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("Process...", i * 1000) // この処理が失敗するかもしれないようなものの時、deferにしておくことでいい感じになる
// 		}()
// 	}
// 	fmt.Println("channelがCloseされました")
// }

// // Main3 はmain()に実行しています
// func Main3() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	// Producer
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	// Consumer
// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	close(ch) // Wait()後にclose()することで、既定の数の受信をし終えたときconsumer()のforループを終了させることができる しないとfor後の処理を行うことができない
// 	fmt.Println("Done!")
// }


// // fan-out, fan-in
// package main
// import(
// 	"fmt"
// )

// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first <-chan int, second chan<- int) { // <-を書いて明示的に送受信を表しておくとコードが読みやすくなることがある
// 	defer close(second)
// 	for f := range first {
// 		second <- f * 2
// 	}
// }

// func multi4(second chan int, third chan int) {
// 	defer close(third)
// 	for s := range second {
// 		third <- s * 4
// 	}
// }

// // Main3 はmain()で実行 second
// func Main3() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }


// select, default, for, break
package main 
import(
	"fmt"
	"time"
	"sync"
)

// Main3 はmain()で実行
func Main3() {
	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// OuterLoop:
	// 	for {
	// 		select {
	// 		case <-tick:
	// 			fmt.Println("tick.")
	// 		case <-boom:
	// 			fmt.Println("Boom!")
	// 			break OuterLoop // ループ名を指定してbreakすれば良い
	// 			// ここにただのbreakを書いてもselectのcase <-boomがbreakされるだけなので、forループにより復活する
	// 			// return だとその後の処理ができなくなる
	// 		default: // どのケースにも当てはまらない場合
	// 			fmt.Println("    .")
	// 			time.Sleep(50 * time.Millisecond)
	// 		}
	// 		// ここならbreakできる
	// 	}
	// fmt.Println("##################")


	// sync.Mutex
	// c := make(map[string]int)のかわりに
	c := Counter{v: make(map[string]int)}
	go func() { // ☆ ☆同士で同一のmapにRWするとconcurrent errorが発生する -> sync.Mutexで解決してみる
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	go func() { // ☆
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.v, c.Value("key"))
}

// Counter はMutex用のstruct
type Counter struct {
	v map[string]int
	mux sync.Mutex // LockとUnlockを使いたい物をstructでMutexとまとめて管理するという認識
}

// Inc ...
func (c *Counter) Inc(key string) { // Increment
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

// Value ...
func (c *Counter) Value(key string) int { // 読み込み用
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}