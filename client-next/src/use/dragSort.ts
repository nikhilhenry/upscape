export default function useDragSort(itemList: Object[]) {
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
    const itemOne = itemList[fromIndex];
    const itemTwo = itemList[toIndex];

    itemList[fromIndex] = itemTwo;
    itemList[toIndex] = itemOne;

    console.log(itemOne, itemTwo);
  }

  return {
    dragStart,
    dragOver,
    dragDrop,
    dragEnter,
    dragLeave,
  };
}
