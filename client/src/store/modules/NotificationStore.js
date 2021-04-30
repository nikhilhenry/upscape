export default{
  namespaced:true,

  state:{
    notification:{},
  },

  mutations:{
    STORE_NOTIFICATION:(state,newNotification)=>{
      state.notification = newNotification
    },
    DELETE_TAG:(state)=>{
      state.notifications = {}
    },
  },

  actions:{
    storeNotification:(context,newTag)=>{
      if(newTag) context.commit("STORE_NOTIFICATION",newNotification)
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