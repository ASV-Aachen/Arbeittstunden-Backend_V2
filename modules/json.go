package modules

import "time"

type JsonMemberRolle struct {
	Rolle string
}

type JsonReduction struct {
	Season               int
	Reduction_in_percent float32
	Note                 string
}

type JsonProject struct {
	Name         string
	Description  string
	First_season int
}

type Json_Member_item_hour struct {
	Member   string
	Duration int
}

type Json_Project_Item struct {
	Season      int
	Date        time.Time
	Title       string
	Description string
	Approved    bool
	Countable   bool
	Work        []Json_Member_item_hour
}
