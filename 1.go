package main

import "fmt"

type club struct {
	id                  string
	menang, kalah, seri int
}

const NMAX = 123

type arrClub [NMAX]club

func main() {
	var data arrClub
	var n int
	inputData(&data, &n)
	sorting(&data, n)
	print(data, n)
}

func inputData(A *arrClub, n *int) {
	/* I.S. data tim bola sudah tersedia pada piranti masukan
	   F.S. array A berisi sejumlah data n pertandingan */
	var id string
	*n = 0
	fmt.Scan(&id)
	for id != "None" && *n < 123 {
		A[*n].id = id
		fmt.Scan(&A[*n].menang, &A[*n].kalah, &A[*n].seri)
		*n++
		fmt.Scan(&id)
	}
}

func sorting(A *arrClub, n int) {
	/* I.S. terdefinisi array A yang terdiri dari sejumlah n data tim bola
	   F.S. array A terurut mengecil berdasarkan total poin menggunakan algoritma SELECTION SORT ATAU INSERTION SORT*/
	var i, pass int
	var temp club

	for pass = 1; pass <= n-1; pass++ {
		i = pass
		temp = A[pass]
		for i > 0 && (temp.menang*4+temp.seri) > (A[i-1].menang*4+A[i-1].seri) {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
}

func print(A arrClub, n int) {
	/* I.S. terdefinisi array A yang berisi n data tim bola
	   F.S menampilkan 10 tim yang memiliki poin terbesar */
	var i, total int
	i = 0
	for i < 10 && i < n {
		total = A[i].menang*4 + A[i].seri
		fmt.Println(A[i].id, total)
		i++
	}
}
