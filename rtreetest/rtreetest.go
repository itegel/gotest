package main

//import _ "net/http/pprof"
import (
	"fmt"
	"github.com/dhconnelly/rtreego"
	//"log"
	"math/rand"
	//"net/http"
	"os"
	"runtime/pprof"
	"time"
)

const num = 1000000

type Node struct {
	loc *rtreego.Rect
	id  int64
}

func (n *Node) Bounds() *rtreego.Rect {
	return n.loc
}
func generateRandNodes() [num]*Node {
	var tmpp1 rtreego.Point
	var tmpr *rtreego.Rect
	maxradius := 3
	var tmpradius int
	var nodes [num]*Node
	for i := 0; i < num; i++ {
		tmpp1 = rtreego.Point{float64(rand.Int() % num), float64(rand.Int() % num)}
		tmpradius = rand.Int()%maxradius + 1
		tmpr, _ = rtreego.NewRect(tmpp1, []float64{float64(tmpradius), float64(tmpradius)})

		nodes[i] = &Node{tmpr, int64(i)}
	}
	fmt.Println(len(nodes))
	//	for _, node := range nodes {
	//	fmt.Println(node.loc, node.id)
	//}
	return nodes
}

func main() {
	//	go func() {
	//		log.Println(http.ListenAndServe("localhost:6060", nil))
	//	}()

	f, _ := os.Create("profile_file")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	rt := rtreego.NewTree(2, 25, 50)
	t0 := time.Now()
	nodes := generateRandNodes()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	for _, node := range nodes {
		//fmt.Println(node.loc, node.id)
		rt.Insert(node)
	}
	t2 := time.Now()
	fmt.Printf("The call took %v to run.\n", t2.Sub(t1))

	fmt.Println(rt.Size())

	bb, _ := rtreego.NewRect(rtreego.Point{num / 2, num / 2}, []float64{num / 20, num / 20})

	results := rt.SearchIntersect(bb)
	t3 := time.Now()
	fmt.Printf("The call took %v to run.\n", t3.Sub(t2))

	fmt.Println(len(results))
	//for _, node := range results {
	//if realnode, ok := node.(*Node); ok {
	//			fmt.Println(realnode.loc, realnode.id)
	//} else {
	//	fmt.Println("not a type of Node")
	//}
	//	}
}
