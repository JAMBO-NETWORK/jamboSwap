/** axios封装
 * 请求拦截、相应拦截、错误统一处理
 */
// 环境的切换
// axios.defaults.baseURL = 'http://192.168.1.105:6001/'
axios.defaults.timeout = 50000
// post请求头
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'


// 请求拦截器

axios.interceptors.request.use(config => {
  return config
})

function get(url, params) {
  return new Promise((resolve, reject) => {
    axios
      .get(url, {
        params: params
      })
      .then(res => {
        resolve(res.data)
      })
      .catch(err => {
        reject(err.msg)
      })
  })
}
// qs.stringify(data)
function post(url, data) {
	
	axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
	axios.defaults.headers.post['Accept'] = 'application/x-www-form-urlencoded;charset=UTF-8'
	// if(headering){
		// axios.defaults.headers.post['Content-Type'] = 'application/json, text/plain, */*'
	// }
  return new Promise((resolve, reject) => {
		console.log(data,Qs.stringify(data))
    axios
      .post(url,data)
      .then(res => {
        resolve(res.data)
      })
      .catch(err => {
        reject(err.msg)
      })
  })
}
// qs.stringify(data)
function upload(url, data) {
	axios.defaults.headers.post['Content-Type'] ='multipart/form-data'
	// axios.defaults.headers.post['process-Data'] = false
  return new Promise((resolve, reject) => {
    axios
      .post(url, data)
      .then(res => {
        resolve(res.data)
      })
      .catch(err => {
        reject(err.msg)
      })
  })
}