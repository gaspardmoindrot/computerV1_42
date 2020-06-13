package	main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func suppr_space_poly(poly string) string {
	reduce := strings.Split(poly, " ")
	str := strings.Join(reduce, "")
	reduce = strings.Split(str, "\t")
	str = strings.Join(reduce, "")
	reduce = strings.Split(str, "\n")
	str = strings.Join(reduce, "")
	return (str)
}

func transform_polynome(poly string) map[float64]float64 {
	var i, sign, pos, equal, a int = 0, 1, 0, 0, 0
	var nb1, nb2 float64 = 0, 0
	var i_prev int = 0
	var err error

	m := make(map[float64]float64)
	reduce := suppr_space_poly(poly)

	for i < len(reduce) {
		if ((reduce[i] == '+' || reduce[i] == '-' || reduce[i] == '=') && pos > 0) {
			fmt.Println("syntax error : there cannot be two signs in a row")
			os.Exit(0)
		} else if (reduce[i] == '=' && equal == 1) {
			fmt.Println("syntax error : there cannot be two equals in a row")
			os.Exit(0)
		}
		pos = 0
		if (reduce[i] == '+') {
			pos = pos + 1
			sign = 1
		} else if (reduce[i] == '-') {
			pos = pos + 1
			sign = -1
		} else if (reduce[i] == '=') {
			equal = 1
			sign = 1
		} else {
			if (equal == 1) {
				sign = -sign
			}
			if (reduce[i] >= '0' && reduce[i] <= '9') {
				i_prev = i
				for (i < len(reduce) && ((reduce[i] >= '0' && reduce[i] <= '9') || reduce[i] == '.')) {
					i++
				}
				nb1, err = strconv.ParseFloat(reduce[i_prev:i], 64)
				if (err != nil) {
					fmt.Println("syntax error : bad number / a number is missing")
					os.Exit(0)
				}
				nb1 = nb1 * float64(sign)
				if (i >= len(reduce)) {
					nb2 = 0
				} else if (reduce[i] != '*' && (reduce[i] == '+' || reduce[i] == '-' || reduce[i] == '=')) {
					nb2 = 0
				} else if (reduce[i] == '*') {
					if (i + 1 >= len(reduce)) {
						fmt.Println("syntax error : '*' can't be the end of an equation")
						os.Exit(0)
					} else if (reduce[i + 1] != 'X') {
						fmt.Println("syntax error : you need to write an 'X' after a '*'")
						os.Exit(0)
					}
					i = i + 2
					if (i >= len(reduce)) {
						nb2 = 1
					} else if (reduce[i] == '+' || reduce[i] == '-' || reduce[i] == '=') {
						nb2 = 1
					} else if (i + 1 >= len(reduce)) {
						fmt.Println("syntax error : problem with a power")
						os.Exit(0)
					} else if (reduce[i] != '^' && (reduce[i] < '0' || reduce[i] > '9')) {
						fmt.Println("syntax error : problem with a power")
						os.Exit(0)
					} else {
						i++
						i_prev = i
						for (i < len(reduce) && ((reduce[i] >= '0' && reduce[i] <= '9') || reduce[i] == '.')) {
							i++
						}
						nb2, err = strconv.ParseFloat(reduce[i_prev:i], 64)
						if (err != nil) {
							fmt.Println("syntax error : bad number / a number is missing")
							os.Exit(0)
						}
					}
				} else {
					fmt.Println("syntax error : problem in the syntax of a number")
					os.Exit(0)
				}
			} else if (reduce[i] == 'X') {
				nb1 = float64(sign)
				i++
				if (i >= len(reduce)) {
					nb2 = 1
				} else if (reduce[i] == '+' || reduce[i] == '-' || reduce[i] == '=') {
					nb2 = 1
				} else if (i + 1 >= len(reduce)) {
					fmt.Println("syntax error : problem with a power")
					os.Exit(0)
				} else if (reduce[i] != '^' && (reduce[i] < '0' || reduce [i] > '9')) {
					fmt.Println("syntax error : problem with a power")
					os.Exit(0)
				} else {
					i++
					i_prev = i
					for (i < len(reduce) && ((reduce[i] >= '0' && reduce[i] <= '9') || reduce[i] == '.')) {
						i++
					}
					nb2, err = strconv.ParseFloat(reduce[i_prev:i], 64)
					if (err != nil) {
						fmt.Println("syntax error : bad number / a number is missing")
						os.Exit(0)
					}
				}
			} else {
				fmt.Println("syntax error : problem in the syntax of a number")
				os.Exit(0)
			}
			i--
			if (equal > 0) {
				a++
			}
			m[nb2] += nb1
		}
		i++
	}
	if (equal == 0) {
		fmt.Println("syntax error : we need an equal sign")
		os.Exit(0)
	}
	if (a == 0) {
		fmt.Println("syntax error : we need something after the equal sign")
		os.Exit(0)
	}
	return (m)
}

func calculate_degre_polynome(m map[float64]float64) float64 {
	var max float64

	max = 0
	for key, element := range m {
		if (max < key && element != 0) {
			max = key
		}
		if (key / float64(int(key)) > 1 || key / float64(int(key)) < 1) {
			return (-1)
		}
	}
	return (max)
}

func trier_map(m map[float64]float64) map[float64]float64 {
	var min float64
	m_return := make(map[float64]float64)
	a := 0
	b := 0
	pass := 0

	fmt.Print("Reduced form: ")
	for len(m) > 0 {
		a = 0
		for key, _ := range m {
			if (a == 0) {
				a = 1
				min = key
			} else if (min > key) {
				min = key
			}
		}
		m_return[min] = m[min]
		if (m[min] < 0) {
			fmt.Print("- ")
			b = 2
			fmt.Print(-m[min], " * X^", min, " ")
			pass++
		} else if (m[min] > 0) {
			if (b == 0) {
				fmt.Print("")
			} else {
				 fmt.Print("+ ")
			}
			b = 2
			fmt.Print(m[min], " * X^", min, " ")
			pass++
		}
		delete(m, min)
	}
	if (pass != 0) {
		fmt.Println("= 0")
	} else {
		fmt.Println("0 * X^0 = 0")
	}
	return(m_return)
}

func sqrt(x float64) float64 {
	i := 0
	z := 1.0

	for i < 20 {
		z -= (z*z - x) / (2*z)
		i++
	}
	return (z)
}

func main() {
	inter := false

	if len(os.Args) < 2 {
		fmt.Println("usage : ./computerV1 name_polynome -inter=bool")
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		if (strings.Compare("-inter", os.Args[2]) == 0) {
			inter = true
		}
	}
	map_polynome := transform_polynome(os.Args[1])
	map_polynome = trier_map(map_polynome)
	degre_poly := calculate_degre_polynome(map_polynome)
	if (degre_poly == -1) {
		fmt.Println("syntax error : a power can't be with a comma")
		os.Exit(0)
	} else if (degre_poly > 2 ) {
		fmt.Println("Polynomial degree:", degre_poly)
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		os.Exit(0)
	}
	fmt.Println("Polynomial degree:", degre_poly)
	if (degre_poly == 0) {
		if (inter == true) {
			fmt.Println("Our polynomial is of the form : A = 0")
			fmt.Println("With A equal to", map_polynome[0])
		}
		if (map_polynome[0] == 0) {
			fmt.Println("All the reals are solutions of the equation")
		} else {
			fmt.Println("There is no solution")
		}
	} else if (degre_poly == 1){
		if (inter == true) {
			fmt.Println("Our polynomial is of the form : A*X + B = 0")
			fmt.Println("With A equal to", map_polynome[1], "and B equal to", map_polynome[0])
			fmt.Println("The solution of this equation is X = -B / A")
		}
		fmt.Println("The solution is:")
		fmt.Println(-map_polynome[0] / map_polynome[1])
	} else {
		if (inter == true) {
			fmt.Println("Our polynome is of the form : A*X^2 + B*X + C = 0")
			fmt.Println("With A equal to", map_polynome[2], ", B equal to", map_polynome[1], "and C equal to", map_polynome[0])
		}
		delta := map_polynome[1] * map_polynome[1] - 4 * map_polynome[2] * map_polynome[0]
		if (inter == true) {
			fmt.Println("We have the discriminant (= B^2 - 4AC) equal to", delta)
		}
		if (delta > 0) {
			fmt.Println("Discriminant is strictly positive, the two solutions are:")
			fmt.Println((-map_polynome[1] - sqrt(delta)) / (2 * map_polynome[2]))
			fmt.Println((-map_polynome[1] + sqrt(delta)) / (2 * map_polynome[2]))
		} else if (delta == 0) {
			fmt.Println("Discriminant is equal to zero, the solution is:")
			fmt.Println(-map_polynome[1] / (2 * map_polynome[2]))
		} else {
			fmt.Println("Discriminant is strictly negative, the two solutions are:")
			fmt.Print(-map_polynome[1] / (2 * map_polynome[2]), " ")
			if (-sqrt(-delta) / (2 * map_polynome[2]) > 0) {
				fmt.Print("+ ")
				fmt.Print(-sqrt(-delta) / (2 * map_polynome[2]), "i")
			} else {
				fmt.Print("- ")
				fmt.Print(sqrt(-delta) / (2 * map_polynome[2]), "i")
			}
			fmt.Println()
			fmt.Print(-map_polynome[1] / (2 * map_polynome[2]), " ")
			if (sqrt(-delta) / (2 * map_polynome[2]) > 0) {
				fmt.Print("+ ")
				fmt.Print(sqrt(-delta) / (2 * map_polynome[2]), "i")
			} else {
				fmt.Print("- ")
				fmt.Print(-sqrt(-delta) / (2 * map_polynome[2]), "i")
			}
			fmt.Println()
		}
	}
}
