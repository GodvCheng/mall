package impl

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"mall/model"
	"mall/model/dto"
	"mall/server/dao"
	"reflect"
	"strconv"
	"time"
)

var goodsDao = dao.NewGoodsDao()

type GoodsServiceImpl struct {
}

func (g *GoodsServiceImpl) ListGoodsByCategory2(category2Id, current, pageSize int) (goodsList []*model.Goods, err error) {
	goodsList = goodsDao.ListGoodsByCategory2(category2Id, current, pageSize)
	if len(goodsList) == 0 {
		log.Errorln("根据二级分类查询商品列表为空")
		err = errors.New("请先添加商品哦")
	}
	return
}

func (g *GoodsServiceImpl) EditGoodsDetail(detail *model.GoodsDetail) (err error) {
	n := goodsDao.EditGoodsDetail(detail)
	if n == 0 {
		log.Errorln("修改商品详情失败")
		err = errors.New("修改失败，请再次尝试")
	}
	return nil
}

func (g *GoodsServiceImpl) DeleteGoodsDetail(id int) (err error) {
	n := goodsDao.DeleteGoodsDetail(id)
	if n == 0 {
		log.Errorln("删除商品详情失败")
		//err = errors.New("删除商品详情失败")
	}
	return
}

func (g *GoodsServiceImpl) AddGoodsDetail(detail *model.GoodsDetail) (err error) {
	n := goodsDao.CreateGoodsDetail(detail)
	if n == 0 {
		log.Errorln("添加商品详情失败")
		//err = errors.New("添加商品详情失败")
	}
	return nil
}

func (g *GoodsServiceImpl) DeleteCategory2(id int) (err error) {
	n := goodsDao.DeleteCategory2(id)
	if n == 0 {
		log.Errorln("删除二级分类失败")
		err = errors.New("删除失败，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) DeleteCategory1(id int) (err error) {
	n := goodsDao.DeleteCategory1(id)
	if n == 0 {
		log.Errorln("删除一级分类失败")
		err = errors.New("删除失败，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) ListCategory1() (category1List []*model.Category1, err error) {
	category1List = goodsDao.ListCategory1()
	for _, category1 := range category1List {
		category2List := goodsDao.ListCategory2(int(category1.ID))
		for _, category2 := range category2List {
			if category1.ID == category2.Category1Id {
				category1.Category2List = category2List
			}
		}
	}
	if len(category1List) == 0 {
		log.Errorln("商品一级分类为空")
		return nil, errors.New("请先添加一级分类")
	}
	return
}

func (g *GoodsServiceImpl) ListCategory2(category1Id int) (category2List []*model.Category2, err error) {
	category2List = goodsDao.ListCategory2(category1Id)
	if len(category2List) == 0 {
		log.Errorln("商品二级分类为空")
		return nil, errors.New("请先添加二级分类")
	}
	return
}

func (g *GoodsServiceImpl) Category1Info(id int) (category1 *model.Category1, err error) {
	category1 = goodsDao.Category1Info(id)
	if reflect.DeepEqual(category1, model.Category1{}) {
		log.Errorln("好像出错了，请再次尝试")
		return nil, errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) Category2Info(id int) (category2 *model.Category2, err error) {
	category2 = goodsDao.Category2Info(id)
	if reflect.DeepEqual(category2, model.Category2{}) {
		log.Errorln("获取商品二级分类详情失败")
		return nil, errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) UpdateCategory1(category1 *model.Category1) (err error) {
	n := goodsDao.UpdateCategory1(category1)
	if n == 0 {
		log.Errorln("好像出错了，请再次尝试")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) UpdateCategory2(category2 *model.Category2) (err error) {
	n := goodsDao.UpdateCategory2(category2)
	if n == 0 {
		log.Errorln("更新二级种类失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) CreateCategory1(category1 *model.Category1) (err error) {
	n := goodsDao.CreateCategory1(category1)
	if n == 0 {
		log.Errorln("创建一级分类失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) CreateCategory2(category2 *model.Category2) (err error) {
	var n int = goodsDao.CreateCategory2(category2)
	if n == 0 {
		log.Errorln("创建二级分类失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) EnableGoods(id int) (err error) {
	n := goodsDao.EnableGoods(id)
	if n == 0 {
		log.Errorln("商品上架失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return nil
}

func (g *GoodsServiceImpl) DisableGoods(id int) (err error) {
	n := goodsDao.DisableGoods(id)
	if n == 0 {
		log.Errorln("商品下架失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) ListGoodsImg(id int) (imgList []string, err error) {
	imgList = goodsDao.ListGoodsImg(id)
	if len(imgList) == 0 {
		log.Errorf("商品图片为空")
		return nil, errors.New("请先上传商品图片哦")
	}
	return imgList, nil
}

func (g *GoodsServiceImpl) DeleteGoods(id int) (err error) {
	num := goodsDao.DeleteGoods(id)
	if num == 0 {
		log.Errorln("删除商品失败")
		err = errors.New("好像出错了，请再次尝试")
		return
	}
	//删除商品详情
	err1 := g.DeleteGoodsDetail(id)
	if err1 != nil {
		err = err1
	}
	return
}

func (g *GoodsServiceImpl) ShowGoods(id int) (goods *dto.GoodsDto, err error) {
	goods = goodsDao.ShowGoods(id)
	if reflect.DeepEqual(*goods, dto.GoodsDto{}) {
		log.Errorln("查询商品信息失败")
		err = errors.New("好像出错了，请再次尝试")
	}
	return
}

func (g *GoodsServiceImpl) ListGoods(goods *model.Goods, current, pageSize int) (goodsList []*dto.GoodsDto, total int, err error) {
	goodsList, total = goodsDao.ListGoods(goods, current, pageSize)
	if len(goodsList) == 0 || total == 0 {
		log.Errorln("商品列表为空")
		return nil, 0, errors.New("请先添加商品哦")
	}
	return
}
func (g *GoodsServiceImpl) CreateGoods(goods *model.Goods) (err error) {
	goods.GoodsNo = strconv.Itoa(int(time.Now().Unix()))
	n := goodsDao.CreateGoods(goods)
	if n == 0 {
		log.Errorln("商品信息创建失败")
		err = errors.New("商品创建失败，请再次尝试")
		return
	}
	//添加商品详情
	var goodsDetail model.GoodsDetail
	goodsDetail.ID = goods.ID
	goodsDetail.Detail = goods.Detail
	err1 := g.AddGoodsDetail(&goodsDetail)
	if err1 != nil {
		err = err1
		return
	}
	return
}

func (g *GoodsServiceImpl) UpdateGoods(goods *model.Goods) (err error) {
	n := goodsDao.UpdateGoods(goods)
	if n == 0 {
		log.Errorln("商品信息更新失败")
		err = errors.New("好像出错了，请再次尝试")
		return
	}
	var goodsDetail model.GoodsDetail
	goodsDetail.ID = goods.ID
	goodsDetail.Detail = goods.Detail
	//修改商品详情
	err1 := g.EditGoodsDetail(&goodsDetail)
	if err1 != nil {
		err = err1
		return
	}
	return
}
