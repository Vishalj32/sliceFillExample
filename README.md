## Efficiently Filling Arrays and Slices in Go: A Performance Guide

When working on a Go project, you might need to fill a slice or array with a repeating pattern. Recently, I explored different approaches to achieve this efficiently and discovered a powerful way to boost performance. In this article, I’ll walk through various techniques to fill a slice in Go, highlight their performance differences, and explain why one method stands out as significantly faster.

---

### Why Slice Filling Matters

For my toy project, I needed to fill a background buffer with a specific RGB color pattern. Optimizing this operation significantly improved my achievable frame rate. The insights I gained could be useful for anyone working with large datasets, graphics programming, or high-performance applications.

All tests were performed on a buffer of 73,437 bytes, allocated as:

```go
var bigSlice = make([]byte, 73437)
```

Let’s compare three common ways to fill this slice.

---

## 1. Index-based Loop

The simplest way is to iterate through each index and set the value.

```go
func FillSliceIndex(slice []byte, value byte) {
    for j := 0; j < len(slice); j++ {
        slice[j] = value
    }
}

func Benchmark_FillsliceIndex(b *testing.B) {
    slice := make([]byte, 73437)
    for i := 0; i < b.N; i++ {
        FillSliceIndex(slice, 65)
    }
}
```

**Benchmark Results:**

| Name | Executions | Time/Op | Bytes/Op | Allocs/Op |
|------|------------|---------|----------|-----------|
| Benchmark_FillsliceIndex-4 | 48,555 | 23,044 ns/op | 0 B/op | 0 allocs/op |

This approach is straightforward but relatively slow due to the per-element indexing and bounds checking.

---

## 2. Range-based Loop

Using a `range` loop provides a slight performance improvement:

```go
func FillSliceRange(slice []byte, value byte) {
    for j := range slice {
        slice[j] = value
    }
}

func Benchmark_FillsliceRange(b *testing.B) {
    slice := make([]byte, 73437)
    for i := 0; i < b.N; i++ {
        FillSliceRange(slice, 66)
    }
}
```

**Benchmark Results:**

| Name | Executions | Time/Op | Bytes/Op | Allocs/Op |
|------|------------|---------|----------|-----------|
| Benchmark_FillsliceRange-4 | 48,687 | 22,890 ns/op | 0 B/op | 0 allocs/op |

This is faster than the index-based approach but still not optimal.

---

## 3. The Copy Trick (Efficient Method)

The most efficient approach leverages Go's built-in `copy` function to fill the slice incrementally:

```go
func FillSliceCopyTrick(slice []byte, value byte) {
    slice[0] = value
    for j := 1; j < len(slice); j *= 2 {
        copy(slice[j:], slice[:j])
    }
}

func Benchmark_FillsliceCopyTrick(b *testing.B) {
    slice := make([]byte, 73437)
    for i := 0; i < b.N; i++ {
        FillSliceCopyTrick(slice, 67)
    }
}
```

**Benchmark Results:**

| Name | Executions | Time/Op | Bytes/Op | Allocs/Op |
|------|------------|---------|----------|-----------|
| Benchmark_FillsliceCopyTrick-4 | 1,259,342 | 1,025 ns/op | 0 B/op | 0 allocs/op |

The improvement here is dramatic—about 30x faster than the basic loop!

---

## 4. Filling with a Multi-Element Pattern

If you need to fill the slice with a repeating multi-byte pattern, you can adapt the copy trick easily:

```go
func FillSlicePatternCopyTrick(slice []byte, pattern []byte) {
    copy(slice, pattern)
    for j := len(pattern); j < len(slice); j *= 2 {
        copy(slice[j:], slice[:j])
    }
}

func Benchmark_FillslicePatternCopyTrick(b *testing.B) {
    slice := make([]byte, 73437)
    pattern := []byte{0xde, 0xad, 0xbe, 0xef}
    for i := 0; i < b.N; i++ {
        FillSlicePatternCopyTrick(slice, pattern)
    }
}
```

**Benchmark Results:**

| Name | Executions | Time/Op | Bytes/Op | Allocs/Op |
|------|------------|---------|----------|-----------|
| Benchmark_FillslicePatternCopyTrick-4 | 1,271,686 | 926.3 ns/op | 0 B/op | 0 allocs/op |

---

## Summary of Benchmark Results

| Name | Executions | Time/Op | Bytes/Op | Allocs/Op |
|------|------------|---------|----------|-----------|
| Benchmark_FillsliceIndex-10 | 48,555 | 23,044 ns/op | 0 B/op | 0 allocs/op |
| Benchmark_FillsliceRange-10 | 48,687 | 22,890 ns/op | 0 B/op | 0 allocs/op |
| Benchmark_FillsliceCopyTrick-10 | 1,259,342 | 1,025 ns/op | 0 B/op | 0 allocs/op |
| Benchmark_FillslicePatternCopyTrick-10 | 1,271,686 | 926.3 ns/op | 0 B/op | 0 allocs/op |

---

## How and Why the Copy Trick Works

The `copy` function avoids the overhead of indexing and bounds checking each element. Here’s how it works:

- The first value (or pattern) is loaded into the slice.
- Each call to `copy` duplicates twice the amount of data as the previous iteration.
- This exponential growth reduces the number of copy operations required, amortizing the cost.
- The final copy naturally stops when the slice is filled—no bounds checks are needed.

Here’s a visual breakdown:

| Current Data | Action | New Data |
|--------------|--------|----------|
| ['A'] | Copy index 0 to 1 | ['A', 'A'] |
| ['A', 'A'] | Copy first 2 elements to index 2 | ['A', 'A', 'A', 'A'] |
| ['A', 'A', 'A', 'A'] | Copy first 4 elements to index 4 | ['A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'] |

Each iteration doubles the amount of work carried out per call to `copy`, leading to a very efficient fill operation.

---

## Final Thoughts

If you ever need to efficiently fill a slice or array in Go, especially for large datasets, the copy trick is a powerful and elegant solution. It’s faster, avoids unnecessary allocations, and leverages built-in optimizations for block memory operations.

I hope this guide helps you optimize your Go code and improve your application’s performance. Let me know if you discover any other cool slice-filling techniques!

Happy coding! 🚀

