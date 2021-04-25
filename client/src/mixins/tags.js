import getTags from '@/api/tagGet.js'
import  {mapGetters} from 'vuex'

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
  }
}