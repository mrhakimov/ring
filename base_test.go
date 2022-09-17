package ring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	buf := New()
	for i := 0; i < defaultBufferSize*10; i++ {
		require.Equal(t, i%defaultBufferSize, buf.end)

		buf.Insert(i)

		require.Equal(t, (i+1)%defaultBufferSize, buf.end)

		require.False(t, buf.Empty())
		require.LessOrEqual(t, buf.Size(), defaultBufferSize)
	}
}

func TestPopNewEmpty(t *testing.T) {
	buf := New()

	require.True(t, buf.Empty())
	require.Equal(t, 0, buf.Size())

	require.Equal(t, 0, buf.start)
	require.Equal(t, buf.start, buf.end)
}

func TestPopNonEmpty(t *testing.T) {
	buf := New()

	for i := 0; i < defaultBufferSize*10; i++ {
		buf.Insert(i)
	}

	require.False(t, buf.Empty())
	require.Equal(t, defaultBufferSize, buf.Size())

	for i := 0; i < defaultBufferSize; i++ {
		require.Equal(t, defaultBufferSize*9+i, buf.Pop())
		require.Equal(t, defaultBufferSize-i-1, buf.Size())
	}

	require.True(t, buf.Empty())
}

func TestInsertsPops_LessThanSize(t *testing.T) {
	buf := New()

	for i := 0; i < defaultBufferSize-1; i++ {
		buf.Insert(i)
	}

	require.False(t, buf.Empty())
	require.Equal(t, defaultBufferSize-1, buf.Size())

	require.Equal(t, 0, buf.start)
	require.Equal(t, defaultBufferSize-1, buf.end)

	for i := 0; i < defaultBufferSize-1; i++ {
		require.Equal(t, i, buf.start)
		require.Equal(t, i, buf.Pop())
		require.Equal(t, i+1, buf.start)

		require.Equal(t, defaultBufferSize-1, buf.end)
	}

	require.True(t, buf.Empty())
	require.Equal(t, defaultBufferSize-1, buf.start)
	require.Equal(t, defaultBufferSize-1, buf.end)
}

func TestInsertsPops_MoreThanSize(t *testing.T) {
	buf := New()

	for i := 0; i < defaultBufferSize+1; i++ {
		buf.Insert(i)
	}

	require.False(t, buf.Empty())
	require.Equal(t, defaultBufferSize, buf.Size())

	require.Equal(t, 1, buf.start)
	require.Equal(t, 1, buf.end)

	for i := 0; i < defaultBufferSize; i++ {
		require.Equal(t, (i+1)%defaultBufferSize, buf.start)
		require.Equal(t, i+1, buf.Pop())
		require.Equal(t, (i+2)%defaultBufferSize, buf.start)
	}

	require.True(t, buf.Empty())
	require.Equal(t, 1, buf.start)
	require.Equal(t, 1, buf.end)
}
