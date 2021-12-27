import { toRaw } from "vue";

export default function useDragSort(itemList: any, setter: Function) {
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
    const itemTwo = rawList[toIndex];

    rawList[fromIndex] = itemTwo;
    rawList[toIndex] = itemOne;

    itemList.value = rawList;

    console.log(itemOne, itemTwo);

    console.log(setter);
  }

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
