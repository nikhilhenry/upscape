<template>
  <div class="date-range">
    <div class="wrapper">
      <i class="las la-angle-left" @click="dateIndex--"></i>
      <div class="content">
      <h3 class="day">{{formattedDate[0]}}</h3>
      <p>{{formattedDate[1]}}</p>
      </div>
      <i class="las la-angle-right" @click="incrementDate" :class="{'is-disabled':isDisabled}"></i>
    </div>
  </div>
</template>

<script>
import getDate from '@/functions/getDate'

export default {
  name:'TasksDateRange',
  computed:{
    formattedDate(){
      return getDate(this.dateIndex)
    },
    isDisabled(){
      // disabled when dateIndex is greater than 0
      return this.dateIndex > 0 ? true:false
    }
  },
  watch:{
    formattedDate: function(date){
      // send query date to parent component
      this.$emit('update:query',date[2])
    }
  },
  data(){
    return{
      dateIndex:0,
    }
  },
  methods:{
    incrementDate:function(){
      // limit range to 1
      if(this.dateIndex>=1) return
      this.dateIndex++
    }
  },
  mounted(){
    this.$emit('update:query','today')
  }
}
</script>

<style lang="scss" scoped>
.wrapper{
  display:flex;
  flex-direction: row;
  align-items: center;
  text-align: center;

  .las{
    font-size: 2.5rem;
  }

  .is-disabled{
    color:rgba(darken($text-primary,20%),.5);
    cursor:not-allowed;
  }

  h3{
    margin-left: 1.2rem;
    margin-right: 1.2rem;
    margin-bottom:5px;
    margin-top:0;
    font-size:24px;
    text-transform: capitalize;
  }

  p{
    margin-top:0;
    margin-bottom: 0;
    margin-left: 1rem;
    margin-right: 1rem;
    font-size: 18px;
    text-transform: uppercase;
  }
}
</style>