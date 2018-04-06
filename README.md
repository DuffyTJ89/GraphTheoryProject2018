# GraphTheoryProject2018



I am a third year Software development student at GMIT and this repo is for my Graph Theory project.

# Aims

  - The main aim of this project was to make a Go program that can build a non-deterministic finite automaton from a regular expression and then check to see if that regular expression matches a user inputted string.
  
# Run the application

The machine you intend to clone this repo to must have GO installed to run this program. To install GO:
- Instructions here https://golang.org/doc/install

Clone the repo and open git bash (or similar) in the folder and then type :

`go run main.go`

The program will run and you will be asked to enter the Regular expression you want to use. Then you will be asked to input a string to check against that regular expression. The result will be printed out to your screen.

There are smaller programs also contained in this repo that do smaller jobs. e.g shunt.go implements the shunting yard algroithm which can be used to change Regular expressions from Infix notation to pofix notation. Running Main.go will give you all these programs together in one. These programs are included to show how the project was written stage by stage.

# ScreenShots:

![](https://i.imgur.com/Ciaoy3d.png)

 a.b means an a followed by a b 
 a|b means an a or a b.
 
 ![](https://i.imgur.com/9iRjSGD.png)
 
 a* means any number of a's

# Technologies used:

This project is written in the GO programming language. 

# Resources used :
- Class Notes and Instructions
- StackOverFlow 
- Regular Expression Matching Can Be Simple And Fast by Russ Cox   https://swtch.com/~rsc/regexp/regexp1.html
- https://dillinger.io/ to edit markdown and check in real time.



Author : Thomas Duffy, GMIT.


