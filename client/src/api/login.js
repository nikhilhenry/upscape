import axios from 'axios';

export default async function(password){

  let response
  await axios.post('/api/auth/login',password)
    .then(res=>{
      response = res.data
    })
    .catch(error=>{
      response = error.response.data
      throw error
    })

    return response
}