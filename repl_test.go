package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanInput(t *testing.T) {
    t.Run("empty string", func(t *testing.T) {
        
        var input string

        cleanedWords := cleanText(input)

        assert.Empty(t, cleanedWords)
     
    })

    t.Run("should lowercase uppercase", func(t *testing.T) {

        input := "Bulbasaur"

        cleanedWords := cleanText(input)

        require.Len(t, cleanedWords, 1)
        assert.Equal(t, "bulbasaur", cleanedWords[0])
    })
}
