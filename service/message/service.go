package product

import (
	"context"
	"sort"
	"strings"

	"github.com/mahfuz110244/project1/entity"
)

func GetTextInfo(ctx context.Context, message string) ([]entity.Message, error) {
	wordArray := strings.Split(message, " ")
	mapWord := make(map[string]int)
	for _, v := range wordArray {
		mapWord[v] = mapWord[v] + 1
	}

	mapMessage := make([]entity.Message, 0)
	for key, value := range mapWord {
		instance := entity.Message{
			Word:       key,
			Occurrence: value,
		}
		mapMessage = append(mapMessage, instance)
	}
	sort.Slice(mapMessage, func(i, j int) bool {
		return mapMessage[i].Occurrence > mapMessage[j].Occurrence
	})
	if len(mapMessage) > 10 {
		return mapMessage[:10], nil
	}
	return mapMessage, nil
}
