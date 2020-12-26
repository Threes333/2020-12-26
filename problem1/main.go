package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

const (
	//黑桃
	Spade = 0
	//红桃
	Hearts = 1
	//梅花
	Club = 2
	//方块
	Diamond = 3
)

type Poker struct {
	Num int
	Flower int
}

func (p Poker)PokerSelf()string  {
	var buffer string

	switch p.Flower {
	case Spade:
		buffer += "♤"
	case Hearts:
		buffer += "♡"
	case Club:
		buffer += "♧"
	case Diamond:
		buffer += "♢"
	}
	switch p.Num {
	case 13:
		buffer += "2"
	case 12:
		buffer += "A"
	case 11:
		buffer += "K"
	case 10:
		buffer += "Q"
	case 9:
		buffer += "J"
	default:
		buffer += strconv.Itoa(p.Num+2)
	}

	return buffer
}

func CreatePokers()(pokers Pokers)  {
	for i := 1; i < 14; i++ {
		for j := 0; j < 4; j++ {
			pokers = append(pokers,Poker{
				Num:    i,
				Flower: j,
			})
		}
	}
	return
}

type Pokers []Poker

func (p Pokers)Print()  {
	for _, i2 := range p {
		fmt.Print(i2.PrintPoker()," ")
	}
	fmt.Println()
}

func (p Poker) PrintPoker()string  {
	return p.PokerSelf()
}
func DaLuan(p Pokers)Pokers  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pokers := make(Pokers,len(p))
	perm := r.Perm(len(p))
	for i,index := range perm{
		pokers[i] = p[index]
	}
	return pokers
}
/*type Interface interface {
	Len() int
	Less(i int, j int) bool
	Swap(i int, j int)
}*/
func (p Pokers) Len()int {
	return len(p)
}
func (p Pokers) Less(i int ,j int)bool {
	if p[i].Num < p[j].Num{
		return true
	}else if p[i].Num == p[j].Num{
		return p[i].Flower < p[j].Flower
	}
	return  false
}
func (p Pokers) Swap(i int, j int) {
	p[i],p[j] = p[j],p[i]
}
func main() {
	pokers := make(Pokers,0)
	pokers = CreatePokers()
	fmt.Println("洗牌前:")
	pokers.Print()
	pokers = DaLuan(pokers)
	fmt.Println("洗牌后:")
	pokers.Print()
	sort.Sort(pokers)
	fmt.Println("排序后:")
	pokers.Print()
}
