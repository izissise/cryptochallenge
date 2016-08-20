package characterFrequency

import str "strings"
import "math"

// CharacterFrequency return a frequency score on the string
// passed as a parameter corresponding as its letter appering in english
func CharacterFrequency(binarystring string) int {
	freqsEnglish := map[byte]float32{'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253,
		'e': 12.702, 'f': 2.228, 'g': 2.015, 'h': 6.094, 'i': 6.966, 'j': 0.153, 'k': 0.772,
		'l': 4.025, 'm': 2.406, 'n': 6.749, 'o': 7.507, 'p': 1.929, 'q': 0.095, 'r': 5.987,
		's': 6.327, 't': 9.056, 'u': 2.758, 'v': 0.978, 'w': 2.360, 'x': 0.150, 'y': 1.974,
		'z': 0.074}

	freqs := make(map[byte]int)
	freqsPercent := make(map[byte]float32)

	for _, c := range binarystring {
		d := (str.ToLower(string(c)))[0]
		(freqs[d])++
	}

	for k := range freqsEnglish {
		freqsPercent[k] = 100 * float32(freqs[k]) / float32(len(binarystring))
	}

	frequencyScore := 0
	for k := range freqsPercent {
		frequencyScore += int(math.Abs(float64(freqsEnglish[k] - freqsPercent[k])))
	}
	return frequencyScore
}

func MapFrequency(str string) map[byte]float32 {
  freq := make(map[byte]float32)
  
  for i := 0; i < 26; i++ {
   pos := byte(i) + 'a';
   for j := 0; j < len(str); j++ {
     if (str[j] == byte(i)) {
       freq[pos] += 1.0
     }
   }
   freq[pos] /= float32(len(str));
  }
  return freq;
}