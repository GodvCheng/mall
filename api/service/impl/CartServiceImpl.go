package impl

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"mall/api/dao"
	"mall/api/util"
	"mall/model/dto"
	"strconv"
)

var redisCli = util.Rdb
var cartDao = dao.NewCartDao()

type CartServiceImpl struct {
}

func (c CartServiceImpl) DelFromCart(key, filed string) (err error) {
	err = redisCli.HDel(context.Background(), key, filed).Err()
	if err != nil {
		log.Errorf("删除购物项失败：%v", err)
		err = errors.New("好像出错了，请稍后再试")
	}
	return
}

func (c CartServiceImpl) ShowCartInfo(key string) (cartInfoList []*dto.CartDto, err error) {
	result, _ := redisCli.HGetAll(context.Background(), key).Result()
	mId, _ := strconv.Atoi(key)
	for goodsId, goodsNum := range result {
		cartInfo := cartDao.ShowCartInfo(goodsId)
		num, _ := strconv.Atoi(goodsNum)
		cartInfo.MemberId = uint(mId)
		cartInfo.Number = uint(num)
		cartInfoList = append(cartInfoList, cartInfo)
	}
	return
}

func (c CartServiceImpl) AddToCart(cart *dto.CartDto) error {
	key := strconv.Itoa(int(cart.MemberId))
	field := strconv.Itoa(int(cart.Id))
	value := int64(cart.Number)
	err := redisCli.HIncrBy(context.Background(), key, field, value).Err()
	if err != nil {
		log.Errorf("向购物车添加商品失败：%v", err)
		err = errors.New("好像出错了，请再试一次")
	}
	return err
}
