package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var papan [][]byte
var n int
var queenpos []int
var count int
var anim bool 

func main() {
	fmt.Print("Masukkan nama file : ")
	scannerInput := bufio.NewScanner(os.Stdin)
	if scannerInput.Scan() {
		filename := scannerInput.Text()
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Gagal membuka file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		papan = bacapapan(scanner)
	}

	fmt.Print("Tampilkan animasi visualisasi? (Ya/Tidak): ")
	animScanner := bufio.NewScanner(os.Stdin)
	if animScanner.Scan() {
		animJawaban := strings.TrimSpace(animScanner.Text())
		if strings.EqualFold(animJawaban, "Ya") {
			anim = true
		} else {
			anim = false
		}
	}

	start := time.Now()
	found := solve()
	end := time.Since(start)

	if anim {
		fmt.Print("\033[H\033[2J")
	}
	
	printpapan()
	if found {
		fmt.Printf("Waktu pencarian: %v\n", end)
		fmt.Printf("Banyak kasus yang ditinjau: %d kasus\n", count)
	} else {
		fmt.Println("Tidak ada solusi")
		fmt.Printf("Waktu pencarian: %v\n", end)
		fmt.Printf("Banyak kasus yang ditinjau: %d kasus\n", count)
	}

	fmt.Print("Apakah Anda ingin menyimpan solusi? (Ya/Tidak) ")
	reader := bufio.NewReader(os.Stdin)
	jawaban, _ := reader.ReadString('\n')
	jawaban = strings.TrimSpace(jawaban)

	if strings.EqualFold(jawaban, "Ya") {
		fmt.Print("Masukkan nama file output : ")
		outScanner := bufio.NewScanner(os.Stdin)
		if outScanner.Scan() {
			outName := outScanner.Text()
			simpan(found, end, outName)
		}
	}
}

func bacapapan(scanner *bufio.Scanner) [][]byte {
	var totalhuruf []byte

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			totalhuruf = append(totalhuruf, []byte(line)...)
		}
	}

	total := len(totalhuruf)
	root := int(math.Sqrt(float64(total)))

	if root < 4 {
		fmt.Println("Tidak valid: Ukuran papan minimal 4x4.")
		os.Exit(1)
	}

	n = root
	queenpos = make([]int, n)

	b := make([][]byte, n)
	index := 0
	for i := 0; i < n; i++ {
		b[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			b[i][j] = totalhuruf[index]
			index++
		}
	}

	return b
}

func solve() bool {
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}

	for {
		count++
		if anim {
			copy(queenpos, perm)
			printLive()
		}

		if valid(perm) {
			copy(queenpos, perm)
			return true
		}
		if !permutasi(perm) {
			break
		}
	}

	return false
}

func valid(perm []int) bool {
	used := make(map[byte]bool)

	for i := 0; i < n; i++ {
		row := i
		col := perm[i]

		region := papan[row][col]
		if used[region] {
			return false
		}
		used[region] = true

		for j := 0; j < i; j++ {
			if mutlak(row-j) <= 1 && mutlak(col-perm[j]) <= 1 {
				return false
			}
		}
	}

	return true
}

func permutasi(a []int) bool {
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}

	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}

	a[i], a[j] = a[j], a[i]
	reverse(a[i+1:])
	return true
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func printLive() {
	fmt.Print("\033[H\033[2J")
	printpapan()
	fmt.Printf("Mencoba konfigurasi ke: %d\n", count)
	time.Sleep(10 * time.Millisecond)
}

func printpapan() {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if queenpos[i] == j {
				fmt.Print("#")
			} else {
				fmt.Printf("%c", papan[i][j])
			}
		}
		fmt.Println()
	}
}

func simpan(found bool, end time.Duration, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Gagal menyimpan file:", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	if found {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if queenpos[i] == j {
					fmt.Fprint(writer, "#")
				} else {
					fmt.Fprintf(writer, "%c", papan[i][j])
				}
			}
			fmt.Fprintln(writer)
		}
	} else {
		fmt.Fprintln(writer, "Tidak ada solusi")
	}

	fmt.Fprintf(writer, "Waktu pencarian: %v\n", end)
	fmt.Fprintf(writer, "Banyak kasus yang ditinjau: %d kasus\n", count)
	writer.Flush()

	fmt.Printf("Solusi berhasil disimpan ke %s\n", filename)
}

func mutlak(x int) int {
	if x < 0 {
		return -x
	}
	return x
}