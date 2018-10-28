package _752_Open_the_Lock

/*https://leetcode.com/problems/open-the-lock/

You have a lock in front of you with 4 circular wheels. Each wheel has 10 slots: '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'. The wheels can rotate freely and wrap around: for example we can turn '9' to be '0', or '0' to be '9'. Each move consists of turning one wheel one slot.

The lock initially starts at '0000', a string representing the state of the 4 wheels.

You are given a list of deadends dead ends, meaning if the lock displays any of these codes, the wheels of the lock will stop turning and you will be unable to open it.

Given a target representing the value of the wheels that will unlock the lock, return the minimum total number of turns required to open the lock, or -1 if it is impossible.

Example 1:
Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
Output: 6
Explanation:
A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
because the wheels of the lock become stuck after the display becomes the dead end "0102".
Example 2:
Input: deadends = ["8888"], target = "0009"
Output: 1
Explanation:
We can turn the last wheel in reverse to move from "0000" -> "0009".
Example 3:
Input: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
Output: -1
Explanation:
We can't reach the target without getting stuck.
Example 4:
Input: deadends = ["0000"], target = "8888"
Output: -1
Note:
The length of deadends will be in the range [1, 500].
target will not be in the list deadends.
Every string in deadends and the string target will be a string of 4 digits from the 10,000 possibilities '0000' to '9999'.
*/

var queue []string

func enqueue(e string) {
    queue = append(queue, e)
}

func dequeue() string {
    if len(queue) == 0 {
        return ""
    }

    e := queue[0]
    queue = queue[1:]
    return e
}

func change(e string, char int, up bool) string {
    // 0 -> 48
    // 9 -> 57
    modch := func(b byte) byte {
        if up {
            return 48 + (b-48+1)%10
        }
        return 48 + (b-48-1)%10
    }
    switch char {
    case 0:
        return string([]byte{modch(e[0]), e[1], e[2], e[3]})
    case 1:
        return string([]byte{e[0], modch(e[1]), e[2], e[3]})
    case 2:
        return string([]byte{e[0], e[1], modch(e[2]), e[3]})
    case 3:
        return string([]byte{e[0], e[1], e[2], modch(e[3])})
    }
    return ""
}

func openLock(deadends []string, target string) int {
    de := map[string]bool{}
    for _, d := range deadends {
        de[d] = true
    }
    used := map[string]bool{}
    level := 0
    enqueue("0000")
    for len(queue) > 0 {
        //fmt.Println(len(queue))
        size := len(queue)
        for size > 0 {
            e := dequeue()
            if e == target {
                return level
            }

            if de[e] || used[e] {
                size--
                continue
            }

            used[e] = true
            for i := 0; i < 4; i++ {
                up := change(e, i, true)
                if !used[up]{
                    enqueue(up)
                }
                down := change(e, i, false)
                if !used[down]{
                    enqueue(down)
                }
            }
            size--
        }
        level++
    }
    return -1
}
