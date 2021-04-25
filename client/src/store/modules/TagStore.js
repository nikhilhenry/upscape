export default{
  namespaced:true,

  state:{
    tags:[],
    tagsLoaded:false,
  },

  mutations:{
    STORE_TAGS:(state,tagList)=>{
      const tags = state.tags.concat(tagList)
      state.tags = tags
    },
    STORE_TAG:(state,newTag)=>{
      state.tags.push(newTag)
    },
    DELETE_TAG:(state,tagIndex)=>{
      state.tags.splice(tagIndex,1)
    },
    SET_TAGS_LOADED:(state,isLoaded)=>{
      state.tagsLoaded = isLoaded
    }
  },

  actions:{
    storeTags:(context,userTags)=>{
      context.commit("STORE_TAGS",userTags)
      if(userTags.length) context.commit("SET_TAGS_LOADED",true)
    },
    storeTag:(context,newTag)=>{
      if(newTag) context.commit("STORE_TAG",newTag)
    },
    deleteTag:(context,tagId)=>{
      if(tagId){
        const index = context.state.tasks.findIndex(tag => tag._id === tagId)
        context.commit("DELETE_TAG",index)
      }
    }
  },

  getters:{
    getTags:(state)=>{
      return state.tags
    }
  }
}