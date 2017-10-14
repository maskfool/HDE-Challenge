package main

import (
	"fmt"
)

func main() {
	var n_cases int

Num_test_cases:

	if    _, err := fmt.Scan(&n_cases);    err != nil {
		fmt.Printf("\n")
		goto Num_test_cases
	}else if (n_cases < 1 || n_cases > 100) {

		goto Num_test_cases
	}

	args := make([][]int, n_cases)

	counter := 0
process_label:
	num_input := num_args_input()
	args_temp := process_args(num_input)
	args[counter] = args_temp[:]
	counter++
	if counter >= n_cases{
		sum_sq_calculate(args, n_cases)

		return
	}
	goto process_label

}

func num_args_input()(i int){

Num_args_label:


	if    _, err := fmt.Scan(&i);    err != nil {
		fmt.Printf("\n")
		goto Num_args_label

	}else if (i < 0 || i > 100){
		fmt.Printf("\n")
		goto Num_args_label

	}
	return
}

func process_args(i int)(arg [] int){
	if i == 0{
		arg = append(arg, 0)
		return
	}

	var scan_val_temp, sq_sum int



	j:= 0
args_input_label:

	if    _, err := fmt.Scan(&scan_val_temp);    err != nil {
		fmt.Printf("\n")
		goto args_input_label
	}

	if scan_val_temp < -100 || scan_val_temp > 100 {
		fmt.Printf("\n")
		goto args_input_label
	}

	j++
	if j == i && scan_val_temp < 0{
		fmt.Printf("\n")
		fmt.Println(arg)

		fmt.Printf("\n")
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

	return
}

func sum_sq_calculate(args [][] int, n_cases int){

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
		fmt.Printf("\n")
		counter++
		if counter >= n_cases{

			return
		}
		goto each_input_row
	}

	return
}
