package main
import(
"os"
"fmt"
"io/ioutil"
"strings")

func main() {

counts :=  make(map[string]int)
for _, filename := range os.Args[1:] {
data , err := ioutil.ReadFile(filename)
if err != nil {
fmt.Fprintf(os.Stderr , "duplicate : %v \n",err)
continue
}
for _ ,line := range strings.Split(string(data) ,"\n") {
counts[line]++}

}
for line ,n := range counts {

if n >1 {
fmt.Printf("te duplicate %d \t %s \n",n ,line)}}
}
