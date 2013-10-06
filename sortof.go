package sortof

import (
    "fmt"
    "sort"
)

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

type Organ struct {
    Name   string
    Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

// ByWeight implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func main() {
    s := []*Organ{
        {"brain", 1340},
        {"heart", 290},
        {"liver", 1494},
        {"pancreas", 131},
        {"prostate", 62},
        {"spleen", 162},
    }

    sort.Sort(ByWeight{s})
    /*
       fmt.Println("Organs by weight:")
       printOrgans(s)

       sort.Sort(ByName{s})
       fmt.Println("Organs by name:")
       printOrgans(s)
    */
    a13 := []int{866, 700, 148, 587, 434, 898, 828, 893, 126, 657, 801, 868, 542}
    fmt.Printf("data: %v\n", a13)
    mergesort(a13)
    fmt.Printf("result: %v\n", a13)
}

func printOrgans(s []*Organ) {
    for _, o := range s {
        fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func bubblesort(arr []int) {
    for i := range arr {
        for j := range arr {
            if arr[i] < arr[j] {
                tmp := arr[j]
                arr[j] = arr[i]
                arr[i] = tmp
            }
        }
    }
}

func merge(arr []int, i1, e1, i2, e2 int) {
    result := make([]int, e2-i1)
    base := i1
    for i := 0; i < len(result); i++ {
        if i2 == e2 {
            result[i] = arr[i1]
            i1++
        } else if i1 == e1 {
            result[i] = arr[i2]
            i2++
        } else if arr[i1] < arr[i2] {
            result[i] = arr[i1]
            i1++
        } else {
            result[i] = arr[i2]
            i2++
        }
    }
    for i := 0; i < len(result); i++ {
        arr[base+i] = result[i]
    }
}

func mergesort(arr []int) {
    runlen := 1
    for runlen < len(arr) {
        runstart := 0
        for runstart < len(arr) {
            runend := min(runstart+runlen, len(arr))
            runstart2 := min(runstart+runlen, len(arr))
            runend2 := min(runstart+runlen*2, len(arr))
            merge(arr, runstart, runend, runstart2, runend2)
            runstart = runend2
        }
        runlen = runlen * 2
    }
}

func Ints(arr []int) {
    mergesort(arr)
}
