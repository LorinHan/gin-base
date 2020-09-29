package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
 * @description: 将字符串转换为uint类型
 * @author: Lorin
 * @time: 2020/9/29 上午11:52
 */
func StrToUint(str string) uint {
	intNum, _ := strconv.Atoi(str)
	uintNum := uint(intNum)
	return uintNum
}

/**
 * @description: 将手机号中间4位隐藏为星号
 * @params: phone -> 手机号码
 * @return: 处理后的手机号码
 * @author: Lorin
 * @time: 2020/9/29 上午11:52
 */
func EncryptedPhone(phone string) string {
	old := ""
	for k, v := range phone {
		if k >= 3 && k <= 6 {
			old = old + string(v)
		}
	}
	return strings.Replace(phone, old, "****", -1)
}

/**
 * @description: 生成随机字符串
 * @params: l -> 随机字符串的长度
 * @return: 指定长度的随机字符串
 * @author: Lorin
 * @time: 2020/9/29 上午11:53
 */
func RandomStr(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * @description: 对map数组进行排序
 * @description: Key：指定为map结构中被排序的字段
 * @description: MapList：被排序的map数组
 * @description: 调用sort进行排序 -> sort.Sort(&mapsSort)
 * @author: Lorin
 * @time: 2020/9/29 上午11:54
 */
type MapsSort struct {
	Key     string
	MapList []map[string]interface{}
}

func (m *MapsSort) Len() int {
	return len(m.MapList)
}

func (m *MapsSort) Less(i, j int) bool {
	var ivalue float64
	var jvalue float64
	var err error
	switch m.MapList[i][m.Key].(type) {
	case string:
		ivalue, err = strconv.ParseFloat(m.MapList[i][m.Key].(string), 64)
		if err != nil {
			return true
		}
	case int:
		ivalue = float64(m.MapList[i][m.Key].(int))
	case float64:
		ivalue = m.MapList[i][m.Key].(float64)
	case int64:
		ivalue = float64(m.MapList[i][m.Key].(int64))
	}
	switch m.MapList[j][m.Key].(type) {
	case string:
		jvalue, err = strconv.ParseFloat(m.MapList[j][m.Key].(string), 64)
		if err != nil {
			return true
		}
	case int:
		jvalue = float64(m.MapList[j][m.Key].(int))
	case float64:
		jvalue = m.MapList[j][m.Key].(float64)
	case int64:
		jvalue = float64(m.MapList[j][m.Key].(int64))
	}
	return ivalue < jvalue
}

func (m *MapsSort) Swap(i, j int) {
	m.MapList[i], m.MapList[j] = m.MapList[j], m.MapList[i]
}
