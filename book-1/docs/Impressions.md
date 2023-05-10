## 1. Install/upgrade
easy install manual, but not perfect tooling. So confused information about upgrading Golang, I can't do after 4-5 pages from google search. Uninstall/Install process for upgrading seems as bad decision.

## 2. simple function syntax
```
func main() {
    fmt.Println("Hello World")
}
```
Why func ? Why not just: 
```
function main() {
    fmt.Println("Hello World")
}
```
or just:
```
main(){
    fmt.Println("Hello World")
}
```

Something like you're not so lazy to type "function" but after 4 characters you realize that you're a little bit lazy anyway, so you just type "func" :D

## 3. Variable naming in book
In chapter 4 author tell you that variable naming very important. In chapter 5 you see "i" variables everywhere.

## 4. Array, Slice, Map
Seems like slices, arrays and maps have very poor implementation. Without generics syntax very ugly.

## 5. Defer functions
First impression about golang defer functions is that a point of many unexpected errors, not really readable

## 6. Interface with struct 
don't require implementation for compile but runtime error. WTF ??? see lol.go in chapter 9

## 7. Interface with struct 2
I like way where you separate methods from struct but implementation in golang seems not so powerful enough. 
Need more researches.

## 8. Goroutines
Very easy and powerful construction but not easy to understand channels blocks

## 9. Biggest LOL
"Go was designed to be a language that encourages good software engineering practices. An important part of high quality software is code reuse – embodied in the principle “Don't Repeat Yourself.”"

## 10. Recompile. Very good point
"It speeds up the compiler by only requiring recompilation of smaller chunks of a program. Although we use the package fmt, we don't have to recompile it every time we change our program."