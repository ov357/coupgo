// Couples gagnants
// enter cotes

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bldmatrix(c []float64, lt []int, sel int) [20][20]int {
	// c = cotes, lt = liste type, sel = selection voulue 1=0..100 / 2=100..200 / 3=200+
	tblgain := [20][20]int{}
	var gains float64
	for k, v := range c {
		for k1, v1 := range c {
			if k != k1 {
				gains = v * v1 / 2
				tmp := int(gains)
				tblgain[k][k1] = tmp
			}
		}
	}
	return tblgain
}

func filterg2(g [20][20]int, f1, f2 int, bases []int, gain int) {
	var c, c2, c3, nbfois int = 0, 0, 0, 0
	//nbfois += 0
	rand.Seed(time.Now().UTC().UnixNano())
	for k, i := range g {
		for k1, j := range i {
			if k1 > k {
				for _, j2 := range bases {
					//fmt.Println(j, j2)
					if k1+1 == j2 || k+1 == j2 {
						if k != k1 {
							if j >= f1 && j <= f2 {
								c += 1
								y := rand.Perm(3)
								if y[0] == 1 || y[0] == 2 {
									if j > 0 {
										nbfois = int(gain/j) + 1
										fmt.Printf("cg [%2d,%2d] = %4d x %2d \n", k+1, k1+1, j, nbfois)
										c2 += 1
										c3 += nbfois
									}
								}
							} // if j
						} // if k
					} // if j == j2
				} // for j2
			}
		} // for j
	} // for i
	fmt.Printf("cbs = %d cout = %d sur total de %d - esperance gain = %4d", c2, c3, c, gain)
}
func filterg(g [20][20]int, f1, f2 int, gain int) {
	var c, c2, c3, nbfois int = 0, 0, 0, 0
	//gain := 1000
	rand.Seed(time.Now().UTC().UnixNano())
	for k, i := range g {
		for k1, j := range i {
			if k1 > k {
				if k != k1 {
					if j >= f1 && j <= f2 {
						c += 1
						y := rand.Perm(3)
						if y[0] == 1 || y[0] == 2 {
							if j > 0 {
								nbfois = int(gain/j) + 1
								fmt.Printf("cg [%2d,%2d] = %4d x %2d \n", k+1, k1+1, j, nbfois)
								//fmt.Println(k+1, k1+1, j)
								c2 += 1
								c3 += nbfois
							}
						}
					} // if j
				} // if k
			}
		} // for j
	} // for i
	fmt.Printf("XX cbs = %d cout = %d sur total de %d - esperance gain = %4d \n", c2, c3, c, gain)
}

func play1(lt []int, cotes []float64, min, max int) {
	// lt := []int{0, 6, 4, 9, 11, 12, 1, 8, 7, 3, 5, 2, 10} //
	// cotes := []float64{20, 34, 16, 6.9, 14, 6.2, 15, 32, 41, 20.2, 14.9, 19, 12.5, 30.5, 23.6, 5.5, 56, 60}
	tm := time.Now().Format(time.RFC850)
	data := "R1C4 - " + tm + "\n"
	g := [20][20]int{}
	s := 1
	g = bldmatrix(cotes, lt, s)
	//min, max, gain := 40, 99, 300
	gain := 300
	fmt.Printf("%ssans bases : [%3d - %3d] gain = %4d \n", data, min, max, gain)
	fmt.Println(len(cotes), cotes)
	filterg(g, min, max, gain)
}

func main() {
	lt := []int{0, 6, 4, 9, 11, 12, 1, 8, 7, 3, 5, 2, 10} //
	// set to zero cotes above 13e
	cotes := []float64{14, 6, 29, 14, 17, 4, 22, 15, 26, 11.7, 14.8, 25, 13, 8, 22, 25, 16}
	// generate matrix CG
	//s := 1
	//g = bldmatrix(cotes, lt, s)
	//r1 := iselect(fcbs, lt) // one round for favorites
	// sans bases
	min, max := 40, 99
	play1(lt, cotes, min, max)
	min, max = 100, 199
	play1(lt, cotes, min, max)
	// avec bases
	//b := []int{}
	//gain := 250 // gain net voulu
	//m1 := 140
	//m2 := 340
	//gain = 250 //rembourse le jeu filterg
	//fmt.Printf("Avec bases  [%3d - %3d] : %2d \n", m1, m2, b)
	//filterg2(g, m1, m2, b, gain)
}
