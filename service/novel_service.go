package service

import (
	"go-boot-starter/models"
	"strconv"
)

type NovelService struct {
	PageNo   int
	PageSize int
	Id       int
}

func (n *NovelService) ListNovelIndex() (int, []*models.NovelIndex, error) {
	novel_index_db := models.NovelIndex{}
	total, err := novel_index_db.Count("")
	if err != nil {
		return 0, nil, err
	}
	list, err := novel_index_db.List(0, 0, "")
	if err != nil {
		return 0, nil, err
	}
	start := (n.PageNo - 1) * n.PageSize
	end := n.PageNo * n.PageSize
	if end > total {
		end = total
	}
	return total, list[start:end], err
}

func (n *NovelService) GetNovelIndex() (*models.NovelIndex, error) {
	novel_index_db := models.NovelIndex{}

	list, err := novel_index_db.List(0, 0, "")
	if err != nil {
		return nil, err
	}
	for _, index := range list {
		if index.Id == strconv.Itoa(n.Id) {
			return index, nil
		}
	}
	return nil, err
}

func (n *NovelService) ListNovelContent() (int, []*models.NovelContent, error) {
	novel_index_db := models.NovelContent{}
	//total, err := novel_index_db.Count("")
	//if err != nil {
	//	return 0, nil, err
	//}
	list, err := novel_index_db.List(0)
	if err != nil {
		return 0, nil, err
	}
	//start := (n.PageNo - 1) * n.PageSize
	//end := n.PageNo * n.PageSize
	//if end > total {
	//	end = total
	//}
	return 0, list, err
}
