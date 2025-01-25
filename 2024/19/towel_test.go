package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPossibilities(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}

		memo := map[string]int{}
		assert.Equal(t, 1, countPossibilities(memo, towels, "wrr"), "placing wr followed by r or r followed by wr should count as 1 possibility")
		assert.Equal(t, 2, countPossibilities(memo, towels, "brwrr"))
		assert.Equal(t, 1, countPossibilities(memo, towels, "bggr"))
		assert.Equal(t, 4, countPossibilities(memo, towels, "gbbr"))
		assert.Equal(t, 6, countPossibilities(memo, towels, "rrbgbr"))
		assert.Zero(t, countPossibilities(memo, towels, "ubwu"))
		assert.Equal(t, 1, countPossibilities(memo, towels, "bwurrg"))
		assert.Equal(t, 2, countPossibilities(memo, towels, "brgr"))
		assert.Zero(t, countPossibilities(memo, towels, "bbrgwb"))
	})

	t.Run("scenario 1", func(t *testing.T) {
		towels := []string{"wuwuwb", "brwguu", "bbb", "ubb", "rrgw", "rwguuw", "wuwuug", "rwu", "wgb", "urrg", "bburgrgw", "brrwub", "rgwr", "uw", "gbb", "bgggrr", "wuw", "bgb", "gbgwgb", "uu", "gwwu", "bwub", "ugg", "rbr", "gggwuw", "rwgu", "gbbuwugg", "ubwwbg", "gru", "rbw", "uuw", "uwu", "ubuw", "wrg", "grb", "uru", "gur", "bbgg", "bug", "rgrb", "wgur", "br", "bwwgruu", "ubgbw", "ruwg", "buu", "uuwrrgr", "wrw", "bbgr", "ggbg", "burbw", "wgu", "gwgg", "u", "grbw", "rubrwb", "bwu", "wbwruwg", "wwuurg", "wbbrbgu", "rgb", "gwbww", "uruu", "brgwr", "ruru", "grrgub", "bbg", "guuur", "bwbgr", "bggb", "bbguw", "rww", "uwrr", "uwrgr", "ur", "wgrb", "rbb", "rbrub", "rguggb", "gug", "rwb", "wru", "ubbbbub", "urug", "burgur", "urrgbw", "rgr", "wbr", "uuww", "wrr", "urw", "uwurwb", "buw", "wwg", "bgg", "ggw", "gbubgw", "bw", "wrub", "gwg", "rgwbr", "rwbruu", "bwug", "uwbr", "urbbbgw", "ggb", "burrubbu", "wr", "rur", "gbbu", "rrw", "rbbrwgb", "rwwww", "ggurw", "wugwggw", "rwuggg", "brguggb", "brb", "rgur", "bgbw", "bwgr", "rgu", "bugww", "rbwr", "gurrgur", "grwbr", "uug", "bwru", "bb", "rguurbbg", "wububgg", "wwu", "wbw", "ugrw", "ugbwuw", "rwrrgw", "gub", "rbrrguu", "wrgg", "bguuw", "gbugrw", "wuur", "bg", "wwbgr", "rwr", "gwu", "rrbgwgbw", "bur", "wrug", "wrgbu", "gbbw", "rwg", "uwgwb", "bgbbbb", "rrwwuu", "rrgg", "gwgur", "ruww", "grrgrb", "gbuwwg", "ggrb", "uugwgu", "ubbrgu", "rbbru", "wurw", "b", "uwurr", "ubbr", "rgugwbr", "bgbwbgg", "gr", "rwww", "gubgbgr", "bgr", "ug", "grr", "wuuruur", "gw", "rub", "uwggu", "bgw", "uwg", "gwr", "bbrgr", "bbru", "urb", "buuu", "bbwguwb", "ww", "wrbww", "gbwbwrw", "grwr", "gww", "gbr", "uggbgu", "wruu", "brwgu", "gbwgu", "r", "wuwr", "rurr", "ugr", "ubgbr", "ubbwg", "rggrb", "wrggb", "ggggbg", "uruwg", "uwbgb", "bgwwwur", "gbruu", "grrbgb", "rwgbuwrw", "bu", "rwwr", "rgrwru", "uwr", "wrrur", "rug", "rwrwgr", "gwbwbrru", "uggbru", "wrb", "ruwbu", "uugbr", "wbgg", "rbu", "bgwwr", "ggr", "rwbwgr", "rbbg", "buurug", "wgguu", "wrwwg", "gguw", "rrg", "guww", "rgg", "uguwbw", "bgu", "grw", "rgugggw", "uwrw", "rburw", "gbwgur", "wugbubg", "uww", "wbb", "ggu", "ruwr", "rrbb", "wwr", "gg", "bwrbuw", "ubu", "rrr", "ruubgg", "wgbrr", "wg", "grgb", "gubw", "buubg", "buur", "bwuw", "wugwwuug", "rwgbw", "guwg", "bbbw", "brrww", "rgbgb", "rubr", "bbbu", "wwbbggr", "bggbuww", "bwugwgu", "rbubr", "grbbu", "grubgu", "bwg", "wgr", "wwbww", "wu", "bgbrw", "wug", "wwrw", "ugb", "uwgr", "rr", "gbggr", "gwrr", "grrrw", "urg", "rbg", "rgw", "buuwuw", "uwgbr", "wrbu", "wuggb", "ruwu", "wggwb", "wwubbb", "uwwgbwu", "wgg", "wrbuwgbb", "ugub", "wbu", "wbuwwgu", "wrwuwbuw", "ugwgbbuu", "bwgb", "wgw", "rgbwrrbb", "uur", "brug", "wrbrg", "ugw", "rru", "rugru", "rggg", "rbrgwr", "guur", "grwugrr", "gugu", "gubg", "uubbu", "gbwugb", "rw", "uuburg", "gwwbu", "grg", "gbu", "bbwbrbg", "ubr", "wur", "rrubggu", "bru", "gwbb", "bbr", "bwbbrg", "brrbggbr", "wbbgw", "brw", "guu", "rbww", "bwb", "wgurrbu", "bgrwbr", "bbruwu", "rrugg", "rb", "bwr", "grrbbbbw", "bwbggwwg", "wbg", "bub", "gwrrb", "ubg", "grrwgu", "brrwru", "bruwg", "brr", "rwbwrbg", "urr", "ggwbu", "wgbrbgb", "gwuw", "wugwb", "rrgbww", "grgug", "gwugg", "gu", "ubuuuu", "wgrgbr", "bbu", "ggg", "bguwb", "ggrbrubr", "guuu", "wggrwwuw", "ubrwgbr", "bugwu", "wwug", "gb", "wwwbg", "ru", "uwb", "rrbuuu", "buugw", "rubwb", "rrb", "uub", "rbubb", "ggbrb", "brgw", "bww", "wwbr", "guw", "bgur", "ggwg", "wurbr", "wbubbr", "ubub", "gwb", "bgwu", "wwb", "wurwuuu", "gbw", "ubrbur", "uuuw", "rbur", "bgwg", "grgbbbuu", "ubw", "rgbg", "ugu", "bwgwg", "wwggbr", "rguw", "wub", "wwru", "bwrrwgu", "uuu", "uurrwwr", "bguw", "grru", "buww", "bbggrb", "ubgb", "brbrurw", "ub", "brgb", "bwbwuwg", "brub", "wwbw", "wwbwwrg", "www", "rwbubb", "wbrugggw", "grbug", "ruu", "urubu", "gbbbrwu", "bbugg", "w", "gwwrb", "bgruuwrg", "wugbg"}
		sort.Slice(towels, func(i, j int) bool {
			return len(towels[i]) > len(towels[j])
		})
		memo := map[string]int{}
		assert.Zero(t, countPossibilities(memo, towels, "wwggwbwruubbwurrgrggbuuwwwgwbrwubggurbrwugguruggggbrwggbbrg"))
	})
}
