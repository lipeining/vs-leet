/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */
/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function (root) {
    if (root == null) {
        return root
    }
    let queue = []
    queue.push(root)
    while (queue.length != 0) {
        const length = queue.length
        for (let i = 0; i < length; i++) {
            const prev = queue[i]
            if (i != length - 1) {
                prev.next = queue[i + 1]
            } else {
                prev.next = null
            }
            if (prev.left) {
                queue.push(prev.left)
            }
            if (prev.right) {
                queue.push(prev.right)
            }
        }
        // console.log(length, queue)
        queue.splice(0, length)
        // console.log(length, queue)
    }
    return root
};