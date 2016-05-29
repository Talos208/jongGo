package main

import (
	"math/rand"
	"log"
	"time"
	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/engine"
	"fmt"
)

type Hai uint8

const (
	Manzu Hai = iota + 1
	Pinzu = Manzu + 9
	Sozu = Pinzu + 9
	Tuhai = Sozu + 9

)

type MentuType uint8

const (
	Shuntu MentuType = iota + 1
	Koutu
	Toitu
)


var MentuIndex  = []int{
	0,8,1,7,2,6,3,5,4,
}

type Mentu struct {
	hai	Hai
	mentuType	MentuType
}

func mkMents() ([]float64, []float64) {
	pais := make([]float64, 9)
	mnts := make([]float64, 25)

	// 面子
	i := 0
	for i < 4 {
		if rand.Intn(3) < 1 {
			// 刻子
			ix := rand.Intn(9)
			if pais[ix] > 1 {
				continue
			}
			//log.Print("刻子:", ix)
			mnts[ix] += 1
			pais[ix] += 3
		} else {
			// 順子
			ix := MentuIndex[2 + rand.Intn(7)]
			if pais[ix - 1] > 3 || pais[ix    ] > 3 || pais[ix + 1] > 3 {
				continue
			}
			//log.Print("順子:", ix - 1, "～", ix + 1)
			mnts[ix + 8] += 1
			pais[ix - 1] += 1
			pais[ix    ] += 1
			pais[ix + 1] += 1
		}
		i++
	}
	// 頭
	for {
		ix := rand.Intn(9)
		if pais[ix] > 2 {
			continue
		}
		//log.Print("対子:", ix)
		mnts[ix + 9 + 7] += 1
		pais[ix] += 2
		break
	}
	for i,_ := range(mnts) {
		mnts[i] *= 0.25
	}

	return pais, mnts
}

func dump(datum []float64) {
	out := "[ "
	for _, data := range(datum) {
		out += fmt.Sprintf("%3.2f ", data * 4)
	}
	out += "]"
	log.Print(out)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 入力：0-8の9個
	// 中間層：試しに出力と同じに
	// 出力：順子2-8、刻子0-8、対子0-8の25個
	nn := neural.NewNetwork(9, []int{25,25})
	nn.RandomizeSynapses()

	eng := engine.New(nn)
	eng.Start()


	for j:=0;j < 50000;j++ {
		//log.Print(pais)
		pais, mnts := mkMents()

		eng.Learn(pais, mnts, 1.0)
	}

	// 試し
	pais, m0 := mkMents()

	mnts := eng.Calculate(pais)
	log.Print(pais)
	log.Print("")
	dump(m0[:9])
	dump(mnts[:9])
	log.Print("")
	dump(m0[9:16])
	dump(mnts[9:16])
	log.Print("")
	dump(m0[16:])
	dump(mnts[16:])

}
