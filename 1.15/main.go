package main

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  // justString = v[:100] - некорректно, ссылаемся на часть большой строки, из-за этого GC не почистит ее

  /* 
  корректно - копируем часть от большой строки и создаем новую строку, 
  благодаря чему активных ссылок на изначальную строку не будет
  */ 
  justString = string([]byte(v[:100])) 
}

func main() {
  someFunc()
}
