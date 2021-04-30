export default{
  namespaced:true,

  state:{
    notification:null,
  },

  mutations:{
    STORE_NOTIFICATION:(state,newNotification)=>{
      state.notification = newNotification
    },
    DELETE_NOTIFICATION:(state)=>{
      state.notification = null
    },
  },

  actions:{
    storeNotification:(context,newNotification)=>{
      if(newNotification) context.commit("STORE_NOTIFICATION",newNotification)

      // set timeout for notification
      setTimeout(()=>context.commit("DELETE_NOTIFICATION"),3000)
    },
    deleteNotification:(context)=>{
        context.commit("DELETE_NOTIFICATION")
    }
  },

  getters:{
    getNotification:(state)=>{
      return state.notification
    },   
  }
}