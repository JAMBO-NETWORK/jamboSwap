// const generalURL = 'http://192.168.1.132:8062';
const generalURL = "http://47.74.86.87:8062";
const hasAdvertisement = () => get(generalURL+'/v1/hasAdvertisement')//是否创建了广告
const saveFile = val =>upload(generalURL+'/v1/saveFile',val)//保存图片
const saveAdvertisement = val =>post(generalURL+'/v1/saveAdvertisement',val)//保存广告

const hideAdvertisement = val =>post(generalURL+'/v1/hideAdvertisement',val)//隐藏广告
const showAdvertisement = val =>post(generalURL+'/v1/openAdvertisement',val)//显示广告
const updateSortNum = val =>post(generalURL+'/v1/updateSortNum',val)//修改序号

const removeAdvertisement = val =>post(generalURL+'/v1/removeAdvertisement',val)//移除广告
const getAdvlist = val =>get(generalURL+'/v1/list',val)//获取广告列表
const getTotal = val =>get(generalURL+'/v1/getTotalTvlAndJamPrice',val)//获取总质押量和jamPrice
const liquidityInfo = val =>get(generalURL+'/v1/liquidityInfo',val)//	获取流动池信息
