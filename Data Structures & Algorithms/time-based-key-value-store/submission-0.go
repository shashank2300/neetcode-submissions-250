type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	timestamp int
	value string	
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.m[key]; !exists {
		return ""
	}

	pairs := this.m[key]
	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})
	if idx == 0 {
		return ""
	}

	return pairs[idx-1].value
}
