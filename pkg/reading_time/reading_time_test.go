package reading_time

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeToRead(t *testing.T) {
	const input1 = "this is some sample text"
	result1 := TimeToRead(input1)
	assert.Equal(t, "1 second read", result1)

	const input2 = "This is 280 words of random, whimsical text that dances across the page like a mischievous pixie. Flummox the quixotic zephyr and juxtapose the kerfuffle, while absquatulating the mellifluous serendipity. Obfuscate the syzygy and discombobulate the onomatopoeia, as you defenestrate the hullabaloo with lackadaisical brouhaha. Persnickety tintinnabulation lollygags through gobbledygook, while balderdash canoodles and skedaddles rambunctiously. Flibbertigibbet cattywampus widdershins poppycock bamboozles the gardyloo, taradiddle collywobbles, and snollygoster jackanapes mollycoddle. Folderol rannygazoo skullduggery wabbit klutzes donnybrook slangwhanger comeuppance pronk bivouac. Rapscallion billingsgate blunderbuss hornswoggle pandiculation codswallop sockdolager snool widdiful snickersnee wafture furbelow gasconade. Scaramouch quidnunc limerance gallivants rigmarole fartlek sesquipedalian diphthong oxymoron zigzag whippersnapper kerfuffle filibuster flabbergast. Malarkey snafu quagmire akimbo didgeridoo foofaraw flapdoodle wabbit bumbershoot curmudgeon tatterdemalion flibbertigibbet widdershins hootenanny. Mnemonic scuttlebutt tommyrot whirligig cockalorum sialoquent lollygag brouhaha mudlark flummox discombobulate blatherskite jackanapes kerfuffle canoodle. Cattywampus skedaddle fartlek widdiful snollygoster taradiddle collywobbles folderol rannygazoo skullduggery wabbit klutz donnybrook slangwhanger comeuppance. Pronk bivouac rapscallion billingsgate blunderbuss hornswoggle pandiculation codswallop sockdolager snool widdiful snickersnee wafture furbelow. Let the words flow like a babbling brook, tumbling and frolicking through a verdant meadow of imagination. Embrace the whimsy and let your mind wander to the beat of its own peculiar drum. "
	result2 := TimeToRead(input2)
	assert.Equal(t, "1 minute read", result2)

	var builder strings.Builder

	// Write strings to the builder
	builder.WriteString(input2)
	builder.WriteString(input2)
	builder.WriteString(input2)
	builder.WriteString(input2)

	input3 := builder.String()
	words := strings.Fields(input3)
	fmt.Printf("The string contains %d words.\n", len(words))

	// Get the concatenated string
	result3 := TimeToRead(input3)
	assert.Equal(t, "3 minute read", result3)
}
