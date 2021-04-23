import axios from 'axios'

export default{
  namespaced:true,

  state:{
    auth_token:localStorage.getItem('auth_token') || '',
    isAuthenticated:false,
    profilePic:'',
    fullname:''
  },

  mutations:{
    STORE_AUTH_TOKEN:(state,userToken)=>{
      state.auth_token = userToken;
    },
    SET_AUTHENTICATED:(state,isAuthenticated)=>{
      state.isAuthenticated=isAuthenticated;
    },
    STORE_FULL_NAME:(state,userFullName)=>{
      state.fullname = userFullName
    },
    STORE_PROFILE_PIC:(state,userProfilePic)=>{
      state.profilePicUrl = userProfilePic
    },    
  },

  actions:{
    // store token
    storeAuthToken:(context,userToken)=>{

      // store token to local storage
      localStorage.setItem('auth_token',userToken);
      
      context.commit('STORE_AUTH_TOKEN',userToken);

     // set header globally
     axios.defaults.headers.common['auth_token'] = userToken

      // set authenticated status
      context.commit("SET_AUTHENTICATED",true);
    }
  },

  getters:{
    getAuthToken:state=>{
      return state.auth_token
    },
    getAuthenticatedState:state=>{
      return state.isAuthenticated
    },
    getFirstName:state=>{
      return state.fullname.split(" ")[0];
    },
    getProfilePicURL:state=>{
      return state.profilePicUrl
    },
  }
}