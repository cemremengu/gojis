package lexer

import (
	"testing"

	"github.com/gojisvm/gojis/internal/parser/token"
	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	require := require.New(t)

	s1 := func(l *Lexer) state {
		require.Equal(0, l.start)
		require.Equal(0, l.pos)

		require.Equal('\u007f', l.next())

		require.Equal(0, l.start)
		require.Equal(1, l.pos)

		require.Equal('\u07ff', l.next())

		require.Equal(0, l.start)
		require.Equal(2, l.pos)

		require.Equal('\u7fff', l.next())

		require.Equal(0, l.start)
		require.Equal(3, l.pos)

		require.True(l.eof())
		require.Equal(0, l.start)
		require.Equal(3, l.pos)

		return nil
	}
	//___________________________ 1byte 2byte 3byte
	testWithDataAndState([]byte("\u007f\u07ff\u7fff"), s1)
}

func TestAcceptWord(t *testing.T) {
	require := require.New(t)

	s1 := func(l *Lexer) state {
		require.Equal(0, l.start)
		require.Equal(0, l.pos)

		require.True(l.acceptWord("hello"))

		require.Equal(0, l.start)
		require.Equal(5, l.pos)

		l.emit(token.Unknown)

		require.Equal(5, l.start)
		require.Equal(5, l.pos)

		require.False(l.acceptWord("world"))

		require.Equal(5, l.start)
		require.Equal(5, l.pos)

		require.True(l.acceptWord("\u7fff"))

		require.Equal(5, l.start)
		require.Equal(6, l.pos)

		l.emit(token.Unknown)

		require.Equal(6, l.start)
		require.Equal(6, l.pos)

		require.True(l.acceptWord("abc"))

		require.Equal(6, l.start)
		require.Equal(9, l.pos)

		l.emit(token.Unknown)

		require.Equal(9, l.start)
		require.Equal(9, l.pos)

		require.True(l.eof())

		return nil
	}
	testWithDataAndState([]byte("hello\u7fffabc"), s1)
}

func testWithDataAndState(data []byte, initial state) {
	l := newWithInitialState(data, initial)

	go func() {
		err := l.StartLexing()
		if err != nil {
			panic(err)
		}
	}()

	for range l.TokenStream().Tokens() {
		// drain all tokens
	}
}