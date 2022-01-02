import { toRaw, WritableComputedRef } from "vue";

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

    rawList.splice(toIndex, 1, itemOne);

    itemList.value = rawList;
  }

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
