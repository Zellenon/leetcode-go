package p3

func mapmax(m map[string]int) int {
	buffer := 0;
	for _, v := range m {
		if v > buffer {
			buffer = v;
		}
	}
	return buffer
}

func lengthOfLongestSubstring(s string) int {
	index := 0;
	longest := 0;
	max_len := len(s);
	char_set := make(map[string]int);

	for index + longest < max_len {
		next_letter := s[index+longest:index+longest+1];
		char_set[next_letter] += 1;
		if mapmax(char_set) > 1 {
			char_set[s[index:index+1]] -= 1;
			index += 1;
		} else {
			longest += 1;
		}
	}
	return longest;
}
