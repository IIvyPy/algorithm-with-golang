package question

// 递归解法
func myPow(x float64, n int) float64 {
	if x == 0{
		return 0
	}
	if n == 0{
		return 1
	}
	if n == 1{
		return x
	}

	if n < 0{
		return myPow(1/x, -n)
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	}else{
		return myPow(x*x, (n-1)/2)*x
	}
}

// 迭代版本
func myPow2(x float64, n int) float64{
	if x == 0{
		return 0
	}

	if n < 0{
		n = -n
		x = 1/x
	}

	if n == 1{
		return x
	}
	var pow float64 = 1
	for n != 0{
		if n & 1 == 1{
			pow = pow * x
		}
		x = x * x
		n = n >> 1
	}
	return pow
}

