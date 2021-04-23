const getGreeting = function(){
  const hrs = new Date().getHours();
  if(hrs <12 ) return 'Good Morning'
  if(hrs >= 12 && hrs <= 17) return 'Good Afternoon'
  if(hrs >= 17 && hrs <= 24) return 'Good Evening'
  else return 'Welcome'
}

module.exports = getGreeting