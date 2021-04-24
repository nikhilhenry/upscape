import axios from 'axios'

export default async function(range='today'){
  const response = await axios.get(`api/task?range=${range}`)
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}

