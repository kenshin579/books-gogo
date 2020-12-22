package set

import log "github.com/sirupsen/logrus"

func hasDupeRune1(s string) bool {
	runeSet := map[rune]bool{}

	for _, r := range s {
		log.Info("runeSet", runeSet)
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}
	return false
}

func hasDupeRune2(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		log.Info("runeSet", runeSet)
		if _, exists := runeSet[r]; exists {
			log.WithFields(log.Fields{
				"runeSet[r]": runeSet[r],
				"exists":     exists,
			}).Info("test")
			return true
		}

		runeSet[r] = struct{}{}
	}
	return false
}
