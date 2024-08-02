package hitopia

type LargestPalindrom struct {
	InputString string
	K           int
}

func (lp LargestPalindrom) GenerateResult() string {
	PalindromResult := makeLargestPalindrome(lp.InputString, lp.K)
	return PalindromResult
}

func makeLargestPalindrome(s string, k int) string {
	arr := []rune(s)
	changed := make([]bool, len(arr))

	if !makePalindrome(arr, 0, len(arr)-1, k, changed) {
		return "-1"
	}

	maximizePalindrome(arr, 0, len(arr)-1, &k, changed)
	return string(arr)
}

func makePalindrome(arr []rune, left int, right int, k int, changed []bool) bool {
	if left >= right {
		return k >= 0
	}

	if arr[left] != arr[right] {
		if k == 0 {
			return false
		}

		if arr[left] > arr[right] {
			arr[right] = arr[left]
		} else {
			arr[left] = arr[right]
		}

		changed[left] = true
		changed[right] = true
		k--
	}

	return makePalindrome(arr, left+1, right-1, k, changed)
}

func maximizePalindrome(arr []rune, left int, right int, k *int, changed []bool) {
	if left >= right || *k <= 0 {
		return
	}

	if arr[left] < '9' {
		if arr[left] == arr[right] {
			if changed[left] || changed[right] {
				arr[left] = '9'
				arr[right] = '9'
				*k--
			} else if *k >= 2 {
				arr[left] = '9'
				arr[right] = '9'
				*k -= 2
			}
		} else {
			if *k >= 1 {
				arr[left] = '9'
				arr[right] = '9'
				*k--
			}
		}
	}

	maximizePalindrome(arr, left+1, right-1, k, changed)
}
