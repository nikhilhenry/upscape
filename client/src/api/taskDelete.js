import axios from 'axios'

export default async function(task_id){
  const response = await axios.delete(`api/task/${task_id}`)
    .then(res=>{
      return res.data
    })
    .catch(err=>{
      return err.response.data
    })

    return response
}