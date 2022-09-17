package ring

const defaultBufferSize = 10

type Ring struct {
	start int
	end   int
	size  int

	buf []int
}

func New() *Ring {
	return &Ring{
		buf: make([]int, defaultBufferSize),
	}
}

func (r *Ring) Size() int {
	return r.size
}

func (r *Ring) Empty() bool {
	return r.size == 0
}

func (r *Ring) Insert(val int) {
	r.buf[r.end] = val
	r.end++
	r.end %= defaultBufferSize

	r.size++
	if r.size > defaultBufferSize {
		r.size = defaultBufferSize
		r.start++
		r.start %= defaultBufferSize
	}
}

func (r *Ring) Pop() int {
	first := r.buf[r.start]
	r.start++
	r.start %= defaultBufferSize

	r.size--

	return first
}
