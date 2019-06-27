package main

/*
	是一个存放可重用对象的值的容器，设计的目的是存放已经分配的但是暂时不用的对象，在需要用到的时候直接从pool中取。
	原因：由于golang内建的GC机制会影响应用的性能，为了减少GC，golang提供了对象重用的机制，也就是sync.Pool对象池；
*/

func main() {

}
