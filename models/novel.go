package models

import (
	"encoding/json"
	"os"
)

type NovelIndex struct {
	Id            string `json:"id"`
	Index         string `json:"index"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	Category      string `json:"category"`
	LocalChapter  string `json:"local_chapter"`
	RemoteChapter string `json:"remote_chapter"`
	CreateDate    string `json:"create_date"`
	ImgCover      string `json:"img_cover"`
	Desc          string `json:"desc"`
}

func (n *NovelIndex) List(a, b int, searchKey string) ([]*NovelIndex, error) {
	file, err := os.ReadFile("novel_index.json")
	if err != nil {
		return nil, err
	}
	var indexList []*NovelIndex
	if err = json.Unmarshal(file, &indexList); err != nil {
		return nil, err
	}
	return indexList, err
}

func (n *NovelIndex) Count(searchKey string) (int, error) {
	file, err := os.ReadFile("novel_index.json")
	if err != nil {
		return 0, err
	}
	var indexList []*NovelIndex
	if err = json.Unmarshal(file, &indexList); err != nil {
		return 0, err
	}
	return len(indexList), err
}

type NovelContent struct {
	Id          string `json:"id"`
	IndexId     string `json:"index_id"`
	ContentId   string `json:"content_id"`
	Title       string `json:"title"`
	EntityId    string `json:"entity_id"`
	UrlLink     string `json:"url_link"`
	VolumeTitle string `json:"volume_title"`
	BodyHtml    string `json:"bodyHtml"`
}

func (n *NovelContent) List(indexId int) ([]*NovelContent, error) {
	file, err := os.ReadFile("novel_content_5.json")
	if err != nil {
		return nil, err
	}
	var indexList []*NovelContent
	if err = json.Unmarshal(file, &indexList); err != nil {
		return nil, err
	}
	return indexList, err
}

func (n *NovelContent) GetByContentId(contentId string) (*NovelContent, error) {
	file, err := os.ReadFile("novel_content_5.json")
	if err != nil {
		return nil, err
	}
	var indexList []*NovelContent
	if err = json.Unmarshal(file, &indexList); err != nil {
		return nil, err
	}
	for _, content := range indexList {
		if content.ContentId == contentId {
			return content, nil
		}
	}
	return nil, err
}
