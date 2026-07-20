# 621. Task Scheduler

- Difficulty: medium
- Link: https://leetcode.com/problems/task-scheduler/

## Approach

Two approaches are implemented. `leastInterval` is a greedy simulation: count
each task, push the counts into a max-heap, then process one cooldown frame of
`n+1` slots at a time — pop the most frequent remaining tasks, decrement and
stash their leftovers, and count an idle only when the heap runs dry but
leftovers still exist (so trailing idles are never charged). `leastIntervalMath`
uses the closed-form arrangement `(maxCount-1)*(n+1) + numMax`, where `maxCount`
is the highest task frequency and `numMax` is how many task types hit it; the
answer is floored at `len(tasks)` for the case where tasks pack tightly enough
to need no idles at all. The result is a single scalar, so the test compares
with `==` and guards the heap variant with a timeout in case a bad idle
condition loops forever.

- Time:  O(t + m log m), t = number of tasks, m = distinct task types (<= 26)
- Space: O(m) for the counts / heap
