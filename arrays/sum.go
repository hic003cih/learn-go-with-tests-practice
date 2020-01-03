package arrays

/* func Sum(numbers [5]int) (sum int) {
	sum = 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
*/

//修改為任意size的傳入值

/* func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
*/

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//使用sum將每個numbers加總並返回
/* func SumAll(numbersToSum ...[]int) (sum []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sum[i] = Sum(numbers)
	}

	return
} */

//使用append把切片追加新值
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

/* func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	fmt.Println(sums)
	return sums
} */

// ...表示傳入多個 int array
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			//將0 存在sum array的後面
			sums = append(sums, 0)
		} else {
			//取到从索引 1 到最后一个元素
			tail := numbers[1:]
			//將總和出來的array存在sum array的後面
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
