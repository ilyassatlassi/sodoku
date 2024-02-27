package piscine

func Split(s, sep string) []string {
  count := 0
  slice := []string{}
  for i := 0 ; i < len(s) - len(sep); i++ {
  //   for j := 1 ; j < len(s); j++ {
  // }
    if s[i: i + len(sep)] == sep {
      
      slice = append(slice,s[count:i]) 
      count = i +len(sep)
    }
  }
      slice = append(slice,s[count:len(s)]) 
return slice
}