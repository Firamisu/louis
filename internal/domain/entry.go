package domain

import "github.com/firamisu/louis/internal/dictclient"

type Entry struct {
	Word     string    `json:"word"`
	Phonetic string    `json:"phonetic"`
	Origin   string    `json:"origin"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}

func FromDictResponse(res dictclient.DictResponse) Entry {
	e := Entry{
		Word:     res.Word,
		Phonetic: res.Phonetic,
		Origin:   res.Origin,
	}

	if len(res.Meanings) != 0 {
		for _, m := range res.Meanings {
			mappedMeaning := Meaning{
				PartOfSpeech: m.PartOfSpeech,
			}

			if len(m.Definitions) != 0 {
				for _, d := range m.Definitions {
					mappedMeaning.Definitions = append(mappedMeaning.Definitions, Definition{
						Definition: d.Definition,
						Example:    d.Example,
					})
				}
			}

			e.Meanings = append(e.Meanings, mappedMeaning)

		}
	}

	return e
}
