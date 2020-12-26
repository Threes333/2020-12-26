package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int ,5000)
var (
	res = 0
	lock sync.Mutex
)
func check(n int) bool {
	num := make([]int,0)
	for i := 1;i * i <= n;i++{
		if n%i == 0{
			if n/i == i || i==1{
				num = append(num,i)
			}else {
				num = append(num, i, n/i)
			}
		}
	}
	sum := 0
	for _,v := range num{
		sum += v
	}
	if sum == n{
		return true
	}
	return false
}
//求素数只需将check(i) 换成CheckPrimer(i)即可
func save(n int, m int) {
	for i:=n;i<m;i++{
		if check(i){
			ch <- i
		}
	}
	lock.Lock()   // 在进行存储操作之前，先加锁
	res += 1
	lock.Unlock()  // 当存储完毕后，进行解锁
	if res == 13{
		close(ch)
	}
}
func CheckPrime(n int) bool {
	for i:=2;i*i <= n;i++{
		if n%i == 0{
			return false
		}
	}
	return true
}

func main() {
	num := make([]int,0)
	for i:=1;i<=123456;i+=10000{
		if i+100 > 123456{
			go save(i,123456)
		}else{
			go save(i,i+10000)
		}
	}
	for i := range ch{
		num = append(num,i)
	}
	fmt.Println(num)
	fmt.Println("ok")
}
