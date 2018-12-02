package main

import (
	"fmt"
	"strings"
)

// A structure to store a pattern and
// which character is missing
type reduced struct {
	index   int
	pattern string
}

var lines = strings.Split(data, "\n")

func main() {
	fmt.Println("checksum", partA())
	fmt.Println("id (1)", partB())
	fmt.Println("id (2)", partB2())
}

// Calculate checksum (product) of the number of boxes with
// as at least two and three repeated characters
func partA() int {
	checksumTwo := 0
	checksumThree := 0

	llen := len(lines[0])

	for _, line := range lines {

		// Count the occurance of each char
		b := make([]byte, llen)
		for i := 0; i < llen; i++ {
			b[line[i]-'a']++
		}

		// Find how many of the chars occurs
		// two or three times
		twos := 0
		threes := 0
		for i := 0; i < llen; i++ {
			if b[i] == 2 {
				twos++
			} else if b[i] == 3 {
				threes++
			}
		}
		if twos > 0 {
			checksumTwo++
		}
		if threes > 0 {
			checksumThree++
		}
	}
	return checksumTwo * checksumThree
}

// First try by looping over each line
// and then all lines after it
func partB() string {
	cLines := len(lines)
	lLen := len(lines[0])

	for i, line := range lines {
	outer:
		for j := i + 1; j < cLines; j++ {
			otherLine := lines[j]

			cDiff := 0
			for k := 0; k < lLen; k++ {
				if line[k] != otherLine[k] {
					cDiff++
					if cDiff > 1 {
						continue outer
					}
				}
			}
			// Found two lines with one differing character
			// return string of all characters which are in common
			common := make([]byte, 0, lLen-1)
			for k := 0; k < lLen; k++ {
				if line[k] == otherLine[k] {
					common = append(common, line[k])
				}
			}
			return string(common)
		}
	}
	return ""
}

// Second try by storing all lines and which character is
// removed in a map and check against this. Fewer iterations
// but much more storage allocated
func partB2() string {
	lLen := len(lines[0])
	patterns := make(map[reduced]bool)

	for _, line := range lines {
		for j := 0; j < lLen; j++ {
			pattern := line[:j] + line[j+1:]
			r := reduced{index: j, pattern: pattern}
			if _, ok := patterns[r]; ok {
				return r.pattern
			}
			patterns[r] = true
		}
	}
	return ""
}

var data = `ohvflkatysoimjxbunazgwcdpr
ohoflkctysmiqjxbufezgwcdpr
ohvflkatysciqwxfunezgwcdpr
fhvflyatysmiqjxbunazgwcdpr
ohvhlkatysmiqjxbunhzgwcdxr
ohvflbatykmiqjxbunezgscdpr
ohvflkatasaiqjxbbnezgwcdpr
ohvflkatyymiqjxrunetgwcdpr
ohvflkatbsmiqhxbunezgwcdpw
oheflkytysmiqjxbuntzgwcdpr
ohvflkatrsmiqjibunezgwcupr
ohvflkaiysmiqjxbunkzgwkdpr
ohvilkutysmiqjxbuoezgwcdpr
phvflkatysmkqjxbulezgwcdpr
ohvflkatnsmiqjxbznezgpcdpr
ohvylkatysriqjobunezgwcdpr
ohvflkatytmiqjxbunezrwcypr
ohvonkatysmiqjxbunezgwxdpr
ohvflkatgsmoqjxyunezgwcdpr
ohvflkbtqsmicjxbunezgwcdpr
ohvflkatysmgqjqbunezgwcdvr
ohvtlkatyrmiqjxbunezgwcdpi
ohvflkatyskovjxbunezgwcdpr
ohvflkayysmipjxbunezgwcdpu
ohvalkltysmiqjxbunezgecdpr
ohvflkatysmiqjxiunezgnndpr
ohvflkatyomiqjxbbnezgwcdpp
ohvflkatysmiqjxbuoezgncdpy
omvflkvtysmiqjxwunezgwcdpr
ohvflkatynmicjxbunezgwpdpr
ohvflkatyqmaqjxbunezvwcdpr
ohbfhkatysmiqjxbunezgwcdqr
ohvflkatesmiqjvbunezpwcdpr
ohvflkatysmsqjxiunezgwcdhr
ohvfjkatysmwqjxbunezgwcddr
ohvflkanysmiqjxbunwkgwcdpr
ohqflkatysmiqjxbuuezgwcddr
ohvflkatysmvqjxbznlzgwcdpr
ohvflkatysmiqjxbunjzwwqdpr
ohvfjkatysmiqxxbunezgwcupr
chvfxkatysmiqjxxunezgwcdpr
uhvflkatitmiqjxbunezgwcdpr
ohvflbatysmiqjxbuntzgwcdor
ohvflkmtysmmqjxbunexgwcdpr
ohvflsatysmyqjxjunezgwcdpr
ohvfskatysmiqjjbunezgwcdpg
ohvflkatysniqjxbunexgwcrpr
ohvfekatysmiqjxbunedswcdpr
ohvfltatysmjqjxbunezghcdpr
ohvflkatydmiqjxvunezggcdpr
oavflkatysmiqjxtunazgwcdpr
ohvflkltysmiqjxbuzeugwcdpr
ohbflkatysmiqjybuuezgwcdpr
ehvfzkatysmiqjxbuhezgwcdpr
odvflkatssmiqjxbunezgwcdpj
ohvflkatysmiqjzbufezgwbdpr
jhvflkdtysmiqqxbunezgwcdpr
ohvflkatysmiqjwbunengwcnpr
ohvfskatysmiqjxbxuezgwcdpr
ohvflkatysmiqjobvnezgwcrpr
ohvrlkatysmiqjxbwnezgrcdpr
ofvflkatysmiqjxbunezpwcdwr
ohvfxdatyomiqjxbunezgwcdpr
yhvflkatydmiqjxbubezgwcdpr
ohvflkatysdiqjxbuneztwcspr
ohvflkatydmiquxbunezgwcbpr
ohvflkatysmiqcxbukezgwcdwr
ohvflkntasmiqjxbunezghcdpr
lhvflkatysmiqjxbunezqwckpr
ehifikatysmiqjxbunezgwcdpr
ohvflkatysmiqjcbutezgwcdpm
ohvflkatjssiqrxbunezgwcdpr
oyvflkavysmiqjxlunezgwcdpr
orvflkgtysmiqjxbukezgwcdpr
ihvflkatysmiqaxbunpzgwcdpr
ohvflkatusmiqjxbbnezgwchpr
ohvflkatysbiqjxvuneugwcdpr
ohvflkatysmiqjcbungzgwcwpr
ovvflkatysmidjxbunezgscdpr
ohvflqatysmiljxbunfzgwcdpr
ghvfokatysmiqjxbunqzgwcdpr
nxvflkatysmxqjxbunezgwcdpr
ohvflkatysmiqjxbexezgwrdpr
ohvfrkatysmhqjxbuntzgwcdpr
ohvflkvtysmiqjxocnezgwcdpr
ohvglkgtysmiqjxnunezgwcdpr
ohvflkatysmnqjxbunecgwqdpr
oyvflkatysgiqjxbcnezgwcdpr
ofvflkatysmiqjxbunfzgwcdpg
otvflkttysmiqjxbunezgwmdpr
ohvflkvtysmiqjbbunezgzcdpr
ahvflkatysyiqjxbunezvwcdpr
ohiflkatysmydjxbunezgwcdpr
ohvfwkatysmvqjxbunezwwcdpr
ohvflkatysbiqjxbunergwodpr
hhvsdkatysmiqjxbunezgwcdpr
ihvflkwtysmiqjxbunezgacdpr
ohvfljatysmiqcxbunuzgwcdpr
ohvflkatysqiqlwbunezgwcdpr
ohvflkauysmkqjxwunezgwcdpr
ohvflkatysmoqjqbunezgwodpr
ohvslkvtysmipjxbunezgwcdpr
olvflkatysmiujxbunezgwctpr
osvflxatysmiqjxbenezgwcdpr
orvflkhtysmiqjxbinezgwcdpr
ohcflkatystiqjxbunezbwcdpr
ohcflkatyfmifjxbunezgwcdpr
ohvflkatdsmiqjxbrnezgwcdpt
ohvflkatysmiqjxbwnqzawcdpr
oevflkakysmiqjxbunezgwcdpt
ofvflkatysmiqjxbunbqgwcdpr
ohvflkatysmdqjxbunefqwcdpr
ohvklkalysmiqjxbunezgwcepr
ocvflhatysmiqjxbunezzwcdpr
uhvflkatysmiqmxbunezgwcxpr
ohvflkatyshikjhbunezgwcdpr
lbvflkatysmoqjxbunezgwcdpr
ohvflkatssmuqjxbunezgscdpr
ohvflkatysmifyxbuvezgwcdpr
ohvfikatysmiqjxbunezgwfupr
ohvmlkaiysmiqjxqunezgwcdpr
ohvflkatysmiqjxiunpzgwcdpo
lhvflkatysmpqjxbenezgwcdpr
ohvflkatysmiqjobunengwczpr
ohoflkatysniqjxbunezgccdpr
ohvfxkatysmiqjgbunyzgwcdpr
ohvflkytysmiljxbubezgwcdpr
hhvsdkatysmiqjxjunezgwcdpr
ohvflkatysmiqjtuunezgwcdpt
ohvfdkxtysmiqjubunezgwcdpr
ohxflkatysmiyjxbunezgwcdhr
ohvflkatysmiqjibunezgwcppd
ohvflkatysmihjxbunezgwcdhj
ohvflkatysmiqjxronezgwcdvr
ofrflxatysmiqjxbunezgwcdpr
ohvwlkatysmiqjxounezgscdpr
ohvflkatcodiqjxbunezgwcdpr
oqvflkatysmiqjxbunebgwmdpr
ohvflmatysmisjxbunezqwcdpr
ovvflkatysmiqjxbuxezgwcdpe
ohvflkatysmdejxbuneztwcdpr
hhvflkathsmiqjxbwnezgwcdpr
ohkflkatlsmsqjxbunezgwcdpr
ohvflkktysmizjxhunezgwcdpr
ohzflkatysmiqjrbunezgwcdpj
ohuflwatysmiqjxbunezgwcdgr
ohvflkatysmiqvxmunpzgwcdpr
xhvflkwtysmiqjxbunezgwjdpr
whvflkatysmiqjxbunezgzcopr
ohvflkayysmiqjxuznezgwcdpr
khvflkasysmiqjxbunezgwcdpv
ohvflkatylmiqjxbpnozgwcdpr
ohvflkgtysziqjxbunezgwgdpr
ohvfljaiysmiqjxbuvezgwcdpr
ohvflkxtyslizjxbunezgwcdpr
ohzflkatysmiqjxbcnezgwcdar
ohvflkatysmiqjxbisecgwcdpr
shvflkatyjmiqjkbunezgwcdpr
mhvflkatysmiqjxvunezgwcdpk
ohfflkatysmiqjxbunczgwcppr
ohvflkatysmiqjkzunezgwcdpc
ohvflkatysmifjxbuneygwctpr
ohvflkatysmimjbbunezgwcdpe
ohvflkatjsciqjxbunezgwcdpa
ohvxlkatysmitjxbunezswcdpr
ohvslkatfsmiqjxbunezgwudpr
ohvflkatysmiqexbugezgwcdnr
onvflkatysmiqjxkunezgtcdpr
fhsflkalysmiqjxbunezgwcdpr
oyvflkatysmiqjobxnezgwcdpr
ohvflkatysmiqjxbunezswgdvr
phvflkatyymiqjxvunezgwcdpr
oivflzutysmiqjxbunezgwcdpr
ohvflkftysmiqjxbunezkwcopr
ohvflkatysmwnjxbunezgwcdpp
ohvflkatysmiqkxcunezgwndpr
phvklkatysmiqjhbunezgwcdpr
ohvflrawysmiqjxbunhzgwcdpr
ohvflkatysmiqjxbunecgwcdig
ohvflpakysmiqjxbunezgwrdpr
odvflkatykmiqjxbunezglcdpr
ohtflkatysiiqjxblnezgwcdpr
lhvfpkatysmiqjxbupezgwcdpr
ohvflkatdsmiqjpbunezgwcdps
ohvflkztysmiqjxvunezgwjdpr
ohvflbatysmxqoxbunezgwcdpr
ohvklkaigsmiqjxbunezgwcdpr
ohvfgkawysmiqjxbunezgwcdur
ohvflkatyskpqjlbunezgwcdpr
ohvflkatyqmiqjhbupezgwcdpr
ohqflkatysmiqjxzonezgwcdpr
ohxfnkatyymiqjxbunezgwcdpr
ohmflkatpsmiqjxbunezgwcdpw
ohvflkatysmiqjibnnewgwcdpr
vevflkatysmiqjxbunezgwcypr
ohvflkatydmiqwxbungzgwcdpr
ohsrlkatysmiqjxbcnezgwcdpr
ohvflkptyvmiqexbunezgwcdpr
opzflkatysmiqjxrunezgwcdpr
ohvflkitysmiqjxcunezgwcmpr
ohvflkatysmhhjxblnezgwcdpr
ohvflkatysfiqjxbunrzgwmdpr
ohvflkatyamibjxbunezgwcdpf
ohvflkalysmigjxbunezggcdpr
ohvflkatwsmisjxbunezgdcdpr
dhvflkatysmlqjxbunszgwcdpr
ohvflkatysmiqjxbueeygwcbpr
ohvflkatgsmiqjnbunezhwcdpr
svvflkatysmiqjxbunezgwckpr
opvflkatysmiqpxbufezgwcdpr
ohnvlkatysmiqjxbunezglcdpr
phvflkutysjiqjxbunezgwcdpr
ohvflabtysmiqjjbunezgwcdpr
ouvflkatysmiqjsbunezgwcdpk
osvflkatysmijjxbunezgwcypr
owvflkatysmiqjxbukxzgwcdpr
ohvfliatvsmiljxbunezgwcdpr
ohvflkatysmiqjxbumezbwtdpr
ohvflkatyfcicjxbunezgwcdpr
ohvflkatysmiqldbunezgfcdpr
oqvflkatysmiqixkunezgwcdpr
ohvflkatysmiqjxbulezgicdpe
ohvflkatysmiqjxbuniegwcdpl
ohvflkatysmiqjwbunbzgwcdhr
ohvflkatysmiqjdbunezgwwdkr
ohqflkytysmiqjxbunezgwcdpc
ohvflkatysmigjxbunezqwwdpr
ohvfloatysmiqjpbumezgwcdpr
ohvklkathkmiqjxbunezgwcdpr
ohvflkstjsmiqjxbunezgwctpr
ohvvlkatysmiqjxbunewgwcdir
ohnflkatysmiqjxbunszgwcdlr
ohvflkatysmnqjxbunezgxcdlr
ohvfrkatysmiqjxbonezgwcdor
ihvflkatysmiqjxbuneogwcxpr
ohvflkatysmiqjxbunecgwcccr
owvflkatysmivjxbunezgwjdpr
ohvflkgtysmiqjxbunczhwcdpr
ohyqlkatysmiqjxbunezgwcypr
ohvflkatysmiqjvbunezuwcdpw
ohvflkathsmiqmxbuoezgwcdpr
ehvjlkajysmiqjxbunezgwcdpr
ohvflkltysmiqjxblnezgwjdpr
oovflkvtfsmiqjxbunezgwcdpr
olvfzkatysmiqjxyunezgwcdpr
ohvflkatysqitjxbunezgncdpr
yhvflkatysmkqjxbunazgwcdpr
zlvolkatysmiqjxbunezgwcdpr
ohvflpatysmiqjxbunezgwcapb
ohvflkatysmuqjxbunezgfcdur`
