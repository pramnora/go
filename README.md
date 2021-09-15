# go
My own Go programming language codes repository

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

### Go programming language IDE's/Integrated Development Environment editor programs  

As well as using 'Windows Notepad' to write your Go programs in...which is *not* a proper programming language editor;  
Go programs can also be set up to run inside of the following IDE's:  

- Visual Studio Code  
- Eclipse   
- etc.  

-----

## Links...

Official site...  
https://www.golang.org  

Do Go programming online...  
http://play.golang.org/

-----

## Videos...  

### YouTube  




