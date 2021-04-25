import axios from 'axios';

export default async function(payload){
  const reponse = await axios.post('api/tag',payload)
    .then(res=>{
      return res.data
    })
    .catch(err =>{
      return err.reponse.data
    })

    return reponse
}