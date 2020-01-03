package arrays

//學習檢驗用,先註解起來
/* func main() {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Println(got)
} */
/*
func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		fmt.Println("numbers: %v", numbers)
		fmt.Println("numbersToSum: %v", numbersToSum)
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			//取到从索引 1 到最后一个元素
			tail := numbers[1:]
			fmt.Println("tail: %v", tail)
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
*/
