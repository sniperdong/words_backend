package biz

type AddWordsRequest struct {
	Word              string `json:"spell"`
	Means             string `json:"means"`
	Pronounce         string `json:"pronounce"`
	Sentences         string `json:"sentences"`
	Property          int    `json:"property"`
	Plural            string `json:"plural"`
	PastTense         string `json:"past_tense"`
	PastParticiple    string `json:"past_participle"`
	PresentParticiple string `json:"present_participle"`
}
