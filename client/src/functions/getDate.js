// when given date index (0,1,-1,2) the function must 
// return any array with 3 elements:
// 1. Day (ie wednesday or today or tomorrow)
// 2. Formated date (APR 24 2021)
// 3. Query date (24/04/21)


const getDate = function(dateIndex){
  if (typeof dateIndex !="number") return null

  const formattedDate  =[]

  // today with index 0
  if(dateIndex == 0){
    formattedDate[0] = 'today'
    formattedDate[1] = getFormattedDate(new Date())
    formattedDate[2] = getQueryDate(new Date())

    return formattedDate
  }

  // get date from index
  const date = new Date().setDate(new Date().getDate + dateIndex)
  
  formattedDate[0] = date.toLocaleString('default',{weekday:'long'})
  formattedDate[1] = getFormattedDate(date)
  formattedDate[2] = getQueryDate(date)

  return formattedDate
}

module.exports = getDate



function getFormattedDate(date){
  const formattedDate = date.toLocaleString('default',{month:'short'}).toUpperCase()+' '+new Date(date).getDate()+' '+ new Date(date).getFullYear()
  return formattedDate
}

function getQueryDate(date){
  const queryDate = new Date(date).getDate() +'/'+new Date(date).getMonth()+1+'/'+new Date(date).getFullYear();
  return queryDate
}