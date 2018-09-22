package main

import "fmt"

/**

4. Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	m := len(nums1)
	n := len(nums2)
	mi := (m + n) / 2

	if (m + n) % 2 == 0 {//偶数个
		mi -= 1
	}

	i, j := 0, 0

	for i < m && j < n {
		if i + j >= mi {
			break
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	//继续迭代，确保到一半了
	for ; i < m && i + j < mi; i++ {
	}
	for ; j < n && i + j < mi; j++ {
	}

	//已经找到一半了

	r := float64(0)
	if (m + n) % 2 == 0 { //偶数个
		//需要找两个
		if i < m && j < n {
			//获取第一个
			if nums1[i] < nums2[j] {
				r += float64(nums1[i])
				i += 1;
			} else {
				r += float64(nums2[j])
				j += 1
			}
			//获取第二个
			if i < m && j < n {
				if nums1[i] < nums2[j] {
					r += float64(nums1[i])
				} else {
					r += float64(nums2[j])
				}

			} else {
				if i < m {
					r += float64(nums1[i])
				} else {
					r += float64(nums2[j])
				}
			}
		} else {//直接获取两个
			if i < m {
				r = float64(nums1[i] + nums1[i + 1])
			} else {
				r = float64(nums2[j] + nums2[j + 1])
			}
		}

		r /= 2
	} else {//找一个就行
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				r = float64(nums1[i])
			} else {
				r = float64(nums2[j])
			}
		} else {
			if i < m {
				r = float64(nums1[i])
			} else {
				r = float64(nums2[j])
			}
		}
	}

	return r
}

func main()  {

	fmt.Println(findMedianSortedArrays([]int{1}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2,4}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
	fmt.Println(findMedianSortedArrays([]int{1,2,3,4,5}, []int{10}))
	fmt.Println(findMedianSortedArrays([]int{1,2,3,4,5}, []int{0}))
}