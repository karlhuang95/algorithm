package isAnagram

import "sync"

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	//创建两个仿bitmap的数组
	var wg sync.WaitGroup
	var wordstable_s =  [26]int{}
	var wordstable_t =  [26]int{}
	wg.Add(2)
	go func(){
		makewordtable(&wordstable_s,s)
		wg.Done()
	}()
	go func(){
		makewordtable(&wordstable_t,t)
		wg.Done()
	}()
	wg.Wait()
	//最后比较一下两个数组就好了
	return wordstable_s == wordstable_t

}

func makewordtable(wordtable *[26]int, s string) {
	for i:=0;i<len(s);i++{
		//使用asc2码 计算对应仿bitmap数组的下标
		index := s[i] - 'a'
		wordtable[index] ++
	}
}

