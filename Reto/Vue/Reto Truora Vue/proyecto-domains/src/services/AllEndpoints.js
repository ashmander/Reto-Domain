import axios from 'axios'
const url = "http://localhost:7777/endpoints"

export default {
    getEndpoints() {
        return axios.get(url);
    } 
}