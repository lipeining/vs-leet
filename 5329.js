/**
 * @param {number[]} arr
 * @return {number}
 */
var minSetSize = function (arr) {
    const counter = {}
    arr.forEach(item => {
        if (!counter[item]) {
            counter[item] = 0
        }
        counter[item]++
    })

    const length = arr.length
    const half = Math.floor(length / 2)
    const keys = Object.keys(counter)
    let ans = arr.length
    const dfs = (counter, keys, index, path, sum) => {
        if (sum >= half) {

            if (path.length < ans) {
                console.log('path', path)
                ans = path.length
            }
            return
        }
        for (let i = index; i < keys.length; i++) {
            path.push(keys[i])
            dfs(counter, keys, i + 1, path, sum + counter[keys[i]])
            path.pop()
            // dfs(counter, keys, i, [...path, keys[i]], sum+counter[keys[i]])
        }
    }

    keys.sort((a, b) => {
        return counter[b] - counter[a]
    })
    console.log(counter, keys, half, ans)
    // dfs(counter, keys, 0, [], 0)
    // return ans
    // 下面是大佬的做法：
    let sum = 0
    let j = 0
    // 只是贪心算法就足够了
    for (let i = 0; i < keys.length; i++) {
        sum += counter[keys[i]]
        j++
        if (sum >= half) {
            break
        }
    }
    console.log('result', j)
    return j
};

function test() {
    const arr = [9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19];
    minSetSize(arr)
}

test()
