package controllers

import (
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type SearchController struct {
	BaseController
}

func (s *SearchController) GetIndex() {
	userId := s.getUserId()
	var keywordService services.KeywordsService
	defaultKeyword := keywordService.GetDefalutKeyWorld()
	hotKeywords := keywordService.GetHotKeyWorldList()

	var serachHistory services.SearchHistoryService
	historyKeyworkList := serachHistory.GetUserSearchHistory(userId)
	utils.HttpSuccess(struct {
		DefaultKeyword     *models.NideshopKeywords        `json:"defaultKeyword"`
		HistoryKeyworkList []*models.NideshopSearchHistory `json:"historyKeyworkList"`
		HotKeywordList     []*models.NideshopKeywords      `json:"hotKeywordList"`
	}{
		defaultKeyword,
		historyKeyworkList,
		hotKeywords,
	}, s.Ctx)
}

//模糊查询关键字
func (s *SearchController) SearchHelper() {
	keyworld := s.GetString("keyword")
	var keyworldService services.KeywordsService
	list := keyworldService.FuzzyQueryKeywordList(keyworld)
	utils.HttpSuccess(list, s.Ctx)
}
