package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func main() {
	// log.Println(DuplicateEncode("Supralapsarian"))
	// log.Println(findSubstring("aqkgwyolaososilnvnquimtcawuwtqrigjnxfjtcwvuoxgtzfehciaocgsossobzyrmxjutsudkrxhzwtmsythfecwdmdznwuopelbqjyuqtkymglfdlnlwhkpkoovqkwjwfcjnaannjuynslpkcrzeqgxnanunxpbqbzxvuozahizdsxflvysenhkisrnwtggsoirefeshuhjjooxpuoisicpncrnqvirkfjfeczyynvpdmweexheegpfitcnromuibrmfvvlkqeglsuxcyobplyelhwjfcbnnkapkyommzhqwsrjylepywakzfzhvxghfmnckkxcdagkmffhrnfklinamnrgicdzzwtvbordqvkomjespinaqrqhnpngymosaucbotyimoyofzkvrudcipafbykwshygeeexiefqczuhldwikewhunxemxininwftuvpnoyjirgqmjvvwlydyrdxqxglsbyxgeafvwrznbljpjtbrjubyyrjjkjqjcfziaekqsevbzsiebjhdemivfmwuyfrayvqfvmvlvdjurwrrtubhtiqwzyugpgcfreybdyirtitvkoymaotyrzpftvyyacbwohgrhcpwuqxjjjjzipsmeqtumkvummqfsdysilplbjzggkbwohoayzzapfjwnlzchwyymptxhrvalnoklboawyoecosrkbwjngbfdzfamfiqfwyancpatfccxpflvlirzgwygedqxcgnxcusunjtlyscbwpbqvvuexialqevkcwzduvkuuyofjmjalctxfgemxbfyvfudqketktzyrxcbroefhjukjmynohrvvsajempoiunxbztjualnpcdznfxsrizgtfbdkcrmtvxbgetcgucsrjpzjbtufbfnaxkipqkisbchzymbsqymjxjuojbxoljpbcgwkycvxymoawzopjuaypkgwwpxggjakhpfqbsjaznatalnznezsncyqgflwtdphpnfjjxgykupkbqdwpcnuncvyjsep{-truncated-}{-truncated-}", 10))
	RemainderSorting([]string{"Colorado", "Utah", "Montana", "Wisconsin", "Oregon", "Maine"})
	EncodeManager(&Manager{"Bekzhan", "SEO", 27, 1})
}

type T struct {
	Key   string
	Value int
	Len   int
}

func RemainderSorting(strArr []string) []string {
	result := []string{}
	type T struct {
		Word   string
		Modulo int
		Len    int
	}
	seqStruct := []T{}

	// m[word] = []int{len(word) % 3, len(word)}
	for _, word := range strArr {
		seqStruct = append(seqStruct, T{word, len(word) % 3, len(word)})
	}
	//two step sort
	sort.Slice(seqStruct, func(i, j int) bool {
		if seqStruct[i].Modulo == seqStruct[j].Modulo {
			return seqStruct[i].Len > seqStruct[j].Len
		} else {
			return seqStruct[i].Modulo < seqStruct[j].Modulo
		}
	})

	for _, obj := range seqStruct {
		result = append(result, obj.Word)
	}
	log.Println(result)
	return result
}

// (io.Reader, error)
func EncodeManager(manager *Manager) {
	result := "{"
	if manager.FullName != "" {
		result += `"full_name:"` + manager.FullName + ","
	}
	if manager.Position != "" {
		result += `"position:"` + manager.Position + ","
	}
	if manager.Age != 0 {
		result += `"age:"` + strconv.Itoa(int(manager.Age)) + ","
	}
	if manager.YearsInCompany != 0 {
		result += `"full_year_company:"` + strconv.Itoa(int(manager.YearsInCompany))
	}
	result += "}"
	log.Println(result)
	// return io.Reader, nil
}

func parallelProccessing() {

}

func findSubstring(s string, k int32) string {
	maxCountVowel := 0
	indexStart := 0
	kInt := int(k)

	for i := range s {
		if len(s)-i >= kInt {
			// count :=
			count := countVowel(s[i : kInt+i]) // get count, each slice word
			if maxCountVowel <= count {
				indexStart = i
				maxCountVowel = count
			}
			if maxCountVowel == kInt {
				indexStart = i
				break
			}
		}
	}

	log.Println(maxCountVowel)
	if maxCountVowel == 0 {
		return "Not found!"
	}
	return s[indexStart : indexStart+kInt]
}

func countVowel(slice string) int {
	// vowels := make(map[rune]bool, 5)
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for _, char := range slice {
		for _, vowel := range vowels {
			if vowel == char {
				count++
			}
		}
	}
	return count // return count
}

// for _, c := range word { t[c]++ }
// for _, c := range word {

func DuplicateEncode(word string) string {

	m := make(map[rune][]int, 10)
	word = strings.ToLower(word)
	result := ""

	for i, ch := range word {
		if _, ok := m[ch]; ok {
			arr := m[ch]         //[0,2]  get prev value
			arr = append(arr, i) // append, prev and current
			m[ch] = arr          // rewrite data
			// obj.Duplicate[ch] = append(obj.Duplicate[ch], i)
		} else {
			m[ch] = []int{i} //map[w]=[0,]
		}
	}
	// log.Println(m)
	//compare len & value in array indexes == char index

	//case same letter
	if len(m[rune(word[0])]) == len(word) {
		for i := 0; i < len(word); i++ {
			result += ")"
		}
	} else {

		for idxChar, char := range word { //0,1,2,3,4,5..
			for key, indexs := range m { //[2,3], [4]
				if key == char {
					if len(indexs) == 1 && indexs[0] == idxChar {
						result += "("
					}
					if len(indexs) == 2 && (indexs[0] == idxChar || indexs[1] == idxChar) {
						result += ")"
					}
					if len(indexs) == 3 && (indexs[0] == idxChar || indexs[1] == idxChar || indexs[2] == idxChar) {
						result += ")"
					}
					if len(indexs) == 4 && (indexs[0] == idxChar || indexs[1] == idxChar || indexs[2] == idxChar || indexs[3] == idxChar) {
						result += ")"
					}
					// if ok := helper(len(indexs), indexs, idxChar); ok {
					// 	result += ")"
					// } else {
				}
			}
		}
	}

	return result
}
func helper(n int, seq []int, currentCharIdx int) bool {

	for i := 0; i < n; i++ { //< 3
		if seq[i] == currentCharIdx {
			return true
		}
	}
	return false
}
