import Vue from 'vue'

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

    deleteReflection:(context,taskId)=>{
      if(taskId){
        const index = context.state.taks.findIndex(task => task._id === taskId)
        context.commit("DELETE_TASK",index)
      }
    }
  },

  getters:{
    getTasks:(state)=>{
      return state.tasks
    },
    getTask:(state) => (taskId)=>{
      const task = state.tasks.find(task => task._id === taskId)
      return task
    }
  }
}