package impl

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"mall/api/dao"
	"mall/model"
)

type GoodsServiceImpl struct {
}

func (g GoodsServiceImpl) GetCategory2Info(id int) (goodsInfo *model.Category2, err error) {
	goodsInfo, i := goodsDao.LoadCategory2Info(id)
	if i == 0 {
		log.Errorln("查询二级分类详情失败，为空")
		err = errors.New("好像出错了，请刷新试试")
	}
	return
}

func (g GoodsServiceImpl) GetGoodsInfo(id int) (goods *model.Goods, err error) {
	goods, n := goodsDao.LoadGoodsInfo(id)
	if n == 0 {
		log.Errorln("查询商品信息失败，商品信息为空")
		err = errors.New("好像出错了，请稍后再试")
	}
	return
}

func (g GoodsServiceImpl) GetGoodsListByCategory2Id(category2Id int) (goodsList []*model.Goods, err error) {
	goodsList = goodsDao.ListGoodsByCategory2Id(category2Id)
	if len(goodsList) == 0 {
		log.Errorln("查询商品列表失败，商品列表为空")
		err = errors.New("还没有该类商品哦，请购买其他种类商品")
	}
	return
}

func (g GoodsServiceImpl) GetBannerImg(goodsId int) (url []string, err error) {
	//TODO 将首页banner数据缓存到redis
	url = goodsDao.LoadBannerImg(goodsId)
	if len(url) == 0 {
		log.Errorf("获取banner图片失败")
		err = errors.New("好像出错了，请稍后再试")
	}
	return
}

var goodsDao = dao.NewGoodsDao()

func (g GoodsServiceImpl) GetCategories() (categories []*model.Category1, err error) {
	categories = goodsDao.LoadCategories()
	for _, category1 := range categories {
		category2List := goodsDao.ListCategory2(int(category1.ID))
		for _, category2 := range category2List {
			if category1.ID == category2.Category1Id {
				category1.Category2List = category2List
			}
		}
	}
	if len(categories) == 0 {
		log.Errorf("获取商品分类失败，商品分类为空")
		err = errors.New("好像出错了，请稍后再试")
	}
	return
}
