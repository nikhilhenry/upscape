import axios from 'axios';

export default async function(id,payload){
  const response = await axios.put(`api/objective/${id}`,payload)
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}