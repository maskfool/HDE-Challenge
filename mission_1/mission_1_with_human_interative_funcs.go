package main

import (
	"fmt"
)

func main() {
	var n_cases int

    Num_test_cases:
	fmt.Printf("please input the number of test cases\n")
	if    _, err := fmt.Scan(&n_cases);    err != nil {
		fmt.Printf("the input should be an integer\n")
		goto Num_test_cases
	}else if (n_cases < 1 || n_cases > 100) {
		fmt.Printf("the input should be between 1 and 100\n")
		goto Num_test_cases
	}

	fmt.Printf("your selected number of test cases is \n")
	fmt.Println(n_cases)

	 args := make([][]int, n_cases)

	counter := 0
	process_label:
	num_input := num_args_input_HI()
	args_temp := process_args_HI(num_input)
	args[counter] = args_temp[:]
	counter++
	if counter >= n_cases{
		fmt.Printf("a comprehensive output:\n")
		fmt.Printf("all arguments after excepting negatives: \n")
		fmt.Print(args)
		fmt.Printf("\n")
		fmt.Printf("sum_square results:\n")
		sum_sq_calculate_HI(args, n_cases)

		return
	}
	goto process_label

}

func num_args_input_HI()(i int){

    Num_args_label:
	fmt.Printf("please input the number of arguments\n")

	if    _, err := fmt.Scan(&i);    err != nil {
		fmt.Printf("the input should be an integer\n")
		goto Num_args_label

	}else if (i < 0 || i > 100){
		fmt.Printf("the input should be between 0 and 100\n")
		goto Num_args_label

	}else{
		fmt.Printf("your selected number of arguments is ")
		fmt.Println(i)
	}
	return
}

func process_args_HI(i int)(arg [] int){
	if i == 0{
		arg = append(arg, 0)
		return
	}

	var scan_val_temp, sq_sum int


	fmt.Printf("please input the arguments\n")
	j:= 0
    args_input_label:

	if    _, err := fmt.Scan(&scan_val_temp);    err != nil {
		fmt.Printf("the input should be an integer\n")
		goto args_input_label
	}

	if scan_val_temp < -100 || scan_val_temp > 100 {
		fmt.Printf("the input should be with [-100 100]\n")
		goto args_input_label
	}

	j++
	if j == i && scan_val_temp < 0{
		fmt.Printf("arguments after excepting negatives: \n")
		fmt.Println(arg)

		fmt.Printf("sum of all inputs' square : \n")
		fmt.Println(sq_sum)

		return
	}

	if scan_val_temp < 0 {
		goto args_input_label
	}

	arg = append(arg, scan_val_temp)

	sq_sum += scan_val_temp * scan_val_temp


	if j < i {
		goto args_input_label
	}
	fmt.Printf("arguments after excepting negatives: \n")
	fmt.Println(arg)

	fmt.Printf("sum of all input's square : \n")
	fmt.Println(sq_sum)

	return
}

func sum_sq_calculate_HI(args [][] int, n_cases int){

	sq_sum := make([]int, n_cases)


	counter := 0
	each_input_row:
	length_each_input_row := len(args[counter])
	counter_j := 0
	each_input_calculate:
	sq_sum[counter] = args[counter][counter_j] * args[counter][counter_j] + sq_sum[counter]
	counter_j++
	if counter_j < length_each_input_row {
		goto each_input_calculate
	}else{
		fmt.Print(sq_sum[counter])
		counter++
		if counter >= n_cases{

			return
		}
		goto each_input_row
		}

	return
}