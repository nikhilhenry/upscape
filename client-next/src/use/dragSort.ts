import { computed, toRaw, WritableComputedRef } from "vue";

export default function useDragSort(itemList: WritableComputedRef<{}[]>) {
  let dragStartIndex: number, dragEndIndex: number;

  // list components
  function dragStart(event: any, index: number) {
    event.dataTransfer.dropEffect = "move";
    event.dataTransfer.effectAllowed = "move";

    dragStartIndex = index;
  }

  // li funcitons
  // 'dragover'
  function dragOver() {}
  // 'drop'
  function dragDrop(event: any, index: number) {
    dragEndIndex = index;
    swapItems(dragStartIndex, dragEndIndex);
  }
  // 'dragenter'
  function dragEnter() {}
  // 'dragleave'
  function dragLeave() {}

  function swapItems(fromIndex: number, toIndex: number) {
    const rawList = toRaw(itemList.value);
    const itemOne = rawList[fromIndex];

    rawList.splice(fromIndex, 1);

    const arrayStart = rawList.slice(0, toIndex - 1);
    const arrayEnd = rawList.slice(toIndex, rawList.length);

    arrayStart.push(itemOne);
    const computedArray = arrayStart.concat(arrayEnd);

    itemList.value = computedArray;
  }

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
