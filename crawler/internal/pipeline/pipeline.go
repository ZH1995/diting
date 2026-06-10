package pipeline

import "github.com/ZH1995/diting/crawler/internal/model"

type Pipeline interface {
	Process(data *model.HotItem) error
	Close() error
}
