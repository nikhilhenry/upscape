import axios from 'axios'

export default async function(){
  const response = await axios.get('/api/objective')
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}

