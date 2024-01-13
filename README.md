# go
My own Go programming language codes repository  

Document last updated: *Fri 12 Jan 2024 18:48 PM GMT*  

-----

## Setting up Go programming language to work on Windows 10...

### Where to download Go programming language from...

1> Go to:   
https://www.golang.org  
...where you can download the Go programming language.

Then, download the appropriate [.msi] file...; then, after the file has downloaded...left double click on it to run...; repeatedly, clicking [Next] button should set it up to work on Windows 10...accepting all of the default setting options.

### How to test if set up has, actually, worked...

Next, open up a DOS prompt...and, type in...  
C:\>go [ENTER]  
...and, a list of Go commands should, automatically, be returned.  

You can also try doing...    
C:\>go version  
go version go1.17.1 windows/amd46  

### How to run [.go] programs you've written...

In order to run Go programs you've written yourself...  

Save the program files using filename extension: [.go];   
next, open up a DOS command prompt window, and, type...    

C:\go run filename.go  
...then, you should be able to observe the programs output.   

### Go programming language IDE's/Integrated Development Environment editor programs  

As well as using 'Windows Notepad' to write your Go programs in...which is *not* a proper programming language editor;  
Go programs can also be set up to run inside of the following IDE's:  

- Visual Studio Code  
- Eclipse   
- etc.  

-----

### NOTES  

First, let me state I'm a not an 'expert' in regards to the Go programming language; therefore, my comments are purely the observations of a complete and total 'amateur'.

However, I've heard it said that Go was invented by the people at Google; and, combines elements of...

- C  
- Java    
- Python    

In particular,  

- it's simpler than C/Java...in that there are no semi-colons used at the end of line endings    
- and, there is no use of brackets inside of for statements     
- etc.  

-----

## Example codes  

### Program: Example 1: Print string literal: Hello, world!

>>package main  
>>
>>import "fmt"  
>>
>>func main(){  
>> fmt.Println("Hello, world!")  
>>}  

### Program: Example 2: Print maths literal: 1+1

>>package main  
>>
>>import "fmt"  
>>
>>func main(){  
>> fmt.Println(1+1)    
>>}  

### Program: Example 3: Combining string literal/and, maths literal...using the comma as a separator: (,)

>>package main  
>>
>>import "fmt"  
>>
>>func main(){  
>> fmt.Println("1 + 1 = ",1+1)    
>>}  

### Program: Example 4: Print out the 12 x Tables...

>>package main  

>>import "fmt"   

>>func main() {  
>>	for x:=1;x<13;x++{  
>>	 fmt.Println(x," X ",12," =",x*12)  
>>	}  
>>}  

-----

## Links...

Official site...  
https://www.golang.org  

Do Go programming online...  
http://play.golang.org/

## GitHub code repositories...

https://github.com/GoesToEleven/GolangTraining/blob/master/01_getting-started/01_helloWorld/main.go  

## Tutorials...  

W3Schools Tutorial  
- https://www.w3schools.com/go/index.php  

GoLangDojo.com/Cheatsheet [.pdf]  
- https://github.com/thegolangdojo/cheatsheet/blob/main/cheatsheet.pdf  

-----

## Videos...  

### YouTube  

General search results listing for: Go language tutorial  
https://www.youtube.com/results?search_query=go+language+tutorial  

Learn Go in 12 Minutes  
https://www.youtube.com/watch?v=C8LgvuEBraI  

Go / Golang Crash Course (Traversy Media)  
https://www.youtube.com/watch?v=SqrbIlUwR0U  

Golang in under an hour (2021)  
https://www.youtube.com/watch?v=N0fIANJkwic  

Go Tutorial Basic | Golang  
https://www.youtube.com/watch?v=ty49_v1tV44  

The Go Language: What Makes it Different? - Jay McGavren  
https://www.youtube.com/watch?v=FEFXjRoac_U  

Learn GO Fast: Full Tutorial  
- https://www.youtube.com/watch?v=8uiZC0l4Ajw  

How I Would Re-Learn Go, Advice after 1 Year and 1 Production App    
- https://www.youtube.com/watch?v=OqdBixi_y1s

How to build a Golang Backend in 5 minutes  
- https://www.youtube.com/watch?v=IwplIbwJtD0     

GoLang: 10+ UNIQUE Concepts/Conventions that Beginners Should Know About!  
- https://www.youtube.com/watch?v=CK5rLpZk5A8

7 Deadly Mistakes Beginner Go Developers Make (and how to fix them)  
- https://www.youtube.com/watch?v=biGr232TBwc  

This Is The BEST Way To Structure Your GO Projects  
- https://www.youtube.com/watch?v=dxPakeBsgl4  
 
Why I Barely Use Go Anymore...  
- https://www.youtube.com/watch?v=tQrOfj57U-4  

This Go package was archived. What do we do now?   
- https://www.youtube.com/watch?v=j584uJhWWhE  

Why are Companies Migrating from Java to Go? (GoLang Dojo)  
- https://www.youtube.com/watch?v=dfyaTa0bSNE
  
Why are Companies Migrating from Javascript to Go!?   
- https://www.youtube.com/watch?v=45w7bgOXC3c  

The One BIG Reason to Learn Google's Go Language  
- https://www.youtube.com/watch?v=rQQcIGqp0OU

The Real Reason why Go is not so Popular  
- https://www.youtube.com/watch?v=2g0MxYf6FEM

Go vs Rust: Which To Learn In 2024?   
- https://www.youtube.com/watch?v=LjIe4w_-vzk

- 
 






