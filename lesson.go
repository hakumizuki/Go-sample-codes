// ※とにかくGolangのdocをみよう

// パッケージ名 mainは必須
package main

// 使うパッケージをimportする
import (
	"fmt"
	"os/user"
	"time"
	"strings"
	"strconv"
	"os"
	"log"
	"io"
)

/* 定数(宣言した時点でコンパイラに解釈はされているが、実行されていないので、型宣言しないし、intもoverflowしない。大して変数はoverflowする) 定数宣言の上にはコメントがあるといいらしい(おそらくパブリックのみ)*/
const (
	Username = "dick"
	Music    = "rock"
)

// 最終的に実行されるmain()関数
func main() {
	fmt.Println("Hello, World!", time.Now())
	fmt.Println(user.Current())

	// 変数宣言（import同様に（）でまとめて宣言可能） ※Goでは下のようにコードを揃えるのがいいとされている
	var (
		i        int     = 1
		str      string  = "fuck"
		f64      float64 = 1.22
		t, f     bool    = true, false
		emptyStr string
		u8       uint8   = 255
	)

	// メイン関数内でのみ、short variable decralationが使える
	xi := 1
	//値の変更
	xi = 45

	//float32を指定したい場合はvarで宣言する
	xf64 := 1.36
	fafaffy := 2.2223

	xs := "suck"
	xt, xf := true, false

	fmt.Println(i, str, f64, t, f, emptyStr)
	fmt.Println(xi, xf64, xs, xt, xf)

	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", fafaffy)

	fmt.Println(Username, Music)

	fmt.Printf("Type=%T Value=%v\n", u8, u8)

	// この書き方だとアスキーコードが表示される(72)
	fmt.Println("Hello World!"[0])

	// string()でキャストすることで表示できる(H)
	fmt.Println(string("Hello World!"[0]))

	// Goで文字列の文字を変更したい時 strings pkgのReplaceを使う 第四引数は変換したい個数 また、コピーして表示しているので、元のstrは変更されていない
	str = "Hello"
	fmt.Println(strings.Replace(str, "l", "r", 1))
	fmt.Println(strings.Replace(str, "l", "r", 2))
	fmt.Println(strings.Replace(str, "l", "r", 0))
	fmt.Println(str)

	// 文字列に含まれているかどうかの検出は、Containsを使う
	fmt.Println(strings.Contains(str, "llo"))
	fmt.Println(strings.Contains(str, ""))
	fmt.Println(strings.Contains(str, " "))


	// リテラル(``で囲むのがいい)
	fmt.Println("\"")
	fmt.Println(
		`マジで勘弁
マジ勘弁
本当勘弁
Ben Ben Ben☆
"著：間次 菅勉"`)


	// 論理演算子は同じ && || !


	// 型変換
	var xst int = 1
	xnd := float64(xst)
	fmt.Printf("%T %v %f\n", xnd, xnd, xnd)

	// stringをintにしたい場合 Goではキャストできないため、strconv pkgのAtoi() を使う※Ascii to integer
	stringToCast := "08"
	i, err := strconv.Atoi(stringToCast) // err を _ にしてエラーハンドリングをしないこともできる
	if err != nil {
		fmt.Println("Error converting")
	} else {
		fmt.Printf("%T %v\n", i, i)
	}


	// 配列 Goの配列はサイズが変更できない->appendなども使えない
	var a [2]int
	a[0] = 100
	a[1] = 200

	var b [2]int = [2]int{100, 200}
	fmt.Println(a, b)


	//slice いわゆる配列
	var c []int = []int{300, 400}
	c = append(c, 500)
	fmt.Println(c)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2:4]) // 2:4は2から4という意味(:2は0から2、2:は2から最後まで、:のみは初めから最後までの意) ※インデックスの数え方として、値の前にインデックスがあるとして考える

	// 入れ子スライス 最後にコンマがほしいらしい
	nnn := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(nnn)

	// スライスのmakeとcapacity varで宣言する時との違いは、空配列になるかnil 配列になるか
	m := make([]int, 3, 5) //引数一つ(ex; 3 -> 3, 3)
	fmt.Printf("len=%d cap=%d val=%v\n", len(m), cap(m), m)

	// 演習
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
	}
	fmt.Println(c)


	// map(辞書型)
	mm := map[string]int{"apple": 100, "banana": 200} 
	fmt.Println(mm)
	fmt.Println(mm["apple"])
	mm["new"] = 300
	fmt.Println(mm)

	result, isExist := mm["banana"]
	result2, isExist2 := mm["nothing"]
	fmt.Println(result, isExist)
	fmt.Println(result2, isExist2) //二つを引数に取ると、値と存在の有無が返ってくる 二つ目の結果はなくても良い

	// mapの初期化 makeはメモリー上に空のmapが確保されるが、varで宣言した場合はnil mapになるため、追加したりできない スライスも同様
	// var mmnil map[string]int -> nil
	mm2 := make(map[string]int)
	mm2["pc"] = 5000
	fmt.Println(mm2)


	// バイト型
	byby := []byte{72, 73}
	fmt.Println(byby)
	fmt.Println(string(byby))

	// 関数、クロージャは省きます


	// 可変長引数
	foo(1, 2)
	foo(3, 4, 5, 6)


	// for continue break
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue // continueが実行されるとその後の処理はスキップされて次の周にいく
		}

		if i > 5 {
			fmt.Println("break")
			break // forループ終了
		}
		fmt.Println(i)
	}

	// forループの省略した書き方
	sum := 1
	for ; sum < 10; { // ;は両方とも省略して良い
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)


	// range indexとvalueをとってきてくれる
	// スライスの場合
	l := []string{"swift", "ruby", "go"}
	for i, v := range l { // indexを使いたくない時は、_ にする
		fmt.Println(i, v)
	}
	// mapの場合
	mappp := map[string]int{"iMac": 20, "macbook": 13}
	for k, v := range mappp { // valueが必要ない場合は書かなくて良いが、keyが必要ない場合は_にする
		fmt.Printf("Key=%s Val=%v\n", k, v)
	}


	// switch
	pcos := "mac"
	switch pcos { // switch宣言時にex-> switch os := thisFuncReturnsMac() {} としても良い
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows...")
	default: // 書かない場合は何も実行されない
		fmt.Println("Default")
	}

	currentTime := time.Now()
	fmt.Println(currentTime.Hour())
	switch { // このようにして条件をcaseにすることもできる
	case currentTime.Hour() < 12:
		fmt.Println("morning!!")
	case currentTime.Hour() >= 12:
		fmt.Println("afternoon!!")
	}


	// defer 遅延実行
	fooDefer()
	// defer を使ってファイルを開いて読み込む処理をしてみる
	file, _ := os.Open("./lesson.go")
	defer file.Close() // こうしておけばClose()し忘れることがなくなる

	data := make([]byte, 100, 100)
	file.Read(data)
	fmt.Println(string(data))


	// logging Print系は普通のlog Fatal系はその時点でプログラムが終了する
	log.Println("logging!!")
	log.Printf("%T %v", "test", "test")
	/* log.Fatalln("fatal error!!!!!!!!!!!!!!!!!") */

	// loggingの設定 この辺はよくわからん
	LoggingSettings("test.log")


	// エラーハンドリング ※ :=で宣言する時、複数ある宣言のうち一つがイニシャライズできていれば実行できる イニシャライズされなかったものは上書きされる
	/* if err = os.Chdir("test"); err != nil {
		こんな感じにもif文を書ける
	} */

	// panic & recover
	save()
	fmt.Println("OK?")

	// 演習２ CLEARED!!!
	// lll := []int{100, 200, 4, 6, 5, 7, 78, 3, 43}
	// var num int
	// for i, v := range lll {
	// 	if i == 0 {
	// 		num = v
	// 	} else {
	// 		if v < num {
	// 			num = v
	// 		}
	// 	}
	// }
	// fmt.Println("the smallest num =", num) // Success!!!
	// これをcontinueを使って書けば、ネストが少なくなってreadableになる ☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆
	lll := []int{100, 200, 4, 6, 5, 7, 78, 3, 43}
	var min int // 最小値を出したいという目的似合うようにnumではなくminと名付けるべき
	for i, v := range lll {

		if i == 0 {
			min = v
			continue // ☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆
		}

		if min >= v {
			min = v
		}

	}


	// ポインタ
	fmt.Println("ここからポインタ=ここからポインタ=ここからポインタ=ここからポインタ")

	var nnnn int = 100
	fmt.Println(nnnn)

	fmt.Println(&nnnn) // &はメモリーのアドレスであることを示す

	var pppp *int = &nnnn // *---はポインタ型であることを示す
	fmt.Println(pppp)
	fmt.Println(*pppp) // *をアドレスの前につけることで、そのアドレスのメモリーにある値を出す

	// dereference する
	nnder := 100
	one(&nnder)
	fmt.Println(nnder, "この数値はdereferenceされた！")


	// varとnew, makeの違い
	/* var integer *int は、メモリーに領域を確保していないため、&integerでもアドレスは返らず、nilとなる
	そのため値に対してアクセスして変化させることはできない -> new()で初期化して、メモリー領域を確保する */

	// newとmakeの違い new()はポインタ型を作るときに使われる
	makeSlice := make([]int, 5, 5)
	makeMap   := make(map[string]int)
	fmt.Printf("%T\n", makeSlice)
	fmt.Printf("%T\n", makeMap)
	var newPointerInt *int = new(int) // *intはintが入るアドレスという解釈で、この場合はnewPointerIntにその新しいアドレスを代入している
	newPointerInt2 := new(int)
	fmt.Printf("%T\n", newPointerInt)
	fmt.Printf("%T\n", newPointerInt2)


	// struct
	fmt.Println("ここからSTRUCT=ここからSTRUCT=ここからSTRUCT=ここからSTRUCT=ここからSTRUCT=ここからSTRUCT")
	vertex := Vertex{} // 値を入れずに初期化すると、デフォルト値が入る
	fmt.Println(vertex)
	vertex2 := Vertex{1, 2, "vertex2", "VERTEX2"}
	fmt.Println(vertex2)
	vertex3 := Vertex{X: 1, S2: "VERTEX3"}
	fmt.Println(vertex3)

	var vertex4 Vertex // nilにならない
	fmt.Println(vertex4, "VERTEX4")

	// 下の5, 6, 7は同じことを意味するが、6がぱっと見でわかりやすく、よく使われる
	var vertex5 *Vertex = new(Vertex)
	fmt.Println(vertex5, "*VERTEX5")
	fmt.Printf("%T VERTEX5\n", vertex5)
	vertex6 := &Vertex{}
	fmt.Println(vertex6, "*VERTEX6")
	fmt.Printf("%T VERTEX6\n", vertex6)
	vertex7 := new(Vertex)
	fmt.Println(vertex7, "*VERTEX7")
	fmt.Printf("%T VERTEX7\n", vertex7)

	// ※TIPs!!!!! mapやスライスはnew, makeを使い、structは&などで宣言するのが今はいい感じ


	// structの値を変更する
	vertexV := &Vertex{1, 2, "3", "4"}
	changeVertex(vertexV)
	fmt.Println(*vertexV)
	vertexF := Vertex{5, 6, "7", "8"}
	changeVertex2(vertexF)
	fmt.Println(vertexF)
} // End of main()




// 可変長引数 [1, 2, 3]...で展開できるのはrubyと同じ
func foo(params ...int) {
	fmt.Println(len(params), params)
}

// deferの処理は、関数が実行され終わった後に実行される
func fooDefer() {
	defer fmt.Println("defer foo")
	defer fmt.Println("stacking defer last") // 複数のdeferは後から書かれたものから実行される(stacking defer)
	defer fmt.Println("stacking defer second")
	defer fmt.Println("stacking defer first")
	fmt.Println("Hello foo")
}

// LoggingSettings logをファイルに書き込む
func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

// panic golangでは推奨されていません -> しっかりとエラーハンドリングで対処する
func thirdPartyConnectDB() {
	panic("unabled to connect database!")
}
func save() {
	defer func() {
		rec := recover() // recover()にpanicの例外結果が入り、強制終了させないようにしている
		fmt.Println(rec)
	}()
	thirdPartyConnectDB()
}

// dereference (ポインタ)
func one(x *int) {
	*x = 1
}

// Vertex ... STRUCT
type Vertex struct {
	X int
	Y int
	S1, S2 string
}

func changeVertex(v *Vertex) { // アドレスを取得してそのメモリーの値を変えることで値が変わる
	v.X = 1000 // vは本来は*vだが、structではこうかける
}
func changeVertex2(v Vertex) { // この場合、vはコピーされたものになるため、実体を変化させることはできない
	v.X = 1000
}

// Area ポインタレシーバーと値レシーバー
func (v Vertex) Area() int { // 値レシーバー このように書くと、v.Area()としてメソッドを呼び出せる 関数でArea(v)よりもわかりやすい
	return v.X * v.Y
}
// Scale ...
func (v *Vertex) Scale(i int) { // ポインタレシーバー
	v.X *= i
	v.Y *= i
}

// New コンストラクタ
func New(x, y int, s1, s2 string) *Vertex { // 自作New以外にもpkg.Newでstructを返すものが多い
	return &Vertex{x, y, s1, s2}
}

// Embedded structを継承させるとき
/* 
	type Vertex3D struct { 
		Vertex // こうするだけでいい
		z int
	}

	// Embeddedのコンストラクタ
		func New(x, y, z int, s1, s2 string) *Vertex3D {
			return &Vertex3D{Vertex{x, y, s1, s2}, z} // ネストする
		}
*/

// non-structのメソッド
type myInt int // Swiftのextensionみたいな感じ

func (i myInt) Twice() int {
	return int(i * 2)
}

// ダックタイピング interface
/*
	type Human interface { // そのinterfaceに準拠しているだけで、定義された関数を持つと判断できることと、その関数を使い忘れないことがメリットか
		MyName() string // この関数を持っていなければならない
	}

	type Person struct {
		Name string
	}

	func (p *Person) MyName() string {
		p.Name = "Mr." + p.Name // 本体も書き換わる
		return p.Name
	}

	func main() {
		var person1 Human = &Person{"Jobs"}
		fmt.Println("I'm", person1.MyName())
	}
*/


// タイプアサーション + switch type文
/*
	func do(i interface{}) { // interface{}はどんな型も許容する
		switch v := i.(type) {
		case int:
			fmt.Println(v *= 2)
		case string:
			fmt.Println(v + "!")
		default: 
			fmt.Printf("I don't know %T", v)
		}

		ii := i.(int) // タイプアサーション 空のinterfaceをint型にしている
		ii *= 2
		fmt.Println(ii)
	}

	func main() {
		do(10)
		do("Mac")
		do(true)
	}
*/

// Stringer fmt.Println()の表示内容を変えたいときに使う
/*
	type X string
	func (x X) String() string {
		return fmt.Sprintf("I'm %v", x) // stringを返す type X がAとBの値を持つものだとしても、Aのみを使うことでBを表示させないようにすることができる
	}
	var x X = "XXX"
	fmt.Println(x)
*/

// カスタムエラー
/*
	type UserNotFound struct { // 1.自分なりのエラーをstructにする
		Username string
	}
	func (e *UserNotFound) Error() string { // 2.*UserNotFoundに対してError()をつくる もしアドレスではなく値を引数とすると比較できなくなり、同一のエラーとして検知してしまう、複数箇所で同じエラーが起こるときに問題が発生する
		return fmt.Sprintf("User not found: %v\n", e.Username) // structのエラーの値を返す
	}

	func someFunc() error {
		if ok := false; ok { // okならnilを返す
			return nil
		} 
		return &UserNotFound{Username: "Jobs"} // okでないのならつくったエラーを返す
	}

	func main() {
		if err1 := someFunc(); err1 != nil {
			fmt.Println(err1)
		}
	}
*/
