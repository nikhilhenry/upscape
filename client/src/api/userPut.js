import axios from 'axios';

export default async function(payload){
  const response = await axios.put('api/user',payload)
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}