package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAcceptWord(t *testing.T) {
	require := require.New(t)

	s1 := func(l *Lexer) state {
		require.Equal(0, l.start)
		require.Equal(0, l.pos)
		require.Equal(0, l.width)

		require.True(l.acceptWord("hello"))

		require.Equal(0, l.start)
		require.Equal(5, l.pos)
		require.Equal(1, l.width)

		l.emit()

		require.Equal(5, l.start)
		require.Equal(5, l.pos)
		require.Equal(1, l.width)

		require.False(l.acceptWord("world"))

		require.Equal(5, l.start)
		require.Equal(5, l.pos)
		require.Equal(1, l.width)

		require.True(l.acceptWord("abc"))

		require.Equal(5, l.start)
		require.Equal(8, l.pos)
		require.Equal(1, l.width)

		l.emit()

		require.Equal(8, l.start)
		require.Equal(8, l.pos)
		require.Equal(1, l.width)

		require.True(l.IsEOF())

		return nil
	}

	testWithDataAndState([]byte("helloabc"), s1)
}

func testWithDataAndState(data []byte, initial state) {
	l := newWithInitialState(data, initial)
	go l.StartLexing()

	for range l.TokenStream().Tokens() {
		// drain all tokens
	}
}
