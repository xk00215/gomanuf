package manuf

import (
   "log"
   "testing"
)

func TestSearch(t *testing.T) {
   d1 ,b1:= Search("08:e2:44:45:0b:04")
   if len(d1) == 0 {
       t.Fatal("数据有问题")
   }
   log.Println(d1,b1)

   //d2 := Search("24:1f:a0:17:6d:9b")
   //if len(d2) == 0 {
   //    t.Fatal("数据有问题")
   //}

}
