import {axios} from '../axios/axios'

export const getToken = async (payload) => {
   return await axios.post(`${process.env.VUE_APP_URL_API}/token`,payload)
   .then(({data}) => {
        return data;
   })
   .catch((err) => {
       return err;
   })
}

export const searchEmailsByTerm = async (payload) => {
    let token = localStorage.getItem("token");
    return await axios.get(`${process.env.VUE_APP_URL_API}/search/${payload.term}-${payload.from}-${payload.maxResults}`,{
        headers: {
                    Authorization: `Bearer ${token}`,
            }
        })
    .then(({data}) => {
         return data;
    })
    .catch((err) => {
        return err;
    })
 }

 