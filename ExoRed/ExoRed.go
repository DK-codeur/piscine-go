package main
import "fmt"

func main(){
  for i:='z';i<'a'; i-- {
	fmt.Println(i)
  }
}

func main(){
	for i:='z'; i >= 'a'; i-- {
	  fmt.Print(string(i))
	}
	fmt.Println("")
  }

  
  func main(){
	for i:='A'; i <= 'Z'; i++ {
	  if i%2==1{
		fmt.Print(strings.ToLower(string(i)))
	  } else {
		fmt.Print(string(i))
	  }
	}
	fmt.Println("")
  }

Écrire un programme //qui affiche l'alphabet à l'envers, avec les lettres paires en majuscule, et les lettres impaires en minuscule, suivi d'un newline('\n').

  func main(){
	for i:='Z'; i >= 'A'; i-- {
	  if i%2==0{
		fmt.Print(strings.ToLower(string(i)))
	  } else {
		fmt.Print(string(i))
	  }
	}
	fmt.Println("")
  }
