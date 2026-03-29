/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	type Event struct {
		time int
		enter int
	}
	events := make([]Event, 0)
	for _, interval := range intervals {
		events = append(events, Event{interval.start, 1})
		events = append(events, Event{interval.end, -1})
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].enter < events[j].enter
		}
		return events[i].time < events[j].time
	})

	cur, ans := 0, 0
	for _, e := range events {
		if e.enter > 0 {
			cur++
			ans = max(ans, cur)
		} else {
			cur--
		}
	}
	return ans
}
