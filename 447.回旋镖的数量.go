/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
// 	int num=0;
// 	int dis=0;
// 	unordered_map<int,int> n;
// 	for(int i=0;i<points.size();++i){
// 		n.clear();
// 		for(int j=0;j<points.size();++j){
// 			if(i!=j){
// 				dis=pow(points[i][0]-points[j][0],2)+pow(points[i][1]-points[j][1],2);
// 				n[dis]++;//同长度线段计数
// 				if (n[dis]>1)
// 					num+=2*(n[dis]-1);//一次加入等长线段,增加2*(n-1)个回旋镖
// 			}
// 		}
// 	}
// 	return num;

// 作者：zhanzhengyu
// 链接：https://leetcode-cn.com/problems/number-of-boomerangs/solution/hui-xuan-biao-de-shu-liang-by-zhanzhengyu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
num:= 0

length := len(points)
for i:=0;i<length;i++ {
	smap := make(map[int]int)
	for j:=0;j<length;j++ {
		if i!= j {
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]
			dis := x*x + y*y
			c,ok := smap[dis]
			if ok {
				smap[dis] = c + 1
				num = num + 2*c
			} else {
				smap[dis] = 1
			}
		}
	}
}
return num

}
// @lc code=end

