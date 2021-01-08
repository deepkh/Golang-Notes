# 目標

1. golang 的範例筆記

-------------------------------------------------

## 線上練習 go 
 -> [Golang Playground](https://play.golang.org/)

-------------------------------------------------

 ## 參考文獻 (References)
 - [官方Effective Go](https://golang.org/doc/effective_go.html)
 - [michaelchen golang 教學](https://michaelchen.tech/golang-programming/)
 - [Golang standard library](https://golang.org/pkg/)
 - [Go maps in action](https://blog.golang.org/maps)
-------------------------------------------------

## [基本型態(Basic Type)](https://michaelchen.tech/golang-programming/data-type/)
- 布林值(真或假)：
	- bool
- 字串：
	- string
- 整數(Integer)
	- 有號數
		- int  int8  int16  int32  int64 
	- 無號數
		- uint uint8 uint16 uint32 uint64 uintptr
- 符點數(float)
	- float32 (含小數點 7 位)
		```golang
		package main
		import "fmt"
		func main() {
			var f1 float32 = 0.1234567123 * 10;
			var f2 float32 = 0.1234567123;
			fmt.Printf("%v\n", f1)
			fmt.Printf("%v\n", f2)
		}
		```
	- float64 (含小數點 15 位)
		```golang
		package main
		import "fmt"
		func main() {
			var f1 float64 = 0.1234567890123456789 * 10;
			var f2 float64 = 0.1234567890123456789;
			fmt.Printf("%v\n", f1)
			fmt.Printf("%v\n", f2)
		}
		```
 
- byte // alias for uint8
- rune // alias for int32, represents a Unicode code point

-------------------------------------------------

## [條件敘述(if and switch)](https://michaelchen.tech/golang-programming/selection/)

-------------------------------------------------

## [迴圈(for loop)](https://michaelchen.tech/golang-programming/iteration/)

-------------------------------------------------

## [陣列 (Array) 和切片 (Slice)](https://michaelchen.tech/golang-programming/array-slice/)

## [向量 (Vector) 和矩陣 (Matrix)](https://michaelchen.tech/golang-cookbook/vector-matrix/)

-------------------------------------------------

## [映射 (Map)](https://michaelchen.tech/golang-programming/map/)
- 字串 key 與字串 value 的映射
	```golang
	package main
	import "fmt"

	//https://stackoverflow.com/a/31064737/11474144
	func MapInit() map[string]string {
		fmt.Printf("\nMapInit\n");

		//case 1
		// Initializes a map with space for 15 items before reallocation
		// the first 15 items added to the map will not require any map resizing
		m1 := make(map[string]string, 15);
		fmt.Printf("len:%d %v\n", len(m1), m1);

		//case 2
		// Initializes a map with an entry relating the name "bob" to the string "bbb"
		m2 := map[string]string{"bob": "bbb"};
		fmt.Printf("len:%d %v\n", len(m2), m2);

		//case 3
		// Initializes a map with three entry
		m3 := map[string]string {
			"aaa": "AAA",
			"bbb": "BBB",
			"ccc": "CCC",
		};
		fmt.Printf("len:%d %v\n", len(m3), m3);
		return m3;
	}

	func MapIterate(m map[string]string) {
		fmt.Printf("\nMapIterate\n");
		for key,value := range m {
			fmt.Printf("key:%v value:%v\n", key, value);
		}
	}

	func MapRemoveByKey(m map[string]string, key string) {
		fmt.Printf("\nMapRemoveByKey %v\n", key);
		delete(m, key);
	}

	func MapInsertKeyValuePair(m map[string]string, key string, value string) {
		fmt.Printf("\nMapInsertKeyValuePair %v %v\n", key, value);
		m[key] = value;
	}

	func MapGetValueByKey(m map[string]string, key string) {
		fmt.Printf("\nMapGetValueByKey %v\n", key);
		value, ok := m[key];
		if ok {
			fmt.Printf("the value is %v\n", value);
		} else {
			fmt.Printf("the key is not found %v\n", key);
		}
	}

	func main() {
		m := MapInit();
		MapIterate(m);

		MapRemoveByKey(m, "bbb");
		MapIterate(m);

		MapInsertKeyValuePair(m, "ddd", "DDD");
		MapIterate(m);

		MapGetValueByKey(m, "aaa");
		MapGetValueByKey(m, "aaA");
	}

	```

## [結構(Struct)](https://michaelchen.tech/golang-programming/struct/)

-------------------------------------------------

## [指標(Pointer)](https://michaelchen.tech/golang-programming/pointer/)
- [Pass by pointer vs pass by value in Go](https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/)

```golang
package main
import "fmt"

type Point struct {
	x float64
	y float64
}

func pass_by_pointer(p *Point, new_x float64) {
	p.x = new_x;
}

func pass_by_value(p Point, new_x float64) {
	p.x = new_x;
}

func main() {
	var p *Point = &Point{x: 1.0, y: 2.0};
	pass_by_pointer(p, 12.0);
	fmt.Printf("%v\n", *p); //{12 2}

	var p2 Point = Point{x: 1.0, y: 2.0};
	pass_by_value(p2, 12.0);
	fmt.Printf("%v\n", p2); //{1 2}
}
```

-------------------------------------------------

## [函式(Function)](https://michaelchen.tech/golang-programming/function/)
```golang
package main
import "fmt"

//代入一個值與回傳一個值
func test1(n1 int) int {
	return n1;
}

//代入二個值與回傳二個值
func test2(n1 int, n2 int) (int, int) {
	return n1, n2;
}

func main() {
	//宣告變數但不指定型別
	var n1, n2 = test2(1, 2);
	fmt.Printf("a: %v %v\n", n1, n2);

	//https://stackoverflow.com/questions/53404305/when-to-use-var-or-in-go/53404332
	//宣告變數與指定型別
	var n3 int;
	fmt.Printf("b: %v\n", n3);

	n3 = test1(3);
	fmt.Printf("c: %v\n", n3);

	// := 不使用 var 宣告
	n4 := 123;
	n4 = test1(4);
	fmt.Printf("d: %v\n", n4);

	// 省略第一個回傳值
	_, n5 := test2(4, 5);
	fmt.Printf("e: %v\n", n5);
}
```

-------------------------------------------------

## [類別 (Class) 和物件 (Object)](https://michaelchen.tech/golang-programming/class-object/)
- 物件最大的意途是抽象化與封裝，將重覆的程序、業務邏輯與附加邏輯隱藏在物件裡讓其他的物件可以重覆呼叫與使用。與函數不同點在於物件裡有成員變數(member variable)，並且物件有繼承可以延伸父類別的行為、介面提供附加邏輯的具體實作。
- 但 Golang 並沒有提供繼承，只有介面。其原因為繼承會增加耦合的成本，無法在動態時期改變具現實例。

```golang
package main
import "fmt"

type PointX struct {
	x float64
}

func (p *PointX) X() float64 {
	return p.x;
}

func (p *PointX) SetX(x float64) {
	p.x = x;
}

type PointY struct {
	y float64
};

func (p *PointY) Y() float64 {
	return p.y;
}

func (p *PointY) SetY(y float64) {
	p.y = y;
}

//embbeded struct of PointX and PointY
type PointXYZ struct {
	PointX
	PointY
	z float64
};

//override PointX's SetX function
func (p *PointXYZ) SetX(x float64) {
	p.PointX.x = x * x;
}

func (p *PointXYZ) Z() float64 {
	return p.z;
}

func (p *PointXYZ) SetZ(z float64) {
	p.z = z;
}

func main() {
	pxy := &PointXYZ{};
	pxy.SetX(3.0);
	pxy.SetY(6.0);
	pxy.SetZ(9.0);
	fmt.Printf("%v %v %v\n", pxy.X(), pxy.Y(), pxy.Z());
}
```

-------------------------------------------------

## [介面 (Interface) 實踐繼承和多型](https://michaelchen.tech/golang-programming/interface/)
- 業務邏輯可以透過介面將附加邏輯給隔開，而實現達到開放封閉原則。最好的例子就是外掛(Plugin-In)。譬如 Foobar2000 音樂播放器的業務邏輯為
	- 檔案讀取模組 (Source Reader) -> 解碼模組 (Decoder) ->  放音模組 (Audio Playback)
- 業務邏輯的主要流程盡量不會變更，變更的為上述三個附加邏輯的模組，當模組化後面對需求，對程式碼的改動是透過增加新程式碼進行的，而不是更改現有的程式碼。檔案讀取模組可以是"檔案讀取"或"串流讀取"或未來的新模組，解碼模組可以是"MP3解碼模組"或"AAC解碼模組或未來的新模組。
- 模組式的設計就是透過介面提供一致性的接口讓業務邏輯去使用。

```golang

package main
import "fmt"

type Reader interface {
	Read(buf *Buf) error
}

type Decoder interface {
		Decode(buf *Buf) error;
}

type Playback interface {
		Playback(buf *Buf) error;
}

type Buf struct {
	in_buf []byte;
	in_len int;
	out_buf []byte;
	out_len int;
}

func NewBuf(in_alloc_len int, out_alloc_len int) *Buf {
	return &Buf{
		make([]byte, in_alloc_len),
		0,
		make([]byte, out_alloc_len),
		0,
	};
}

func (p *Buf) InBuf() []byte {
	return p.in_buf;
}

func (p *Buf) SetInLen(n int) {
	p.in_len = n;
}

func (p *Buf) InLen() int {
	return p.in_len;
}

func (p *Buf) OutBuf() []byte {
	return p.out_buf;
}		

func (p *Buf) OutLen() int {
	return p.out_len;
}

func (p *Buf) SetOutLen(n int) {
	p.out_len = n;
}

func ReverseBuf(buf *Buf) {
	n := buf.InLen();
		for i := 0; i < n/2; i++ {
				buf.OutBuf()[i] = buf.InBuf()[n-1-i];
		buf.OutBuf()[n-1-i] = buf.InBuf()[i];
	}
	buf.OutBuf()[n/2] = buf.InBuf()[n/2];
	buf.SetOutLen(buf.InLen());
}

type FileReader struct {

}

func NewFileReader() *FileReader {
	return &FileReader{};
}

func (p *FileReader) Read(buf *Buf) error {
	for i:=0; i< len(buf.InBuf()); i++ {
		buf.InBuf()[i] = 'A' + byte(i);
	}
	buf.SetInLen(len(buf.InBuf()));
	return nil;
}


type Mp3Decoder struct {

}

func NewMp3Decoder() *Mp3Decoder {
	return &Mp3Decoder{};
}

func (p *Mp3Decoder) Info() string {
	return "128kbps,16bit,2ch";
}

func (p *Mp3Decoder) Decode(buf *Buf) error {
	ReverseBuf(buf);
	return nil;
}

type DirectSoundPlayback struct {

}

func NewDirectSoundPlayback() *DirectSoundPlayback {
	return &DirectSoundPlayback{};
}

func (p *DirectSoundPlayback) Playback(buf *Buf) error {
	fmt.Printf("Read from reader %v %v\n", buf.InBuf(), string(buf.InBuf()));
	fmt.Printf("Decode from decoder %v %v\n", buf.OutBuf(), string(buf.OutBuf()));
	return nil;
}


func main() {
	var buf *Buf = NewBuf(26, 26);

	var reader Reader = NewFileReader();
	var decoder Decoder = NewMp3Decoder();
	var playback Playback = NewDirectSoundPlayback();

	reader.Read(buf);
	decoder.Decode(buf);
	playback.Playback(buf);
}
```

-------------------------------------------------

## interface{} 變數
- interface{} 與 golang interface 並不是同一件事， interface{} 類似於 C 語言裡的 void * 變數，用於承接任何指標類型的實例。

```golang
package main
import "fmt"

type A struct {

}

func (p *A) Name() string {
	return "A";
}

type B struct {

}

func (p *B) Name() string {
	return "B";
}

type C struct {

}

func (p *C) Name() string {
	return "C";
}

func main() {
	interface_map :=  map[string]interface{} {
		"A": &A{},
		"B": &B{},
		"C": &C{},
	};

	//casting with error code
	vlaue_a, ok_a:= interface_map["A"];
	if ok_a {
		a, oka_2 := vlaue_a.(*A);
		if oka_2 {
			fmt.Printf("%v\n", a.Name());
		}
	}

	//casting without check error code
	value_b := interface_map["B"];
	if value_b != nil {
		b := value_b.(*B);
		if b != nil {
			fmt.Printf("%v\n", b.Name());
		}
	}

	//use .(type) to check which class instance is
	switch interface_map["C"].(type) {
		case *C:
			c := interface_map["C"].(*C);
			fmt.Printf("%v\n", c.Name());
			break;
	}
}
```
	
-------------------------------------------------

## 函數式程式設計 (Functional Programming)
- 閉包 (Closure)

```golang
package main
import "fmt"
 
func main() {
	n := 1;
	f := func() int {
		n += 1;
		return n;
	};
	fmt.Printf("%v\n", f());
}
```

- Callback function: 類似 C 的 function pointer

```golang
package main
import "fmt"

func DoAlsaAudioCapture(pcm_callback func([]float64) error) {
	pcm_data := make([]float64, 64);
	pcm_callback(pcm_data);
	
	pcm_data2 := make([]float64, 128);
	pcm_callback(pcm_data2);
	
	pcm_data3 := make([]float64, 256);
	pcm_callback(pcm_data3);
}
 
func main() {
	pcm_callback := func(pcm_data []float64) error {
		fmt.Printf("Receving number %v of audio samples\n", len(pcm_data));
		return nil;
	};
	DoAlsaAudioCapture(pcm_callback);
}
```

-------------------------------------------------

## 使用 Json Marshal/Unmarshal
- 當成員有了 json tag 那麼第一個字元必須大寫代表 export

```golang
package main
import (
	"fmt"
	"encoding/json"
)

type ServerConfig struct {
	ListenPort int		`json:"listen_port"`
	EnableSsl bool		`json:"enable_ssl"`
}

func main() {
	json_bytes := []byte(`
{
	"listen_port": 443,
	"enable_ssl": false
}
`);
	server_config := &ServerConfig { };
	json.Unmarshal(json_bytes, &server_config);
	fmt.Printf("%v\n", server_config);
	
	json_bytes, _ = json.Marshal(server_config);
	fmt.Printf("%v\n", string(json_bytes));
	
	json_bytes, _ = json.MarshalIndent(server_config, "", "\t")
	fmt.Printf("%v\n", string(json_bytes));
}
```

-------------------------------------------------

## [Custom JSON Marshalling in Go](http://choly.ca/post/go-json-marshalling/)
- 客製化 json marshal/unmarshal 的輸入/輸出欄位

```golang
package main
import (
	"fmt"
	"encoding/json"
)

type ServerConfig struct {
	ListenPort int		`json:"listen_port"`
	EnableSsl bool		`json:"enable_ssl"`
}

func (p *ServerConfig) MarshalJSON() ([]byte, error)  {
	return json.Marshal(*p);
}

func (p *ServerConfig) UnmarshalJson(data []byte) error {
	return json.Unmarshal(data, p);
}
 
func main() {
	json_bytes := []byte(`
{
	"listen_port": 443,
	"enable_ssl": false
}
`);
	server_config := &ServerConfig { };
	server_config.UnmarshalJson(json_bytes);
	fmt.Printf("%v\n", server_config);
	
	json_bytes, _ = server_config.MarshalJSON();
	fmt.Printf("%v\n", string(json_bytes));
}
```

-------------------------------------------------

## [Read/write from/to file](https://stackoverflow.com/a/9739903/11474144)

-------------------------------------------------


