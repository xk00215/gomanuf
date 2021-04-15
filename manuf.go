package manuf

import (
   "os"
   "bufio"
   "strings"
   "io"
   "strconv"
   "runtime"
   "path"
)

const hexDigit = "0123456789ABCDEF"

var d map[int]interface{}
type vendor struct{
   d string
   b string
}
func init() {
   d = make(map[int]interface{})
   _, file, _, _ := runtime.Caller(0)
   f := path.Join(path.Dir(file), "manuf")
   err := readLine(f, func(s string) {
       l := strings.Split(s, "\t")
       if len(l) > 3 {
           parse(l[0], l[2],l[3])
       }else{
           if len(l) > 2 {
               parse(l[0], l[2],"")
           }
       }

   })
   if err != nil {
       panic(err)
   }
}

func parse(mac, comment,brand string) {
   var vd vendor
   g := strings.Split(mac, "/")
   m := strings.Split(g[0], ":")
   var b int
   if len(g) != 2 {
       b = 48 - len(m) * 8
   } else {
       b, _ = strconv.Atoi(g[1])
   }
   if _, ok := d[b]; !ok {
       d[b] = make(map[uint64]vendor)
   }
   vd.d=comment
   vd.b=brand
   d[b].(map[uint64]vendor)[b2uint64(m)] = vd
}
func b2uint64(sList []string) uint64 {
   var t uint64
   for i, b := range sList {
       l := strings.Index(hexDigit, string(b[0]))
       r := strings.Index(hexDigit, string(b[1]))
       t += uint64((l << 4) + r) << uint8((6 - i - 1) * 8)
   }

   return t
}

func Search(mac string) (string,string) {
   s := strings.Split(strings.ToUpper(mac) , ":")
   bint := b2uint64(s)
   for b := range d {
       k := 48 - b
       bint = (bint >> uint8(k)) << uint8(k)
       if _, ok := d[b].(map[uint64]vendor)[bint]; ok {
           return d[b].(map[uint64]vendor)[bint].d, d[b].(map[uint64]vendor)[bint].b
       }
   }
   return "",""
}

func readLine(fileName string, handler func(string)) error {
   f, err := os.Open(fileName)
   defer f.Close()
   if err != nil {
       return err
   }
   buf := bufio.NewReader(f)
   for {
       line, err := buf.ReadString('\n')
       line = strings.TrimSpace(line)
       handler(line)
       if err != nil {
           if err == io.EOF {
               return nil
           }
           return err
       }
   }
   return nil
}

