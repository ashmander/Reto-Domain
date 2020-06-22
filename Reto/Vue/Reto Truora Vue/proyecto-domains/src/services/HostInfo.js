import axios from 'axios'
const url = "http://localhost:5555/domains/"

export default {
    getInfo(domain) {
        return axios.get(url+domain);
    } 
}