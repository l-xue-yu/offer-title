package main

import "sync"

/*
实现一个单例，指一个类型（类）的实例化，在整个程序生命周期中只有一个。
知识点：
1.如何保证协程安全
1.1加锁
1.2once.Do
*/

type Person struct {
	Name string
	Sex  string
}

var singlePerson *Person
var once sync.Once

var mu sync.Mutex

/*
思路1加锁
*/
func GetSinglePerson2() *Person {
	mu.Lock()
	defer mu.Unlock()
	if singlePerson == nil {
		singlePerson = &Person{}
	}
	return singlePerson
}

//加锁优化,双重锁
func GetsinglePerson3() *Person {
	if singlePerson == nil {
		mu.Lock()
		defer mu.Unlock()
		if singlePerson == nil {
			singlePerson = &Person{}
		}
	}

}

/*
思路2：使用onceDo,保证只有一次实例化。
感觉思路其实和上面的双重锁一样的……
*/
func GetSinglePerson() *Person {
	once.Do(func() {
		singlePerson = &Person{}
	})
	return singlePerson
}

/*
once.Do知识点：使用互斥锁和原子操作来完成的，先用原子操作拿到done，判断是否执行过，执行过返回，没执行过，锁住这段代码，再判断一次，是否没执行，没执行，设置为1，表示已执行。

Once结构体
type Once struct {
    m    Mutex
    done uint32   // 初始值为0表示还未执行过，1表示已经执行过
}

func (o *Once) Do(f func()) {
    // done==1表示已经执行过了，直接结束返回
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    // 锁住对象，避免并发问题
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        // 先将done设置为1，再执行f函数
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}

type Once struct {
	done uint32 // 初始值为0表示还未执行过，1表示已经执行过
	m    Mutex  //互斥锁，防止数据竞态
}


func (o *Once) Do(f func()) {
	//原子读取done，判断是否已经执行过
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	//互斥锁，锁住这段代码，防止并发资源竞争
	o.m.Lock()
	defer o.m.Unlock()
	//这里不需要原子操作，因为，上面已经是锁住的资源，保证同一时间，只有一个协程能调用这块代码
	if o.done == 0 {
		//调用匿名函数后将done原子操作加1，
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
*/
