import axios from 'axios'

export default async function(objective_id){
  const response = await axios.delete(`api/objective/${objective_id}`)
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}