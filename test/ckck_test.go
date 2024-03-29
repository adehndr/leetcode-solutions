package test

import (
	"fmt"
	"math"
	"project/data"
	"project/util"
	"project/util/arraysproblem"
	"project/util/greedy"
	sortingUtil "project/util/sorting"
	"testing"
)

func TestBinarSearch(t *testing.T) {
	anInput := []int{7, 9, 2, 3, 5, 0, 1, 8, 6, 7}
	target := 3
	mergeSort(anInput)
	fmt.Println(anInput)
	fmt.Println(binarSearch(anInput, target))
}

func binarSearch(input []int, target int) (result int) {
	result = -1

	l := 0
	r := len(input) - 1

	for l <= r {
		mid := l + ((r - l) / 2)
		if target > input[mid] {
			l = mid + 1
		} else if target < input[mid] {
			r = mid - 1
		} else {
			result = mid
			return result
		}
	}

	return result
}

func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2

	tmpLeft := []int{}
	tmpRight := []int{}

	tmpLeft = append(tmpLeft, arr[:mid]...)
	tmpRight = append(tmpRight, arr[mid:]...)

	mergeSort(tmpLeft)
	mergeSort(tmpRight)

	sort(arr, tmpLeft, tmpRight)

}

func sort(arr, left, right []int) {
	l := 0
	i := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			arr[i] = left[l]
			l += 1
			i += 1
		} else {
			arr[i] = right[r]
			r += 1
			i += 1
		}
	}
	for l < len(left) {
		arr[i] = left[l]
		l += 1
		i += 1
	}

	for r < len(right) {
		arr[i] = right[r]
		r += 1
		i += 1
	}
}

func TestReceiveChannel(t *testing.T) {
	tmpC := getAChannel(2)
	fmt.Println(<-tmpC)
}

func getAChannel(input int) <-chan int {
	aChan := make(chan int)
	go func() {
		aChan <- input
	}()
	return aChan
}

func TestString(t *testing.T) {
	input := "abcba@"
	input = filter(input)
	fmt.Println(polindrom(input))
}

func filter(input string) (output string) {
	for i := 0; i < len(input); i++ {
		if input[i] < 123 && input[i] > 96 {
			output += string(input[i])
		}
	}
	return output
}

func polindrom(input string) bool {
	length := len(input) / 2
	for i := 0; i < length; i++ {
		fmt.Println(string(input[i]), string(input[len(input)-1-i]))
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}

type TempStruct struct {
	Name string
}

type SpecialStruct interface {
	sum(int) int
	sum2(input int) int
}

func (a *TempStruct) sum(int) int {
	return 0
}

func (a TempStruct) sum2(input int) int {
	return input
}

func receiveTheInterface(a SpecialStruct) {
	a.sum(2)
	a.sum2(52)
}

func TestSendInterface(t *testing.T) {
	a := &TempStruct{Name: "ade"}
	receiveTheInterface(a)
}

func TestTempStruct(t *testing.T) {
	tmpStr := new(TempStruct)
	tmpStr.Name = "Ade"
	fmt.Println(tmpStr, (*tmpStr).Name)
}

func TestLongestPalindrom(t *testing.T) {
	anInput := "civilwartestingwhetherthatnaptionorannartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesaherebutitcanneverforgetwhatthedidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichthewhofoughtherehavethusfarsonobladvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichthegavethelastpfullmeasureofdevotionthatweherehighlresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebthepeopleforthepeopleshallnotperishfromtheearth"
	longestPalindrom(anInput)
}

func longestPalindrom(s string) int {
	aMap := make(map[string]int)
	for _, v := range s {
		if value, exist := aMap[string(v)]; exist {
			tempVal := value
			tempVal += 1
			aMap[string(v)] = tempVal
		} else {
			aMap[string(v)] = 1
		}
	}
	highestOdd := 0
	var keHighestOdd string
	result := 0
	resultAll := 0
	for k, v := range aMap {
		if v%2 == 1 && v > highestOdd {
			highestOdd = v
			keHighestOdd = k
		}
		if v%2 == 1 {
			result += (v - 1)
		} else {
			result += v
		}
		resultAll += v
	}

	if _, ok := aMap[keHighestOdd]; ok {
		result = result + 1
	}
	return result
}

func TestMissingNumbers(t *testing.T) {
	// Generate number from 1 to length of arra
	// Create Map with ke is the int and the value is count
	// if equal, count ++
	// lastl, check if value is 0, if zero return the missing (ke)
	fmt.Println(arraysproblem.MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func TestFindDisappearedNumbers(t *testing.T) {
	fmt.Println(arraysproblem.FindDisappearedNumbers2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(arraysproblem.FindDisappearedNumbersNeetCode([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func TestBuSellStock(t *testing.T) {
	fmt.Println(greedy.BuySellStock([]int{7, 1, 5, 3, 6, 4}))
}

func TestCountBits(t *testing.T) {
	fmt.Println(util.CountBits2(3))
	fmt.Println(util.CountBits3(3))
}

func TestSortedSquareRoot(t *testing.T) {
	input := []int{-5,-1,0,3,9}
	fmt.Println(util.SortedSquareRoot(input))
}

func TestStack(t *testing.T) {
	util.BackspaceString("ab#c", "ad#c")
}

func TestMajoritElement(t *testing.T) {
	fmt.Println(sortingUtil.MajorityElement([]int{2, 1, 1, 1, 1, 2, 2}))
}

func TestFibonnacoMemoizeBottomUp(t *testing.T) {
	fmt.Println(util.FibonnacoMemoizeBottomUp(2))
}

func TestConvert1DTo2D(t *testing.T) {
	anOrig := []int{3}
	m := 1
	n := 2
	fmt.Println(util.Convert1DTo2D(anOrig, m, n))
}

func TestSaveIndexOnList(t *testing.T) {
	aList := []int{0, 1, 0, 2, 0}
	aMap := map[int][]int{}
	for idx, val := range aList {
		if val == 0 {
			aMap[val] = append(aMap[val], idx)
		}
	}
	fmt.Println(aMap)
}

func TestTwoSums(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 22
	fmt.Println(util.TwoSums(nums, target))
}

func TestIsSubsequence(t *testing.T) {
	fmt.Println(util.IsSubsequence("kaaabcefg", "kzzazzazzazzzzazzbzzczzdzzezzfzzgzz"))
}

func TestListNode(t *testing.T) {
	/* 	aNode := &util.ListNode{Val: 1}
	   	aNode.AddNode(&util.ListNode{Val: 3})
	   	aNode.AddNode(&util.ListNode{Val: 6}) */
	aNode := util.InitNode([]int{1, 2, 3, 4, 5})
	aNode.Traverse()
	util.HasCycler(aNode)
}

func TestReverseLinkedList(t *testing.T) {
	aNode := &data.NodeLeetCode{Val: 0}
	tmpNode := aNode
	for i := 1; i < 3; i++ {
		currHead := &data.NodeLeetCode{Val: i}
		aNode.Next = currHead
		aNode = currHead
	}

	bNode := tmpNode
	for bNode != nil {
		fmt.Printf("%d -> ", bNode.Val)
		bNode = bNode.Next
	}
	fmt.Println()

	newTmpNode := util.ReverseList(tmpNode)

	for newTmpNode != nil {
		fmt.Printf("%d -> ", newTmpNode.Val)
		newTmpNode = newTmpNode.Next
	}
	fmt.Println()
	fmt.Println(newTmpNode)

}

func TestRemoveLinkedListElements(t *testing.T) {
	aNode := data.InitNodeLeetCode([]int{7, 7, 7, 7})
	aNode.TraverseNodeLeetCode()
	newNode := util.RemoveLinkedListElements(aNode, 7)
	newNode.TraverseNodeLeetCode()
}

func TestPrintTriagle(t *testing.T) {
	n := 5
	mid := (n / 2) + 1
	for i := 1; i <= n; i++ {
		if i <= mid {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		} else {
			for j := 0; j <= (n - i); j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(
		arraysproblem.SingleNumberBitManipulation([]int{1, 2, 2}),
	)
}

func TestSumUpBruteForce(t *testing.T) {
	a1 := []int{-1,3,8,2,9,5}
	a2 := []int{4,1,2,10,5,20}
	target := 24

	init := a1[0] + a2[0]
	result := make([]int,2)
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			tmp := target - (v1 + v2) 
			if tmp <= (target - init) && tmp >= 0{
				result[0] = v1
				result[1] = v2
				init = v1 + v2
			}
		}
	}
	fmt.Println(result)
}

func TestSumUpEfficient(t *testing.T) {
	a1 := []int{-1,3,8,2,9,5}
	a2 := []int{4,1,2,10,5,20}
	target := 24

	hMap := make(map[int]int)
	for i, v := range a1 {
		hMap[v] = i
	}
	i := 0
	result := make([]int,2)
	outer:
	for {
		counter := int(math.Pow(-1,float64(i)))
		counter *= i
		target += counter

		for idx, v := range a2 {
			if v, ok := hMap[target-v]; ok {
				result[0] = v
				result[1] = idx 
				break outer
			}
		}
		fmt.Println("ini i ke", i, "dengan target", target)
		if i == 2 {
			break
		}
		i += 1
	}
	fmt.Println(result)
}