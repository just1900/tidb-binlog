package bitmap

// Bitmap used to check slices all fetched
// unsupported concurrency
type Bitmap struct {
	Value   []uint8
	Total   int
	Current int
}

// NewBitmap returns Bitmap instance
func NewBitmap(total int) *Bitmap {
	if total == 0 {
		return nil
	}

	bm := &Bitmap{
		Value: make([]uint8, (total+7)/8),
		Total: total,
	}

	// mask useless bit
	if total%8 > 0 {
		bm.Value[len(bm.Value)-1] = 0xFF ^ (1<<uint(total%8) - 1)
	}
	return bm
}

// Set sets a bit in bitmap
func (b *Bitmap) Set(index int) bool {
	mask := uint8(1 << uint(index%8))
	bucket := b.Value[index/8]

	if bucket|mask != bucket {
		b.Value[index/8] = bucket | mask
		b.Current++
		return true
	}

	return false
}

// Completed checks whether bitmap is completed (all bits set)
func (b *Bitmap) Completed() bool {
	return b.Current == b.Total
}