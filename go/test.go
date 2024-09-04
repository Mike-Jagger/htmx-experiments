package main

// func add(x, y int) int {
// 	return x + y
// }

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func split(sum int) (x, y int) {
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

// var i, j int = 1, 2

// var (
// 	Bool bool = true

// 	String string = "string variable"

// 	Int int = 69
// 	Int8 int8
// 	Int16 int16
// 	Int32 int32
// 	Int64 int64

// 	Uint uint
// 	Uint8 uint8
// 	Uint16 uint16
// 	Uint32 uint32
// 	Uint64 uint64 = 1<<64 - 1
// 	Uintptr uintptr

// 	Byte byte

// 	Rune rune

// 	FLoat32 float32
// 	Float64 float64

// 	Complex64 complex64
// 	Complex128 complex128 = cmplx.Sqrt(-5 + 12i)
// )

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}
// 	return fmt.Sprint(math.Sqrt(x))
// }

// func pow(x, n, lim float64) float64 {
// 	if v:= math.Pow(x, n); v < lim {
// 		return v
// 	} else {
// 		fmt.Printf("%g >= %g\n", v, lim)
// 	}

// 	return lim
// }

// func PrintTypeNValue(_var interface{}) {
// 	fmt.Printf("Type: %T Value: %v \n", _var, _var)
// }

// func Sqrt(x float64) float64 {
// 	var z float64 = float64(1)

// 	var prev float64

// 	for ; z > math.Sqrt(x) || prev != z; {
// 		fmt.Printf("z:%v x:%v \n", z, math.Sqrt(x))
// 		z -= (z*z - x) / (2 * z)

// 		prev = z
// 	}

// 	return z
// }

// func willDeferPrints(a, b string) {
// 	defer fmt.Println(b)

// 	fmt.Println(a)
// }

// func CopyFile(dstName, srcName string) (written int64, err error) {
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 		return
// 	}
// 	defer src.Close()

// 	dst, err := os.Create(dstName)
// 	if err != nil {
// 		return
// 	}
// 	defer dst.Close()

// 	return io.Copy(src, dst)
// }

// func f() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Printf("Recoverd in f %v\n", r)
// 		}
// 	}()

// 	fmt.Println("Calling g")
// 	g(0)
// 	fmt.Println("Returned normally from g")

// }

// func g(i int64) {
// 	if i > 3 {
// 		fmt.Println("Panicking!")
// 		panic(fmt.Sprintf("%v", i))
// 	}

// 	defer fmt.Printf("Defer G %v\n", i)
// 	fmt.Printf("Print in G %v\n", i)

// 	g(i + 1)
// }

// func Pic(dx, dy int) [][]uint8 {
// 	pic := make([][]uint8, dy)

// 	for i := range pic {
// 		pic[i] = make([]uint8, dx)
// 	}

// 	return pic
// }

// func WordCount(s string) map[string]int {
// 	stringSlice := strings.Fields(s)

// 	if len(stringSlice) == 0 {
// 		return map[string]int{
// 			"": 0,
// 		}
// 	}

// 	wordCountMap := make(map[string]int)

// 	for _, v := range stringSlice {
// 		if wordCountMap[v] != 0 {
// 			wordCountMap[v] += 1
// 		} else {
// 			wordCountMap[v] = 1
// 		}
// 	}

// 	return wordCountMap
// }

// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 2)
// }

// func square(x, y float64) float64 {
// 	return x * y
// }

// func modifyPatientCount(action bool, countPtr *int) func (int) int{
// 	if action {
// 		return func(amount int) int {
// 			*countPtr += amount
// 			return *countPtr
// 		}
// 	} else {
// 		return func(amount int) int {
// 			*countPtr -= amount
// 			return *countPtr
// 		}
// 	}
// }

// type MyFLoat float64

// func (f MyFLoat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}

// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X * v.X + v.Y * v.Y)
// }

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X * v.X + v.Y * v.Y)
// }

// func (v *Vertex) Scale(f float64) {
// 	(*v).X = (*v).X * f
// 	(*v).Y = (*v).Y * f
// }

// func (v Vertex) getScale(f float64) (x, y float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f

// 	x = v.X
// 	y = v.Y
// 	return
// }

// type I interface {
// 	M()
// }

// type T struct {
// 	S string
// }

// func (t *T) M() {
// 	if t == nil {
// 		fmt.Println("<nil>")
// 		return
// 	}

// 	fmt.Println((*t).S)
// }

// type F float64

// func (f F) M() {
// 	fmt.Println(f)
// }

// func describe(i I) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func describeI(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

// func isPanicking(i interface{}) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("panicking because of wrong type assertion:", r)
// 		}
// 	}()

// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("is of type int: %T\n", v)
// 	case float64:
// 		fmt.Printf("is of type float: %T\n", v)
// 	default:
// 		fmt.Printf("is of other type: %T", v)
// 	}
// }

// type Person struct {
// 	Name string
// 	Age int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

// type IPAddr [4]byte

// func (ipAddr IPAddr) String() string {
// 	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
// }

// type MyError struct {
// 	Reason string
// 	Service string
// 	Time time.Time
// }

// func (error *MyError) Error() string {
// 	return fmt.Sprintf("Service: %v \nTIme: %v \nReason: %v", error.Service, error.Time,error.Reason)
// }

// func run() (int, error) {
// 	return 0, &MyError{
// 		"Couldn't find shop in database",
// 		"Verfication",
// 		time.Now(),
// 	}
// }

// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, ErrNegativeSqrt(x)
// 	}
// 	return math.Sqrt(x), nil
// }

// type err string

// func (reason err) Error() string {
// 	return fmt.Sprintf("%v", string(reason))
// }

// func copyToBuffer(src *strings.Reader, dest *[]byte, chunkSize int) ([]byte, error) {
// 	if cap(*dest) < src.Len() {
// 		return make([]byte, 0), err("Destination buffer too small")
// 	}

// 	b := make([]byte, chunkSize)

// 	j := 0
// 	for i := 0; i <= int(math.Round(float64(src.Len()/chunkSize))); i++ {
// 		n, err := src.Read(b)

// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b)

// 		for k := 0; k < chunkSize; k++ {
// 			(*dest)[j] = b[k]
// 			j++
// 		}

// 		if err == io.EOF {
// 			break
// 		}
// 	}

// 	return *dest, nil
// }

// func Index[T comparable](s []T, x T) int {
// 	for i, v := range s {
// 		if v == x {
// 			return i
// 		}
// 	}

// 	return -1
// }

// type integerQueue[T any] interface {
// 	pop() Node[T]
// 	shift() Node[T]
// 	peek() Node[T]
// }

// type List[T any] struct {
// 	head *Node[T]
// 	length int
// }

// type Node[T any] struct {
// 	next *Node[T]
// 	val T
// }

// func (list *List[T]) peek() Node[T] {
// 	if list.head != nil {
// 		return *list.head
// 	}
// 	panic("There is no value in list")
// }

// func (list *List[T]) shift() Node[T] {
// 	var node Node[T]
// 	var head *Node[T] = list.head

// 	node = *head
// 	list.head = (*head).next

// 	return node
// }

// func (list *List[T]) pop() Node[T] {
// 	var head *Node[T] = list.head
// 	var node Node[T]

// 	var curr *Node[T] = head

// 	if curr.next == nil {
// 		list.head = nil
// 		return *curr
// 	}

// 	for curr != nil {
// 		if curr.next.next == nil {
// 			node = *((*curr).next)
// 			curr.next = nil
// 		}

// 		curr = curr.next
// 	}

// 	return node
// }

// func say(s int) {
// 	for i :=0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func Walk(tree *tree.Tree, ch chan int) {
// 	if tree == nil {
// 		return
// 	}

// 	Walk(tree.Left, ch)
// 	ch <- tree.Value
// 	Walk(tree.Right, ch)
// }

// func Same(t1, t2 *tree.Tree) bool {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)

// 	for i := 0; i < 10; i++ {
// 		// fmt.Println(<- ch1)
// 		// fmt.Println(<- ch2)
// 		n1 := <- ch1
// 		n2 := <- ch2

// 		if n1 != n2 {
// 			return false
// 		}
// 	}

// 	return true
// }

// type Fetcher interface {
// 	// Fetch returns the body of URL and
// 	// a slice of URLs found on that page.
// 	Fetch(url string) (body string, urls []string, err error)
// }

// // Crawl uses fetcher to recursively crawl
// // pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher, cache *URLCache) {
// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:

// 	if depth <= 0 {
// 		return
// 	}

// 	if !cache.inCache(url) {
// 		cache.addURL(url)
// 		body, urls, err := fetcher.Fetch(url)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Printf("found: %s %q\n", url, body)
// 		for _, u := range urls {
// 			Crawl(u, depth-1, fetcher, cache)
// 		}
// 	} else {
// 		// fmt.Printf("\nAlready in cache: %v \n\n", url)
// 	}
// }

// // fakeFetcher is Fetcher that returns canned results.
// type fakeFetcher map[string]*fakeResult

// type fakeResult struct {
// 	body string
// 	urls []string
// }

// func (f fakeFetcher) Fetch(url string) (string, []string, error) {
// 	if res, ok := f[url]; ok {
// 		return res.body, res.urls, nil
// 	}
// 	return "", nil, fmt.Errorf("not found: %s", url)
// }

// // fetcher is a populated fakeFetcher.
// var fetcher = fakeFetcher{
// 	"https://golang.org/": &fakeResult{
// 		"The Go Programming Language",
// 		[]string{
// 			"https://golang.org/pkg/",
// 			"https://golang.org/cmd/",
// 		},
// 	},
// 	"https://golang.org/pkg/": &fakeResult{
// 		"Packages",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/cmd/",
// 			"https://golang.org/pkg/fmt/",
// 			"https://golang.org/pkg/os/",
// 		},
// 	},
// 	"https://golang.org/pkg/fmt/": &fakeResult{
// 		"Package fmt",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// 	"https://golang.org/pkg/os/": &fakeResult{
// 		"Package os",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// }

// type URLCache struct {
// 	mu sync.Mutex
// 	store map[string]bool
// }

// func (cache *URLCache) inCache(key string) bool {
// 	cache.mu.Lock()
// 	defer cache.mu.Unlock()

// 	_, ok := (*cache).store[key]

// 	return ok
// }

// func (cache *URLCache) addURL(key string) {
// 	cache.mu.Lock()

// 	((*cache).store)[key] = true

// 	cache.mu.Unlock()
// }



func main() {
	// var customCache URLCache = URLCache{store: make(map[string]bool)}

	// Crawl("https://golang.org/", 4, fetcher, &customCache)

	// fmt.Println(customCache)

	// fmt.Println(Same(tree.New(1), tree.New(2)))

	// i := 5
	// i_ptr := &i
	// go say(*i_ptr)
	// i++
	// say(*i_ptr)

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Error:", r)
	// 	}
	// }()
	// var node1 Node[int] = Node[int]{nil, 1} 
	// var node2 Node[int] = Node[int]{nil, 2}


	// var listInterfaceToBeUsed List[int] = List[int]{&node1, 1}

	// var list integerQueue[int] = &listInterfaceToBeUsed
	
	// node1.next = &node2

	// fmt.Println(list.peek())

	// fmt.Println(list.shift())
	// fmt.Println(list.peek())

	// fmt.Println(list.pop())
	// fmt.Println(list.peek())
	
	// var intArr []int = []int{10, 20, 15, -10}
	// fmt.Println(Index(intArr, 2))
	// fmt.Println(Index(intArr, 20))

	// var strArr []string = []string{"foo", "bar", "baz"}
	// fmt.Println(Index(strArr, "hello"))
	// fmt.Println(Index(strArr, "baz"))
	
	// b := make([]byte, 16)
	// r := strings.NewReader("Hello, Reader!!!")

	// val, err := copyToBuffer(r, &b, 8)

	// fmt.Printf("String: %q Value: %v", val[:], err)

	// val, err := sqrt(64)

	// fmt.Println(val, err)

	// val, err = sqrt(-2)
	// fmt.Println(val, err)

	// value, err := run()

	// if err != nil && value == 0 {
	// 	fmt.Printf("Error while running: \n%v \nValue: \t%v \n", err, value)
	// }


	// hosts := map[string]IPAddr {
	// 	"loopback": {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }

	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }

	// a := Person{"John Doe", 8}
	// var description string = a.String()
	// fmt.Println(description)
	// fmt.Println(a)
	
	// var i interface{} = "Hello world"

	// describeI(i)

	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok) 

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// isPanicking(i)

	// var i I

	// var t *T
	// i = t
	// describe(i)
	// i.M()
	
	// describe(t)
	// t.M()

	// i = &T{"Hello"}
	// describe(i)
	// i.M()

	// t := T{"String type"}
	// i = &t
	// describe(i)
	// i.M()
	
	// f := F(math.Pi)
	// i = f
	// describe(f)
	// i.M()

	// topLeftCorner := Vertex{3, 4}
	// f := MyFLoat(-math.Sqrt2)
	
	// fmt.Println(topLeftCorner.Abs())
	// fmt.Println(Abs(topLeftCorner))

	// fmt.Println(f.Abs())

	// topLeftCorner.Scale(8)
	// fmt.Println(topLeftCorner)

	// fmt.Println(topLeftCorner.getScale(1./8.))
	// fmt.Println(topLeftCorner)
	// patientCount := 0
	// patientCountPtr := &patientCount

	// addPatient, removePatient := modifyPatientCount(true, patientCountPtr), modifyPatientCount(false, patientCountPtr)
	
	// for count := 0; count < 100; count++ {
	// 	fmt.Println(addPatient(2))
	// 	fmt.Println(removePatient(1))
	// }
	
	// fmt.Println(compute(math.Pow))
	// fmt.Println(compute(square))

	// fmt.Println(WordCount("This is a sentence with a few a's with an a"))

	// fmt.Printf("%v", Pic(6, 6))

	// f()
	// fmt.Println("Returned normally from f")


	// fmt.Println(Sqrt(32))

	// fmt.Print("Go runs on: ")

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }

	// saturday := time.Saturday

	// switch today := time.Now().Weekday(); saturday { 
	// case today:
	// 	fmt.Println("It's Saturday!")
	// case today + 1:
	// 	fmt.Println("It's tomorrow")
	// case today + 2:
	// 	fmt.Println("It's in 2 days")
	// default:
	// 	fmt.Println("It is too far away")
	// }

	// hour := time.Now().Hour()

	// switch {
	// case hour < 12:
	// 	fmt.Println("Good morning!")
	// case hour < 17:
	// 	fmt.Println("Good afternoon")
	// default:
	// 	fmt.Print("Good evening")
	// }

	// willDeferPrints("hello", "world")
	// fmt.Println(
	// 	pow(3, 2, 9),
	// 	pow(2, 2, 10),
	// )


	// PrintTypeNValue(Bool)
	// PrintTypeNValue(String)
	// PrintTypeNValue(Int)
	// PrintTypeNValue(Uint64)
	// PrintTypeNValue(Complex128)

	// fmt.Println()

	// PrintTypeNValue(Bool)
	// PrintTypeNValue(Int32)
	// PrintTypeNValue(Uint)
	// PrintTypeNValue(Float64)
	// PrintTypeNValue(Rune)
	// PrintTypeNValue(Byte)
	// PrintTypeNValue(Complex64)

	// fmt.Println()

	// PrintTypeNValue(float64(Int))

	// fmt.Println()

	// integerValue := 24
	// fmt.Println(math.Sqrt(2.22 + float64(integerValue)))

	// fmt.Println()

	// fmt.Printf("%T\n", 2)
	// fmt.Printf("%T\n", 2.0)
	// fmt.Printf("%T\n", 0 + 0i)

	// fmt.Println()

	// const cInt = 2
	// const cString = "Constant string"
	// const cBool = true
	// const cChar = "世界"

	// PrintTypeNValue(cInt)
	// PrintTypeNValue(cChar)
	// PrintTypeNValue(cString)
	// PrintTypeNValue(cBool)

	// fmt.Println()

	// const (
	// 	Big = 1 << 100
	// 	Small = Big >> 99
	// )

	// fmt.Printf("This is going to return float even though constant \n%v \n", Big * 0.1)
	// fmt.Printf("This is going to return int even though constant \n%v", Small * 10 + 1)




	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i is %v and sum is %v \n", i, sum)
	// 	sum = i * 10
	// } 


	// var c bool = true
	
	// for c {
	// 	if !c {
	// 		fmt.Println("If it comes here then not while loop")
	// 		break
	// 	}

	// 	fmt.Printf("Value of c is %v \n", c)

	// 	c = false
	// }
	
	// fmt.Printf("Checking if condition is met and loop breaks, then while loop simulated \n")
	

	// fmt.Println(sqrt(2), sqrt(-4))
	// var c, python = 1, true
	// java := "It is just Java bro, but with the dot:column declaration"

	// fmt.Println(add(1, 2))

	// a, b := swap("Hello", "World!")

	// fmt.Println(a, b)

	// fmt.Println(split(100))

	// fmt.Println(i, j, c, python, java)
}	