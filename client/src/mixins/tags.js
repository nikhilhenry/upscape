import getTags from '@/api/tagGet.js'
import addTag from '@/api/tagPost.js'

import  {mapActions, mapGetters} from 'vuex'

export default{

  computed:{
    ...mapGetters({
      tagOptions:'tag/getTags'
    })
  },
  async mounted(){
    // check if tags isLoaded
    if(!this.$store.state.tag.tagsLoaded){
      // query tags
      const tags = await getTags()

      // save tags to store
      this.$store.dispatch('tag/storeTags',tags)
    }
  },
  methods:{
    createTag: async function(newTagName){
      const randomColor = '#' + Math.floor(Math.random()*16777215).toString(16);

      // create tag 
      const tag = {name:newTagName,color:randomColor}

      // send tag to api
      const savedTag = await addTag(tag)

      // save tag to store
      this.storeTag(savedTag)


      // if used in multiselect
      if(this.selectedTags){
        this.selectedTags.push(savedTag)
      }
    },
    ...mapActions({
      storeTag:'tag/storeTag'
    })
  }
}