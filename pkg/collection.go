package pkg

type Gitmoji struct {
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type GitmojiCollection struct {
	Gitmojis []Gitmoji `json:"gitmojis"`
}

func LoadGitmojiCollection(personalSource string) (GitmojiCollection, error) {
	def, err := LoadDefaultGitmojiCollection()

	if err != nil {
		return GitmojiCollection{}, err
	}

	personal, err := LoadPersonalGitmojiCollection(personalSource)

	if err != nil {
		return GitmojiCollection{}, err
	}

	def.Gitmojis = append(def.Gitmojis, personal.Gitmojis...)

	return def, nil
}