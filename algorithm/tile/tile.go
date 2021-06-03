package tile

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
	Z int
}

const (
	Max = 6
	Min = 1
)

// 箱子平铺算法
func TileAlg(positions []*Position) *Position {
	weights := make([]float64, 0, len(positions))
	for _, position := range positions {
		// 同轨时，Y都相等，不需要做考虑
		weight := convertXAndZ2float64(position.X, position.Z)
		weights = append(weights, weight)
	}
	sort.Float64s(weights)
	z := 1
	x := -1
	// 采用左右指针，分段从第一层开始遍历weights
	length := len(weights)
	left := 0
	// right为每一轨最多能存储的箱数
	maxStore := Max - Min
	right := maxStore - 1
	if right >= length {
		right = length - 1
	}
	for {
		startX, startZ := convertFloat64ToXAndZ(weights[left])
		if startX != Min {
			// 第一个库位没有箱子，提前返回
			x = Min
			break
		}
		endX, endZ := convertFloat64ToXAndZ(weights[right])
		if startZ == endZ && endX-startX == maxStore-1 {
			// 说明该层已经铺满，进入下一层
			left = right
			right += right
			z++
			if right >= length {
				right = length - 1
			}
			continue
		}
		if startZ != endZ {
			// 如果该轨道的已存储的库位的startZ不等与endZ，说明该层还没有铺满
			for {
				// 定位到endZ等于StartZ的地方
				right--
				_, endZ = convertFloat64ToXAndZ(weights[right])
				if startZ == endZ {
					break
				}
			}
		}
		// layerWeights := weights[left:right+1]
		// fmt.Println("[TileStorageStrategy]" + fmt.Sprintf("开始遍历%d层, %v", startZ, layerWeights))
		// 从左开始遍历找到空位
		lastX := startX
		for {
			left++
			nextX, _ := convertFloat64ToXAndZ(weights[left])
			if nextX-lastX > 1 {
				// 在两个点之间存在空位
				x = lastX + 1
				break
			}
			if left == right {
				// 已经到了该层的最后一个箱子
				x = nextX + 1
				break
			}
			lastX = nextX
		}
		if x != -1 {
			break
		}
	}
	if x == -1 {
		// 没有找到可存储的库位
		fmt.Println("can not find suitable storage node on track")
		return nil
	}
	return &Position{
		X: x,
		Y: 1,
		Z: z,
	}
}

func MultiTileAlg(positions []*Position, num int) []*Position {
	result := make([]*Position, 0, num)
	if len(positions) == 0 {
		lastX := Min - 1
		z := 1
		for len(result) < num {
			lastX += 1
			result = append(result, &Position{
				X: lastX,
				Y: 1,
				Z: z,
			})
			if lastX == Max-1 {
				lastX = Min - 1
				z += 1
			}
		}
		return result
	}
	weightsMap := make(map[float64]*Position)
	for _, position := range positions {
		// 同轨时，Y都相等，不需要做考虑
		weight := convertXAndZ2float64(position.X, position.Z)
		weightsMap[weight] = position
	}
	x := Min
	z := 1
	for len(result) < num {
		if x == Max-1 {
			x = Min
			z += 1
		}
		weight := convertXAndZ2float64(x, z)
		position := weightsMap[weight]
		if position == nil {
			result = append(result, &Position{
				X: x,
				Y: 1,
				Z: z,
			})
		}
		x++
	}
	return result
}

func convertXAndZ2float64(x, z int) float64 {
	// z在前，x在后，优先排序高度
	return float64(z) + (float64(x) / float64(1000))
}
func convertFloat64ToXAndZ(num float64) (int, int) {
	z := int(num) % 32
	x, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%v", num), ".")[1])
	return x, z

}
