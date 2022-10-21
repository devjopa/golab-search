import {axios} from '../axios/axios'
/**
*
* @param {Object} payload
*
* get access token
*/
export const getToken = async (payload) => {
   return await axios.post('http://localhost:3100/api/token',payload)
   .then(({data}) => {
        return data;
   })
   .catch((err) => {
       return err;
   })
}

export const searchEmailsByTerm = async (payload) => {
    return await axios.get(`http://localhost:3100/api/search/${payload.term}-${payload.from}-${payload.maxResults}`)
    .then(({data}) => {
         return data;
    })
    .catch((err) => {
        return err;
    })
 }

 