const generalURL ='http://192.168.1.132:8062'
// const generalURL = 'http://119.8.103.132:8062';
const hasAdvertisement = () => get(generalURL+'/v1/hasAdvertisement')//是否创建了广告
const saveFile = val =>upload(generalURL+'/v1/saveFile',val)//保存图片
const saveAdvertisement = val =>post(generalURL+'/v1/saveAdvertisement',val)//保存广告

const removeAdvertisement = val =>post(generalURL+'/v1/removeAdvertisement',val)//移除广告
const getAdvlist = val =>get(generalURL+'/v1/list',val)//获取广告列表
const getTotal = val =>get(generalURL+'/v1/getTotalTvlAndJamPrice',val)//获取总质押量和jamPrice
const liquidityInfo = val =>get(generalURL+'/v1/liquidityInfo',val)//	获取流动池信息
