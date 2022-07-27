/**
 * @author Emove
 * @date 2021/1/26
 */
package num_equiv_domino_pairs_1128

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := [100]int{}
	ans := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}

//给定一个二维平面及平面上的N个点列表Points，其中第i个点的坐标为Points[i]=[Xi,Yi].
//请找出一条直线，其通过的点的数目最多.
//设穿过最多点的直线所穿过的全部点编号从小到大排序的列表为S.
//你仅需返回[S[0],S[1]]作为答案，若有多条直线穿过了相同数量的点.
//则选择S[0]值较小的直线返回，S[0]相同则选择S[1]值较小的直线返回
