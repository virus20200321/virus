package problems

import "fmt"

/*
void findSolutions(n, other params) :
    if (found a solution) :
        solutionsFound = solutionsFound + 1;
        displaySolution();
        if (solutionsFound >= solutionTarget) :
            System.exit(0);
        return

    for (val = first to last) :
        if (isValid(val, n)) :
            applyValue(val, n);
            findSolutions(n+1, other params);
            removeValue(val, n);
*/
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	combinationSumHelper(candidates, 0, target, []int{}, &result)
	return result
}

func combinationSumHelper(candidates []int, n int, target int, curNums []int, result *[][]int) {

	if target == 0 {
		*result = append(*result, curNums)
		fmt.Println("found", curNums)
		return
	}

	for i := n; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		curNums = append(curNums, candidates[i])
		combinationSumHelper(candidates, i, target-candidates[i], curNums, result)
		curNums = curNums[:len(curNums)-1]
	}
}
