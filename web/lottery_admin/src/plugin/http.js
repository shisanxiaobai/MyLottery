import Vue from "vue"
import axios from "axios"

let Url='http://localhost:9002/'

axios.defaults.baseURL=Url

axios.interceptors.request.use(config => {
    config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    return config
  })
  

Vue.prototype.$http=axios

export {Url}