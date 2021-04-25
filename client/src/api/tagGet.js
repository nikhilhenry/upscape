import axios from 'axios'

export default async function(){
  const response = await axios.get('/api/tag')
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}

