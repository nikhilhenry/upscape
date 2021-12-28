<template>
  <div class="date-range">
    <div class="flex flex-row items-center text-center">
      <i class="las la-angle-left text-4xl" @click="dateIndex--"></i>
      <div class="content">
        <h3 class="mx-5 mb-2 text-2xl font-medium capitalize">
          {{ formattedDate[0] }}
        </h3>
        <p class="my-0 mx-4 text-xl uppercase">{{ formattedDate[1] }}</p>
      </div>
      <i
        class="las la-angle-right text-4xl"
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
.is-disabled {
  color: rgba(darken($text-primary, 20%), 0.5);
  cursor: not-allowed;
}
</style>
