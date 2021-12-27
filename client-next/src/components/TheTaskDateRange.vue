<template>
  <div class="date-range">
    <div class="wrapper">
      <i class="las la-angle-left" @click="dateIndex--"></i>
      <div class="content">
        <h3 class="day">{{ formattedDate[0] }}</h3>
        <p>{{ formattedDate[1] }}</p>
      </div>
      <i
        class="las la-angle-right"
        @click="incrementDate"
        :class="{ 'is-disabled': isDisabled }"
      ></i>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref, watch } from "vue";
import getDate from "../functions/getDate";

export default defineComponent({
  setup(_, { emit }) {
    const dateIndex = ref(0);

    const formattedDate = computed(() => {
      return getDate(dateIndex.value) || [];
    });

    const incrementDate = () => {
      if (dateIndex.value > 0) return;
      dateIndex.value++;
    };

    watch(formattedDate, (queryDate) => {
      // send query date to parent component
      if (!queryDate) return null;
      emit("update-date-query", queryDate[2]);
    });

    const isDisabled = computed(() => {
      return dateIndex.value > 0 ? true : false;
    });

    return {
      dateIndex,
      formattedDate,
      incrementDate,
      isDisabled,
    };
  },
});
</script>

<style lang="scss" scoped>
.wrapper {
  display: flex;
  flex-direction: row;
  align-items: center;
  text-align: center;

  .las {
    font-size: 2.5rem;
  }

  .is-disabled {
    color: rgba(darken($text-primary, 20%), 0.5);
    cursor: not-allowed;
  }

  h3 {
    margin-left: 1.2rem;
    margin-right: 1.2rem;
    margin-bottom: 5px;
    margin-top: 0;
    font-size: 24px;
    text-transform: capitalize;
  }

  p {
    margin-top: 0;
    margin-bottom: 0;
    margin-left: 1rem;
    margin-right: 1rem;
    font-size: 18px;
    text-transform: uppercase;
  }
}
</style>
