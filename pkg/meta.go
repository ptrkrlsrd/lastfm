package lastfm

// Bio ...
type Bio struct {
	Links     Links  `json:"links,omitempty"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

// Stats ...
type Stats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}
