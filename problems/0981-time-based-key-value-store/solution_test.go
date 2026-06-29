package timebasedkeyvaluestore

import "testing"

// assertGet checks Get(key, ts) returns want. Design problem (a type with
// methods), so this is an operation-sequence test rather than the registry.
func assertGet(t *testing.T, tm *TimeMap, key string, ts int, want string) {
	t.Helper()
	if got := tm.Get(key, ts); got != want {
		t.Errorf("Get(%q, %d) = %q, want %q", key, ts, got, want)
	}
}

func TestTimeMap(t *testing.T) {
	t.Run("leetcode example", func(t *testing.T) {
		tm := Constructor()
		tm.Set("foo", "bar", 1)
		assertGet(t, &tm, "foo", 1, "bar")
		assertGet(t, &tm, "foo", 3, "bar") // no set at 3 -> latest <= 3 is ts=1
		tm.Set("foo", "bar2", 4)
		assertGet(t, &tm, "foo", 4, "bar2")
		assertGet(t, &tm, "foo", 5, "bar2")
	})

	t.Run("query before earliest timestamp", func(t *testing.T) {
		tm := Constructor()
		tm.Set("k", "v", 10)
		assertGet(t, &tm, "k", 5, "")  // earlier than any set
		assertGet(t, &tm, "k", 9, "")  // still earlier
		assertGet(t, &tm, "k", 10, "v") // exact
		assertGet(t, &tm, "k", 11, "v") // later
	})

	t.Run("unknown key", func(t *testing.T) {
		tm := Constructor()
		tm.Set("a", "1", 1)
		assertGet(t, &tm, "b", 100, "")
	})

	t.Run("multiple keys are independent", func(t *testing.T) {
		tm := Constructor()
		tm.Set("a", "a1", 1)
		tm.Set("b", "b1", 2)
		tm.Set("a", "a2", 3)
		assertGet(t, &tm, "a", 2, "a1") // a has a1@1, a2@3 -> at ts=2 it's a1
		assertGet(t, &tm, "a", 3, "a2")
		assertGet(t, &tm, "b", 2, "b1")
		assertGet(t, &tm, "b", 1, "") // b's first set is at ts=2
	})

	t.Run("binary search picks largest timestamp <= query", func(t *testing.T) {
		tm := Constructor()
		tm.Set("k", "v5", 5)
		tm.Set("k", "v10", 10)
		tm.Set("k", "v15", 15)
		tm.Set("k", "v20", 20)
		tm.Set("k", "v25", 25)
		assertGet(t, &tm, "k", 4, "")    // before first
		assertGet(t, &tm, "k", 5, "v5")  // exact first
		assertGet(t, &tm, "k", 12, "v10") // between -> floor
		assertGet(t, &tm, "k", 15, "v15") // exact middle
		assertGet(t, &tm, "k", 24, "v20") // between
		assertGet(t, &tm, "k", 100, "v25") // after last
	})

	t.Run("same timestamp overwrite is not required; distinct ts", func(t *testing.T) {
		tm := Constructor()
		tm.Set("x", "first", 1)
		tm.Set("x", "second", 2)
		assertGet(t, &tm, "x", 1, "first")
		assertGet(t, &tm, "x", 2, "second")
	})
}
