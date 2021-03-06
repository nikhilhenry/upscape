import Vue from 'vue'

import updateTaskById from '@/api/taskPut.js'

export default{
  namespaced:true,

  state:{
    tasks:[],
    tasksLoaded:false,
    totalDuration:false,
  },

  mutations:{
    STORE_TASKS:(state,taskList)=>{
      const tasks = state.tasks.concat(taskList);
      state.tasks = tasks
    },
    SET_TASKS:(state,taskList)=>{
      state.tasks = taskList
    },
    SET_TASKS_LOADED:(state,isLoaded)=>{
      state.tasksLoaded = isLoaded
    },
    UPDATE_TASK:(state,payload)=>{
      Vue.set(state.tasks,payload.index,payload.updatedTask)
    },
    STORE_TASK:(state,newTask)=>{
      state.tasks.push(newTask)
    },
    DELETE_TASK:(state,taskIndex)=>{
      state.tasks.splice(taskIndex,1)
    }
  },
  actions:{
    storeTasks:(context,userTasks)=>{
      context.commit("STORE_TASKS",userTasks)
      if(userTasks.length){
        context.commit("SET_TASKS_LOADED",true)
      }
    },

    updateTask:(context,updatedTask)=>{
      if(updatedTask){
        const index = context.state.tasks.findIndex(task=>task._id === updatedTask._id)
        context.commit("UPDATE_TASK",{index,updatedTask})
      }
    },

    storeTask:(context,newTask)=>{
      if(newTask){
        context.commit("STORE_TASK",newTask)
      }
    },

    deleteTask:(context,taskId)=>{
      if(taskId){
        const index = context.state.tasks.findIndex(task => task._id === taskId)
        context.commit("DELETE_TASK",index)
      }
    },
    setNewTaskList:(context,taskList)=>{
      context.commit("SET_TASKS_LOADED",false)
      context.commit("SET_TASKS",taskList)
    },
    updateTaskList:(context,taskList)=>{

      taskList.forEach(async (task,index)=>{

        // set task api id to index of ordered taskList
        task.id = index

        // send to api
        await updateTaskById(task._id,{id:task.id})
      })

      context.commit("SET_TASKS",taskList)

    }
  },

  getters:{
    getTasks:(state)=>{
      return state.tasks
    },
    getTask:(state) => (taskId)=>{
      const task = state.tasks.find(task => task._id === taskId)
      return task
    },
    getRemainingTime:(state)=>{
      const tasks = state.tasks

      if(!tasks){
        return 0
      }

      let totalTime = 0

      // filter tasks for uncompleted
      const uncompletedTasks = tasks.filter(task=>{return !task.completed})

      // loop through all uncompleted tasks and add time
      uncompletedTasks.forEach(task=>{

        totalTime+=task.duration
      })

      return (totalTime/60).toFixed(2)      
    },
    getTotalTime:(state)=>{
      const tasks = state.tasks

      if(!tasks){
        return 0
      }

      let totalTime = 0

      tasks.forEach(task=>{

        totalTime+=task.duration
      })
      // round time to 2 dp in hours
      return (totalTime/60).toFixed(2)
    }
  }
}